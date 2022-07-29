// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// option optimize_for = LITE_RUNTIME;

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.2
// source: protocol/common/recommend.proto

package rec5

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

type ResponseRecommendItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        []int64   `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`                          //creativeid
	ChildId   []int64   `protobuf:"varint,2,rep,packed,name=child_id,json=childId,proto3" json:"child_id,omitempty"` //派生广告id,动态创意id
	Score     []float32 `protobuf:"fixed32,3,rep,packed,name=score,proto3" json:"score,omitempty"`                   // ctr
	Cvr       []float32 `protobuf:"fixed32,4,rep,packed,name=cvr,proto3" json:"cvr,omitempty"`
	Rvr       []float32 `protobuf:"fixed32,5,rep,packed,name=rvr,proto3" json:"rvr,omitempty"`
	Cvr2      []float32 `protobuf:"fixed32,6,rep,packed,name=cvr2,proto3" json:"cvr2,omitempty"`
	RankScore []float32 `protobuf:"fixed32,7,rep,packed,name=rank_score,json=rankScore,proto3" json:"rank_score,omitempty"` // 粗排打分, 不启用时size为0
	StyleId   []int32   `protobuf:"varint,8,rep,packed,name=style_id,json=styleId,proto3" json:"style_id,omitempty"`        // 样式ID
	//原始打分列表，所有创意ID对应的ctr,cvr,rvr,cvr2 打平放到一个列表中（获取不到数据，这边会填充默认值）（按照顺序排列：ctr,cvr,rvr,cvr2）
	//逻辑计算规则：wild_score_info.size() / id.size() = 4 (代表目前存储了4个类型的数据)
	WildScoreInfo []float32 `protobuf:"fixed32,9,rep,packed,name=wild_score_info,json=wildScoreInfo,proto3" json:"wild_score_info,omitempty"`
}

func (x *ResponseRecommendItem) Reset() {
	*x = ResponseRecommendItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_common_recommend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRecommendItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRecommendItem) ProtoMessage() {}

func (x *ResponseRecommendItem) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_common_recommend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRecommendItem.ProtoReflect.Descriptor instead.
func (*ResponseRecommendItem) Descriptor() ([]byte, []int) {
	return file_protocol_common_recommend_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseRecommendItem) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ResponseRecommendItem) GetChildId() []int64 {
	if x != nil {
		return x.ChildId
	}
	return nil
}

func (x *ResponseRecommendItem) GetScore() []float32 {
	if x != nil {
		return x.Score
	}
	return nil
}

func (x *ResponseRecommendItem) GetCvr() []float32 {
	if x != nil {
		return x.Cvr
	}
	return nil
}

func (x *ResponseRecommendItem) GetRvr() []float32 {
	if x != nil {
		return x.Rvr
	}
	return nil
}

func (x *ResponseRecommendItem) GetCvr2() []float32 {
	if x != nil {
		return x.Cvr2
	}
	return nil
}

func (x *ResponseRecommendItem) GetRankScore() []float32 {
	if x != nil {
		return x.RankScore
	}
	return nil
}

func (x *ResponseRecommendItem) GetStyleId() []int32 {
	if x != nil {
		return x.StyleId
	}
	return nil
}

func (x *ResponseRecommendItem) GetWildScoreInfo() []float32 {
	if x != nil {
		return x.WildScoreInfo
	}
	return nil
}

type RequestRecommendItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             []int64   `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`                          //creativeid
	ChildId        []int64   `protobuf:"varint,2,rep,packed,name=child_id,json=childId,proto3" json:"child_id,omitempty"` //派生广告id,动态创意id
	Appid          []int64   `protobuf:"varint,3,rep,packed,name=appid,proto3" json:"appid,omitempty"`
	Stage          []int32   `protobuf:"varint,4,rep,packed,name=stage,proto3" json:"stage,omitempty"`                                         //ocpc阶段, 非ocpc广告=0, ocpc第一阶段=1, ocpc第二阶段=2, ocpc第三阶段=3
	CostType       []int32   `protobuf:"varint,5,rep,packed,name=costType,proto3" json:"costType,omitempty"`                                   // 区分广告的计费类型，如是否cpt
	SecondCostType []int32   `protobuf:"varint,6,rep,packed,name=secondCostType,proto3" json:"secondCostType,omitempty"`                       //ocpc阶段, 第二目标转化目标
	ScoreTypes     []int32   `protobuf:"varint,7,rep,packed,name=scoreTypes,proto3" json:"scoreTypes,omitempty"`                               // 1-ctr; 2-cvr; 4-rvr; 8-cvr2;
	AdType         []AD_TYPE `protobuf:"varint,8,rep,packed,name=adType,proto3,enum=rec5.AD_TYPE" json:"adType,omitempty"`                     // enum AdType {普通广告、通投广告、动态创意};
	Bid            []float32 `protobuf:"fixed32,9,rep,packed,name=bid,proto3" json:"bid,omitempty"`                                            // cpc广告主出价， ocpc风控出价
	TitleType      []int32   `protobuf:"varint,10,rep,packed,name=titleType,proto3" json:"titleType,omitempty"`                                //dpa广告标题类型：0-模板标题；1商品标题；2-动态标题
	SecondSmartBid []float32 `protobuf:"fixed32,11,rep,packed,name=secondSmartBid,proto3" json:"secondSmartBid,omitempty"`                     // 双出价广告-第二目标风控出价
	MaterialId     []int64   `protobuf:"varint,12,rep,packed,name=material_id,json=materialId,proto3" json:"material_id,omitempty"`            // 物料ID列表，所有创意ID对应的物料ID打平后放到一个列表
	MaterialIdPos  []int64   `protobuf:"varint,13,rep,packed,name=material_id_pos,json=materialIdPos,proto3" json:"material_id_pos,omitempty"` // 每个创意ID对应的物料ID列表在所有物料ID列表中的结束偏移，比如第一个创意ID对应四个物料ID，则material_id_pos中第一个元素为4，创意对应的物料ID列表为[上一个创意结束偏移, 当前创意结束偏移)
	StyleId        []int32   `protobuf:"varint,14,rep,packed,name=style_id,json=styleId,proto3" json:"style_id,omitempty"`                     // 样式ID列表，所有创意ID对应的样式ID打平后放到一个列表
	StyleIdPos     []int32   `protobuf:"varint,15,rep,packed,name=style_id_pos,json=styleIdPos,proto3" json:"style_id_pos,omitempty"`          // 每个创意ID对应的样式ID列表在所有样式ID列表中的结束偏移，比如第一个创意ID对应四个样式ID，则style_id_pos中第一个元素为4，创意对应的样式ID列表为[上一个创意结束偏移, 当前创意结束偏移)
}

func (x *RequestRecommendItem) Reset() {
	*x = RequestRecommendItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_common_recommend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRecommendItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRecommendItem) ProtoMessage() {}

func (x *RequestRecommendItem) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_common_recommend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRecommendItem.ProtoReflect.Descriptor instead.
func (*RequestRecommendItem) Descriptor() ([]byte, []int) {
	return file_protocol_common_recommend_proto_rawDescGZIP(), []int{1}
}

func (x *RequestRecommendItem) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *RequestRecommendItem) GetChildId() []int64 {
	if x != nil {
		return x.ChildId
	}
	return nil
}

func (x *RequestRecommendItem) GetAppid() []int64 {
	if x != nil {
		return x.Appid
	}
	return nil
}

func (x *RequestRecommendItem) GetStage() []int32 {
	if x != nil {
		return x.Stage
	}
	return nil
}

func (x *RequestRecommendItem) GetCostType() []int32 {
	if x != nil {
		return x.CostType
	}
	return nil
}

