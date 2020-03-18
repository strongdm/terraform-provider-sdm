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
// source: spec.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

// AlreadyExistsError is used when an entity already exists in the system
type AlreadyExistsError struct {
	Entity               string   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlreadyExistsError) Reset()         { *m = AlreadyExistsError{} }
func (m *AlreadyExistsError) String() string { return proto.CompactTextString(m) }
func (*AlreadyExistsError) ProtoMessage()    {}
func (*AlreadyExistsError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{0}
}

func (m *AlreadyExistsError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlreadyExistsError.Unmarshal(m, b)
}
func (m *AlreadyExistsError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlreadyExistsError.Marshal(b, m, deterministic)
}
func (m *AlreadyExistsError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlreadyExistsError.Merge(m, src)
}
func (m *AlreadyExistsError) XXX_Size() int {
	return xxx_messageInfo_AlreadyExistsError.Size(m)
}
func (m *AlreadyExistsError) XXX_DiscardUnknown() {
	xxx_messageInfo_AlreadyExistsError.DiscardUnknown(m)
}

var xxx_messageInfo_AlreadyExistsError proto.InternalMessageInfo

func (m *AlreadyExistsError) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

// NotFoundError is used when an entity does not exist in the system
type NotFoundError struct {
	Entity               string   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotFoundError) Reset()         { *m = NotFoundError{} }
func (m *NotFoundError) String() string { return proto.CompactTextString(m) }
func (*NotFoundError) ProtoMessage()    {}
func (*NotFoundError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{1}
}

func (m *NotFoundError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotFoundError.Unmarshal(m, b)
}
func (m *NotFoundError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotFoundError.Marshal(b, m, deterministic)
}
func (m *NotFoundError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotFoundError.Merge(m, src)
}
func (m *NotFoundError) XXX_Size() int {
	return xxx_messageInfo_NotFoundError.Size(m)
}
func (m *NotFoundError) XXX_DiscardUnknown() {
	xxx_messageInfo_NotFoundError.DiscardUnknown(m)
}

var xxx_messageInfo_NotFoundError proto.InternalMessageInfo

func (m *NotFoundError) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

// BadRequestError identifies a bad request sent by the client
type BadRequestError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BadRequestError) Reset()         { *m = BadRequestError{} }
func (m *BadRequestError) String() string { return proto.CompactTextString(m) }
func (*BadRequestError) ProtoMessage()    {}
func (*BadRequestError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{2}
}

func (m *BadRequestError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BadRequestError.Unmarshal(m, b)
}
func (m *BadRequestError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BadRequestError.Marshal(b, m, deterministic)
}
func (m *BadRequestError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadRequestError.Merge(m, src)
}
func (m *BadRequestError) XXX_Size() int {
	return xxx_messageInfo_BadRequestError.Size(m)
}
func (m *BadRequestError) XXX_DiscardUnknown() {
	xxx_messageInfo_BadRequestError.DiscardUnknown(m)
}

var xxx_messageInfo_BadRequestError proto.InternalMessageInfo

// AuthenticationError is used to specify an authentication failure condition
type AuthenticationError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationError) Reset()         { *m = AuthenticationError{} }
func (m *AuthenticationError) String() string { return proto.CompactTextString(m) }
func (*AuthenticationError) ProtoMessage()    {}
func (*AuthenticationError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{3}
}

func (m *AuthenticationError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationError.Unmarshal(m, b)
}
func (m *AuthenticationError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationError.Marshal(b, m, deterministic)
}
func (m *AuthenticationError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationError.Merge(m, src)
}
func (m *AuthenticationError) XXX_Size() int {
	return xxx_messageInfo_AuthenticationError.Size(m)
}
func (m *AuthenticationError) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationError.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationError proto.InternalMessageInfo

// PermissionError is used to specify a permissions violation
type PermissionError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PermissionError) Reset()         { *m = PermissionError{} }
func (m *PermissionError) String() string { return proto.CompactTextString(m) }
func (*PermissionError) ProtoMessage()    {}
func (*PermissionError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{4}
}

