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
// source: workflow_roles.proto

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

// WorkflowRolesCreateRequest specifies the workflowID and roleID of a new
// workflow role to be created.
type WorkflowRolesCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *CreateRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Parameters to define the new WorkflowRole.
	WorkflowRole *WorkflowRole `protobuf:"bytes,2,opt,name=workflow_role,json=workflowRole,proto3" json:"workflow_role,omitempty"`
}

func (x *WorkflowRolesCreateRequest) Reset() {
	*x = WorkflowRolesCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_roles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowRolesCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRolesCreateRequest) ProtoMessage() {}

func (x *WorkflowRolesCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_roles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRolesCreateRequest.ProtoReflect.Descriptor instead.
func (*WorkflowRolesCreateRequest) Descriptor() ([]byte, []int) {
	return file_workflow_roles_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowRolesCreateRequest) GetMeta() *CreateRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkflowRolesCreateRequest) GetWorkflowRole() *WorkflowRole {
	if x != nil {
		return x.WorkflowRole
	}
	return nil
}

// WorkflowRolesCreateResponse reports how the WorkflowRole was created in the system.
type WorkflowRolesCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *CreateResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The created workflow role.
	WorkflowRole *WorkflowRole `protobuf:"bytes,2,opt,name=workflow_role,json=workflowRole,proto3" json:"workflow_role,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *WorkflowRolesCreateResponse) Reset() {
	*x = WorkflowRolesCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_roles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowRolesCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRolesCreateResponse) ProtoMessage() {}

func (x *WorkflowRolesCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_roles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRolesCreateResponse.ProtoReflect.Descriptor instead.
func (*WorkflowRolesCreateResponse) Descriptor() ([]byte, []int) {
	return file_workflow_roles_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowRolesCreateResponse) GetMeta() *CreateResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkflowRolesCreateResponse) GetWorkflowRole() *WorkflowRole {
	if x != nil {
		return x.WorkflowRole
	}
	return nil
}

func (x *WorkflowRolesCreateResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// WorkflowRoleGetRequest specifies which WorkflowRole to retrieve.
type WorkflowRoleGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the WorkflowRole to retrieve.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WorkflowRoleGetRequest) Reset() {
	*x = WorkflowRoleGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_roles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowRoleGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRoleGetRequest) ProtoMessage() {}

func (x *WorkflowRoleGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_roles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRoleGetRequest.ProtoReflect.Descriptor instead.
func (*WorkflowRoleGetRequest) Descriptor() ([]byte, []int) {
	return file_workflow_roles_proto_rawDescGZIP(), []int{2}
}

func (x *WorkflowRoleGetRequest) GetMeta() *GetRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkflowRoleGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// WorkflowRoleGetResponse returns a requested WorkflowRole.
type WorkflowRoleGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested WorkflowRole.
	WorkflowRole *WorkflowRole `protobuf:"bytes,2,opt,name=workflow_role,json=workflowRole,proto3" json:"workflow_role,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *WorkflowRoleGetResponse) Reset() {
	*x = WorkflowRoleGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_roles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowRoleGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRoleGetResponse) ProtoMessage() {}

func (x *WorkflowRoleGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_roles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRoleGetResponse.ProtoReflect.Descriptor instead.
func (*WorkflowRoleGetResponse) Descriptor() ([]byte, []int) {
	return file_workflow_roles_proto_rawDescGZIP(), []int{3}
}

func (x *WorkflowRoleGetResponse) GetMeta() *GetResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkflowRoleGetResponse) GetWorkflowRole() *WorkflowRole {
	if x != nil {
		return x.WorkflowRole
	}
	return nil
}

func (x *WorkflowRoleGetResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// WorkflowRolesDeleteRequest specifies the ID of a WorkflowRole to be deleted.
type WorkflowRolesDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *DeleteRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the WorkflowRole to delete.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WorkflowRolesDeleteRequest) Reset() {
	*x = WorkflowRolesDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_roles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowRolesDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRolesDeleteRequest) ProtoMessage() {}

func (x *WorkflowRolesDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_roles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRolesDeleteRequest.ProtoReflect.Descriptor instead.
func (*WorkflowRolesDeleteRequest) Descriptor() ([]byte, []int) {
	return file_workflow_roles_proto_rawDescGZIP(), []int{4}
}

func (x *WorkflowRolesDeleteRequest) GetMeta() *DeleteRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkflowRolesDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// WorkflowRolesDeleteResponse reports how the WorkflowRole was deleted in the system.
type WorkflowRolesDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *DeleteResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,2,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *WorkflowRolesDeleteResponse) Reset() {
	*x = WorkflowRolesDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_roles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowRolesDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRolesDeleteResponse) ProtoMessage() {}

