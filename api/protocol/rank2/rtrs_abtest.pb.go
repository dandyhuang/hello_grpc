// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.2
// source: protocol/rank2/rtrs_abtest.proto

package rank2

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

type AbtestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pageid uint64 `protobuf:"varint,1,opt,name=pageid,proto3" json:"pageid,omitempty"`
	Imei   string `protobuf:"bytes,2,opt,name=imei,proto3" json:"imei,omitempty"`
}

func (x *AbtestReq) Reset() {
	*x = AbtestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbtestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbtestReq) ProtoMessage() {}

func (x *AbtestReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[0]
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
	return file_protocol_rank2_rtrs_abtest_proto_rawDescGZIP(), []int{0}
}

func (x *AbtestReq) GetPageid() uint64 {
	if x != nil {
		return x.Pageid
	}
	return 0
}

func (x *AbtestReq) GetImei() string {
	if x != nil {
		return x.Imei
	}
	return ""
}

type AbtestV2Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sceneid       uint64 `protobuf:"varint,1,opt,name=sceneid,proto3" json:"sceneid,omitempty"`                                 //场景id
	Userid        string `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid,omitempty"`                                    //用于查找白名单
	Hashid        string `protobuf:"bytes,3,opt,name=hashid,proto3" json:"hashid,omitempty"`                                    //用于计算哈希值
	PageRecognize string `protobuf:"bytes,4,opt,name=page_recognize,json=pageRecognize,proto3" json:"page_recognize,omitempty"` //指定场景查找结点
}

func (x *AbtestV2Req) Reset() {
	*x = AbtestV2Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbtestV2Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbtestV2Req) ProtoMessage() {}

func (x *AbtestV2Req) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbtestV2Req.ProtoReflect.Descriptor instead.
func (*AbtestV2Req) Descriptor() ([]byte, []int) {
	return file_protocol_rank2_rtrs_abtest_proto_rawDescGZIP(), []int{1}
}

func (x *AbtestV2Req) GetSceneid() uint64 {
	if x != nil {
		return x.Sceneid
	}
	return 0
}

func (x *AbtestV2Req) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *AbtestV2Req) GetHashid() string {
	if x != nil {
		return x.Hashid
	}
	return ""
}

func (x *AbtestV2Req) GetPageRecognize() string {
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
		mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbtestRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbtestRsp) ProtoMessage() {}

func (x *AbtestRsp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[2]
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
	return file_protocol_rank2_rtrs_abtest_proto_rawDescGZIP(), []int{2}
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

type AbNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BHead         bool              `protobuf:"varint,1,opt,name=bHead,proto3" json:"bHead,omitempty"` // 是否根节点
	BHashMix      bool              `protobuf:"varint,2,opt,name=bHashMix,proto3" json:"bHashMix,omitempty"`
	BHashDate     bool              `protobuf:"varint,3,opt,name=bHashDate,proto3" json:"bHashDate,omitempty"`
	Pageid        uint64            `protobuf:"varint,4,opt,name=pageid,proto3" json:"pageid,omitempty"`   // 结点ID
	Sceneid       uint64            `protobuf:"varint,5,opt,name=sceneid,proto3" json:"sceneid,omitempty"` // 场景
	Version       int32             `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	AbDefPage     uint64            `protobuf:"varint,7,opt,name=abDefPage,proto3" json:"abDefPage,omitempty"`
	ChannelSize   uint64            `protobuf:"varint,8,opt,name=channelSize,proto3" json:"channelSize,omitempty"` // hash
	HashFix       uint64            `protobuf:"varint,9,opt,name=hashFix,proto3" json:"hashFix,omitempty"`
	HashDateDelay int64             `protobuf:"varint,10,opt,name=hashDateDelay,proto3" json:"hashDateDelay,omitempty"`
	NodeType      int32             `protobuf:"varint,11,opt,name=nodeType,proto3" json:"nodeType,omitempty"`           // 结点类型
	NodeAlgorithm int32             `protobuf:"varint,12,opt,name=nodeAlgorithm,proto3" json:"nodeAlgorithm,omitempty"` // hash算法类型
	JsName        string            `protobuf:"bytes,13,opt,name=jsName,proto3" json:"jsName,omitempty"`
	AbChannels    []string          `protobuf:"bytes,14,rep,name=abChannels,proto3" json:"abChannels,omitempty"`                                                                                       // hash(***) [0,20] -> A,  abChannels[0] = "A"  abChannels[19] = "A"
	SubNodes      map[string]uint64 `protobuf:"bytes,15,rep,name=subNodes,proto3" json:"subNodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`  // 子节点名--》pageID映射
	ExpParams     map[string]string `protobuf:"bytes,16,rep,name=expParams,proto3" json:"expParams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 配置的参数
	FixUser       map[string]uint64 `protobuf:"bytes,17,rep,name=fixUser,proto3" json:"fixUser,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`    // 白名单imei-> 根节点映射
	Recognize     []string          `protobuf:"bytes,18,rep,name=recognize,proto3" json:"recognize,omitempty"`                                                                                         // 该结点流量识别标识
}

