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
// source: role_resources_history.proto

package v1

import (
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

// RoleResourceHistoryListRequest specifies criteria for retrieving a list of
// RoleResourceHistory records.
type RoleResourceHistoryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *RoleResourceHistoryListRequest) Reset() {
	*x = RoleResourceHistoryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_resources_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleResourceHistoryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleResourceHistoryListRequest) ProtoMessage() {}

func (x *RoleResourceHistoryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_resources_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleResourceHistoryListRequest.ProtoReflect.Descriptor instead.
func (*RoleResourceHistoryListRequest) Descriptor() ([]byte, []int) {
	return file_role_resources_history_proto_rawDescGZIP(), []int{0}
}

func (x *RoleResourceHistoryListRequest) GetMeta() *ListRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RoleResourceHistoryListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// RoleResourceHistoryListResponse returns a list of RoleResourceHistory records that meet
// the criteria of an RoleResourceHistoryListRequest.
type RoleResourceHistoryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	History []*RoleResourceHistory `protobuf:"bytes,2,rep,name=history,proto3" json:"history,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *RoleResourceHistoryListResponse) Reset() {
	*x = RoleResourceHistoryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_resources_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleResourceHistoryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleResourceHistoryListResponse) ProtoMessage() {}

func (x *RoleResourceHistoryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_resources_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleResourceHistoryListResponse.ProtoReflect.Descriptor instead.
func (*RoleResourceHistoryListResponse) Descriptor() ([]byte, []int) {
	return file_role_resources_history_proto_rawDescGZIP(), []int{1}
}

func (x *RoleResourceHistoryListResponse) GetMeta() *ListResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RoleResourceHistoryListResponse) GetHistory() []*RoleResourceHistory {
	if x != nil {
		return x.History
	}
	return nil
}

func (x *RoleResourceHistoryListResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// RoleResourceHistory records the state of a RoleResource at a given point in time,
// where every change (create or delete) to a RoleResource produces an
// RoleResourceHistory record.
type RoleResourceHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the Activity that produced this change to the RoleResource.
	// May be empty for some system-initiated updates.
	ActivityId string `protobuf:"bytes,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	// The time at which the RoleResource state was recorded.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The complete RoleResource state at this time.
	RoleResource *RoleResource `protobuf:"bytes,3,opt,name=role_resource,json=roleResource,proto3" json:"role_resource,omitempty"`
	// If this RoleResource was deleted, the time it was deleted.
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *RoleResourceHistory) Reset() {
	*x = RoleResourceHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_resources_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleResourceHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleResourceHistory) ProtoMessage() {}

func (x *RoleResourceHistory) ProtoReflect() protoreflect.Message {
	mi := &file_role_resources_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleResourceHistory.ProtoReflect.Descriptor instead.
func (*RoleResourceHistory) Descriptor() ([]byte, []int) {
	return file_role_resources_history_proto_rawDescGZIP(), []int{2}
}

func (x *RoleResourceHistory) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *RoleResourceHistory) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *RoleResourceHistory) GetRoleResource() *RoleResource {
	if x != nil {
		return x.RoleResource
	}
	return nil
}

func (x *RoleResourceHistory) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_role_resources_history_proto protoreflect.FileDescriptor

