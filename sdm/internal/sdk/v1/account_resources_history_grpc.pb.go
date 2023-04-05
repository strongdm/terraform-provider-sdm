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

// AccountResourcesHistoryClient is the client API for AccountResourcesHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountResourcesHistoryClient interface {
	// List gets a list of AccountResourceHistory records matching a given set of criteria.
	List(ctx context.Context, in *AccountResourceHistoryListRequest, opts ...grpc.CallOption) (*AccountResourceHistoryListResponse, error)
}

type accountResourcesHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountResourcesHistoryClient(cc grpc.ClientConnInterface) AccountResourcesHistoryClient {
	return &accountResourcesHistoryClient{cc}
}

func (c *accountResourcesHistoryClient) List(ctx context.Context, in *AccountResourceHistoryListRequest, opts ...grpc.CallOption) (*AccountResourceHistoryListResponse, error) {
	out := new(AccountResourceHistoryListResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountResourcesHistory/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountResourcesHistoryServer is the server API for AccountResourcesHistory service.
// All implementations must embed UnimplementedAccountResourcesHistoryServer
// for forward compatibility
type AccountResourcesHistoryServer interface {
	// List gets a list of AccountResourceHistory records matching a given set of criteria.
	List(context.Context, *AccountResourceHistoryListRequest) (*AccountResourceHistoryListResponse, error)
	mustEmbedUnimplementedAccountResourcesHistoryServer()
}

// UnimplementedAccountResourcesHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedAccountResourcesHistoryServer struct {
}

func (UnimplementedAccountResourcesHistoryServer) List(context.Context, *AccountResourceHistoryListRequest) (*AccountResourceHistoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAccountResourcesHistoryServer) mustEmbedUnimplementedAccountResourcesHistoryServer() {
}

// UnsafeAccountResourcesHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountResourcesHistoryServer will
// result in compilation errors.
type UnsafeAccountResourcesHistoryServer interface {
	mustEmbedUnimplementedAccountResourcesHistoryServer()
}

func RegisterAccountResourcesHistoryServer(s grpc.ServiceRegistrar, srv AccountResourcesHistoryServer) {
	s.RegisterService(&_AccountResourcesHistory_serviceDesc, srv)
}

func _AccountResourcesHistory_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountResourceHistoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountResourcesHistoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountResourcesHistory/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountResourcesHistoryServer).List(ctx, req.(*AccountResourceHistoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountResourcesHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AccountResourcesHistory",
	HandlerType: (*AccountResourcesHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AccountResourcesHistory_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_resources_history.proto",
}
