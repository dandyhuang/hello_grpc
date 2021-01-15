// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: rtrs_schedule.proto

package so_rtrs

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ParamKV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *ParamKV) Reset() {
	*x = ParamKV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtrs_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamKV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamKV) ProtoMessage() {}

func (x *ParamKV) ProtoReflect() protoreflect.Message {
	mi := &file_rtrs_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamKV.ProtoReflect.Descriptor instead.
func (*ParamKV) Descriptor() ([]byte, []int) {
	return file_rtrs_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *ParamKV) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ParamKV) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri    string     `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Query  []*ParamKV `protobuf:"bytes,2,rep,name=query,proto3" json:"query,omitempty"`
	Header []*ParamKV `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty"`
	Cookie []*ParamKV `protobuf:"bytes,4,rep,name=cookie,proto3" json:"cookie,omitempty"`
	Body   []byte     `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtrs_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_rtrs_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_rtrs_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Request) GetQuery() []*ParamKV {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *Request) GetHeader() []*ParamKV {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Request) GetCookie() []*ParamKV {
	if x != nil {
		return x.Cookie
	}
	return nil
}

func (x *Request) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mime   string     `protobuf:"bytes,1,opt,name=mime,proto3" json:"mime,omitempty"`
	Header []*ParamKV `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty"`
	Code   uint32     `protobuf:"fixed32,3,opt,name=code,proto3" json:"code,omitempty"`
	Body   []byte     `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtrs_schedule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_rtrs_schedule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_rtrs_schedule_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *Response) GetHeader() []*ParamKV {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Response) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_rtrs_schedule_proto protoreflect.FileDescriptor

var file_rtrs_schedule_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x74, 0x72, 0x73, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6f, 0x5f, 0x72, 0x74, 0x72, 0x73, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x4b, 0x56, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0xc6, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x12, 0x2f, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6f, 0x5f, 0x72, 0x74, 0x72, 0x73, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4b, 0x56, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6f, 0x5f, 0x72, 0x74, 0x72, 0x73, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4b, 0x56,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6f, 0x5f, 0x72, 0x74,
	0x72, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x4b, 0x56, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22,
	0x79, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12,
	0x31, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x6f, 0x5f, 0x72, 0x74, 0x72, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4b, 0x56, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x07,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x73, 0x6f,
	0x5f, 0x72, 0x74, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rtrs_schedule_proto_rawDescOnce sync.Once
	file_rtrs_schedule_proto_rawDescData = file_rtrs_schedule_proto_rawDesc
)

func file_rtrs_schedule_proto_rawDescGZIP() []byte {
	file_rtrs_schedule_proto_rawDescOnce.Do(func() {
		file_rtrs_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_rtrs_schedule_proto_rawDescData)
	})
	return file_rtrs_schedule_proto_rawDescData
}

var file_rtrs_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rtrs_schedule_proto_goTypes = []interface{}{
	(*ParamKV)(nil),  // 0: so_rtrs.schedule.ParamKV
	(*Request)(nil),  // 1: so_rtrs.schedule.Request
	(*Response)(nil), // 2: so_rtrs.schedule.Response
}
var file_rtrs_schedule_proto_depIdxs = []int32{
	0, // 0: so_rtrs.schedule.Request.query:type_name -> so_rtrs.schedule.ParamKV
	0, // 1: so_rtrs.schedule.Request.header:type_name -> so_rtrs.schedule.ParamKV
	0, // 2: so_rtrs.schedule.Request.cookie:type_name -> so_rtrs.schedule.ParamKV
	0, // 3: so_rtrs.schedule.Response.header:type_name -> so_rtrs.schedule.ParamKV
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rtrs_schedule_proto_init() }
func file_rtrs_schedule_proto_init() {
	if File_rtrs_schedule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rtrs_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamKV); i {
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
		file_rtrs_schedule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_rtrs_schedule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_rtrs_schedule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rtrs_schedule_proto_goTypes,
		DependencyIndexes: file_rtrs_schedule_proto_depIdxs,
		MessageInfos:      file_rtrs_schedule_proto_msgTypes,
	}.Build()
	File_rtrs_schedule_proto = out.File
	file_rtrs_schedule_proto_rawDesc = nil
	file_rtrs_schedule_proto_goTypes = nil
	file_rtrs_schedule_proto_depIdxs = nil
}
