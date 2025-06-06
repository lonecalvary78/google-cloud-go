// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bttest

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"google.golang.org/grpc/metadata"

	btapb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	btpb "cloud.google.com/go/bigtable/apiv2/bigtablepb"
	"cloud.google.com/go/internal/testutil"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type closeSpyListener struct {
	net.Listener
	closed bool
}

func (s *closeSpyListener) Close() error {
	s.closed = true
	return s.Listener.Close()
}

func TestNewServer(t *testing.T) {
	t.Run("TCP", func(t *testing.T) {
		srv, err := NewServer("localhost:0")
		if err != nil {
			t.Fatalf("NewServer() error = %v", err)
		}
		if srv == nil {
			t.Fatal("NewServer() returned nil server")
		}
		defer srv.Close()
		if srv.Addr == "" {
			t.Error("NewServer() returned server with empty Addr")
		}
	})

	t.Run("Unix", func(t *testing.T) {
		tmpDir := t.TempDir()
		sockPath := tmpDir + "/bttest.sock"
		srv, err := NewServer(sockPath)
		if err != nil {
			t.Fatalf("NewServer() error = %v", err)
		}
		if srv == nil {
			t.Fatal("NewServer() returned nil server")
		}

		if srv.Addr != sockPath {
			t.Errorf("srv.Addr got %q, want %q", srv.Addr, sockPath)
		}

		srv.Close()

		// Check that the unix socket file was removed.
		if _, err := os.Stat(sockPath); !os.IsNotExist(err) {
			t.Errorf("socket file %q was not removed after Close()", sockPath)
		}
	})

	t.Run("WithListener", func(t *testing.T) {
		l, err := net.Listen("tcp", "localhost:0")
		if err != nil {
			t.Fatalf("net.Listen() error = %v", err)
		}

		// Create a spy that can tell us if its Close method was called.
		spy := &closeSpyListener{Listener: l}
		srv, err := NewServerWithListener(spy) // This should not call spy.Close().
		if err != nil {
			l.Close()
			t.Fatalf("NewServerWithListener() error = %v", err)
		}

		srv.Close() // This should NOT call the Close method on our spy.
		if spy.closed {
			t.Error("Listener was closed by the server, but it should not have been.")
		}

		// Clean up the real listener now that the test is done.
		l.Close()

		// Validate that the listener is now actually closed.
		_, err = l.Accept()
		if err == nil {
			t.Fatal("l.Accept() should have failed for a closed listener, but it did not")
		}
		if !errors.Is(err, net.ErrClosed) {
			t.Errorf("Expected net.ErrClosed, but got a different error: %v", err)
		}
	})
}

func TestConcurrentMutationsReadModifyAndGC(t *testing.T) {
	s := &server{
		tables: make(map[string]*table),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	if _, err := s.CreateTable(
		ctx,
		&btapb.CreateTableRequest{Parent: "cluster", TableId: "t"}); err != nil {
		t.Fatal(err)
	}
	const name = `cluster/tables/t`
	tbl := s.tables[name]
	req := &btapb.ModifyColumnFamiliesRequest{
		Name: name,
		Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{{
			Id:  "cf",
			Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Create{Create: &btapb.ColumnFamily{}},
		}},
	}
	_, err := s.ModifyColumnFamilies(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	req = &btapb.ModifyColumnFamiliesRequest{
		Name: name,
		Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{{
			Id: "cf",
			Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Update{Update: &btapb.ColumnFamily{
				GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}},
			}},
		}},
	}
	if _, err := s.ModifyColumnFamilies(ctx, req); err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	var ts int64
	ms := func() []*btpb.Mutation {
		return []*btpb.Mutation{{
			Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
				FamilyName:      "cf",
				ColumnQualifier: []byte(`col`),
				TimestampMicros: atomic.AddInt64(&ts, 1000),
			}},
		}}
	}

	rmw := func() *btpb.ReadModifyWriteRowRequest {
		return &btpb.ReadModifyWriteRowRequest{
			TableName: name,
			RowKey:    []byte(fmt.Sprint(rand.Intn(100))),
			Rules: []*btpb.ReadModifyWriteRule{{
				FamilyName:      "cf",
				ColumnQualifier: []byte("col"),
				Rule:            &btpb.ReadModifyWriteRule_IncrementAmount{IncrementAmount: 1},
			}},
		}
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ctx.Err() == nil {
				req := &btpb.MutateRowRequest{
					TableName: name,
					RowKey:    []byte(fmt.Sprint(rand.Intn(100))),
					Mutations: ms(),
				}
				if _, err := s.MutateRow(ctx, req); err != nil {
					panic(err) // can't use t.Fatal in goroutine
				}
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ctx.Err() == nil {
				_, _ = s.ReadModifyWriteRow(ctx, rmw())
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			tbl.gc()
		}()
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Error("Concurrent mutations and GCs haven't completed after 1s")
	}
}

func TestCreateTableResponse(t *testing.T) {
	// We need to ensure that invoking CreateTable returns
	// the  ColumnFamilies as well as Granularity.
	// See issue https://github.com/googleapis/google-cloud-go/issues/1512.
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()
	got, err := s.CreateTable(ctx, &btapb.CreateTableRequest{
		Parent:  "projects/issue-1512/instances/instance",
		TableId: "table",
		Table: &btapb.Table{
			ColumnFamilies: map[string]*btapb.ColumnFamily{
				"cf1": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 123}}},
				"cf2": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 456}}},
			},
		},
	})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}

	want := &btapb.Table{
		Name: "projects/issue-1512/instances/instance/tables/table",
		// If no Granularity was specified, we should get back "MILLIS".
		Granularity: btapb.Table_MILLIS,
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf1": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 123}}},
			"cf2": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 456}}},
		},
	}
	if diff := testutil.Diff(got, want); diff != "" {
		t.Fatalf("Response mismatch: got - want +\n%s", diff)
	}
}

func TestCreateTableWithFamily(t *testing.T) {
	// The Go client currently doesn't support creating a table with column families
	// in one operation but it is allowed by the API. This must still be supported by the
	// fake server so this test lives here instead of in the main bigtable
	// integration test.
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf1": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 123}}},
			"cf2": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 456}}},
		},
	}
	cTbl, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}
	tbl, err := s.GetTable(ctx, &btapb.GetTableRequest{Name: cTbl.Name})
	if err != nil {
		t.Fatalf("Getting table: %v", err)
	}
	cf := tbl.ColumnFamilies["cf1"]
	if cf == nil {
		t.Fatalf("Missing col family cf1")
	}
	if got, want := cf.GcRule.GetMaxNumVersions(), int32(123); got != want {
		t.Errorf("Invalid MaxNumVersions: wanted:%d, got:%d", want, got)
	}
	cf = tbl.ColumnFamilies["cf2"]
	if cf == nil {
		t.Fatalf("Missing col family cf2")
	}
	if got, want := cf.GcRule.GetMaxNumVersions(), int32(456); got != want {
		t.Errorf("Invalid MaxNumVersions: wanted:%d, got:%d", want, got)
	}
}

func TestGetPartitionsByTableName(t *testing.T) {
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf1": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 123}}},
			"cf2": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 456}}},
		},
	}
	_, err1 := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t1", Table: &newTbl})
	if err1 != nil {
		t.Fatalf("Creating table: %v", err1)
	}

	newTbl = btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf3": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 567}}},
			"cf4": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 890}}},
		},
	}
	_, err2 := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t2", Table: &newTbl})
	if err2 != nil {
		t.Fatalf("Creating table: %v", err2)
	}

	tblNamePrefix := "cluster" + "/tables/"

	// A random table name doesn't return partitions.
	partitions := s.GetPartitionsByTableName(tblNamePrefix + "random")
	if partitions != nil {
		t.Fatalf("Getting partitions for table random")
	}

	partitions = s.GetPartitionsByTableName(tblNamePrefix + "t1")
	if len(partitions) != 10 {
		t.Fatalf("Getting partitions for table t1")
	}

	partitions = s.GetPartitionsByTableName(tblNamePrefix + "t2")
	if len(partitions) != 10 {
		t.Fatalf("Getting partitions for table t2")
	}
}

type MockSampleRowKeysServer struct {
	responses []*btpb.SampleRowKeysResponse
	grpc.ServerStream
}

func (s *MockSampleRowKeysServer) Send(resp *btpb.SampleRowKeysResponse) error {
	s.responses = append(s.responses, resp)
	return nil
}

func TestSampleRowKeys(t *testing.T) {
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tbl, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}

	// Populate the table
	val := []byte("value")
	rowCount := 1000
	for i := 0; i < rowCount; i++ {
		req := &btpb.MutateRowRequest{
			TableName: tbl.Name,
			RowKey:    []byte("row-" + strconv.Itoa(i)),
			Mutations: []*btpb.Mutation{{
				Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
					FamilyName:      "cf",
					ColumnQualifier: []byte("col"),
					TimestampMicros: 1000,
					Value:           val,
				}},
			}},
		}
		if _, err := s.MutateRow(ctx, req); err != nil {
			t.Fatalf("Populating table: %v", err)
		}
	}

	mock := &MockSampleRowKeysServer{}
	if err := s.SampleRowKeys(&btpb.SampleRowKeysRequest{TableName: tbl.Name}, mock); err != nil {
		t.Errorf("SampleRowKeys error: %v", err)
	}
	if len(mock.responses) == 0 {
		t.Fatal("Response count: got 0, want > 0")
	}
	// Make sure the offset of the penultimate response is the offset of the final row
	got := mock.responses[len(mock.responses)-2].OffsetBytes
	want := int64((rowCount - 1) * len(val))
	if got != want {
		t.Errorf("Invalid penultimate offset: got %d, want %d", got, want)
	}

	// Make sure the offset of the final response is the offset of all the rows
	got = mock.responses[len(mock.responses)-1].OffsetBytes
	want = int64(rowCount * len(val))
	if got != want {
		t.Errorf("Invalid final offset: got %d, want %d", got, want)
	}

	// Make sure the key of the final response is empty
	gotLastKey := mock.responses[len(mock.responses)-1].RowKey
	wantLastKey := ""
	if got != want {
		t.Errorf("Invalid final RowKey: got %s, want %s", gotLastKey, wantLastKey)
	}
}

type AntagonistFunction func(s *server, attempts int, tblName string, finished chan (bool))

func SampleRowKeysConcurrentTest(t *testing.T, antagonist AntagonistFunction) {
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tbl, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}

	// Populate the table
	populate := func() {
		rowCount := 100
		for i := 0; i < rowCount; i++ {
			req := &btpb.MutateRowRequest{
				TableName: tbl.Name,
				RowKey:    []byte("row-" + strconv.Itoa(i)),
				Mutations: []*btpb.Mutation{{
					Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
						FamilyName:      "cf",
						ColumnQualifier: []byte("col"),
						TimestampMicros: 1000,
						Value:           []byte("value"),
					}},
				}},
			}
			if _, err := s.MutateRow(ctx, req); err != nil {
				t.Fatalf("Populating table: %v", err)
			}
		}
	}

	attempts := 500
	finished := make(chan bool)
	go func() {
		populate()
		mock := &MockSampleRowKeysServer{}
		for i := 0; i < attempts; i++ {
			if err := s.SampleRowKeys(&btpb.SampleRowKeysRequest{TableName: tbl.Name}, mock); err != nil {
				t.Errorf("SampleRowKeys error: %v", err)
			}
		}
		finished <- true
	}()
	go antagonist(s, attempts, tbl.Name, finished)
	for i := 0; i < 2; i++ {
		select {
		case <-finished:
		case <-time.After(2 * time.Second):
			t.Fatalf("Timeout waiting for task %d\n", i)
		}
	}
}

