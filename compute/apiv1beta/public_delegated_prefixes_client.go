// Copyright 2025 Google LLC
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

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"sort"
	"time"

	computepb "cloud.google.com/go/compute/apiv1beta/computepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newPublicDelegatedPrefixesClientHook clientHook

// PublicDelegatedPrefixesCallOptions contains the retry settings for each method of PublicDelegatedPrefixesClient.
type PublicDelegatedPrefixesCallOptions struct {
	AggregatedList []gax.CallOption
	Announce       []gax.CallOption
	Delete         []gax.CallOption
	Get            []gax.CallOption
	Insert         []gax.CallOption
	List           []gax.CallOption
	Patch          []gax.CallOption
	Withdraw       []gax.CallOption
}

func defaultPublicDelegatedPrefixesRESTCallOptions() *PublicDelegatedPrefixesCallOptions {
	return &PublicDelegatedPrefixesCallOptions{
		AggregatedList: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		Announce: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		Delete: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		Get: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		Insert: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		List: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		Patch: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		Withdraw: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
	}
}

// internalPublicDelegatedPrefixesClient is an interface that defines the methods available from Google Compute Engine API.
type internalPublicDelegatedPrefixesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListPublicDelegatedPrefixesRequest, ...gax.CallOption) *PublicDelegatedPrefixesScopedListPairIterator
	Announce(context.Context, *computepb.AnnouncePublicDelegatedPrefixeRequest, ...gax.CallOption) (*Operation, error)
	Delete(context.Context, *computepb.DeletePublicDelegatedPrefixeRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetPublicDelegatedPrefixeRequest, ...gax.CallOption) (*computepb.PublicDelegatedPrefix, error)
	Insert(context.Context, *computepb.InsertPublicDelegatedPrefixeRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListPublicDelegatedPrefixesRequest, ...gax.CallOption) *PublicDelegatedPrefixIterator
	Patch(context.Context, *computepb.PatchPublicDelegatedPrefixeRequest, ...gax.CallOption) (*Operation, error)
	Withdraw(context.Context, *computepb.WithdrawPublicDelegatedPrefixeRequest, ...gax.CallOption) (*Operation, error)
}

// PublicDelegatedPrefixesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The PublicDelegatedPrefixes API.
type PublicDelegatedPrefixesClient struct {
	// The internal transport-dependent client.
	internalClient internalPublicDelegatedPrefixesClient

	// The call options for this service.
	CallOptions *PublicDelegatedPrefixesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PublicDelegatedPrefixesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PublicDelegatedPrefixesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PublicDelegatedPrefixesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList lists all PublicDelegatedPrefix resources owned by the specific project across all scopes. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *PublicDelegatedPrefixesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListPublicDelegatedPrefixesRequest, opts ...gax.CallOption) *PublicDelegatedPrefixesScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Announce announces the specified PublicDelegatedPrefix in the given region.
func (c *PublicDelegatedPrefixesClient) Announce(ctx context.Context, req *computepb.AnnouncePublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Announce(ctx, req, opts...)
}

