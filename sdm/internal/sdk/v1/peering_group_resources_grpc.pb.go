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

// PeeringGroupResourcesClient is the client API for PeeringGroupResources service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeeringGroupResourcesClient interface {
	// Create attaches a Resource to a PeeringGroup
	Create(ctx context.Context, in *PeeringGroupResourceCreateRequest, opts ...grpc.CallOption) (*PeeringGroupResourceCreateResponse, error)
	// Delete detaches a Resource to a PeeringGroup
	Delete(ctx context.Context, in *PeeringGroupResourceDeleteRequest, opts ...grpc.CallOption) (*PeeringGroupResourceDeleteResponse, error)
	// Get reads the information of one peering group to resource attachment.
	Get(ctx context.Context, in *PeeringGroupResourceGetRequest, opts ...grpc.CallOption) (*PeeringGroupResourceGetResponse, error)
	// List gets a list of peering group resource attachments.
	List(ctx context.Context, in *PeeringGroupResourceListRequest, opts ...grpc.CallOption) (*PeeringGroupResourceListResponse, error)
}

type peeringGroupResourcesClient struct {
	cc grpc.ClientConnInterface
}

func NewPeeringGroupResourcesClient(cc grpc.ClientConnInterface) PeeringGroupResourcesClient {
	return &peeringGroupResourcesClient{cc}
}

func (c *peeringGroupResourcesClient) Create(ctx context.Context, in *PeeringGroupResourceCreateRequest, opts ...grpc.CallOption) (*PeeringGroupResourceCreateResponse, error) {
	out := new(PeeringGroupResourceCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.PeeringGroupResources/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringGroupResourcesClient) Delete(ctx context.Context, in *PeeringGroupResourceDeleteRequest, opts ...grpc.CallOption) (*PeeringGroupResourceDeleteResponse, error) {
	out := new(PeeringGroupResourceDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.PeeringGroupResources/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringGroupResourcesClient) Get(ctx context.Context, in *PeeringGroupResourceGetRequest, opts ...grpc.CallOption) (*PeeringGroupResourceGetResponse, error) {
	out := new(PeeringGroupResourceGetResponse)
	err := c.cc.Invoke(ctx, "/v1.PeeringGroupResources/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringGroupResourcesClient) List(ctx context.Context, in *PeeringGroupResourceListRequest, opts ...grpc.CallOption) (*PeeringGroupResourceListResponse, error) {
	out := new(PeeringGroupResourceListResponse)
	err := c.cc.Invoke(ctx, "/v1.PeeringGroupResources/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeeringGroupResourcesServer is the server API for PeeringGroupResources service.
// All implementations must embed UnimplementedPeeringGroupResourcesServer
// for forward compatibility
type PeeringGroupResourcesServer interface {
	// Create attaches a Resource to a PeeringGroup
	Create(context.Context, *PeeringGroupResourceCreateRequest) (*PeeringGroupResourceCreateResponse, error)
	// Delete detaches a Resource to a PeeringGroup
	Delete(context.Context, *PeeringGroupResourceDeleteRequest) (*PeeringGroupResourceDeleteResponse, error)
	// Get reads the information of one peering group to resource attachment.
	Get(context.Context, *PeeringGroupResourceGetRequest) (*PeeringGroupResourceGetResponse, error)
	// List gets a list of peering group resource attachments.
	List(context.Context, *PeeringGroupResourceListRequest) (*PeeringGroupResourceListResponse, error)
	mustEmbedUnimplementedPeeringGroupResourcesServer()
}

// UnimplementedPeeringGroupResourcesServer must be embedded to have forward compatible implementations.
type UnimplementedPeeringGroupResourcesServer struct {
}

func (UnimplementedPeeringGroupResourcesServer) Create(context.Context, *PeeringGroupResourceCreateRequest) (*PeeringGroupResourceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPeeringGroupResourcesServer) Delete(context.Context, *PeeringGroupResourceDeleteRequest) (*PeeringGroupResourceDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPeeringGroupResourcesServer) Get(context.Context, *PeeringGroupResourceGetRequest) (*PeeringGroupResourceGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPeeringGroupResourcesServer) List(context.Context, *PeeringGroupResourceListRequest) (*PeeringGroupResourceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPeeringGroupResourcesServer) mustEmbedUnimplementedPeeringGroupResourcesServer() {}

// UnsafePeeringGroupResourcesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeeringGroupResourcesServer will
// result in compilation errors.
type UnsafePeeringGroupResourcesServer interface {
	mustEmbedUnimplementedPeeringGroupResourcesServer()
}

func RegisterPeeringGroupResourcesServer(s grpc.ServiceRegistrar, srv PeeringGroupResourcesServer) {
	s.RegisterService(&_PeeringGroupResources_serviceDesc, srv)
}

func _PeeringGroupResources_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringGroupResourceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringGroupResourcesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PeeringGroupResources/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringGroupResourcesServer).Create(ctx, req.(*PeeringGroupResourceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringGroupResources_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringGroupResourceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringGroupResourcesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PeeringGroupResources/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringGroupResourcesServer).Delete(ctx, req.(*PeeringGroupResourceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringGroupResources_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringGroupResourceGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringGroupResourcesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PeeringGroupResources/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringGroupResourcesServer).Get(ctx, req.(*PeeringGroupResourceGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringGroupResources_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringGroupResourceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringGroupResourcesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PeeringGroupResources/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringGroupResourcesServer).List(ctx, req.(*PeeringGroupResourceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeeringGroupResources_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PeeringGroupResources",
	HandlerType: (*PeeringGroupResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PeeringGroupResources_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PeeringGroupResources_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PeeringGroupResources_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PeeringGroupResources_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peering_group_resources.proto",
}