// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/broadcast/v1/push.proto

package bilibili_broadcast_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//
type LinkType int32

const (
	LinkType_LINK_TYPE_UNKNOWN LinkType = 0
	LinkType_LINK_TYPE_BANGUMI LinkType = 1
	LinkType_LINK_TYPE_VIDEO   LinkType = 2
	LinkType_LINK_TYPE_LIVE    LinkType = 3
)

var LinkType_name = map[int32]string{
	0: "LINK_TYPE_UNKNOWN",
	1: "LINK_TYPE_BANGUMI",
	2: "LINK_TYPE_VIDEO",
	3: "LINK_TYPE_LIVE",
}

var LinkType_value = map[string]int32{
	"LINK_TYPE_UNKNOWN": 0,
	"LINK_TYPE_BANGUMI": 1,
	"LINK_TYPE_VIDEO":   2,
	"LINK_TYPE_LIVE":    3,
}

func (x LinkType) String() string {
	return proto.EnumName(LinkType_name, int32(x))
}

func (LinkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5397841bc98f44da, []int{0}
}

// 业务类型
type PushMessageResp_Biz int32

const (
	// 未知
	PushMessageResp_BIZ_UNKNOWN PushMessageResp_Biz = 0
	// 视频
	PushMessageResp_BIZ_VIDEO PushMessageResp_Biz = 1
	// 直播
	PushMessageResp_BIZ_LIVE PushMessageResp_Biz = 2
	// 活动
	PushMessageResp_BIZ_ACTIVITY PushMessageResp_Biz = 3
)

var PushMessageResp_Biz_name = map[int32]string{
	0: "BIZ_UNKNOWN",
	1: "BIZ_VIDEO",
	2: "BIZ_LIVE",
	3: "BIZ_ACTIVITY",
}

var PushMessageResp_Biz_value = map[string]int32{
	"BIZ_UNKNOWN":  0,
	"BIZ_VIDEO":    1,
	"BIZ_LIVE":     2,
	"BIZ_ACTIVITY": 3,
}

func (x PushMessageResp_Biz) String() string {
	return proto.EnumName(PushMessageResp_Biz_name, int32(x))
}

func (PushMessageResp_Biz) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5397841bc98f44da, []int{2, 0}
}

// 消息类型
type PushMessageResp_Type int32

const (
	// 未知
	PushMessageResp_TYPE_UNKNOWN PushMessageResp_Type = 0
	// 默认
	PushMessageResp_TYPE_DEFAULT PushMessageResp_Type = 1
	// 热门
	PushMessageResp_TYPE_HOT PushMessageResp_Type = 2
	// 实时
	PushMessageResp_TYPE_REALTIME PushMessageResp_Type = 3
	// 推荐
	PushMessageResp_TYPE_RECOMMEND PushMessageResp_Type = 4
)

var PushMessageResp_Type_name = map[int32]string{
	0: "TYPE_UNKNOWN",
	1: "TYPE_DEFAULT",
	2: "TYPE_HOT",
	3: "TYPE_REALTIME",
	4: "TYPE_RECOMMEND",
}

var PushMessageResp_Type_value = map[string]int32{
	"TYPE_UNKNOWN":   0,
	"TYPE_DEFAULT":   1,
	"TYPE_HOT":       2,
	"TYPE_REALTIME":  3,
	"TYPE_RECOMMEND": 4,
}

func (x PushMessageResp_Type) String() string {
	return proto.EnumName(PushMessageResp_Type_name, int32(x))
}

func (PushMessageResp_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5397841bc98f44da, []int{2, 1}
}

// 展示未知
type PushMessageResp_Position int32

const (
	// 未知
	PushMessageResp_POS_UNKNOWN PushMessageResp_Position = 0
	// 顶部
	PushMessageResp_POS_TOP PushMessageResp_Position = 1
)

var PushMessageResp_Position_name = map[int32]string{
	0: "POS_UNKNOWN",
	1: "POS_TOP",
}

var PushMessageResp_Position_value = map[string]int32{
	"POS_UNKNOWN": 0,
	"POS_TOP":     1,
}

func (x PushMessageResp_Position) String() string {
	return proto.EnumName(PushMessageResp_Position_name, int32(x))
}

func (PushMessageResp_Position) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5397841bc98f44da, []int{2, 2}
}

//
type PageBlackList struct {
	//
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageBlackList) Reset()         { *m = PageBlackList{} }
func (m *PageBlackList) String() string { return proto.CompactTextString(m) }
func (*PageBlackList) ProtoMessage()    {}
func (*PageBlackList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5397841bc98f44da, []int{0}
}