func TestSampleRowKeysVsDropRowRange(t *testing.T) {
	SampleRowKeysConcurrentTest(t, func(s *server, attempts int, tblName string, finished chan (bool)) {
		ctx := context.Background()
		for i := 0; i < attempts; i++ {
			req := &btapb.DropRowRangeRequest{
				Name:   tblName,
				Target: &btapb.DropRowRangeRequest_DeleteAllDataFromTable{DeleteAllDataFromTable: true},
			}
			if _, err := s.DropRowRange(ctx, req); err != nil {
				t.Fatalf("Dropping all rows: %v", err)
			}
		}
		finished <- true
	})
}

func TestSampleRowKeysVsModifyColumnFamilies(t *testing.T) {
	SampleRowKeysConcurrentTest(t, func(s *server, attempts int, tblName string, finished chan (bool)) {
		ctx := context.Background()
		for i := 0; i < attempts; i++ {
			req := &btapb.ModifyColumnFamiliesRequest{
				Name: tblName,
				Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{{
					Id:  "cf2",
					Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Create{Create: &btapb.ColumnFamily{}},
				}},
			}
			if _, err := s.ModifyColumnFamilies(ctx, req); err != nil {
				t.Fatalf("Creating column family cf2: %v", err)
			}
			rowCount := 100
			for i := 0; i < rowCount; i++ {
				req := &btpb.MutateRowRequest{
					TableName: tblName,
					RowKey:    []byte("row-" + strconv.Itoa(i)),
					Mutations: []*btpb.Mutation{{
						Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
							FamilyName:      "cf2",
							ColumnQualifier: []byte("col"),
							TimestampMicros: 1000,
							Value:           []byte("value"),
						}},
					}},
				}
				if _, err := s.MutateRow(ctx, req); err != nil {
					t.Fatalf("Populating table: %v", err)
				}
			}
			req = &btapb.ModifyColumnFamiliesRequest{
				Name: tblName,
				Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{{
					Id:  "cf2",
					Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Drop{Drop: true},
				}},
			}
			if _, err := s.ModifyColumnFamilies(ctx, req); err != nil {
				t.Fatalf("Dropping column family cf2: %v", err)
			}
		}
		finished <- true
	})
}

func TestModifyColumnFamilies(t *testing.T) {
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()
	tblInfo, err := populateTable(ctx, s)
	if err != nil {
		t.Fatal(err)
	}

	readRows := func(expectChunks, expectCols, expectFams int) {
		t.Helper()
		mock := &MockReadRowsServer{}
		req := &btpb.ReadRowsRequest{TableName: tblInfo.Name}
		if err := s.ReadRows(req, mock); err != nil {
			t.Fatalf("ReadRows error: %v", err)
		}
		cols := map[string]bool{}
		fams := map[string]bool{}
		chunks := 0
		for _, r := range mock.responses {
			for _, c := range r.Chunks {
				chunks++
				colName := c.FamilyName.Value + "." + string(c.Qualifier.Value)
				cols[colName] = true
				fams[c.FamilyName.Value] = true
			}
		}
		if got, want := len(fams), expectFams; got != want {
			t.Errorf("col count: got %d, want %d", got, want)
		}
		if got, want := len(cols), expectCols; got != want {
			t.Errorf("col count: got %d, want %d", got, want)
		}
		if got, want := chunks, expectChunks; got != want {
			t.Errorf("chunk count: got %d, want %d", got, want)
		}
	}

	readRows(27, 9, 3)

	// Now drop the middle column.
	if _, err := s.ModifyColumnFamilies(ctx, &btapb.ModifyColumnFamiliesRequest{
		Name: tblInfo.Name,
		Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{{
			Id:  "cf1",
			Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Drop{Drop: true},
		}},
	}); err != nil {
		t.Fatalf("ModifyColumnFamilies error: %v", err)
	}

	readRows(18, 6, 2)

	// adding the column back should not re-create the data.
	if _, err := s.ModifyColumnFamilies(ctx, &btapb.ModifyColumnFamiliesRequest{
		Name: tblInfo.Name,
		Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{{
			Id:  "cf1",
			Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Create{Create: &btapb.ColumnFamily{}},
		}},
	}); err != nil {
		t.Fatalf("ModifyColumnFamilies error: %v", err)
	}

	readRows(18, 6, 2)
}

func TestGC(t *testing.T) {
	// Create server
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()

	colFamilyID := "colFam"
	colName := "colName"
	rowKey := "rowKey"

	// Create table with max age gc rule
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{},
	}
	newTbl.ColumnFamilies[colFamilyID] = &btapb.ColumnFamily{GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxAge{MaxAge: durationpb.New(time.Millisecond)}}}

	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatal(err)
	}

	// Populate the table
	for i := 0; i < 2; i++ {
		req := &btpb.MutateRowRequest{
			TableName: tblInfo.Name,
			RowKey:    []byte(rowKey),
			Mutations: []*btpb.Mutation{{
				Mutation: &btpb.Mutation_SetCell_{
					SetCell: &btpb.Mutation_SetCell{
						FamilyName:      colFamilyID,
						ColumnQualifier: []byte(colName),
						TimestampMicros: 1000, // MaxAge is 1ms
						Value:           []byte{},
					}},
			}},
		}
		if _, err := s.MutateRow(ctx, req); err != nil {
			t.Fatal(err)
		}
	}
	if err != nil {
		t.Fatal(err)
	}

	// Sleep till maxAge passes
	time.Sleep(2 * time.Millisecond)

	// Trigger gc
	tbl := s.tables[tblInfo.Name]
	tbl.gc()

	// Verify that the row was deleted after garbage collection
	readRowsReq := &btpb.ReadRowsRequest{
		TableName: tblInfo.Name,
		Rows:      &btpb.RowSet{RowKeys: [][]byte{[]byte(rowKey)}},
	}
	mock := &MockReadRowsServer{}
	if err = s.ReadRows(readRowsReq, mock); err != nil {
		t.Errorf("ReadRows error: %v", err)
	}
	if got, want := len(mock.responses), 0; got != want {
		t.Errorf("response count: got %d, want %d", got, want)
	}
}

func TestDropRowRange(t *testing.T) {
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}

	tbl := s.tables[tblInfo.Name]

	// Populate the table
	prefixes := []string{"AAA", "BBB", "CCC", "DDD"}
	count := 3
	doWrite := func() {
		for _, prefix := range prefixes {
			for i := 0; i < count; i++ {
				req := &btpb.MutateRowRequest{
					TableName: tblInfo.Name,
					RowKey:    []byte(prefix + strconv.Itoa(i)),
					Mutations: []*btpb.Mutation{{
						Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
							FamilyName:      "cf",
							ColumnQualifier: []byte("col"),
							TimestampMicros: 1000,
							Value:           []byte{},
						}},
					}},
				}
				if _, err := s.MutateRow(ctx, req); err != nil {
					t.Fatalf("Populating table: %v", err)
				}
			}
		}
	}

	doWrite()
	tblSize := tbl.rows.Len()
	req := &btapb.DropRowRangeRequest{
		Name:   tblInfo.Name,
		Target: &btapb.DropRowRangeRequest_RowKeyPrefix{RowKeyPrefix: []byte("AAA")},
	}
	if _, err = s.DropRowRange(ctx, req); err != nil {
		t.Fatalf("Dropping first range: %v", err)
	}
	got, want := tbl.rows.Len(), tblSize-count
	if got != want {
		t.Errorf("Row count after first drop: got %d (%v), want %d", got, tbl.rows, want)
	}

	req = &btapb.DropRowRangeRequest{
		Name:   tblInfo.Name,
		Target: &btapb.DropRowRangeRequest_RowKeyPrefix{RowKeyPrefix: []byte("DDD")},
	}
	if _, err = s.DropRowRange(ctx, req); err != nil {
		t.Fatalf("Dropping second range: %v", err)
	}
	got, want = tbl.rows.Len(), tblSize-(2*count)
	if got != want {
		t.Errorf("Row count after second drop: got %d (%v), want %d", got, tbl.rows, want)
	}

	req = &btapb.DropRowRangeRequest{
		Name:   tblInfo.Name,
		Target: &btapb.DropRowRangeRequest_RowKeyPrefix{RowKeyPrefix: []byte("XXX")},
	}
	if _, err = s.DropRowRange(ctx, req); err != nil {
		t.Fatalf("Dropping invalid range: %v", err)
	}
	got, want = tbl.rows.Len(), tblSize-(2*count)
	if got != want {
		t.Errorf("Row count after invalid drop: got %d (%v), want %d", got, tbl.rows, want)
	}

	req = &btapb.DropRowRangeRequest{
		Name:   tblInfo.Name,
		Target: &btapb.DropRowRangeRequest_DeleteAllDataFromTable{DeleteAllDataFromTable: true},
	}
	if _, err = s.DropRowRange(ctx, req); err != nil {
		t.Fatalf("Dropping all data: %v", err)
	}
	got, want = tbl.rows.Len(), 0
	if got != want {
		t.Errorf("Row count after drop all: got %d, want %d", got, want)
	}

	// Test that we can write rows, delete some and then write them again.
	count = 1
	doWrite()

	req = &btapb.DropRowRangeRequest{
		Name:   tblInfo.Name,
		Target: &btapb.DropRowRangeRequest_DeleteAllDataFromTable{DeleteAllDataFromTable: true},
	}
	if _, err = s.DropRowRange(ctx, req); err != nil {
		t.Fatalf("Dropping all data: %v", err)
	}
	got, want = tbl.rows.Len(), 0
	if got != want {
		t.Errorf("Row count after drop all: got %d, want %d", got, want)
	}

	doWrite()
	got, want = tbl.rows.Len(), len(prefixes)
	if got != want {
		t.Errorf("Row count after rewrite: got %d, want %d", got, want)
	}

	req = &btapb.DropRowRangeRequest{
		Name:   tblInfo.Name,
		Target: &btapb.DropRowRangeRequest_RowKeyPrefix{RowKeyPrefix: []byte("BBB")},
	}
	if _, err = s.DropRowRange(ctx, req); err != nil {
		t.Fatalf("Dropping range: %v", err)
	}
	doWrite()
	got, want = tbl.rows.Len(), len(prefixes)
	if got != want {
		t.Errorf("Row count after drop range: got %d, want %d", got, want)
	}
}

type MockReadRowsServer struct {
	responses []*btpb.ReadRowsResponse
	ctx       context.Context
	grpc.ServerStream
}

func (s *MockReadRowsServer) Send(resp *btpb.ReadRowsResponse) error {
	s.responses = append(s.responses, resp)
	return nil
}

func (s *MockReadRowsServer) Context() context.Context {
	if s.ctx == nil {
		return context.Background()
	}
	return s.ctx
}

func TestCheckTimestampMaxValue(t *testing.T) {
	// Test that max Timestamp value can be passed in TimestampMicros without error
	// and that max Timestamp is the largest valid value in Millis.
	// See issue https://github.com/googleapis/google-cloud-go/issues/1790
	ctx := context.Background()
	s := &server{
		tables: make(map[string]*table),
	}
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf0": {},
		},
	}
	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "issue-1790", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}
	var maxTimestamp int64 = math.MaxInt64 - math.MaxInt64%1000
	mreq1 := &btpb.MutateRowRequest{
		TableName: tblInfo.Name,
		RowKey:    []byte("row"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
				FamilyName:      "cf0",
				ColumnQualifier: []byte("col"),
				TimestampMicros: maxTimestamp,
				Value:           []byte{},
			}},
		}},
	}
	if _, err := s.MutateRow(ctx, mreq1); err != nil {
		t.Fatalf("TimestampMicros wasn't set: %v", err)
	}

	mreq2 := &btpb.MutateRowRequest{
		TableName: tblInfo.Name,
		RowKey:    []byte("row"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
				FamilyName:      "cf0",
				ColumnQualifier: []byte("col"),
				TimestampMicros: maxTimestamp + 1000,
				Value:           []byte{},
			}},
		}},
	}
	if _, err := s.MutateRow(ctx, mreq2); err == nil {
		t.Fatalf("want TimestampMicros rejection, got acceptance: %v", err)
	}
}

