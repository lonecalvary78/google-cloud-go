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
// source: google/cloud/aiplatform/v1beta1/example_store.proto

package aiplatformpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// The logic to use for filtering.
type ExamplesArrayFilter_ArrayOperator int32

const (
	// Not specified. This value should not be used.
	ExamplesArrayFilter_ARRAY_OPERATOR_UNSPECIFIED ExamplesArrayFilter_ArrayOperator = 0
	// The metadata array field in the example must contain at least one of the
	// values.
	ExamplesArrayFilter_CONTAINS_ANY ExamplesArrayFilter_ArrayOperator = 1
	// The metadata array field in the example must contain all of the values.
	ExamplesArrayFilter_CONTAINS_ALL ExamplesArrayFilter_ArrayOperator = 2
)

// Enum value maps for ExamplesArrayFilter_ArrayOperator.
var (
	ExamplesArrayFilter_ArrayOperator_name = map[int32]string{
		0: "ARRAY_OPERATOR_UNSPECIFIED",
		1: "CONTAINS_ANY",
		2: "CONTAINS_ALL",
	}
	ExamplesArrayFilter_ArrayOperator_value = map[string]int32{
		"ARRAY_OPERATOR_UNSPECIFIED": 0,
		"CONTAINS_ANY":               1,
		"CONTAINS_ALL":               2,
	}
)

func (x ExamplesArrayFilter_ArrayOperator) Enum() *ExamplesArrayFilter_ArrayOperator {
	p := new(ExamplesArrayFilter_ArrayOperator)
	*p = x
	return p
}

func (x ExamplesArrayFilter_ArrayOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExamplesArrayFilter_ArrayOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1beta1_example_store_proto_enumTypes[0].Descriptor()
}

func (ExamplesArrayFilter_ArrayOperator) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1beta1_example_store_proto_enumTypes[0]
}

func (x ExamplesArrayFilter_ArrayOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExamplesArrayFilter_ArrayOperator.Descriptor instead.
func (ExamplesArrayFilter_ArrayOperator) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescGZIP(), []int{4, 0}
}

// Represents an executable service to manage and retrieve examples.
type ExampleStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the ExampleStore. This is a unique
	// identifier. Format:
	// projects/{project}/locations/{location}/exampleStores/{example_store}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Display name of the ExampleStore.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. Description of the ExampleStore.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Timestamp when this ExampleStore was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this ExampleStore was most recently updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Required. Example Store config.
	ExampleStoreConfig *ExampleStoreConfig `protobuf:"bytes,6,opt,name=example_store_config,json=exampleStoreConfig,proto3" json:"example_store_config,omitempty"`
}

func (x *ExampleStore) Reset() {
	*x = ExampleStore{}
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleStore) ProtoMessage() {}

func (x *ExampleStore) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleStore.ProtoReflect.Descriptor instead.
func (*ExampleStore) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExampleStore) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ExampleStore) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ExampleStore) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ExampleStore) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ExampleStore) GetExampleStoreConfig() *ExampleStoreConfig {
	if x != nil {
		return x.ExampleStoreConfig
	}
	return nil
}

// Configuration for the Example Store.
type ExampleStoreConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The embedding model to be used for vector embedding.
	// Immutable.
	// Supported models:
	// * "textembedding-gecko@003"
	// * "text-embedding-004"
	// * "text-embedding-005"
	// * "text-multilingual-embedding-002"
	VertexEmbeddingModel string `protobuf:"bytes,1,opt,name=vertex_embedding_model,json=vertexEmbeddingModel,proto3" json:"vertex_embedding_model,omitempty"`
}

func (x *ExampleStoreConfig) Reset() {
	*x = ExampleStoreConfig{}
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleStoreConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleStoreConfig) ProtoMessage() {}

