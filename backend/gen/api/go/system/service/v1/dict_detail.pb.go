// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: system/service/v1/dict_detail.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 数据字典详情
type DictDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`               // ID
	DictId      *uint32                `protobuf:"varint,2,opt,name=dictId,proto3,oneof" json:"dictId,omitempty"` // ID
	OrderNo     *int32                 `protobuf:"varint,3,opt,name=orderNo,proto3,oneof" json:"orderNo,omitempty"`
	Label       *string                `protobuf:"bytes,4,opt,name=label,proto3,oneof" json:"label,omitempty"`               // 字典标签
	Value       *string                `protobuf:"bytes,5,opt,name=value,proto3,oneof" json:"value,omitempty"`               // 字典值
	CreatorId   *uint32                `protobuf:"varint,10,opt,name=creatorId,proto3,oneof" json:"creatorId,omitempty"`     // 创建者ID
	CreatorName *uint32                `protobuf:"varint,11,opt,name=creatorName,proto3,oneof" json:"creatorName,omitempty"` // 创建者名字
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,30,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,31,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"`
	DeleteTime  *timestamppb.Timestamp `protobuf:"bytes,32,opt,name=delete_time,json=deleteTime,proto3,oneof" json:"delete_time,omitempty"`
}

func (x *DictDetail) Reset() {
	*x = DictDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_v1_dict_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictDetail) ProtoMessage() {}

func (x *DictDetail) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_dict_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictDetail.ProtoReflect.Descriptor instead.
func (*DictDetail) Descriptor() ([]byte, []int) {
	return file_system_service_v1_dict_detail_proto_rawDescGZIP(), []int{0}
}

func (x *DictDetail) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DictDetail) GetDictId() uint32 {
	if x != nil && x.DictId != nil {
		return *x.DictId
	}
	return 0
}

func (x *DictDetail) GetOrderNo() int32 {
	if x != nil && x.OrderNo != nil {
		return *x.OrderNo
	}
	return 0
}

func (x *DictDetail) GetLabel() string {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return ""
}

func (x *DictDetail) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

func (x *DictDetail) GetCreatorId() uint32 {
	if x != nil && x.CreatorId != nil {
		return *x.CreatorId
	}
	return 0
}

func (x *DictDetail) GetCreatorName() uint32 {
	if x != nil && x.CreatorName != nil {
		return *x.CreatorName
	}
	return 0
}

func (x *DictDetail) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *DictDetail) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *DictDetail) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

// 查询字典详情列表 - 答复
type ListDictDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*DictDetail `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListDictDetailResponse) Reset() {
	*x = ListDictDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_v1_dict_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDictDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDictDetailResponse) ProtoMessage() {}

func (x *ListDictDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_dict_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDictDetailResponse.ProtoReflect.Descriptor instead.
func (*ListDictDetailResponse) Descriptor() ([]byte, []int) {
	return file_system_service_v1_dict_detail_proto_rawDescGZIP(), []int{1}
}

func (x *ListDictDetailResponse) GetItems() []*DictDetail {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListDictDetailResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 查询字典详情 - 请求
type GetDictDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code *string `protobuf:"bytes,3,opt,name=code,proto3,oneof" json:"code,omitempty"`
}

func (x *GetDictDetailRequest) Reset() {
	*x = GetDictDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_v1_dict_detail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDictDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDictDetailRequest) ProtoMessage() {}

func (x *GetDictDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_dict_detail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDictDetailRequest.ProtoReflect.Descriptor instead.
func (*GetDictDetailRequest) Descriptor() ([]byte, []int) {
	return file_system_service_v1_dict_detail_proto_rawDescGZIP(), []int{2}
}

func (x *GetDictDetailRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetDictDetailRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

// 创建字典详情 - 请求
type CreateDictDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detail     *DictDetail `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
	OperatorId *uint32     `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *CreateDictDetailRequest) Reset() {
	*x = CreateDictDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_v1_dict_detail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDictDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDictDetailRequest) ProtoMessage() {}

func (x *CreateDictDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_dict_detail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDictDetailRequest.ProtoReflect.Descriptor instead.
func (*CreateDictDetailRequest) Descriptor() ([]byte, []int) {
	return file_system_service_v1_dict_detail_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDictDetailRequest) GetDetail() *DictDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *CreateDictDetailRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 更新字典详情 - 请求
type UpdateDictDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detail     *DictDetail `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
	OperatorId *uint32     `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *UpdateDictDetailRequest) Reset() {
	*x = UpdateDictDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_v1_dict_detail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDictDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDictDetailRequest) ProtoMessage() {}

func (x *UpdateDictDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_dict_detail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDictDetailRequest.ProtoReflect.Descriptor instead.
func (*UpdateDictDetailRequest) Descriptor() ([]byte, []int) {
	return file_system_service_v1_dict_detail_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDictDetailRequest) GetDetail() *DictDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *UpdateDictDetailRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 删除字典详情 - 请求
type DeleteDictDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OperatorId *uint32 `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *DeleteDictDetailRequest) Reset() {
	*x = DeleteDictDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_v1_dict_detail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDictDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDictDetailRequest) ProtoMessage() {}