// Delete deletes the specified PublicDelegatedPrefix in the given region.
func (c *PublicDelegatedPrefixesClient) Delete(ctx context.Context, req *computepb.DeletePublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified PublicDelegatedPrefix resource in the given region.
func (c *PublicDelegatedPrefixesClient) Get(ctx context.Context, req *computepb.GetPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.PublicDelegatedPrefix, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a PublicDelegatedPrefix in the specified project in the given region using the parameters that are included in the request.
func (c *PublicDelegatedPrefixesClient) Insert(ctx context.Context, req *computepb.InsertPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List lists the PublicDelegatedPrefixes for a project in the given region.
func (c *PublicDelegatedPrefixesClient) List(ctx context.Context, req *computepb.ListPublicDelegatedPrefixesRequest, opts ...gax.CallOption) *PublicDelegatedPrefixIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified PublicDelegatedPrefix resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *PublicDelegatedPrefixesClient) Patch(ctx context.Context, req *computepb.PatchPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// Withdraw withdraws the specified PublicDelegatedPrefix in the given region.
func (c *PublicDelegatedPrefixesClient) Withdraw(ctx context.Context, req *computepb.WithdrawPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Withdraw(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type publicDelegatedPrefixesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// operationClient is used to call the operation-specific management service.
	operationClient *RegionOperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing PublicDelegatedPrefixesClient
	CallOptions **PublicDelegatedPrefixesCallOptions

	logger *slog.Logger
}

// NewPublicDelegatedPrefixesRESTClient creates a new public delegated prefixes rest client.
//
// The PublicDelegatedPrefixes API.
func NewPublicDelegatedPrefixesRESTClient(ctx context.Context, opts ...option.ClientOption) (*PublicDelegatedPrefixesClient, error) {
	clientOpts := append(defaultPublicDelegatedPrefixesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultPublicDelegatedPrefixesRESTCallOptions()
	c := &publicDelegatedPrefixesRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	o := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opC, err := NewRegionOperationsRESTClient(ctx, o...)
	if err != nil {
		return nil, err
	}
	c.operationClient = opC

	return &PublicDelegatedPrefixesClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultPublicDelegatedPrefixesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://compute.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *publicDelegatedPrefixesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *publicDelegatedPrefixesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	if err := c.operationClient.Close(); err != nil {
		return err
	}
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *publicDelegatedPrefixesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AggregatedList lists all PublicDelegatedPrefix resources owned by the specific project across all scopes. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *publicDelegatedPrefixesRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListPublicDelegatedPrefixesRequest, opts ...gax.CallOption) *PublicDelegatedPrefixesScopedListPairIterator {
	it := &PublicDelegatedPrefixesScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListPublicDelegatedPrefixesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]PublicDelegatedPrefixesScopedListPair, string, error) {
		resp := &computepb.PublicDelegatedPrefixAggregatedList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(uint32(math.MaxInt32))
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/compute/beta/projects/%v/aggregated/publicDelegatedPrefixes", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.IncludeAllScopes != nil {
			params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}
		if req != nil && req.ServiceProjectNumber != nil {
			params.Add("serviceProjectNumber", fmt.Sprintf("%v", req.GetServiceProjectNumber()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "AggregatedList")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp

		elems := make([]PublicDelegatedPrefixesScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, PublicDelegatedPrefixesScopedListPair{k, v})
		}
		sort.Slice(elems, func(i, j int) bool { return elems[i].Key < elems[j].Key })

		return elems, resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Announce announces the specified PublicDelegatedPrefix in the given region.
func (c *publicDelegatedPrefixesRESTClient) Announce(ctx context.Context, req *computepb.AnnouncePublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/beta/projects/%v/regions/%v/publicDelegatedPrefixes/%v/announce", req.GetProject(), req.GetRegion(), req.GetPublicDelegatedPrefix())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "public_delegated_prefix", url.QueryEscape(req.GetPublicDelegatedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Announce[0:len((*c.CallOptions).Announce):len((*c.CallOptions).Announce)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "Announce")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// Delete deletes the specified PublicDelegatedPrefix in the given region.
func (c *publicDelegatedPrefixesRESTClient) Delete(ctx context.Context, req *computepb.DeletePublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/beta/projects/%v/regions/%v/publicDelegatedPrefixes/%v", req.GetProject(), req.GetRegion(), req.GetPublicDelegatedPrefix())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "public_delegated_prefix", url.QueryEscape(req.GetPublicDelegatedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "Delete")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// Get returns the specified PublicDelegatedPrefix resource in the given region.
func (c *publicDelegatedPrefixesRESTClient) Get(ctx context.Context, req *computepb.GetPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.PublicDelegatedPrefix, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/beta/projects/%v/regions/%v/publicDelegatedPrefixes/%v", req.GetProject(), req.GetRegion(), req.GetPublicDelegatedPrefix())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "public_delegated_prefix", url.QueryEscape(req.GetPublicDelegatedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.PublicDelegatedPrefix{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "Get")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// Insert creates a PublicDelegatedPrefix in the specified project in the given region using the parameters that are included in the request.
func (c *publicDelegatedPrefixesRESTClient) Insert(ctx context.Context, req *computepb.InsertPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetPublicDelegatedPrefixResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/beta/projects/%v/regions/%v/publicDelegatedPrefixes", req.GetProject(), req.GetRegion())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "Insert")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// List lists the PublicDelegatedPrefixes for a project in the given region.
func (c *publicDelegatedPrefixesRESTClient) List(ctx context.Context, req *computepb.ListPublicDelegatedPrefixesRequest, opts ...gax.CallOption) *PublicDelegatedPrefixIterator {
	it := &PublicDelegatedPrefixIterator{}
	req = proto.Clone(req).(*computepb.ListPublicDelegatedPrefixesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.PublicDelegatedPrefix, string, error) {
		resp := &computepb.PublicDelegatedPrefixList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(uint32(math.MaxInt32))
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/compute/beta/projects/%v/regions/%v/publicDelegatedPrefixes", req.GetProject(), req.GetRegion())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "List")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Patch patches the specified PublicDelegatedPrefix resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *publicDelegatedPrefixesRESTClient) Patch(ctx context.Context, req *computepb.PatchPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetPublicDelegatedPrefixResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/beta/projects/%v/regions/%v/publicDelegatedPrefixes/%v", req.GetProject(), req.GetRegion(), req.GetPublicDelegatedPrefix())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "public_delegated_prefix", url.QueryEscape(req.GetPublicDelegatedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "Patch")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// Withdraw withdraws the specified PublicDelegatedPrefix in the given region.
func (c *publicDelegatedPrefixesRESTClient) Withdraw(ctx context.Context, req *computepb.WithdrawPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/beta/projects/%v/regions/%v/publicDelegatedPrefixes/%v/withdraw", req.GetProject(), req.GetRegion(), req.GetPublicDelegatedPrefix())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "public_delegated_prefix", url.QueryEscape(req.GetPublicDelegatedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Withdraw[0:len((*c.CallOptions).Withdraw):len((*c.CallOptions).Withdraw)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "Withdraw")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}
