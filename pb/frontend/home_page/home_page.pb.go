// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.12.4
// source: home_page.proto

package home_page

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetHomepageReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetHomepageReq) Reset() {
	*x = GetHomepageReq{}
	mi := &file_home_page_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHomepageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomepageReq) ProtoMessage() {}

func (x *GetHomepageReq) ProtoReflect() protoreflect.Message {
	mi := &file_home_page_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomepageReq.ProtoReflect.Descriptor instead.
func (*GetHomepageReq) Descriptor() ([]byte, []int) {
	return file_home_page_proto_rawDescGZIP(), []int{0}
}

type GetHomepageResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetHomepageResp) Reset() {
	*x = GetHomepageResp{}
	mi := &file_home_page_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHomepageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomepageResp) ProtoMessage() {}

func (x *GetHomepageResp) ProtoReflect() protoreflect.Message {
	mi := &file_home_page_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomepageResp.ProtoReflect.Descriptor instead.
func (*GetHomepageResp) Descriptor() ([]byte, []int) {
	return file_home_page_proto_rawDescGZIP(), []int{1}
}

type GetShowcaseReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Which         string                 `protobuf:"bytes,1,opt,name=which,proto3" json:"which,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShowcaseReq) Reset() {
	*x = GetShowcaseReq{}
	mi := &file_home_page_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShowcaseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShowcaseReq) ProtoMessage() {}

func (x *GetShowcaseReq) ProtoReflect() protoreflect.Message {
	mi := &file_home_page_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShowcaseReq.ProtoReflect.Descriptor instead.
func (*GetShowcaseReq) Descriptor() ([]byte, []int) {
	return file_home_page_proto_rawDescGZIP(), []int{2}
}

func (x *GetShowcaseReq) GetWhich() string {
	if x != nil {
		return x.Which
	}
	return ""
}

type GetShowcaseResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShowcaseResp) Reset() {
	*x = GetShowcaseResp{}
	mi := &file_home_page_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShowcaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShowcaseResp) ProtoMessage() {}

func (x *GetShowcaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_home_page_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShowcaseResp.ProtoReflect.Descriptor instead.
func (*GetShowcaseResp) Descriptor() ([]byte, []int) {
	return file_home_page_proto_rawDescGZIP(), []int{3}
}

type VerifyHomepageStatusReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyHomepageStatusReq) Reset() {
	*x = VerifyHomepageStatusReq{}
	mi := &file_home_page_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyHomepageStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyHomepageStatusReq) ProtoMessage() {}

func (x *VerifyHomepageStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_home_page_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyHomepageStatusReq.ProtoReflect.Descriptor instead.
func (*VerifyHomepageStatusReq) Descriptor() ([]byte, []int) {
	return file_home_page_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyHomepageStatusReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyHomepageStatusResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyHomepageStatusResp) Reset() {
	*x = VerifyHomepageStatusResp{}
	mi := &file_home_page_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyHomepageStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyHomepageStatusResp) ProtoMessage() {}

func (x *VerifyHomepageStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_home_page_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyHomepageStatusResp.ProtoReflect.Descriptor instead.
func (*VerifyHomepageStatusResp) Descriptor() ([]byte, []int) {
	return file_home_page_proto_rawDescGZIP(), []int{5}
}

var File_home_page_proto protoreflect.FileDescriptor

var file_home_page_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x22, 0x10, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x22, 0x11,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x26, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x68, 0x69, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x77, 0x68, 0x69, 0x63, 0x68, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2f, 0x0a, 0x17,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1a, 0x0a,
	0x18, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x32, 0x80, 0x02, 0x0a, 0x0b, 0x48, 0x6f,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x12, 0x19, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x14, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x22, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x3e, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6d, 0x69, 0x6c,
	0x65, 0x31, 0x35, 0x35, 0x39, 0x2f, 0x64, 0x79, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x70, 0x62, 0x2f,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x3b, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_home_page_proto_rawDescOnce sync.Once
	file_home_page_proto_rawDescData = file_home_page_proto_rawDesc
)

func file_home_page_proto_rawDescGZIP() []byte {
	file_home_page_proto_rawDescOnce.Do(func() {
		file_home_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_home_page_proto_rawDescData)
	})
	return file_home_page_proto_rawDescData
}

var file_home_page_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_home_page_proto_goTypes = []any{
	(*GetHomepageReq)(nil),           // 0: home_page.GetHomepageReq
	(*GetHomepageResp)(nil),          // 1: home_page.GetHomepageResp
	(*GetShowcaseReq)(nil),           // 2: home_page.GetShowcaseReq
	(*GetShowcaseResp)(nil),          // 3: home_page.GetShowcaseResp
	(*VerifyHomepageStatusReq)(nil),  // 4: home_page.VerifyHomepageStatusReq
	(*VerifyHomepageStatusResp)(nil), // 5: home_page.VerifyHomepageStatusResp
}
var file_home_page_proto_depIdxs = []int32{
	0, // 0: home_page.HomeService.GetHomepage:input_type -> home_page.GetHomepageReq
	2, // 1: home_page.HomeService.GetShowcase:input_type -> home_page.GetShowcaseReq
	4, // 2: home_page.HomeService.VerifyHomepageStatus:input_type -> home_page.VerifyHomepageStatusReq
	1, // 3: home_page.HomeService.GetHomepage:output_type -> home_page.GetHomepageResp
	3, // 4: home_page.HomeService.GetShowcase:output_type -> home_page.GetShowcaseResp
	5, // 5: home_page.HomeService.VerifyHomepageStatus:output_type -> home_page.VerifyHomepageStatusResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_home_page_proto_init() }
func file_home_page_proto_init() {
	if File_home_page_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_home_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_home_page_proto_goTypes,
		DependencyIndexes: file_home_page_proto_depIdxs,
		MessageInfos:      file_home_page_proto_msgTypes,
	}.Build()
	File_home_page_proto = out.File
	file_home_page_proto_rawDesc = nil
	file_home_page_proto_goTypes = nil
	file_home_page_proto_depIdxs = nil
}
