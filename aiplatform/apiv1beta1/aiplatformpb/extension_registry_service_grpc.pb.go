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
// source: google/cloud/aiplatform/v1beta1/extension_registry_service.proto

package aiplatformpb

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
	ExtensionRegistryService_ImportExtension_FullMethodName = "/google.cloud.aiplatform.v1beta1.ExtensionRegistryService/ImportExtension"
	ExtensionRegistryService_GetExtension_FullMethodName    = "/google.cloud.aiplatform.v1beta1.ExtensionRegistryService/GetExtension"
	ExtensionRegistryService_ListExtensions_FullMethodName  = "/google.cloud.aiplatform.v1beta1.ExtensionRegistryService/ListExtensions"
	ExtensionRegistryService_UpdateExtension_FullMethodName = "/google.cloud.aiplatform.v1beta1.ExtensionRegistryService/UpdateExtension"
	ExtensionRegistryService_DeleteExtension_FullMethodName = "/google.cloud.aiplatform.v1beta1.ExtensionRegistryService/DeleteExtension"
)

// ExtensionRegistryServiceClient is the client API for ExtensionRegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtensionRegistryServiceClient interface {
	// Imports an Extension.
	ImportExtension(ctx context.Context, in *ImportExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Gets an Extension.
	GetExtension(ctx context.Context, in *GetExtensionRequest, opts ...grpc.CallOption) (*Extension, error)
	// Lists Extensions in a location.
	ListExtensions(ctx context.Context, in *ListExtensionsRequest, opts ...grpc.CallOption) (*ListExtensionsResponse, error)
	// Updates an Extension.
	UpdateExtension(ctx context.Context, in *UpdateExtensionRequest, opts ...grpc.CallOption) (*Extension, error)
	// Deletes an Extension.
	DeleteExtension(ctx context.Context, in *DeleteExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type extensionRegistryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExtensionRegistryServiceClient(cc grpc.ClientConnInterface) ExtensionRegistryServiceClient {
	return &extensionRegistryServiceClient{cc}
}

func (c *extensionRegistryServiceClient) ImportExtension(ctx context.Context, in *ImportExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ExtensionRegistryService_ImportExtension_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionRegistryServiceClient) GetExtension(ctx context.Context, in *GetExtensionRequest, opts ...grpc.CallOption) (*Extension, error) {
	out := new(Extension)
	err := c.cc.Invoke(ctx, ExtensionRegistryService_GetExtension_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionRegistryServiceClient) ListExtensions(ctx context.Context, in *ListExtensionsRequest, opts ...grpc.CallOption) (*ListExtensionsResponse, error) {
	out := new(ListExtensionsResponse)
	err := c.cc.Invoke(ctx, ExtensionRegistryService_ListExtensions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionRegistryServiceClient) UpdateExtension(ctx context.Context, in *UpdateExtensionRequest, opts ...grpc.CallOption) (*Extension, error) {
	out := new(Extension)
	err := c.cc.Invoke(ctx, ExtensionRegistryService_UpdateExtension_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionRegistryServiceClient) DeleteExtension(ctx context.Context, in *DeleteExtensionRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ExtensionRegistryService_DeleteExtension_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionRegistryServiceServer is the server API for ExtensionRegistryService service.
// All implementations should embed UnimplementedExtensionRegistryServiceServer
// for forward compatibility
type ExtensionRegistryServiceServer interface {
	// Imports an Extension.
	ImportExtension(context.Context, *ImportExtensionRequest) (*longrunningpb.Operation, error)
	// Gets an Extension.
	GetExtension(context.Context, *GetExtensionRequest) (*Extension, error)
	// Lists Extensions in a location.
	ListExtensions(context.Context, *ListExtensionsRequest) (*ListExtensionsResponse, error)
	// Updates an Extension.
	UpdateExtension(context.Context, *UpdateExtensionRequest) (*Extension, error)
	// Deletes an Extension.
	DeleteExtension(context.Context, *DeleteExtensionRequest) (*longrunningpb.Operation, error)
}

// UnimplementedExtensionRegistryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedExtensionRegistryServiceServer struct {
}

func (UnimplementedExtensionRegistryServiceServer) ImportExtension(context.Context, *ImportExtensionRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportExtension not implemented")
}
func (UnimplementedExtensionRegistryServiceServer) GetExtension(context.Context, *GetExtensionRequest) (*Extension, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExtension not implemented")
}
func (UnimplementedExtensionRegistryServiceServer) ListExtensions(context.Context, *ListExtensionsRequest) (*ListExtensionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExtensions not implemented")
}
func (UnimplementedExtensionRegistryServiceServer) UpdateExtension(context.Context, *UpdateExtensionRequest) (*Extension, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExtension not implemented")
}
func (UnimplementedExtensionRegistryServiceServer) DeleteExtension(context.Context, *DeleteExtensionRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExtension not implemented")
}

// UnsafeExtensionRegistryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtensionRegistryServiceServer will
// result in compilation errors.
type UnsafeExtensionRegistryServiceServer interface {
	mustEmbedUnimplementedExtensionRegistryServiceServer()
}

func RegisterExtensionRegistryServiceServer(s grpc.ServiceRegistrar, srv ExtensionRegistryServiceServer) {
	s.RegisterService(&ExtensionRegistryService_ServiceDesc, srv)
}

func _ExtensionRegistryService_ImportExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionRegistryServiceServer).ImportExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionRegistryService_ImportExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionRegistryServiceServer).ImportExtension(ctx, req.(*ImportExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionRegistryService_GetExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionRegistryServiceServer).GetExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionRegistryService_GetExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionRegistryServiceServer).GetExtension(ctx, req.(*GetExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionRegistryService_ListExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExtensionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionRegistryServiceServer).ListExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionRegistryService_ListExtensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionRegistryServiceServer).ListExtensions(ctx, req.(*ListExtensionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionRegistryService_UpdateExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionRegistryServiceServer).UpdateExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionRegistryService_UpdateExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionRegistryServiceServer).UpdateExtension(ctx, req.(*UpdateExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionRegistryService_DeleteExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionRegistryServiceServer).DeleteExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionRegistryService_DeleteExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionRegistryServiceServer).DeleteExtension(ctx, req.(*DeleteExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExtensionRegistryService_ServiceDesc is the grpc.ServiceDesc for ExtensionRegistryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExtensionRegistryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1beta1.ExtensionRegistryService",
	HandlerType: (*ExtensionRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportExtension",
			Handler:    _ExtensionRegistryService_ImportExtension_Handler,
		},
		{
			MethodName: "GetExtension",
			Handler:    _ExtensionRegistryService_GetExtension_Handler,
		},
		{
			MethodName: "ListExtensions",
			Handler:    _ExtensionRegistryService_ListExtensions_Handler,
		},
		{
			MethodName: "UpdateExtension",
			Handler:    _ExtensionRegistryService_UpdateExtension_Handler,
		},
		{
			MethodName: "DeleteExtension",
			Handler:    _ExtensionRegistryService_DeleteExtension_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1beta1/extension_registry_service.proto",
}
