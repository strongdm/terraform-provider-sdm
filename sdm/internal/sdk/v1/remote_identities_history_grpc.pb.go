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

// RemoteIdentitiesHistoryClient is the client API for RemoteIdentitiesHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type RemoteIdentitiesHistoryClient interface {
	// List gets a list of RemoteIdentityHistory records matching a given set of criteria.
	List(ctx context.Context, in *RemoteIdentityHistoryListRequest, opts ...grpc.CallOption) (*RemoteIdentityHistoryListResponse, error)
}

type remoteIdentitiesHistoryClient struct {
	cc grpc.ClientConnInterface
}

// Deprecated: Do not use.
func NewRemoteIdentitiesHistoryClient(cc grpc.ClientConnInterface) RemoteIdentitiesHistoryClient {
	return &remoteIdentitiesHistoryClient{cc}
}

func (c *remoteIdentitiesHistoryClient) List(ctx context.Context, in *RemoteIdentityHistoryListRequest, opts ...grpc.CallOption) (*RemoteIdentityHistoryListResponse, error) {
	out := new(RemoteIdentityHistoryListResponse)
	err := c.cc.Invoke(ctx, "/v1.RemoteIdentitiesHistory/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteIdentitiesHistoryServer is the server API for RemoteIdentitiesHistory service.
// All implementations must embed UnimplementedRemoteIdentitiesHistoryServer
// for forward compatibility
//
// Deprecated: Do not use.
type RemoteIdentitiesHistoryServer interface {
	// List gets a list of RemoteIdentityHistory records matching a given set of criteria.
	List(context.Context, *RemoteIdentityHistoryListRequest) (*RemoteIdentityHistoryListResponse, error)
	mustEmbedUnimplementedRemoteIdentitiesHistoryServer()
}

// UnimplementedRemoteIdentitiesHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteIdentitiesHistoryServer struct {
}

func (UnimplementedRemoteIdentitiesHistoryServer) List(context.Context, *RemoteIdentityHistoryListRequest) (*RemoteIdentityHistoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRemoteIdentitiesHistoryServer) mustEmbedUnimplementedRemoteIdentitiesHistoryServer() {
}

// UnsafeRemoteIdentitiesHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteIdentitiesHistoryServer will
// result in compilation errors.
type UnsafeRemoteIdentitiesHistoryServer interface {
	mustEmbedUnimplementedRemoteIdentitiesHistoryServer()
}

// Deprecated: Do not use.
func RegisterRemoteIdentitiesHistoryServer(s grpc.ServiceRegistrar, srv RemoteIdentitiesHistoryServer) {
	s.RegisterService(&_RemoteIdentitiesHistory_serviceDesc, srv)
}

func _RemoteIdentitiesHistory_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteIdentityHistoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIdentitiesHistoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RemoteIdentitiesHistory/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIdentitiesHistoryServer).List(ctx, req.(*RemoteIdentityHistoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteIdentitiesHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RemoteIdentitiesHistory",
	HandlerType: (*RemoteIdentitiesHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _RemoteIdentitiesHistory_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote_identities_history.proto",
}
