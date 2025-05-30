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

package css

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	csspb "cloud.google.com/go/shopping/css/apiv1/csspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newCssProductInputsClientHook clientHook

// CssProductInputsCallOptions contains the retry settings for each method of CssProductInputsClient.
type CssProductInputsCallOptions struct {
	InsertCssProductInput []gax.CallOption
	UpdateCssProductInput []gax.CallOption
	DeleteCssProductInput []gax.CallOption
}

func defaultCssProductInputsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("css.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("css.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("css.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://css.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCssProductInputsCallOptions() *CssProductInputsCallOptions {
	return &CssProductInputsCallOptions{
		InsertCssProductInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateCssProductInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteCssProductInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultCssProductInputsRESTCallOptions() *CssProductInputsCallOptions {
	return &CssProductInputsCallOptions{
		InsertCssProductInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateCssProductInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteCssProductInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalCssProductInputsClient is an interface that defines the methods available from CSS API.
type internalCssProductInputsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	InsertCssProductInput(context.Context, *csspb.InsertCssProductInputRequest, ...gax.CallOption) (*csspb.CssProductInput, error)
	UpdateCssProductInput(context.Context, *csspb.UpdateCssProductInputRequest, ...gax.CallOption) (*csspb.CssProductInput, error)
	DeleteCssProductInput(context.Context, *csspb.DeleteCssProductInputRequest, ...gax.CallOption) error
}

// CssProductInputsClient is a client for interacting with CSS API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to use CssProductInput resource.
// This service helps to insert/update/delete CSS Products.
type CssProductInputsClient struct {
	// The internal transport-dependent client.
	internalClient internalCssProductInputsClient

	// The call options for this service.
	CallOptions *CssProductInputsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CssProductInputsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CssProductInputsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CssProductInputsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// InsertCssProductInput uploads a CssProductInput to your CSS Center account. If an
// input with the same contentLanguage, identity, feedLabel and feedId already
// exists, this method replaces that entry.
//
// After inserting, updating, or deleting a CSS Product input, it may
// take several minutes before the processed CSS Product can be retrieved.
func (c *CssProductInputsClient) InsertCssProductInput(ctx context.Context, req *csspb.InsertCssProductInputRequest, opts ...gax.CallOption) (*csspb.CssProductInput, error) {
	return c.internalClient.InsertCssProductInput(ctx, req, opts...)
}

// UpdateCssProductInput updates the existing Css Product input in your CSS Center account.
//
// After inserting, updating, or deleting a CSS Product input, it may take
// several minutes before the processed Css Product can be retrieved.
func (c *CssProductInputsClient) UpdateCssProductInput(ctx context.Context, req *csspb.UpdateCssProductInputRequest, opts ...gax.CallOption) (*csspb.CssProductInput, error) {
	return c.internalClient.UpdateCssProductInput(ctx, req, opts...)
}

// DeleteCssProductInput deletes a CSS Product input from your CSS Center account.
//
// After a delete it may take several minutes until the input is no longer
// available.
func (c *CssProductInputsClient) DeleteCssProductInput(ctx context.Context, req *csspb.DeleteCssProductInputRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteCssProductInput(ctx, req, opts...)
}

// cssProductInputsGRPCClient is a client for interacting with CSS API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type cssProductInputsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CssProductInputsClient
	CallOptions **CssProductInputsCallOptions

	// The gRPC API client.
	cssProductInputsClient csspb.CssProductInputsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewCssProductInputsClient creates a new css product inputs service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to use CssProductInput resource.
// This service helps to insert/update/delete CSS Products.
func NewCssProductInputsClient(ctx context.Context, opts ...option.ClientOption) (*CssProductInputsClient, error) {
	clientOpts := defaultCssProductInputsGRPCClientOptions()
	if newCssProductInputsClientHook != nil {
		hookOpts, err := newCssProductInputsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CssProductInputsClient{CallOptions: defaultCssProductInputsCallOptions()}

	c := &cssProductInputsGRPCClient{
		connPool:               connPool,
		cssProductInputsClient: csspb.NewCssProductInputsServiceClient(connPool),
		CallOptions:            &client.CallOptions,
		logger:                 internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *cssProductInputsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *cssProductInputsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *cssProductInputsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type cssProductInputsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing CssProductInputsClient
	CallOptions **CssProductInputsCallOptions

	logger *slog.Logger
}

// NewCssProductInputsRESTClient creates a new css product inputs service rest client.
//
// Service to use CssProductInput resource.
// This service helps to insert/update/delete CSS Products.
func NewCssProductInputsRESTClient(ctx context.Context, opts ...option.ClientOption) (*CssProductInputsClient, error) {
	clientOpts := append(defaultCssProductInputsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultCssProductInputsRESTCallOptions()
	c := &cssProductInputsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &CssProductInputsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultCssProductInputsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://css.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://css.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://css.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://css.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *cssProductInputsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *cssProductInputsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *cssProductInputsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *cssProductInputsGRPCClient) InsertCssProductInput(ctx context.Context, req *csspb.InsertCssProductInputRequest, opts ...gax.CallOption) (*csspb.CssProductInput, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).InsertCssProductInput[0:len((*c.CallOptions).InsertCssProductInput):len((*c.CallOptions).InsertCssProductInput)], opts...)
	var resp *csspb.CssProductInput
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.cssProductInputsClient.InsertCssProductInput, req, settings.GRPC, c.logger, "InsertCssProductInput")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *cssProductInputsGRPCClient) UpdateCssProductInput(ctx context.Context, req *csspb.UpdateCssProductInputRequest, opts ...gax.CallOption) (*csspb.CssProductInput, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "css_product_input.name", url.QueryEscape(req.GetCssProductInput().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateCssProductInput[0:len((*c.CallOptions).UpdateCssProductInput):len((*c.CallOptions).UpdateCssProductInput)], opts...)
	var resp *csspb.CssProductInput
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.cssProductInputsClient.UpdateCssProductInput, req, settings.GRPC, c.logger, "UpdateCssProductInput")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *cssProductInputsGRPCClient) DeleteCssProductInput(ctx context.Context, req *csspb.DeleteCssProductInputRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteCssProductInput[0:len((*c.CallOptions).DeleteCssProductInput):len((*c.CallOptions).DeleteCssProductInput)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.cssProductInputsClient.DeleteCssProductInput, req, settings.GRPC, c.logger, "DeleteCssProductInput")
		return err
	}, opts...)
	return err
}

// InsertCssProductInput uploads a CssProductInput to your CSS Center account. If an
// input with the same contentLanguage, identity, feedLabel and feedId already
// exists, this method replaces that entry.
//
// After inserting, updating, or deleting a CSS Product input, it may
// take several minutes before the processed CSS Product can be retrieved.
func (c *cssProductInputsRESTClient) InsertCssProductInput(ctx context.Context, req *csspb.InsertCssProductInputRequest, opts ...gax.CallOption) (*csspb.CssProductInput, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetCssProductInput()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/cssProductInputs:insert", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetFeedId() != 0 {
		params.Add("feedId", fmt.Sprintf("%v", req.GetFeedId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).InsertCssProductInput[0:len((*c.CallOptions).InsertCssProductInput):len((*c.CallOptions).InsertCssProductInput)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &csspb.CssProductInput{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "InsertCssProductInput")
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

// UpdateCssProductInput updates the existing Css Product input in your CSS Center account.
//
// After inserting, updating, or deleting a CSS Product input, it may take
// several minutes before the processed Css Product can be retrieved.
func (c *cssProductInputsRESTClient) UpdateCssProductInput(ctx context.Context, req *csspb.UpdateCssProductInputRequest, opts ...gax.CallOption) (*csspb.CssProductInput, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetCssProductInput()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetCssProductInput().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		field, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(field[1:len(field)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "css_product_input.name", url.QueryEscape(req.GetCssProductInput().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateCssProductInput[0:len((*c.CallOptions).UpdateCssProductInput):len((*c.CallOptions).UpdateCssProductInput)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &csspb.CssProductInput{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateCssProductInput")
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

// DeleteCssProductInput deletes a CSS Product input from your CSS Center account.
//
// After a delete it may take several minutes until the input is no longer
// available.
func (c *cssProductInputsRESTClient) DeleteCssProductInput(ctx context.Context, req *csspb.DeleteCssProductInputRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req != nil && req.SupplementalFeedId != nil {
		params.Add("supplementalFeedId", fmt.Sprintf("%v", req.GetSupplementalFeedId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		_, err = executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteCssProductInput")
		return err
	}, opts...)
}
