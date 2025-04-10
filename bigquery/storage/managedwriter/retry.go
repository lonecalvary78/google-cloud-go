// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package managedwriter

import (
	"errors"
	"io"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/googleapis/gax-go/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	defaultRetryAttempts = 4
)

// This retry predicate is used for higher level retries, enqueing appends onto to a bidi
// channel and evaluating whether an append should be retried (re-enqueued).
func retryPredicate(err error) bool {
	if err == nil {
		return false
	}

	s, ok := status.FromError(err)
	// non-status based error conditions.
	if !ok {
		// EOF can happen in the case of connection close.
		if errors.Is(err, io.EOF) {
			return true
		}
		// All other non-status errors are treated as non-retryable (including context errors).
		return false
	}
	switch s.Code() {
	case codes.Aborted,
		codes.Canceled,
		codes.DeadlineExceeded,
		codes.FailedPrecondition,
		codes.Internal,
		codes.Unavailable:
		return true
	case codes.ResourceExhausted:
		if strings.HasPrefix(s.Message(), "Exceeds 'AppendRows throughput' quota") {
			// Note: internal b/246031522 opened to give this a structured error
			// and avoid string parsing.  Should be a QuotaFailure or similar.
			return true
		}
	}
	return false
}

// unaryRetryer is for retrying a unary-style operation, like (re)-opening the bidi connection.
type unaryRetryer struct {
	bo gax.Backoff
}

func (ur *unaryRetryer) Retry(err error) (time.Duration, bool) {
	shouldRetry := retryPredicate(err)
	return ur.bo.Pause(), shouldRetry
}

// statelessRetryer is used for backing off within a continuous process, like processing the responses
// from the receive side of the bidi stream.  An individual item in that process has a notion of an attempt
// count, and we use maximum retries as a way of evicting bad items.
type statelessRetryer struct {
	mu sync.Mutex // guards r
	r  *rand.Rand

	minBackoff  time.Duration
	jitter      time.Duration
	maxAttempts int
}

func newStatelessRetryer() *statelessRetryer {
	return &statelessRetryer{
		r:           rand.New(rand.NewSource(time.Now().UnixNano())),
		minBackoff:  50 * time.Millisecond,
		jitter:      time.Second,
		maxAttempts: defaultRetryAttempts,
	}
}

func (sr *statelessRetryer) pause() time.Duration {
	jitter := sr.jitter.Nanoseconds()
	if jitter > 0 {
		sr.mu.Lock()
		jitter = sr.r.Int63n(jitter)
		sr.mu.Unlock()
	}
	pause := sr.minBackoff.Nanoseconds() + jitter
	return time.Duration(pause)
}

func (sr *statelessRetryer) Retry(err error, attemptCount int) (time.Duration, bool) {
	if attemptCount >= sr.maxAttempts {
		return 0, false
	}
	if retryPredicate(err) {
		return sr.pause(), true
	}
	return 0, false
}

// shouldReconnect is akin to a retry predicate, in that it evaluates whether we should force
// our bidi stream to close/reopen based on the responses error.  Errors here signal that no
// further appends will succeed.
func shouldReconnect(err error) bool {

	// io.EOF is the typical not connected signal.
	if errors.Is(err, io.EOF) {
		return true
	}
	// Backend responses that trigger reconnection on send.
	reconnectCodes := []codes.Code{
		codes.Aborted,
		codes.Canceled,
		codes.Unavailable,
		codes.DeadlineExceeded,
	}
	if s, ok := status.FromError(err); ok {
		for _, c := range reconnectCodes {
			if s.Code() == c {
				return true
			}
		}
	}
	return false
}