func (x *AbNodeInfo) Reset() {
	*x = AbNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbNodeInfo) ProtoMessage() {}

func (x *AbNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbNodeInfo.ProtoReflect.Descriptor instead.
func (*AbNodeInfo) Descriptor() ([]byte, []int) {
	return file_protocol_rank2_rtrs_abtest_proto_rawDescGZIP(), []int{3}
}

func (x *AbNodeInfo) GetBHead() bool {
	if x != nil {
		return x.BHead
	}
	return false
}

func (x *AbNodeInfo) GetBHashMix() bool {
	if x != nil {
		return x.BHashMix
	}
	return false
}

func (x *AbNodeInfo) GetBHashDate() bool {
	if x != nil {
		return x.BHashDate
	}
	return false
}

func (x *AbNodeInfo) GetPageid() uint64 {
	if x != nil {
		return x.Pageid
	}
	return 0
}

func (x *AbNodeInfo) GetSceneid() uint64 {
	if x != nil {
		return x.Sceneid
	}
	return 0
}

func (x *AbNodeInfo) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *AbNodeInfo) GetAbDefPage() uint64 {
	if x != nil {
		return x.AbDefPage
	}
	return 0
}

func (x *AbNodeInfo) GetChannelSize() uint64 {
	if x != nil {
		return x.ChannelSize
	}
	return 0
}

func (x *AbNodeInfo) GetHashFix() uint64 {
	if x != nil {
		return x.HashFix
	}
	return 0
}

func (x *AbNodeInfo) GetHashDateDelay() int64 {
	if x != nil {
		return x.HashDateDelay
	}
	return 0
}

func (x *AbNodeInfo) GetNodeType() int32 {
	if x != nil {
		return x.NodeType
	}
	return 0
}

func (x *AbNodeInfo) GetNodeAlgorithm() int32 {
	if x != nil {
		return x.NodeAlgorithm
	}
	return 0
}

func (x *AbNodeInfo) GetJsName() string {
	if x != nil {
		return x.JsName
	}
	return ""
}

func (x *AbNodeInfo) GetAbChannels() []string {
	if x != nil {
		return x.AbChannels
	}
	return nil
}

func (x *AbNodeInfo) GetSubNodes() map[string]uint64 {
	if x != nil {
		return x.SubNodes
	}
	return nil
}

func (x *AbNodeInfo) GetExpParams() map[string]string {
	if x != nil {
		return x.ExpParams
	}
	return nil
}

func (x *AbNodeInfo) GetFixUser() map[string]uint64 {
	if x != nil {
		return x.FixUser
	}
	return nil
}

func (x *AbNodeInfo) GetRecognize() []string {
	if x != nil {
		return x.Recognize
	}
	return nil
}

type AbNodeTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id2NodeInfo map[uint64]*AbNodeInfo `protobuf:"bytes,1,rep,name=id2NodeInfo,proto3" json:"id2NodeInfo,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Sceneid2Id  map[uint64]uint64      `protobuf:"bytes,2,rep,name=sceneid2Id,proto3" json:"sceneid2Id,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Recog2Id    map[string]uint64      `protobuf:"bytes,3,rep,name=recog2Id,proto3" json:"recog2Id,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *AbNodeTree) Reset() {
	*x = AbNodeTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbNodeTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbNodeTree) ProtoMessage() {}

func (x *AbNodeTree) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbNodeTree.ProtoReflect.Descriptor instead.
func (*AbNodeTree) Descriptor() ([]byte, []int) {
	return file_protocol_rank2_rtrs_abtest_proto_rawDescGZIP(), []int{4}
}

func (x *AbNodeTree) GetId2NodeInfo() map[uint64]*AbNodeInfo {
	if x != nil {
		return x.Id2NodeInfo
	}
	return nil
}

func (x *AbNodeTree) GetSceneid2Id() map[uint64]uint64 {
	if x != nil {
		return x.Sceneid2Id
	}
	return nil
}

func (x *AbNodeTree) GetRecog2Id() map[string]uint64 {
	if x != nil {
		return x.Recog2Id
	}
	return nil
}

type AbtReqNodeTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sceneid uint64 `protobuf:"varint,1,opt,name=sceneid,proto3" json:"sceneid,omitempty"` //场景id
}

func (x *AbtReqNodeTree) Reset() {
	*x = AbtReqNodeTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbtReqNodeTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbtReqNodeTree) ProtoMessage() {}

func (x *AbtReqNodeTree) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbtReqNodeTree.ProtoReflect.Descriptor instead.
func (*AbtReqNodeTree) Descriptor() ([]byte, []int) {
	return file_protocol_rank2_rtrs_abtest_proto_rawDescGZIP(), []int{5}
}

func (x *AbtReqNodeTree) GetSceneid() uint64 {
	if x != nil {
		return x.Sceneid
	}
	return 0
}

type AbtRspNodeTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeTree   *AbNodeTree `protobuf:"bytes,1,opt,name=nodeTree,proto3" json:"nodeTree,omitempty"`
	DebugError string      `protobuf:"bytes,2,opt,name=debug_error,json=debugError,proto3" json:"debug_error,omitempty"`
}

func (x *AbtRspNodeTree) Reset() {
	*x = AbtRspNodeTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbtRspNodeTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbtRspNodeTree) ProtoMessage() {}

func (x *AbtRspNodeTree) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_rank2_rtrs_abtest_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbtRspNodeTree.ProtoReflect.Descriptor instead.
func (*AbtRspNodeTree) Descriptor() ([]byte, []int) {
	return file_protocol_rank2_rtrs_abtest_proto_rawDescGZIP(), []int{6}
}

func (x *AbtRspNodeTree) GetNodeTree() *AbNodeTree {
	if x != nil {
		return x.NodeTree
	}
	return nil
}

func (x *AbtRspNodeTree) GetDebugError() string {
	if x != nil {
		return x.DebugError
	}
	return ""
}

var File_protocol_rank2_rtrs_abtest_proto protoreflect.FileDescriptor

