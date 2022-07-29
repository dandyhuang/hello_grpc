// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.2
// source: protocol/abt/rtrs_abtest.proto

package abt

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
	PageName string `protobuf:"bytes,7,opt,name=pageName,proto3" json:"pageName,omitempty"`
}

func (x *AlgMultiTestInfo) Reset() {
	*x = AlgMultiTestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_abt_rtrs_abtest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgMultiTestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgMultiTestInfo) ProtoMessage() {}

func (x *AlgMultiTestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_abt_rtrs_abtest_proto_msgTypes[0]
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
	return file_protocol_abt_rtrs_abtest_proto_rawDescGZIP(), []int{0}
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

func (x *AlgMultiTestInfo) GetPageName() string {
	if x != nil {
		return x.PageName
	}
	return ""
}

type AllTestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VecTestInfo []*AlgMultiTestInfo `protobuf:"bytes,1,rep,name=vecTestInfo,proto3" json:"vecTestInfo,omitempty"`
	DebugError  string              `protobuf:"bytes,2,opt,name=debug_error,json=debugError,proto3" json:"debug_error,omitempty"`
}

func (x *AllTestInfo) Reset() {
	*x = AllTestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_abt_rtrs_abtest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllTestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllTestInfo) ProtoMessage() {}

func (x *AllTestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_abt_rtrs_abtest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllTestInfo.ProtoReflect.Descriptor instead.
func (*AllTestInfo) Descriptor() ([]byte, []int) {
	return file_protocol_abt_rtrs_abtest_proto_rawDescGZIP(), []int{1}
}

func (x *AllTestInfo) GetVecTestInfo() []*AlgMultiTestInfo {
	if x != nil {
		return x.VecTestInfo
	}
	return nil
}

func (x *AllTestInfo) GetDebugError() string {
	if x != nil {
		return x.DebugError
	}
	return ""
}

type AbtestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sceneid       uint64 `protobuf:"varint,1,opt,name=sceneid,proto3" json:"sceneid,omitempty"`                                 //场景id
	Userid        string `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid,omitempty"`                                    //用于查找白名单
	Hashid        string `protobuf:"bytes,3,opt,name=hashid,proto3" json:"hashid,omitempty"`                                    //用于计算哈希值
	PageRecognize string `protobuf:"bytes,4,opt,name=page_recognize,json=pageRecognize,proto3" json:"page_recognize,omitempty"` //指定场景查找结点
}

func (x *AbtestReq) Reset() {
	*x = AbtestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_abt_rtrs_abtest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbtestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbtestReq) ProtoMessage() {}

func (x *AbtestReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_abt_rtrs_abtest_proto_msgTypes[2]
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
	return file_protocol_abt_rtrs_abtest_proto_rawDescGZIP(), []int{2}
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

func (x *AbtestReq) GetPageRecognize() string {
	if x != nil {
		return x.PageRecognize
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
		mi := &file_protocol_abt_rtrs_abtest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbtestRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbtestRsp) ProtoMessage() {}

func (x *AbtestRsp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_abt_rtrs_abtest_proto_msgTypes[3]
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
	return file_protocol_abt_rtrs_abtest_proto_rawDescGZIP(), []int{3}
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

var File_protocol_abt_rtrs_abtest_proto protoreflect.FileDescriptor

var file_protocol_abt_rtrs_abtest_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x62, 0x74, 0x2f, 0x72,
	0x74, 0x72, 0x73, 0x5f, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x62, 0x74, 0x22, 0xae, 0x02, 0x0a, 0x10, 0x41, 0x6c, 0x67, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65,
	0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x42, 0x0a, 0x09, 0x6d, 0x61,
	0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x61, 0x62, 0x74, 0x2e, 0x41, 0x6c, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x78, 0x70, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x65, 0x78, 0x70, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x3c, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x67, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x0b, 0x76, 0x65, 0x63, 0x54, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x62, 0x74,
	0x2e, 0x41, 0x6c, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x7c, 0x0a, 0x09, 0x41, 0x62, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x22, 0x62, 0x0a,
	0x09, 0x41, 0x62, 0x74, 0x65, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x61, 0x62, 0x74, 0x2e, 0x41, 0x6c, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x67, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x42, 0x22, 0x5a, 0x20, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x62, 0x74,
	0x2f, 0x3b, 0x61, 0x62, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_abt_rtrs_abtest_proto_rawDescOnce sync.Once
	file_protocol_abt_rtrs_abtest_proto_rawDescData = file_protocol_abt_rtrs_abtest_proto_rawDesc
)

func file_protocol_abt_rtrs_abtest_proto_rawDescGZIP() []byte {
	file_protocol_abt_rtrs_abtest_proto_rawDescOnce.Do(func() {
		file_protocol_abt_rtrs_abtest_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_abt_rtrs_abtest_proto_rawDescData)
	})
	return file_protocol_abt_rtrs_abtest_proto_rawDescData
}

var file_protocol_abt_rtrs_abtest_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protocol_abt_rtrs_abtest_proto_goTypes = []interface{}{
	(*AlgMultiTestInfo)(nil), // 0: abt.AlgMultiTestInfo
	(*AllTestInfo)(nil),      // 1: abt.AllTestInfo
	(*AbtestReq)(nil),        // 2: abt.AbtestReq
	(*AbtestRsp)(nil),        // 3: abt.AbtestRsp
	nil,                      // 4: abt.AlgMultiTestInfo.MapParamsEntry
}
var file_protocol_abt_rtrs_abtest_proto_depIdxs = []int32{
	4, // 0: abt.AlgMultiTestInfo.mapParams:type_name -> abt.AlgMultiTestInfo.MapParamsEntry
	0, // 1: abt.AllTestInfo.vecTestInfo:type_name -> abt.AlgMultiTestInfo
	0, // 2: abt.AbtestRsp.page_items:type_name -> abt.AlgMultiTestInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protocol_abt_rtrs_abtest_proto_init() }
func file_protocol_abt_rtrs_abtest_proto_init() {
	if File_protocol_abt_rtrs_abtest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_abt_rtrs_abtest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_protocol_abt_rtrs_abtest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllTestInfo); i {
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
		file_protocol_abt_rtrs_abtest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_protocol_abt_rtrs_abtest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_protocol_abt_rtrs_abtest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_abt_rtrs_abtest_proto_goTypes,
		DependencyIndexes: file_protocol_abt_rtrs_abtest_proto_depIdxs,
		MessageInfos:      file_protocol_abt_rtrs_abtest_proto_msgTypes,
	}.Build()
	File_protocol_abt_rtrs_abtest_proto = out.File
	file_protocol_abt_rtrs_abtest_proto_rawDesc = nil
	file_protocol_abt_rtrs_abtest_proto_goTypes = nil
	file_protocol_abt_rtrs_abtest_proto_depIdxs = nil
}
