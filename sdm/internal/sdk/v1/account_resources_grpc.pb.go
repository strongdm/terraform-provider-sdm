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

// AccountResourcesClient is the client API for AccountResources service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountResourcesClient interface {
	// List gets a list of AccountResource records matching a given set of criteria.
	List(ctx context.Context, in *AccountResourceListRequest, opts ...grpc.CallOption) (*AccountResourceListResponse, error)
}

type accountResourcesClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountResourcesClient(cc grpc.ClientConnInterface) AccountResourcesClient {
	return &accountResourcesClient{cc}
}

func (c *accountResourcesClient) List(ctx context.Context, in *AccountResourceListRequest, opts ...grpc.CallOption) (*AccountResourceListResponse, error) {
	out := new(AccountResourceListResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountResources/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountResourcesServer is the server API for AccountResources service.
// All implementations must embed UnimplementedAccountResourcesServer
// for forward compatibility
type AccountResourcesServer interface {
	// List gets a list of AccountResource records matching a given set of criteria.
	List(context.Context, *AccountResourceListRequest) (*AccountResourceListResponse, error)
	mustEmbedUnimplementedAccountResourcesServer()
}

// UnimplementedAccountResourcesServer must be embedded to have forward compatible implementations.
type UnimplementedAccountResourcesServer struct {
}

func (UnimplementedAccountResourcesServer) List(context.Context, *AccountResourceListRequest) (*AccountResourceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAccountResourcesServer) mustEmbedUnimplementedAccountResourcesServer() {}

// UnsafeAccountResourcesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountResourcesServer will
// result in compilation errors.
type UnsafeAccountResourcesServer interface {
	mustEmbedUnimplementedAccountResourcesServer()
}

func RegisterAccountResourcesServer(s grpc.ServiceRegistrar, srv AccountResourcesServer) {
	s.RegisterService(&_AccountResources_serviceDesc, srv)
}

func _AccountResources_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountResourceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountResourcesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountResources/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountResourcesServer).List(ctx, req.(*AccountResourceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountResources_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AccountResources",
	HandlerType: (*AccountResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AccountResources_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_resources.proto",
}