func (x *ExampleStoreConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleStoreConfig.ProtoReflect.Descriptor instead.
func (*ExampleStoreConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescGZIP(), []int{1}
}

func (x *ExampleStoreConfig) GetVertexEmbeddingModel() string {
	if x != nil {
		return x.VertexEmbeddingModel
	}
	return ""
}

// The metadata filters that will be used to remove or fetch
// StoredContentsExamples. If a field is unspecified, then no filtering for that
// field will be applied.
type StoredContentsExampleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The search keys for filtering. Only examples with one of the
	// specified search keys
	// ([StoredContentsExample.search_key][google.cloud.aiplatform.v1beta1.StoredContentsExample.search_key])
	// are eligible to be returned.
	SearchKeys []string `protobuf:"bytes,1,rep,name=search_keys,json=searchKeys,proto3" json:"search_keys,omitempty"`
	// Optional. The function names for filtering.
	FunctionNames *ExamplesArrayFilter `protobuf:"bytes,2,opt,name=function_names,json=functionNames,proto3" json:"function_names,omitempty"`
}

func (x *StoredContentsExampleFilter) Reset() {
	*x = StoredContentsExampleFilter{}
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoredContentsExampleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredContentsExampleFilter) ProtoMessage() {}

func (x *StoredContentsExampleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredContentsExampleFilter.ProtoReflect.Descriptor instead.
func (*StoredContentsExampleFilter) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescGZIP(), []int{2}
}

func (x *StoredContentsExampleFilter) GetSearchKeys() []string {
	if x != nil {
		return x.SearchKeys
	}
	return nil
}

func (x *StoredContentsExampleFilter) GetFunctionNames() *ExamplesArrayFilter {
	if x != nil {
		return x.FunctionNames
	}
	return nil
}

// The metadata filters that will be used to search StoredContentsExamples.
// If a field is unspecified, then no filtering for that field will be applied
type StoredContentsExampleParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The query to use to retrieve similar StoredContentsExamples.
	//
	// Types that are assignable to Query:
	//
	//	*StoredContentsExampleParameters_SearchKey
	//	*StoredContentsExampleParameters_ContentSearchKey_
	Query isStoredContentsExampleParameters_Query `protobuf_oneof:"query"`
	// Optional. The function names for filtering.
	FunctionNames *ExamplesArrayFilter `protobuf:"bytes,3,opt,name=function_names,json=functionNames,proto3" json:"function_names,omitempty"`
}

func (x *StoredContentsExampleParameters) Reset() {
	*x = StoredContentsExampleParameters{}
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoredContentsExampleParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredContentsExampleParameters) ProtoMessage() {}

func (x *StoredContentsExampleParameters) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredContentsExampleParameters.ProtoReflect.Descriptor instead.
func (*StoredContentsExampleParameters) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescGZIP(), []int{3}
}

func (m *StoredContentsExampleParameters) GetQuery() isStoredContentsExampleParameters_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *StoredContentsExampleParameters) GetSearchKey() string {
	if x, ok := x.GetQuery().(*StoredContentsExampleParameters_SearchKey); ok {
		return x.SearchKey
	}
	return ""
}

func (x *StoredContentsExampleParameters) GetContentSearchKey() *StoredContentsExampleParameters_ContentSearchKey {
	if x, ok := x.GetQuery().(*StoredContentsExampleParameters_ContentSearchKey_); ok {
		return x.ContentSearchKey
	}
	return nil
}

func (x *StoredContentsExampleParameters) GetFunctionNames() *ExamplesArrayFilter {
	if x != nil {
		return x.FunctionNames
	}
	return nil
}

type isStoredContentsExampleParameters_Query interface {
	isStoredContentsExampleParameters_Query()
}

type StoredContentsExampleParameters_SearchKey struct {
	// The exact search key to use for retrieval.
	SearchKey string `protobuf:"bytes,1,opt,name=search_key,json=searchKey,proto3,oneof"`
}

type StoredContentsExampleParameters_ContentSearchKey_ struct {
	// The chat history to use to generate the search key for retrieval.
	ContentSearchKey *StoredContentsExampleParameters_ContentSearchKey `protobuf:"bytes,2,opt,name=content_search_key,json=contentSearchKey,proto3,oneof"`
}