func TestReadRows(t *testing.T) {
	ctx := context.Background()
	s := &server{
		tables: make(map[string]*table),
	}
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf0": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}
	mreq := &btpb.MutateRowRequest{
		TableName: tblInfo.Name,
		RowKey:    []byte("row"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
				FamilyName:      "cf0",
				ColumnQualifier: []byte("col"),
				TimestampMicros: 1000,
				Value:           []byte{},
			}},
		}},
	}
	if _, err := s.MutateRow(ctx, mreq); err != nil {
		t.Fatalf("Populating table: %v", err)
	}

	for _, rowset := range []*btpb.RowSet{
		{RowKeys: [][]byte{[]byte("row")}},
		{RowRanges: []*btpb.RowRange{{StartKey: &btpb.RowRange_StartKeyClosed{StartKeyClosed: []byte("")}}}},
		{RowRanges: []*btpb.RowRange{{StartKey: &btpb.RowRange_StartKeyClosed{StartKeyClosed: []byte("r")}}}},
		{RowRanges: []*btpb.RowRange{{
			StartKey: &btpb.RowRange_StartKeyClosed{StartKeyClosed: []byte("")},
			EndKey:   &btpb.RowRange_EndKeyOpen{EndKeyOpen: []byte("s")},
		}}},
	} {
		mock := &MockReadRowsServer{}
		req := &btpb.ReadRowsRequest{TableName: tblInfo.Name, Rows: rowset}
		if err = s.ReadRows(req, mock); err != nil {
			t.Fatalf("ReadRows error: %v", err)
		}
		if got, want := len(mock.responses), 1; got != want {
			t.Errorf("%+v: response count: got %d, want %d", rowset, got, want)
		}
	}
}

// withFeatureFlags set the feature flags the client supports in the
// `bigtable-features` header sent on each request.
func withFeatureFlags() metadata.MD {
	ffStr := ""
	ff := btpb.FeatureFlags{
		ReverseScans:             true,
		LastScannedRowResponses:  true,
		ClientSideMetricsEnabled: false, // Not suppported in emulator
	}
	b, err := proto.Marshal(&ff)
	if err == nil {
		ffStr = base64.URLEncoding.EncodeToString(b)
	}
	return metadata.Pairs("bigtable-features", ffStr)
}

func TestReadRowsLastScannedRow(t *testing.T) {
	ctx := context.Background()
	s := &server{
		tables: make(map[string]*table),
	}
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf0": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}
	mreq := &btpb.MutateRowRequest{
		TableName: tblInfo.Name,
		RowKey:    []byte("row"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
				FamilyName:      "cf0",
				ColumnQualifier: []byte("col"),
				TimestampMicros: 1000,
				Value:           []byte{},
			}},
		}},
	}
	if _, err := s.MutateRow(ctx, mreq); err != nil {
		t.Fatalf("Populating table: %v", err)
	}

	for _, rowset := range []*btpb.RowSet{
		{RowKeys: [][]byte{[]byte("row")}},
		{RowRanges: []*btpb.RowRange{{StartKey: &btpb.RowRange_StartKeyClosed{StartKeyClosed: []byte("")}}}},
		{RowRanges: []*btpb.RowRange{{StartKey: &btpb.RowRange_StartKeyClosed{StartKeyClosed: []byte("r")}}}},
		{RowRanges: []*btpb.RowRange{{
			StartKey: &btpb.RowRange_StartKeyClosed{StartKeyClosed: []byte("")},
			EndKey:   &btpb.RowRange_EndKeyOpen{EndKeyOpen: []byte("s")},
		}}},
	} {
		featureFlags := withFeatureFlags()
		ctx := metadata.NewIncomingContext(context.Background(), featureFlags)

		mock := &MockReadRowsServer{ctx: ctx}
		filter := &btpb.RowFilter{Filter: &btpb.RowFilter_BlockAllFilter{BlockAllFilter: true}}
		req := &btpb.ReadRowsRequest{TableName: tblInfo.Name, Rows: rowset, Filter: filter}
		if err = s.ReadRows(req, mock); err != nil {
			t.Fatalf("ReadRows error: %v", err)
		}
		if got, want := len(mock.responses), 1; got != want {
			t.Errorf("%+v: response count: got %d, want %d", rowset, got, want)
		}
		expect := &btpb.ReadRowsResponse{LastScannedRowKey: []byte("row")}
		if !proto.Equal(mock.responses[0], expect) {
			t.Errorf("%+v: response: got %+v, want %+v", rowset, mock.responses[0], expect)
		}
	}
}

func TestReadRowsError(t *testing.T) {
	ctx := context.Background()
	s := &server{
		tables: make(map[string]*table),
	}
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf0": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}
	mreq := &btpb.MutateRowRequest{
		TableName: tblInfo.Name,
		RowKey:    []byte("row"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
				FamilyName:      "cf0",
				ColumnQualifier: []byte("col"),
				TimestampMicros: 1000,
				Value:           []byte{},
			}},
		}},
	}
	if _, err := s.MutateRow(ctx, mreq); err != nil {
		t.Fatalf("Populating table: %v", err)
	}

	mock := &MockReadRowsServer{}
	req := &btpb.ReadRowsRequest{
		TableName: tblInfo.Name, Filter: &btpb.RowFilter{
			Filter: &btpb.RowFilter_RowKeyRegexFilter{RowKeyRegexFilter: []byte("[")},
		}, // Invalid regex.
	}
	if err = s.ReadRows(req, mock); err == nil {
		t.Fatal("ReadRows got no error, want error")
	}
}

func TestReadRowsErrorOnEmptyRegex(t *testing.T) {
	ctx := context.Background()
	s := &server{
		tables: make(map[string]*table),
	}
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf0": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}
	mreq := &btpb.MutateRowRequest{
		TableName: tblInfo.Name,
		RowKey:    []byte("row"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
				FamilyName:      "cf0",
				ColumnQualifier: []byte("col"),
				TimestampMicros: 1000,
				Value:           []byte{},
			}},
		}},
	}
	if _, err := s.MutateRow(ctx, mreq); err != nil {
		t.Fatalf("Populating table: %v", err)
	}

	mock := &MockReadRowsServer{}
	req := &btpb.ReadRowsRequest{
		TableName: tblInfo.Name, Filter: &btpb.RowFilter{
			Filter: &btpb.RowFilter_RowKeyRegexFilter{RowKeyRegexFilter: []byte("")},
		}, // Empty regexes should be rejected.
	}
	if err = s.ReadRows(req, mock); err == nil {
		t.Fatal("ReadRows got no error, want error")
	}
}

func TestReadRowsAfterDeletion(t *testing.T) {
	ctx := context.Background()
	s := &server{
		tables: make(map[string]*table),
	}
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf0": {},
		},
	}
	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{
		Parent: "cluster", TableId: "t", Table: &newTbl,
	})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}
	populateTable(ctx, s)
	dreq := &btpb.MutateRowRequest{
		TableName: tblInfo.Name,
		RowKey:    []byte("row"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_DeleteFromRow_{
				DeleteFromRow: &btpb.Mutation_DeleteFromRow{},
			},
		}},
	}
	if _, err := s.MutateRow(ctx, dreq); err != nil {
		t.Fatalf("Deleting from table: %v", err)
	}

	mock := &MockReadRowsServer{}
	req := &btpb.ReadRowsRequest{TableName: tblInfo.Name}
	if err = s.ReadRows(req, mock); err != nil {
		t.Fatalf("ReadRows error: %v", err)
	}
	if got, want := len(mock.responses), 0; got != want {
		t.Errorf("response count: got %d, want %d", got, want)
	}
}

func TestReadRowsOrder(t *testing.T) {
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf0": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}
	count := 3
	mcf := func(i int) *btapb.ModifyColumnFamiliesRequest {
		return &btapb.ModifyColumnFamiliesRequest{
			Name: tblInfo.Name,
			Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{{
				Id:  "cf" + strconv.Itoa(i),
				Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Create{Create: &btapb.ColumnFamily{}},
			}},
		}
	}
	for i := 1; i <= count; i++ {
		_, err = s.ModifyColumnFamilies(ctx, mcf(i))
		if err != nil {
			t.Fatal(err)
		}
	}
	// Populate the table
	for fc := 0; fc < count; fc++ {
		for cc := count; cc > 0; cc-- {
			for tc := 0; tc < count; tc++ {
				req := &btpb.MutateRowRequest{
					TableName: tblInfo.Name,
					RowKey:    []byte("row"),
					Mutations: []*btpb.Mutation{{
						Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
							FamilyName:      "cf" + strconv.Itoa(fc),
							ColumnQualifier: []byte("col" + strconv.Itoa(cc)),
							TimestampMicros: int64((tc + 1) * 1000),
							Value:           []byte{},
						}},
					}},
				}
				if _, err := s.MutateRow(ctx, req); err != nil {
					t.Fatalf("Populating table: %v", err)
				}
			}
		}
	}
	req := &btpb.ReadRowsRequest{
		TableName: tblInfo.Name,
		Rows:      &btpb.RowSet{RowKeys: [][]byte{[]byte("row")}},
	}
	mock := &MockReadRowsServer{}
	if err = s.ReadRows(req, mock); err != nil {
		t.Errorf("ReadRows error: %v", err)
	}
	if len(mock.responses) == 0 {
		t.Fatal("Response count: got 0, want > 0")
	}
	if len(mock.responses[0].Chunks) != 27 {
		t.Fatalf("Chunk count: got %d, want 27", len(mock.responses[0].Chunks))
	}
	testOrder := func(ms *MockReadRowsServer) {
		var prevFam, prevCol string
		var prevTime int64
		for _, cc := range ms.responses[0].Chunks {
			if prevFam == "" {
				prevFam = cc.FamilyName.Value
				prevCol = string(cc.Qualifier.Value)
				prevTime = cc.TimestampMicros
				continue
			}
			if cc.FamilyName.Value < prevFam {
				t.Errorf("Family order is not correct: got %s < %s", cc.FamilyName.Value, prevFam)
			} else if cc.FamilyName.Value == prevFam {
				if string(cc.Qualifier.Value) < prevCol {
					t.Errorf("Column order is not correct: got %s < %s", string(cc.Qualifier.Value), prevCol)
				} else if string(cc.Qualifier.Value) == prevCol {
					if cc.TimestampMicros > prevTime {
						t.Errorf("cell order is not correct: got %d > %d", cc.TimestampMicros, prevTime)
					}
				}
			}
			prevFam = cc.FamilyName.Value
			prevCol = string(cc.Qualifier.Value)
			prevTime = cc.TimestampMicros
		}
	}
	testOrder(mock)

	// Read with interleave filter
	inter := &btpb.RowFilter_Interleave{}
	fnr := &btpb.RowFilter{Filter: &btpb.RowFilter_FamilyNameRegexFilter{FamilyNameRegexFilter: "cf1"}}
	cqr := &btpb.RowFilter{Filter: &btpb.RowFilter_ColumnQualifierRegexFilter{ColumnQualifierRegexFilter: []byte("col2")}}
	inter.Filters = append(inter.Filters, fnr, cqr)
	req = &btpb.ReadRowsRequest{
		TableName: tblInfo.Name,
		Rows:      &btpb.RowSet{RowKeys: [][]byte{[]byte("row")}},
		Filter: &btpb.RowFilter{
			Filter: &btpb.RowFilter_Interleave_{Interleave: inter},
		},
	}

	mock = &MockReadRowsServer{}
	if err = s.ReadRows(req, mock); err != nil {
		t.Errorf("ReadRows error: %v", err)
	}
	if len(mock.responses) == 0 {
		t.Fatal("Response count: got 0, want > 0")
	}
	if len(mock.responses[0].Chunks) != 18 {
		t.Fatalf("Chunk count: got %d, want 18", len(mock.responses[0].Chunks))
	}
	testOrder(mock)

	// Check order after ReadModifyWriteRow
	rmw := func(i int) *btpb.ReadModifyWriteRowRequest {
		return &btpb.ReadModifyWriteRowRequest{
			TableName: tblInfo.Name,
			RowKey:    []byte("row"),
			Rules: []*btpb.ReadModifyWriteRule{{
				FamilyName:      "cf3",
				ColumnQualifier: []byte("col" + strconv.Itoa(i)),
				Rule:            &btpb.ReadModifyWriteRule_IncrementAmount{IncrementAmount: 1},
			}},
		}
	}
	for i := count; i > 0; i-- {
		if _, err := s.ReadModifyWriteRow(ctx, rmw(i)); err != nil {
			t.Fatal(err)
		}
	}
	req = &btpb.ReadRowsRequest{
		TableName: tblInfo.Name,
		Rows:      &btpb.RowSet{RowKeys: [][]byte{[]byte("row")}},
	}
	mock = &MockReadRowsServer{}
	if err = s.ReadRows(req, mock); err != nil {
		t.Errorf("ReadRows error: %v", err)
	}
	if len(mock.responses) == 0 {
		t.Fatal("Response count: got 0, want > 0")
	}
	if len(mock.responses[0].Chunks) != 30 {
		t.Fatalf("Chunk count: got %d, want 30", len(mock.responses[0].Chunks))
	}
	testOrder(mock)
}

