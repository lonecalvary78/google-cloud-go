// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.7
// source: google/cloud/bigquery/storage/v1beta1/table_reference.proto

package storagepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Table reference that includes just the 3 strings needed to identify a table.
type TableReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The assigned project ID of the project.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The ID of the dataset in the above project.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The ID of the table in the above dataset.
	TableId string `protobuf:"bytes,3,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
}

func (x *TableReference) Reset() {
	*x = TableReference{}
	mi := &file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableReference) ProtoMessage() {}

func (x *TableReference) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableReference.ProtoReflect.Descriptor instead.
func (*TableReference) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDescGZIP(), []int{0}
}

func (x *TableReference) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *TableReference) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *TableReference) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

// All fields in this message optional.
type TableModifiers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The snapshot time of the table. If not set, interpreted as now.
	SnapshotTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=snapshot_time,json=snapshotTime,proto3" json:"snapshot_time,omitempty"`
}

func (x *TableModifiers) Reset() {
	*x = TableModifiers{}
	mi := &file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableModifiers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableModifiers) ProtoMessage() {}

func (x *TableModifiers) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableModifiers.ProtoReflect.Descriptor instead.
func (*TableModifiers) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDescGZIP(), []int{1}
}

func (x *TableModifiers) GetSnapshotTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SnapshotTime
	}
	return nil
}

var File_google_cloud_bigquery_storage_v1beta1_table_reference_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0x51, 0x0a, 0x0e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x42, 0x85, 0x01, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x13, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70,
	0x62, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDescData = file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDesc
)

func file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDescData)
	})
	return file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDescData
}

var file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_goTypes = []any{
	(*TableReference)(nil),        // 0: google.cloud.bigquery.storage.v1beta1.TableReference
	(*TableModifiers)(nil),        // 1: google.cloud.bigquery.storage.v1beta1.TableModifiers
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_depIdxs = []int32{
	2, // 0: google.cloud.bigquery.storage.v1beta1.TableModifiers.snapshot_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_init() }
func file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_init() {
	if File_google_cloud_bigquery_storage_v1beta1_table_reference_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_depIdxs,
		MessageInfos:      file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_storage_v1beta1_table_reference_proto = out.File
	file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_rawDesc = nil
	file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_goTypes = nil
	file_google_cloud_bigquery_storage_v1beta1_table_reference_proto_depIdxs = nil
}
