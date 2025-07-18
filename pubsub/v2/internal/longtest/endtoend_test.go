// Copyright 2025 Google LLC
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

package longtest_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"cloud.google.com/go/internal/testutil"
	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	timeout                 = time.Minute * 10
	ackDeadline             = time.Second * 10
	nMessages               = 1e4
	acceptableDupPercentage = 1
	numAcceptableDups       = int(nMessages * acceptableDupPercentage / 100)
	resourcePrefix          = "endtoend"
)

// The end-to-end pumps many messages into a topic and tests that they are all
// delivered to each subscription for the topic. It also tests that messages
// are not unexpectedly redelivered.
func TestEndToEnd_Dupes(t *testing.T) {
	t.Skip("https://github.com/googleapis/google-cloud-go/issues/1752")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, topic, cleanup := prepareEndToEndTest(ctx, t)
	defer cleanup()
	subPrefix := fmt.Sprintf("%s-%d", resourcePrefix, time.Now().UnixNano())

	// Two subscriptions to the same topic.
	var err error
	var subs [2]*pubsub.Subscription
	for i := 0; i < len(subs); i++ {
		subs[i], err = client.CreateSubscription(ctx, fmt.Sprintf("%s-%d", subPrefix, i), pubsub.SubscriptionConfig{
			Topic:       topic,
			AckDeadline: ackDeadline,
		})
		if err != nil {
			t.Fatalf("CreateSub error: %v", err)
		}
		defer subs[i].Delete(ctx)
	}

	err = publish(ctx, topic, nMessages)
	topic.Stop()
	if err != nil {
		t.Fatalf("publish: %v", err)
	}

	// recv provides an indication that messages are still arriving.
	recv := make(chan struct{})
	// We have two subscriptions to our topic.
	// Each subscription will get a copy of each published message.
	var wg sync.WaitGroup
	cctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	consumers := []*consumer{
		{
			counts:    make(map[string]int),
			recv:      recv,
			durations: []time.Duration{time.Hour},
			done:      make(chan struct{}),
		},
		{
			counts:    make(map[string]int),
			recv:      recv,
			durations: []time.Duration{ackDeadline, ackDeadline, ackDeadline / 2, ackDeadline / 2, time.Hour},
			done:      make(chan struct{}),
		},
	}
	for i, con := range consumers {
		con := con
		sub := subs[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			con.consume(ctx, t, sub)
		}()
	}
	// Wait for a while after the last message before declaring quiescence.
	// We wait a multiple of the ack deadline, for two reasons:
	// 1. To detect if messages are redelivered after having their ack
	//    deadline extended.
	// 2. To wait for redelivery of messages that were en route when a Receive
	//    is canceled. This can take considerably longer than the ack deadline.
	quiescenceDur := ackDeadline * 6
	quiescenceTimer := time.NewTimer(quiescenceDur)

loop:
	for {
		select {
		case <-recv:
			// Reset timer so we wait quiescenceDur after the last message.
			// See https://godoc.org/time#Timer.Reset for why the Stop
			// and channel drain are necessary.
			if !quiescenceTimer.Stop() {
				<-quiescenceTimer.C
			}
			quiescenceTimer.Reset(quiescenceDur)

		case <-quiescenceTimer.C:
			cancel()
			log.Println("quiesced")
			break loop

		case <-cctx.Done():
			t.Fatal("timed out")
		}
	}
	wg.Wait()
	close(recv)
	for i, con := range consumers {
		var numDups int
		var zeroes int
		for _, v := range con.counts {
			if v == 0 {
				zeroes++
			}
			numDups += v - 1
		}

		if zeroes > 0 {
			t.Errorf("Consumer %d: %d messages never arrived", i, zeroes)
		} else if numDups > numAcceptableDups {
			t.Errorf("Consumer %d: Willing to accept %d dups (%v%% duplicated of %d messages), but got %d", i, numAcceptableDups, acceptableDupPercentage, int(nMessages), numDups)
		}
	}

	for i, con := range consumers {
		select {
		case <-con.done:
		case <-time.After(15 * time.Second):
			t.Fatalf("timed out waiting for consumer %d to finish", i)
		}
	}
}