func (x *RequestRecommendItem) GetSecondCostType() []int32 {
	if x != nil {
		return x.SecondCostType
	}
	return nil
}

func (x *RequestRecommendItem) GetScoreTypes() []int32 {
	if x != nil {
		return x.ScoreTypes
	}
	return nil
}

func (x *RequestRecommendItem) GetAdType() []AD_TYPE {
	if x != nil {
		return x.AdType
	}
	return nil
}

func (x *RequestRecommendItem) GetBid() []float32 {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *RequestRecommendItem) GetTitleType() []int32 {
	if x != nil {
		return x.TitleType
	}
	return nil
}

func (x *RequestRecommendItem) GetSecondSmartBid() []float32 {
	if x != nil {
		return x.SecondSmartBid
	}
	return nil
}

func (x *RequestRecommendItem) GetMaterialId() []int64 {
	if x != nil {
		return x.MaterialId
	}
	return nil
}

func (x *RequestRecommendItem) GetMaterialIdPos() []int64 {
	if x != nil {
		return x.MaterialIdPos
	}
	return nil
}

func (x *RequestRecommendItem) GetStyleId() []int32 {
	if x != nil {
		return x.StyleId
	}
	return nil
}

func (x *RequestRecommendItem) GetStyleIdPos() []int32 {
	if x != nil {
		return x.StyleIdPos
	}
	return nil
}

