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

// RolesClient is the client API for Roles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolesClient interface {
	// Create registers a new Role.
	Create(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error)
	// Get reads one Role by ID.
	Get(ctx context.Context, in *RoleGetRequest, opts ...grpc.CallOption) (*RoleGetResponse, error)
	// Update replaces all the fields of a Role by ID.
	Update(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleUpdateResponse, error)
	// Delete removes a Role by ID.
	Delete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error)
	// List gets a list of Roles matching a given set of criteria.
	List(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListResponse, error)
}

type rolesClient struct {
	cc grpc.ClientConnInterface
}

func NewRolesClient(cc grpc.ClientConnInterface) RolesClient {
	return &rolesClient{cc}
}

func (c *rolesClient) Create(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error) {
	out := new(RoleCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.Roles/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Get(ctx context.Context, in *RoleGetRequest, opts ...grpc.CallOption) (*RoleGetResponse, error) {
	out := new(RoleGetResponse)
	err := c.cc.Invoke(ctx, "/v1.Roles/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Update(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleUpdateResponse, error) {
	out := new(RoleUpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.Roles/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Delete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error) {
	out := new(RoleDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.Roles/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) List(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListResponse, error) {
	out := new(RoleListResponse)
	err := c.cc.Invoke(ctx, "/v1.Roles/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServer is the server API for Roles service.
// All implementations must embed UnimplementedRolesServer
// for forward compatibility
type RolesServer interface {
	// Create registers a new Role.
	Create(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error)
	// Get reads one Role by ID.
	Get(context.Context, *RoleGetRequest) (*RoleGetResponse, error)
	// Update replaces all the fields of a Role by ID.
	Update(context.Context, *RoleUpdateRequest) (*RoleUpdateResponse, error)
	// Delete removes a Role by ID.
	Delete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error)
	// List gets a list of Roles matching a given set of criteria.
	List(context.Context, *RoleListRequest) (*RoleListResponse, error)
	mustEmbedUnimplementedRolesServer()
}

// UnimplementedRolesServer must be embedded to have forward compatible implementations.
type UnimplementedRolesServer struct {
}

func (UnimplementedRolesServer) Create(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRolesServer) Get(context.Context, *RoleGetRequest) (*RoleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRolesServer) Update(context.Context, *RoleUpdateRequest) (*RoleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRolesServer) Delete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRolesServer) List(context.Context, *RoleListRequest) (*RoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRolesServer) mustEmbedUnimplementedRolesServer() {}

// UnsafeRolesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolesServer will
// result in compilation errors.
type UnsafeRolesServer interface {
	mustEmbedUnimplementedRolesServer()
}

func RegisterRolesServer(s grpc.ServiceRegistrar, srv RolesServer) {
	s.RegisterService(&_Roles_serviceDesc, srv)
}

func _Roles_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Roles/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Create(ctx, req.(*RoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Roles/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Get(ctx, req.(*RoleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Roles/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Update(ctx, req.(*RoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Roles/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Delete(ctx, req.(*RoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Roles/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).List(ctx, req.(*RoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Roles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Roles",
	HandlerType: (*RolesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Roles_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Roles_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Roles_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Roles_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Roles_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "roles.proto",
}