var file_protocol_rank2_rtrs_abtest_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x32,
	0x2f, 0x72, 0x74, 0x72, 0x73, 0x5f, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x72, 0x74, 0x72, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x1a, 0x21,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x32, 0x2f, 0x72,
	0x74, 0x72, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x37, 0x0a, 0x09, 0x41, 0x62, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x70, 0x61, 0x67, 0x65, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x22, 0x7e, 0x0a, 0x0b, 0x41, 0x62,
	0x74, 0x65, 0x73, 0x74, 0x56, 0x32, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x22, 0x69, 0x0a, 0x09, 0x41, 0x62,
	0x74, 0x65, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x74,
	0x72, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x41, 0x6c, 0x67, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xbd, 0x06, 0x0a, 0x0a, 0x41, 0x62, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x48, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x62, 0x48, 0x65, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x48,
	0x61, 0x73, 0x68, 0x4d, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x48,
	0x61, 0x73, 0x68, 0x4d, 0x69, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x48, 0x61, 0x73, 0x68, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x48, 0x61, 0x73, 0x68,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x62, 0x44, 0x65, 0x66, 0x50, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x62, 0x44, 0x65, 0x66, 0x50, 0x61, 0x67, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x68, 0x46, 0x69, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x68, 0x61, 0x73, 0x68, 0x46, 0x69, 0x78, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61,
	0x73, 0x68, 0x44, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x68, 0x44, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x61, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x6e, 0x6f, 0x64, 0x65, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6a, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x62,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x62, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x75,
	0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72,
	0x74, 0x72, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x41, 0x62, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x73, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x09,
	0x65, 0x78, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x72, 0x74, 0x72, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x41, 0x62, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x78, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x65, 0x78, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x3d, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x18, 0x11, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x74, 0x72, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e,
	0x41, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x46, 0x69, 0x78, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x66, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x1a, 0x3b,
	0x0a, 0x0d, 0x53, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x45,
	0x78, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x69, 0x78,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb5, 0x03, 0x0a, 0x0a, 0x41, 0x62, 0x4e, 0x6f, 0x64, 0x65,
	0x54, 0x72, 0x65, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x69, 0x64, 0x32, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x74, 0x72, 0x73,
	0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x41, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x65,
	0x65, 0x2e, 0x49, 0x64, 0x32, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x69, 0x64, 0x32, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x46, 0x0a, 0x0a, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x69, 0x64, 0x32, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x74, 0x72, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x2e, 0x41, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x65, 0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x69, 0x64, 0x32, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x69, 0x64, 0x32, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x67,
	0x32, 0x49, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x74, 0x72, 0x73,
	0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x41, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x65,
	0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x32, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x32, 0x49, 0x64, 0x1a, 0x56, 0x0a, 0x10, 0x49, 0x64, 0x32,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x72, 0x74, 0x72, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x41, 0x62, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x69, 0x64, 0x32, 0x49, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3b, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x32, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2a, 0x0a,
	0x0e, 0x41, 0x62, 0x74, 0x52, 0x65, 0x71, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x65, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x0e, 0x41, 0x62, 0x74,
	0x52, 0x73, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x65, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x54, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x72, 0x74, 0x72, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x41, 0x62, 0x4e, 0x6f, 0x64,
	0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x65, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x26, 0x5a, 0x24, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x61, 0x6e, 0x6b,
	0x32, 0x2f, 0x3b, 0x72, 0x61, 0x6e, 0x6b, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_rank2_rtrs_abtest_proto_rawDescOnce sync.Once
	file_protocol_rank2_rtrs_abtest_proto_rawDescData = file_protocol_rank2_rtrs_abtest_proto_rawDesc
)

func file_protocol_rank2_rtrs_abtest_proto_rawDescGZIP() []byte {
	file_protocol_rank2_rtrs_abtest_proto_rawDescOnce.Do(func() {
		file_protocol_rank2_rtrs_abtest_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_rank2_rtrs_abtest_proto_rawDescData)
	})
	return file_protocol_rank2_rtrs_abtest_proto_rawDescData
}

