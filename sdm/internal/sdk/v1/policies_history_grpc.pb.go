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

// PoliciesHistoryClient is the client API for PoliciesHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoliciesHistoryClient interface {
	// List gets a list of PolicyHistory records matching a given set of criteria.
	List(ctx context.Context, in *PoliciesHistoryListRequest, opts ...grpc.CallOption) (*PoliciesHistoryListResponse, error)
}

type policiesHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewPoliciesHistoryClient(cc grpc.ClientConnInterface) PoliciesHistoryClient {
	return &policiesHistoryClient{cc}
}

func (c *policiesHistoryClient) List(ctx context.Context, in *PoliciesHistoryListRequest, opts ...grpc.CallOption) (*PoliciesHistoryListResponse, error) {
	out := new(PoliciesHistoryListResponse)
	err := c.cc.Invoke(ctx, "/v1.PoliciesHistory/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoliciesHistoryServer is the server API for PoliciesHistory service.
// All implementations must embed UnimplementedPoliciesHistoryServer
// for forward compatibility
type PoliciesHistoryServer interface {
	// List gets a list of PolicyHistory records matching a given set of criteria.
	List(context.Context, *PoliciesHistoryListRequest) (*PoliciesHistoryListResponse, error)
	mustEmbedUnimplementedPoliciesHistoryServer()
}

// UnimplementedPoliciesHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedPoliciesHistoryServer struct {
}

func (UnimplementedPoliciesHistoryServer) List(context.Context, *PoliciesHistoryListRequest) (*PoliciesHistoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPoliciesHistoryServer) mustEmbedUnimplementedPoliciesHistoryServer() {}

// UnsafePoliciesHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoliciesHistoryServer will
// result in compilation errors.
type UnsafePoliciesHistoryServer interface {
	mustEmbedUnimplementedPoliciesHistoryServer()
}

func RegisterPoliciesHistoryServer(s grpc.ServiceRegistrar, srv PoliciesHistoryServer) {
	s.RegisterService(&_PoliciesHistory_serviceDesc, srv)
}

func _PoliciesHistory_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PoliciesHistoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesHistoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PoliciesHistory/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesHistoryServer).List(ctx, req.(*PoliciesHistoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PoliciesHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PoliciesHistory",
	HandlerType: (*PoliciesHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PoliciesHistory_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policies_history.proto",
}
