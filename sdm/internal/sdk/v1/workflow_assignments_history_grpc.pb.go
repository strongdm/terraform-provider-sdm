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

// WorkflowAssignmentsHistoryClient is the client API for WorkflowAssignmentsHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowAssignmentsHistoryClient interface {
	// List gets a list of WorkflowAssignmentsHistory records matching a given set of criteria.
	List(ctx context.Context, in *WorkflowAssignmentsHistoryListRequest, opts ...grpc.CallOption) (*WorkflowAssignmentsHistoryListResponse, error)
}

type workflowAssignmentsHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowAssignmentsHistoryClient(cc grpc.ClientConnInterface) WorkflowAssignmentsHistoryClient {
	return &workflowAssignmentsHistoryClient{cc}
}

func (c *workflowAssignmentsHistoryClient) List(ctx context.Context, in *WorkflowAssignmentsHistoryListRequest, opts ...grpc.CallOption) (*WorkflowAssignmentsHistoryListResponse, error) {
	out := new(WorkflowAssignmentsHistoryListResponse)
	err := c.cc.Invoke(ctx, "/v1.WorkflowAssignmentsHistory/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowAssignmentsHistoryServer is the server API for WorkflowAssignmentsHistory service.
// All implementations must embed UnimplementedWorkflowAssignmentsHistoryServer
// for forward compatibility
type WorkflowAssignmentsHistoryServer interface {
	// List gets a list of WorkflowAssignmentsHistory records matching a given set of criteria.
	List(context.Context, *WorkflowAssignmentsHistoryListRequest) (*WorkflowAssignmentsHistoryListResponse, error)
	mustEmbedUnimplementedWorkflowAssignmentsHistoryServer()
}

// UnimplementedWorkflowAssignmentsHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowAssignmentsHistoryServer struct {
}

func (UnimplementedWorkflowAssignmentsHistoryServer) List(context.Context, *WorkflowAssignmentsHistoryListRequest) (*WorkflowAssignmentsHistoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedWorkflowAssignmentsHistoryServer) mustEmbedUnimplementedWorkflowAssignmentsHistoryServer() {
}

// UnsafeWorkflowAssignmentsHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowAssignmentsHistoryServer will
// result in compilation errors.
type UnsafeWorkflowAssignmentsHistoryServer interface {
	mustEmbedUnimplementedWorkflowAssignmentsHistoryServer()
}

func RegisterWorkflowAssignmentsHistoryServer(s grpc.ServiceRegistrar, srv WorkflowAssignmentsHistoryServer) {
	s.RegisterService(&_WorkflowAssignmentsHistory_serviceDesc, srv)
}

func _WorkflowAssignmentsHistory_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowAssignmentsHistoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAssignmentsHistoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.WorkflowAssignmentsHistory/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAssignmentsHistoryServer).List(ctx, req.(*WorkflowAssignmentsHistoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowAssignmentsHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.WorkflowAssignmentsHistory",
	HandlerType: (*WorkflowAssignmentsHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _WorkflowAssignmentsHistory_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow_assignments_history.proto",
}