func (m *PageBlackList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageBlackList.Unmarshal(m, b)
}
func (m *PageBlackList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageBlackList.Marshal(b, m, deterministic)
}
func (m *PageBlackList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageBlackList.Merge(m, src)
}
func (m *PageBlackList) XXX_Size() int {
	return xxx_messageInfo_PageBlackList.Size(m)
}
func (m *PageBlackList) XXX_DiscardUnknown() {
	xxx_messageInfo_PageBlackList.DiscardUnknown(m)
}

var xxx_messageInfo_PageBlackList proto.InternalMessageInfo

func (m *PageBlackList) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
type PageView struct {
	//
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageView) Reset()         { *m = PageView{} }
func (m *PageView) String() string { return proto.CompactTextString(m) }
func (*PageView) ProtoMessage()    {}
func (*PageView) Descriptor() ([]byte, []int) {
	return fileDescriptor_5397841bc98f44da, []int{1}
}

func (m *PageView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageView.Unmarshal(m, b)
}
func (m *PageView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageView.Marshal(b, m, deterministic)
}
func (m *PageView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageView.Merge(m, src)
}
func (m *PageView) XXX_Size() int {
	return xxx_messageInfo_PageView.Size(m)
}
func (m *PageView) XXX_DiscardUnknown() {
	xxx_messageInfo_PageView.DiscardUnknown(m)
}

var xxx_messageInfo_PageView proto.InternalMessageInfo

