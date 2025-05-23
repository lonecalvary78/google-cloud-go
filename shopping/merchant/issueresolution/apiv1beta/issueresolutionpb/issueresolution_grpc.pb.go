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
// source: google/shopping/merchant/issueresolution/v1beta/issueresolution.proto

package issueresolutionpb

import (
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
	IssueResolutionService_RenderAccountIssues_FullMethodName = "/google.shopping.merchant.issueresolution.v1beta.IssueResolutionService/RenderAccountIssues"
	IssueResolutionService_RenderProductIssues_FullMethodName = "/google.shopping.merchant.issueresolution.v1beta.IssueResolutionService/RenderProductIssues"
	IssueResolutionService_TriggerAction_FullMethodName       = "/google.shopping.merchant.issueresolution.v1beta.IssueResolutionService/TriggerAction"
)

// IssueResolutionServiceClient is the client API for IssueResolutionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IssueResolutionServiceClient interface {
	// Provide a list of business's account issues with an issue resolution
	// content and available actions. This content and actions are meant to be
	// rendered and shown in third-party applications.
	RenderAccountIssues(ctx context.Context, in *RenderAccountIssuesRequest, opts ...grpc.CallOption) (*RenderAccountIssuesResponse, error)
	// Provide a list of issues for business's product with an issue resolution
	// content and available actions. This content and actions are meant to be
	// rendered and shown in third-party applications.
	RenderProductIssues(ctx context.Context, in *RenderProductIssuesRequest, opts ...grpc.CallOption) (*RenderProductIssuesResponse, error)
	// Start an action. The action can be requested by a business in
	// third-party application. Before the business can request the action, the
	// third-party application needs to show them action specific content and
	// display a user input form.
	//
	// The action can be successfully started only once all `required` inputs are
	// provided. If any `required` input is missing, or invalid value was
	// provided, the service will return 400 error. Validation errors will contain
	// [Ids][google.shopping.merchant.issueresolution.v1beta.InputField.id] for
	// all problematic field together with translated, human readable error
	// messages that can be shown to the user.
	TriggerAction(ctx context.Context, in *TriggerActionRequest, opts ...grpc.CallOption) (*TriggerActionResponse, error)
}

type issueResolutionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIssueResolutionServiceClient(cc grpc.ClientConnInterface) IssueResolutionServiceClient {
	return &issueResolutionServiceClient{cc}
}

func (c *issueResolutionServiceClient) RenderAccountIssues(ctx context.Context, in *RenderAccountIssuesRequest, opts ...grpc.CallOption) (*RenderAccountIssuesResponse, error) {
	out := new(RenderAccountIssuesResponse)
	err := c.cc.Invoke(ctx, IssueResolutionService_RenderAccountIssues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueResolutionServiceClient) RenderProductIssues(ctx context.Context, in *RenderProductIssuesRequest, opts ...grpc.CallOption) (*RenderProductIssuesResponse, error) {
	out := new(RenderProductIssuesResponse)
	err := c.cc.Invoke(ctx, IssueResolutionService_RenderProductIssues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueResolutionServiceClient) TriggerAction(ctx context.Context, in *TriggerActionRequest, opts ...grpc.CallOption) (*TriggerActionResponse, error) {
	out := new(TriggerActionResponse)
	err := c.cc.Invoke(ctx, IssueResolutionService_TriggerAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssueResolutionServiceServer is the server API for IssueResolutionService service.
// All implementations should embed UnimplementedIssueResolutionServiceServer
// for forward compatibility
type IssueResolutionServiceServer interface {
	// Provide a list of business's account issues with an issue resolution
	// content and available actions. This content and actions are meant to be
	// rendered and shown in third-party applications.
	RenderAccountIssues(context.Context, *RenderAccountIssuesRequest) (*RenderAccountIssuesResponse, error)
	// Provide a list of issues for business's product with an issue resolution
	// content and available actions. This content and actions are meant to be
	// rendered and shown in third-party applications.
	RenderProductIssues(context.Context, *RenderProductIssuesRequest) (*RenderProductIssuesResponse, error)
	// Start an action. The action can be requested by a business in
	// third-party application. Before the business can request the action, the
	// third-party application needs to show them action specific content and
	// display a user input form.
	//
	// The action can be successfully started only once all `required` inputs are
	// provided. If any `required` input is missing, or invalid value was
	// provided, the service will return 400 error. Validation errors will contain
	// [Ids][google.shopping.merchant.issueresolution.v1beta.InputField.id] for
	// all problematic field together with translated, human readable error
	// messages that can be shown to the user.
	TriggerAction(context.Context, *TriggerActionRequest) (*TriggerActionResponse, error)
}

// UnimplementedIssueResolutionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedIssueResolutionServiceServer struct {
}

func (UnimplementedIssueResolutionServiceServer) RenderAccountIssues(context.Context, *RenderAccountIssuesRequest) (*RenderAccountIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenderAccountIssues not implemented")
}
func (UnimplementedIssueResolutionServiceServer) RenderProductIssues(context.Context, *RenderProductIssuesRequest) (*RenderProductIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenderProductIssues not implemented")
}
func (UnimplementedIssueResolutionServiceServer) TriggerAction(context.Context, *TriggerActionRequest) (*TriggerActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerAction not implemented")
}

// UnsafeIssueResolutionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IssueResolutionServiceServer will
// result in compilation errors.
type UnsafeIssueResolutionServiceServer interface {
	mustEmbedUnimplementedIssueResolutionServiceServer()
}

func RegisterIssueResolutionServiceServer(s grpc.ServiceRegistrar, srv IssueResolutionServiceServer) {
	s.RegisterService(&IssueResolutionService_ServiceDesc, srv)
}

func _IssueResolutionService_RenderAccountIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderAccountIssuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueResolutionServiceServer).RenderAccountIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IssueResolutionService_RenderAccountIssues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueResolutionServiceServer).RenderAccountIssues(ctx, req.(*RenderAccountIssuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueResolutionService_RenderProductIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderProductIssuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueResolutionServiceServer).RenderProductIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IssueResolutionService_RenderProductIssues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueResolutionServiceServer).RenderProductIssues(ctx, req.(*RenderProductIssuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueResolutionService_TriggerAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueResolutionServiceServer).TriggerAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IssueResolutionService_TriggerAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueResolutionServiceServer).TriggerAction(ctx, req.(*TriggerActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IssueResolutionService_ServiceDesc is the grpc.ServiceDesc for IssueResolutionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IssueResolutionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.merchant.issueresolution.v1beta.IssueResolutionService",
	HandlerType: (*IssueResolutionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RenderAccountIssues",
			Handler:    _IssueResolutionService_RenderAccountIssues_Handler,
		},
		{
			MethodName: "RenderProductIssues",
			Handler:    _IssueResolutionService_RenderProductIssues_Handler,
		},
		{
			MethodName: "TriggerAction",
			Handler:    _IssueResolutionService_TriggerAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/merchant/issueresolution/v1beta/issueresolution.proto",
}
