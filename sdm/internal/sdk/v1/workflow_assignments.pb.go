// Copyright 2020 StrongDM Inc
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
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: workflow_assignments.proto

package v1

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

// WorkflowAssignmentsListRequest specifies criteria for retrieving a list of
// WorkflowAssignment records
type WorkflowAssignmentsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *WorkflowAssignmentsListRequest) Reset() {
	*x = WorkflowAssignmentsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_assignments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowAssignmentsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowAssignmentsListRequest) ProtoMessage() {}

func (x *WorkflowAssignmentsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_assignments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowAssignmentsListRequest.ProtoReflect.Descriptor instead.
func (*WorkflowAssignmentsListRequest) Descriptor() ([]byte, []int) {
	return file_workflow_assignments_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowAssignmentsListRequest) GetMeta() *ListRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkflowAssignmentsListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// WorkflowAssignmentsListResponse returns a list of WorkflowAssignment records that meet
// the criteria of a WorkflowAssignmentsListRequest.
type WorkflowAssignmentsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The matching workflows.
	WorkflowAssignments []*WorkflowAssignment `protobuf:"bytes,2,rep,name=workflow_assignments,json=workflowAssignments,proto3" json:"workflow_assignments,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *WorkflowAssignmentsListResponse) Reset() {
	*x = WorkflowAssignmentsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_assignments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowAssignmentsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowAssignmentsListResponse) ProtoMessage() {}

func (x *WorkflowAssignmentsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_assignments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowAssignmentsListResponse.ProtoReflect.Descriptor instead.
func (*WorkflowAssignmentsListResponse) Descriptor() ([]byte, []int) {
	return file_workflow_assignments_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowAssignmentsListResponse) GetMeta() *ListResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkflowAssignmentsListResponse) GetWorkflowAssignments() []*WorkflowAssignment {
	if x != nil {
		return x.WorkflowAssignments
	}
	return nil
}

func (x *WorkflowAssignmentsListResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// WorkflowAssignment links a Resource to a Workflow. The assigned resources are those that a user can request
// access to via the workflow.
type WorkflowAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workflow id.
	WorkflowId string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	// The resource id.
	ResourceId string `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *WorkflowAssignment) Reset() {
	*x = WorkflowAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_assignments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowAssignment) ProtoMessage() {}

func (x *WorkflowAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_assignments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowAssignment.ProtoReflect.Descriptor instead.
func (*WorkflowAssignment) Descriptor() ([]byte, []int) {
	return file_workflow_assignments_proto_rawDescGZIP(), []int{2}
}

func (x *WorkflowAssignment) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *WorkflowAssignment) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

var File_workflow_assignments_proto protoreflect.FileDescriptor

var file_workflow_assignments_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x1e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3,
	0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a,
	0x10, 0xfa, 0xf8, 0xb3, 0x07, 0x0b, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xd2, 0xf3, 0xb3, 0x07, 0x01,
	0x2a, 0x22, 0x9c, 0x02, 0x0a, 0x1f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x55, 0x0a, 0x14, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05,
	0xb8, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x62, 0x0a, 0x0a, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3, 0x07,
	0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x10,
	0xfa, 0xf8, 0xb3, 0x07, 0x0b, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a,
	0x22, 0x84, 0x01, 0x0a, 0x12, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0xf8,
	0xb3, 0x07, 0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f,
	0xf2, 0xf8, 0xb3, 0x07, 0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x3a, 0x0a, 0xfa, 0xf8, 0xb3,
	0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x32, 0xf5, 0x01, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x80, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2f, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07, 0x03, 0x67, 0x65, 0x74,
	0x82, 0xf9, 0xb3, 0x07, 0x1d, 0xaa, 0xf3, 0xb3, 0x07, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x1a, 0x5b, 0xca, 0xf9, 0xb3, 0x07, 0x17, 0xc2, 0xf9, 0xb3, 0x07, 0x12, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0xca, 0xf9, 0xb3, 0x07, 0x08, 0xd2, 0xf9, 0xb3, 0x07, 0x03, 0x77, 0x61, 0x2d, 0xca, 0xf9, 0xb3,
	0x07, 0x06, 0xca, 0xf9, 0xb3, 0x07, 0x01, 0x2a, 0xca, 0xf9, 0xb3, 0x07, 0x18, 0xca, 0xf9, 0xb3,
	0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0xca, 0xf9, 0xb3, 0x07, 0x05, 0xe8, 0xf9, 0xb3, 0x07, 0x01, 0x42,
	0x97, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x42, 0x1b, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x50, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x76, 0x33, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0xc2, 0x92, 0xb4, 0x07, 0x06, 0xa2, 0x8c, 0xb4, 0x07, 0x01, 0x2a, 0xc2, 0x92, 0xb4, 0x07,
	0x18, 0xa2, 0x8c, 0xb4, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_workflow_assignments_proto_rawDescOnce sync.Once
	file_workflow_assignments_proto_rawDescData = file_workflow_assignments_proto_rawDesc
)

func file_workflow_assignments_proto_rawDescGZIP() []byte {
	file_workflow_assignments_proto_rawDescOnce.Do(func() {
		file_workflow_assignments_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_assignments_proto_rawDescData)
	})
	return file_workflow_assignments_proto_rawDescData
}

var file_workflow_assignments_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_workflow_assignments_proto_goTypes = []interface{}{
	(*WorkflowAssignmentsListRequest)(nil),  // 0: v1.WorkflowAssignmentsListRequest
	(*WorkflowAssignmentsListResponse)(nil), // 1: v1.WorkflowAssignmentsListResponse
	(*WorkflowAssignment)(nil),              // 2: v1.WorkflowAssignment
	(*ListRequestMetadata)(nil),             // 3: v1.ListRequestMetadata
	(*ListResponseMetadata)(nil),            // 4: v1.ListResponseMetadata
	(*RateLimitMetadata)(nil),               // 5: v1.RateLimitMetadata
}
var file_workflow_assignments_proto_depIdxs = []int32{
	3, // 0: v1.WorkflowAssignmentsListRequest.meta:type_name -> v1.ListRequestMetadata
	4, // 1: v1.WorkflowAssignmentsListResponse.meta:type_name -> v1.ListResponseMetadata
	2, // 2: v1.WorkflowAssignmentsListResponse.workflow_assignments:type_name -> v1.WorkflowAssignment
	5, // 3: v1.WorkflowAssignmentsListResponse.rate_limit:type_name -> v1.RateLimitMetadata
	0, // 4: v1.WorkflowAssignments.List:input_type -> v1.WorkflowAssignmentsListRequest
	1, // 5: v1.WorkflowAssignments.List:output_type -> v1.WorkflowAssignmentsListResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_workflow_assignments_proto_init() }
func file_workflow_assignments_proto_init() {
	if File_workflow_assignments_proto != nil {
		return
	}
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_workflow_assignments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowAssignmentsListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workflow_assignments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowAssignmentsListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workflow_assignments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowAssignment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_workflow_assignments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workflow_assignments_proto_goTypes,
		DependencyIndexes: file_workflow_assignments_proto_depIdxs,
		MessageInfos:      file_workflow_assignments_proto_msgTypes,
	}.Build()
	File_workflow_assignments_proto = out.File
	file_workflow_assignments_proto_rawDesc = nil
	file_workflow_assignments_proto_goTypes = nil
	file_workflow_assignments_proto_depIdxs = nil
}
