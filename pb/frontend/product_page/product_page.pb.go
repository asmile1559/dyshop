// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: product_page.proto

package product_page

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListProductsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int64                  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	CategoryName  string                 `protobuf:"bytes,3,opt,name=categoryName,proto3" json:"categoryName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsReq) Reset() {
	*x = ListProductsReq{}
	mi := &file_product_page_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsReq) ProtoMessage() {}

func (x *ListProductsReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsReq.ProtoReflect.Descriptor instead.
func (*ListProductsReq) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{0}
}

func (x *ListProductsReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProductsReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListProductsReq) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

type ListProductsResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsResp) Reset() {
	*x = ListProductsResp{}
	mi := &file_product_page_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResp) ProtoMessage() {}

func (x *ListProductsResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResp.ProtoReflect.Descriptor instead.
func (*ListProductsResp) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{1}
}

type GetProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductReq) Reset() {
	*x = GetProductReq{}
	mi := &file_product_page_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReq) ProtoMessage() {}

func (x *GetProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReq.ProtoReflect.Descriptor instead.
func (*GetProductReq) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductResp) Reset() {
	*x = GetProductResp{}
	mi := &file_product_page_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResp) ProtoMessage() {}

func (x *GetProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResp.ProtoReflect.Descriptor instead.
func (*GetProductResp) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{3}
}

type SearchProductsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Q             string                 `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchProductsReq) Reset() {
	*x = SearchProductsReq{}
	mi := &file_product_page_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchProductsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsReq) ProtoMessage() {}

func (x *SearchProductsReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsReq.ProtoReflect.Descriptor instead.
func (*SearchProductsReq) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{4}
}

func (x *SearchProductsReq) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

type SearchProductsResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchProductsResp) Reset() {
	*x = SearchProductsResp{}
	mi := &file_product_page_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchProductsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsResp) ProtoMessage() {}

func (x *SearchProductsResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsResp.ProtoReflect.Descriptor instead.
func (*SearchProductsResp) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{5}
}

type CreateProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Picture       string                 `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty"`
	Price         float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Amount        uint32                 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductReq) Reset() {
	*x = CreateProductReq{}
	mi := &file_product_page_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductReq) ProtoMessage() {}

func (x *CreateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductReq.ProtoReflect.Descriptor instead.
func (*CreateProductReq) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{6}
}

func (x *CreateProductReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductReq) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *CreateProductReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductReq) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductResp) Reset() {
	*x = CreateProductResp{}
	mi := &file_product_page_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResp) ProtoMessage() {}

func (x *CreateProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResp.ProtoReflect.Descriptor instead.
func (*CreateProductResp) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{7}
}

type ModifyProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description   *string                `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Picture       *string                `protobuf:"bytes,4,opt,name=picture,proto3,oneof" json:"picture,omitempty"`
	Price         *float32               `protobuf:"fixed32,5,opt,name=price,proto3,oneof" json:"price,omitempty"`
	Amount        *uint32                `protobuf:"varint,6,opt,name=amount,proto3,oneof" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ModifyProductReq) Reset() {
	*x = ModifyProductReq{}
	mi := &file_product_page_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModifyProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyProductReq) ProtoMessage() {}

func (x *ModifyProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyProductReq.ProtoReflect.Descriptor instead.
func (*ModifyProductReq) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{8}
}

func (x *ModifyProductReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ModifyProductReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ModifyProductReq) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ModifyProductReq) GetPicture() string {
	if x != nil && x.Picture != nil {
		return *x.Picture
	}
	return ""
}

func (x *ModifyProductReq) GetPrice() float32 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *ModifyProductReq) GetAmount() uint32 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

type ModifyProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ModifyProductResp) Reset() {
	*x = ModifyProductResp{}
	mi := &file_product_page_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModifyProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyProductResp) ProtoMessage() {}

func (x *ModifyProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyProductResp.ProtoReflect.Descriptor instead.
func (*ModifyProductResp) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{9}
}

type DeleteProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductReq) Reset() {
	*x = DeleteProductReq{}
	mi := &file_product_page_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductReq) ProtoMessage() {}

