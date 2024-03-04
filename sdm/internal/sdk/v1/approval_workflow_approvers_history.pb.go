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
// source: approval_workflow_approvers_history.proto

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

// ApprovalWorkflowApproverHistoryListRequest specifies criteria for retrieving a list of
// ApprovalWorkflowApproverHistory records.
type ApprovalWorkflowApproverHistoryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ApprovalWorkflowApproverHistoryListRequest) Reset() {
	*x = ApprovalWorkflowApproverHistoryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_approval_workflow_approvers_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalWorkflowApproverHistoryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalWorkflowApproverHistoryListRequest) ProtoMessage() {}

func (x *ApprovalWorkflowApproverHistoryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_approval_workflow_approvers_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalWorkflowApproverHistoryListRequest.ProtoReflect.Descriptor instead.
func (*ApprovalWorkflowApproverHistoryListRequest) Descriptor() ([]byte, []int) {
	return file_approval_workflow_approvers_history_proto_rawDescGZIP(), []int{0}
}

func (x *ApprovalWorkflowApproverHistoryListRequest) GetMeta() *ListRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ApprovalWorkflowApproverHistoryListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// ApprovalWorkflowApproverHistoryListResponse returns a list of ApprovalWorkflowApproverHistory records that meet
// the criteria of an ApprovalWorkflowApproverHistoryListRequest.
type ApprovalWorkflowApproverHistoryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	History []*ApprovalWorkflowApproverHistory `protobuf:"bytes,2,rep,name=history,proto3" json:"history,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *ApprovalWorkflowApproverHistoryListResponse) Reset() {
	*x = ApprovalWorkflowApproverHistoryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_approval_workflow_approvers_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalWorkflowApproverHistoryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalWorkflowApproverHistoryListResponse) ProtoMessage() {}

func (x *ApprovalWorkflowApproverHistoryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_approval_workflow_approvers_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalWorkflowApproverHistoryListResponse.ProtoReflect.Descriptor instead.
func (*ApprovalWorkflowApproverHistoryListResponse) Descriptor() ([]byte, []int) {
	return file_approval_workflow_approvers_history_proto_rawDescGZIP(), []int{1}
}

func (x *ApprovalWorkflowApproverHistoryListResponse) GetMeta() *ListResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ApprovalWorkflowApproverHistoryListResponse) GetHistory() []*ApprovalWorkflowApproverHistory {
	if x != nil {
		return x.History
	}
	return nil
}

func (x *ApprovalWorkflowApproverHistoryListResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// ApprovalWorkflowApproverHistory records the state of an ApprovalWorkflowApprover at a given point in time,
// where every change (create or delete) to an ApprovalWorkflowApprover produces an
// ApprovalWorkflowApproverHistory record.
type ApprovalWorkflowApproverHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the Activity that produced this change to the ApprovalWorkflowApprover.
	// May be empty for some system-initiated updates.
	ActivityId string `protobuf:"bytes,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	// The time at which the ApprovalWorkflowApprover state was recorded.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The complete ApprovalWorkflowApprover state at this time.
	ApprovalWorkflowApprover *ApprovalWorkflowApprover `protobuf:"bytes,3,opt,name=approval_workflow_approver,json=approvalWorkflowApprover,proto3" json:"approval_workflow_approver,omitempty"`
	// If this ApprovalWorkflowApprover was deleted, the time it was deleted.
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *ApprovalWorkflowApproverHistory) Reset() {
	*x = ApprovalWorkflowApproverHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_approval_workflow_approvers_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalWorkflowApproverHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalWorkflowApproverHistory) ProtoMessage() {}

func (x *ApprovalWorkflowApproverHistory) ProtoReflect() protoreflect.Message {
	mi := &file_approval_workflow_approvers_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalWorkflowApproverHistory.ProtoReflect.Descriptor instead.
func (*ApprovalWorkflowApproverHistory) Descriptor() ([]byte, []int) {
	return file_approval_workflow_approvers_history_proto_rawDescGZIP(), []int{2}
}

func (x *ApprovalWorkflowApproverHistory) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *ApprovalWorkflowApproverHistory) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ApprovalWorkflowApproverHistory) GetApprovalWorkflowApprover() *ApprovalWorkflowApprover {
	if x != nil {
		return x.ApprovalWorkflowApprover
	}
	return nil
}

func (x *ApprovalWorkflowApproverHistory) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_approval_workflow_approvers_history_proto protoreflect.FileDescriptor

