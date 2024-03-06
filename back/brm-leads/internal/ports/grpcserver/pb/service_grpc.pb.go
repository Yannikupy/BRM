// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LeadsServiceClient is the client API for LeadsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeadsServiceClient interface {
	CreateLead(ctx context.Context, in *CreateLeadRequest, opts ...grpc.CallOption) (*CreateLeadResponse, error)
	GetLeads(ctx context.Context, in *GetLeadsRequest, opts ...grpc.CallOption) (*GetLeadsResponse, error)
	GetLeadById(ctx context.Context, in *GetLeadByIdRequest, opts ...grpc.CallOption) (*GetLeadByIdResponse, error)
	UpdateLead(ctx context.Context, in *UpdateLeadRequest, opts ...grpc.CallOption) (*UpdateLeadResponse, error)
	GetStatuses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStatusesResponse, error)
	GetStatusById(ctx context.Context, in *GetStatusByIdRequest, opts ...grpc.CallOption) (*GetStatusByIdResponse, error)
}

type leadsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeadsServiceClient(cc grpc.ClientConnInterface) LeadsServiceClient {
	return &leadsServiceClient{cc}
}

func (c *leadsServiceClient) CreateLead(ctx context.Context, in *CreateLeadRequest, opts ...grpc.CallOption) (*CreateLeadResponse, error) {
	out := new(CreateLeadResponse)
	err := c.cc.Invoke(ctx, "/leads.LeadsService/CreateLead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leadsServiceClient) GetLeads(ctx context.Context, in *GetLeadsRequest, opts ...grpc.CallOption) (*GetLeadsResponse, error) {
	out := new(GetLeadsResponse)
	err := c.cc.Invoke(ctx, "/leads.LeadsService/GetLeads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leadsServiceClient) GetLeadById(ctx context.Context, in *GetLeadByIdRequest, opts ...grpc.CallOption) (*GetLeadByIdResponse, error) {
	out := new(GetLeadByIdResponse)
	err := c.cc.Invoke(ctx, "/leads.LeadsService/GetLeadById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leadsServiceClient) UpdateLead(ctx context.Context, in *UpdateLeadRequest, opts ...grpc.CallOption) (*UpdateLeadResponse, error) {
	out := new(UpdateLeadResponse)
	err := c.cc.Invoke(ctx, "/leads.LeadsService/UpdateLead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leadsServiceClient) GetStatuses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStatusesResponse, error) {
	out := new(GetStatusesResponse)
	err := c.cc.Invoke(ctx, "/leads.LeadsService/GetStatuses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leadsServiceClient) GetStatusById(ctx context.Context, in *GetStatusByIdRequest, opts ...grpc.CallOption) (*GetStatusByIdResponse, error) {
	out := new(GetStatusByIdResponse)
	err := c.cc.Invoke(ctx, "/leads.LeadsService/GetStatusById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeadsServiceServer is the server API for LeadsService service.
// All implementations should embed UnimplementedLeadsServiceServer
// for forward compatibility
type LeadsServiceServer interface {
	CreateLead(context.Context, *CreateLeadRequest) (*CreateLeadResponse, error)
	GetLeads(context.Context, *GetLeadsRequest) (*GetLeadsResponse, error)
	GetLeadById(context.Context, *GetLeadByIdRequest) (*GetLeadByIdResponse, error)
	UpdateLead(context.Context, *UpdateLeadRequest) (*UpdateLeadResponse, error)
	GetStatuses(context.Context, *emptypb.Empty) (*GetStatusesResponse, error)
	GetStatusById(context.Context, *GetStatusByIdRequest) (*GetStatusByIdResponse, error)
}

// UnimplementedLeadsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLeadsServiceServer struct {
}

func (UnimplementedLeadsServiceServer) CreateLead(context.Context, *CreateLeadRequest) (*CreateLeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLead not implemented")
}
func (UnimplementedLeadsServiceServer) GetLeads(context.Context, *GetLeadsRequest) (*GetLeadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeads not implemented")
}
func (UnimplementedLeadsServiceServer) GetLeadById(context.Context, *GetLeadByIdRequest) (*GetLeadByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeadById not implemented")
}
func (UnimplementedLeadsServiceServer) UpdateLead(context.Context, *UpdateLeadRequest) (*UpdateLeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLead not implemented")
}
func (UnimplementedLeadsServiceServer) GetStatuses(context.Context, *emptypb.Empty) (*GetStatusesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatuses not implemented")
}
func (UnimplementedLeadsServiceServer) GetStatusById(context.Context, *GetStatusByIdRequest) (*GetStatusByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatusById not implemented")
}

// UnsafeLeadsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeadsServiceServer will
// result in compilation errors.
type UnsafeLeadsServiceServer interface {
	mustEmbedUnimplementedLeadsServiceServer()
}

func RegisterLeadsServiceServer(s grpc.ServiceRegistrar, srv LeadsServiceServer) {
	s.RegisterService(&LeadsService_ServiceDesc, srv)
}

func _LeadsService_CreateLead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadsServiceServer).CreateLead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leads.LeadsService/CreateLead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadsServiceServer).CreateLead(ctx, req.(*CreateLeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeadsService_GetLeads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadsServiceServer).GetLeads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leads.LeadsService/GetLeads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadsServiceServer).GetLeads(ctx, req.(*GetLeadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeadsService_GetLeadById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeadByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadsServiceServer).GetLeadById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leads.LeadsService/GetLeadById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadsServiceServer).GetLeadById(ctx, req.(*GetLeadByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeadsService_UpdateLead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadsServiceServer).UpdateLead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leads.LeadsService/UpdateLead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadsServiceServer).UpdateLead(ctx, req.(*UpdateLeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeadsService_GetStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadsServiceServer).GetStatuses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leads.LeadsService/GetStatuses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadsServiceServer).GetStatuses(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeadsService_GetStatusById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadsServiceServer).GetStatusById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leads.LeadsService/GetStatusById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadsServiceServer).GetStatusById(ctx, req.(*GetStatusByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LeadsService_ServiceDesc is the grpc.ServiceDesc for LeadsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeadsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "leads.LeadsService",
	HandlerType: (*LeadsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLead",
			Handler:    _LeadsService_CreateLead_Handler,
		},
		{
			MethodName: "GetLeads",
			Handler:    _LeadsService_GetLeads_Handler,
		},
		{
			MethodName: "GetLeadById",
			Handler:    _LeadsService_GetLeadById_Handler,
		},
		{
			MethodName: "UpdateLead",
			Handler:    _LeadsService_UpdateLead_Handler,
		},
		{
			MethodName: "GetStatuses",
			Handler:    _LeadsService_GetStatuses_Handler,
		},
		{
			MethodName: "GetStatusById",
			Handler:    _LeadsService_GetStatusById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
