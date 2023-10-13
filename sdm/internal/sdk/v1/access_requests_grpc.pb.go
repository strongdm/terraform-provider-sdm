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

// AccessRequestsClient is the client API for AccessRequests service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessRequestsClient interface {
	// Lists existing access requests.
	List(ctx context.Context, in *AccessRequestListRequest, opts ...grpc.CallOption) (*AccessRequestListResponse, error)
}

type accessRequestsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessRequestsClient(cc grpc.ClientConnInterface) AccessRequestsClient {
	return &accessRequestsClient{cc}
}

func (c *accessRequestsClient) List(ctx context.Context, in *AccessRequestListRequest, opts ...grpc.CallOption) (*AccessRequestListResponse, error) {
	out := new(AccessRequestListResponse)
	err := c.cc.Invoke(ctx, "/v1.AccessRequests/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessRequestsServer is the server API for AccessRequests service.
// All implementations must embed UnimplementedAccessRequestsServer
// for forward compatibility
type AccessRequestsServer interface {
	// Lists existing access requests.
	List(context.Context, *AccessRequestListRequest) (*AccessRequestListResponse, error)
	mustEmbedUnimplementedAccessRequestsServer()
}

// UnimplementedAccessRequestsServer must be embedded to have forward compatible implementations.
type UnimplementedAccessRequestsServer struct {
}

func (UnimplementedAccessRequestsServer) List(context.Context, *AccessRequestListRequest) (*AccessRequestListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAccessRequestsServer) mustEmbedUnimplementedAccessRequestsServer() {}

// UnsafeAccessRequestsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessRequestsServer will
// result in compilation errors.
type UnsafeAccessRequestsServer interface {
	mustEmbedUnimplementedAccessRequestsServer()
}

func RegisterAccessRequestsServer(s grpc.ServiceRegistrar, srv AccessRequestsServer) {
	s.RegisterService(&_AccessRequests_serviceDesc, srv)
}

func _AccessRequests_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRequestListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRequestsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccessRequests/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRequestsServer).List(ctx, req.(*AccessRequestListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessRequests_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AccessRequests",
	HandlerType: (*AccessRequestsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AccessRequests_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "access_requests.proto",
}
