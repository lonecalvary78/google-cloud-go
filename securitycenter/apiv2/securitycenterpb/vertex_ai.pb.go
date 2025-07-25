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
// source: google/cloud/securitycenter/v2/vertex_ai.proto

package securitycenterpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Vertex AI-related information associated with the finding.
type VertexAi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Datasets associated with the finding.
	Datasets []*VertexAi_Dataset `protobuf:"bytes,1,rep,name=datasets,proto3" json:"datasets,omitempty"`
	// Pipelines associated with the finding.
	Pipelines []*VertexAi_Pipeline `protobuf:"bytes,2,rep,name=pipelines,proto3" json:"pipelines,omitempty"`
}

func (x *VertexAi) Reset() {
	*x = VertexAi{}
	mi := &file_google_cloud_securitycenter_v2_vertex_ai_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VertexAi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VertexAi) ProtoMessage() {}

func (x *VertexAi) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_vertex_ai_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VertexAi.ProtoReflect.Descriptor instead.
func (*VertexAi) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDescGZIP(), []int{0}
}

func (x *VertexAi) GetDatasets() []*VertexAi_Dataset {
	if x != nil {
		return x.Datasets
	}
	return nil
}

func (x *VertexAi) GetPipelines() []*VertexAi_Pipeline {
	if x != nil {
		return x.Pipelines
	}
	return nil
}

// Vertex AI dataset associated with the finding.
type VertexAi_Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the dataset, e.g.
	// projects/{project}/locations/{location}/datasets/2094040236064505856
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The user defined display name of dataset, e.g. plants-dataset
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Data source, such as a BigQuery source URI, e.g.
	// bq://scc-nexus-test.AIPPtest.gsod
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *VertexAi_Dataset) Reset() {
	*x = VertexAi_Dataset{}
	mi := &file_google_cloud_securitycenter_v2_vertex_ai_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VertexAi_Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VertexAi_Dataset) ProtoMessage() {}

func (x *VertexAi_Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_vertex_ai_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VertexAi_Dataset.ProtoReflect.Descriptor instead.
func (*VertexAi_Dataset) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDescGZIP(), []int{0, 0}
}

func (x *VertexAi_Dataset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VertexAi_Dataset) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *VertexAi_Dataset) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

// Vertex AI training pipeline associated with the finding.
type VertexAi_Pipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the pipeline, e.g.
	// projects/{project}/locations/{location}/trainingPipelines/5253428229225578496
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The user-defined display name of pipeline, e.g. plants-classification
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *VertexAi_Pipeline) Reset() {
	*x = VertexAi_Pipeline{}
	mi := &file_google_cloud_securitycenter_v2_vertex_ai_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VertexAi_Pipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VertexAi_Pipeline) ProtoMessage() {}

func (x *VertexAi_Pipeline) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_vertex_ai_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VertexAi_Pipeline.ProtoReflect.Descriptor instead.
func (*VertexAi_Pipeline) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDescGZIP(), []int{0, 1}
}

func (x *VertexAi_Pipeline) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VertexAi_Pipeline) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_google_cloud_securitycenter_v2_vertex_ai_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x5f, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x22, 0xc6, 0x02, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x41, 0x69, 0x12, 0x4c, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x41, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x09, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e,
	0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x41, 0x69, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x09, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x1a, 0x58, 0x0a, 0x07,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x41, 0x0a, 0x08, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xe7, 0x01, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x42, 0x0d, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x41, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x32, 0xca, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x32, 0xea,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDescData = file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDesc
)

func file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDescData
}

var file_google_cloud_securitycenter_v2_vertex_ai_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_securitycenter_v2_vertex_ai_proto_goTypes = []any{
	(*VertexAi)(nil),          // 0: google.cloud.securitycenter.v2.VertexAi
	(*VertexAi_Dataset)(nil),  // 1: google.cloud.securitycenter.v2.VertexAi.Dataset
	(*VertexAi_Pipeline)(nil), // 2: google.cloud.securitycenter.v2.VertexAi.Pipeline
}
var file_google_cloud_securitycenter_v2_vertex_ai_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v2.VertexAi.datasets:type_name -> google.cloud.securitycenter.v2.VertexAi.Dataset
	2, // 1: google.cloud.securitycenter.v2.VertexAi.pipelines:type_name -> google.cloud.securitycenter.v2.VertexAi.Pipeline
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v2_vertex_ai_proto_init() }
func file_google_cloud_securitycenter_v2_vertex_ai_proto_init() {
	if File_google_cloud_securitycenter_v2_vertex_ai_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v2_vertex_ai_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v2_vertex_ai_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v2_vertex_ai_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v2_vertex_ai_proto = out.File
	file_google_cloud_securitycenter_v2_vertex_ai_proto_rawDesc = nil
	file_google_cloud_securitycenter_v2_vertex_ai_proto_goTypes = nil
	file_google_cloud_securitycenter_v2_vertex_ai_proto_depIdxs = nil
}
