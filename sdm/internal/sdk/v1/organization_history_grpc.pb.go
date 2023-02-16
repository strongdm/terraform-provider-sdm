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

// OrganizationHistoryClient is the client API for OrganizationHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationHistoryClient interface {
	// List gets a list of OrganizationHistory records matching a given set of criteria.
	List(ctx context.Context, in *OrganizationHistoryListRequest, opts ...grpc.CallOption) (*OrganizationHistoryListResponse, error)
}

type organizationHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationHistoryClient(cc grpc.ClientConnInterface) OrganizationHistoryClient {
	return &organizationHistoryClient{cc}
}

func (c *organizationHistoryClient) List(ctx context.Context, in *OrganizationHistoryListRequest, opts ...grpc.CallOption) (*OrganizationHistoryListResponse, error) {
	out := new(OrganizationHistoryListResponse)
	err := c.cc.Invoke(ctx, "/v1.OrganizationHistory/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationHistoryServer is the server API for OrganizationHistory service.
// All implementations must embed UnimplementedOrganizationHistoryServer
// for forward compatibility
type OrganizationHistoryServer interface {
	// List gets a list of OrganizationHistory records matching a given set of criteria.
	List(context.Context, *OrganizationHistoryListRequest) (*OrganizationHistoryListResponse, error)
	mustEmbedUnimplementedOrganizationHistoryServer()
}

// UnimplementedOrganizationHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationHistoryServer struct {
}

func (UnimplementedOrganizationHistoryServer) List(context.Context, *OrganizationHistoryListRequest) (*OrganizationHistoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrganizationHistoryServer) mustEmbedUnimplementedOrganizationHistoryServer() {}

// UnsafeOrganizationHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationHistoryServer will
// result in compilation errors.
type UnsafeOrganizationHistoryServer interface {
	mustEmbedUnimplementedOrganizationHistoryServer()
}

func RegisterOrganizationHistoryServer(s grpc.ServiceRegistrar, srv OrganizationHistoryServer) {
	s.RegisterService(&_OrganizationHistory_serviceDesc, srv)
}

func _OrganizationHistory_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationHistoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationHistoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OrganizationHistory/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationHistoryServer).List(ctx, req.(*OrganizationHistoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrganizationHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.OrganizationHistory",
	HandlerType: (*OrganizationHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _OrganizationHistory_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "organization_history.proto",
}