var file_approval_workflow_approvers_history_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x5f, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7,
	0x01, 0x0a, 0x2a, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07,
	0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x28,
	0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3, 0x07, 0x18,
	0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0xb4, 0x02, 0x0a, 0x2b, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x0a, 0xf2, 0xf8,
	0xb3, 0x07, 0x05, 0xb8, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x62, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0xf2, 0xf8,
	0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4, 0xb3,
	0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21, 0x6a, 0x73,
	0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x28, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07,
	0x01, 0x2a, 0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22,
	0xf7, 0x02, 0x0a, 0x1f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x44, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x66, 0x0a, 0x1a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0x52, 0x18, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x45,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a,
	0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x32, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07,
	0x01, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3, 0x07,
	0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x32, 0xa9, 0x02, 0x0a, 0x20, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0xa7,
	0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2,
	0xf3, 0xb3, 0x07, 0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x2c, 0xaa, 0xf3, 0xb3, 0x07,
	0x27, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2d, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73,
	0x2d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x5b, 0xca, 0xf9, 0xb3, 0x07, 0x24, 0xc2,
	0xf9, 0xb3, 0x07, 0x1f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0xca, 0xf9, 0xb3, 0x07, 0x05, 0xd8, 0xf9, 0xb3, 0x07, 0x01, 0xca, 0xf9, 0xb3,
	0x07, 0x06, 0xca, 0xf9, 0xb3, 0x07, 0x01, 0x2a, 0xca, 0xf9, 0xb3, 0x07, 0x18, 0xca, 0xf9, 0xb3,
	0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0xa4, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74,
	0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62,
	0x69, 0x6e, 0x67, 0x42, 0x28, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67,
	0x64, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d,
	0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0xc2, 0x92, 0xb4, 0x07, 0x06, 0xa2, 0x8c, 0xb4, 0x07, 0x01, 0x2a, 0xc2,
	0x92, 0xb4, 0x07, 0x18, 0xa2, 0x8c, 0xb4, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_approval_workflow_approvers_history_proto_rawDescOnce sync.Once
	file_approval_workflow_approvers_history_proto_rawDescData = file_approval_workflow_approvers_history_proto_rawDesc
)

func file_approval_workflow_approvers_history_proto_rawDescGZIP() []byte {
	file_approval_workflow_approvers_history_proto_rawDescOnce.Do(func() {
		file_approval_workflow_approvers_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_approval_workflow_approvers_history_proto_rawDescData)
	})
	return file_approval_workflow_approvers_history_proto_rawDescData
}

var file_approval_workflow_approvers_history_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_approval_workflow_approvers_history_proto_goTypes = []interface{}{
	(*ApprovalWorkflowApproverHistoryListRequest)(nil),  // 0: v1.ApprovalWorkflowApproverHistoryListRequest
	(*ApprovalWorkflowApproverHistoryListResponse)(nil), // 1: v1.ApprovalWorkflowApproverHistoryListResponse
	(*ApprovalWorkflowApproverHistory)(nil),             // 2: v1.ApprovalWorkflowApproverHistory
	(*ListRequestMetadata)(nil),                         // 3: v1.ListRequestMetadata
	(*ListResponseMetadata)(nil),                        // 4: v1.ListResponseMetadata
	(*RateLimitMetadata)(nil),                           // 5: v1.RateLimitMetadata
	(*timestamppb.Timestamp)(nil),                       // 6: google.protobuf.Timestamp
	(*ApprovalWorkflowApprover)(nil),                    // 7: v1.ApprovalWorkflowApprover
}
var file_approval_workflow_approvers_history_proto_depIdxs = []int32{
	3, // 0: v1.ApprovalWorkflowApproverHistoryListRequest.meta:type_name -> v1.ListRequestMetadata
	4, // 1: v1.ApprovalWorkflowApproverHistoryListResponse.meta:type_name -> v1.ListResponseMetadata
	2, // 2: v1.ApprovalWorkflowApproverHistoryListResponse.history:type_name -> v1.ApprovalWorkflowApproverHistory
	5, // 3: v1.ApprovalWorkflowApproverHistoryListResponse.rate_limit:type_name -> v1.RateLimitMetadata
	6, // 4: v1.ApprovalWorkflowApproverHistory.timestamp:type_name -> google.protobuf.Timestamp
	7, // 5: v1.ApprovalWorkflowApproverHistory.approval_workflow_approver:type_name -> v1.ApprovalWorkflowApprover
	6, // 6: v1.ApprovalWorkflowApproverHistory.deleted_at:type_name -> google.protobuf.Timestamp
	0, // 7: v1.ApprovalWorkflowApproversHistory.List:input_type -> v1.ApprovalWorkflowApproverHistoryListRequest
	1, // 8: v1.ApprovalWorkflowApproversHistory.List:output_type -> v1.ApprovalWorkflowApproverHistoryListResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_approval_workflow_approvers_history_proto_init() }
func file_approval_workflow_approvers_history_proto_init() {
	if File_approval_workflow_approvers_history_proto != nil {
		return
	}
	file_approval_workflow_approvers_proto_init()
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_approval_workflow_approvers_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalWorkflowApproverHistoryListRequest); i {
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
		file_approval_workflow_approvers_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalWorkflowApproverHistoryListResponse); i {
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
		file_approval_workflow_approvers_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalWorkflowApproverHistory); i {
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
			RawDescriptor: file_approval_workflow_approvers_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_approval_workflow_approvers_history_proto_goTypes,
		DependencyIndexes: file_approval_workflow_approvers_history_proto_depIdxs,
		MessageInfos:      file_approval_workflow_approvers_history_proto_msgTypes,
	}.Build()
	File_approval_workflow_approvers_history_proto = out.File
	file_approval_workflow_approvers_history_proto_rawDesc = nil
	file_approval_workflow_approvers_history_proto_goTypes = nil
	file_approval_workflow_approvers_history_proto_depIdxs = nil
}