func (x *DeleteDictDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_dict_detail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDictDetailRequest.ProtoReflect.Descriptor instead.
func (*DeleteDictDetailRequest) Descriptor() ([]byte, []int) {
	return file_system_service_v1_dict_detail_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDictDetailRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteDictDetailRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

var File_system_service_v1_dict_detail_proto protoreflect.FileDescriptor

var file_system_service_v1_dict_detail_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x05, 0x0a,
	0x0a, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x14, 0xba, 0x47, 0x11, 0x92, 0x02, 0x0e, 0xe5,
	0xad, 0x97, 0xe5, 0x85, 0xb8, 0xe8, 0xaf, 0xa6, 0xe6, 0x83, 0x85, 0x49, 0x44, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x0e, 0xba, 0x47, 0x0b, 0x92, 0x02, 0x08, 0xe5, 0xad, 0x97, 0xe5, 0x85, 0xb8, 0x49,
	0x44, 0x48, 0x00, 0x52, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2e,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0f, 0xba, 0x47, 0x0c, 0x92, 0x02, 0x09, 0xe6, 0x8e, 0x92, 0xe5, 0xba, 0x8f, 0xe5, 0x8f, 0xb7,
	0x48, 0x01, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x2d,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xba,
	0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe5, 0xad, 0x97, 0xe5, 0x85, 0xb8, 0xe6, 0xa0, 0x87, 0xe7, 0xad,
	0xbe, 0x48, 0x02, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xba, 0x47,
	0x0c, 0x92, 0x02, 0x09, 0xe5, 0xad, 0x97, 0xe5, 0x85, 0xb8, 0xe5, 0x80, 0xbc, 0x48, 0x03, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x11, 0xba, 0x47,
	0x0e, 0x92, 0x02, 0x0b, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe8, 0x80, 0x85, 0x49, 0x44, 0x48,
	0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x3c, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x15, 0xba, 0x47, 0x12, 0x92, 0x02, 0x0f, 0xe5, 0x88, 0x9b, 0xe5,
	0xbb, 0xba, 0xe8, 0x80, 0x85, 0xe5, 0x90, 0x8d, 0xe5, 0xad, 0x97, 0x48, 0x05, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x06,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x07, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x48, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x63, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x23,
	0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35,
	0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x17, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x32, 0xb7, 0x04, 0x0a, 0x11, 0x44, 0x69, 0x63,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x19, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44,
	0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x27, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x58, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x2a, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x5f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x22, 0x00, 0x42, 0xcd, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x69,
	0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68,
	0x69, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x11, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_service_v1_dict_detail_proto_rawDescOnce sync.Once
	file_system_service_v1_dict_detail_proto_rawDescData = file_system_service_v1_dict_detail_proto_rawDesc
)

func file_system_service_v1_dict_detail_proto_rawDescGZIP() []byte {
	file_system_service_v1_dict_detail_proto_rawDescOnce.Do(func() {
		file_system_service_v1_dict_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_service_v1_dict_detail_proto_rawDescData)
	})
	return file_system_service_v1_dict_detail_proto_rawDescData
}