func (x *WorkflowRolesDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_roles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRolesDeleteResponse.ProtoReflect.Descriptor instead.
func (*WorkflowRolesDeleteResponse) Descriptor() ([]byte, []int) {
	return file_workflow_roles_proto_rawDescGZIP(), []int{5}
}

func (x *WorkflowRolesDeleteResponse) GetMeta() *DeleteResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkflowRolesDeleteResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// WorkflowRolesListRequest specifies criteria for retrieving a list of
// WorkflowRole records
type WorkflowRolesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *WorkflowRolesListRequest) Reset() {
	*x = WorkflowRolesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_roles_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowRolesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRolesListRequest) ProtoMessage() {}

func (x *WorkflowRolesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_roles_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRolesListRequest.ProtoReflect.Descriptor instead.
func (*WorkflowRolesListRequest) Descriptor() ([]byte, []int) {
	return file_workflow_roles_proto_rawDescGZIP(), []int{6}
}

func (x *WorkflowRolesListRequest) GetMeta() *ListRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkflowRolesListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// WorkflowRolesListResponse returns a list of WorkflowRole records that meet
// the criteria of a WorkflowRolesListRequest.
type WorkflowRolesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The matching workflow role.
	WorkflowRole []*WorkflowRole `protobuf:"bytes,2,rep,name=workflow_role,json=workflowRole,proto3" json:"workflow_role,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *WorkflowRolesListResponse) Reset() {
	*x = WorkflowRolesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_roles_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowRolesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRolesListResponse) ProtoMessage() {}

func (x *WorkflowRolesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_roles_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRolesListResponse.ProtoReflect.Descriptor instead.
func (*WorkflowRolesListResponse) Descriptor() ([]byte, []int) {
	return file_workflow_roles_proto_rawDescGZIP(), []int{7}
}

func (x *WorkflowRolesListResponse) GetMeta() *ListResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkflowRolesListResponse) GetWorkflowRole() []*WorkflowRole {
	if x != nil {
		return x.WorkflowRole
	}
	return nil
}

func (x *WorkflowRolesListResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// WorkflowRole links a role to a workflow. The linked roles indicate which roles a user must be a part of
// to request access to a resource via the workflow.
type WorkflowRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the WorkflowRole.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The workflow id.
	WorkflowId string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	// The role id.
	RoleId string `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *WorkflowRole) Reset() {
	*x = WorkflowRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_roles_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRole) ProtoMessage() {}

func (x *WorkflowRole) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_roles_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRole.ProtoReflect.Descriptor instead.
func (*WorkflowRole) Descriptor() ([]byte, []int) {
	return file_workflow_roles_proto_rawDescGZIP(), []int{8}
}

func (x *WorkflowRole) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkflowRole) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *WorkflowRole) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

var File_workflow_roles_proto protoreflect.FileDescriptor

