// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: order_page.proto

package order_page

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

type ListOrdersReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersReq) Reset() {
	*x = ListOrdersReq{}
	mi := &file_order_page_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersReq) ProtoMessage() {}

func (x *ListOrdersReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersReq.ProtoReflect.Descriptor instead.
func (*ListOrdersReq) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{0}
}

type ListOrdersResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersResp) Reset() {
	*x = ListOrdersResp{}
	mi := &file_order_page_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResp) ProtoMessage() {}

func (x *ListOrdersResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersResp.ProtoReflect.Descriptor instead.
func (*ListOrdersResp) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{1}
}

type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     uint32                 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Cost          float32                `protobuf:"fixed32,3,opt,name=cost,proto3" json:"cost,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_order_page_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItem) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StreetAddress string                 `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City          string                 `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string                 `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country       string                 `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode       string                 `protobuf:"bytes,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_order_page_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{3}
}

func (x *Address) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

type PlaceOrderReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserCurrency  string                 `protobuf:"bytes,1,opt,name=user_currency,json=userCurrency,proto3" json:"user_currency,omitempty"`
	Address       *Address               `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	OrderItems    []*OrderItem           `protobuf:"bytes,4,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlaceOrderReq) Reset() {
	*x = PlaceOrderReq{}
	mi := &file_order_page_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlaceOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderReq) ProtoMessage() {}

func (x *PlaceOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderReq.ProtoReflect.Descriptor instead.
func (*PlaceOrderReq) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{4}
}

func (x *PlaceOrderReq) GetUserCurrency() string {
	if x != nil {
		return x.UserCurrency
	}
	return ""
}

func (x *PlaceOrderReq) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PlaceOrderReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PlaceOrderReq) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

type PlaceOrderResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlaceOrderResp) Reset() {
	*x = PlaceOrderResp{}
	mi := &file_order_page_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlaceOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderResp) ProtoMessage() {}

func (x *PlaceOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderResp.ProtoReflect.Descriptor instead.
func (*PlaceOrderResp) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{5}
}

func (x *PlaceOrderResp) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderItems    []*OrderItem           `protobuf:"bytes,1,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	OrderId       string                 `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId        uint32                 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserCurrency  string                 `protobuf:"bytes,4,opt,name=user_currency,json=userCurrency,proto3" json:"user_currency,omitempty"`
	Address       *Address               `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Email         string                 `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt     int32                  `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_order_page_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{6}
}

func (x *Order) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetUserCurrency() string {
	if x != nil {
		return x.UserCurrency
	}
	return ""
}

func (x *Order) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Order) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Order) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type GetOrderReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderReq) Reset() {
	*x = GetOrderReq{}
	mi := &file_order_page_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderReq) ProtoMessage() {}

func (x *GetOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderReq.ProtoReflect.Descriptor instead.
func (*GetOrderReq) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{7}
}

func (x *GetOrderReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetOrderResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderResp) Reset() {
	*x = GetOrderResp{}
	mi := &file_order_page_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResp) ProtoMessage() {}

func (x *GetOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResp.ProtoReflect.Descriptor instead.
func (*GetOrderResp) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{8}
}

func (x *GetOrderResp) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type ModifyOrderReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ProductId     uint32                 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      *int32                 `protobuf:"varint,3,opt,name=quantity,proto3,oneof" json:"quantity,omitempty"`
	Cost          *float32               `protobuf:"fixed32,4,opt,name=cost,proto3,oneof" json:"cost,omitempty"`
	UserCurrency  *string                `protobuf:"bytes,5,opt,name=user_currency,json=userCurrency,proto3,oneof" json:"user_currency,omitempty"`
	Address       *Address               `protobuf:"bytes,6,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Email         *string                `protobuf:"bytes,7,opt,name=email,proto3,oneof" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ModifyOrderReq) Reset() {
	*x = ModifyOrderReq{}
	mi := &file_order_page_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModifyOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyOrderReq) ProtoMessage() {}

func (x *ModifyOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyOrderReq.ProtoReflect.Descriptor instead.
func (*ModifyOrderReq) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{9}
}

func (x *ModifyOrderReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ModifyOrderReq) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ModifyOrderReq) GetQuantity() int32 {
	if x != nil && x.Quantity != nil {
		return *x.Quantity
	}
	return 0
}

func (x *ModifyOrderReq) GetCost() float32 {
	if x != nil && x.Cost != nil {
		return *x.Cost
	}
	return 0
}

func (x *ModifyOrderReq) GetUserCurrency() string {
	if x != nil && x.UserCurrency != nil {
		return *x.UserCurrency
	}
	return ""
}

func (x *ModifyOrderReq) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ModifyOrderReq) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

type ModifyOrderResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ModifyOrderResp) Reset() {
	*x = ModifyOrderResp{}
	mi := &file_order_page_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModifyOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyOrderResp) ProtoMessage() {}

func (x *ModifyOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyOrderResp.ProtoReflect.Descriptor instead.
func (*ModifyOrderResp) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{10}
}

type CancelOrderReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       uint32                 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelOrderReq) Reset() {
	*x = CancelOrderReq{}
	mi := &file_order_page_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderReq) ProtoMessage() {}

func (x *CancelOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderReq.ProtoReflect.Descriptor instead.
func (*CancelOrderReq) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{11}
}

func (x *CancelOrderReq) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type CancelOrderResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelOrderResp) Reset() {
	*x = CancelOrderResp{}
	mi := &file_order_page_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderResp) ProtoMessage() {}

func (x *CancelOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderResp.ProtoReflect.Descriptor instead.
func (*CancelOrderResp) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{12}
}

type MarkOrderPaidReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarkOrderPaidReq) Reset() {
	*x = MarkOrderPaidReq{}
	mi := &file_order_page_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarkOrderPaidReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkOrderPaidReq) ProtoMessage() {}

func (x *MarkOrderPaidReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkOrderPaidReq.ProtoReflect.Descriptor instead.
func (*MarkOrderPaidReq) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{13}
}

func (x *MarkOrderPaidReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type MarkOrderPaidResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarkOrderPaidResp) Reset() {
	*x = MarkOrderPaidResp{}
	mi := &file_order_page_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarkOrderPaidResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkOrderPaidResp) ProtoMessage() {}

func (x *MarkOrderPaidResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkOrderPaidResp.ProtoReflect.Descriptor instead.
func (*MarkOrderPaidResp) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{14}
}

var File_order_page_proto protoreflect.FileDescriptor

var file_order_page_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x22, 0x0f,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x22,
	0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x5a, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x8f, 0x01,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0xb1, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x36, 0x0a, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x2b, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xfc, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2d, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x28, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0xbb, 0x02, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01,
	0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88,
	0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x2b, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x11, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x2d, 0x0a, 0x10, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x32, 0xc1, 0x03, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x4d,
	0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x69, 0x64, 0x12, 0x1c, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6d, 0x69, 0x6c, 0x65,
	0x31, 0x35, 0x35, 0x39, 0x2f, 0x64, 0x79, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x70, 0x62, 0x2f, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_order_page_proto_rawDescOnce sync.Once
	file_order_page_proto_rawDescData []byte
)

func file_order_page_proto_rawDescGZIP() []byte {
	file_order_page_proto_rawDescOnce.Do(func() {
		file_order_page_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_page_proto_rawDesc), len(file_order_page_proto_rawDesc)))
	})
	return file_order_page_proto_rawDescData
}

var file_order_page_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_order_page_proto_goTypes = []any{
	(*ListOrdersReq)(nil),     // 0: order_page.ListOrdersReq
	(*ListOrdersResp)(nil),    // 1: order_page.ListOrdersResp
	(*OrderItem)(nil),         // 2: order_page.OrderItem
	(*Address)(nil),           // 3: order_page.Address
	(*PlaceOrderReq)(nil),     // 4: order_page.PlaceOrderReq
	(*PlaceOrderResp)(nil),    // 5: order_page.PlaceOrderResp
	(*Order)(nil),             // 6: order_page.Order
	(*GetOrderReq)(nil),       // 7: order_page.GetOrderReq
	(*GetOrderResp)(nil),      // 8: order_page.GetOrderResp
	(*ModifyOrderReq)(nil),    // 9: order_page.ModifyOrderReq
	(*ModifyOrderResp)(nil),   // 10: order_page.ModifyOrderResp
	(*CancelOrderReq)(nil),    // 11: order_page.CancelOrderReq
	(*CancelOrderResp)(nil),   // 12: order_page.CancelOrderResp
	(*MarkOrderPaidReq)(nil),  // 13: order_page.MarkOrderPaidReq
	(*MarkOrderPaidResp)(nil), // 14: order_page.MarkOrderPaidResp
}
var file_order_page_proto_depIdxs = []int32{
	3,  // 0: order_page.PlaceOrderReq.address:type_name -> order_page.Address
	2,  // 1: order_page.PlaceOrderReq.order_items:type_name -> order_page.OrderItem
	2,  // 2: order_page.Order.order_items:type_name -> order_page.OrderItem
	3,  // 3: order_page.Order.address:type_name -> order_page.Address
	6,  // 4: order_page.GetOrderResp.order:type_name -> order_page.Order
	3,  // 5: order_page.ModifyOrderReq.address:type_name -> order_page.Address
	0,  // 6: order_page.OrderService.ListOrders:input_type -> order_page.ListOrdersReq
	4,  // 7: order_page.OrderService.PlaceOrder:input_type -> order_page.PlaceOrderReq
	7,  // 8: order_page.OrderService.GetOrder:input_type -> order_page.GetOrderReq
	9,  // 9: order_page.OrderService.ModifyOrder:input_type -> order_page.ModifyOrderReq
	11, // 10: order_page.OrderService.CancelOrder:input_type -> order_page.CancelOrderReq
	13, // 11: order_page.OrderService.MarkOrderPaid:input_type -> order_page.MarkOrderPaidReq
	1,  // 12: order_page.OrderService.ListOrders:output_type -> order_page.ListOrdersResp
	5,  // 13: order_page.OrderService.PlaceOrder:output_type -> order_page.PlaceOrderResp
	8,  // 14: order_page.OrderService.GetOrder:output_type -> order_page.GetOrderResp
	10, // 15: order_page.OrderService.ModifyOrder:output_type -> order_page.ModifyOrderResp
	12, // 16: order_page.OrderService.CancelOrder:output_type -> order_page.CancelOrderResp
	14, // 17: order_page.OrderService.MarkOrderPaid:output_type -> order_page.MarkOrderPaidResp
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_order_page_proto_init() }
func file_order_page_proto_init() {
	if File_order_page_proto != nil {
		return
	}
	file_order_page_proto_msgTypes[9].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_page_proto_rawDesc), len(file_order_page_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_page_proto_goTypes,
		DependencyIndexes: file_order_page_proto_depIdxs,
		MessageInfos:      file_order_page_proto_msgTypes,
	}.Build()
	File_order_page_proto = out.File
	file_order_page_proto_goTypes = nil
	file_order_page_proto_depIdxs = nil
}