//推荐信息
type RecommendInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId              int64                            `protobuf:"varint,1,opt,name=positionId,proto3" json:"positionId,omitempty"` // 广告位ID(int64)
	RefreshTimes            int32                            `protobuf:"varint,2,opt,name=refreshTimes,proto3" json:"refreshTimes,omitempty"`
	SspMediaType            int32                            `protobuf:"varint,3,opt,name=ssp_media_type,json=sspMediaType,proto3" json:"ssp_media_type,omitempty"`                                                                      // add by luojianhui 20200715, 1表示自有流量，2表示联盟开发者平台的媒体流量，3表示联盟ssp流量
	ReqItems                *RequestRecommendItem            `protobuf:"bytes,4,opt,name=reqItems,proto3" json:"reqItems,omitempty"`                                                                                                     // 请求广告集合
	ResItems                *ResponseRecommendItem           `protobuf:"bytes,5,opt,name=resItems,proto3" json:"resItems,omitempty"`                                                                                                     // 返回广告集合
	MapContext              map[string]string                `protobuf:"bytes,6,rep,name=mapContext,proto3" json:"mapContext,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`         // 推荐上下文信息
	AreaDldType             int32                            `protobuf:"varint,7,opt,name=areaDldType,proto3" json:"areaDldType,omitempty"`                                                                                              // 1表示 自动下载；2表示手动下载
	ExpName                 string                           `protobuf:"bytes,8,opt,name=expName,proto3" json:"expName,omitempty"`                                                                                                       // 实验信息
	MultiPositionId         []int64                          `protobuf:"varint,9,rep,packed,name=multiPositionId,proto3" json:"multiPositionId,omitempty"`                                                                               // 位次信息
	MultiResItems           map[int64]*ResponseRecommendItem `protobuf:"bytes,10,rep,name=multiResItems,proto3" json:"multiResItems,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //多位次信息返回结果
	MultiPositionIdBidFloor []int64                          `protobuf:"varint,11,rep,packed,name=multiPositionIdBidFloor,proto3" json:"multiPositionIdBidFloor,omitempty"`                                                              // add by shaodi(20210419), 上游传过来的底价，跟广告位相关。 后续跟广告位相关的抽象出一个message
	DldBitCtl               int32                            `protobuf:"varint,12,opt,name=dldBitCtl,proto3" json:"dldBitCtl,omitempty"`                                                                                                 //第一个广告位的dldBitCtl(广告联盟手动自动下载配置能力，按位获取)   取值为-1表示为空   取值大于等于0为有效值
	ClickableBitCtl         int32                            `protobuf:"varint,13,opt,name=clickableBitCtl,proto3" json:"clickableBitCtl,omitempty"`                                                                                     // 联盟广告区域是否可点，按位获取
}

func (x *RecommendInfo) Reset() {
	*x = RecommendInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_common_recommend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendInfo) ProtoMessage() {}

func (x *RecommendInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_common_recommend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendInfo.ProtoReflect.Descriptor instead.
func (*RecommendInfo) Descriptor() ([]byte, []int) {
	return file_protocol_common_recommend_proto_rawDescGZIP(), []int{2}
}

func (x *RecommendInfo) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *RecommendInfo) GetRefreshTimes() int32 {
	if x != nil {
		return x.RefreshTimes
	}
	return 0
}

func (x *RecommendInfo) GetSspMediaType() int32 {
	if x != nil {
		return x.SspMediaType
	}
	return 0
}

func (x *RecommendInfo) GetReqItems() *RequestRecommendItem {
	if x != nil {
		return x.ReqItems
	}
	return nil
}

func (x *RecommendInfo) GetResItems() *ResponseRecommendItem {
	if x != nil {
		return x.ResItems
	}
	return nil
}

func (x *RecommendInfo) GetMapContext() map[string]string {
	if x != nil {
		return x.MapContext
	}
	return nil
}

func (x *RecommendInfo) GetAreaDldType() int32 {
	if x != nil {
		return x.AreaDldType
	}
	return 0
}

func (x *RecommendInfo) GetExpName() string {
	if x != nil {
		return x.ExpName
	}
	return ""
}

func (x *RecommendInfo) GetMultiPositionId() []int64 {
	if x != nil {
		return x.MultiPositionId
	}
	return nil
}

func (x *RecommendInfo) GetMultiResItems() map[int64]*ResponseRecommendItem {
	if x != nil {
		return x.MultiResItems
	}
	return nil
}

func (x *RecommendInfo) GetMultiPositionIdBidFloor() []int64 {
	if x != nil {
		return x.MultiPositionIdBidFloor
	}
	return nil
}

func (x *RecommendInfo) GetDldBitCtl() int32 {
	if x != nil {
		return x.DldBitCtl
	}
	return 0
}

func (x *RecommendInfo) GetClickableBitCtl() int32 {
	if x != nil {
		return x.ClickableBitCtl
	}
	return 0
}

var File_protocol_common_recommend_proto protoreflect.FileDescriptor

var file_protocol_common_recommend_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x72, 0x65, 0x63, 0x35, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x62, 0x74, 0x2f, 0x72, 0x74,
	0x72, 0x73, 0x5f, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x75, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01, 0x0a, 0x15, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x02, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x72, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x02, 0x52, 0x03, 0x63, 0x76, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x76, 0x72, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x02, 0x52, 0x03, 0x72, 0x76, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x76, 0x72,
	0x32, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52, 0x04, 0x63, 0x76, 0x72, 0x32, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x02, 0x52, 0x09, 0x72, 0x61, 0x6e, 0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07,
	0x73, 0x74, 0x79, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x69, 0x6c, 0x64, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x03, 0x28, 0x02,
	0x52, 0x0d, 0x77, 0x69, 0x6c, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0xd6, 0x03, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x63, 0x35, 0x2e, 0x41, 0x44, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x52, 0x06, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x02, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x42, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x02, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x42,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x50, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x74, 0x79, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x73,
	0x74, 0x79, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74,
	0x79, 0x6c, 0x65, 0x49, 0x64, 0x50, 0x6f, 0x73, 0x22, 0x83, 0x06, 0x0a, 0x0d, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x24,
	0x0a, 0x0e, 0x73, 0x73, 0x70, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x73, 0x70, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x63, 0x35, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x08, 0x72, 0x65, 0x71, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x72, 0x65, 0x63, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x43, 0x0a, 0x0a, 0x6d, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x65, 0x63, 0x35,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d,
	0x61, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x6d, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x72,
	0x65, 0x61, 0x44, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x61, 0x72, 0x65, 0x61, 0x44, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x78, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x78, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x4c, 0x0a, 0x0d, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x65, 0x63, 0x35, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x52, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x38,
	0x0a, 0x17, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x42, 0x69, 0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x17, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x42, 0x69, 0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x6c, 0x64, 0x42,
	0x69, 0x74, 0x43, 0x74, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x6c, 0x64,
	0x42, 0x69, 0x74, 0x43, 0x74, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x61,
	0x62, 0x6c, 0x65, 0x42, 0x69, 0x74, 0x43, 0x74, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x69, 0x74, 0x43, 0x74, 0x6c,
	0x1a, 0x3d, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x5d, 0x0a, 0x12, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x63, 0x35, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x36,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x76, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x42, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x72, 0x65, 0x63, 0x35,
	0xa2, 0x02, 0x03, 0x48, 0x4c, 0x57, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_common_recommend_proto_rawDescOnce sync.Once
	file_protocol_common_recommend_proto_rawDescData = file_protocol_common_recommend_proto_rawDesc
)

func file_protocol_common_recommend_proto_rawDescGZIP() []byte {
	file_protocol_common_recommend_proto_rawDescOnce.Do(func() {
		file_protocol_common_recommend_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_common_recommend_proto_rawDescData)
	})
	return file_protocol_common_recommend_proto_rawDescData
}

var file_protocol_common_recommend_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protocol_common_recommend_proto_goTypes = []interface{}{
	(*ResponseRecommendItem)(nil), // 0: rec5.ResponseRecommendItem
	(*RequestRecommendItem)(nil),  // 1: rec5.RequestRecommendItem
	(*RecommendInfo)(nil),         // 2: rec5.RecommendInfo
	nil,                           // 3: rec5.RecommendInfo.MapContextEntry
	nil,                           // 4: rec5.RecommendInfo.MultiResItemsEntry
	(AD_TYPE)(0),                  // 5: rec5.AD_TYPE
}
var file_protocol_common_recommend_proto_depIdxs = []int32{
	5, // 0: rec5.RequestRecommendItem.adType:type_name -> rec5.AD_TYPE
	1, // 1: rec5.RecommendInfo.reqItems:type_name -> rec5.RequestRecommendItem
	0, // 2: rec5.RecommendInfo.resItems:type_name -> rec5.ResponseRecommendItem
	3, // 3: rec5.RecommendInfo.mapContext:type_name -> rec5.RecommendInfo.MapContextEntry
	4, // 4: rec5.RecommendInfo.multiResItems:type_name -> rec5.RecommendInfo.MultiResItemsEntry
	0, // 5: rec5.RecommendInfo.MultiResItemsEntry.value:type_name -> rec5.ResponseRecommendItem
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protocol_common_recommend_proto_init() }
func file_protocol_common_recommend_proto_init() {
	if File_protocol_common_recommend_proto != nil {
		return
	}
	file_protocol_common_enum_proto_init()
	file_protocol_common_user_proto_init()
	file_protocol_abt_rtrs_abtest_proto_init()
	file_protocol_common_uac_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protocol_common_recommend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRecommendItem); i {
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
		file_protocol_common_recommend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRecommendItem); i {
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
		file_protocol_common_recommend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendInfo); i {
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
			RawDescriptor: file_protocol_common_recommend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_common_recommend_proto_goTypes,
		DependencyIndexes: file_protocol_common_recommend_proto_depIdxs,
		MessageInfos:      file_protocol_common_recommend_proto_msgTypes,
	}.Build()
	File_protocol_common_recommend_proto = out.File
	file_protocol_common_recommend_proto_rawDesc = nil
	file_protocol_common_recommend_proto_goTypes = nil
	file_protocol_common_recommend_proto_depIdxs = nil
}