var file_workflow_roles_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x1a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x0a, 0xf2, 0xf8,
	0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x3a, 0x0a, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3,
	0x07, 0x01, 0x22, 0x80, 0x02, 0x0a, 0x1b, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x41, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x0a, 0xf2, 0xf8, 0xb3,
	0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x62, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x2c, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07,
	0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07,
	0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x09,
	0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x0a, 0xfa, 0xf8, 0xb3, 0x07, 0x05,
	0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x22, 0x60, 0x0a, 0x16, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3,
	0xb3, 0x07, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x85, 0x02, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07,
	0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x0d,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x62, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0xf2, 0xf8, 0xb3, 0x07,
	0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01,
	0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e,
	0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x3a, 0x0a, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x22,
	0x73, 0x0a, 0x1a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x0a, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8,
	0xf3, 0xb3, 0x07, 0x01, 0x22, 0xbd, 0x01, 0x0a, 0x1b, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x62, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x2c, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06,
	0xb2, 0xf4, 0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d,
	0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x0a, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8,
	0xf3, 0xb3, 0x07, 0x01, 0x22, 0x77, 0x0a, 0x18, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2,
	0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x3a, 0x0a, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x22, 0xfc, 0x01,
	0x0a, 0x19, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x0d, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f,
	0x6c, 0x65, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb8, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0c,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x62, 0x0a, 0x0a,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3,
	0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8,
	0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x3a, 0x0a, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x22, 0xf6, 0x01, 0x0a,
	0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05,
	0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f,
	0xf2, 0xf8, 0xb3, 0x07, 0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0xf8,
	0xb3, 0x07, 0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x3a, 0x6e, 0xfa, 0xf8, 0xb3, 0x07, 0x69, 0xa8, 0xf3, 0xb3, 0x07,
	0x01, 0xc2, 0xf3, 0xb3, 0x07, 0x59, 0xa2, 0xf3, 0xb3, 0x07, 0x26, 0x74, 0x66, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x74, 0x78,
	0x74, 0xaa, 0xf3, 0xb3, 0x07, 0x29, 0x74, 0x66, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x74, 0x78, 0x74, 0xd2,
	0xf3, 0xb3, 0x07, 0x01, 0x2a, 0x32, 0x8d, 0x04, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x75, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2a, 0x82, 0xf9, 0xb3, 0x07, 0x09, 0xa2, 0xf3, 0xb3, 0x07, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x17, 0xaa, 0xf3, 0xb3, 0x07, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x6d,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x6f, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d,
	0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07, 0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3,
	0x07, 0x1b, 0xaa, 0xf3, 0xb3, 0x07, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2d, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x77, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xf9, 0xb3, 0x07, 0x0b, 0xa2,
	0xf3, 0xb3, 0x07, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x82, 0xf9, 0xb3, 0x07, 0x17, 0xaa,
	0xf3, 0xb3, 0x07, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2d, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x6e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xf9, 0xb3,
	0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07, 0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x17, 0xaa,
	0xf3, 0xb3, 0x07, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2d, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x1a, 0x2d, 0xca, 0xf9, 0xb3, 0x07, 0x11, 0xc2, 0xf9, 0xb3,
	0x07, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0xca, 0xf9,
	0xb3, 0x07, 0x08, 0xd2, 0xf9, 0xb3, 0x07, 0x03, 0x77, 0x72, 0x2d, 0xca, 0xf9, 0xb3, 0x07, 0x05,
	0xe8, 0xf9, 0xb3, 0x07, 0x01, 0x42, 0x69, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72,
	0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x69,
	0x6e, 0x67, 0x42, 0x15, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x50, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f, 0x73,
	0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76,
	0x33, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflow_roles_proto_rawDescOnce sync.Once
	file_workflow_roles_proto_rawDescData = file_workflow_roles_proto_rawDesc
)

func file_workflow_roles_proto_rawDescGZIP() []byte {
	file_workflow_roles_proto_rawDescOnce.Do(func() {
		file_workflow_roles_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_roles_proto_rawDescData)
	})
	return file_workflow_roles_proto_rawDescData
}