func TestReadRowsWithlabelTransformer(t *testing.T) {
	ctx := context.Background()
	s := &server{
		tables: make(map[string]*table),
	}
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf0": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}
	mreq := &btpb.MutateRowRequest{
		TableName: tblInfo.Name,
		RowKey:    []byte("row"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
				FamilyName:      "cf0",
				ColumnQualifier: []byte("col"),
				TimestampMicros: 1000,
				Value:           []byte{},
			}},
		}},
	}
	if _, err := s.MutateRow(ctx, mreq); err != nil {
		t.Fatalf("Populating table: %v", err)
	}

	mock := &MockReadRowsServer{}
	req := &btpb.ReadRowsRequest{
		TableName: tblInfo.Name,
		Filter: &btpb.RowFilter{
			Filter: &btpb.RowFilter_ApplyLabelTransformer{
				ApplyLabelTransformer: "label",
			},
		},
	}
	if err = s.ReadRows(req, mock); err != nil {
		t.Fatalf("ReadRows error: %v", err)
	}

	if got, want := len(mock.responses), 1; got != want {
		t.Fatalf("response count: got %d, want %d", got, want)
	}
	resp := mock.responses[0]
	if got, want := len(resp.Chunks), 1; got != want {
		t.Fatalf("chunks count: got %d, want %d", got, want)
	}
	chunk := resp.Chunks[0]
	if got, want := len(chunk.Labels), 1; got != want {
		t.Fatalf("labels count: got %d, want %d", got, want)
	}
	if got, want := chunk.Labels[0], "label"; got != want {
		t.Fatalf("label: got %s, want %s", got, want)
	}

	mock = &MockReadRowsServer{}
	req = &btpb.ReadRowsRequest{
		TableName: tblInfo.Name,
		Filter: &btpb.RowFilter{
			Filter: &btpb.RowFilter_ApplyLabelTransformer{
				ApplyLabelTransformer: "", // invalid label
			},
		},
	}
	if err = s.ReadRows(req, mock); err == nil {
		t.Fatal("ReadRows want invalid label error, got none")
	}
}

func TestReadRowsReversed(t *testing.T) {
	ctx := context.Background()
	srv := &server{
		tables: make(map[string]*table),
	}
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tbl, err := srv.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}
	entries := []struct {
		row   string
		value []byte
	}{
		{"row1", []byte("a")},
		{"row2", []byte("b")},
	}

	for _, entry := range entries {
		req := &btpb.MutateRowRequest{
			TableName: tbl.Name,
			RowKey:    []byte(entry.row),
			Mutations: []*btpb.Mutation{{
				Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
					FamilyName:      "cf",
					ColumnQualifier: []byte("cq"),
					TimestampMicros: 1000,
					Value:           entry.value,
				}},
			}},
		}
		if _, err := srv.MutateRow(ctx, req); err != nil {
			t.Fatalf("Failed to insert entry %v into server: %v", entry, err)
		}
	}

	serverCtx := metadata.NewIncomingContext(context.Background(), withFeatureFlags())
	rrss := &MockReadRowsServer{ctx: serverCtx}
	rreq := &btpb.ReadRowsRequest{TableName: tbl.Name, Reversed: true}
	if err := srv.ReadRows(rreq, rrss); err != nil {
		t.Fatalf("Failed to read rows: %v", err)
	}

	var gotChunks []*btpb.ReadRowsResponse_CellChunk
	for _, res := range rrss.responses {
		gotChunks = append(gotChunks, res.Chunks...)
	}

	wantChunks := []*btpb.ReadRowsResponse_CellChunk{
		{
			RowKey:          []byte("row2"),
			FamilyName:      &wrapperspb.StringValue{Value: "cf"},
			Qualifier:       &wrapperspb.BytesValue{Value: []byte("cq")},
			TimestampMicros: 1000,
			Value:           []byte("b"),
			RowStatus: &btpb.ReadRowsResponse_CellChunk_CommitRow{
				CommitRow: true,
			},
		},
		{
			RowKey:          []byte("row1"),
			FamilyName:      &wrapperspb.StringValue{Value: "cf"},
			Qualifier:       &wrapperspb.BytesValue{Value: []byte("cq")},
			TimestampMicros: 1000,
			Value:           []byte("a"),
			RowStatus: &btpb.ReadRowsResponse_CellChunk_CommitRow{
				CommitRow: true,
			},
		},
	}
	if diff := cmp.Diff(gotChunks, wantChunks, protocmp.Transform()); diff != "" {
		t.Fatalf("Response chunks mismatch: got: + want -\n%s", diff)
	}
}

func TestCheckAndMutateRowWithoutPredicate(t *testing.T) {
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tbl, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}

	val := []byte("value")
	muts := []*btpb.Mutation{{
		Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
			FamilyName:      "cf",
			ColumnQualifier: []byte("col"),
			TimestampMicros: 1000,
			Value:           val,
		}},
	}}

	mrreq := &btpb.MutateRowRequest{
		TableName: tbl.Name,
		RowKey:    []byte("row-present"),
		Mutations: muts,
	}
	if _, err := s.MutateRow(ctx, mrreq); err != nil {
		t.Fatalf("Populating table: %v", err)
	}

	req := &btpb.CheckAndMutateRowRequest{
		TableName:      tbl.Name,
		RowKey:         []byte("row-not-present"),
		FalseMutations: muts,
	}
	if res, err := s.CheckAndMutateRow(ctx, req); err != nil {
		t.Errorf("CheckAndMutateRow error: %v", err)
	} else if got, want := res.PredicateMatched, false; got != want {
		t.Errorf("Invalid PredicateMatched value: got %t, want %t", got, want)
	}

	req = &btpb.CheckAndMutateRowRequest{
		TableName:      tbl.Name,
		RowKey:         []byte("row-present"),
		FalseMutations: muts,
	}
	if res, err := s.CheckAndMutateRow(ctx, req); err != nil {
		t.Errorf("CheckAndMutateRow error: %v", err)
	} else if got, want := res.PredicateMatched, true; got != want {
		t.Errorf("Invalid PredicateMatched value: got %t, want %t", got, want)
	}
}

