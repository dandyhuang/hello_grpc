// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: rtrs_abtest.proto

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

type AbtestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sceneid uint64 `protobuf:"varint,1,opt,name=sceneid,proto3" json:"sceneid,omitempty"` //场景id
	Userid  string `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid,omitempty"`    //用于查找白名单
	Hashid  string `protobuf:"bytes,3,opt,name=hashid,proto3" json:"hashid,omitempty"`    //用于计算哈希值
}

func (x *AbtestReq) Reset() {
	*x = AbtestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtrs_abtest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbtestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbtestReq) ProtoMessage() {}

func (x *AbtestReq) ProtoReflect() protoreflect.Message {
	mi := &file_rtrs_abtest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbtestReq.ProtoReflect.Descriptor instead.
func (*AbtestReq) Descriptor() ([]byte, []int) {
	return file_rtrs_abtest_proto_rawDescGZIP(), []int{0}
}

func (x *AbtestReq) GetSceneid() uint64 {
	if x != nil {
		return x.Sceneid
	}
	return 0
}

func (x *AbtestReq) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *AbtestReq) GetHashid() string {
	if x != nil {
		return x.Hashid
	}
	return ""
}

type AbtestRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageItems  []*AlgMultiTestInfo `protobuf:"bytes,1,rep,name=page_items,json=pageItems,proto3" json:"page_items,omitempty"`
	DebugError string              `protobuf:"bytes,2,opt,name=debug_error,json=debugError,proto3" json:"debug_error,omitempty"`
}

func (x *AbtestRsp) Reset() {
	*x = AbtestRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtrs_abtest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbtestRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbtestRsp) ProtoMessage() {}

func (x *AbtestRsp) ProtoReflect() protoreflect.Message {
	mi := &file_rtrs_abtest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbtestRsp.ProtoReflect.Descriptor instead.
func (*AbtestRsp) Descriptor() ([]byte, []int) {
	return file_rtrs_abtest_proto_rawDescGZIP(), []int{1}
}

func (x *AbtestRsp) GetPageItems() []*AlgMultiTestInfo {
	if x != nil {
		return x.PageItems
	}
	return nil
}

func (x *AbtestRsp) GetDebugError() string {
	if x != nil {
		return x.DebugError
	}
	return ""
}

var File_rtrs_abtest_proto protoreflect.FileDescriptor

var file_rtrs_abtest_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x74, 0x72, 0x73, 0x5f, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x6f, 0x5f, 0x72, 0x74, 0x72, 0x73, 0x1a, 0x12, 0x72, 0x74,
	0x72, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x55, 0x0a, 0x09, 0x41, 0x62, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x68, 0x61, 0x73, 0x68, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x09, 0x41, 0x62, 0x74, 0x65, 0x73,
	0x74, 0x52, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6f, 0x5f, 0x72, 0x74,
	0x72, 0x73, 0x2e, 0x41, 0x6c, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x65, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rtrs_abtest_proto_rawDescOnce sync.Once
	file_rtrs_abtest_proto_rawDescData = file_rtrs_abtest_proto_rawDesc
)

func file_rtrs_abtest_proto_rawDescGZIP() []byte {
	file_rtrs_abtest_proto_rawDescOnce.Do(func() {
		file_rtrs_abtest_proto_rawDescData = protoimpl.X.CompressGZIP(file_rtrs_abtest_proto_rawDescData)
	})
	return file_rtrs_abtest_proto_rawDescData
}

var file_rtrs_abtest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rtrs_abtest_proto_goTypes = []interface{}{
	(*AbtestReq)(nil),        // 0: so_rtrs.AbtestReq
	(*AbtestRsp)(nil),        // 1: so_rtrs.AbtestRsp
	(*AlgMultiTestInfo)(nil), // 2: so_rtrs.AlgMultiTestInfo
}
var file_rtrs_abtest_proto_depIdxs = []int32{
	2, // 0: so_rtrs.AbtestRsp.page_items:type_name -> so_rtrs.AlgMultiTestInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rtrs_abtest_proto_init() }
func file_rtrs_abtest_proto_init() {
	if File_rtrs_abtest_proto != nil {
		return
	}
	file_rtrs_expinfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rtrs_abtest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbtestReq); i {
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
		file_rtrs_abtest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbtestRsp); i {
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
			RawDescriptor: file_rtrs_abtest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rtrs_abtest_proto_goTypes,
		DependencyIndexes: file_rtrs_abtest_proto_depIdxs,
		MessageInfos:      file_rtrs_abtest_proto_msgTypes,
	}.Build()
	File_rtrs_abtest_proto = out.File
	file_rtrs_abtest_proto_rawDesc = nil
	file_rtrs_abtest_proto_goTypes = nil
	file_rtrs_abtest_proto_depIdxs = nil
}