var file_workflow_roles_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_workflow_roles_proto_goTypes = []interface{}{
	(*WorkflowRolesCreateRequest)(nil),  // 0: v1.WorkflowRolesCreateRequest
	(*WorkflowRolesCreateResponse)(nil), // 1: v1.WorkflowRolesCreateResponse
	(*WorkflowRoleGetRequest)(nil),      // 2: v1.WorkflowRoleGetRequest
	(*WorkflowRoleGetResponse)(nil),     // 3: v1.WorkflowRoleGetResponse
	(*WorkflowRolesDeleteRequest)(nil),  // 4: v1.WorkflowRolesDeleteRequest
	(*WorkflowRolesDeleteResponse)(nil), // 5: v1.WorkflowRolesDeleteResponse
	(*WorkflowRolesListRequest)(nil),    // 6: v1.WorkflowRolesListRequest
	(*WorkflowRolesListResponse)(nil),   // 7: v1.WorkflowRolesListResponse
	(*WorkflowRole)(nil),                // 8: v1.WorkflowRole
	(*CreateRequestMetadata)(nil),       // 9: v1.CreateRequestMetadata
	(*CreateResponseMetadata)(nil),      // 10: v1.CreateResponseMetadata
	(*RateLimitMetadata)(nil),           // 11: v1.RateLimitMetadata
	(*GetRequestMetadata)(nil),          // 12: v1.GetRequestMetadata
	(*GetResponseMetadata)(nil),         // 13: v1.GetResponseMetadata
	(*DeleteRequestMetadata)(nil),       // 14: v1.DeleteRequestMetadata
	(*DeleteResponseMetadata)(nil),      // 15: v1.DeleteResponseMetadata
	(*ListRequestMetadata)(nil),         // 16: v1.ListRequestMetadata
	(*ListResponseMetadata)(nil),        // 17: v1.ListResponseMetadata
}
var file_workflow_roles_proto_depIdxs = []int32{
	9,  // 0: v1.WorkflowRolesCreateRequest.meta:type_name -> v1.CreateRequestMetadata
	8,  // 1: v1.WorkflowRolesCreateRequest.workflow_role:type_name -> v1.WorkflowRole
	10, // 2: v1.WorkflowRolesCreateResponse.meta:type_name -> v1.CreateResponseMetadata
	8,  // 3: v1.WorkflowRolesCreateResponse.workflow_role:type_name -> v1.WorkflowRole
	11, // 4: v1.WorkflowRolesCreateResponse.rate_limit:type_name -> v1.RateLimitMetadata
	12, // 5: v1.WorkflowRoleGetRequest.meta:type_name -> v1.GetRequestMetadata
	13, // 6: v1.WorkflowRoleGetResponse.meta:type_name -> v1.GetResponseMetadata
	8,  // 7: v1.WorkflowRoleGetResponse.workflow_role:type_name -> v1.WorkflowRole
	11, // 8: v1.WorkflowRoleGetResponse.rate_limit:type_name -> v1.RateLimitMetadata
	14, // 9: v1.WorkflowRolesDeleteRequest.meta:type_name -> v1.DeleteRequestMetadata
	15, // 10: v1.WorkflowRolesDeleteResponse.meta:type_name -> v1.DeleteResponseMetadata
	11, // 11: v1.WorkflowRolesDeleteResponse.rate_limit:type_name -> v1.RateLimitMetadata
	16, // 12: v1.WorkflowRolesListRequest.meta:type_name -> v1.ListRequestMetadata
	17, // 13: v1.WorkflowRolesListResponse.meta:type_name -> v1.ListResponseMetadata
	8,  // 14: v1.WorkflowRolesListResponse.workflow_role:type_name -> v1.WorkflowRole
	11, // 15: v1.WorkflowRolesListResponse.rate_limit:type_name -> v1.RateLimitMetadata
	0,  // 16: v1.WorkflowRoles.Create:input_type -> v1.WorkflowRolesCreateRequest
	2,  // 17: v1.WorkflowRoles.Get:input_type -> v1.WorkflowRoleGetRequest
	4,  // 18: v1.WorkflowRoles.Delete:input_type -> v1.WorkflowRolesDeleteRequest
	6,  // 19: v1.WorkflowRoles.List:input_type -> v1.WorkflowRolesListRequest
	1,  // 20: v1.WorkflowRoles.Create:output_type -> v1.WorkflowRolesCreateResponse
	3,  // 21: v1.WorkflowRoles.Get:output_type -> v1.WorkflowRoleGetResponse
	5,  // 22: v1.WorkflowRoles.Delete:output_type -> v1.WorkflowRolesDeleteResponse
	7,  // 23: v1.WorkflowRoles.List:output_type -> v1.WorkflowRolesListResponse
	20, // [20:24] is the sub-list for method output_type
	16, // [16:20] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_workflow_roles_proto_init() }
func file_workflow_roles_proto_init() {
	if File_workflow_roles_proto != nil {
		return
	}
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_workflow_roles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowRolesCreateRequest); i {
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
		file_workflow_roles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowRolesCreateResponse); i {
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
		file_workflow_roles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowRoleGetRequest); i {
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
		file_workflow_roles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowRoleGetResponse); i {
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
		file_workflow_roles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowRolesDeleteRequest); i {
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
		file_workflow_roles_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowRolesDeleteResponse); i {
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
		file_workflow_roles_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowRolesListRequest); i {
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
		file_workflow_roles_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowRolesListResponse); i {
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
		file_workflow_roles_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowRole); i {
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
			RawDescriptor: file_workflow_roles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workflow_roles_proto_goTypes,
		DependencyIndexes: file_workflow_roles_proto_depIdxs,
		MessageInfos:      file_workflow_roles_proto_msgTypes,
	}.Build()
	File_workflow_roles_proto = out.File
	file_workflow_roles_proto_rawDesc = nil
	file_workflow_roles_proto_goTypes = nil
	file_workflow_roles_proto_depIdxs = nil
}
