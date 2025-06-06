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
// source: google/chat/v1/space_read_state.proto

package chatpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// A user's read state within a space, used to identify read and unread
// messages.
type SpaceReadState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the space read state.
	//
	// Format: `users/{user}/spaces/{space}/spaceReadState`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The time when the user's space read state was updated. Usually
	// this corresponds with either the timestamp of the last read message, or a
	// timestamp specified by the user to mark the last read position in a space.
	LastReadTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_read_time,json=lastReadTime,proto3" json:"last_read_time,omitempty"`
}

func (x *SpaceReadState) Reset() {
	*x = SpaceReadState{}
	mi := &file_google_chat_v1_space_read_state_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpaceReadState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceReadState) ProtoMessage() {}

func (x *SpaceReadState) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_space_read_state_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceReadState.ProtoReflect.Descriptor instead.
func (*SpaceReadState) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_space_read_state_proto_rawDescGZIP(), []int{0}
}

func (x *SpaceReadState) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpaceReadState) GetLastReadTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastReadTime
	}
	return nil
}

// Request message for GetSpaceReadState API.
type GetSpaceReadStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the space read state to retrieve.
	//
	// Only supports getting read state for the calling user.
	//
	// To refer to the calling user, set one of the following:
	//
	// - The `me` alias. For example, `users/me/spaces/{space}/spaceReadState`.
	//
	// - Their Workspace email address. For example,
	// `users/user@example.com/spaces/{space}/spaceReadState`.
	//
	// - Their user id. For example,
	// `users/123456789/spaces/{space}/spaceReadState`.
	//
	// Format: users/{user}/spaces/{space}/spaceReadState
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSpaceReadStateRequest) Reset() {
	*x = GetSpaceReadStateRequest{}
	mi := &file_google_chat_v1_space_read_state_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSpaceReadStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpaceReadStateRequest) ProtoMessage() {}

func (x *GetSpaceReadStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_space_read_state_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpaceReadStateRequest.ProtoReflect.Descriptor instead.
func (*GetSpaceReadStateRequest) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_space_read_state_proto_rawDescGZIP(), []int{1}
}

func (x *GetSpaceReadStateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for UpdateSpaceReadState API.
type UpdateSpaceReadStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The space read state and fields to update.
	//
	// Only supports updating read state for the calling user.
	//
	// To refer to the calling user, set one of the following:
	//
	// - The `me` alias. For example, `users/me/spaces/{space}/spaceReadState`.
	//
	// - Their Workspace email address. For example,
	// `users/user@example.com/spaces/{space}/spaceReadState`.
	//
	// - Their user id. For example,
	// `users/123456789/spaces/{space}/spaceReadState`.
	//
	// Format: users/{user}/spaces/{space}/spaceReadState
	SpaceReadState *SpaceReadState `protobuf:"bytes,1,opt,name=space_read_state,json=spaceReadState,proto3" json:"space_read_state,omitempty"`
	// Required. The field paths to update. Currently supported field paths:
	//
	// - `last_read_time`
	//
	// When the `last_read_time` is before the latest message create time, the
	// space appears as unread in the UI.
	//
	// To mark the space as read, set `last_read_time` to any value later (larger)
	// than the latest message create time. The `last_read_time` is coerced to
	// match the latest message create time. Note that the space read state only
	// affects the read state of messages that are visible in the space's
	// top-level conversation. Replies in threads are unaffected by this
	// timestamp, and instead rely on the thread read state.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateSpaceReadStateRequest) Reset() {
	*x = UpdateSpaceReadStateRequest{}
	mi := &file_google_chat_v1_space_read_state_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSpaceReadStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSpaceReadStateRequest) ProtoMessage() {}

func (x *UpdateSpaceReadStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_space_read_state_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSpaceReadStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateSpaceReadStateRequest) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_space_read_state_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSpaceReadStateRequest) GetSpaceReadState() *SpaceReadState {
	if x != nil {
		return x.SpaceReadState
	}
	return nil
}

func (x *UpdateSpaceReadStateRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_google_chat_v1_space_read_state_proto protoreflect.FileDescriptor

var file_google_chat_v1_space_read_state_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x0e, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a,
	0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x3a, 0x63, 0xea, 0x41, 0x60, 0x0a, 0x22, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x7d, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x32, 0x0e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5a, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x10, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0xac, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x3b, 0x63, 0x68, 0x61, 0x74,
	0x70, 0x62, 0xa2, 0x02, 0x0b, 0x44, 0x59, 0x4e, 0x41, 0x50, 0x49, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x70, 0x70, 0x73, 0x5c, 0x43, 0x68, 0x61, 0x74, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x16, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x73, 0x3a, 0x3a, 0x43, 0x68, 0x61,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_chat_v1_space_read_state_proto_rawDescOnce sync.Once
	file_google_chat_v1_space_read_state_proto_rawDescData = file_google_chat_v1_space_read_state_proto_rawDesc
)

func file_google_chat_v1_space_read_state_proto_rawDescGZIP() []byte {
	file_google_chat_v1_space_read_state_proto_rawDescOnce.Do(func() {
		file_google_chat_v1_space_read_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_chat_v1_space_read_state_proto_rawDescData)
	})
	return file_google_chat_v1_space_read_state_proto_rawDescData
}

var file_google_chat_v1_space_read_state_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_chat_v1_space_read_state_proto_goTypes = []any{
	(*SpaceReadState)(nil),              // 0: google.chat.v1.SpaceReadState
	(*GetSpaceReadStateRequest)(nil),    // 1: google.chat.v1.GetSpaceReadStateRequest
	(*UpdateSpaceReadStateRequest)(nil), // 2: google.chat.v1.UpdateSpaceReadStateRequest
	(*timestamppb.Timestamp)(nil),       // 3: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil),       // 4: google.protobuf.FieldMask
}
var file_google_chat_v1_space_read_state_proto_depIdxs = []int32{
	3, // 0: google.chat.v1.SpaceReadState.last_read_time:type_name -> google.protobuf.Timestamp
	0, // 1: google.chat.v1.UpdateSpaceReadStateRequest.space_read_state:type_name -> google.chat.v1.SpaceReadState
	4, // 2: google.chat.v1.UpdateSpaceReadStateRequest.update_mask:type_name -> google.protobuf.FieldMask
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_chat_v1_space_read_state_proto_init() }
func file_google_chat_v1_space_read_state_proto_init() {
	if File_google_chat_v1_space_read_state_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_chat_v1_space_read_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_chat_v1_space_read_state_proto_goTypes,
		DependencyIndexes: file_google_chat_v1_space_read_state_proto_depIdxs,
		MessageInfos:      file_google_chat_v1_space_read_state_proto_msgTypes,
	}.Build()
	File_google_chat_v1_space_read_state_proto = out.File
	file_google_chat_v1_space_read_state_proto_rawDesc = nil
	file_google_chat_v1_space_read_state_proto_goTypes = nil
	file_google_chat_v1_space_read_state_proto_depIdxs = nil
}