func (m *PageView) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
type PushMessageResp struct {
	// Deprecated: 推送任务id，使用string
	OldTaskid int64 `protobuf:"varint,1,opt,name=old_taskid,json=oldTaskid,proto3" json:"old_taskid,omitempty"`
	// 业务
	// 1:是视频 2:是直播 3:是活动
	Biz PushMessageResp_Biz `protobuf:"varint,2,opt,name=biz,proto3,enum=bilibili.broadcast.v1.PushMessageResp_Biz" json:"biz,omitempty"`
	// 类型
	// 1:是默认 2:是热门 3:是实时 4:是推荐
	Type PushMessageResp_Type `protobuf:"varint,3,opt,name=type,proto3,enum=bilibili.broadcast.v1.PushMessageResp_Type" json:"type,omitempty"`
	// 主标题
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// 副标题
	Summary string `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	// 图片地址
	Img string `protobuf:"bytes,6,opt,name=img,proto3" json:"img,omitempty"`
	// 跳转地址
	Link string `protobuf:"bytes,7,opt,name=link,proto3" json:"link,omitempty"`
	// 展示位置，1是顶部
	Position PushMessageResp_Position `protobuf:"varint,8,opt,name=position,proto3,enum=bilibili.broadcast.v1.PushMessageResp_Position" json:"position,omitempty"`
	// 展示时长（单位：秒），默认3秒
	Duration int32 `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	// 失效时间
	Expire int64 `protobuf:"varint,10,opt,name=expire,proto3" json:"expire,omitempty"`
	// 推送任务id
	Taskid string `protobuf:"bytes,11,opt,name=taskid,proto3" json:"taskid,omitempty"`
	// 应用内推送黑名单
	// UGC:     ugc-video-detail
	// PGC:     pgc-video-detail
	// 一起看:   pgc-video-detail-theater
	// 直播:     live-room-detail
	// Story:    ugc-video-detail-vertical
	// 播单黑名单 playlist-video-detail
	PageBlackList []*PageBlackList `protobuf:"bytes,12,rep,name=page_blackList,json=pageBlackList,proto3" json:"page_blackList,omitempty"`
	// 预留pvid
	PageView []*PageView `protobuf:"bytes,13,rep,name=page_view,json=pageView,proto3" json:"page_view,omitempty"`
	// 跳转资源
	TargetResource *TargetResource `protobuf:"bytes,14,opt,name=target_resource,json=targetResource,proto3" json:"target_resource,omitempty"`
	//
	ImageFrame int32 `protobuf:"varint,15,opt,name=image_frame,json=imageFrame,proto3" json:"image_frame,omitempty"`
	//
	ImageMarker int32 `protobuf:"varint,16,opt,name=image_marker,json=imageMarker,proto3" json:"image_marker,omitempty"`
	//
	ImagePosition int32 `protobuf:"varint,17,opt,name=image_position,json=imagePosition,proto3" json:"image_position,omitempty"`
	//
	Job                  int64    `protobuf:"varint,18,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushMessageResp) Reset()         { *m = PushMessageResp{} }
func (m *PushMessageResp) String() string { return proto.CompactTextString(m) }
func (*PushMessageResp) ProtoMessage()    {}
func (*PushMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5397841bc98f44da, []int{2}
}

func (m *PushMessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushMessageResp.Unmarshal(m, b)
}
func (m *PushMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushMessageResp.Marshal(b, m, deterministic)
}
func (m *PushMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMessageResp.Merge(m, src)
}
func (m *PushMessageResp) XXX_Size() int {
	return xxx_messageInfo_PushMessageResp.Size(m)
}
func (m *PushMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_PushMessageResp proto.InternalMessageInfo

func (m *PushMessageResp) GetOldTaskid() int64 {
	if m != nil {
		return m.OldTaskid
	}
	return 0
}

func (m *PushMessageResp) GetBiz() PushMessageResp_Biz {
	if m != nil {
		return m.Biz
	}
	return PushMessageResp_BIZ_UNKNOWN
}

func (m *PushMessageResp) GetType() PushMessageResp_Type {
	if m != nil {
		return m.Type
	}
	return PushMessageResp_TYPE_UNKNOWN
}

func (m *PushMessageResp) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PushMessageResp) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *PushMessageResp) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *PushMessageResp) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *PushMessageResp) GetPosition() PushMessageResp_Position {
	if m != nil {
		return m.Position
	}
	return PushMessageResp_POS_UNKNOWN
}

func (m *PushMessageResp) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *PushMessageResp) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

func (m *PushMessageResp) GetTaskid() string {
	if m != nil {
		return m.Taskid
	}
	return ""
}

func (m *PushMessageResp) GetPageBlackList() []*PageBlackList {
	if m != nil {
		return m.PageBlackList
	}
	return nil
}

func (m *PushMessageResp) GetPageView() []*PageView {
	if m != nil {
		return m.PageView
	}
	return nil
}

func (m *PushMessageResp) GetTargetResource() *TargetResource {
	if m != nil {
		return m.TargetResource
	}
	return nil
}

func (m *PushMessageResp) GetImageFrame() int32 {
	if m != nil {
		return m.ImageFrame
	}
	return 0
}

func (m *PushMessageResp) GetImageMarker() int32 {
	if m != nil {
		return m.ImageMarker
	}
	return 0
}

func (m *PushMessageResp) GetImagePosition() int32 {
	if m != nil {
		return m.ImagePosition
	}
	return 0
}

func (m *PushMessageResp) GetJob() int64 {
	if m != nil {
		return m.Job
	}
	return 0
}

//
type TargetResource struct {
	//直播:   roomid
	//UGC:   avid
	//PGC:   seasonid
	//Story: avid
	//举个例子
	//Type: LINK_TYPE_BANGUMI (番剧)
	//Resource: {"seasonid":"123"}
	//
	//Type: LINK_TYPE_VIDEO (视频)
	//Resource: {"avid":"123"}
	//
	//Type: LINK_TYPE_LIVE (直播)
	//Resource: {"roomid":"123"}
	//
	Type LinkType `protobuf:"varint,1,opt,name=Type,proto3,enum=bilibili.broadcast.v1.LinkType" json:"Type,omitempty"`
	//
	Resource             map[string]string `protobuf:"bytes,2,rep,name=Resource,proto3" json:"Resource,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TargetResource) Reset()         { *m = TargetResource{} }
func (m *TargetResource) String() string { return proto.CompactTextString(m) }
func (*TargetResource) ProtoMessage()    {}
func (*TargetResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_5397841bc98f44da, []int{3}
}

func (m *TargetResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetResource.Unmarshal(m, b)
}
func (m *TargetResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetResource.Marshal(b, m, deterministic)
}
func (m *TargetResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetResource.Merge(m, src)
}
func (m *TargetResource) XXX_Size() int {
	return xxx_messageInfo_TargetResource.Size(m)
}
func (m *TargetResource) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetResource.DiscardUnknown(m)
}

var xxx_messageInfo_TargetResource proto.InternalMessageInfo

func (m *TargetResource) GetType() LinkType {
	if m != nil {
		return m.Type
	}
	return LinkType_LINK_TYPE_UNKNOWN
}

