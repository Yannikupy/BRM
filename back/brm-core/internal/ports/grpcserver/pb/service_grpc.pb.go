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

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreServiceClient interface {
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	CreateCompanyAndOwner(ctx context.Context, in *CreateCompanyAndOwnerRequest, opts ...grpc.CallOption) (*CreateCompanyAndOwnerResponse, error)
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error)
	DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetIndustriesList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetIndustriesListResponse, error)
	GetIndustryById(ctx context.Context, in *GetIndustryByIdRequest, opts ...grpc.CallOption) (*GetIndustryByIdResponse, error)
	CreateEmployee(ctx context.Context, in *CreateEmployeeRequest, opts ...grpc.CallOption) (*CreateEmployeeResponse, error)
	UpdateEmployee(ctx context.Context, in *UpdateEmployeeRequest, opts ...grpc.CallOption) (*UpdateEmployeeResponse, error)
	DeleteEmployee(ctx context.Context, in *DeleteEmployeeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCompanyEmployees(ctx context.Context, in *GetCompanyEmployeesRequest, opts ...grpc.CallOption) (*GetCompanyEmployeesResponse, error)
	GetEmployeeByName(ctx context.Context, in *GetEmployeeByNameRequest, opts ...grpc.CallOption) (*GetEmployeeByNameResponse, error)
	GetEmployeeById(ctx context.Context, in *GetEmployeeByIdRequest, opts ...grpc.CallOption) (*GetEmployeeByIdResponse, error)
	CreateContact(ctx context.Context, in *CreateContactRequest, opts ...grpc.CallOption) (*CreateContactResponse, error)
	UpdateContact(ctx context.Context, in *UpdateContactRequest, opts ...grpc.CallOption) (*UpdateContactResponse, error)
	DeleteContact(ctx context.Context, in *DeleteContactRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetContacts(ctx context.Context, in *GetContactsRequest, opts ...grpc.CallOption) (*GetContactsResponse, error)
	GetContactById(ctx context.Context, in *GetContactByIdRequest, opts ...grpc.CallOption) (*GetContactByIdResponse, error)
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

func (c *coreServiceClient) CreateCompanyAndOwner(ctx context.Context, in *CreateCompanyAndOwnerRequest, opts ...grpc.CallOption) (*CreateCompanyAndOwnerResponse, error) {
	out := new(CreateCompanyAndOwnerResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/CreateCompanyAndOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error) {
	out := new(UpdateCompanyResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/core.CoreService/DeleteCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetIndustriesList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetIndustriesListResponse, error) {
	out := new(GetIndustriesListResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetIndustriesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetIndustryById(ctx context.Context, in *GetIndustryByIdRequest, opts ...grpc.CallOption) (*GetIndustryByIdResponse, error) {
	out := new(GetIndustryByIdResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetIndustryById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) CreateEmployee(ctx context.Context, in *CreateEmployeeRequest, opts ...grpc.CallOption) (*CreateEmployeeResponse, error) {
	out := new(CreateEmployeeResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/CreateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) UpdateEmployee(ctx context.Context, in *UpdateEmployeeRequest, opts ...grpc.CallOption) (*UpdateEmployeeResponse, error) {
	out := new(UpdateEmployeeResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/UpdateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteEmployee(ctx context.Context, in *DeleteEmployeeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/core.CoreService/DeleteEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetCompanyEmployees(ctx context.Context, in *GetCompanyEmployeesRequest, opts ...grpc.CallOption) (*GetCompanyEmployeesResponse, error) {
	out := new(GetCompanyEmployeesResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetCompanyEmployees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetEmployeeByName(ctx context.Context, in *GetEmployeeByNameRequest, opts ...grpc.CallOption) (*GetEmployeeByNameResponse, error) {
	out := new(GetEmployeeByNameResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetEmployeeByName", in, out, opts...)
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

func (c *coreServiceClient) CreateContact(ctx context.Context, in *CreateContactRequest, opts ...grpc.CallOption) (*CreateContactResponse, error) {
	out := new(CreateContactResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/CreateContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) UpdateContact(ctx context.Context, in *UpdateContactRequest, opts ...grpc.CallOption) (*UpdateContactResponse, error) {
	out := new(UpdateContactResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/UpdateContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteContact(ctx context.Context, in *DeleteContactRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/core.CoreService/DeleteContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetContacts(ctx context.Context, in *GetContactsRequest, opts ...grpc.CallOption) (*GetContactsResponse, error) {
	out := new(GetContactsResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetContactById(ctx context.Context, in *GetContactByIdRequest, opts ...grpc.CallOption) (*GetContactByIdResponse, error) {
	out := new(GetContactByIdResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetContactById", in, out, opts...)
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
	CreateCompanyAndOwner(context.Context, *CreateCompanyAndOwnerRequest) (*CreateCompanyAndOwnerResponse, error)
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error)
	DeleteCompany(context.Context, *DeleteCompanyRequest) (*emptypb.Empty, error)
	GetIndustriesList(context.Context, *emptypb.Empty) (*GetIndustriesListResponse, error)
	GetIndustryById(context.Context, *GetIndustryByIdRequest) (*GetIndustryByIdResponse, error)
	CreateEmployee(context.Context, *CreateEmployeeRequest) (*CreateEmployeeResponse, error)
	UpdateEmployee(context.Context, *UpdateEmployeeRequest) (*UpdateEmployeeResponse, error)
	DeleteEmployee(context.Context, *DeleteEmployeeRequest) (*emptypb.Empty, error)
	GetCompanyEmployees(context.Context, *GetCompanyEmployeesRequest) (*GetCompanyEmployeesResponse, error)
	GetEmployeeByName(context.Context, *GetEmployeeByNameRequest) (*GetEmployeeByNameResponse, error)
	GetEmployeeById(context.Context, *GetEmployeeByIdRequest) (*GetEmployeeByIdResponse, error)
	CreateContact(context.Context, *CreateContactRequest) (*CreateContactResponse, error)
	UpdateContact(context.Context, *UpdateContactRequest) (*UpdateContactResponse, error)
	DeleteContact(context.Context, *DeleteContactRequest) (*emptypb.Empty, error)
	GetContacts(context.Context, *GetContactsRequest) (*GetContactsResponse, error)
	GetContactById(context.Context, *GetContactByIdRequest) (*GetContactByIdResponse, error)
}

// UnimplementedCoreServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCoreServiceServer struct {
}

func (UnimplementedCoreServiceServer) GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (UnimplementedCoreServiceServer) CreateCompanyAndOwner(context.Context, *CreateCompanyAndOwnerRequest) (*CreateCompanyAndOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompanyAndOwner not implemented")
}
func (UnimplementedCoreServiceServer) UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedCoreServiceServer) DeleteCompany(context.Context, *DeleteCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (UnimplementedCoreServiceServer) GetIndustriesList(context.Context, *emptypb.Empty) (*GetIndustriesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndustriesList not implemented")
}
func (UnimplementedCoreServiceServer) GetIndustryById(context.Context, *GetIndustryByIdRequest) (*GetIndustryByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndustryById not implemented")
}
func (UnimplementedCoreServiceServer) CreateEmployee(context.Context, *CreateEmployeeRequest) (*CreateEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmployee not implemented")
}
func (UnimplementedCoreServiceServer) UpdateEmployee(context.Context, *UpdateEmployeeRequest) (*UpdateEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (UnimplementedCoreServiceServer) DeleteEmployee(context.Context, *DeleteEmployeeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (UnimplementedCoreServiceServer) GetCompanyEmployees(context.Context, *GetCompanyEmployeesRequest) (*GetCompanyEmployeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyEmployees not implemented")
}
func (UnimplementedCoreServiceServer) GetEmployeeByName(context.Context, *GetEmployeeByNameRequest) (*GetEmployeeByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeByName not implemented")
}
func (UnimplementedCoreServiceServer) GetEmployeeById(context.Context, *GetEmployeeByIdRequest) (*GetEmployeeByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeById not implemented")
}
func (UnimplementedCoreServiceServer) CreateContact(context.Context, *CreateContactRequest) (*CreateContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContact not implemented")
}
func (UnimplementedCoreServiceServer) UpdateContact(context.Context, *UpdateContactRequest) (*UpdateContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContact not implemented")
}
func (UnimplementedCoreServiceServer) DeleteContact(context.Context, *DeleteContactRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContact not implemented")
}
func (UnimplementedCoreServiceServer) GetContacts(context.Context, *GetContactsRequest) (*GetContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContacts not implemented")
}
func (UnimplementedCoreServiceServer) GetContactById(context.Context, *GetContactByIdRequest) (*GetContactByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactById not implemented")
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

func _CoreService_CreateCompanyAndOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyAndOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CreateCompanyAndOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/CreateCompanyAndOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CreateCompanyAndOwner(ctx, req.(*CreateCompanyAndOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/DeleteCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteCompany(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetIndustriesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetIndustriesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetIndustriesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetIndustriesList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetIndustryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndustryByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetIndustryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetIndustryById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetIndustryById(ctx, req.(*GetIndustryByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_CreateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CreateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/CreateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CreateEmployee(ctx, req.(*CreateEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/UpdateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).UpdateEmployee(ctx, req.(*UpdateEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/DeleteEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteEmployee(ctx, req.(*DeleteEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetCompanyEmployees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyEmployeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetCompanyEmployees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetCompanyEmployees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetCompanyEmployees(ctx, req.(*GetCompanyEmployeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetEmployeeByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeeByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetEmployeeByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetEmployeeByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetEmployeeByName(ctx, req.(*GetEmployeeByNameRequest))
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

func _CoreService_CreateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CreateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/CreateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CreateContact(ctx, req.(*CreateContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_UpdateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).UpdateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/UpdateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).UpdateContact(ctx, req.(*UpdateContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/DeleteContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteContact(ctx, req.(*DeleteContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetContacts(ctx, req.(*GetContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetContactById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetContactById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetContactById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetContactById(ctx, req.(*GetContactByIdRequest))
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
			MethodName: "CreateCompanyAndOwner",
			Handler:    _CoreService_CreateCompanyAndOwner_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _CoreService_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _CoreService_DeleteCompany_Handler,
		},
		{
			MethodName: "GetIndustriesList",
			Handler:    _CoreService_GetIndustriesList_Handler,
		},
		{
			MethodName: "GetIndustryById",
			Handler:    _CoreService_GetIndustryById_Handler,
		},
		{
			MethodName: "CreateEmployee",
			Handler:    _CoreService_CreateEmployee_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _CoreService_UpdateEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _CoreService_DeleteEmployee_Handler,
		},
		{
			MethodName: "GetCompanyEmployees",
			Handler:    _CoreService_GetCompanyEmployees_Handler,
		},
		{
			MethodName: "GetEmployeeByName",
			Handler:    _CoreService_GetEmployeeByName_Handler,
		},
		{
			MethodName: "GetEmployeeById",
			Handler:    _CoreService_GetEmployeeById_Handler,
		},
		{
			MethodName: "CreateContact",
			Handler:    _CoreService_CreateContact_Handler,
		},
		{
			MethodName: "UpdateContact",
			Handler:    _CoreService_UpdateContact_Handler,
		},
		{
			MethodName: "DeleteContact",
			Handler:    _CoreService_DeleteContact_Handler,
		},
		{
			MethodName: "GetContacts",
			Handler:    _CoreService_GetContacts_Handler,
		},
		{
			MethodName: "GetContactById",
			Handler:    _CoreService_GetContactById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
