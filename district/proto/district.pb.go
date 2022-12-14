// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: district.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// GetAllDistricts
type GetAllDistrictsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header         `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   []*DistrictData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllDistrictsResponse) Reset() {
	*x = GetAllDistrictsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDistrictsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDistrictsResponse) ProtoMessage() {}

func (x *GetAllDistrictsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDistrictsResponse.ProtoReflect.Descriptor instead.
func (*GetAllDistrictsResponse) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllDistrictsResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetAllDistrictsResponse) GetData() []*DistrictData {
	if x != nil {
		return x.Data
	}
	return nil
}

// GetDistrictByID
type GetDistrictByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *GetDistrictByIDInput `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetDistrictByIDRequest) Reset() {
	*x = GetDistrictByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDistrictByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDistrictByIDRequest) ProtoMessage() {}

func (x *GetDistrictByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDistrictByIDRequest.ProtoReflect.Descriptor instead.
func (*GetDistrictByIDRequest) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{1}
}

func (x *GetDistrictByIDRequest) GetRequest() *GetDistrictByIDInput {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetDistrictByIDInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDistrictByIDInput) Reset() {
	*x = GetDistrictByIDInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDistrictByIDInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDistrictByIDInput) ProtoMessage() {}

func (x *GetDistrictByIDInput) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDistrictByIDInput.ProtoReflect.Descriptor instead.
func (*GetDistrictByIDInput) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{2}
}

func (x *GetDistrictByIDInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDistrictByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header       `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   *DistrictData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetDistrictByIDResponse) Reset() {
	*x = GetDistrictByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDistrictByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDistrictByIDResponse) ProtoMessage() {}

func (x *GetDistrictByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDistrictByIDResponse.ProtoReflect.Descriptor instead.
func (*GetDistrictByIDResponse) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{3}
}

func (x *GetDistrictByIDResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetDistrictByIDResponse) GetData() *DistrictData {
	if x != nil {
		return x.Data
	}
	return nil
}

// GetDistrictsByIDs
type GetDistrictsByIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *GetDistrictsByIDsInput `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetDistrictsByIDsRequest) Reset() {
	*x = GetDistrictsByIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDistrictsByIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDistrictsByIDsRequest) ProtoMessage() {}

func (x *GetDistrictsByIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDistrictsByIDsRequest.ProtoReflect.Descriptor instead.
func (*GetDistrictsByIDsRequest) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{4}
}

func (x *GetDistrictsByIDsRequest) GetRequest() *GetDistrictsByIDsInput {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetDistrictsByIDsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDistrictsByIDsInput) Reset() {
	*x = GetDistrictsByIDsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDistrictsByIDsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDistrictsByIDsInput) ProtoMessage() {}

func (x *GetDistrictsByIDsInput) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDistrictsByIDsInput.ProtoReflect.Descriptor instead.
func (*GetDistrictsByIDsInput) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{5}
}

func (x *GetDistrictsByIDsInput) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetDistrictsByIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header         `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   []*DistrictData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetDistrictsByIDsResponse) Reset() {
	*x = GetDistrictsByIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDistrictsByIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDistrictsByIDsResponse) ProtoMessage() {}

func (x *GetDistrictsByIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDistrictsByIDsResponse.ProtoReflect.Descriptor instead.
func (*GetDistrictsByIDsResponse) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{6}
}

func (x *GetDistrictsByIDsResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetDistrictsByIDsResponse) GetData() []*DistrictData {
	if x != nil {
		return x.Data
	}
	return nil
}

// Base Data
type DistrictData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistrictID   string `protobuf:"bytes,1,opt,name=districtID,proto3" json:"districtID,omitempty"`
	DistrictName string `protobuf:"bytes,2,opt,name=districtName,proto3" json:"districtName,omitempty"`
	CityID       string `protobuf:"bytes,3,opt,name=cityID,proto3" json:"cityID,omitempty"`
}

func (x *DistrictData) Reset() {
	*x = DistrictData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistrictData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistrictData) ProtoMessage() {}

func (x *DistrictData) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistrictData.ProtoReflect.Descriptor instead.
func (*DistrictData) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{7}
}

func (x *DistrictData) GetDistrictID() string {
	if x != nil {
		return x.DistrictID
	}
	return ""
}

func (x *DistrictData) GetDistrictName() string {
	if x != nil {
		return x.DistrictName
	}
	return ""
}

func (x *DistrictData) GetCityID() string {
	if x != nil {
		return x.CityID
	}
	return ""
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message []string `protobuf:"bytes,2,rep,name=message,proto3" json:"message,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{8}
}

