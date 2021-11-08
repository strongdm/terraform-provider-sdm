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
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AccountAttachmentsClient is the client API for AccountAttachments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountAttachmentsClient interface {
	// Create registers a new AccountAttachment.
	Create(ctx context.Context, in *AccountAttachmentCreateRequest, opts ...grpc.CallOption) (*AccountAttachmentCreateResponse, error)
	// Get reads one AccountAttachment by ID.
	Get(ctx context.Context, in *AccountAttachmentGetRequest, opts ...grpc.CallOption) (*AccountAttachmentGetResponse, error)
	// Delete removes a AccountAttachment by ID.
	Delete(ctx context.Context, in *AccountAttachmentDeleteRequest, opts ...grpc.CallOption) (*AccountAttachmentDeleteResponse, error)
	// List gets a list of AccountAttachments matching a given set of criteria.
	List(ctx context.Context, in *AccountAttachmentListRequest, opts ...grpc.CallOption) (*AccountAttachmentListResponse, error)
}

type accountAttachmentsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountAttachmentsClient(cc grpc.ClientConnInterface) AccountAttachmentsClient {
	return &accountAttachmentsClient{cc}
}

func (c *accountAttachmentsClient) Create(ctx context.Context, in *AccountAttachmentCreateRequest, opts ...grpc.CallOption) (*AccountAttachmentCreateResponse, error) {
	out := new(AccountAttachmentCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountAttachments/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAttachmentsClient) Get(ctx context.Context, in *AccountAttachmentGetRequest, opts ...grpc.CallOption) (*AccountAttachmentGetResponse, error) {
	out := new(AccountAttachmentGetResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountAttachments/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAttachmentsClient) Delete(ctx context.Context, in *AccountAttachmentDeleteRequest, opts ...grpc.CallOption) (*AccountAttachmentDeleteResponse, error) {
	out := new(AccountAttachmentDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountAttachments/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAttachmentsClient) List(ctx context.Context, in *AccountAttachmentListRequest, opts ...grpc.CallOption) (*AccountAttachmentListResponse, error) {
	out := new(AccountAttachmentListResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountAttachments/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountAttachmentsServer is the server API for AccountAttachments service.
// All implementations must embed UnimplementedAccountAttachmentsServer
// for forward compatibility
type AccountAttachmentsServer interface {
	// Create registers a new AccountAttachment.
	Create(context.Context, *AccountAttachmentCreateRequest) (*AccountAttachmentCreateResponse, error)
	// Get reads one AccountAttachment by ID.
	Get(context.Context, *AccountAttachmentGetRequest) (*AccountAttachmentGetResponse, error)
	// Delete removes a AccountAttachment by ID.
	Delete(context.Context, *AccountAttachmentDeleteRequest) (*AccountAttachmentDeleteResponse, error)
	// List gets a list of AccountAttachments matching a given set of criteria.
	List(context.Context, *AccountAttachmentListRequest) (*AccountAttachmentListResponse, error)
	mustEmbedUnimplementedAccountAttachmentsServer()
}

// UnimplementedAccountAttachmentsServer must be embedded to have forward compatible implementations.
type UnimplementedAccountAttachmentsServer struct {
}

func (UnimplementedAccountAttachmentsServer) Create(context.Context, *AccountAttachmentCreateRequest) (*AccountAttachmentCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccountAttachmentsServer) Get(context.Context, *AccountAttachmentGetRequest) (*AccountAttachmentGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccountAttachmentsServer) Delete(context.Context, *AccountAttachmentDeleteRequest) (*AccountAttachmentDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAccountAttachmentsServer) List(context.Context, *AccountAttachmentListRequest) (*AccountAttachmentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAccountAttachmentsServer) mustEmbedUnimplementedAccountAttachmentsServer() {}

// UnsafeAccountAttachmentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountAttachmentsServer will
// result in compilation errors.
type UnsafeAccountAttachmentsServer interface {
	mustEmbedUnimplementedAccountAttachmentsServer()
}

func RegisterAccountAttachmentsServer(s grpc.ServiceRegistrar, srv AccountAttachmentsServer) {
	s.RegisterService(&_AccountAttachments_serviceDesc, srv)
}

func _AccountAttachments_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAttachmentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAttachmentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountAttachments/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAttachmentsServer).Create(ctx, req.(*AccountAttachmentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAttachments_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAttachmentGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAttachmentsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountAttachments/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAttachmentsServer).Get(ctx, req.(*AccountAttachmentGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAttachments_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAttachmentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAttachmentsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountAttachments/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAttachmentsServer).Delete(ctx, req.(*AccountAttachmentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAttachments_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAttachmentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAttachmentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountAttachments/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAttachmentsServer).List(ctx, req.(*AccountAttachmentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountAttachments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AccountAttachments",
	HandlerType: (*AccountAttachmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AccountAttachments_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AccountAttachments_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccountAttachments_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AccountAttachments_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_attachments.proto",
}