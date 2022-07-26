// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: telegraf.proto

package tfapi

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

type PathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// path不能为空
	// @inject_tag: uri:"path" form:"path" validate:"ne="
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path" uri:"path" form:"path" validate:"ne="`
}

func (x *PathRequest) Reset() {
	*x = PathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegraf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathRequest) ProtoMessage() {}

func (x *PathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telegraf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathRequest.ProtoReflect.Descriptor instead.
func (*PathRequest) Descriptor() ([]byte, []int) {
	return file_telegraf_proto_rawDescGZIP(), []int{0}
}

func (x *PathRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// path不能为空
	// @inject_tag: validate:"ne="
	Path    string `protobuf:"bytes,1,opt,name=path,proto3" json:"path" validate:"ne="`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content"`
}

func (x *ContentRequest) Reset() {
	*x = ContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegraf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentRequest) ProtoMessage() {}

func (x *ContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telegraf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentRequest.ProtoReflect.Descriptor instead.
func (*ContentRequest) Descriptor() ([]byte, []int) {
	return file_telegraf_proto_rawDescGZIP(), []int{1}
}

func (x *ContentRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ContentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegraf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telegraf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_telegraf_proto_rawDescGZIP(), []int{2}
}

func (x *CommonResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content"`
}

func (x *ContentResponse) Reset() {
	*x = ContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegraf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentResponse) ProtoMessage() {}

func (x *ContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telegraf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentResponse.ProtoReflect.Descriptor instead.
func (*ContentResponse) Descriptor() ([]byte, []int) {
	return file_telegraf_proto_rawDescGZIP(), []int{3}
}

func (x *ContentResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ContentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ContentResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type FilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Files   []*FileInfo `protobuf:"bytes,3,rep,name=files,proto3" json:"files"`
}

func (x *FilesResponse) Reset() {
	*x = FilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegraf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesResponse) ProtoMessage() {}

func (x *FilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telegraf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesResponse.ProtoReflect.Descriptor instead.
func (*FilesResponse) Descriptor() ([]byte, []int) {
	return file_telegraf_proto_rawDescGZIP(), []int{4}
}

func (x *FilesResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FilesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FilesResponse) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegraf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_telegraf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_telegraf_proto_rawDescGZIP(), []int{5}
}

func (x *FileInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_telegraf_proto protoreflect.FileDescriptor

var file_telegraf_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x74, 0x66, 0x61, 0x70, 0x69, 0x22, 0x21, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x3e, 0x0a, 0x0e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x0e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a, 0x0f, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x64, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x66, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x1e, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x32, 0xd6, 0x02, 0x0a, 0x08,
	0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x66, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x74, 0x66, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x66, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x12, 0x12, 0x2e, 0x74, 0x66,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x74, 0x66, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x74, 0x66, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x66, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x05, 0x54, 0x6f, 0x75, 0x63, 0x68, 0x12, 0x12, 0x2e, 0x74, 0x66, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x74, 0x66, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x74, 0x66, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x66, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x05, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x74, 0x66, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x66,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x74, 0x66, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telegraf_proto_rawDescOnce sync.Once
	file_telegraf_proto_rawDescData = file_telegraf_proto_rawDesc
)

func file_telegraf_proto_rawDescGZIP() []byte {
	file_telegraf_proto_rawDescOnce.Do(func() {
		file_telegraf_proto_rawDescData = protoimpl.X.CompressGZIP(file_telegraf_proto_rawDescData)
	})
	return file_telegraf_proto_rawDescData
}

var file_telegraf_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_telegraf_proto_goTypes = []interface{}{
	(*PathRequest)(nil),     // 0: tfapi.PathRequest
	(*ContentRequest)(nil),  // 1: tfapi.ContentRequest
	(*CommonResponse)(nil),  // 2: tfapi.CommonResponse
	(*ContentResponse)(nil), // 3: tfapi.ContentResponse
	(*FilesResponse)(nil),   // 4: tfapi.FilesResponse
	(*FileInfo)(nil),        // 5: tfapi.FileInfo
}
var file_telegraf_proto_depIdxs = []int32{
	5, // 0: tfapi.FilesResponse.files:type_name -> tfapi.FileInfo
	0, // 1: tfapi.Telegraf.GetFile:input_type -> tfapi.PathRequest
	0, // 2: tfapi.Telegraf.GetDir:input_type -> tfapi.PathRequest
	1, // 3: tfapi.Telegraf.Update:input_type -> tfapi.ContentRequest
	0, // 4: tfapi.Telegraf.Touch:input_type -> tfapi.PathRequest
	0, // 5: tfapi.Telegraf.Delete:input_type -> tfapi.PathRequest
	0, // 6: tfapi.Telegraf.Exist:input_type -> tfapi.PathRequest
	3, // 7: tfapi.Telegraf.GetFile:output_type -> tfapi.ContentResponse
	4, // 8: tfapi.Telegraf.GetDir:output_type -> tfapi.FilesResponse
	2, // 9: tfapi.Telegraf.Update:output_type -> tfapi.CommonResponse
	2, // 10: tfapi.Telegraf.Touch:output_type -> tfapi.CommonResponse
	2, // 11: tfapi.Telegraf.Delete:output_type -> tfapi.CommonResponse
	2, // 12: tfapi.Telegraf.Exist:output_type -> tfapi.CommonResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_telegraf_proto_init() }
func file_telegraf_proto_init() {
	if File_telegraf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_telegraf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathRequest); i {
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
		file_telegraf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentRequest); i {
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
		file_telegraf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
		file_telegraf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentResponse); i {
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
		file_telegraf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesResponse); i {
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
		file_telegraf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
			RawDescriptor: file_telegraf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_telegraf_proto_goTypes,
		DependencyIndexes: file_telegraf_proto_depIdxs,
		MessageInfos:      file_telegraf_proto_msgTypes,
	}.Build()
	File_telegraf_proto = out.File
	file_telegraf_proto_rawDesc = nil
	file_telegraf_proto_goTypes = nil
	file_telegraf_proto_depIdxs = nil
}