func (m *TargetResource) GetResource() map[string]string {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterEnum("bilibili.broadcast.v1.LinkType", LinkType_name, LinkType_value)
	proto.RegisterEnum("bilibili.broadcast.v1.PushMessageResp_Biz", PushMessageResp_Biz_name, PushMessageResp_Biz_value)
	proto.RegisterEnum("bilibili.broadcast.v1.PushMessageResp_Type", PushMessageResp_Type_name, PushMessageResp_Type_value)
	proto.RegisterEnum("bilibili.broadcast.v1.PushMessageResp_Position", PushMessageResp_Position_name, PushMessageResp_Position_value)
	proto.RegisterType((*PageBlackList)(nil), "bilibili.broadcast.v1.PageBlackList")
	proto.RegisterType((*PageView)(nil), "bilibili.broadcast.v1.PageView")
	proto.RegisterType((*PushMessageResp)(nil), "bilibili.broadcast.v1.PushMessageResp")
	proto.RegisterType((*TargetResource)(nil), "bilibili.broadcast.v1.TargetResource")
	proto.RegisterMapType((map[string]string)(nil), "bilibili.broadcast.v1.TargetResource.ResourceEntry")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/broadcast/v1/push.proto", fileDescriptor_5397841bc98f44da)
}

var fileDescriptor_5397841bc98f44da = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xea, 0x46,
	0x10, 0xad, 0x6d, 0x12, 0xcc, 0x00, 0xc6, 0xd9, 0xf6, 0x5e, 0xad, 0xa8, 0xaa, 0x50, 0xd4, 0x5b,
	0xd1, 0x5b, 0xc9, 0x34, 0xe4, 0xa5, 0x6a, 0x23, 0x55, 0x90, 0x38, 0xad, 0xc5, 0xa7, 0x5c, 0x87,
	0x34, 0x7d, 0x41, 0x0b, 0x6c, 0x9c, 0x2d, 0x06, 0x5b, 0xb6, 0x21, 0x25, 0xbf, 0xb0, 0xef, 0xfd,
	0x43, 0xd5, 0xae, 0x6d, 0x12, 0xd2, 0xa6, 0xe2, 0x01, 0x69, 0xcf, 0xf1, 0x99, 0xc3, 0xec, 0xcc,
	0xec, 0xc0, 0x37, 0x6e, 0x18, 0xcc, 0x26, 0x24, 0x60, 0xcd, 0x29, 0xf3, 0x18, 0xff, 0x35, 0xa7,
	0xa1, 0x4f, 0xe6, 0x33, 0x12, 0xc5, 0xcd, 0xcd, 0x59, 0x33, 0x58, 0x47, 0x0f, 0x46, 0x10, 0xfa,
	0xb1, 0x8f, 0xde, 0x65, 0x0a, 0x63, 0xa7, 0x30, 0x36, 0x67, 0xd5, 0xcf, 0x5d, 0xdf, 0x77, 0x3d,
	0xda, 0x14, 0xa2, 0xe9, 0xfa, 0xbe, 0x49, 0x97, 0x41, 0xbc, 0x4d, 0x62, 0xea, 0xa7, 0x50, 0x1e,
	0x11, 0x97, 0x76, 0x3c, 0x32, 0x5b, 0xf4, 0x58, 0x14, 0x23, 0x0d, 0x64, 0x36, 0xc7, 0x52, 0x4d,
	0x6a, 0x14, 0x6c, 0x99, 0xcd, 0xeb, 0x55, 0x50, 0xb9, 0x60, 0xcc, 0xe8, 0xe3, 0xbf, 0xbe, 0xfd,
	0x95, 0x87, 0xca, 0x68, 0x1d, 0x3d, 0xf4, 0x69, 0x14, 0x11, 0x97, 0xda, 0x34, 0x0a, 0xd0, 0x17,
	0x00, 0xbe, 0x37, 0x9f, 0xc4, 0x24, 0x5a, 0xa4, 0x5a, 0xc5, 0x2e, 0xf8, 0xde, 0xdc, 0x11, 0x04,
	0xba, 0x00, 0x65, 0xca, 0x9e, 0xb0, 0x5c, 0x93, 0x1a, 0x5a, 0xeb, 0xa3, 0xf1, 0x9f, 0x19, 0x1b,
	0xaf, 0x3c, 0x8d, 0x0e, 0x7b, 0xb2, 0x79, 0x18, 0xfa, 0x09, 0x72, 0xf1, 0x36, 0xa0, 0x58, 0x11,
	0xe1, 0xdf, 0x1e, 0x18, 0xee, 0x6c, 0x03, 0x6a, 0x8b, 0x40, 0xf4, 0x19, 0x1c, 0xc5, 0x2c, 0xf6,
	0x28, 0xce, 0x89, 0x4b, 0x24, 0x00, 0x61, 0xc8, 0x47, 0xeb, 0xe5, 0x92, 0x84, 0x5b, 0x7c, 0x24,
	0xf8, 0x0c, 0x22, 0x1d, 0x14, 0xb6, 0x74, 0xf1, 0xb1, 0x60, 0xf9, 0x11, 0x21, 0xc8, 0x79, 0x6c,
	0xb5, 0xc0, 0x79, 0x41, 0x89, 0x33, 0xea, 0x82, 0x1a, 0xf8, 0x11, 0x8b, 0x99, 0xbf, 0xc2, 0xaa,
	0x48, 0xad, 0x79, 0x60, 0x6a, 0xa3, 0x34, 0xcc, 0xde, 0x19, 0xa0, 0x2a, 0xa8, 0xf3, 0x75, 0x48,
	0x84, 0x59, 0xa1, 0x26, 0x35, 0x8e, 0xec, 0x1d, 0x46, 0xef, 0xe1, 0x98, 0xfe, 0x19, 0xb0, 0x90,
	0x62, 0x10, 0x85, 0x4d, 0x11, 0xe7, 0xd3, 0x82, 0x17, 0x45, 0x5a, 0x29, 0x42, 0x5d, 0xd0, 0x02,
	0xe2, 0xd2, 0xc9, 0x34, 0x6b, 0x2f, 0x2e, 0xd5, 0x94, 0x46, 0xb1, 0xf5, 0xd5, 0x5b, 0xe9, 0xbd,
	0x1c, 0x05, 0xbb, 0x1c, 0xec, 0x4d, 0xc6, 0x05, 0x14, 0x84, 0xd9, 0x86, 0xd1, 0x47, 0x5c, 0x16,
	0x3e, 0xa7, 0xff, 0xe3, 0xc3, 0x27, 0xc6, 0x56, 0x83, 0x6c, 0x76, 0x06, 0x50, 0x89, 0x49, 0xe8,
	0xd2, 0x78, 0x12, 0xd2, 0xc8, 0x5f, 0x87, 0x33, 0x8a, 0xb5, 0x9a, 0xd4, 0x28, 0xb6, 0x3e, 0xbc,
	0xe1, 0xe1, 0x08, 0xb5, 0x9d, 0x8a, 0x6d, 0x2d, 0xde, 0xc3, 0xe8, 0x14, 0x8a, 0x6c, 0xc9, 0xd3,
	0xb9, 0x0f, 0xc9, 0x92, 0xe2, 0x8a, 0xa8, 0x14, 0x08, 0xea, 0x9a, 0x33, 0xe8, 0x4b, 0x28, 0x25,
	0x82, 0x25, 0x09, 0x17, 0x34, 0xc4, 0xba, 0x50, 0x24, 0x41, 0x7d, 0x41, 0xa1, 0x0f, 0xa0, 0x25,
	0x92, 0x5d, 0xf7, 0x4e, 0x84, 0xa8, 0x2c, 0xd8, 0xac, 0x37, 0x7c, 0x08, 0xfe, 0xf0, 0xa7, 0x18,
	0x89, 0x92, 0xf3, 0x63, 0xdd, 0x04, 0xa5, 0xc3, 0x9e, 0x50, 0x05, 0x8a, 0x1d, 0xeb, 0xf7, 0xc9,
	0xcd, 0xa0, 0x3b, 0x18, 0xde, 0x0e, 0xf4, 0x4f, 0x50, 0x19, 0x0a, 0x9c, 0x18, 0x5b, 0x57, 0xe6,
	0x50, 0x97, 0x50, 0x09, 0x54, 0x0e, 0x7b, 0xd6, 0xd8, 0xd4, 0x65, 0xa4, 0x43, 0x89, 0xa3, 0xf6,
	0xa5, 0x63, 0x8d, 0x2d, 0xe7, 0x4e, 0x57, 0xea, 0x13, 0xc8, 0xf1, 0xd9, 0xe4, 0x5f, 0x9c, 0xbb,
	0x91, 0xf9, 0xc2, 0x28, 0x63, 0xae, 0xcc, 0xeb, 0xf6, 0x4d, 0xcf, 0x49, 0xbc, 0x04, 0xf3, 0xcb,
	0xd0, 0xd1, 0x65, 0x74, 0x02, 0x65, 0x81, 0x6c, 0xb3, 0xdd, 0x73, 0xac, 0xbe, 0xa9, 0x2b, 0x08,
	0x81, 0x96, 0x52, 0x97, 0xc3, 0x7e, 0xdf, 0x1c, 0x5c, 0xe9, 0xb9, 0x7a, 0x03, 0xd4, 0xdd, 0x2d,
	0x2a, 0x50, 0x1c, 0x0d, 0x7f, 0x7d, 0xf1, 0x1f, 0x45, 0xc8, 0x73, 0xc2, 0x19, 0x8e, 0x74, 0xa9,
	0xfe, 0xb7, 0x04, 0xda, 0x7e, 0xc5, 0xd1, 0x79, 0x92, 0x9d, 0x78, 0xc3, 0xda, 0x9b, 0xad, 0xee,
	0xb1, 0xd5, 0x22, 0x79, 0x60, 0xe2, 0x2a, 0x43, 0x50, 0x33, 0x03, 0x2c, 0x8b, 0x19, 0x39, 0x3f,
	0xa8, 0xbf, 0x46, 0x76, 0x30, 0x57, 0x71, 0xb8, 0xb5, 0x77, 0x26, 0xd5, 0x1f, 0xa1, 0xbc, 0xf7,
	0x89, 0x77, 0x63, 0x41, 0xb7, 0xe9, 0x16, 0xe2, 0x47, 0xfe, 0xa8, 0x37, 0xc4, 0x5b, 0x53, 0xb1,
	0x55, 0x0a, 0x76, 0x02, 0x7e, 0x90, 0xbf, 0x97, 0x3e, 0x12, 0x50, 0xb3, 0xfc, 0xd0, 0x3b, 0x38,
	0xe9, 0x59, 0x83, 0xee, 0xe4, 0x55, 0xa5, 0xf7, 0xe8, 0x4e, 0x7b, 0xf0, 0xf3, 0x4d, 0xdf, 0xd2,
	0x25, 0xf4, 0x29, 0x54, 0x9e, 0xe9, 0xa4, 0x9f, 0x32, 0x2f, 0xf1, 0x33, 0x29, 0xba, 0xaa, 0xb4,
	0x7e, 0x83, 0x1c, 0x7f, 0xd4, 0x68, 0x04, 0xa5, 0x5b, 0x12, 0xcf, 0xb2, 0xd7, 0x8d, 0xde, 0x1b,
	0xc9, 0xda, 0x35, 0xb2, 0xb5, 0x6b, 0x98, 0x7c, 0xed, 0x56, 0xbf, 0x3e, 0x6c, 0x33, 0x7c, 0x27,
	0x4d, 0x8f, 0x45, 0xe4, 0xf9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xe5, 0x10, 0x81, 0x02,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PushClient is the client API for Push service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PushClient interface {
	WatchMessage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Push_WatchMessageClient, error)
}