func (m *PermissionError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionError.Unmarshal(m, b)
}
func (m *PermissionError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionError.Marshal(b, m, deterministic)
}
func (m *PermissionError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionError.Merge(m, src)
}
func (m *PermissionError) XXX_Size() int {
	return xxx_messageInfo_PermissionError.Size(m)
}
func (m *PermissionError) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionError.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionError proto.InternalMessageInfo

// InternalError is used to specify an internal system error
type InternalError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalError) Reset()         { *m = InternalError{} }
func (m *InternalError) String() string { return proto.CompactTextString(m) }
func (*InternalError) ProtoMessage()    {}
func (*InternalError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{5}
}

func (m *InternalError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalError.Unmarshal(m, b)
}
func (m *InternalError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalError.Marshal(b, m, deterministic)
}
func (m *InternalError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalError.Merge(m, src)
}
func (m *InternalError) XXX_Size() int {
	return xxx_messageInfo_InternalError.Size(m)
}
func (m *InternalError) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalError.DiscardUnknown(m)
}

var xxx_messageInfo_InternalError proto.InternalMessageInfo

// RateLimitError is used for rate limit excess condition
type RateLimitError struct {
	RateLimit            *RateLimitMetadata `protobuf:"bytes,1,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RateLimitError) Reset()         { *m = RateLimitError{} }
func (m *RateLimitError) String() string { return proto.CompactTextString(m) }
func (*RateLimitError) ProtoMessage()    {}
func (*RateLimitError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{6}
}

func (m *RateLimitError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitError.Unmarshal(m, b)
}
func (m *RateLimitError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitError.Marshal(b, m, deterministic)
}
func (m *RateLimitError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitError.Merge(m, src)
}
func (m *RateLimitError) XXX_Size() int {
	return xxx_messageInfo_RateLimitError.Size(m)
}
func (m *RateLimitError) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitError.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitError proto.InternalMessageInfo

func (m *RateLimitError) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// CreateRequestMetadata is reserved for future use.
type CreateRequestMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequestMetadata) Reset()         { *m = CreateRequestMetadata{} }
func (m *CreateRequestMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateRequestMetadata) ProtoMessage()    {}
func (*CreateRequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{7}
}

func (m *CreateRequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequestMetadata.Unmarshal(m, b)
}
func (m *CreateRequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequestMetadata.Marshal(b, m, deterministic)
}
func (m *CreateRequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequestMetadata.Merge(m, src)
}
func (m *CreateRequestMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateRequestMetadata.Size(m)
}
func (m *CreateRequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequestMetadata proto.InternalMessageInfo

// CreateResponseMetadata is reserved for future use.
type CreateResponseMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponseMetadata) Reset()         { *m = CreateResponseMetadata{} }
func (m *CreateResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateResponseMetadata) ProtoMessage()    {}
func (*CreateResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{8}
}

func (m *CreateResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponseMetadata.Unmarshal(m, b)
}
func (m *CreateResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponseMetadata.Marshal(b, m, deterministic)
}
func (m *CreateResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponseMetadata.Merge(m, src)
}
func (m *CreateResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateResponseMetadata.Size(m)
}
func (m *CreateResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponseMetadata proto.InternalMessageInfo

// GetRequestMetadata is reserved for future use.
type GetRequestMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequestMetadata) Reset()         { *m = GetRequestMetadata{} }
func (m *GetRequestMetadata) String() string { return proto.CompactTextString(m) }
func (*GetRequestMetadata) ProtoMessage()    {}
func (*GetRequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{9}
}

func (m *GetRequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequestMetadata.Unmarshal(m, b)
}
func (m *GetRequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequestMetadata.Marshal(b, m, deterministic)
}
func (m *GetRequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequestMetadata.Merge(m, src)
}
func (m *GetRequestMetadata) XXX_Size() int {
	return xxx_messageInfo_GetRequestMetadata.Size(m)
}
func (m *GetRequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequestMetadata proto.InternalMessageInfo

// GetResponseMetadata is reserved for future use.
type GetResponseMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponseMetadata) Reset()         { *m = GetResponseMetadata{} }
func (m *GetResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*GetResponseMetadata) ProtoMessage()    {}
func (*GetResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{10}
}

func (m *GetResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponseMetadata.Unmarshal(m, b)
}
func (m *GetResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponseMetadata.Marshal(b, m, deterministic)
}
func (m *GetResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponseMetadata.Merge(m, src)
}
func (m *GetResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_GetResponseMetadata.Size(m)
}
func (m *GetResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponseMetadata proto.InternalMessageInfo

// UpdateRequestMetadata is reserved for future use.
type UpdateRequestMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequestMetadata) Reset()         { *m = UpdateRequestMetadata{} }
func (m *UpdateRequestMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateRequestMetadata) ProtoMessage()    {}
func (*UpdateRequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{11}
}

func (m *UpdateRequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequestMetadata.Unmarshal(m, b)
}
func (m *UpdateRequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequestMetadata.Marshal(b, m, deterministic)
}
func (m *UpdateRequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequestMetadata.Merge(m, src)
}
func (m *UpdateRequestMetadata) XXX_Size() int {
	return xxx_messageInfo_UpdateRequestMetadata.Size(m)
}
func (m *UpdateRequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequestMetadata proto.InternalMessageInfo

// UpdateResponseMetadata is reserved for future use.
type UpdateResponseMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponseMetadata) Reset()         { *m = UpdateResponseMetadata{} }
func (m *UpdateResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateResponseMetadata) ProtoMessage()    {}
func (*UpdateResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{12}
}

func (m *UpdateResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponseMetadata.Unmarshal(m, b)
}
func (m *UpdateResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponseMetadata.Marshal(b, m, deterministic)
}
func (m *UpdateResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponseMetadata.Merge(m, src)
}
func (m *UpdateResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_UpdateResponseMetadata.Size(m)
}
func (m *UpdateResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponseMetadata proto.InternalMessageInfo

// DeleteRequestMetadata is reserved for future use.
type DeleteRequestMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequestMetadata) Reset()         { *m = DeleteRequestMetadata{} }
func (m *DeleteRequestMetadata) String() string { return proto.CompactTextString(m) }
func (*DeleteRequestMetadata) ProtoMessage()    {}
func (*DeleteRequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{13}
}

func (m *DeleteRequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequestMetadata.Unmarshal(m, b)
}
func (m *DeleteRequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequestMetadata.Marshal(b, m, deterministic)
}
func (m *DeleteRequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequestMetadata.Merge(m, src)
}
func (m *DeleteRequestMetadata) XXX_Size() int {
	return xxx_messageInfo_DeleteRequestMetadata.Size(m)
}
func (m *DeleteRequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequestMetadata proto.InternalMessageInfo

// DeleteResponseMetadata is reserved for future use.
type DeleteResponseMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponseMetadata) Reset()         { *m = DeleteResponseMetadata{} }
func (m *DeleteResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*DeleteResponseMetadata) ProtoMessage()    {}
func (*DeleteResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{14}
}

func (m *DeleteResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponseMetadata.Unmarshal(m, b)
}
func (m *DeleteResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponseMetadata.Marshal(b, m, deterministic)
}
func (m *DeleteResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponseMetadata.Merge(m, src)
}
func (m *DeleteResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_DeleteResponseMetadata.Size(m)
}
func (m *DeleteResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponseMetadata proto.InternalMessageInfo

// ListRequestMetadata specifies paging parameters for listing entities. If this
// metadata is not provided, the default behavior is to return the first page of
// entities, along with a cursor which can be used to fetch the remaining pages.
type ListRequestMetadata struct {
	// The cursor specifies where to start fetching entities in the total list
	// of all entities. If the cursor is non-empty, the page and limit
	// parameters are ignored. See ListResponseMetadata.next_cursor.
	Cursor string `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// The page number to fetch. Use of this parameter is not recommended. Use
	// the cursor instead.
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// The number of entities to fetch in a single page. If not specified, a
	// default value will be used.
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequestMetadata) Reset()         { *m = ListRequestMetadata{} }
func (m *ListRequestMetadata) String() string { return proto.CompactTextString(m) }
func (*ListRequestMetadata) ProtoMessage()    {}
func (*ListRequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{15}
}