func (x *DeleteProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductReq.ProtoReflect.Descriptor instead.
func (*DeleteProductReq) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteProductReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductResp) Reset() {
	*x = DeleteProductResp{}
	mi := &file_product_page_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductResp) ProtoMessage() {}

func (x *DeleteProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_page_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductResp.ProtoReflect.Descriptor instead.
func (*DeleteProductResp) Descriptor() ([]byte, []int) {
	return file_product_page_proto_rawDescGZIP(), []int{11}
}

var File_product_page_proto protoreflect.FileDescriptor

var file_product_page_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x65, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x21, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x71, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x90, 0x01, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x13, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0xf3, 0x01, 0x0a, 0x10, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x22, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x32, 0xff, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x52, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6d, 0x69, 0x6c, 0x65, 0x31, 0x35, 0x35, 0x39,
	0x2f, 0x64, 0x79, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_product_page_proto_rawDescOnce sync.Once
	file_product_page_proto_rawDescData []byte
)

func file_product_page_proto_rawDescGZIP() []byte {
	file_product_page_proto_rawDescOnce.Do(func() {
		file_product_page_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_page_proto_rawDesc), len(file_product_page_proto_rawDesc)))
	})
	return file_product_page_proto_rawDescData
}

var file_product_page_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_product_page_proto_goTypes = []any{
	(*ListProductsReq)(nil),    // 0: product_page.ListProductsReq
	(*ListProductsResp)(nil),   // 1: product_page.ListProductsResp
	(*GetProductReq)(nil),      // 2: product_page.GetProductReq
	(*GetProductResp)(nil),     // 3: product_page.GetProductResp
	(*SearchProductsReq)(nil),  // 4: product_page.SearchProductsReq
	(*SearchProductsResp)(nil), // 5: product_page.SearchProductsResp
	(*CreateProductReq)(nil),   // 6: product_page.CreateProductReq
	(*CreateProductResp)(nil),  // 7: product_page.CreateProductResp
	(*ModifyProductReq)(nil),   // 8: product_page.ModifyProductReq
	(*ModifyProductResp)(nil),  // 9: product_page.ModifyProductResp
	(*DeleteProductReq)(nil),   // 10: product_page.DeleteProductReq
	(*DeleteProductResp)(nil),  // 11: product_page.DeleteProductResp
}
var file_product_page_proto_depIdxs = []int32{
	0,  // 0: product_page.ProductService.ListProducts:input_type -> product_page.ListProductsReq
	2,  // 1: product_page.ProductService.GetProduct:input_type -> product_page.GetProductReq
	4,  // 2: product_page.ProductService.SearchProducts:input_type -> product_page.SearchProductsReq
	6,  // 3: product_page.ProductService.CreateProduct:input_type -> product_page.CreateProductReq
	8,  // 4: product_page.ProductService.ModifyProduct:input_type -> product_page.ModifyProductReq
	10, // 5: product_page.ProductService.DeleteProduct:input_type -> product_page.DeleteProductReq
	1,  // 6: product_page.ProductService.ListProducts:output_type -> product_page.ListProductsResp
	3,  // 7: product_page.ProductService.GetProduct:output_type -> product_page.GetProductResp
	5,  // 8: product_page.ProductService.SearchProducts:output_type -> product_page.SearchProductsResp
	7,  // 9: product_page.ProductService.CreateProduct:output_type -> product_page.CreateProductResp
	9,  // 10: product_page.ProductService.ModifyProduct:output_type -> product_page.ModifyProductResp
	11, // 11: product_page.ProductService.DeleteProduct:output_type -> product_page.DeleteProductResp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_product_page_proto_init() }
func file_product_page_proto_init() {
	if File_product_page_proto != nil {
		return
	}
	file_product_page_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_page_proto_rawDesc), len(file_product_page_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_page_proto_goTypes,
		DependencyIndexes: file_product_page_proto_depIdxs,
		MessageInfos:      file_product_page_proto_msgTypes,
	}.Build()
	File_product_page_proto = out.File
	file_product_page_proto_goTypes = nil
	file_product_page_proto_depIdxs = nil
}
