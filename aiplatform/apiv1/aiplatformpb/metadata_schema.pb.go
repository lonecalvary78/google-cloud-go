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
// source: google/cloud/aiplatform/v1/metadata_schema.proto

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

// Describes the type of the MetadataSchema.
type MetadataSchema_MetadataSchemaType int32

const (
	// Unspecified type for the MetadataSchema.
	MetadataSchema_METADATA_SCHEMA_TYPE_UNSPECIFIED MetadataSchema_MetadataSchemaType = 0
	// A type indicating that the MetadataSchema will be used by Artifacts.
	MetadataSchema_ARTIFACT_TYPE MetadataSchema_MetadataSchemaType = 1
	// A typee indicating that the MetadataSchema will be used by Executions.
	MetadataSchema_EXECUTION_TYPE MetadataSchema_MetadataSchemaType = 2
	// A state indicating that the MetadataSchema will be used by Contexts.
	MetadataSchema_CONTEXT_TYPE MetadataSchema_MetadataSchemaType = 3
)

// Enum value maps for MetadataSchema_MetadataSchemaType.
var (
	MetadataSchema_MetadataSchemaType_name = map[int32]string{
		0: "METADATA_SCHEMA_TYPE_UNSPECIFIED",
		1: "ARTIFACT_TYPE",
		2: "EXECUTION_TYPE",
		3: "CONTEXT_TYPE",
	}
	MetadataSchema_MetadataSchemaType_value = map[string]int32{
		"METADATA_SCHEMA_TYPE_UNSPECIFIED": 0,
		"ARTIFACT_TYPE":                    1,
		"EXECUTION_TYPE":                   2,
		"CONTEXT_TYPE":                     3,
	}
)

func (x MetadataSchema_MetadataSchemaType) Enum() *MetadataSchema_MetadataSchemaType {
	p := new(MetadataSchema_MetadataSchemaType)
	*p = x
	return p
}

func (x MetadataSchema_MetadataSchemaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetadataSchema_MetadataSchemaType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_metadata_schema_proto_enumTypes[0].Descriptor()
}

func (MetadataSchema_MetadataSchemaType) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_metadata_schema_proto_enumTypes[0]
}

func (x MetadataSchema_MetadataSchemaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetadataSchema_MetadataSchemaType.Descriptor instead.
func (MetadataSchema_MetadataSchemaType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDescGZIP(), []int{0, 0}
}

// Instance of a general MetadataSchema.
type MetadataSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the MetadataSchema.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The version of the MetadataSchema. The version's format must match
	// the following regular expression: `^[0-9]+[.][0-9]+[.][0-9]+$`, which would
	// allow to order/compare different versions. Example: 1.0.0, 1.0.1, etc.
	SchemaVersion string `protobuf:"bytes,2,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	// Required. The raw YAML string representation of the MetadataSchema. The
	// combination of [MetadataSchema.version] and the schema name given by
	// `title` in [MetadataSchema.schema] must be unique within a MetadataStore.
	//
	// The schema is defined as an OpenAPI 3.0.2
	// [MetadataSchema
	// Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#schemaObject)
	Schema string `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	// The type of the MetadataSchema. This is a property that identifies which
	// metadata types will use the MetadataSchema.
	SchemaType MetadataSchema_MetadataSchemaType `protobuf:"varint,4,opt,name=schema_type,json=schemaType,proto3,enum=google.cloud.aiplatform.v1.MetadataSchema_MetadataSchemaType" json:"schema_type,omitempty"`
	// Output only. Timestamp when this MetadataSchema was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Description of the Metadata Schema
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *MetadataSchema) Reset() {
	*x = MetadataSchema{}
	mi := &file_google_cloud_aiplatform_v1_metadata_schema_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetadataSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataSchema) ProtoMessage() {}

func (x *MetadataSchema) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_metadata_schema_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataSchema.ProtoReflect.Descriptor instead.
func (*MetadataSchema) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataSchema) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetadataSchema) GetSchemaVersion() string {
	if x != nil {
		return x.SchemaVersion
	}
	return ""
}

func (x *MetadataSchema) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *MetadataSchema) GetSchemaType() MetadataSchema_MetadataSchemaType {
	if x != nil {
		return x.SchemaType
	}
	return MetadataSchema_METADATA_SCHEMA_TYPE_UNSPECIFIED
}

func (x *MetadataSchema) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *MetadataSchema) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_google_cloud_aiplatform_v1_metadata_schema_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x04, 0x0a, 0x0e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x5e, 0x0a, 0x0b, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x73, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41,
	0x52, 0x54, 0x49, 0x46, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x03, 0x3a, 0x99, 0x01, 0xea, 0x41, 0x95, 0x01, 0x0a, 0x28, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x69, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x7b,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x7d,
	0x42, 0xd1, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x42, 0x13, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDescData = file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_metadata_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1_metadata_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1_metadata_schema_proto_goTypes = []any{
	(MetadataSchema_MetadataSchemaType)(0), // 0: google.cloud.aiplatform.v1.MetadataSchema.MetadataSchemaType
	(*MetadataSchema)(nil),                 // 1: google.cloud.aiplatform.v1.MetadataSchema
	(*timestamppb.Timestamp)(nil),          // 2: google.protobuf.Timestamp
}
var file_google_cloud_aiplatform_v1_metadata_schema_proto_depIdxs = []int32{
	0, // 0: google.cloud.aiplatform.v1.MetadataSchema.schema_type:type_name -> google.cloud.aiplatform.v1.MetadataSchema.MetadataSchemaType
	2, // 1: google.cloud.aiplatform.v1.MetadataSchema.create_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_metadata_schema_proto_init() }
func file_google_cloud_aiplatform_v1_metadata_schema_proto_init() {
	if File_google_cloud_aiplatform_v1_metadata_schema_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_metadata_schema_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_metadata_schema_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1_metadata_schema_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1_metadata_schema_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_metadata_schema_proto = out.File
	file_google_cloud_aiplatform_v1_metadata_schema_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_metadata_schema_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_metadata_schema_proto_depIdxs = nil
}