func TestCheckAndMutateRowWithPredicate(t *testing.T) {
	ctx := context.Background()
	srv := &server{tables: make(map[string]*table)}

	tblReq := &btapb.CreateTableRequest{
		Parent:  "issue-1435",
		TableId: "table_id",
		Table: &btapb.Table{
			ColumnFamilies: map[string]*btapb.ColumnFamily{
				"cf": {},
				"df": {},
				"ef": {},
				"ff": {},
				"zf": {},
			},
		},
	}
	tbl, err := srv.CreateTable(ctx, tblReq)
	if err != nil {
		t.Fatalf("Failed to create the table: %v", err)
	}

	entries := []struct {
		row                         string
		value                       []byte
		familyName, columnQualifier string
	}{
		{"row1", []byte{0x11}, "cf", "cq"},
		{"row2", []byte{0x1a}, "df", "dq"},
		{"row3", []byte{'a'}, "ef", "eq"},
		{"row4", []byte{'b'}, "ff", "fq"},
	}

	for _, entry := range entries {
		req := &btpb.MutateRowRequest{
			TableName: tbl.Name,
			RowKey:    []byte(entry.row),
			Mutations: []*btpb.Mutation{{
				Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
					FamilyName:      entry.familyName,
					ColumnQualifier: []byte(entry.columnQualifier),
					TimestampMicros: 1000,
					Value:           entry.value,
				}},
			}},
		}
		if _, err := srv.MutateRow(ctx, req); err != nil {
			t.Fatalf("Failed to insert entry %v into server: %v", entry, err)
		}
	}

	var bogusMutations = []*btpb.Mutation{{
		Mutation: &btpb.Mutation_DeleteFromFamily_{
			DeleteFromFamily: &btpb.Mutation_DeleteFromFamily{
				FamilyName: "bogus_family",
			},
		},
	}}

	tests := []struct {
		req       *btpb.CheckAndMutateRowRequest
		wantMatch bool
		name      string

		// if wantState is nil, that means we don't care to check
		// what the state of the world is.
		wantState []*btpb.ReadRowsResponse_CellChunk
	}{
		{
			req: &btpb.CheckAndMutateRowRequest{
				TableName: tbl.Name,
				RowKey:    []byte("row1"),
				PredicateFilter: &btpb.RowFilter{
					Filter: &btpb.RowFilter_RowKeyRegexFilter{
						RowKeyRegexFilter: []byte("not-one"),
					},
				},
				TrueMutations: bogusMutations,
			},
			name: "no match",
		},
		{
			req: &btpb.CheckAndMutateRowRequest{
				TableName: tbl.Name,
				RowKey:    []byte("row1"),
				PredicateFilter: &btpb.RowFilter{
					Filter: &btpb.RowFilter_RowKeyRegexFilter{
						RowKeyRegexFilter: []byte("ro.+"),
					},
				},
				FalseMutations: bogusMutations,
			},
			wantMatch: true,
			name:      "rowkey regex",
		},
		{
			req: &btpb.CheckAndMutateRowRequest{
				TableName: tbl.Name,
				RowKey:    []byte("row1"),
				PredicateFilter: &btpb.RowFilter{
					Filter: &btpb.RowFilter_PassAllFilter{
						PassAllFilter: true,
					},
				},
				FalseMutations: bogusMutations,
			},
			wantMatch: true,
			name:      "pass all",
		},
		{
			req: &btpb.CheckAndMutateRowRequest{
				TableName: tbl.Name,
				RowKey:    []byte("row1"),
				PredicateFilter: &btpb.RowFilter{
					Filter: &btpb.RowFilter_BlockAllFilter{
						BlockAllFilter: true,
					},
				},
				FalseMutations: []*btpb.Mutation{
					{
						Mutation: &btpb.Mutation_SetCell_{
							SetCell: &btpb.Mutation_SetCell{
								FamilyName:      "zf",
								Value:           []byte("foo"),
								TimestampMicros: 2000,
								ColumnQualifier: []byte("et"),
							},
						},
					},
				},
			},
			name:      "BlockAll for row1",
			wantMatch: false,
			wantState: []*btpb.ReadRowsResponse_CellChunk{
				{
					RowKey: []byte("row1"),
					FamilyName: &wrapperspb.StringValue{
						Value: "cf",
					},
					Qualifier: &wrapperspb.BytesValue{
						Value: []byte("cq"),
					},
					TimestampMicros: 1000,
					Value:           []byte{0x11},
				},
				{
					RowKey: []byte("row1"),
					FamilyName: &wrapperspb.StringValue{
						Value: "zf",
					},
					Qualifier: &wrapperspb.BytesValue{
						Value: []byte("et"),
					},
					TimestampMicros: 2000,
					Value:           []byte("foo"),
					RowStatus: &btpb.ReadRowsResponse_CellChunk_CommitRow{
						CommitRow: true,
					},
				},
				{
					RowKey: []byte("row2"),
					FamilyName: &wrapperspb.StringValue{
						Value: "df",
					},
					Qualifier: &wrapperspb.BytesValue{
						Value: []byte("dq"),
					},
					TimestampMicros: 1000,
					Value:           []byte{0x1a},
					RowStatus: &btpb.ReadRowsResponse_CellChunk_CommitRow{
						CommitRow: true,
					},
				},
				{
					RowKey: []byte("row3"),
					FamilyName: &wrapperspb.StringValue{
						Value: "ef",
					},
					Qualifier: &wrapperspb.BytesValue{
						Value: []byte("eq"),
					},
					TimestampMicros: 1000,
					Value:           []byte("a"),
					RowStatus: &btpb.ReadRowsResponse_CellChunk_CommitRow{
						CommitRow: true,
					},
				},
				{
					RowKey: []byte("row4"),
					FamilyName: &wrapperspb.StringValue{
						Value: "ff",
					},
					Qualifier: &wrapperspb.BytesValue{
						Value: []byte("fq"),
					},
					TimestampMicros: 1000,
					Value:           []byte("b"),
					RowStatus: &btpb.ReadRowsResponse_CellChunk_CommitRow{
						CommitRow: true,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := srv.CheckAndMutateRow(ctx, tt.req)
			if err != nil {
				t.Fatalf("CheckAndMutateRow error: %v", err)
			}
			got, want := res.PredicateMatched, tt.wantMatch
			if got != want {
				t.Fatalf("Invalid PredicateMatched value: got %t, want %t\nRequest: %+v", got, want, tt.req)
			}

			if tt.wantState == nil {
				return
			}

			rreq := &btpb.ReadRowsRequest{TableName: tbl.Name}
			mock := &MockReadRowsServer{}
			if err = srv.ReadRows(rreq, mock); err != nil {
				t.Fatalf("ReadRows error: %v", err)
			}

			// Collect all the cellChunks
			var gotCellChunks []*btpb.ReadRowsResponse_CellChunk
			for _, res := range mock.responses {
				gotCellChunks = append(gotCellChunks, res.Chunks...)
			}
			sort.Slice(gotCellChunks, func(i, j int) bool {
				ci, cj := gotCellChunks[i], gotCellChunks[j]
				return compareCellChunks(ci, cj)
			})
			wantCellChunks := tt.wantState[0:]
			sort.Slice(wantCellChunks, func(i, j int) bool {
				return compareCellChunks(wantCellChunks[i], wantCellChunks[j])
			})

			// bttest for some reason undeterministically returns:
			//      RowStatus: &bigtable.ReadRowsResponse_CellChunk_CommitRow{CommitRow: true},
			// so we'll ignore that field during comparison.
			scrubRowStatus := func(cs []*btpb.ReadRowsResponse_CellChunk) []*btpb.ReadRowsResponse_CellChunk {
				for _, c := range cs {
					c.RowStatus = nil
				}
				return cs
			}
			diff := cmp.Diff(scrubRowStatus(gotCellChunks), scrubRowStatus(wantCellChunks), cmp.Comparer(proto.Equal))
			if diff != "" {
				t.Fatalf("unexpected response: %s", diff)
			}
		})
	}
}

// compareCellChunks is a comparator that is passed
// into sort.Slice to stably sort cell chunks.
func compareCellChunks(ci, cj *btpb.ReadRowsResponse_CellChunk) bool {
	if bytes.Compare(ci.RowKey, cj.RowKey) > 0 {
		return false
	}
	if bytes.Compare(ci.Value, cj.Value) > 0 {
		return false
	}
	return ci.FamilyName.GetValue() < cj.FamilyName.GetValue()
}

func TestServer_ReadModifyWriteRow(t *testing.T) {
	s := &server{
		tables: make(map[string]*table),
	}

	ctx := context.Background()
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tbl, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		t.Fatalf("Creating table: %v", err)
	}

	req := &btpb.ReadModifyWriteRowRequest{
		TableName: tbl.Name,
		RowKey:    []byte("row-key"),
		Rules: []*btpb.ReadModifyWriteRule{
			{
				FamilyName:      "cf",
				ColumnQualifier: []byte("q1"),
				Rule: &btpb.ReadModifyWriteRule_AppendValue{
					AppendValue: []byte("a"),
				},
			},
			// multiple ops for same cell
			{
				FamilyName:      "cf",
				ColumnQualifier: []byte("q1"),
				Rule: &btpb.ReadModifyWriteRule_AppendValue{
					AppendValue: []byte("b"),
				},
			},
			// different cell whose qualifier should sort before the prior rules
			{
				FamilyName:      "cf",
				ColumnQualifier: []byte("q0"),
				Rule: &btpb.ReadModifyWriteRule_IncrementAmount{
					IncrementAmount: 1,
				},
			},
		},
	}

	got, err := s.ReadModifyWriteRow(ctx, req)
	if err != nil {
		t.Fatalf("ReadModifyWriteRow error: %v", err)
	}

	want := &btpb.ReadModifyWriteRowResponse{
		Row: &btpb.Row{
			Key: []byte("row-key"),
			Families: []*btpb.Family{{
				Name: "cf",
				Columns: []*btpb.Column{
					{
						Qualifier: []byte("q0"),
						Cells: []*btpb.Cell{{
							Value: []byte{0, 0, 0, 0, 0, 0, 0, 1},
						}},
					},
					{
						Qualifier: []byte("q1"),
						Cells: []*btpb.Cell{{
							Value: []byte("ab"),
						}},
					},
				},
			}},
		},
	}

	scrubTimestamp := func(resp *btpb.ReadModifyWriteRowResponse) *btpb.ReadModifyWriteRowResponse {
		for _, fam := range resp.GetRow().GetFamilies() {
			for _, col := range fam.GetColumns() {
				for _, cell := range col.GetCells() {
					cell.TimestampMicros = 0
				}
			}
		}
		return resp
	}
	diff := cmp.Diff(scrubTimestamp(got), scrubTimestamp(want), cmp.Comparer(proto.Equal))
	if diff != "" {
		t.Errorf("unexpected response: %s", diff)
	}
}

// helper function to populate table data
func populateTable(ctx context.Context, s *server) (*btapb.Table, error) {
	newTbl := btapb.Table{
		ColumnFamilies: map[string]*btapb.ColumnFamily{
			"cf0": {GcRule: &btapb.GcRule{Rule: &btapb.GcRule_MaxNumVersions{MaxNumVersions: 1}}},
		},
	}
	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{Parent: "cluster", TableId: "t", Table: &newTbl})
	if err != nil {
		return nil, err
	}
	count := 3
	mcf := func(i int) *btapb.ModifyColumnFamiliesRequest {
		return &btapb.ModifyColumnFamiliesRequest{
			Name: tblInfo.Name,
			Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{{
				Id:  "cf" + strconv.Itoa(i),
				Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Create{Create: &btapb.ColumnFamily{}},
			}},
		}
	}
	for i := 1; i <= count; i++ {
		_, err = s.ModifyColumnFamilies(ctx, mcf(i))
		if err != nil {
			return nil, err
		}
	}
	// Populate the table
	for fc := 0; fc < count; fc++ {
		for cc := count; cc > 0; cc-- {
			for tc := 0; tc < count; tc++ {
				req := &btpb.MutateRowRequest{
					TableName: tblInfo.Name,
					RowKey:    []byte("row"),
					Mutations: []*btpb.Mutation{{
						Mutation: &btpb.Mutation_SetCell_{
							SetCell: &btpb.Mutation_SetCell{
								FamilyName:      "cf" + strconv.Itoa(fc),
								ColumnQualifier: []byte("col" + strconv.Itoa(cc)),
								TimestampMicros: int64((tc + 1) * 1000),
								Value:           []byte{},
							}},
					}},
				}
				if _, err := s.MutateRow(ctx, req); err != nil {
					return nil, err
				}
			}
		}
	}

	return tblInfo, nil
}

func TestFilters(t *testing.T) {
	tests := []struct {
		in   *btpb.RowFilter
		code codes.Code
		out  int
	}{
		{in: &btpb.RowFilter{Filter: &btpb.RowFilter_BlockAllFilter{BlockAllFilter: true}}, out: 0},
		{in: &btpb.RowFilter{Filter: &btpb.RowFilter_BlockAllFilter{BlockAllFilter: false}}, code: codes.InvalidArgument},
		{in: &btpb.RowFilter{Filter: &btpb.RowFilter_PassAllFilter{PassAllFilter: true}}, out: 1},
		{in: &btpb.RowFilter{Filter: &btpb.RowFilter_PassAllFilter{PassAllFilter: false}}, code: codes.InvalidArgument},
	}

	ctx := context.Background()

	s := &server{
		tables: make(map[string]*table),
	}

	tblInfo, err := populateTable(ctx, s)
	if err != nil {
		t.Fatal(err)
	}

	req := &btpb.ReadRowsRequest{
		TableName: tblInfo.Name,
		Rows:      &btpb.RowSet{RowKeys: [][]byte{[]byte("row")}},
	}

	for _, tc := range tests {
		req.Filter = tc.in

		mock := &MockReadRowsServer{}
		err := s.ReadRows(req, mock)
		if tc.code != codes.OK {
			s, _ := status.FromError(err)
			if s.Code() != tc.code {
				t.Errorf("error code: got %d, want %d", s.Code(), tc.code)
			}
			continue
		}

		if err != nil {
			t.Errorf("ReadRows error: %v", err)
			continue
		}

		if len(mock.responses) != tc.out {
			t.Errorf("Response count: got %d, want %d", len(mock.responses), tc.out)
			continue
		}
	}
}

func TestMutateRowsAggregate_AddToCell(t *testing.T) {
	ctx := context.Background()

	s := &server{
		tables: make(map[string]*table),
	}

	tblInfo, err := populateTable(ctx, s)
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.ModifyColumnFamilies(ctx, &btapb.ModifyColumnFamiliesRequest{
		Name: tblInfo.Name,
		Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{{
			Id: "sum",
			Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Create{
				Create: &btapb.ColumnFamily{
					ValueType: &btapb.Type{
						Kind: &btapb.Type_AggregateType{
							AggregateType: &btapb.Type_Aggregate{
								InputType: &btapb.Type{
									Kind: &btapb.Type_Int64Type{},
								},
								Aggregator: &btapb.Type_Aggregate_Sum_{
									Sum: &btapb.Type_Aggregate_Sum{},
								},
							},
						},
					},
				},
			}},
		}})

	if err != nil {
		t.Fatal(err)
	}

	_, err = s.MutateRow(ctx, &btpb.MutateRowRequest{
		TableName: tblInfo.GetName(),
		RowKey:    []byte("row1"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_AddToCell_{AddToCell: &btpb.Mutation_AddToCell{
				FamilyName:      "sum",
				ColumnQualifier: &btpb.Value{Kind: &btpb.Value_RawValue{RawValue: []byte("col1")}},
				Timestamp:       &btpb.Value{Kind: &btpb.Value_RawTimestampMicros{RawTimestampMicros: 0}},
				Input:           &btpb.Value{Kind: &btpb.Value_IntValue{IntValue: 1}},
			}},
		}},
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = s.MutateRow(ctx, &btpb.MutateRowRequest{
		TableName: tblInfo.GetName(),
		RowKey:    []byte("row1"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_AddToCell_{AddToCell: &btpb.Mutation_AddToCell{
				FamilyName:      "sum",
				ColumnQualifier: &btpb.Value{Kind: &btpb.Value_RawValue{RawValue: []byte("col1")}},
				Timestamp:       &btpb.Value{Kind: &btpb.Value_RawTimestampMicros{RawTimestampMicros: 0}},
				Input:           &btpb.Value{Kind: &btpb.Value_IntValue{IntValue: 2}},
			}},
		}},
	})

	if err != nil {
		t.Fatal(err)
	}

	mock := &MockReadRowsServer{}
	err = s.ReadRows(&btpb.ReadRowsRequest{
		TableName: tblInfo.GetName(),
		Rows: &btpb.RowSet{
			RowKeys: [][]byte{
				[]byte("row1"),
			},
		}}, mock)
	if err != nil {
		t.Fatal(err)
	}
	got := mock.responses[0]

	if !bytes.Equal(got.Chunks[0].Value, binary.BigEndian.AppendUint64([]byte{}, 3)) {
		t.Error()
	}
}