func (x *Header) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Header) GetMessage() []string {
	if x != nil {
		return x.Message
	}
	return nil
}

type OptionalString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	HasSet bool   `protobuf:"varint,2,opt,name=hasSet,proto3" json:"hasSet,omitempty"`
}

func (x *OptionalString) Reset() {
	*x = OptionalString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionalString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionalString) ProtoMessage() {}

func (x *OptionalString) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionalString.ProtoReflect.Descriptor instead.
func (*OptionalString) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{9}
}

func (x *OptionalString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *OptionalString) GetHasSet() bool {
	if x != nil {
		return x.HasSet
	}
	return false
}

var File_district_proto protoreflect.FileDescriptor

var file_district_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x52, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x6f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x42, 0x79, 0x49,
	0x44, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6a, 0x0a, 0x0c, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x69, 0x74, 0x79, 0x49, 0x44, 0x22, 0x3a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61,
	0x73, 0x53, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x61, 0x73, 0x53,
	0x65, 0x74, 0x32, 0x9b, 0x02, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x21, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73,
	0x42, 0x79, 0x49, 0x44, 0x73, 0x12, 0x22, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x42, 0x79, 0x49,
	0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_district_proto_rawDescOnce sync.Once
	file_district_proto_rawDescData = file_district_proto_rawDesc
)

func file_district_proto_rawDescGZIP() []byte {
	file_district_proto_rawDescOnce.Do(func() {
		file_district_proto_rawDescData = protoimpl.X.CompressGZIP(file_district_proto_rawDescData)
	})
	return file_district_proto_rawDescData
}

var file_district_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_district_proto_goTypes = []interface{}{
	(*GetAllDistrictsResponse)(nil),   // 0: district.GetAllDistrictsResponse
	(*GetDistrictByIDRequest)(nil),    // 1: district.GetDistrictByIDRequest
	(*GetDistrictByIDInput)(nil),      // 2: district.GetDistrictByIDInput
	(*GetDistrictByIDResponse)(nil),   // 3: district.GetDistrictByIDResponse
	(*GetDistrictsByIDsRequest)(nil),  // 4: district.GetDistrictsByIDsRequest
	(*GetDistrictsByIDsInput)(nil),    // 5: district.GetDistrictsByIDsInput
	(*GetDistrictsByIDsResponse)(nil), // 6: district.GetDistrictsByIDsResponse
	(*DistrictData)(nil),              // 7: district.DistrictData
	(*Header)(nil),                    // 8: district.Header
	(*OptionalString)(nil),            // 9: district.OptionalString
	(*emptypb.Empty)(nil),             // 10: google.protobuf.Empty
}
var file_district_proto_depIdxs = []int32{
	8,  // 0: district.GetAllDistrictsResponse.header:type_name -> district.Header
	7,  // 1: district.GetAllDistrictsResponse.data:type_name -> district.DistrictData
	2,  // 2: district.GetDistrictByIDRequest.request:type_name -> district.GetDistrictByIDInput
	8,  // 3: district.GetDistrictByIDResponse.header:type_name -> district.Header
	7,  // 4: district.GetDistrictByIDResponse.data:type_name -> district.DistrictData
	5,  // 5: district.GetDistrictsByIDsRequest.request:type_name -> district.GetDistrictsByIDsInput
	8,  // 6: district.GetDistrictsByIDsResponse.header:type_name -> district.Header
	7,  // 7: district.GetDistrictsByIDsResponse.data:type_name -> district.DistrictData
	10, // 8: district.DistrictService.GetAllDistricts:input_type -> google.protobuf.Empty
	1,  // 9: district.DistrictService.GetDistrictByID:input_type -> district.GetDistrictByIDRequest
	4,  // 10: district.DistrictService.GetDistrictsByIDs:input_type -> district.GetDistrictsByIDsRequest
	0,  // 11: district.DistrictService.GetAllDistricts:output_type -> district.GetAllDistrictsResponse
	3,  // 12: district.DistrictService.GetDistrictByID:output_type -> district.GetDistrictByIDResponse
	6,  // 13: district.DistrictService.GetDistrictsByIDs:output_type -> district.GetDistrictsByIDsResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_district_proto_init() }
func file_district_proto_init() {
	if File_district_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_district_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDistrictsResponse); i {
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
		file_district_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDistrictByIDRequest); i {
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
		file_district_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDistrictByIDInput); i {
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
		file_district_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDistrictByIDResponse); i {
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
		file_district_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDistrictsByIDsRequest); i {
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
		file_district_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDistrictsByIDsInput); i {
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
		file_district_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDistrictsByIDsResponse); i {
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
		file_district_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistrictData); i {
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
		file_district_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_district_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionalString); i {
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
			RawDescriptor: file_district_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_district_proto_goTypes,
		DependencyIndexes: file_district_proto_depIdxs,
		MessageInfos:      file_district_proto_msgTypes,
	}.Build()
	File_district_proto = out.File
	file_district_proto_rawDesc = nil
	file_district_proto_goTypes = nil
	file_district_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DistrictServiceClient is the client API for DistrictService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DistrictServiceClient interface {
	GetAllDistricts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllDistrictsResponse, error)
	GetDistrictByID(ctx context.Context, in *GetDistrictByIDRequest, opts ...grpc.CallOption) (*GetDistrictByIDResponse, error)
	GetDistrictsByIDs(ctx context.Context, in *GetDistrictsByIDsRequest, opts ...grpc.CallOption) (*GetDistrictsByIDsResponse, error)
}

type districtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDistrictServiceClient(cc grpc.ClientConnInterface) DistrictServiceClient {
	return &districtServiceClient{cc}
}

func (c *districtServiceClient) GetAllDistricts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllDistrictsResponse, error) {
	out := new(GetAllDistrictsResponse)
	err := c.cc.Invoke(ctx, "/district.DistrictService/GetAllDistricts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *districtServiceClient) GetDistrictByID(ctx context.Context, in *GetDistrictByIDRequest, opts ...grpc.CallOption) (*GetDistrictByIDResponse, error) {
	out := new(GetDistrictByIDResponse)
	err := c.cc.Invoke(ctx, "/district.DistrictService/GetDistrictByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *districtServiceClient) GetDistrictsByIDs(ctx context.Context, in *GetDistrictsByIDsRequest, opts ...grpc.CallOption) (*GetDistrictsByIDsResponse, error) {
	out := new(GetDistrictsByIDsResponse)
	err := c.cc.Invoke(ctx, "/district.DistrictService/GetDistrictsByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistrictServiceServer is the server API for DistrictService service.
type DistrictServiceServer interface {
	GetAllDistricts(context.Context, *emptypb.Empty) (*GetAllDistrictsResponse, error)
	GetDistrictByID(context.Context, *GetDistrictByIDRequest) (*GetDistrictByIDResponse, error)
	GetDistrictsByIDs(context.Context, *GetDistrictsByIDsRequest) (*GetDistrictsByIDsResponse, error)
}

// UnimplementedDistrictServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDistrictServiceServer struct {
}

func (*UnimplementedDistrictServiceServer) GetAllDistricts(context.Context, *emptypb.Empty) (*GetAllDistrictsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDistricts not implemented")
}
func (*UnimplementedDistrictServiceServer) GetDistrictByID(context.Context, *GetDistrictByIDRequest) (*GetDistrictByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDistrictByID not implemented")
}
func (*UnimplementedDistrictServiceServer) GetDistrictsByIDs(context.Context, *GetDistrictsByIDsRequest) (*GetDistrictsByIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDistrictsByIDs not implemented")
}

func RegisterDistrictServiceServer(s *grpc.Server, srv DistrictServiceServer) {
	s.RegisterService(&_DistrictService_serviceDesc, srv)
}

func _DistrictService_GetAllDistricts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistrictServiceServer).GetAllDistricts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/district.DistrictService/GetAllDistricts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistrictServiceServer).GetAllDistricts(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistrictService_GetDistrictByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDistrictByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistrictServiceServer).GetDistrictByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/district.DistrictService/GetDistrictByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistrictServiceServer).GetDistrictByID(ctx, req.(*GetDistrictByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistrictService_GetDistrictsByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDistrictsByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistrictServiceServer).GetDistrictsByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/district.DistrictService/GetDistrictsByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistrictServiceServer).GetDistrictsByIDs(ctx, req.(*GetDistrictsByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistrictService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "district.DistrictService",
	HandlerType: (*DistrictServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllDistricts",
			Handler:    _DistrictService_GetAllDistricts_Handler,
		},
		{
			MethodName: "GetDistrictByID",
			Handler:    _DistrictService_GetDistrictByID_Handler,
		},
		{
			MethodName: "GetDistrictsByIDs",
			Handler:    _DistrictService_GetDistrictsByIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "district.proto",
}