type pushClient struct {
	cc *grpc.ClientConn
}

func NewPushClient(cc *grpc.ClientConn) PushClient {
	return &pushClient{cc}
}

func (c *pushClient) WatchMessage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Push_WatchMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Push_serviceDesc.Streams[0], "/bilibili.broadcast.v1.Push/WatchMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushWatchMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Push_WatchMessageClient interface {
	Recv() (*PushMessageResp, error)
	grpc.ClientStream
}

type pushWatchMessageClient struct {
	grpc.ClientStream
}

func (x *pushWatchMessageClient) Recv() (*PushMessageResp, error) {
	m := new(PushMessageResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PushServer is the server API for Push service.
type PushServer interface {
	WatchMessage(*emptypb.Empty, Push_WatchMessageServer) error
}

// UnimplementedPushServer can be embedded to have forward compatible implementations.
type UnimplementedPushServer struct {
}

func (*UnimplementedPushServer) WatchMessage(req *emptypb.Empty, srv Push_WatchMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchMessage not implemented")
}

func RegisterPushServer(s *grpc.Server, srv PushServer) {
	s.RegisterService(&_Push_serviceDesc, srv)
}

func _Push_WatchMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushServer).WatchMessage(m, &pushWatchMessageServer{stream})
}

type Push_WatchMessageServer interface {
	Send(*PushMessageResp) error
	grpc.ServerStream
}

type pushWatchMessageServer struct {
	grpc.ServerStream
}

func (x *pushWatchMessageServer) Send(m *PushMessageResp) error {
	return x.ServerStream.SendMsg(m)
}

var _Push_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.v1.Push",
	HandlerType: (*PushServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchMessage",
			Handler:       _Push_WatchMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc_api/bilibili/broadcast/v1/push.proto",
}