func (m *ListRequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequestMetadata.Unmarshal(m, b)
}
func (m *ListRequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequestMetadata.Marshal(b, m, deterministic)
}
func (m *ListRequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequestMetadata.Merge(m, src)
}
func (m *ListRequestMetadata) XXX_Size() int {
	return xxx_messageInfo_ListRequestMetadata.Size(m)
}
func (m *ListRequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequestMetadata proto.InternalMessageInfo

func (m *ListRequestMetadata) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *ListRequestMetadata) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListRequestMetadata) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// ListResponseMetadata contains paging information about the remaining entities
// in a list request.
type ListResponseMetadata struct {
	// A cursor to fetch the next page. If the cursor is an empty string, there
	// are no more entities to fetch. If the cursor is non-empty, make another
	// list request and pass the cursor value in the metadata.
	NextCursor string `protobuf:"bytes,1,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	// The total count of all entities matching the criteria of a list request.
	// Note that this value may change between page requests.
	Total                int32    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponseMetadata) Reset()         { *m = ListResponseMetadata{} }
func (m *ListResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*ListResponseMetadata) ProtoMessage()    {}
func (*ListResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{16}
}

func (m *ListResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponseMetadata.Unmarshal(m, b)
}
func (m *ListResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponseMetadata.Marshal(b, m, deterministic)
}
func (m *ListResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponseMetadata.Merge(m, src)
}
func (m *ListResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_ListResponseMetadata.Size(m)
}
func (m *ListResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponseMetadata proto.InternalMessageInfo

func (m *ListResponseMetadata) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

func (m *ListResponseMetadata) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

// RateLimitMetadata contains information about remaining requests avaialable
// to the user over some timeframe.
type RateLimitMetadata struct {
	// How many total requests the user/token is authorized to make before being
	// rate limited.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// How many remaining requests out of the limit are still avaialable.
	Remaining int64 `protobuf:"varint,2,opt,name=remaining,proto3" json:"remaining,omitempty"`
	// The time when remaining will be reset to limit.
	ResetAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=reset_at,json=resetAt,proto3" json:"reset_at,omitempty"`
	// The bucket this user/token is associated with, which may be shared between
	// multiple users/tokens.
	Bucket               string   `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitMetadata) Reset()         { *m = RateLimitMetadata{} }
func (m *RateLimitMetadata) String() string { return proto.CompactTextString(m) }
func (*RateLimitMetadata) ProtoMessage()    {}
func (*RateLimitMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{17}
}

func (m *RateLimitMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitMetadata.Unmarshal(m, b)
}
func (m *RateLimitMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitMetadata.Marshal(b, m, deterministic)
}
func (m *RateLimitMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitMetadata.Merge(m, src)
}
func (m *RateLimitMetadata) XXX_Size() int {
	return xxx_messageInfo_RateLimitMetadata.Size(m)
}
func (m *RateLimitMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitMetadata proto.InternalMessageInfo

func (m *RateLimitMetadata) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RateLimitMetadata) GetRemaining() int64 {
	if m != nil {
		return m.Remaining
	}
	return 0
}

func (m *RateLimitMetadata) GetResetAt() *timestamp.Timestamp {
	if m != nil {
		return m.ResetAt
	}
	return nil
}

func (m *RateLimitMetadata) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func init() {
	proto.RegisterType((*AlreadyExistsError)(nil), "v1.AlreadyExistsError")
	proto.RegisterType((*NotFoundError)(nil), "v1.NotFoundError")
	proto.RegisterType((*BadRequestError)(nil), "v1.BadRequestError")
	proto.RegisterType((*AuthenticationError)(nil), "v1.AuthenticationError")
	proto.RegisterType((*PermissionError)(nil), "v1.PermissionError")
	proto.RegisterType((*InternalError)(nil), "v1.InternalError")
	proto.RegisterType((*RateLimitError)(nil), "v1.RateLimitError")
	proto.RegisterType((*CreateRequestMetadata)(nil), "v1.CreateRequestMetadata")
	proto.RegisterType((*CreateResponseMetadata)(nil), "v1.CreateResponseMetadata")
	proto.RegisterType((*GetRequestMetadata)(nil), "v1.GetRequestMetadata")
	proto.RegisterType((*GetResponseMetadata)(nil), "v1.GetResponseMetadata")
	proto.RegisterType((*UpdateRequestMetadata)(nil), "v1.UpdateRequestMetadata")
	proto.RegisterType((*UpdateResponseMetadata)(nil), "v1.UpdateResponseMetadata")
	proto.RegisterType((*DeleteRequestMetadata)(nil), "v1.DeleteRequestMetadata")
	proto.RegisterType((*DeleteResponseMetadata)(nil), "v1.DeleteResponseMetadata")
	proto.RegisterType((*ListRequestMetadata)(nil), "v1.ListRequestMetadata")
	proto.RegisterType((*ListResponseMetadata)(nil), "v1.ListResponseMetadata")
	proto.RegisterType((*RateLimitMetadata)(nil), "v1.RateLimitMetadata")
}

func init() { proto.RegisterFile("spec.proto", fileDescriptor_423806180556987f) }

var fileDescriptor_423806180556987f = []byte{
	// 723 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x4e, 0xdb, 0x4c,
	0x14, 0x96, 0x09, 0x49, 0xe0, 0xa0, 0xfc, 0x3f, 0x1d, 0x12, 0x48, 0x23, 0x2e, 0x8e, 0xd5, 0x05,
	0xaa, 0x1a, 0x9b, 0xd0, 0x55, 0xb3, 0x22, 0xe1, 0x52, 0x21, 0x41, 0x85, 0x52, 0xaa, 0x6e, 0x2a,
	0xa1, 0x89, 0x33, 0x38, 0x56, 0xe3, 0x19, 0x33, 0x73, 0x4c, 0xe0, 0x35, 0xfa, 0x08, 0x5d, 0x75,
	0x59, 0x29, 0x6f, 0xd2, 0x45, 0x1f, 0xa4, 0x6c, 0x50, 0x57, 0xd5, 0xd8, 0x31, 0x4d, 0xd2, 0x88,
	0x4a, 0x5d, 0x25, 0xf3, 0x5d, 0xc7, 0x67, 0xec, 0x01, 0x50, 0x21, 0x73, 0xed, 0x50, 0x0a, 0x14,
	0x64, 0xee, 0xba, 0x5e, 0x29, 0x88, 0x10, 0x7d, 0xc1, 0x55, 0x02, 0x55, 0x5e, 0xc4, 0x3f, 0x6e,
	0xcd, 0x63, 0xbc, 0xa6, 0x06, 0xd4, 0xf3, 0x98, 0x74, 0x46, 0x0a, 0x87, 0x72, 0x2e, 0x90, 0x8e,
	0xab, 0xb7, 0x3c, 0x21, 0xbc, 0x3e, 0x73, 0xe2, 0x55, 0x27, 0xba, 0x74, 0xd0, 0x0f, 0x98, 0x42,
	0x1a, 0x84, 0x89, 0xc0, 0xba, 0x04, 0xd2, 0xec, 0x4b, 0x46, 0xbb, 0xb7, 0x87, 0x37, 0xbe, 0x42,
	0x75, 0x28, 0xa5, 0x90, 0xa4, 0x06, 0x39, 0xc6, 0xd1, 0xc7, 0xdb, 0xb2, 0x61, 0x1a, 0xdb, 0x8b,
	0xad, 0xd2, 0x8f, 0xfb, 0x61, 0x7e, 0xf9, 0xf3, 0xdd, 0x30, 0x9f, 0x3b, 0x8c, 0xf1, 0xaf, 0x77,
	0xc3, 0xbc, 0xd1, 0x1e, 0x89, 0x1a, 0xd5, 0x9f, 0xf7, 0xc3, 0xfc, 0xba, 0xa6, 0x67, 0x84, 0x69,
	0x69, 0xce, 0xfa, 0x00, 0x85, 0x37, 0x02, 0x8f, 0x44, 0xc4, 0xbb, 0xff, 0x54, 0xb1, 0xae, 0x2b,
	0xd6, 0x34, 0x3d, 0x99, 0xa3, 0x55, 0x59, 0xab, 0x0e, 0xff, 0xb7, 0x68, 0xb7, 0xcd, 0xae, 0x22,
	0xa6, 0x30, 0xc6, 0x1b, 0x9b, 0xda, 0xf0, 0x54, 0x1b, 0xa6, 0x29, 0x6d, 0xc9, 0x58, 0xaf, 0x60,
	0xa5, 0x19, 0x61, 0x4f, 0xc7, 0xbb, 0xf1, 0xc8, 0x12, 0x9b, 0xa5, 0x6d, 0x1b, 0xda, 0x36, 0x8b,
	0xd6, 0xd6, 0x65, 0xdd, 0x76, 0xc6, 0x64, 0xe0, 0x2b, 0xf5, 0x60, 0x1b, 0x6b, 0x9b, 0xa2, 0xb4,
	0x25, 0x6f, 0xd5, 0xa0, 0x70, 0xcc, 0x91, 0x49, 0x4e, 0xfb, 0x89, 0x61, 0xec, 0x79, 0x26, 0x08,
	0x2d, 0x2f, 0x58, 0x57, 0xf0, 0x5f, 0x9b, 0x22, 0x3b, 0xf1, 0x03, 0x3f, 0xd9, 0x33, 0xd9, 0x03,
	0x90, 0x14, 0xd9, 0x45, 0x5f, 0x43, 0xf1, 0xc8, 0x96, 0x76, 0x4b, 0xf6, 0x75, 0xdd, 0x7e, 0xd0,
	0x9d, 0x32, 0xa4, 0x5d, 0x8a, 0xb4, 0x05, 0x7a, 0x92, 0xd9, 0x64, 0x7c, 0x8b, 0x32, 0xa5, 0x1b,
	0x1b, 0xba, 0xb1, 0xac, 0x1b, 0xa7, 0xb2, 0xb5, 0x72, 0xc1, 0x5a, 0x83, 0xd2, 0xbe, 0x64, 0x14,
	0xd9, 0x68, 0x54, 0x69, 0x9c, 0xf5, 0x0c, 0x56, 0x53, 0x42, 0x85, 0x82, 0x2b, 0x96, 0x32, 0x0d,
	0xd0, 0x89, 0xd9, 0x2f, 0xba, 0xc8, 0x2a, 0x02, 0x79, 0xcd, 0x70, 0xda, 0x5b, 0x85, 0x95, 0x18,
	0x7d, 0xc4, 0xb8, 0x06, 0xa5, 0x77, 0x61, 0x77, 0x76, 0x6f, 0x4a, 0x3c, 0x6e, 0x3f, 0x60, 0x7d,
	0x36, 0xd3, 0x9e, 0x12, 0x8f, 0xd8, 0xdf, 0xc3, 0xca, 0x89, 0xaf, 0xa6, 0xf7, 0x4d, 0x56, 0x21,
	0xe7, 0x46, 0x52, 0x09, 0x99, 0xbc, 0x9c, 0xed, 0xd1, 0x8a, 0x10, 0x98, 0x0f, 0xa9, 0xc7, 0xca,
	0x73, 0xa6, 0xb1, 0x9d, 0x6d, 0xc7, 0xff, 0x49, 0x11, 0xb2, 0xc9, 0xa1, 0x64, 0x62, 0x30, 0x59,
	0x58, 0xa7, 0x50, 0x4c, 0x82, 0x27, 0xcb, 0xc9, 0x16, 0x2c, 0x71, 0x76, 0x83, 0x17, 0x13, 0xf1,
	0xa0, 0xa1, 0xfd, 0xa4, 0xa2, 0x08, 0x59, 0x14, 0x48, 0xfb, 0xa3, 0x8e, 0x64, 0x61, 0x7d, 0x33,
	0xe0, 0xc9, 0x1f, 0x27, 0x4d, 0xcc, 0xb4, 0x5a, 0xc7, 0x64, 0x26, 0x0e, 0x3e, 0x21, 0xc8, 0x36,
	0x2c, 0x4a, 0x16, 0x50, 0x9f, 0xfb, 0xdc, 0x8b, 0x13, 0x33, 0x53, 0xaf, 0x47, 0x4a, 0x92, 0x26,
	0x2c, 0x48, 0xa6, 0x18, 0x5e, 0xd0, 0xe4, 0x49, 0x96, 0x76, 0x2b, 0x76, 0x72, 0x79, 0xd8, 0xe9,
	0xe5, 0x61, 0x9f, 0xa7, 0x97, 0xc7, 0x44, 0x48, 0x3e, 0xf6, 0x35, 0x91, 0x58, 0x90, 0xeb, 0x44,
	0xee, 0x47, 0x86, 0xe5, 0xf9, 0xf8, 0x93, 0x1e, 0x17, 0x8d, 0x98, 0xf1, 0xe1, 0xb7, 0xee, 0x8d,
	0x4f, 0xcd, 0xef, 0x06, 0x19, 0xc0, 0xc2, 0x5b, 0x94, 0x82, 0x7b, 0x07, 0xa7, 0x64, 0xf3, 0xbc,
	0xc7, 0x4c, 0x9f, 0x5f, 0x4a, 0xaa, 0x50, 0x46, 0x2e, 0x46, 0x92, 0x99, 0xd4, 0x75, 0x99, 0x52,
	0x66, 0xf3, 0xec, 0xd8, 0xb6, 0x0e, 0xc6, 0xb4, 0x56, 0x0f, 0x31, 0x54, 0x0d, 0xc7, 0x19, 0x0c,
	0x06, 0xb6, 0x8a, 0xd1, 0x6e, 0x60, 0xbb, 0x22, 0x70, 0xba, 0xc2, 0x55, 0x0e, 0x0d, 0x7d, 0xa7,
	0x52, 0x54, 0x51, 0x18, 0x0a, 0x89, 0x7b, 0xe3, 0xfc, 0x6e, 0xb6, 0x6e, 0xef, 0xd8, 0x3b, 0x95,
	0x65, 0x1a, 0xfa, 0x13, 0xc6, 0xe7, 0xc6, 0x9c, 0x3c, 0x82, 0xea, 0x09, 0xa3, 0x92, 0x9b, 0x81,
	0xd0, 0xdd, 0x1d, 0x11, 0xa1, 0x89, 0x3d, 0x66, 0xaa, 0x51, 0xad, 0xde, 0x07, 0xa9, 0xfe, 0xb5,
	0x1a, 0xd6, 0x5d, 0x11, 0xfc, 0xa6, 0x74, 0xd3, 0x75, 0xdd, 0x0e, 0xfb, 0x51, 0xd0, 0xf1, 0xb9,
	0xd7, 0xc9, 0xc5, 0x33, 0x7d, 0xf9, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x04, 0x7a, 0x31, 0xed,
	0x05, 0x00, 0x00,
}