func TestMutateRowsAggregate_MergeToCell(t *testing.T) {
	ctx := context.Background()

	s := &server{
		tables: make(map[string]*table),
	}

	tblInfo, err := populateTable(ctx, s)
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.ModifyColumnFamilies(ctx, &btapb.ModifyColumnFamiliesRequest{
		Name: tblInfo.Name,
		Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{{
			Id: "sum",
			Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Create{
				Create: &btapb.ColumnFamily{
					ValueType: &btapb.Type{
						Kind: &btapb.Type_AggregateType{
							AggregateType: &btapb.Type_Aggregate{
								InputType: &btapb.Type{
									Kind: &btapb.Type_Int64Type{},
								},
								Aggregator: &btapb.Type_Aggregate_Sum_{
									Sum: &btapb.Type_Aggregate_Sum{},
								},
							},
						},
					},
				},
			}},
		}})

	if err != nil {
		t.Fatal(err)
	}

	_, err = s.MutateRow(ctx, &btpb.MutateRowRequest{
		TableName: tblInfo.GetName(),
		RowKey:    []byte("row1"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_MergeToCell_{MergeToCell: &btpb.Mutation_MergeToCell{
				FamilyName:      "sum",
				ColumnQualifier: &btpb.Value{Kind: &btpb.Value_RawValue{RawValue: []byte("col1")}},
				Timestamp:       &btpb.Value{Kind: &btpb.Value_RawTimestampMicros{RawTimestampMicros: 0}},
				Input:           &btpb.Value{Kind: &btpb.Value_RawValue{RawValue: binary.BigEndian.AppendUint64([]byte{}, 1)}},
			}},
		}},
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = s.MutateRow(ctx, &btpb.MutateRowRequest{
		TableName: tblInfo.GetName(),
		RowKey:    []byte("row1"),
		Mutations: []*btpb.Mutation{{
			Mutation: &btpb.Mutation_MergeToCell_{MergeToCell: &btpb.Mutation_MergeToCell{
				FamilyName:      "sum",
				ColumnQualifier: &btpb.Value{Kind: &btpb.Value_RawValue{RawValue: []byte("col1")}},
				Timestamp:       &btpb.Value{Kind: &btpb.Value_RawTimestampMicros{RawTimestampMicros: 0}},
				Input:           &btpb.Value{Kind: &btpb.Value_RawValue{RawValue: binary.BigEndian.AppendUint64([]byte{}, 2)}},
			}},
		}},
	})

	if err != nil {
		t.Fatal(err)
	}

	mock := &MockReadRowsServer{}
	err = s.ReadRows(&btpb.ReadRowsRequest{
		TableName: tblInfo.GetName(),
		Rows: &btpb.RowSet{
			RowKeys: [][]byte{
				[]byte("row1"),
			},
		}}, mock)
	if err != nil {
		t.Fatal(err)
	}
	got := mock.responses[0]

	if !bytes.Equal(got.Chunks[0].Value, binary.BigEndian.AppendUint64([]byte{}, 3)) {
		t.Error()
	}
}

func Test_Mutation_DeleteFromColumn(t *testing.T) {
	ctx := context.Background()

	s := &server{
		tables: make(map[string]*table),
	}

	tblInfo, err := populateTable(ctx, s)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		in   *btpb.MutateRowRequest
		fail bool
	}{
		{
			in: &btpb.MutateRowRequest{
				TableName: tblInfo.Name,
				RowKey:    []byte("row"),
				Mutations: []*btpb.Mutation{{
					Mutation: &btpb.Mutation_DeleteFromColumn_{DeleteFromColumn: &btpb.Mutation_DeleteFromColumn{
						FamilyName:      "cf1",
						ColumnQualifier: []byte("col1"),
						TimeRange: &btpb.TimestampRange{
							StartTimestampMicros: 2000,
							EndTimestampMicros:   1000,
						},
					}},
				}},
			},
			fail: true,
		},
		{
			in: &btpb.MutateRowRequest{
				TableName: tblInfo.Name,
				RowKey:    []byte("row"),
				Mutations: []*btpb.Mutation{{
					Mutation: &btpb.Mutation_DeleteFromColumn_{DeleteFromColumn: &btpb.Mutation_DeleteFromColumn{
						FamilyName:      "cf2",
						ColumnQualifier: []byte("col2"),
						TimeRange: &btpb.TimestampRange{
							StartTimestampMicros: 1000,
							EndTimestampMicros:   2000,
						},
					}},
				}},
			},
			fail: false,
		},
		{
			in: &btpb.MutateRowRequest{
				TableName: tblInfo.Name,
				RowKey:    []byte("row"),
				Mutations: []*btpb.Mutation{{
					Mutation: &btpb.Mutation_DeleteFromColumn_{DeleteFromColumn: &btpb.Mutation_DeleteFromColumn{
						FamilyName:      "cf3",
						ColumnQualifier: []byte("col3"),
						TimeRange: &btpb.TimestampRange{
							StartTimestampMicros: 1000,
							EndTimestampMicros:   0,
						},
					}},
				}},
			},
			fail: false,
		},
		{
			in: &btpb.MutateRowRequest{
				TableName: tblInfo.Name,
				RowKey:    []byte("row"),
				Mutations: []*btpb.Mutation{{
					Mutation: &btpb.Mutation_DeleteFromColumn_{DeleteFromColumn: &btpb.Mutation_DeleteFromColumn{
						FamilyName:      "cf4",
						ColumnQualifier: []byte("col4"),
						TimeRange: &btpb.TimestampRange{
							StartTimestampMicros: 0,
							EndTimestampMicros:   1000,
						},
					}},
				}},
			},
			fail: true,
		},
	}
	for _, test := range tests {
		_, err = s.MutateRow(ctx, test.in)

		if err != nil && !test.fail {
			t.Errorf("expected passed got failure for : %v \n with err: %v", test.in, err)
		}

		if err == nil && test.fail {
			t.Errorf("expected failure got passed for : %v", test)
		}
	}
}

func TestFilterRow(t *testing.T) {
	row := &row{
		key: "row",
		families: map[string]*family{
			"fam": {
				name: "fam",
				cells: map[string][]cell{
					"col": {{ts: 1000, value: []byte("val")}},
				},
			},
		},
	}
	for _, test := range []struct {
		filter *btpb.RowFilter
		want   bool
	}{
		// The regexp-based filters perform whole-string, case-sensitive matches.
		{&btpb.RowFilter{Filter: &btpb.RowFilter_RowKeyRegexFilter{RowKeyRegexFilter: []byte("row")}}, true},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_RowKeyRegexFilter{RowKeyRegexFilter: []byte("ro")}}, false},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_RowKeyRegexFilter{RowKeyRegexFilter: []byte("ROW")}}, false},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_RowKeyRegexFilter{RowKeyRegexFilter: []byte("moo")}}, false},

		{&btpb.RowFilter{Filter: &btpb.RowFilter_FamilyNameRegexFilter{FamilyNameRegexFilter: "fam"}}, true},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_FamilyNameRegexFilter{FamilyNameRegexFilter: "f.*"}}, true},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_FamilyNameRegexFilter{FamilyNameRegexFilter: "[fam]+"}}, true},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_FamilyNameRegexFilter{FamilyNameRegexFilter: "fa"}}, false},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_FamilyNameRegexFilter{FamilyNameRegexFilter: "FAM"}}, false},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_FamilyNameRegexFilter{FamilyNameRegexFilter: "moo"}}, false},

		{&btpb.RowFilter{Filter: &btpb.RowFilter_ColumnQualifierRegexFilter{ColumnQualifierRegexFilter: []byte("col")}}, true},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_ColumnQualifierRegexFilter{ColumnQualifierRegexFilter: []byte("co")}}, false},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_ColumnQualifierRegexFilter{ColumnQualifierRegexFilter: []byte("COL")}}, false},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_ColumnQualifierRegexFilter{ColumnQualifierRegexFilter: []byte("moo")}}, false},

		{&btpb.RowFilter{Filter: &btpb.RowFilter_ValueRegexFilter{ValueRegexFilter: []byte("val")}}, true},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_ValueRegexFilter{ValueRegexFilter: []byte("va")}}, false},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_ValueRegexFilter{ValueRegexFilter: []byte("VAL")}}, false},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_ValueRegexFilter{ValueRegexFilter: []byte("moo")}}, false},

		{&btpb.RowFilter{Filter: &btpb.RowFilter_TimestampRangeFilter{TimestampRangeFilter: &btpb.TimestampRange{StartTimestampMicros: int64(0), EndTimestampMicros: int64(1000)}}}, false},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_TimestampRangeFilter{TimestampRangeFilter: &btpb.TimestampRange{StartTimestampMicros: int64(1000), EndTimestampMicros: int64(2001)}}}, true},
	} {
		got, err := filterRow(test.filter, row.copy())
		if err != nil {
			t.Errorf("%s: got unexpected error: %v", prototext.Format(test.filter), err)
		}
		if got != test.want {
			t.Errorf("%s: got %t, want %t", prototext.Format(test.filter), got, test.want)
		}
	}
}

func TestFilterRowWithErrors(t *testing.T) {
	row := &row{
		key: "row",
		families: map[string]*family{
			"fam": {
				name: "fam",
				cells: map[string][]cell{
					"col": {{ts: 1000, value: []byte("val")}},
				},
			},
		},
	}
	for _, test := range []struct {
		badRegex *btpb.RowFilter
	}{
		{&btpb.RowFilter{Filter: &btpb.RowFilter_RowKeyRegexFilter{RowKeyRegexFilter: []byte("[")}}},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_FamilyNameRegexFilter{FamilyNameRegexFilter: "["}}},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_ColumnQualifierRegexFilter{ColumnQualifierRegexFilter: []byte("[")}}},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_ValueRegexFilter{ValueRegexFilter: []byte("[")}}},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_Chain_{
			Chain: &btpb.RowFilter_Chain{
				Filters: []*btpb.RowFilter{
					{Filter: &btpb.RowFilter_ValueRegexFilter{ValueRegexFilter: []byte("[")}},
				},
			},
		}}},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_Condition_{
			Condition: &btpb.RowFilter_Condition{
				PredicateFilter: &btpb.RowFilter{Filter: &btpb.RowFilter_ValueRegexFilter{ValueRegexFilter: []byte("[")}},
			},
		}}},

		{&btpb.RowFilter{Filter: &btpb.RowFilter_RowSampleFilter{RowSampleFilter: 0.0}}}, // 0.0 is invalid.
		{&btpb.RowFilter{Filter: &btpb.RowFilter_RowSampleFilter{RowSampleFilter: 1.0}}}, // 1.0 is invalid.
	} {
		got, err := filterRow(test.badRegex, row.copy())
		if got != false {
			t.Errorf("%s: got true, want false", prototext.Format(test.badRegex))
		}
		if err == nil {
			t.Errorf("%s: got no error, want error", prototext.Format(test.badRegex))
		}
	}
}

func TestFilterRowWithRowSampleFilter(t *testing.T) {
	prev := randFloat
	randFloat = func() float64 { return 0.5 }
	defer func() { randFloat = prev }()
	for _, test := range []struct {
		p    float64
		want bool
	}{
		{0.1, false}, // Less than random float. Return no rows.
		{0.5, false}, // Equal to random float. Return no rows.
		{0.9, true},  // Greater than random float. Return all rows.
	} {
		got, err := filterRow(&btpb.RowFilter{Filter: &btpb.RowFilter_RowSampleFilter{RowSampleFilter: test.p}}, &row{})
		if err != nil {
			t.Fatalf("%f: %v", test.p, err)
		}
		if got != test.want {
			t.Errorf("%v: got %t, want %t", test.p, got, test.want)
		}
	}
}

