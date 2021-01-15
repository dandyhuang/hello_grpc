// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: rtrs_expinfo.proto

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

//目的：支持多层实验平台，通过配置化方式实现多组实验统计
//代表流量经过该层，所附带的算法参数和系统参数
type AlgMultiTestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestID int32 `protobuf:"varint,1,opt,name=testID,proto3" json:"testID,omitempty"` //暂未使用
	RuleID int32 `protobuf:"varint,2,opt,name=ruleID,proto3" json:"ruleID,omitempty"` //标识实验流量进入哪个so执行，灰度过程中ruleID 和实验参数中都填入
	//实验参数 由实验者和系统设定参数，各自识别并执行
	//算法参数:被相应的so识别并执行
	//系统参数:，IDC路由参数,参数
	//参数名，参数值 比如 ruleID="10"
	MapParams map[string]string `protobuf:"bytes,3,rep,name=mapParams,proto3" json:"mapParams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 模块层标识
	// 1 -> scoring
	// 2 -> sorting
	// 3 -> filter
	// R(ascii) -> retrieval
	// ....
	ExpLayer uint32 `protobuf:"varint,4,opt,name=expLayer,proto3" json:"expLayer,omitempty"`
	ExpName  string `protobuf:"bytes,5,opt,name=expName,proto3" json:"expName,omitempty"` // 实验名称
	Pageid   uint64 `protobuf:"varint,6,opt,name=pageid,proto3" json:"pageid,omitempty"`
}

func (x *AlgMultiTestInfo) Reset() {
	*x = AlgMultiTestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtrs_expinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgMultiTestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgMultiTestInfo) ProtoMessage() {}

func (x *AlgMultiTestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rtrs_expinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgMultiTestInfo.ProtoReflect.Descriptor instead.
func (*AlgMultiTestInfo) Descriptor() ([]byte, []int) {
	return file_rtrs_expinfo_proto_rawDescGZIP(), []int{0}
}

func (x *AlgMultiTestInfo) GetTestID() int32 {
	if x != nil {
		return x.TestID
	}
	return 0
}

func (x *AlgMultiTestInfo) GetRuleID() int32 {
	if x != nil {
		return x.RuleID
	}
	return 0
}

func (x *AlgMultiTestInfo) GetMapParams() map[string]string {
	if x != nil {
		return x.MapParams
	}
	return nil
}

func (x *AlgMultiTestInfo) GetExpLayer() uint32 {
	if x != nil {
		return x.ExpLayer
	}
	return 0
}

func (x *AlgMultiTestInfo) GetExpName() string {
	if x != nil {
		return x.ExpName
	}
	return ""
}

func (x *AlgMultiTestInfo) GetPageid() uint64 {
	if x != nil {
		return x.Pageid
	}
	return 0
}

var File_rtrs_expinfo_proto protoreflect.FileDescriptor

var file_rtrs_expinfo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x74, 0x72, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x6f, 0x5f, 0x72, 0x74, 0x72, 0x73, 0x22, 0x96, 0x02,
	0x0a, 0x10, 0x41, 0x6c, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x75,
	0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65,
	0x49, 0x44, 0x12, 0x46, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x6f, 0x5f, 0x72, 0x74, 0x72, 0x73, 0x2e,
	0x41, 0x6c, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x4d, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x09, 0x6d, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78,
	0x70, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x78,
	0x70, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x69, 0x64, 0x1a, 0x3c, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rtrs_expinfo_proto_rawDescOnce sync.Once
	file_rtrs_expinfo_proto_rawDescData = file_rtrs_expinfo_proto_rawDesc
)

func file_rtrs_expinfo_proto_rawDescGZIP() []byte {
	file_rtrs_expinfo_proto_rawDescOnce.Do(func() {
		file_rtrs_expinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_rtrs_expinfo_proto_rawDescData)
	})
	return file_rtrs_expinfo_proto_rawDescData
}

var file_rtrs_expinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rtrs_expinfo_proto_goTypes = []interface{}{
	(*AlgMultiTestInfo)(nil), // 0: so_rtrs.AlgMultiTestInfo
	nil,                      // 1: so_rtrs.AlgMultiTestInfo.MapParamsEntry
}
var file_rtrs_expinfo_proto_depIdxs = []int32{
	1, // 0: so_rtrs.AlgMultiTestInfo.mapParams:type_name -> so_rtrs.AlgMultiTestInfo.MapParamsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rtrs_expinfo_proto_init() }
func file_rtrs_expinfo_proto_init() {
	if File_rtrs_expinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rtrs_expinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgMultiTestInfo); i {
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
			RawDescriptor: file_rtrs_expinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rtrs_expinfo_proto_goTypes,
		DependencyIndexes: file_rtrs_expinfo_proto_depIdxs,
		MessageInfos:      file_rtrs_expinfo_proto_msgTypes,
	}.Build()
	File_rtrs_expinfo_proto = out.File
	file_rtrs_expinfo_proto_rawDesc = nil
	file_rtrs_expinfo_proto_goTypes = nil
	file_rtrs_expinfo_proto_depIdxs = nil
}