func TestEndToEnd_LongProcessingTime(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, topic, cleanup := prepareEndToEndTest(ctx, t)
	defer cleanup()
	subPrefix := fmt.Sprintf("%s-%d", resourcePrefix, time.Now().UnixNano())

	// Two subscriptions to the same topic.
	sub, err := client.CreateSubscription(ctx, subPrefix+"-00", pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: ackDeadline,
	})
	if err != nil {
		t.Fatalf("CreateSub error: %v", err)
	}
	defer sub.Delete(ctx)

	// Tests the issue found in https://github.com/googleapis/google-cloud-go/issues/1247.
	sub.ReceiveSettings.Synchronous = true
	sub.ReceiveSettings.MaxOutstandingMessages = 500

	err = publish(ctx, topic, 500)
	topic.Stop()
	if err != nil {
		t.Fatalf("publish: %v", err)
	}

	// recv provides an indication that messages are still arriving.
	recv := make(chan struct{})
	consumer := consumer{
		counts:    make(map[string]int),
		recv:      recv,
		durations: []time.Duration{time.Hour},
		processingDelay: func() time.Duration {
			return time.Duration(1+rand.Int63n(120)) * time.Second
		},
		done: make(chan struct{}),
	}
	go consumer.consume(ctx, t, sub)
	// Wait for a while after the last message before declaring quiescence.
	// We wait a multiple of the ack deadline, for two reasons:
	// 1. To detect if messages are redelivered after having their ack
	//    deadline extended.
	// 2. To wait for redelivery of messages that were en route when a Receive
	//    is canceled. This can take considerably longer than the ack deadline.
	quiescenceDur := 12 * ackDeadline
	quiescenceTimer := time.NewTimer(quiescenceDur)
loop:
	for {
		select {
		case <-recv:
			// Reset timer so we wait quiescenceDur after the last message.
			// See https://godoc.org/time#Timer.Reset for why the Stop
			// and channel drain are necessary.
			if !quiescenceTimer.Stop() {
				<-quiescenceTimer.C
			}
			quiescenceTimer.Reset(quiescenceDur)

		case <-quiescenceTimer.C:
			cancel()
			log.Println("quiesced")
			break loop

		case <-ctx.Done():
			t.Fatal("timed out")
		}
	}
	close(recv)
	var numDups int
	var zeroes int
	for _, v := range consumer.counts {
		if v == 0 {
			zeroes++
		}
		numDups += v - 1
	}

	if zeroes > 0 {
		t.Errorf("%d messages never arrived", zeroes)
	} else if numDups > numAcceptableDups {
		t.Errorf("Willing to accept %d dups (%v duplicated of %d messages), but got %d", numAcceptableDups, acceptableDupPercentage, int(nMessages), numDups)
	}

	select {
	case <-consumer.done:
	case <-time.After(15 * time.Second):
		t.Fatal("timed out waiting for consumer to finish")
	}
}