func TestFilterRowWithBinaryColumnQualifier(t *testing.T) {
	rs := []byte{128, 128}
	row := &row{
		key: string(rs),
		families: map[string]*family{
			"fam": {
				name: "fam",
				cells: map[string][]cell{
					string(rs): {{ts: 1000, value: []byte("val")}},
				},
			},
		},
	}
	for _, test := range []struct {
		filter string
		want   bool
	}{
		{`\x80\x80`, true},      // succeeds, exact match
		{`\x80\x81`, false},     // fails
		{`\x80`, false},         // fails, because the regexp must match the entire input
		{`\x80*`, true},         // succeeds: 0 or more 128s
		{`[\x7f\x80]{2}`, true}, // succeeds: exactly two of either 127 or 128
		{`\C{2}`, true},         // succeeds: two bytes
	} {
		got, _ := filterRow(&btpb.RowFilter{Filter: &btpb.RowFilter_ColumnQualifierRegexFilter{ColumnQualifierRegexFilter: []byte(test.filter)}}, row.copy())
		if got != test.want {
			t.Errorf("%v: got %t, want %t", test.filter, got, test.want)
		}
	}
}

func TestFilterRowWithUnicodeColumnQualifier(t *testing.T) {
	rs := []byte("a§b")
	row := &row{
		key: string(rs),
		families: map[string]*family{
			"fam": {
				name: "fam",
				cells: map[string][]cell{
					string(rs): {{ts: 1000, value: []byte("val")}},
				},
			},
		},
	}
	for _, test := range []struct {
		filter string
		want   bool
	}{
		{`a§b`, true},        // succeeds, exact match
		{`a\xC2\xA7b`, true}, // succeeds, exact match
		{`a\xC2.+`, true},    // succeeds, prefix match
		{`a\xC2\C{2}`, true}, // succeeds, prefix match
		{`a\xC.+`, false},    // fails, prefix match, bad escape
		{`a§.+`, true},       // succeeds, prefix match
		{`.+§b`, true},       // succeeds, suffix match
		{`.§b`, true},        // succeeds
		{`a§c`, false},       // fails
		{`§b`, false},        // fails, because the regexp must match the entire input
		{`.*§.*`, true},      // succeeds: anything with a §
		{`.+§.+`, true},      // succeeds: anything with a § in the middle
		{`a\C{2}b`, true},    // succeeds: § is two bytes
		{`\C{4}`, true},      // succeeds: four bytes
	} {
		got, _ := filterRow(&btpb.RowFilter{Filter: &btpb.RowFilter_ColumnQualifierRegexFilter{ColumnQualifierRegexFilter: []byte(test.filter)}}, row.copy())
		if got != test.want {
			t.Errorf("%v: got %t, want %t", test.filter, got, test.want)
		}
	}
}

// Test that a single column qualifier with the interleave filter returns
// the correct result and not return every single row.
// See Issue https://github.com/googleapis/google-cloud-go/issues/1399
func TestFilterRowWithSingleColumnQualifier(t *testing.T) {
	ctx := context.Background()
	srv := &server{tables: make(map[string]*table)}

	tblReq := &btapb.CreateTableRequest{
		Parent:  "issue-1399",
		TableId: "table_id",
		Table: &btapb.Table{
			ColumnFamilies: map[string]*btapb.ColumnFamily{
				"cf": {},
			},
		},
	}
	tbl, err := srv.CreateTable(ctx, tblReq)
	if err != nil {
		t.Fatalf("Failed to create the table: %v", err)
	}

	entries := []struct {
		row   string
		value []byte
	}{
		{"row1", []byte{0x11}},
		{"row2", []byte{0x1a}},
		{"row3", []byte{'a'}},
		{"row4", []byte{'b'}},
	}

	for _, entry := range entries {
		req := &btpb.MutateRowRequest{
			TableName: tbl.Name,
			RowKey:    []byte(entry.row),
			Mutations: []*btpb.Mutation{{
				Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
					FamilyName:      "cf",
					ColumnQualifier: []byte("cq"),
					TimestampMicros: 1000,
					Value:           entry.value,
				}},
			}},
		}
		if _, err := srv.MutateRow(ctx, req); err != nil {
			t.Fatalf("Failed to insert entry %v into server: %v", entry, err)
		}
	}

	// After insertion now it is time for querying.
	req := &btpb.ReadRowsRequest{
		TableName: tbl.Name,
		Filter: &btpb.RowFilter{Filter: &btpb.RowFilter_Chain_{
			Chain: &btpb.RowFilter_Chain{Filters: []*btpb.RowFilter{{
				Filter: &btpb.RowFilter_Interleave_{
					Interleave: &btpb.RowFilter_Interleave{
						Filters: []*btpb.RowFilter{{Filter: &btpb.RowFilter_Condition_{
							Condition: &btpb.RowFilter_Condition{
								PredicateFilter: &btpb.RowFilter{Filter: &btpb.RowFilter_Chain_{
									Chain: &btpb.RowFilter_Chain{Filters: []*btpb.RowFilter{
										{
											Filter: &btpb.RowFilter_ValueRangeFilter{ValueRangeFilter: &btpb.ValueRange{
												StartValue: &btpb.ValueRange_StartValueClosed{
													StartValueClosed: []byte("a"),
												},
												EndValue: &btpb.ValueRange_EndValueClosed{EndValueClosed: []byte("a")},
											}},
										},
										{Filter: &btpb.RowFilter_PassAllFilter{PassAllFilter: true}},
									}},
								}},
								TrueFilter: &btpb.RowFilter{Filter: &btpb.RowFilter_PassAllFilter{PassAllFilter: true}},
							},
						}},
							{Filter: &btpb.RowFilter_BlockAllFilter{BlockAllFilter: true}},
						},
					},
				},
			},
				{Filter: &btpb.RowFilter_PassAllFilter{PassAllFilter: true}},
			}},
		}},
	}

	rrss := new(MockReadRowsServer)
	if err := srv.ReadRows(req, rrss); err != nil {
		t.Fatalf("Failed to read rows: %v", err)
	}

	if g, w := len(rrss.responses), 1; g != w {
		t.Fatalf("Results/Streamed chunks mismatch:: got %d want %d", g, w)
	}

	got := rrss.responses[0]
	// Only row3 should be matched.
	want := &btpb.ReadRowsResponse{
		Chunks: []*btpb.ReadRowsResponse_CellChunk{
			{
				RowKey:          []byte("row3"),
				FamilyName:      &wrapperspb.StringValue{Value: "cf"},
				Qualifier:       &wrapperspb.BytesValue{Value: []byte("cq")},
				TimestampMicros: 1000,
				Value:           []byte("a"),
				RowStatus: &btpb.ReadRowsResponse_CellChunk_CommitRow{
					CommitRow: true,
				},
			},
		},
	}
	if diff := cmp.Diff(got, want, cmp.Comparer(proto.Equal)); diff != "" {
		t.Fatalf("Response mismatch: got: + want -\n%s", diff)
	}
}

func TestValueFilterRowWithAlternationInRegex(t *testing.T) {
	// Test that regex alternation is applied properly.
	// See Issue https://github.com/googleapis/google-cloud-go/issues/1499
	ctx := context.Background()
	srv := &server{tables: make(map[string]*table)}

	tblReq := &btapb.CreateTableRequest{
		Parent:  "issue-1499",
		TableId: "table_id",
		Table: &btapb.Table{
			ColumnFamilies: map[string]*btapb.ColumnFamily{
				"cf": {},
			},
		},
	}
	tbl, err := srv.CreateTable(ctx, tblReq)
	if err != nil {
		t.Fatalf("Failed to create the table: %v", err)
	}

	entries := []struct {
		row   string
		value []byte
	}{
		{"row1", []byte("")},
		{"row2", []byte{'x'}},
		{"row3", []byte{'a'}},
		{"row4", []byte{'m'}},
	}

	for _, entry := range entries {
		req := &btpb.MutateRowRequest{
			TableName: tbl.Name,
			RowKey:    []byte(entry.row),
			Mutations: []*btpb.Mutation{{
				Mutation: &btpb.Mutation_SetCell_{SetCell: &btpb.Mutation_SetCell{
					FamilyName:      "cf",
					ColumnQualifier: []byte("cq"),
					TimestampMicros: 1000,
					Value:           entry.value,
				}},
			}},
		}
		if _, err := srv.MutateRow(ctx, req); err != nil {
			t.Fatalf("Failed to insert entry %v into server: %v", entry, err)
		}
	}

	// After insertion now it is time for querying.
	req := &btpb.ReadRowsRequest{
		TableName: tbl.Name,
		Rows:      &btpb.RowSet{},
		Filter: &btpb.RowFilter{
			Filter: &btpb.RowFilter_ValueRegexFilter{
				ValueRegexFilter: []byte("|a"),
			},
		},
	}

	rrss := new(MockReadRowsServer)
	if err := srv.ReadRows(req, rrss); err != nil {
		t.Fatalf("Failed to read rows: %v", err)
	}

	var gotChunks []*btpb.ReadRowsResponse_CellChunk
	for _, res := range rrss.responses {
		gotChunks = append(gotChunks, res.Chunks...)
	}

	// Only row1 "" and row3 "a" should be matched.
	wantChunks := []*btpb.ReadRowsResponse_CellChunk{
		{
			RowKey:          []byte("row1"),
			FamilyName:      &wrapperspb.StringValue{Value: "cf"},
			Qualifier:       &wrapperspb.BytesValue{Value: []byte("cq")},
			TimestampMicros: 1000,
			Value:           []byte(""),
			RowStatus: &btpb.ReadRowsResponse_CellChunk_CommitRow{
				CommitRow: true,
			},
		},
		{
			RowKey:          []byte("row3"),
			FamilyName:      &wrapperspb.StringValue{Value: "cf"},
			Qualifier:       &wrapperspb.BytesValue{Value: []byte("cq")},
			TimestampMicros: 1000,
			Value:           []byte("a"),
			RowStatus: &btpb.ReadRowsResponse_CellChunk_CommitRow{
				CommitRow: true,
			},
		},
	}
	if diff := cmp.Diff(gotChunks, wantChunks, cmp.Comparer(proto.Equal)); diff != "" {
		t.Fatalf("Response chunks mismatch: got: + want -\n%s", diff)
	}
}

func TestMutateRowEmptyMutationErrors(t *testing.T) {
	srv := &server{tables: make(map[string]*table)}
	ctx := context.Background()
	req := &btpb.MutateRowRequest{
		TableName: "mytable",
		RowKey:    []byte("r"),
		Mutations: []*btpb.Mutation{},
	}

	resp, err := srv.MutateRow(ctx, req)
	if resp != nil ||
		fmt.Sprint(err) !=
			"rpc error: code = InvalidArgument"+
				" desc = No mutations provided" {
		t.Fatalf("Failed to error %s", err)
	}
}

type bigtableTestingMutateRowsServer struct {
	grpc.ServerStream
}

func (x *bigtableTestingMutateRowsServer) Send(m *btpb.MutateRowsResponse) error {
	return nil
}

func TestMutateRowsEmptyMutationErrors(t *testing.T) {
	srv := &server{tables: make(map[string]*table)}
	req := &btpb.MutateRowsRequest{
		TableName: "mytable",
		Entries: []*btpb.MutateRowsRequest_Entry{
			{Mutations: []*btpb.Mutation{}},
			{Mutations: []*btpb.Mutation{}},
		},
	}

	err := srv.MutateRows(req, &bigtableTestingMutateRowsServer{})
	if fmt.Sprint(err) !=
		"rpc error: code = InvalidArgument "+
			"desc = No mutations provided" {
		t.Fatalf("Failed to error %s", err)
	}
}