func (*StoredContentsExampleParameters_SearchKey) isStoredContentsExampleParameters_Query() {}

func (*StoredContentsExampleParameters_ContentSearchKey_) isStoredContentsExampleParameters_Query() {}

// Filters for examples' array metadata fields. An array field is example
// metadata where multiple values are attributed to a single example.
type ExamplesArrayFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The values by which to filter examples.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// Required. The operator logic to use for filtering.
	ArrayOperator ExamplesArrayFilter_ArrayOperator `protobuf:"varint,2,opt,name=array_operator,json=arrayOperator,proto3,enum=google.cloud.aiplatform.v1beta1.ExamplesArrayFilter_ArrayOperator" json:"array_operator,omitempty"`
}

func (x *ExamplesArrayFilter) Reset() {
	*x = ExamplesArrayFilter{}
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExamplesArrayFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExamplesArrayFilter) ProtoMessage() {}

func (x *ExamplesArrayFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExamplesArrayFilter.ProtoReflect.Descriptor instead.
func (*ExamplesArrayFilter) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescGZIP(), []int{4}
}

func (x *ExamplesArrayFilter) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *ExamplesArrayFilter) GetArrayOperator() ExamplesArrayFilter_ArrayOperator {
	if x != nil {
		return x.ArrayOperator
	}
	return ExamplesArrayFilter_ARRAY_OPERATOR_UNSPECIFIED
}

