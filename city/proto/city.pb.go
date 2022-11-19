// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: city.proto

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

// GetAllCities
type GetAllCitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header     `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   []*CityData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllCitiesResponse) Reset() {
	*x = GetAllCitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCitiesResponse) ProtoMessage() {}

func (x *GetAllCitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCitiesResponse.ProtoReflect.Descriptor instead.
func (*GetAllCitiesResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllCitiesResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetAllCitiesResponse) GetData() []*CityData {
	if x != nil {
		return x.Data
	}
	return nil
}

// GetCityByID
type GetCityByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *GetCityByIDInput `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetCityByIDRequest) Reset() {
	*x = GetCityByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityByIDRequest) ProtoMessage() {}

func (x *GetCityByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityByIDRequest.ProtoReflect.Descriptor instead.
func (*GetCityByIDRequest) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{1}
}

func (x *GetCityByIDRequest) GetRequest() *GetCityByIDInput {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetCityByIDInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCityByIDInput) Reset() {
	*x = GetCityByIDInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityByIDInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityByIDInput) ProtoMessage() {}

func (x *GetCityByIDInput) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityByIDInput.ProtoReflect.Descriptor instead.
func (*GetCityByIDInput) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{2}
}

func (x *GetCityByIDInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCityByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   *CityData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCityByIDResponse) Reset() {
	*x = GetCityByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityByIDResponse) ProtoMessage() {}

func (x *GetCityByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityByIDResponse.ProtoReflect.Descriptor instead.
func (*GetCityByIDResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{3}
}

func (x *GetCityByIDResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetCityByIDResponse) GetData() *CityData {
	if x != nil {
		return x.Data
	}
	return nil
}

// GetCitiesByIDs
type GetCitiesByIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *GetCitiesByIDsInput `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetCitiesByIDsRequest) Reset() {
	*x = GetCitiesByIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCitiesByIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCitiesByIDsRequest) ProtoMessage() {}

func (x *GetCitiesByIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCitiesByIDsRequest.ProtoReflect.Descriptor instead.
func (*GetCitiesByIDsRequest) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{4}
}

func (x *GetCitiesByIDsRequest) GetRequest() *GetCitiesByIDsInput {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetCitiesByIDsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCitiesByIDsInput) Reset() {
	*x = GetCitiesByIDsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCitiesByIDsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCitiesByIDsInput) ProtoMessage() {}

func (x *GetCitiesByIDsInput) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCitiesByIDsInput.ProtoReflect.Descriptor instead.
func (*GetCitiesByIDsInput) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{5}
}

func (x *GetCitiesByIDsInput) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetCitiesByIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header     `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   []*CityData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCitiesByIDsResponse) Reset() {
	*x = GetCitiesByIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCitiesByIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCitiesByIDsResponse) ProtoMessage() {}

func (x *GetCitiesByIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCitiesByIDsResponse.ProtoReflect.Descriptor instead.
func (*GetCitiesByIDsResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{6}
}

func (x *GetCitiesByIDsResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetCitiesByIDsResponse) GetData() []*CityData {
	if x != nil {
		return x.Data
	}
	return nil
}

// Base Data
type CityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityID     string `protobuf:"bytes,1,opt,name=cityID,proto3" json:"cityID,omitempty"`
	CityName   string `protobuf:"bytes,2,opt,name=cityName,proto3" json:"cityName,omitempty"`
	CityType   string `protobuf:"bytes,3,opt,name=cityType,proto3" json:"cityType,omitempty"`
	ProvinceID string `protobuf:"bytes,4,opt,name=provinceID,proto3" json:"provinceID,omitempty"`
}

func (x *CityData) Reset() {
	*x = CityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityData) ProtoMessage() {}

func (x *CityData) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityData.ProtoReflect.Descriptor instead.
func (*CityData) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{7}
}

func (x *CityData) GetCityID() string {
	if x != nil {
		return x.CityID
	}
	return ""
}

func (x *CityData) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *CityData) GetCityType() string {
	if x != nil {
		return x.CityType
	}
	return ""
}

func (x *CityData) GetProvinceID() string {
	if x != nil {
		return x.ProvinceID
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
		mi := &file_city_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[8]
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
	return file_city_proto_rawDescGZIP(), []int{8}
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
		mi := &file_city_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionalString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionalString) ProtoMessage() {}

func (x *OptionalString) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[9]
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
	return file_city_proto_rawDescGZIP(), []int{9}
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

var File_city_proto protoreflect.FileDescriptor

var file_city_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x60, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x46, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x43, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4c,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x49, 0x44, 0x22, 0x3a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x3e, 0x0a, 0x0e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x53, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x61, 0x73, 0x53, 0x65, 0x74, 0x32,
	0xe8, 0x01, 0x0a, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x12, 0x1b, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79,
	0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_proto_rawDescOnce sync.Once
	file_city_proto_rawDescData = file_city_proto_rawDesc
)

func file_city_proto_rawDescGZIP() []byte {
	file_city_proto_rawDescOnce.Do(func() {
		file_city_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_proto_rawDescData)
	})
	return file_city_proto_rawDescData
}