func TestMutateRowEmptyRowKeyErrors(t *testing.T) {
	srv := &server{tables: make(map[string]*table)}
	ctx := context.Background()

	const tableID = "mytable"
	tblReq := &btapb.CreateTableRequest{
		Parent:  "cluster",
		TableId: tableID,
		Table: &btapb.Table{
			ColumnFamilies: map[string]*btapb.ColumnFamily{"cf": {}},
		},
	}
	if _, err := srv.CreateTable(ctx, tblReq); err != nil {
		t.Fatalf("Failed to create the table: %v", err)
	}

	const name = "cluster/tables/" + tableID
	req := &btpb.MutateRowRequest{
		TableName: name,
		RowKey:    []byte(""),
		Mutations: []*btpb.Mutation{
			{
				Mutation: &btpb.Mutation_SetCell_{
					SetCell: &btpb.Mutation_SetCell{
						FamilyName:      "cf",
						ColumnQualifier: []byte("col"),
						TimestampMicros: 1000,
						Value:           []byte("hello, world!"),
					},
				},
			},
		},
	}

	resp, err := srv.MutateRow(ctx, req)
	if resp != nil || err == nil || err.Error() !=
		"rpc error: code = InvalidArgument"+
			" desc = Row keys must be non-empty" {
		t.Fatalf("Failed to produce the expected error: %s", err)
	}
}

func TestFilterRowCellsPerRowLimitFilterTruthiness(t *testing.T) {
	row := &row{
		key: "row",
		families: map[string]*family{
			"fam": {
				name: "fam",
				cells: map[string][]cell{
					"col1": {{ts: 1000, value: []byte("val2")}},
					"col2": {
						{ts: 1000, value: []byte("val2")},
						{ts: 1000, value: []byte("val3")},
					},
				},
				colNames: []string{"col1", "col2"},
			},
		},
	}
	for _, test := range []struct {
		filter *btpb.RowFilter
		want   bool
	}{
		// The regexp-based filters perform whole-string, case-sensitive matches.
		{&btpb.RowFilter{Filter: &btpb.RowFilter_CellsPerRowOffsetFilter{CellsPerRowOffsetFilter: 1}}, true},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_CellsPerRowOffsetFilter{CellsPerRowOffsetFilter: 2}}, true},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_CellsPerRowOffsetFilter{CellsPerRowOffsetFilter: 3}}, false},
		{&btpb.RowFilter{Filter: &btpb.RowFilter_CellsPerRowOffsetFilter{CellsPerRowOffsetFilter: 4}}, false},
	} {
		got, err := filterRow(test.filter, row.copy())
		if err != nil {
			t.Errorf("%s: got unexpected error: %v", prototext.Format(test.filter), err)
		}
		if got != test.want {
			t.Errorf("%s: got %t, want %t", prototext.Format(test.filter), got, test.want)
		}
	}
}

func TestAuthorizedViewApis(t *testing.T) {
	s := &server{
		tables: make(map[string]*table),
	}
	ctx := context.Background()
	_, err := populateTable(ctx, s)
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.CreateAuthorizedView(ctx, &btapb.CreateAuthorizedViewRequest{})
	if fmt.Sprint(err) !=
		"rpc error: code = Unimplemented desc = the emulator does not currently support authorized views" {
		t.Fatalf("Failed to error %s", err)
	}

	_, err = s.GetAuthorizedView(ctx, &btapb.GetAuthorizedViewRequest{})
	if fmt.Sprint(err) !=
		"rpc error: code = Unimplemented desc = the emulator does not currently support authorized views" {
		t.Fatalf("Failed to error %s", err)
	}

	_, err = s.ListAuthorizedViews(ctx, &btapb.ListAuthorizedViewsRequest{})
	if fmt.Sprint(err) !=
		"rpc error: code = Unimplemented desc = the emulator does not currently support authorized views" {
		t.Fatalf("Failed to error %s", err)
	}

	_, err = s.UpdateAuthorizedView(ctx, &btapb.UpdateAuthorizedViewRequest{})
	if fmt.Sprint(err) !=
		"rpc error: code = Unimplemented desc = the emulator does not currently support authorized views" {
		t.Fatalf("Failed to error %s", err)
	}

	_, err = s.DeleteAuthorizedView(ctx, &btapb.DeleteAuthorizedViewRequest{Name: "av_name"})
	if fmt.Sprint(err) !=
		"rpc error: code = Unimplemented desc = the emulator does not currently support authorized views" {
		t.Fatalf("Failed to error %s", err)
	}
}

func TestUpdateGCPolicyOnAggregateColumnFamily(t *testing.T) {
	ctx := context.Background()

	s := &server{
		tables: make(map[string]*table),
	}

	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{
		Parent:  "cluster",
		TableId: "t",
		Table: &btapb.Table{
			ColumnFamilies: map[string]*btapb.ColumnFamily{
				"sum": {
					ValueType: &btapb.Type{
						Kind: &btapb.Type_AggregateType{
							AggregateType: &btapb.Type_Aggregate{
								InputType: &btapb.Type{
									Kind: &btapb.Type_Int64Type{},
								},
								Aggregator: &btapb.Type_Aggregate_Sum_{
									Sum: &btapb.Type_Aggregate_Sum{},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if tblInfo.ColumnFamilies["sum"].
		GetValueType().
		GetAggregateType().
		GetSum() == nil {
		t.Fatal("Unexpected aggregate column family type at start of test")
	}

	if tblInfo.ColumnFamilies["sum"].
		GetGcRule().
		GetMaxNumVersions() == 1 {
		t.Fatal("Unexpected GC policy state at start of test")
	}

	tblInfo, err = s.ModifyColumnFamilies(ctx, &btapb.ModifyColumnFamiliesRequest{
		Name: tblInfo.Name,
		Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{
			{
				Id: "sum",
				// UpdateMask intentionally left empty, which the server will
				// implicitly interpret as a gc_rule update.
				Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Update{
					Update: &btapb.ColumnFamily{
						GcRule: &btapb.GcRule{
							Rule: &btapb.GcRule_MaxNumVersions{
								MaxNumVersions: 1,
							},
						},
						// HACK: Intentionally include an invalid type
						// update, which should be ignored since it isn't
						// present in the UpdateMask.
						ValueType: &btapb.Type{
							Kind: &btapb.Type_AggregateType{
								AggregateType: &btapb.Type_Aggregate{
									InputType: &btapb.Type{
										Kind: &btapb.Type_Int64Type{},
									},
									Aggregator: &btapb.Type_Aggregate_Max_{
										Max: &btapb.Type_Aggregate_Max{},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if tblInfo.ColumnFamilies["sum"].
		GetValueType().
		GetAggregateType().
		GetSum() == nil {
		t.Fatal("Aggregate type was updated when it should not have been")
	}

	if tblInfo.ColumnFamilies["sum"].
		GetGcRule().
		GetMaxNumVersions() != 1 {
		t.Fatal("GC policy was not updated when it should have been")
	}

	tblInfo, err = s.ModifyColumnFamilies(ctx, &btapb.ModifyColumnFamiliesRequest{
		Name: tblInfo.Name,
		Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{
			{
				Id: "sum",
				// Including UpdateMask in the request this time.
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{"gc_rule"},
				},
				Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Update{
					Update: &btapb.ColumnFamily{
						GcRule: &btapb.GcRule{
							Rule: &btapb.GcRule_MaxNumVersions{
								MaxNumVersions: 2,
							},
						},
						// HACK: Intentionally including an invalid type
						// update, which should be ignored since it isn't
						// present in the UpdateMask.
						ValueType: &btapb.Type{
							Kind: &btapb.Type_AggregateType{
								AggregateType: &btapb.Type_Aggregate{
									InputType: &btapb.Type{
										Kind: &btapb.Type_Int64Type{},
									},
									Aggregator: &btapb.Type_Aggregate_Max_{
										Max: &btapb.Type_Aggregate_Max{},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if tblInfo.ColumnFamilies["sum"].
		GetValueType().
		GetAggregateType().
		GetSum() == nil {
		t.Fatal("Aggregate type was updated when it should not have been")
	}

	if tblInfo.ColumnFamilies["sum"].
		GetGcRule().
		GetMaxNumVersions() != 2 {
		t.Fatal("GC policy was not updated when it should have been")
	}
}

func TestCannotUpdateTypeOfAggregateColumnFamily(t *testing.T) {
	ctx := context.Background()

	s := &server{
		tables: make(map[string]*table),
	}

	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{
		Parent:  "cluster",
		TableId: "t",
		Table: &btapb.Table{
			ColumnFamilies: map[string]*btapb.ColumnFamily{
				"sum": {
					ValueType: &btapb.Type{
						Kind: &btapb.Type_AggregateType{
							AggregateType: &btapb.Type_Aggregate{
								InputType: &btapb.Type{
									Kind: &btapb.Type_Int64Type{},
								},
								Aggregator: &btapb.Type_Aggregate_Sum_{
									Sum: &btapb.Type_Aggregate_Sum{},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if tblInfo.ColumnFamilies["sum"].
		GetValueType().
		GetAggregateType().
		GetSum() == nil {
		t.Fatal("Unexpected aggregate column family type at start of test")
	}

	_, err = s.ModifyColumnFamilies(ctx, &btapb.ModifyColumnFamiliesRequest{
		Name: tblInfo.Name,
		Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{
			{
				Id: "sum",
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{"value_type"},
				},
				Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Update{
					Update: &btapb.ColumnFamily{
						ValueType: &btapb.Type{
							Kind: &btapb.Type_AggregateType{
								AggregateType: &btapb.Type_Aggregate{
									InputType: &btapb.Type{
										Kind: &btapb.Type_Int64Type{},
									},
									Aggregator: &btapb.Type_Aggregate_Max_{
										Max: &btapb.Type_Aggregate_Max{},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err == nil {
		t.Fatal("ModifyColumnFamilies was supposed to return an error, but it did not")
	}

	tblInfo, err = s.GetTable(ctx, &btapb.GetTableRequest{Name: tblInfo.Name})
	if err != nil {
		t.Fatal(err)
	}

	if tblInfo.ColumnFamilies["sum"].
		GetValueType().
		GetAggregateType().
		GetSum() == nil {
		t.Fatal("Aggregate type was updated when it should not have been")
	}
}

func TestInvalidUpdateMaskInColumnFamilyUpdate(t *testing.T) {
	ctx := context.Background()

	s := &server{
		tables: make(map[string]*table),
	}

	tblInfo, err := s.CreateTable(ctx, &btapb.CreateTableRequest{
		Parent:  "cluster",
		TableId: "t",
		Table: &btapb.Table{
			ColumnFamilies: map[string]*btapb.ColumnFamily{
				"sum": {
					ValueType: &btapb.Type{
						Kind: &btapb.Type_AggregateType{
							AggregateType: &btapb.Type_Aggregate{
								InputType: &btapb.Type{
									Kind: &btapb.Type_Int64Type{},
								},
								Aggregator: &btapb.Type_Aggregate_Sum_{
									Sum: &btapb.Type_Aggregate_Sum{},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if tblInfo.ColumnFamilies["sum"].
		GetGcRule().
		GetMaxNumVersions() == 1 {
		t.Fatal("Unexpected GC policy state at start of test")
	}

	_, err = s.ModifyColumnFamilies(ctx, &btapb.ModifyColumnFamiliesRequest{
		Name: tblInfo.Name,
		Modifications: []*btapb.ModifyColumnFamiliesRequest_Modification{
			{
				Id: "sum",
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{"bad", "gc_rule"},
				},
				Mod: &btapb.ModifyColumnFamiliesRequest_Modification_Update{
					Update: &btapb.ColumnFamily{
						GcRule: &btapb.GcRule{
							Rule: &btapb.GcRule_MaxNumVersions{
								MaxNumVersions: 1,
							},
						},
					},
				},
			},
		},
	})
	if err == nil {
		t.Fatal("ModifyColumnFamilies was supposed to return an error, but it did not")
	}

	tblInfo, err = s.GetTable(ctx, &btapb.GetTableRequest{Name: tblInfo.Name})
	if err != nil {
		t.Fatal(err)
	}

	if tblInfo.ColumnFamilies["sum"].
		GetGcRule().
		GetMaxNumVersions() == 1 {
		t.Fatal("GC policy was updated when it should not have been")
	}
}
