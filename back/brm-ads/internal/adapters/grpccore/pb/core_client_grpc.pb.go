// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: core_client.proto

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

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreServiceClient interface {
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	GetEmployeeById(ctx context.Context, in *GetEmployeeByIdRequest, opts ...grpc.CallOption) (*GetEmployeeByIdResponse, error)
	GetIndustryId(ctx context.Context, in *GetIndustryIdRequest, opts ...grpc.CallOption) (*GetIndustryIdResponse, error)
}

type coreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreServiceClient(cc grpc.ClientConnInterface) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetEmployeeById(ctx context.Context, in *GetEmployeeByIdRequest, opts ...grpc.CallOption) (*GetEmployeeByIdResponse, error) {
	out := new(GetEmployeeByIdResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetEmployeeById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetIndustryId(ctx context.Context, in *GetIndustryIdRequest, opts ...grpc.CallOption) (*GetIndustryIdResponse, error) {
	out := new(GetIndustryIdResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetIndustryId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServiceServer is the server API for CoreService service.
// All implementations should embed UnimplementedCoreServiceServer
// for forward compatibility
type CoreServiceServer interface {
	GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error)
	GetEmployeeById(context.Context, *GetEmployeeByIdRequest) (*GetEmployeeByIdResponse, error)
	GetIndustryId(context.Context, *GetIndustryIdRequest) (*GetIndustryIdResponse, error)
}

// UnimplementedCoreServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCoreServiceServer struct {
}

func (UnimplementedCoreServiceServer) GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (UnimplementedCoreServiceServer) GetEmployeeById(context.Context, *GetEmployeeByIdRequest) (*GetEmployeeByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeById not implemented")
}
func (UnimplementedCoreServiceServer) GetIndustryId(context.Context, *GetIndustryIdRequest) (*GetIndustryIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndustryId not implemented")
}

// UnsafeCoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreServiceServer will
// result in compilation errors.
type UnsafeCoreServiceServer interface {
	mustEmbedUnimplementedCoreServiceServer()
}

func RegisterCoreServiceServer(s grpc.ServiceRegistrar, srv CoreServiceServer) {
	s.RegisterService(&CoreService_ServiceDesc, srv)
}

func _CoreService_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetEmployeeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeeByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetEmployeeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetEmployeeById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetEmployeeById(ctx, req.(*GetEmployeeByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetIndustryId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndustryIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetIndustryId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetIndustryId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetIndustryId(ctx, req.(*GetIndustryIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoreService_ServiceDesc is the grpc.ServiceDesc for CoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompany",
			Handler:    _CoreService_GetCompany_Handler,
		},
		{
			MethodName: "GetEmployeeById",
			Handler:    _CoreService_GetEmployeeById_Handler,
		},
		{
			MethodName: "GetIndustryId",
			Handler:    _CoreService_GetIndustryId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core_client.proto",
}