// The chat history to use to generate the search key for retrieval.
type StoredContentsExampleParameters_ContentSearchKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The conversation for generating a search key.
	Contents []*Content `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
	// Required. The method of generating a search key.
	SearchKeyGenerationMethod *StoredContentsExample_SearchKeyGenerationMethod `protobuf:"bytes,2,opt,name=search_key_generation_method,json=searchKeyGenerationMethod,proto3" json:"search_key_generation_method,omitempty"`
}

func (x *StoredContentsExampleParameters_ContentSearchKey) Reset() {
	*x = StoredContentsExampleParameters_ContentSearchKey{}
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoredContentsExampleParameters_ContentSearchKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredContentsExampleParameters_ContentSearchKey) ProtoMessage() {}

func (x *StoredContentsExampleParameters_ContentSearchKey) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredContentsExampleParameters_ContentSearchKey.ProtoReflect.Descriptor instead.
func (*StoredContentsExampleParameters_ContentSearchKey) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescGZIP(), []int{3, 0}
}

func (x *StoredContentsExampleParameters_ContentSearchKey) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *StoredContentsExampleParameters_ContentSearchKey) GetSearchKeyGenerationMethod() *StoredContentsExample_SearchKeyGenerationMethod {
	if x != nil {
		return x.SearchKeyGenerationMethod
	}
	return nil
}

var File_google_cloud_aiplatform_v1beta1_example_store_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf9, 0x03, 0x0a, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x6a,
	0x0a, 0x14, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x12, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x90, 0x01, 0xea, 0x41, 0x8c,
	0x01, 0x0a, 0x26, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x45, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x2f, 0x7b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d,
	0x2a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x32,
	0x0c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x4f, 0x0a,
	0x12, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x16, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x5f, 0x65, 0x6d,
	0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x14, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78,
	0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xa5,
	0x01, 0x0a, 0x1b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x60, 0x0a, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0xaa, 0x04, 0x0a, 0x1f, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x81, 0x01, 0x0a, 0x12,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x10, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12,
	0x60, 0x0a, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x1a, 0xf6, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x49, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x96, 0x01, 0x0a, 0x1c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x19, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x22, 0xf7, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x6e, 0x0a, 0x0e, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x53, 0x0a, 0x0d, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x52, 0x52,
	0x41, 0x59, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e,
	0x54, 0x41, 0x49, 0x4e, 0x53, 0x5f, 0x41, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43,
	0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x42, 0xe8, 0x01,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x11, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41,
	0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61,
	0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_example_store_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_aiplatform_v1beta1_example_store_proto_goTypes = []any{
	(ExamplesArrayFilter_ArrayOperator)(0),                   // 0: google.cloud.aiplatform.v1beta1.ExamplesArrayFilter.ArrayOperator
	(*ExampleStore)(nil),                                     // 1: google.cloud.aiplatform.v1beta1.ExampleStore
	(*ExampleStoreConfig)(nil),                               // 2: google.cloud.aiplatform.v1beta1.ExampleStoreConfig
	(*StoredContentsExampleFilter)(nil),                      // 3: google.cloud.aiplatform.v1beta1.StoredContentsExampleFilter
	(*StoredContentsExampleParameters)(nil),                  // 4: google.cloud.aiplatform.v1beta1.StoredContentsExampleParameters
	(*ExamplesArrayFilter)(nil),                              // 5: google.cloud.aiplatform.v1beta1.ExamplesArrayFilter
	(*StoredContentsExampleParameters_ContentSearchKey)(nil), // 6: google.cloud.aiplatform.v1beta1.StoredContentsExampleParameters.ContentSearchKey
	(*timestamppb.Timestamp)(nil),                            // 7: google.protobuf.Timestamp
	(*Content)(nil),                                          // 8: google.cloud.aiplatform.v1beta1.Content
	(*StoredContentsExample_SearchKeyGenerationMethod)(nil),  // 9: google.cloud.aiplatform.v1beta1.StoredContentsExample.SearchKeyGenerationMethod
}
var file_google_cloud_aiplatform_v1beta1_example_store_proto_depIdxs = []int32{
	7, // 0: google.cloud.aiplatform.v1beta1.ExampleStore.create_time:type_name -> google.protobuf.Timestamp
	7, // 1: google.cloud.aiplatform.v1beta1.ExampleStore.update_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.cloud.aiplatform.v1beta1.ExampleStore.example_store_config:type_name -> google.cloud.aiplatform.v1beta1.ExampleStoreConfig
	5, // 3: google.cloud.aiplatform.v1beta1.StoredContentsExampleFilter.function_names:type_name -> google.cloud.aiplatform.v1beta1.ExamplesArrayFilter
	6, // 4: google.cloud.aiplatform.v1beta1.StoredContentsExampleParameters.content_search_key:type_name -> google.cloud.aiplatform.v1beta1.StoredContentsExampleParameters.ContentSearchKey
	5, // 5: google.cloud.aiplatform.v1beta1.StoredContentsExampleParameters.function_names:type_name -> google.cloud.aiplatform.v1beta1.ExamplesArrayFilter
	0, // 6: google.cloud.aiplatform.v1beta1.ExamplesArrayFilter.array_operator:type_name -> google.cloud.aiplatform.v1beta1.ExamplesArrayFilter.ArrayOperator
	8, // 7: google.cloud.aiplatform.v1beta1.StoredContentsExampleParameters.ContentSearchKey.contents:type_name -> google.cloud.aiplatform.v1beta1.Content
	9, // 8: google.cloud.aiplatform.v1beta1.StoredContentsExampleParameters.ContentSearchKey.search_key_generation_method:type_name -> google.cloud.aiplatform.v1beta1.StoredContentsExample.SearchKeyGenerationMethod
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_example_store_proto_init() }
func file_google_cloud_aiplatform_v1beta1_example_store_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_example_store_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_content_proto_init()
	file_google_cloud_aiplatform_v1beta1_example_proto_init()
	file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes[3].OneofWrappers = []any{
		(*StoredContentsExampleParameters_SearchKey)(nil),
		(*StoredContentsExampleParameters_ContentSearchKey_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_example_store_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_example_store_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1beta1_example_store_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_example_store_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_example_store_proto = out.File
	file_google_cloud_aiplatform_v1beta1_example_store_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_example_store_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_example_store_proto_depIdxs = nil
}