// publish publishes n messages to topic.
func publish(ctx context.Context, topic *pubsub.Topic, n int) error {
	var rs []*pubsub.PublishResult
	for i := 0; i < n; i++ {
		m := &pubsub.Message{Data: []byte(fmt.Sprintf("msg %d", i))}
		rs = append(rs, topic.Publish(ctx, m))
	}
	for _, r := range rs {
		_, err := r.Get(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// consumer consumes messages according to its configuration.
type consumer struct {
	// A consumer will spin out a Receive for each duration, which will be
	// canceled after each duration and the next one spun up. For example, if
	// there are 5 3 second durations, then there will be 5 3 second Receives.
	durations []time.Duration

	// A value is sent to recv each time process is called.
	recv chan struct{}

	// How long to wait for before acking.
	processingDelay func() time.Duration

	mu         sync.Mutex
	counts     map[string]int // msgID: recvdAmt
	totalRecvd int

	// Done consuming.
	done chan struct{}
}

// consume reads messages from a subscription, and keeps track of what it receives in mc.
// After consume returns, the caller should wait on wg to ensure that no more updates to mc will be made.
func (c *consumer) consume(ctx context.Context, t *testing.T, sub *pubsub.Subscription) {
	defer close(c.done)
	for _, dur := range c.durations {
		ctx2, cancel := context.WithTimeout(ctx, dur)
		defer cancel()
		id := sub.String()[len(sub.String())-1:]
		t.Logf("%s: start receive", id)
		prev := c.totalRecvd
		err := sub.Receive(ctx2, c.process)
		t.Logf("%s: end receive; read %d", id, c.totalRecvd-prev)
		if err != nil && !errors.Is(err, context.Canceled) {
			panic(err)
		}
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

// process handles a message and records it in mc.
func (c *consumer) process(_ context.Context, m *pubsub.Message) {
	c.mu.Lock()
	c.counts[m.ID]++
	c.totalRecvd++
	c.mu.Unlock()
	c.recv <- struct{}{}

	var delay time.Duration
	if c.processingDelay == nil {
		delay = time.Duration(rand.Intn(int(ackDeadline * 3)))
	} else {
		delay = c.processingDelay()
	}

	// Simulate time taken to process m, while continuing to process more messages.
	// Some messages will need to have their ack deadline extended due to this delay.
	time.AfterFunc(delay, func() {
		m.Ack()
	})
}

// Remember to call cleanup!
func prepareEndToEndTest(ctx context.Context, t *testing.T) (*pubsub.Client, *pubsub.Topic, func()) {
	if testing.Short() {
		t.Skip("Integration tests skipped in short mode")
	}
	ts := testutil.TokenSource(ctx, pubsub.ScopePubSub, pubsub.ScopeCloudPlatform)
	if ts == nil {
		t.Skip("Integration tests skipped. See CONTRIBUTING.md for details")
	}

	now := time.Now()
	topicName := fmt.Sprintf("%s-%d", resourcePrefix, now.UnixNano())

	client, err := pubsub.NewClient(ctx, testutil.ProjID(), option.WithTokenSource(ts))
	if err != nil {
		t.Fatalf("Creating client error: %v", err)
	}

	// Don't stop the test if cleanup failed.
	if err := cleanupSubscription(ctx, client); err != nil {
		t.Logf("Pre-test subscription cleanup failed: %v", err)
	}
	if err := cleanupTopic(ctx, client); err != nil {
		t.Logf("Pre-test topic cleanup failed: %v", err)
	}

	var topic *pubsub.Topic
	if topic, err = client.CreateTopic(ctx, topicName); err != nil {
		t.Fatalf("CreateTopic error: %v", err)
	}

	return client, topic, func() {
		topic.Delete(ctx)
		client.Close()
	}
}

// cleanupTopic deletes stale testing topics.
func cleanupTopic(ctx context.Context, client *pubsub.Client) error {
	if testing.Short() {
		return nil // Don't clean up in short mode.
	}
	// Delete topics which were	created a while ago.
	const expireAge = 24 * time.Hour

	it := client.Topics(ctx)
	for {
		t, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return err
		}
		// Take timestamp from id.
		tID := t.ID()
		p := strings.Split(tID, "-")

		// Only delete resources created from the endtoend test.
		// Otherwise, this will affect other tests running midflight.
		if p[0] == resourcePrefix {
			tCreated := p[len(p)-1]
			timestamp, err := strconv.ParseInt(tCreated, 10, 64)
			if err != nil {
				continue
			}
			timeTCreated := time.Unix(0, timestamp)
			if time.Since(timeTCreated) > expireAge {
				log.Printf("deleting topic %q", tID)
				if err := t.Delete(ctx); err != nil {
					return fmt.Errorf("Delete topic: %v: %v", t.String(), err)
				}
			}
		}
	}
	return nil
}

// cleanupSubscription deletes stale testing subscriptions.
func cleanupSubscription(ctx context.Context, client *pubsub.Client) error {
	if testing.Short() {
		return nil // Don't clean up in short mode.
	}
	// Delete subscriptions which were created a while ago.
	const expireAge = 24 * time.Hour

	it := client.Subscriptions(ctx)
	for {
		s, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return err
		}
		sID := s.ID()
		p := strings.Split(sID, "-")

		// Only delete resources created from the endtoend test.
		// Otherwise, this will affect other tests running midflight.
		if p[0] == resourcePrefix {
			sCreated := p[len(p)-2]
			timestamp, err := strconv.ParseInt(sCreated, 10, 64)
			if err != nil {
				continue
			}
			timeSCreated := time.Unix(0, timestamp)
			if time.Since(timeSCreated) > expireAge {
				log.Printf("deleting subscription %q", sID)
				if err := s.Delete(ctx); err != nil {
					return fmt.Errorf("Delete subscription: %v: %v", s.String(), err)
				}
			}
		}
	}
	return nil
}
