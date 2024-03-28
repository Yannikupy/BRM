// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActiveLeadsAmount     uint64  `protobuf:"varint,1,opt,name=active_leads_amount,json=activeLeadsAmount,proto3" json:"active_leads_amount,omitempty"`
	ActiveLeadsPrice      uint64  `protobuf:"varint,2,opt,name=active_leads_price,json=activeLeadsPrice,proto3" json:"active_leads_price,omitempty"`
	ClosedLeadsAmount     uint64  `protobuf:"varint,3,opt,name=closed_leads_amount,json=closedLeadsAmount,proto3" json:"closed_leads_amount,omitempty"`
	ClosedLeadsPrice      uint64  `protobuf:"varint,4,opt,name=closed_leads_price,json=closedLeadsPrice,proto3" json:"closed_leads_price,omitempty"`
	ActiveAdsAmount       uint64  `protobuf:"varint,5,opt,name=active_ads_amount,json=activeAdsAmount,proto3" json:"active_ads_amount,omitempty"`
	CompanyAbsoluteRating float64 `protobuf:"fixed64,6,opt,name=company_absolute_rating,json=companyAbsoluteRating,proto3" json:"company_absolute_rating,omitempty"`
	CompanyRelativeRating float64 `protobuf:"fixed64,7,opt,name=company_relative_rating,json=companyRelativeRating,proto3" json:"company_relative_rating,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetActiveLeadsAmount() uint64 {
	if x != nil {
		return x.ActiveLeadsAmount
	}
	return 0
}

func (x *Data) GetActiveLeadsPrice() uint64 {
	if x != nil {
		return x.ActiveLeadsPrice
	}
	return 0
}

func (x *Data) GetClosedLeadsAmount() uint64 {
	if x != nil {
		return x.ClosedLeadsAmount
	}
	return 0
}

func (x *Data) GetClosedLeadsPrice() uint64 {
	if x != nil {
		return x.ClosedLeadsPrice
	}
	return 0
}

func (x *Data) GetActiveAdsAmount() uint64 {
	if x != nil {
		return x.ActiveAdsAmount
	}
	return 0
}

func (x *Data) GetCompanyAbsoluteRating() float64 {
	if x != nil {
		return x.CompanyAbsoluteRating
	}
	return 0
}

func (x *Data) GetCompanyRelativeRating() float64 {
	if x != nil {
		return x.CompanyRelativeRating
	}
	return 0
}

type GetCompanyMainPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId uint64 `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *GetCompanyMainPageRequest) Reset() {
	*x = GetCompanyMainPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyMainPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyMainPageRequest) ProtoMessage() {}

func (x *GetCompanyMainPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyMainPageRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyMainPageRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCompanyMainPageRequest) GetCompanyId() uint64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type GetCompanyMainPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCompanyMainPageResponse) Reset() {
	*x = GetCompanyMainPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyMainPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyMainPageResponse) ProtoMessage() {}

func (x *GetCompanyMainPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyMainPageResponse.ProtoReflect.Descriptor instead.
func (*GetCompanyMainPageResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCompanyMainPageResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type SubmitClosedLeadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId uint64 `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Submit    bool   `protobuf:"varint,2,opt,name=submit,proto3" json:"submit,omitempty"`
}

func (x *SubmitClosedLeadRequest) Reset() {
	*x = SubmitClosedLeadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitClosedLeadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitClosedLeadRequest) ProtoMessage() {}

func (x *SubmitClosedLeadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitClosedLeadRequest.ProtoReflect.Descriptor instead.
func (*SubmitClosedLeadRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitClosedLeadRequest) GetCompanyId() uint64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *SubmitClosedLeadRequest) GetSubmit() bool {
	if x != nil {
		return x.Submit
	}
	return false
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xde, 0x02, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x13,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x4c, 0x65, 0x61, 0x64, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x4c, 0x65, 0x61, 0x64, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x4c,
	0x65, 0x61, 0x64, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x4c, 0x65,
	0x61, 0x64, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x64, 0x73, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x41, 0x62,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x22, 0x3a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x22, 0x3d, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x61,
	0x69, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x50, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x4c,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x32, 0xb9, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x4d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x61, 0x69,
	0x6e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4c, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x4c,
	0x65, 0x61, 0x64, 0x12, 0x1e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x05, 0x5a,
	0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_proto_goTypes = []interface{}{
	(*Data)(nil),                       // 0: stats.Data
	(*GetCompanyMainPageRequest)(nil),  // 1: stats.GetCompanyMainPageRequest
	(*GetCompanyMainPageResponse)(nil), // 2: stats.GetCompanyMainPageResponse
	(*SubmitClosedLeadRequest)(nil),    // 3: stats.SubmitClosedLeadRequest
	(*emptypb.Empty)(nil),              // 4: google.protobuf.Empty
}
var file_service_proto_depIdxs = []int32{
	0, // 0: stats.GetCompanyMainPageResponse.data:type_name -> stats.Data
	1, // 1: stats.StatsService.GetCompanyMainPage:input_type -> stats.GetCompanyMainPageRequest
	3, // 2: stats.StatsService.SubmitClosedLead:input_type -> stats.SubmitClosedLeadRequest
	2, // 3: stats.StatsService.GetCompanyMainPage:output_type -> stats.GetCompanyMainPageResponse
	4, // 4: stats.StatsService.SubmitClosedLead:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompanyMainPageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompanyMainPageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitClosedLeadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
