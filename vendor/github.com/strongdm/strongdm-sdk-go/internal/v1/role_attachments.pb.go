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
// source: role_attachments.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// RoleAttachmentCreateRequest specifies what kind of RoleAttachments should be registered in
// the organizations fleet.
type RoleAttachmentCreateRequest struct {
	// Reserved for future use.
	Meta *CreateRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Parameters to define the new RoleAttachment.
	RoleAttachment       *RoleAttachment `protobuf:"bytes,2,opt,name=role_attachment,json=roleAttachment,proto3" json:"role_attachment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RoleAttachmentCreateRequest) Reset()         { *m = RoleAttachmentCreateRequest{} }
func (m *RoleAttachmentCreateRequest) String() string { return proto.CompactTextString(m) }
func (*RoleAttachmentCreateRequest) ProtoMessage()    {}
func (*RoleAttachmentCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32d84988da10a62, []int{0}
}

func (m *RoleAttachmentCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAttachmentCreateRequest.Unmarshal(m, b)
}
func (m *RoleAttachmentCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAttachmentCreateRequest.Marshal(b, m, deterministic)
}
func (m *RoleAttachmentCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAttachmentCreateRequest.Merge(m, src)
}
func (m *RoleAttachmentCreateRequest) XXX_Size() int {
	return xxx_messageInfo_RoleAttachmentCreateRequest.Size(m)
}
func (m *RoleAttachmentCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAttachmentCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAttachmentCreateRequest proto.InternalMessageInfo

func (m *RoleAttachmentCreateRequest) GetMeta() *CreateRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleAttachmentCreateRequest) GetRoleAttachment() *RoleAttachment {
	if m != nil {
		return m.RoleAttachment
	}
	return nil
}

// RoleAttachmentCreateResponse reports how the RoleAttachments were created in the system.
type RoleAttachmentCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The created RoleAttachment.
	RoleAttachment *RoleAttachment `protobuf:"bytes,2,opt,name=role_attachment,json=roleAttachment,proto3" json:"role_attachment,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleAttachmentCreateResponse) Reset()         { *m = RoleAttachmentCreateResponse{} }
func (m *RoleAttachmentCreateResponse) String() string { return proto.CompactTextString(m) }
func (*RoleAttachmentCreateResponse) ProtoMessage()    {}
func (*RoleAttachmentCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32d84988da10a62, []int{1}
}

func (m *RoleAttachmentCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAttachmentCreateResponse.Unmarshal(m, b)
}
func (m *RoleAttachmentCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAttachmentCreateResponse.Marshal(b, m, deterministic)
}
func (m *RoleAttachmentCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAttachmentCreateResponse.Merge(m, src)
}
func (m *RoleAttachmentCreateResponse) XXX_Size() int {
	return xxx_messageInfo_RoleAttachmentCreateResponse.Size(m)
}
func (m *RoleAttachmentCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAttachmentCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAttachmentCreateResponse proto.InternalMessageInfo

func (m *RoleAttachmentCreateResponse) GetMeta() *CreateResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleAttachmentCreateResponse) GetRoleAttachment() *RoleAttachment {
	if m != nil {
		return m.RoleAttachment
	}
	return nil
}

func (m *RoleAttachmentCreateResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// RoleAttachmentGetRequest specifies which RoleAttachment to retrieve.
type RoleAttachmentGetRequest struct {
	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the RoleAttachment to retrieve.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleAttachmentGetRequest) Reset()         { *m = RoleAttachmentGetRequest{} }
func (m *RoleAttachmentGetRequest) String() string { return proto.CompactTextString(m) }
func (*RoleAttachmentGetRequest) ProtoMessage()    {}
func (*RoleAttachmentGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32d84988da10a62, []int{2}
}

func (m *RoleAttachmentGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAttachmentGetRequest.Unmarshal(m, b)
}
func (m *RoleAttachmentGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAttachmentGetRequest.Marshal(b, m, deterministic)
}
func (m *RoleAttachmentGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAttachmentGetRequest.Merge(m, src)
}
func (m *RoleAttachmentGetRequest) XXX_Size() int {
	return xxx_messageInfo_RoleAttachmentGetRequest.Size(m)
}
func (m *RoleAttachmentGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAttachmentGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAttachmentGetRequest proto.InternalMessageInfo

func (m *RoleAttachmentGetRequest) GetMeta() *GetRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleAttachmentGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// RoleAttachmentGetResponse returns a requested RoleAttachment.
type RoleAttachmentGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested RoleAttachment.
	RoleAttachment *RoleAttachment `protobuf:"bytes,2,opt,name=role_attachment,json=roleAttachment,proto3" json:"role_attachment,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleAttachmentGetResponse) Reset()         { *m = RoleAttachmentGetResponse{} }
func (m *RoleAttachmentGetResponse) String() string { return proto.CompactTextString(m) }
func (*RoleAttachmentGetResponse) ProtoMessage()    {}
func (*RoleAttachmentGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32d84988da10a62, []int{3}
}

func (m *RoleAttachmentGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAttachmentGetResponse.Unmarshal(m, b)
}
func (m *RoleAttachmentGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAttachmentGetResponse.Marshal(b, m, deterministic)
}
func (m *RoleAttachmentGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAttachmentGetResponse.Merge(m, src)
}
func (m *RoleAttachmentGetResponse) XXX_Size() int {
	return xxx_messageInfo_RoleAttachmentGetResponse.Size(m)
}
func (m *RoleAttachmentGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAttachmentGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAttachmentGetResponse proto.InternalMessageInfo

func (m *RoleAttachmentGetResponse) GetMeta() *GetResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleAttachmentGetResponse) GetRoleAttachment() *RoleAttachment {
	if m != nil {
		return m.RoleAttachment
	}
	return nil
}

func (m *RoleAttachmentGetResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// RoleAttachmentDeleteRequest identifies a RoleAttachment by ID to delete.
type RoleAttachmentDeleteRequest struct {
	// Reserved for future use.
	Meta *DeleteRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the RoleAttachment to delete.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleAttachmentDeleteRequest) Reset()         { *m = RoleAttachmentDeleteRequest{} }
func (m *RoleAttachmentDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*RoleAttachmentDeleteRequest) ProtoMessage()    {}
func (*RoleAttachmentDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32d84988da10a62, []int{4}
}

func (m *RoleAttachmentDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAttachmentDeleteRequest.Unmarshal(m, b)
}
func (m *RoleAttachmentDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAttachmentDeleteRequest.Marshal(b, m, deterministic)
}
func (m *RoleAttachmentDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAttachmentDeleteRequest.Merge(m, src)
}
func (m *RoleAttachmentDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_RoleAttachmentDeleteRequest.Size(m)
}
func (m *RoleAttachmentDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAttachmentDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAttachmentDeleteRequest proto.InternalMessageInfo

func (m *RoleAttachmentDeleteRequest) GetMeta() *DeleteRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleAttachmentDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// RoleAttachmentDeleteResponse returns information about a RoleAttachment that was deleted.
type RoleAttachmentDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,2,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleAttachmentDeleteResponse) Reset()         { *m = RoleAttachmentDeleteResponse{} }
func (m *RoleAttachmentDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*RoleAttachmentDeleteResponse) ProtoMessage()    {}
func (*RoleAttachmentDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32d84988da10a62, []int{5}
}

func (m *RoleAttachmentDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAttachmentDeleteResponse.Unmarshal(m, b)
}
func (m *RoleAttachmentDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAttachmentDeleteResponse.Marshal(b, m, deterministic)
}
func (m *RoleAttachmentDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAttachmentDeleteResponse.Merge(m, src)
}
func (m *RoleAttachmentDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_RoleAttachmentDeleteResponse.Size(m)
}
func (m *RoleAttachmentDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAttachmentDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAttachmentDeleteResponse proto.InternalMessageInfo

func (m *RoleAttachmentDeleteResponse) GetMeta() *DeleteResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleAttachmentDeleteResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// RoleAttachmentListRequest specifies criteria for retrieving a list of RoleAttachments.
type RoleAttachmentListRequest struct {
	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter               string   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleAttachmentListRequest) Reset()         { *m = RoleAttachmentListRequest{} }
func (m *RoleAttachmentListRequest) String() string { return proto.CompactTextString(m) }
func (*RoleAttachmentListRequest) ProtoMessage()    {}
func (*RoleAttachmentListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32d84988da10a62, []int{6}
}

func (m *RoleAttachmentListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAttachmentListRequest.Unmarshal(m, b)
}
func (m *RoleAttachmentListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAttachmentListRequest.Marshal(b, m, deterministic)
}
func (m *RoleAttachmentListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAttachmentListRequest.Merge(m, src)
}
func (m *RoleAttachmentListRequest) XXX_Size() int {
	return xxx_messageInfo_RoleAttachmentListRequest.Size(m)
}
func (m *RoleAttachmentListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAttachmentListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAttachmentListRequest proto.InternalMessageInfo

func (m *RoleAttachmentListRequest) GetMeta() *ListRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleAttachmentListRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// RoleAttachmentListResponse returns a list of RoleAttachments that meet the criteria of a
// RoleAttachmentListRequest.
type RoleAttachmentListResponse struct {
	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	RoleAttachments []*RoleAttachment `protobuf:"bytes,2,rep,name=role_attachments,json=roleAttachments,proto3" json:"role_attachments,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleAttachmentListResponse) Reset()         { *m = RoleAttachmentListResponse{} }
func (m *RoleAttachmentListResponse) String() string { return proto.CompactTextString(m) }
func (*RoleAttachmentListResponse) ProtoMessage()    {}
func (*RoleAttachmentListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32d84988da10a62, []int{7}
}

func (m *RoleAttachmentListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAttachmentListResponse.Unmarshal(m, b)
}
func (m *RoleAttachmentListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAttachmentListResponse.Marshal(b, m, deterministic)
}
func (m *RoleAttachmentListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAttachmentListResponse.Merge(m, src)
}
func (m *RoleAttachmentListResponse) XXX_Size() int {
	return xxx_messageInfo_RoleAttachmentListResponse.Size(m)
}
func (m *RoleAttachmentListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAttachmentListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAttachmentListResponse proto.InternalMessageInfo

func (m *RoleAttachmentListResponse) GetMeta() *ListResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleAttachmentListResponse) GetRoleAttachments() []*RoleAttachment {
	if m != nil {
		return m.RoleAttachments
	}
	return nil
}

func (m *RoleAttachmentListResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// A RoleAttachment assigns a role to a composite role.
type RoleAttachment struct {
	// Unique identifier of the RoleAttachment.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The id of the composite role of this RoleAttachment.
	CompositeRoleId string `protobuf:"bytes,2,opt,name=composite_role_id,json=compositeRoleId,proto3" json:"composite_role_id,omitempty"`
	// The id of the attached role of this RoleAttachment.
	AttachedRoleId       string   `protobuf:"bytes,3,opt,name=attached_role_id,json=attachedRoleId,proto3" json:"attached_role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleAttachment) Reset()         { *m = RoleAttachment{} }
func (m *RoleAttachment) String() string { return proto.CompactTextString(m) }
func (*RoleAttachment) ProtoMessage()    {}
func (*RoleAttachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32d84988da10a62, []int{8}
}

func (m *RoleAttachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAttachment.Unmarshal(m, b)
}
func (m *RoleAttachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAttachment.Marshal(b, m, deterministic)
}
func (m *RoleAttachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAttachment.Merge(m, src)
}
func (m *RoleAttachment) XXX_Size() int {
	return xxx_messageInfo_RoleAttachment.Size(m)
}
func (m *RoleAttachment) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAttachment.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAttachment proto.InternalMessageInfo

func (m *RoleAttachment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RoleAttachment) GetCompositeRoleId() string {
	if m != nil {
		return m.CompositeRoleId
	}
	return ""
}

func (m *RoleAttachment) GetAttachedRoleId() string {
	if m != nil {
		return m.AttachedRoleId
	}
	return ""
}

func init() {
	proto.RegisterType((*RoleAttachmentCreateRequest)(nil), "v1.RoleAttachmentCreateRequest")
	proto.RegisterType((*RoleAttachmentCreateResponse)(nil), "v1.RoleAttachmentCreateResponse")
	proto.RegisterType((*RoleAttachmentGetRequest)(nil), "v1.RoleAttachmentGetRequest")
	proto.RegisterType((*RoleAttachmentGetResponse)(nil), "v1.RoleAttachmentGetResponse")
	proto.RegisterType((*RoleAttachmentDeleteRequest)(nil), "v1.RoleAttachmentDeleteRequest")
	proto.RegisterType((*RoleAttachmentDeleteResponse)(nil), "v1.RoleAttachmentDeleteResponse")
	proto.RegisterType((*RoleAttachmentListRequest)(nil), "v1.RoleAttachmentListRequest")
	proto.RegisterType((*RoleAttachmentListResponse)(nil), "v1.RoleAttachmentListResponse")
	proto.RegisterType((*RoleAttachment)(nil), "v1.RoleAttachment")
}

func init() { proto.RegisterFile("role_attachments.proto", fileDescriptor_c32d84988da10a62) }

var fileDescriptor_c32d84988da10a62 = []byte{
	// 872 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x41, 0x8f, 0xdb, 0x44,
	0x18, 0x95, 0x9d, 0x25, 0xa1, 0x83, 0xd8, 0x94, 0x01, 0xb6, 0x59, 0x37, 0xb4, 0x66, 0x8a, 0x44,
	0xb5, 0xdd, 0xc4, 0xca, 0x12, 0xa9, 0x52, 0x4e, 0xec, 0xb6, 0xd2, 0xaa, 0x62, 0x57, 0xa0, 0x9c,
	0xb8, 0xa0, 0x68, 0xd6, 0xfe, 0xea, 0x58, 0xd8, 0x1e, 0xe3, 0x99, 0x4d, 0x2a, 0x55, 0xbd, 0x70,
	0xe5, 0x56, 0xc4, 0x1f, 0xe0, 0x54, 0x71, 0x42, 0xca, 0x85, 0x0b, 0x12, 0x2a, 0xb7, 0x72, 0xe4,
	0x80, 0xb8, 0x73, 0xe3, 0x04, 0xb9, 0x50, 0x4e, 0x68, 0xc6, 0x76, 0xb2, 0xe3, 0x38, 0x91, 0x50,
	0xb9, 0x70, 0x4a, 0x32, 0xdf, 0x37, 0xef, 0xbd, 0xef, 0xcd, 0xcb, 0xd8, 0x68, 0x27, 0x65, 0x21,
	0x8c, 0xa8, 0x10, 0xd4, 0x1d, 0x47, 0x10, 0x0b, 0xde, 0x4d, 0x52, 0x26, 0x18, 0x36, 0x27, 0x3d,
	0xab, 0xed, 0x33, 0xe6, 0x87, 0xe0, 0xd0, 0x24, 0x70, 0x68, 0x1c, 0x33, 0x41, 0x45, 0xc0, 0xe2,
	0xbc, 0xc3, 0xda, 0x57, 0x1f, 0x6e, 0xc7, 0x87, 0xb8, 0xc3, 0xa7, 0xd4, 0xf7, 0x21, 0x75, 0x58,
	0xa2, 0x3a, 0x2a, 0xba, 0x5f, 0xcd, 0x4b, 0xf9, 0x4f, 0xc4, 0x13, 0x70, 0xb3, 0xef, 0xe4, 0x2b,
	0x03, 0x5d, 0x1d, 0xb2, 0x10, 0x0e, 0x17, 0x22, 0xee, 0xa4, 0x40, 0x05, 0x0c, 0xe1, 0xb3, 0x73,
	0xe0, 0x02, 0x77, 0xd0, 0x56, 0x04, 0x82, 0xb6, 0x0c, 0xdb, 0xb8, 0xf9, 0xca, 0xc1, 0x6e, 0x77,
	0xd2, 0xeb, 0x6a, 0x0d, 0xa7, 0x20, 0xa8, 0x47, 0x05, 0x1d, 0xaa, 0x36, 0x7c, 0x8c, 0x9a, 0xa5,
	0x99, 0x5a, 0xa6, 0xda, 0x89, 0xe5, 0x4e, 0x9d, 0xe8, 0x08, 0xfd, 0xf9, 0xd7, 0xac, 0xf1, 0xd2,
	0xb7, 0xf3, 0x59, 0xc3, 0x18, 0x6e, 0xa7, 0x5a, 0x8d, 0xfc, 0x61, 0xa0, 0x76, 0xb5, 0x2e, 0x9e,
	0xb0, 0x98, 0x03, 0x1e, 0x68, 0xc2, 0xac, 0x8b, 0xc2, 0xb2, 0x8e, 0x42, 0x99, 0x46, 0xf3, 0xdf,
	0xaa, 0xc4, 0xef, 0x23, 0x94, 0x52, 0x01, 0xa3, 0x30, 0x88, 0x02, 0xd1, 0xaa, 0x29, 0x8c, 0x37,
	0x15, 0x06, 0x15, 0x70, 0x22, 0x17, 0x2b, 0x55, 0x5c, 0x4a, 0x8b, 0xf2, 0x00, 0xfd, 0x2d, 0x97,
	0x9f, 0xc8, 0x65, 0x72, 0x86, 0x5a, 0x3a, 0xf7, 0x31, 0x88, 0xe2, 0x1c, 0xf6, 0xb4, 0x71, 0x77,
	0x24, 0xc7, 0xb2, 0x5a, 0x3a, 0x04, 0x0b, 0x99, 0x81, 0xa7, 0x26, 0xba, 0xa4, 0xd1, 0x9a, 0x81,
	0x47, 0x7e, 0x37, 0xd0, 0x6e, 0x05, 0x49, 0x6e, 0xea, 0x6d, 0x8d, 0xe5, 0xca, 0x82, 0xe5, 0xff,
	0xea, 0xe8, 0xb8, 0x1c, 0xee, 0xbb, 0x10, 0xc2, 0xc6, 0x70, 0x6b, 0x0d, 0xff, 0xc2, 0xd7, 0x27,
	0x2b, 0x79, 0x2d, 0x90, 0xd6, 0xe7, 0x55, 0xef, 0xd8, 0xe0, 0xae, 0x6e, 0x8a, 0xf9, 0x82, 0xa6,
	0x84, 0xe5, 0x04, 0x9c, 0x04, 0x7c, 0x91, 0xb3, 0x5b, 0xab, 0x09, 0xb8, 0x50, 0x2e, 0x19, 0x42,
	0x50, 0xfd, 0x7e, 0x10, 0x0a, 0x48, 0x2b, 0x4c, 0xc9, 0x2b, 0xe4, 0x17, 0x03, 0x59, 0x55, 0x74,
	0xb9, 0x2d, 0xfb, 0x1a, 0x5f, 0x6b, 0xc9, 0xa7, 0x9b, 0x92, 0x13, 0xde, 0x43, 0x97, 0xcb, 0x57,
	0x66, 0xcb, 0xb4, 0x6b, 0x1b, 0x73, 0xf6, 0x9d, 0x92, 0xd3, 0xd4, 0x73, 0xc6, 0x5f, 0x3c, 0x68,
	0xe4, 0xfb, 0x1a, 0xda, 0xd6, 0x19, 0x71, 0x5f, 0x25, 0xc4, 0x50, 0x66, 0xbc, 0x23, 0x77, 0x5d,
	0xff, 0x7a, 0x3e, 0x6b, 0x98, 0xf7, 0xee, 0xaa, 0xcd, 0xcf, 0xe6, 0xb3, 0x46, 0x69, 0x87, 0xcc,
	0x0e, 0xfe, 0x18, 0xbd, 0xe6, 0xb2, 0x28, 0x61, 0x3c, 0x10, 0x30, 0x52, 0xf3, 0x2d, 0x62, 0xb6,
	0x2f, 0x41, 0xde, 0x95, 0x20, 0xcd, 0x3b, 0x45, 0x8b, 0x44, 0xc8, 0x11, 0x7f, 0x28, 0x60, 0xb7,
	0xe4, 0xe2, 0xb0, 0xe9, 0x6a, 0x3d, 0x1e, 0x3e, 0x45, 0x97, 0x33, 0xab, 0xc0, 0x5b, 0x00, 0xd7,
	0x14, 0xf0, 0x0d, 0x09, 0x7c, 0x4d, 0x02, 0xd7, 0xd7, 0xe3, 0x6d, 0x17, 0x9b, 0x33, 0xb8, 0xc1,
	0x4f, 0xc6, 0xe3, 0xc3, 0x0f, 0x0f, 0x4e, 0xf1, 0x07, 0x0f, 0x6d, 0x12, 0x78, 0x64, 0x60, 0x93,
	0x94, 0x76, 0x0e, 0xfa, 0x7d, 0xb2, 0x6f, 0x93, 0x95, 0x09, 0x54, 0xb9, 0xd3, 0xeb, 0xf5, 0x64,
	0xb5, 0xac, 0x82, 0x0c, 0x48, 0xda, 0xe9, 0xf7, 0xfb, 0xe4, 0x91, 0x0c, 0xa6, 0x90, 0x3a, 0x5e,
	0x56, 0x94, 0x2c, 0x04, 0x95, 0xd2, 0xa7, 0xf3, 0x59, 0xe3, 0x13, 0xb9, 0x7c, 0x53, 0xdc, 0x1f,
	0xc1, 0x03, 0x1a, 0x25, 0x21, 0x70, 0xa7, 0x74, 0xf8, 0xa3, 0x14, 0x38, 0x3b, 0x4f, 0x5d, 0xe8,
	0x8a, 0x07, 0xe2, 0x9b, 0xf9, 0xac, 0x71, 0x6b, 0x53, 0xb3, 0x3c, 0xcb, 0xd1, 0xb2, 0xff, 0xe0,
	0xc7, 0x2d, 0xd4, 0x1c, 0x96, 0x52, 0xf1, 0xab, 0x81, 0xea, 0xd9, 0x63, 0x04, 0x5f, 0x5f, 0x4d,
	0x94, 0xf6, 0xe4, 0xb3, 0xec, 0xf5, 0x0d, 0x59, 0x78, 0xc9, 0x17, 0xc6, 0xe3, 0xc3, 0x80, 0xf8,
	0x88, 0x9c, 0x00, 0x4d, 0x63, 0x7b, 0xcc, 0xa6, 0xb6, 0x60, 0x76, 0x44, 0x3f, 0x05, 0x9b, 0xda,
	0xa5, 0xe8, 0x1c, 0x8e, 0x85, 0x48, 0xf8, 0xc0, 0x71, 0xa6, 0xd3, 0x69, 0x97, 0x8b, 0x94, 0xc5,
	0xbe, 0x17, 0x75, 0x5d, 0x16, 0x39, 0x1e, 0x73, 0xb9, 0x7a, 0x0d, 0xe0, 0x90, 0x4e, 0x02, 0x17,
	0xb8, 0x53, 0x52, 0x7d, 0x23, 0xa3, 0xfd, 0xfc, 0xe7, 0xdf, 0xbe, 0x34, 0x77, 0xc9, 0x1b, 0xce,
	0xa4, 0x57, 0x9e, 0x9d, 0x0f, 0x8c, 0x3d, 0x0c, 0xa8, 0x76, 0x0c, 0x02, 0xb7, 0x57, 0x65, 0x2f,
	0x9f, 0x24, 0xd6, 0x5b, 0x6b, 0xaa, 0xf9, 0x44, 0x6f, 0x2b, 0x92, 0xab, 0x78, 0xb7, 0x8a, 0xc4,
	0x79, 0x18, 0x78, 0x8f, 0x70, 0x8c, 0xea, 0xd9, 0xc5, 0x56, 0xe5, 0xa0, 0x76, 0xbd, 0x56, 0x39,
	0xa8, 0xdf, 0x89, 0x05, 0xdf, 0xde, 0x06, 0x3e, 0x17, 0x6d, 0xc9, 0x1b, 0x03, 0x57, 0x28, 0xbf,
	0x70, 0x73, 0x59, 0xd7, 0xd6, 0x95, 0x73, 0xa6, 0xb6, 0x62, 0xda, 0xc1, 0x95, 0xf6, 0x59, 0xad,
	0x67, 0xcf, 0x67, 0x8d, 0xd7, 0x9f, 0x3e, 0x5f, 0xf9, 0x03, 0x1f, 0xdd, 0x46, 0x6d, 0x97, 0x45,
	0xcb, 0x13, 0xa3, 0x49, 0x20, 0x99, 0x92, 0xf0, 0x3c, 0x3a, 0x0b, 0x62, 0xff, 0xe8, 0x4a, 0xe9,
	0xb0, 0x3e, 0xca, 0x0b, 0x67, 0x75, 0xf5, 0x02, 0xf6, 0xde, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x55, 0x19, 0xe5, 0xb6, 0x05, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoleAttachmentsClient is the client API for RoleAttachments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoleAttachmentsClient interface {
	// Create registers a new RoleAttachment.
	Create(ctx context.Context, in *RoleAttachmentCreateRequest, opts ...grpc.CallOption) (*RoleAttachmentCreateResponse, error)
	// Get reads one RoleAttachment by ID.
	Get(ctx context.Context, in *RoleAttachmentGetRequest, opts ...grpc.CallOption) (*RoleAttachmentGetResponse, error)
	// Delete removes a RoleAttachment by ID.
	Delete(ctx context.Context, in *RoleAttachmentDeleteRequest, opts ...grpc.CallOption) (*RoleAttachmentDeleteResponse, error)
	// List gets a list of RoleAttachments matching a given set of criteria.
	List(ctx context.Context, in *RoleAttachmentListRequest, opts ...grpc.CallOption) (*RoleAttachmentListResponse, error)
}

type roleAttachmentsClient struct {
	cc *grpc.ClientConn
}

func NewRoleAttachmentsClient(cc *grpc.ClientConn) RoleAttachmentsClient {
	return &roleAttachmentsClient{cc}
}

func (c *roleAttachmentsClient) Create(ctx context.Context, in *RoleAttachmentCreateRequest, opts ...grpc.CallOption) (*RoleAttachmentCreateResponse, error) {
	out := new(RoleAttachmentCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleAttachments/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleAttachmentsClient) Get(ctx context.Context, in *RoleAttachmentGetRequest, opts ...grpc.CallOption) (*RoleAttachmentGetResponse, error) {
	out := new(RoleAttachmentGetResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleAttachments/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleAttachmentsClient) Delete(ctx context.Context, in *RoleAttachmentDeleteRequest, opts ...grpc.CallOption) (*RoleAttachmentDeleteResponse, error) {
	out := new(RoleAttachmentDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleAttachments/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleAttachmentsClient) List(ctx context.Context, in *RoleAttachmentListRequest, opts ...grpc.CallOption) (*RoleAttachmentListResponse, error) {
	out := new(RoleAttachmentListResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleAttachments/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleAttachmentsServer is the server API for RoleAttachments service.
type RoleAttachmentsServer interface {
	// Create registers a new RoleAttachment.
	Create(context.Context, *RoleAttachmentCreateRequest) (*RoleAttachmentCreateResponse, error)
	// Get reads one RoleAttachment by ID.
	Get(context.Context, *RoleAttachmentGetRequest) (*RoleAttachmentGetResponse, error)
	// Delete removes a RoleAttachment by ID.
	Delete(context.Context, *RoleAttachmentDeleteRequest) (*RoleAttachmentDeleteResponse, error)
	// List gets a list of RoleAttachments matching a given set of criteria.
	List(context.Context, *RoleAttachmentListRequest) (*RoleAttachmentListResponse, error)
}

// UnimplementedRoleAttachmentsServer can be embedded to have forward compatible implementations.
type UnimplementedRoleAttachmentsServer struct {
}

func (*UnimplementedRoleAttachmentsServer) Create(ctx context.Context, req *RoleAttachmentCreateRequest) (*RoleAttachmentCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRoleAttachmentsServer) Get(ctx context.Context, req *RoleAttachmentGetRequest) (*RoleAttachmentGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRoleAttachmentsServer) Delete(ctx context.Context, req *RoleAttachmentDeleteRequest) (*RoleAttachmentDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedRoleAttachmentsServer) List(ctx context.Context, req *RoleAttachmentListRequest) (*RoleAttachmentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterRoleAttachmentsServer(s *grpc.Server, srv RoleAttachmentsServer) {
	s.RegisterService(&_RoleAttachments_serviceDesc, srv)
}

func _RoleAttachments_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAttachmentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleAttachmentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleAttachments/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleAttachmentsServer).Create(ctx, req.(*RoleAttachmentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleAttachments_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAttachmentGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleAttachmentsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleAttachments/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleAttachmentsServer).Get(ctx, req.(*RoleAttachmentGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleAttachments_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAttachmentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleAttachmentsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleAttachments/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleAttachmentsServer).Delete(ctx, req.(*RoleAttachmentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleAttachments_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAttachmentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleAttachmentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleAttachments/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleAttachmentsServer).List(ctx, req.(*RoleAttachmentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoleAttachments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RoleAttachments",
	HandlerType: (*RoleAttachmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RoleAttachments_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RoleAttachments_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RoleAttachments_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RoleAttachments_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role_attachments.proto",
}
