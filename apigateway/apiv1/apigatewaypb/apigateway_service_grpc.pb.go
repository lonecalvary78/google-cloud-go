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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.7
// source: google/cloud/apigateway/v1/apigateway_service.proto

package apigatewaypb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ApiGatewayService_ListGateways_FullMethodName    = "/google.cloud.apigateway.v1.ApiGatewayService/ListGateways"
	ApiGatewayService_GetGateway_FullMethodName      = "/google.cloud.apigateway.v1.ApiGatewayService/GetGateway"
	ApiGatewayService_CreateGateway_FullMethodName   = "/google.cloud.apigateway.v1.ApiGatewayService/CreateGateway"
	ApiGatewayService_UpdateGateway_FullMethodName   = "/google.cloud.apigateway.v1.ApiGatewayService/UpdateGateway"
	ApiGatewayService_DeleteGateway_FullMethodName   = "/google.cloud.apigateway.v1.ApiGatewayService/DeleteGateway"
	ApiGatewayService_ListApis_FullMethodName        = "/google.cloud.apigateway.v1.ApiGatewayService/ListApis"
	ApiGatewayService_GetApi_FullMethodName          = "/google.cloud.apigateway.v1.ApiGatewayService/GetApi"
	ApiGatewayService_CreateApi_FullMethodName       = "/google.cloud.apigateway.v1.ApiGatewayService/CreateApi"
	ApiGatewayService_UpdateApi_FullMethodName       = "/google.cloud.apigateway.v1.ApiGatewayService/UpdateApi"
	ApiGatewayService_DeleteApi_FullMethodName       = "/google.cloud.apigateway.v1.ApiGatewayService/DeleteApi"
	ApiGatewayService_ListApiConfigs_FullMethodName  = "/google.cloud.apigateway.v1.ApiGatewayService/ListApiConfigs"
	ApiGatewayService_GetApiConfig_FullMethodName    = "/google.cloud.apigateway.v1.ApiGatewayService/GetApiConfig"
	ApiGatewayService_CreateApiConfig_FullMethodName = "/google.cloud.apigateway.v1.ApiGatewayService/CreateApiConfig"
	ApiGatewayService_UpdateApiConfig_FullMethodName = "/google.cloud.apigateway.v1.ApiGatewayService/UpdateApiConfig"
	ApiGatewayService_DeleteApiConfig_FullMethodName = "/google.cloud.apigateway.v1.ApiGatewayService/DeleteApiConfig"
)

