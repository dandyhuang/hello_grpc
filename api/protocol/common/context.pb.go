// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.2
// source: protocol/common/context.proto

package common

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

// 广告位信息
type PositionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId int64 `protobuf:"varint,1,opt,name=positionId,proto3" json:"positionId,omitempty"` // 广告位ID(int64)
	BidFloor   int64 `protobuf:"varint,2,opt,name=bidFloor,proto3" json:"bidFloor,omitempty"`     // 广告位底价
}

func (x *PositionInfo) Reset() {
	*x = PositionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_common_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionInfo) ProtoMessage() {}

func (x *PositionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_common_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionInfo.ProtoReflect.Descriptor instead.
func (*PositionInfo) Descriptor() ([]byte, []int) {
	return file_protocol_common_context_proto_rawDescGZIP(), []int{0}
}

func (x *PositionInfo) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *PositionInfo) GetBidFloor() int64 {
	if x != nil {
		return x.BidFloor
	}
	return 0
}

// 广告上下文信息
type AdContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MultiPositionInfo []*PositionInfo   `protobuf:"bytes,1,rep,name=multiPositionInfo,proto3" json:"multiPositionInfo,omitempty"` // 位次信息
	RefreshTimes      int32             `protobuf:"varint,2,opt,name=refreshTimes,proto3" json:"refreshTimes,omitempty"`
	SspMediaType      int32             `protobuf:"varint,3,opt,name=ssp_media_type,json=sspMediaType,proto3" json:"ssp_media_type,omitempty"`                                                              // add by luojianhui 20200715, 1表示自有流量，2表示联盟开发者平台的媒体流量，3表示联盟ssp流量
	MapContext        map[string]string `protobuf:"bytes,4,rep,name=mapContext,proto3" json:"mapContext,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 推荐上下文信息
	AreaDldType       int32             `protobuf:"varint,5,opt,name=areaDldType,proto3" json:"areaDldType,omitempty"`                                                                                      // 1表示 自动下载；2表示手动下载
}

func (x *AdContext) Reset() {
	*x = AdContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_common_context_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdContext) ProtoMessage() {}

func (x *AdContext) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_common_context_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdContext.ProtoReflect.Descriptor instead.
func (*AdContext) Descriptor() ([]byte, []int) {
	return file_protocol_common_context_proto_rawDescGZIP(), []int{1}
}

func (x *AdContext) GetMultiPositionInfo() []*PositionInfo {
	if x != nil {
		return x.MultiPositionInfo
	}
	return nil
}

func (x *AdContext) GetRefreshTimes() int32 {
	if x != nil {
		return x.RefreshTimes
	}
	return 0
}

func (x *AdContext) GetSspMediaType() int32 {
	if x != nil {
		return x.SspMediaType
	}
	return 0
}

func (x *AdContext) GetMapContext() map[string]string {
	if x != nil {
		return x.MapContext
	}
	return nil
}

func (x *AdContext) GetAreaDldType() int32 {
	if x != nil {
		return x.AreaDldType
	}
	return 0
}

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId    uint32      `protobuf:"varint,1,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	AlgName    string      `protobuf:"bytes,2,opt,name=alg_name,json=algName,proto3" json:"alg_name,omitempty"`
	AlgGroup   string      `protobuf:"bytes,3,opt,name=alg_group,json=algGroup,proto3" json:"alg_group,omitempty"`
	AlgVersion string      `protobuf:"bytes,4,opt,name=alg_version,json=algVersion,proto3" json:"alg_version,omitempty"`
	ReqnUm     int32       `protobuf:"varint,5,opt,name=reqn_um,json=reqnUm,proto3" json:"reqn_um,omitempty"` // 请求中的广告数
	Ad         *AdContext  `protobuf:"bytes,60,opt,name=ad,proto3" json:"ad,omitempty"`
	CpcContext *CPCContext `protobuf:"bytes,61,opt,name=cpc_context,json=cpcContext,proto3" json:"cpc_context,omitempty"`
	CpdContext *CPDContext `protobuf:"bytes,62,opt,name=cpd_context,json=cpdContext,proto3" json:"cpd_context,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_common_context_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_common_context_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_protocol_common_context_proto_rawDescGZIP(), []int{2}
}

func (x *Context) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *Context) GetAlgName() string {
	if x != nil {
		return x.AlgName
	}
	return ""
}

func (x *Context) GetAlgGroup() string {
	if x != nil {
		return x.AlgGroup
	}
	return ""
}

func (x *Context) GetAlgVersion() string {
	if x != nil {
		return x.AlgVersion
	}
	return ""
}

func (x *Context) GetReqnUm() int32 {
	if x != nil {
		return x.ReqnUm
	}
	return 0
}

func (x *Context) GetAd() *AdContext {
	if x != nil {
		return x.Ad
	}
	return nil
}

func (x *Context) GetCpcContext() *CPCContext {
	if x != nil {
		return x.CpcContext
	}
	return nil
}

func (x *Context) GetCpdContext() *CPDContext {
	if x != nil {
		return x.CpdContext
	}
	return nil
}

var File_protocol_common_context_proto protoreflect.FileDescriptor

var file_protocol_common_context_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x70, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a,
	0x0c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x69, 0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x62, 0x69, 0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x22, 0xbd, 0x02, 0x0a, 0x09, 0x41, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x42, 0x0a, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12,
	0x24, 0x0a, 0x0e, 0x73, 0x73, 0x70, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x73, 0x70, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x6d, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x4d, 0x61, 0x70,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6d, 0x61,
	0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x72, 0x65, 0x61,
	0x44, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61,
	0x72, 0x65, 0x61, 0x44, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x3d, 0x0a, 0x0f, 0x4d, 0x61,
	0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa3, 0x02, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x6c, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x6c, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x67, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x6c, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x6e, 0x5f, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x71, 0x6e,
	0x55, 0x6d, 0x12, 0x21, 0x0a, 0x02, 0x61, 0x64, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x52, 0x02, 0x61, 0x64, 0x12, 0x33, 0x0a, 0x0b, 0x63, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x50, 0x43, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0a,
	0x63, 0x70, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x63, 0x70,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x3e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x50, 0x44, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x0a, 0x63, 0x70, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42,
	0x28, 0x5a, 0x26, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protocol_common_context_proto_rawDescOnce sync.Once
	file_protocol_common_context_proto_rawDescData = file_protocol_common_context_proto_rawDesc
)

func file_protocol_common_context_proto_rawDescGZIP() []byte {
	file_protocol_common_context_proto_rawDescOnce.Do(func() {
		file_protocol_common_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_common_context_proto_rawDescData)
	})
	return file_protocol_common_context_proto_rawDescData
}

var file_protocol_common_context_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocol_common_context_proto_goTypes = []interface{}{
	(*PositionInfo)(nil), // 0: common.PositionInfo
	(*AdContext)(nil),    // 1: common.AdContext
	(*Context)(nil),      // 2: common.Context
	nil,                  // 3: common.AdContext.MapContextEntry
	(*CPCContext)(nil),   // 4: common.CPCContext
	(*CPDContext)(nil),   // 5: common.CPDContext
}
var file_protocol_common_context_proto_depIdxs = []int32{
	0, // 0: common.AdContext.multiPositionInfo:type_name -> common.PositionInfo
	3, // 1: common.AdContext.mapContext:type_name -> common.AdContext.MapContextEntry
	1, // 2: common.Context.ad:type_name -> common.AdContext
	4, // 3: common.Context.cpc_context:type_name -> common.CPCContext
	5, // 4: common.Context.cpd_context:type_name -> common.CPDContext
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protocol_common_context_proto_init() }
func file_protocol_common_context_proto_init() {
	if File_protocol_common_context_proto != nil {
		return
	}
	file_protocol_common_cpc_context_proto_init()
	file_protocol_common_cpd_context_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protocol_common_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionInfo); i {
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
		file_protocol_common_context_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdContext); i {
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
		file_protocol_common_context_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
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
			RawDescriptor: file_protocol_common_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_common_context_proto_goTypes,
		DependencyIndexes: file_protocol_common_context_proto_depIdxs,
		MessageInfos:      file_protocol_common_context_proto_msgTypes,
	}.Build()
	File_protocol_common_context_proto = out.File
	file_protocol_common_context_proto_rawDesc = nil
	file_protocol_common_context_proto_goTypes = nil
	file_protocol_common_context_proto_depIdxs = nil
}