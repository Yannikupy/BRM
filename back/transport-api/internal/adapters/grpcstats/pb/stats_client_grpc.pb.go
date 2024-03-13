// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: stats_client.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsServiceClient interface {
	GetCompanyMainPage(ctx context.Context, in *GetCompanyMainPageRequest, opts ...grpc.CallOption) (*GetCompanyMainPageResponse, error)
}

type statsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsServiceClient(cc grpc.ClientConnInterface) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) GetCompanyMainPage(ctx context.Context, in *GetCompanyMainPageRequest, opts ...grpc.CallOption) (*GetCompanyMainPageResponse, error) {
	out := new(GetCompanyMainPageResponse)
	err := c.cc.Invoke(ctx, "/stats.StatsService/GetCompanyMainPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServiceServer is the server API for StatsService service.
// All implementations should embed UnimplementedStatsServiceServer
// for forward compatibility
type StatsServiceServer interface {
	GetCompanyMainPage(context.Context, *GetCompanyMainPageRequest) (*GetCompanyMainPageResponse, error)
}

// UnimplementedStatsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStatsServiceServer struct {
}

func (UnimplementedStatsServiceServer) GetCompanyMainPage(context.Context, *GetCompanyMainPageRequest) (*GetCompanyMainPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyMainPage not implemented")
}

// UnsafeStatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServiceServer will
// result in compilation errors.
type UnsafeStatsServiceServer interface {
	mustEmbedUnimplementedStatsServiceServer()
}

func RegisterStatsServiceServer(s grpc.ServiceRegistrar, srv StatsServiceServer) {
	s.RegisterService(&StatsService_ServiceDesc, srv)
}

func _StatsService_GetCompanyMainPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyMainPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetCompanyMainPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.StatsService/GetCompanyMainPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetCompanyMainPage(ctx, req.(*GetCompanyMainPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatsService_ServiceDesc is the grpc.ServiceDesc for StatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stats.StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompanyMainPage",
			Handler:    _StatsService_GetCompanyMainPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stats_client.proto",
}