// ApiGatewayServiceClient is the client API for ApiGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiGatewayServiceClient interface {
	// Lists Gateways in a given project and location.
	ListGateways(ctx context.Context, in *ListGatewaysRequest, opts ...grpc.CallOption) (*ListGatewaysResponse, error)
	// Gets details of a single Gateway.
	GetGateway(ctx context.Context, in *GetGatewayRequest, opts ...grpc.CallOption) (*Gateway, error)
	// Creates a new Gateway in a given project and location.
	CreateGateway(ctx context.Context, in *CreateGatewayRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates the parameters of a single Gateway.
	UpdateGateway(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a single Gateway.
	DeleteGateway(ctx context.Context, in *DeleteGatewayRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Lists Apis in a given project and location.
	ListApis(ctx context.Context, in *ListApisRequest, opts ...grpc.CallOption) (*ListApisResponse, error)
	// Gets details of a single Api.
	GetApi(ctx context.Context, in *GetApiRequest, opts ...grpc.CallOption) (*Api, error)
	// Creates a new Api in a given project and location.
	CreateApi(ctx context.Context, in *CreateApiRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates the parameters of a single Api.
	UpdateApi(ctx context.Context, in *UpdateApiRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a single Api.
	DeleteApi(ctx context.Context, in *DeleteApiRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Lists ApiConfigs in a given project and location.
	ListApiConfigs(ctx context.Context, in *ListApiConfigsRequest, opts ...grpc.CallOption) (*ListApiConfigsResponse, error)
	// Gets details of a single ApiConfig.
	GetApiConfig(ctx context.Context, in *GetApiConfigRequest, opts ...grpc.CallOption) (*ApiConfig, error)
	// Creates a new ApiConfig in a given project and location.
	CreateApiConfig(ctx context.Context, in *CreateApiConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates the parameters of a single ApiConfig.
	UpdateApiConfig(ctx context.Context, in *UpdateApiConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a single ApiConfig.
	DeleteApiConfig(ctx context.Context, in *DeleteApiConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type apiGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiGatewayServiceClient(cc grpc.ClientConnInterface) ApiGatewayServiceClient {
	return &apiGatewayServiceClient{cc}
}

func (c *apiGatewayServiceClient) ListGateways(ctx context.Context, in *ListGatewaysRequest, opts ...grpc.CallOption) (*ListGatewaysResponse, error) {
	out := new(ListGatewaysResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_ListGateways_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) GetGateway(ctx context.Context, in *GetGatewayRequest, opts ...grpc.CallOption) (*Gateway, error) {
	out := new(Gateway)
	err := c.cc.Invoke(ctx, ApiGatewayService_GetGateway_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) CreateGateway(ctx context.Context, in *CreateGatewayRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ApiGatewayService_CreateGateway_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) UpdateGateway(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ApiGatewayService_UpdateGateway_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) DeleteGateway(ctx context.Context, in *DeleteGatewayRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ApiGatewayService_DeleteGateway_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) ListApis(ctx context.Context, in *ListApisRequest, opts ...grpc.CallOption) (*ListApisResponse, error) {
	out := new(ListApisResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_ListApis_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) GetApi(ctx context.Context, in *GetApiRequest, opts ...grpc.CallOption) (*Api, error) {
	out := new(Api)
	err := c.cc.Invoke(ctx, ApiGatewayService_GetApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) CreateApi(ctx context.Context, in *CreateApiRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ApiGatewayService_CreateApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) UpdateApi(ctx context.Context, in *UpdateApiRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ApiGatewayService_UpdateApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) DeleteApi(ctx context.Context, in *DeleteApiRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ApiGatewayService_DeleteApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) ListApiConfigs(ctx context.Context, in *ListApiConfigsRequest, opts ...grpc.CallOption) (*ListApiConfigsResponse, error) {
	out := new(ListApiConfigsResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_ListApiConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) GetApiConfig(ctx context.Context, in *GetApiConfigRequest, opts ...grpc.CallOption) (*ApiConfig, error) {
	out := new(ApiConfig)
	err := c.cc.Invoke(ctx, ApiGatewayService_GetApiConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) CreateApiConfig(ctx context.Context, in *CreateApiConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ApiGatewayService_CreateApiConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) UpdateApiConfig(ctx context.Context, in *UpdateApiConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ApiGatewayService_UpdateApiConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) DeleteApiConfig(ctx context.Context, in *DeleteApiConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ApiGatewayService_DeleteApiConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiGatewayServiceServer is the server API for ApiGatewayService service.
// All implementations should embed UnimplementedApiGatewayServiceServer
// for forward compatibility
type ApiGatewayServiceServer interface {
	// Lists Gateways in a given project and location.
	ListGateways(context.Context, *ListGatewaysRequest) (*ListGatewaysResponse, error)
	// Gets details of a single Gateway.
	GetGateway(context.Context, *GetGatewayRequest) (*Gateway, error)
	// Creates a new Gateway in a given project and location.
	CreateGateway(context.Context, *CreateGatewayRequest) (*longrunningpb.Operation, error)
	// Updates the parameters of a single Gateway.
	UpdateGateway(context.Context, *UpdateGatewayRequest) (*longrunningpb.Operation, error)
	// Deletes a single Gateway.
	DeleteGateway(context.Context, *DeleteGatewayRequest) (*longrunningpb.Operation, error)
	// Lists Apis in a given project and location.
	ListApis(context.Context, *ListApisRequest) (*ListApisResponse, error)
	// Gets details of a single Api.
	GetApi(context.Context, *GetApiRequest) (*Api, error)
	// Creates a new Api in a given project and location.
	CreateApi(context.Context, *CreateApiRequest) (*longrunningpb.Operation, error)
	// Updates the parameters of a single Api.
	UpdateApi(context.Context, *UpdateApiRequest) (*longrunningpb.Operation, error)
	// Deletes a single Api.
	DeleteApi(context.Context, *DeleteApiRequest) (*longrunningpb.Operation, error)
	// Lists ApiConfigs in a given project and location.
	ListApiConfigs(context.Context, *ListApiConfigsRequest) (*ListApiConfigsResponse, error)
	// Gets details of a single ApiConfig.
	GetApiConfig(context.Context, *GetApiConfigRequest) (*ApiConfig, error)
	// Creates a new ApiConfig in a given project and location.
	CreateApiConfig(context.Context, *CreateApiConfigRequest) (*longrunningpb.Operation, error)
	// Updates the parameters of a single ApiConfig.
	UpdateApiConfig(context.Context, *UpdateApiConfigRequest) (*longrunningpb.Operation, error)
	// Deletes a single ApiConfig.
	DeleteApiConfig(context.Context, *DeleteApiConfigRequest) (*longrunningpb.Operation, error)
}

// UnimplementedApiGatewayServiceServer should be embedded to have forward compatible implementations.
type UnimplementedApiGatewayServiceServer struct {
}

func (UnimplementedApiGatewayServiceServer) ListGateways(context.Context, *ListGatewaysRequest) (*ListGatewaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGateways not implemented")
}
func (UnimplementedApiGatewayServiceServer) GetGateway(context.Context, *GetGatewayRequest) (*Gateway, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGateway not implemented")
}
func (UnimplementedApiGatewayServiceServer) CreateGateway(context.Context, *CreateGatewayRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGateway not implemented")
}
func (UnimplementedApiGatewayServiceServer) UpdateGateway(context.Context, *UpdateGatewayRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGateway not implemented")
}
func (UnimplementedApiGatewayServiceServer) DeleteGateway(context.Context, *DeleteGatewayRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGateway not implemented")
}
func (UnimplementedApiGatewayServiceServer) ListApis(context.Context, *ListApisRequest) (*ListApisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApis not implemented")
}
func (UnimplementedApiGatewayServiceServer) GetApi(context.Context, *GetApiRequest) (*Api, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApi not implemented")
}
func (UnimplementedApiGatewayServiceServer) CreateApi(context.Context, *CreateApiRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApi not implemented")
}
func (UnimplementedApiGatewayServiceServer) UpdateApi(context.Context, *UpdateApiRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApi not implemented")
}
func (UnimplementedApiGatewayServiceServer) DeleteApi(context.Context, *DeleteApiRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApi not implemented")
}
func (UnimplementedApiGatewayServiceServer) ListApiConfigs(context.Context, *ListApiConfigsRequest) (*ListApiConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApiConfigs not implemented")
}
func (UnimplementedApiGatewayServiceServer) GetApiConfig(context.Context, *GetApiConfigRequest) (*ApiConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiConfig not implemented")
}
func (UnimplementedApiGatewayServiceServer) CreateApiConfig(context.Context, *CreateApiConfigRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApiConfig not implemented")
}
func (UnimplementedApiGatewayServiceServer) UpdateApiConfig(context.Context, *UpdateApiConfigRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApiConfig not implemented")
}
func (UnimplementedApiGatewayServiceServer) DeleteApiConfig(context.Context, *DeleteApiConfigRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApiConfig not implemented")
}

// UnsafeApiGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiGatewayServiceServer will
// result in compilation errors.
type UnsafeApiGatewayServiceServer interface {
	mustEmbedUnimplementedApiGatewayServiceServer()
}

func RegisterApiGatewayServiceServer(s grpc.ServiceRegistrar, srv ApiGatewayServiceServer) {
	s.RegisterService(&ApiGatewayService_ServiceDesc, srv)
}

func _ApiGatewayService_ListGateways_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).ListGateways(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_ListGateways_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).ListGateways(ctx, req.(*ListGatewaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_GetGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).GetGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_GetGateway_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).GetGateway(ctx, req.(*GetGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_CreateGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).CreateGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_CreateGateway_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).CreateGateway(ctx, req.(*CreateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_UpdateGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).UpdateGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_UpdateGateway_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).UpdateGateway(ctx, req.(*UpdateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_DeleteGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).DeleteGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_DeleteGateway_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).DeleteGateway(ctx, req.(*DeleteGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_ListApis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).ListApis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_ListApis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).ListApis(ctx, req.(*ListApisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_GetApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).GetApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_GetApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).GetApi(ctx, req.(*GetApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_CreateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).CreateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_CreateApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).CreateApi(ctx, req.(*CreateApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_UpdateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).UpdateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_UpdateApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).UpdateApi(ctx, req.(*UpdateApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_DeleteApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).DeleteApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_DeleteApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).DeleteApi(ctx, req.(*DeleteApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_ListApiConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApiConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).ListApiConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_ListApiConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).ListApiConfigs(ctx, req.(*ListApiConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_GetApiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).GetApiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_GetApiConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).GetApiConfig(ctx, req.(*GetApiConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_CreateApiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).CreateApiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_CreateApiConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).CreateApiConfig(ctx, req.(*CreateApiConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_UpdateApiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApiConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).UpdateApiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_UpdateApiConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).UpdateApiConfig(ctx, req.(*UpdateApiConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_DeleteApiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApiConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).DeleteApiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_DeleteApiConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).DeleteApiConfig(ctx, req.(*DeleteApiConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiGatewayService_ServiceDesc is the grpc.ServiceDesc for ApiGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.apigateway.v1.ApiGatewayService",
	HandlerType: (*ApiGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGateways",
			Handler:    _ApiGatewayService_ListGateways_Handler,
		},
		{
			MethodName: "GetGateway",
			Handler:    _ApiGatewayService_GetGateway_Handler,
		},
		{
			MethodName: "CreateGateway",
			Handler:    _ApiGatewayService_CreateGateway_Handler,
		},
		{
			MethodName: "UpdateGateway",
			Handler:    _ApiGatewayService_UpdateGateway_Handler,
		},
		{
			MethodName: "DeleteGateway",
			Handler:    _ApiGatewayService_DeleteGateway_Handler,
		},
		{
			MethodName: "ListApis",
			Handler:    _ApiGatewayService_ListApis_Handler,
		},
		{
			MethodName: "GetApi",
			Handler:    _ApiGatewayService_GetApi_Handler,
		},
		{
			MethodName: "CreateApi",
			Handler:    _ApiGatewayService_CreateApi_Handler,
		},
		{
			MethodName: "UpdateApi",
			Handler:    _ApiGatewayService_UpdateApi_Handler,
		},
		{
			MethodName: "DeleteApi",
			Handler:    _ApiGatewayService_DeleteApi_Handler,
		},
		{
			MethodName: "ListApiConfigs",
			Handler:    _ApiGatewayService_ListApiConfigs_Handler,
		},
		{
			MethodName: "GetApiConfig",
			Handler:    _ApiGatewayService_GetApiConfig_Handler,
		},
		{
			MethodName: "CreateApiConfig",
			Handler:    _ApiGatewayService_CreateApiConfig_Handler,
		},
		{
			MethodName: "UpdateApiConfig",
			Handler:    _ApiGatewayService_UpdateApiConfig_Handler,
		},
		{
			MethodName: "DeleteApiConfig",
			Handler:    _ApiGatewayService_DeleteApiConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/apigateway/v1/apigateway_service.proto",
}
