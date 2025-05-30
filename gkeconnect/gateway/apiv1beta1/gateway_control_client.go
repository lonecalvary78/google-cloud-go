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

package gateway

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	gatewaypb "cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newGatewayControlClientHook clientHook

// GatewayControlCallOptions contains the retry settings for each method of GatewayControlClient.
type GatewayControlCallOptions struct {
	GenerateCredentials []gax.CallOption
}

func defaultGatewayControlRESTCallOptions() *GatewayControlCallOptions {
	return &GatewayControlCallOptions{
		GenerateCredentials: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalGatewayControlClient is an interface that defines the methods available from Connect Gateway API.
type internalGatewayControlClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GenerateCredentials(context.Context, *gatewaypb.GenerateCredentialsRequest, ...gax.CallOption) (*gatewaypb.GenerateCredentialsResponse, error)
}

// GatewayControlClient is a client for interacting with Connect Gateway API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// GatewayControl is the control plane API for Connect Gateway.
type GatewayControlClient struct {
	// The internal transport-dependent client.
	internalClient internalGatewayControlClient

	// The call options for this service.
	CallOptions *GatewayControlCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GatewayControlClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GatewayControlClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *GatewayControlClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GenerateCredentials generateCredentials provides connection information that allows a user to
// access the specified membership using Connect Gateway.
func (c *GatewayControlClient) GenerateCredentials(ctx context.Context, req *gatewaypb.GenerateCredentialsRequest, opts ...gax.CallOption) (*gatewaypb.GenerateCredentialsResponse, error) {
	return c.internalClient.GenerateCredentials(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gatewayControlRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing GatewayControlClient
	CallOptions **GatewayControlCallOptions

	logger *slog.Logger
}

// NewGatewayControlRESTClient creates a new gateway control rest client.
//
// GatewayControl is the control plane API for Connect Gateway.
func NewGatewayControlRESTClient(ctx context.Context, opts ...option.ClientOption) (*GatewayControlClient, error) {
	clientOpts := append(defaultGatewayControlRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultGatewayControlRESTCallOptions()
	c := &gatewayControlRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &GatewayControlClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultGatewayControlRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://connectgateway.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://connectgateway.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://connectgateway.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://connectgateway.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gatewayControlRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gatewayControlRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *gatewayControlRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// GenerateCredentials generateCredentials provides connection information that allows a user to
// access the specified membership using Connect Gateway.
func (c *gatewayControlRESTClient) GenerateCredentials(ctx context.Context, req *gatewaypb.GenerateCredentialsRequest, opts ...gax.CallOption) (*gatewaypb.GenerateCredentialsResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v:generateCredentials", req.GetName())

	params := url.Values{}
	if req.GetForceUseAgent() {
		params.Add("forceUseAgent", fmt.Sprintf("%v", req.GetForceUseAgent()))
	}
	if req.GetKubernetesNamespace() != "" {
		params.Add("kubernetesNamespace", fmt.Sprintf("%v", req.GetKubernetesNamespace()))
	}
	if req.GetOperatingSystem() != 0 {
		params.Add("operatingSystem", fmt.Sprintf("%v", req.GetOperatingSystem()))
	}
	if req.GetVersion() != "" {
		params.Add("version", fmt.Sprintf("%v", req.GetVersion()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GenerateCredentials[0:len((*c.CallOptions).GenerateCredentials):len((*c.CallOptions).GenerateCredentials)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &gatewaypb.GenerateCredentialsResponse{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GenerateCredentials")
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
