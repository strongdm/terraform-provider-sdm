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
// source: health_checks.proto

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

// HealthcheckListRequest specifies criteria for retrieving a list of Healthchecks.
type HealthcheckListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *HealthcheckListRequest) Reset() {
	*x = HealthcheckListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_checks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthcheckListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthcheckListRequest) ProtoMessage() {}

func (x *HealthcheckListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_checks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthcheckListRequest.ProtoReflect.Descriptor instead.
func (*HealthcheckListRequest) Descriptor() ([]byte, []int) {
	return file_health_checks_proto_rawDescGZIP(), []int{0}
}

func (x *HealthcheckListRequest) GetMeta() *ListRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *HealthcheckListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// HealthcheckListResponse returns a list of Healthchecks that meet the criteria of a
// HealthcheckListRequest.
type HealthcheckListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	Healthchecks []*Healthcheck `protobuf:"bytes,2,rep,name=healthchecks,proto3" json:"healthchecks,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *HealthcheckListResponse) Reset() {
	*x = HealthcheckListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_checks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthcheckListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthcheckListResponse) ProtoMessage() {}

func (x *HealthcheckListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_health_checks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthcheckListResponse.ProtoReflect.Descriptor instead.
func (*HealthcheckListResponse) Descriptor() ([]byte, []int) {
	return file_health_checks_proto_rawDescGZIP(), []int{1}
}

func (x *HealthcheckListResponse) GetMeta() *ListResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *HealthcheckListResponse) GetHealthchecks() []*Healthcheck {
	if x != nil {
		return x.Healthchecks
	}
	return nil
}

func (x *HealthcheckListResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// Healthcheck defines the status of the link between a node and a resource
type Healthcheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the healthcheck.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique identifier of the healthcheck resource.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The name of the resource.
	ResourceName string `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Unique identifier of the healthcheck node.
	NodeId string `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// The name of the node.
	NodeName string `protobuf:"bytes,5,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Whether the healthcheck succeeded.
	Healthy bool `protobuf:"varint,6,opt,name=healthy,proto3" json:"healthy,omitempty"`
	// The error if unhealthy
	ErrorMsg string `protobuf:"bytes,7,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	// The time at which the healthcheck state was recorded.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Healthcheck) Reset() {
	*x = Healthcheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_checks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Healthcheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Healthcheck) ProtoMessage() {}

func (x *Healthcheck) ProtoReflect() protoreflect.Message {
	mi := &file_health_checks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Healthcheck.ProtoReflect.Descriptor instead.
func (*Healthcheck) Descriptor() ([]byte, []int) {
	return file_health_checks_proto_rawDescGZIP(), []int{2}
}

func (x *Healthcheck) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Healthcheck) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Healthcheck) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *Healthcheck) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Healthcheck) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Healthcheck) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

func (x *Healthcheck) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *Healthcheck) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_health_checks_proto protoreflect.FileDescriptor

var file_health_checks_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x16, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2,
	0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x3a, 0x28, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8,
	0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0xa0, 0x02, 0x0a, 0x17,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x0a, 0xf2, 0xf8,
	0xb3, 0x07, 0x05, 0xb8, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x62, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x2c, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3,
	0x07, 0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3,
	0x07, 0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52,
	0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x32, 0xfa, 0xf8, 0xb3, 0x07,
	0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01,
	0x2a, 0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72,
	0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0xc1,
	0x03, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1f,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0xf8, 0xb3, 0x07,
	0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x30, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0xf8, 0xb3, 0x07, 0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01,
	0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x34, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0xf8, 0xb3, 0x07, 0x0a, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0xf8, 0xb3, 0x07, 0x0a, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x2c, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0xf8, 0xb3, 0x07, 0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01,
	0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x0f, 0xf2, 0xf8, 0xb3, 0x07, 0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x12, 0x27, 0x0a, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2,
	0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x49, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x0f, 0xf2, 0xf8, 0xb3, 0x07, 0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3,
	0xb3, 0x07, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x3a, 0x32,
	0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2,
	0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x32, 0xc1, 0x01, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x12, 0x68, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07, 0x03,
	0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x15, 0xaa, 0xf3, 0xb3, 0x07, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x1a, 0x47, 0xca,
	0xf9, 0xb3, 0x07, 0x10, 0xc2, 0xf9, 0xb3, 0x07, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0xca, 0xf9, 0xb3, 0x07, 0x05, 0xd8, 0xf9, 0xb3, 0x07, 0x01, 0xca, 0xf9,
	0xb3, 0x07, 0x06, 0xca, 0xf9, 0xb3, 0x07, 0x01, 0x2a, 0xca, 0xf9, 0xb3, 0x07, 0x18, 0xca, 0xf9,
	0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x90, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x6d,
	0x62, 0x69, 0x6e, 0x67, 0x42, 0x14, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x50, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x76, 0x33, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0xc2, 0x92, 0xb4, 0x07, 0x06, 0xa2, 0x8c, 0xb4, 0x07, 0x01, 0x2a, 0xc2, 0x92, 0xb4, 0x07,
	0x18, 0xa2, 0x8c, 0xb4, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_health_checks_proto_rawDescOnce sync.Once
	file_health_checks_proto_rawDescData = file_health_checks_proto_rawDesc
)

func file_health_checks_proto_rawDescGZIP() []byte {
	file_health_checks_proto_rawDescOnce.Do(func() {
		file_health_checks_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_checks_proto_rawDescData)
	})
	return file_health_checks_proto_rawDescData
}

var file_health_checks_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_health_checks_proto_goTypes = []interface{}{
	(*HealthcheckListRequest)(nil),  // 0: v1.HealthcheckListRequest
	(*HealthcheckListResponse)(nil), // 1: v1.HealthcheckListResponse
	(*Healthcheck)(nil),             // 2: v1.Healthcheck
	(*ListRequestMetadata)(nil),     // 3: v1.ListRequestMetadata
	(*ListResponseMetadata)(nil),    // 4: v1.ListResponseMetadata
	(*RateLimitMetadata)(nil),       // 5: v1.RateLimitMetadata
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_health_checks_proto_depIdxs = []int32{
	3, // 0: v1.HealthcheckListRequest.meta:type_name -> v1.ListRequestMetadata
	4, // 1: v1.HealthcheckListResponse.meta:type_name -> v1.ListResponseMetadata
	2, // 2: v1.HealthcheckListResponse.healthchecks:type_name -> v1.Healthcheck
	5, // 3: v1.HealthcheckListResponse.rate_limit:type_name -> v1.RateLimitMetadata
	6, // 4: v1.Healthcheck.timestamp:type_name -> google.protobuf.Timestamp
	0, // 5: v1.HealthChecks.List:input_type -> v1.HealthcheckListRequest
	1, // 6: v1.HealthChecks.List:output_type -> v1.HealthcheckListResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_health_checks_proto_init() }
func file_health_checks_proto_init() {
	if File_health_checks_proto != nil {
		return
	}
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_health_checks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthcheckListRequest); i {
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
		file_health_checks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthcheckListResponse); i {
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
		file_health_checks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Healthcheck); i {
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
			RawDescriptor: file_health_checks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_checks_proto_goTypes,
		DependencyIndexes: file_health_checks_proto_depIdxs,
		MessageInfos:      file_health_checks_proto_msgTypes,
	}.Build()
	File_health_checks_proto = out.File
	file_health_checks_proto_rawDesc = nil
	file_health_checks_proto_goTypes = nil
	file_health_checks_proto_depIdxs = nil
}