var file_system_service_v1_dict_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_system_service_v1_dict_detail_proto_goTypes = []interface{}{
	(*DictDetail)(nil),              // 0: system.service.v1.DictDetail
	(*ListDictDetailResponse)(nil),  // 1: system.service.v1.ListDictDetailResponse
	(*GetDictDetailRequest)(nil),    // 2: system.service.v1.GetDictDetailRequest
	(*CreateDictDetailRequest)(nil), // 3: system.service.v1.CreateDictDetailRequest
	(*UpdateDictDetailRequest)(nil), // 4: system.service.v1.UpdateDictDetailRequest
	(*DeleteDictDetailRequest)(nil), // 5: system.service.v1.DeleteDictDetailRequest
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
	(*v1.PagingRequest)(nil),        // 7: pagination.PagingRequest
	(*emptypb.Empty)(nil),           // 8: google.protobuf.Empty
}
var file_system_service_v1_dict_detail_proto_depIdxs = []int32{
	6,  // 0: system.service.v1.DictDetail.create_time:type_name -> google.protobuf.Timestamp
	6,  // 1: system.service.v1.DictDetail.update_time:type_name -> google.protobuf.Timestamp
	6,  // 2: system.service.v1.DictDetail.delete_time:type_name -> google.protobuf.Timestamp
	0,  // 3: system.service.v1.ListDictDetailResponse.items:type_name -> system.service.v1.DictDetail
	0,  // 4: system.service.v1.CreateDictDetailRequest.detail:type_name -> system.service.v1.DictDetail
	0,  // 5: system.service.v1.UpdateDictDetailRequest.detail:type_name -> system.service.v1.DictDetail
	7,  // 6: system.service.v1.DictDetailService.ListDictDetail:input_type -> pagination.PagingRequest
	2,  // 7: system.service.v1.DictDetailService.GetDictDetail:input_type -> system.service.v1.GetDictDetailRequest
	3,  // 8: system.service.v1.DictDetailService.CreateDictDetail:input_type -> system.service.v1.CreateDictDetailRequest
	4,  // 9: system.service.v1.DictDetailService.UpdateDictDetail:input_type -> system.service.v1.UpdateDictDetailRequest
	5,  // 10: system.service.v1.DictDetailService.DeleteDictDetail:input_type -> system.service.v1.DeleteDictDetailRequest
	2,  // 11: system.service.v1.DictDetailService.GetDictDetailByCode:input_type -> system.service.v1.GetDictDetailRequest
	1,  // 12: system.service.v1.DictDetailService.ListDictDetail:output_type -> system.service.v1.ListDictDetailResponse
	0,  // 13: system.service.v1.DictDetailService.GetDictDetail:output_type -> system.service.v1.DictDetail
	8,  // 14: system.service.v1.DictDetailService.CreateDictDetail:output_type -> google.protobuf.Empty
	8,  // 15: system.service.v1.DictDetailService.UpdateDictDetail:output_type -> google.protobuf.Empty
	8,  // 16: system.service.v1.DictDetailService.DeleteDictDetail:output_type -> google.protobuf.Empty
	0,  // 17: system.service.v1.DictDetailService.GetDictDetailByCode:output_type -> system.service.v1.DictDetail
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_system_service_v1_dict_detail_proto_init() }
func file_system_service_v1_dict_detail_proto_init() {
	if File_system_service_v1_dict_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_service_v1_dict_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictDetail); i {
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
		file_system_service_v1_dict_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDictDetailResponse); i {
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
		file_system_service_v1_dict_detail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDictDetailRequest); i {
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
		file_system_service_v1_dict_detail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDictDetailRequest); i {
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
		file_system_service_v1_dict_detail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDictDetailRequest); i {
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
		file_system_service_v1_dict_detail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDictDetailRequest); i {
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
	file_system_service_v1_dict_detail_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_system_service_v1_dict_detail_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_system_service_v1_dict_detail_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_system_service_v1_dict_detail_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_system_service_v1_dict_detail_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_service_v1_dict_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_service_v1_dict_detail_proto_goTypes,
		DependencyIndexes: file_system_service_v1_dict_detail_proto_depIdxs,
		MessageInfos:      file_system_service_v1_dict_detail_proto_msgTypes,
	}.Build()
	File_system_service_v1_dict_detail_proto = out.File
	file_system_service_v1_dict_detail_proto_rawDesc = nil
	file_system_service_v1_dict_detail_proto_goTypes = nil
	file_system_service_v1_dict_detail_proto_depIdxs = nil
}