var file_role_resources_history_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x1e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x28, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2,
	0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x22, 0x9c, 0x02, 0x0a, 0x1f, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x0a,
	0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb8, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x62, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c,
	0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2,
	0xf4, 0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21,
	0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x28, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3,
	0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74,
	0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x22, 0xc6, 0x02, 0x0a, 0x13, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x41, 0x0a, 0x0d,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x32, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3,
	0x07, 0x01, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3,
	0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72,
	0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x32, 0xec, 0x01, 0x0a, 0x14, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x82, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07,
	0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x1f, 0xaa, 0xf3, 0xb3, 0x07, 0x1a, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x4f, 0xca, 0xf9, 0xb3, 0x07, 0x18, 0xc2,
	0xf9, 0xb3, 0x07, 0x13, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0xca, 0xf9, 0xb3, 0x07, 0x05, 0xd8, 0xf9, 0xb3, 0x07,
	0x01, 0xca, 0xf9, 0xb3, 0x07, 0x06, 0xca, 0xf9, 0xb3, 0x07, 0x01, 0x2a, 0xca, 0xf9, 0xb3, 0x07,
	0x18, 0xca, 0xf9, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x98, 0x01, 0x0a, 0x19, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x42, 0x1c, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x6c, 0x75,
	0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e,
	0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xc2, 0x92, 0xb4, 0x07,
	0x06, 0xa2, 0x8c, 0xb4, 0x07, 0x01, 0x2a, 0xc2, 0x92, 0xb4, 0x07, 0x18, 0xa2, 0x8c, 0xb4, 0x07,
	0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_resources_history_proto_rawDescOnce sync.Once
	file_role_resources_history_proto_rawDescData = file_role_resources_history_proto_rawDesc
)

func file_role_resources_history_proto_rawDescGZIP() []byte {
	file_role_resources_history_proto_rawDescOnce.Do(func() {
		file_role_resources_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_resources_history_proto_rawDescData)
	})
	return file_role_resources_history_proto_rawDescData
}

var file_role_resources_history_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_role_resources_history_proto_goTypes = []interface{}{
	(*RoleResourceHistoryListRequest)(nil),  // 0: v1.RoleResourceHistoryListRequest
	(*RoleResourceHistoryListResponse)(nil), // 1: v1.RoleResourceHistoryListResponse
	(*RoleResourceHistory)(nil),             // 2: v1.RoleResourceHistory
	(*ListRequestMetadata)(nil),             // 3: v1.ListRequestMetadata
	(*ListResponseMetadata)(nil),            // 4: v1.ListResponseMetadata
	(*RateLimitMetadata)(nil),               // 5: v1.RateLimitMetadata
	(*timestamppb.Timestamp)(nil),           // 6: google.protobuf.Timestamp
	(*RoleResource)(nil),                    // 7: v1.RoleResource
}
var file_role_resources_history_proto_depIdxs = []int32{
	3, // 0: v1.RoleResourceHistoryListRequest.meta:type_name -> v1.ListRequestMetadata
	4, // 1: v1.RoleResourceHistoryListResponse.meta:type_name -> v1.ListResponseMetadata
	2, // 2: v1.RoleResourceHistoryListResponse.history:type_name -> v1.RoleResourceHistory
	5, // 3: v1.RoleResourceHistoryListResponse.rate_limit:type_name -> v1.RateLimitMetadata
	6, // 4: v1.RoleResourceHistory.timestamp:type_name -> google.protobuf.Timestamp
	7, // 5: v1.RoleResourceHistory.role_resource:type_name -> v1.RoleResource
	6, // 6: v1.RoleResourceHistory.deleted_at:type_name -> google.protobuf.Timestamp
	0, // 7: v1.RoleResourcesHistory.List:input_type -> v1.RoleResourceHistoryListRequest
	1, // 8: v1.RoleResourcesHistory.List:output_type -> v1.RoleResourceHistoryListResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_role_resources_history_proto_init() }
func file_role_resources_history_proto_init() {
	if File_role_resources_history_proto != nil {
		return
	}
	file_role_resources_proto_init()
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_role_resources_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleResourceHistoryListRequest); i {
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
		file_role_resources_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleResourceHistoryListResponse); i {
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
		file_role_resources_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleResourceHistory); i {
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
			RawDescriptor: file_role_resources_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_resources_history_proto_goTypes,
		DependencyIndexes: file_role_resources_history_proto_depIdxs,
		MessageInfos:      file_role_resources_history_proto_msgTypes,
	}.Build()
	File_role_resources_history_proto = out.File
	file_role_resources_history_proto_rawDesc = nil
	file_role_resources_history_proto_goTypes = nil
	file_role_resources_history_proto_depIdxs = nil
}