var file_city_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_city_proto_goTypes = []interface{}{
	(*GetAllCitiesResponse)(nil),   // 0: city.GetAllCitiesResponse
	(*GetCityByIDRequest)(nil),     // 1: city.GetCityByIDRequest
	(*GetCityByIDInput)(nil),       // 2: city.GetCityByIDInput
	(*GetCityByIDResponse)(nil),    // 3: city.GetCityByIDResponse
	(*GetCitiesByIDsRequest)(nil),  // 4: city.GetCitiesByIDsRequest
	(*GetCitiesByIDsInput)(nil),    // 5: city.GetCitiesByIDsInput
	(*GetCitiesByIDsResponse)(nil), // 6: city.GetCitiesByIDsResponse
	(*CityData)(nil),               // 7: city.CityData
	(*Header)(nil),                 // 8: city.Header
	(*OptionalString)(nil),         // 9: city.OptionalString
	(*emptypb.Empty)(nil),          // 10: google.protobuf.Empty
}
var file_city_proto_depIdxs = []int32{
	8,  // 0: city.GetAllCitiesResponse.header:type_name -> city.Header
	7,  // 1: city.GetAllCitiesResponse.data:type_name -> city.CityData
	2,  // 2: city.GetCityByIDRequest.request:type_name -> city.GetCityByIDInput
	8,  // 3: city.GetCityByIDResponse.header:type_name -> city.Header
	7,  // 4: city.GetCityByIDResponse.data:type_name -> city.CityData
	5,  // 5: city.GetCitiesByIDsRequest.request:type_name -> city.GetCitiesByIDsInput
	8,  // 6: city.GetCitiesByIDsResponse.header:type_name -> city.Header
	7,  // 7: city.GetCitiesByIDsResponse.data:type_name -> city.CityData
	10, // 8: city.CityService.GetAllCities:input_type -> google.protobuf.Empty
	1,  // 9: city.CityService.GetCityByID:input_type -> city.GetCityByIDRequest
	4,  // 10: city.CityService.GetCitiesByIDs:input_type -> city.GetCitiesByIDsRequest
	0,  // 11: city.CityService.GetAllCities:output_type -> city.GetAllCitiesResponse
	3,  // 12: city.CityService.GetCityByID:output_type -> city.GetCityByIDResponse
	6,  // 13: city.CityService.GetCitiesByIDs:output_type -> city.GetCitiesByIDsResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_city_proto_init() }
func file_city_proto_init() {
	if File_city_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCitiesResponse); i {
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
		file_city_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCityByIDRequest); i {
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
		file_city_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCityByIDInput); i {
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
		file_city_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCityByIDResponse); i {
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
		file_city_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCitiesByIDsRequest); i {
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
		file_city_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCitiesByIDsInput); i {
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
		file_city_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCitiesByIDsResponse); i {
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
		file_city_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityData); i {
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
		file_city_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_city_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_city_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_proto_goTypes,
		DependencyIndexes: file_city_proto_depIdxs,
		MessageInfos:      file_city_proto_msgTypes,
	}.Build()
	File_city_proto = out.File
	file_city_proto_rawDesc = nil
	file_city_proto_goTypes = nil
	file_city_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CityServiceClient is the client API for CityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CityServiceClient interface {
	GetAllCities(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllCitiesResponse, error)
	GetCityByID(ctx context.Context, in *GetCityByIDRequest, opts ...grpc.CallOption) (*GetCityByIDResponse, error)
	GetCitiesByIDs(ctx context.Context, in *GetCitiesByIDsRequest, opts ...grpc.CallOption) (*GetCitiesByIDsResponse, error)
}

type cityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCityServiceClient(cc grpc.ClientConnInterface) CityServiceClient {
	return &cityServiceClient{cc}
}

func (c *cityServiceClient) GetAllCities(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllCitiesResponse, error) {
	out := new(GetAllCitiesResponse)
	err := c.cc.Invoke(ctx, "/city.CityService/GetAllCities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) GetCityByID(ctx context.Context, in *GetCityByIDRequest, opts ...grpc.CallOption) (*GetCityByIDResponse, error) {
	out := new(GetCityByIDResponse)
	err := c.cc.Invoke(ctx, "/city.CityService/GetCityByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) GetCitiesByIDs(ctx context.Context, in *GetCitiesByIDsRequest, opts ...grpc.CallOption) (*GetCitiesByIDsResponse, error) {
	out := new(GetCitiesByIDsResponse)
	err := c.cc.Invoke(ctx, "/city.CityService/GetCitiesByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CityServiceServer is the server API for CityService service.
type CityServiceServer interface {
	GetAllCities(context.Context, *emptypb.Empty) (*GetAllCitiesResponse, error)
	GetCityByID(context.Context, *GetCityByIDRequest) (*GetCityByIDResponse, error)
	GetCitiesByIDs(context.Context, *GetCitiesByIDsRequest) (*GetCitiesByIDsResponse, error)
}

// UnimplementedCityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCityServiceServer struct {
}

func (*UnimplementedCityServiceServer) GetAllCities(context.Context, *emptypb.Empty) (*GetAllCitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCities not implemented")
}
func (*UnimplementedCityServiceServer) GetCityByID(context.Context, *GetCityByIDRequest) (*GetCityByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCityByID not implemented")
}
func (*UnimplementedCityServiceServer) GetCitiesByIDs(context.Context, *GetCitiesByIDsRequest) (*GetCitiesByIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCitiesByIDs not implemented")
}

func RegisterCityServiceServer(s *grpc.Server, srv CityServiceServer) {
	s.RegisterService(&_CityService_serviceDesc, srv)
}

func _CityService_GetAllCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).GetAllCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.CityService/GetAllCities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).GetAllCities(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_GetCityByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCityByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).GetCityByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.CityService/GetCityByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).GetCityByID(ctx, req.(*GetCityByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_GetCitiesByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCitiesByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).GetCitiesByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.CityService/GetCitiesByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).GetCitiesByIDs(ctx, req.(*GetCitiesByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "city.CityService",
	HandlerType: (*CityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCities",
			Handler:    _CityService_GetAllCities_Handler,
		},
		{
			MethodName: "GetCityByID",
			Handler:    _CityService_GetCityByID_Handler,
		},
		{
			MethodName: "GetCitiesByIDs",
			Handler:    _CityService_GetCitiesByIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "city.proto",
}