var file_protocol_rank2_rtrs_abtest_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_protocol_rank2_rtrs_abtest_proto_goTypes = []interface{}{
	(*AbtestReq)(nil),        // 0: rtrs_frame.AbtestReq
	(*AbtestV2Req)(nil),      // 1: rtrs_frame.AbtestV2Req
	(*AbtestRsp)(nil),        // 2: rtrs_frame.AbtestRsp
	(*AbNodeInfo)(nil),       // 3: rtrs_frame.AbNodeInfo
	(*AbNodeTree)(nil),       // 4: rtrs_frame.AbNodeTree
	(*AbtReqNodeTree)(nil),   // 5: rtrs_frame.AbtReqNodeTree
	(*AbtRspNodeTree)(nil),   // 6: rtrs_frame.AbtRspNodeTree
	nil,                      // 7: rtrs_frame.AbNodeInfo.SubNodesEntry
	nil,                      // 8: rtrs_frame.AbNodeInfo.ExpParamsEntry
	nil,                      // 9: rtrs_frame.AbNodeInfo.FixUserEntry
	nil,                      // 10: rtrs_frame.AbNodeTree.Id2NodeInfoEntry
	nil,                      // 11: rtrs_frame.AbNodeTree.Sceneid2IdEntry
	nil,                      // 12: rtrs_frame.AbNodeTree.Recog2IdEntry
	(*AlgMultiTestInfo)(nil), // 13: rtrs_frame.AlgMultiTestInfo
}
var file_protocol_rank2_rtrs_abtest_proto_depIdxs = []int32{
	13, // 0: rtrs_frame.AbtestRsp.page_items:type_name -> rtrs_frame.AlgMultiTestInfo
	7,  // 1: rtrs_frame.AbNodeInfo.subNodes:type_name -> rtrs_frame.AbNodeInfo.SubNodesEntry
	8,  // 2: rtrs_frame.AbNodeInfo.expParams:type_name -> rtrs_frame.AbNodeInfo.ExpParamsEntry
	9,  // 3: rtrs_frame.AbNodeInfo.fixUser:type_name -> rtrs_frame.AbNodeInfo.FixUserEntry
	10, // 4: rtrs_frame.AbNodeTree.id2NodeInfo:type_name -> rtrs_frame.AbNodeTree.Id2NodeInfoEntry
	11, // 5: rtrs_frame.AbNodeTree.sceneid2Id:type_name -> rtrs_frame.AbNodeTree.Sceneid2IdEntry
	12, // 6: rtrs_frame.AbNodeTree.recog2Id:type_name -> rtrs_frame.AbNodeTree.Recog2IdEntry
	4,  // 7: rtrs_frame.AbtRspNodeTree.nodeTree:type_name -> rtrs_frame.AbNodeTree
	3,  // 8: rtrs_frame.AbNodeTree.Id2NodeInfoEntry.value:type_name -> rtrs_frame.AbNodeInfo
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_protocol_rank2_rtrs_abtest_proto_init() }
func file_protocol_rank2_rtrs_abtest_proto_init() {
	if File_protocol_rank2_rtrs_abtest_proto != nil {
		return
	}
	file_protocol_rank2_rtrs_expinfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protocol_rank2_rtrs_abtest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_protocol_rank2_rtrs_abtest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbtestV2Req); i {
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
		file_protocol_rank2_rtrs_abtest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_protocol_rank2_rtrs_abtest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbNodeInfo); i {
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
		file_protocol_rank2_rtrs_abtest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbNodeTree); i {
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
		file_protocol_rank2_rtrs_abtest_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbtReqNodeTree); i {
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
		file_protocol_rank2_rtrs_abtest_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbtRspNodeTree); i {
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
			RawDescriptor: file_protocol_rank2_rtrs_abtest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_rank2_rtrs_abtest_proto_goTypes,
		DependencyIndexes: file_protocol_rank2_rtrs_abtest_proto_depIdxs,
		MessageInfos:      file_protocol_rank2_rtrs_abtest_proto_msgTypes,
	}.Build()
	File_protocol_rank2_rtrs_abtest_proto = out.File
	file_protocol_rank2_rtrs_abtest_proto_rawDesc = nil
	file_protocol_rank2_rtrs_abtest_proto_goTypes = nil
	file_protocol_rank2_rtrs_abtest_proto_depIdxs = nil
}
