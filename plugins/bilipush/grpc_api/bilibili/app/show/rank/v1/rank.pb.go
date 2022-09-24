// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/app/show/rank/v1/rank.proto

package bilibili_app_show_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 排行榜列表项
type Item struct {
	// 标题
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// 封面url
	Cover string `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`
	// 参数(稿件avid)
	Param string `protobuf:"bytes,3,opt,name=param,proto3" json:"param,omitempty"`
	// 跳转uri
	Uri string `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
	// 重定向url
	RedirectUrl string `protobuf:"bytes,5,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	// 跳转类型
	// av:视频稿件
	Goto string `protobuf:"bytes,6,opt,name=goto,proto3" json:"goto,omitempty"`
	// 播放数
	Play int32 `protobuf:"varint,7,opt,name=play,proto3" json:"play,omitempty"`
	// 弹幕数
	Danmaku int32 `protobuf:"varint,8,opt,name=danmaku,proto3" json:"danmaku,omitempty"`
	// UP主mid
	Mid int64 `protobuf:"varint,9,opt,name=mid,proto3" json:"mid,omitempty"`
	// UP主昵称
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// UP主头像url
	Face string `protobuf:"bytes,11,opt,name=face,proto3" json:"face,omitempty"`
	// 评论数
	Reply int32 `protobuf:"varint,12,opt,name=reply,proto3" json:"reply,omitempty"`
	// 收藏数
	Favourite int32 `protobuf:"varint,13,opt,name=favourite,proto3" json:"favourite,omitempty"`
	// 发布时间
	PubDate int64 `protobuf:"varint,14,opt,name=pub_date,json=pubDate,proto3" json:"pub_date,omitempty"`
	// 分区tid
	Rid int32 `protobuf:"varint,15,opt,name=rid,proto3" json:"rid,omitempty"`
	// 子分区名
	Rname string `protobuf:"bytes,16,opt,name=rname,proto3" json:"rname,omitempty"`
	// 视频总时长
	Duration int64 `protobuf:"varint,17,opt,name=duration,proto3" json:"duration,omitempty"`
	// 点赞数
	Like int32 `protobuf:"varint,18,opt,name=like,proto3" json:"like,omitempty"`
	// 1P cid
	Cid int64 `protobuf:"varint,19,opt,name=cid,proto3" json:"cid,omitempty"`
	// 综合评分
	Pts int64 `protobuf:"varint,20,opt,name=pts,proto3" json:"pts,omitempty"`
	// 合作视频文案
	Cooperation string `protobuf:"bytes,21,opt,name=cooperation,proto3" json:"cooperation,omitempty"`
	// 属性位
	// 0:未关注 1:已关注
	Attribute int32 `protobuf:"varint,22,opt,name=attribute,proto3" json:"attribute,omitempty"`
	// UP主粉丝数
	Follower int64 `protobuf:"varint,23,opt,name=follower,proto3" json:"follower,omitempty"`
	// UP主认证信息
	OfficialVerify *OfficialVerify `protobuf:"bytes,24,opt,name=official_verify,json=officialVerify,proto3" json:"official_verify,omitempty"`
	// 同一UP收起子项列表
	Children []*Item `protobuf:"bytes,25,rep,name=children,proto3" json:"children,omitempty"`
	// 关系信息
	Relation             *Relation `protobuf:"bytes,26,opt,name=relation,proto3" json:"relation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_025e0746947f3e29, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Item) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *Item) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *Item) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Item) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

func (m *Item) GetGoto() string {
	if m != nil {
		return m.Goto
	}
	return ""
}

func (m *Item) GetPlay() int32 {
	if m != nil {
		return m.Play
	}
	return 0
}

func (m *Item) GetDanmaku() int32 {
	if m != nil {
		return m.Danmaku
	}
	return 0
}

func (m *Item) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

func (m *Item) GetReply() int32 {
	if m != nil {
		return m.Reply
	}
	return 0
}

func (m *Item) GetFavourite() int32 {
	if m != nil {
		return m.Favourite
	}
	return 0
}

func (m *Item) GetPubDate() int64 {
	if m != nil {
		return m.PubDate
	}
	return 0
}

func (m *Item) GetRid() int32 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *Item) GetRname() string {
	if m != nil {
		return m.Rname
	}
	return ""
}

func (m *Item) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Item) GetLike() int32 {
	if m != nil {
		return m.Like
	}
	return 0
}

func (m *Item) GetCid() int64 {
	if m != nil {
		return m.Cid
	}
	return 0
}

func (m *Item) GetPts() int64 {
	if m != nil {
		return m.Pts
	}
	return 0
}

func (m *Item) GetCooperation() string {
	if m != nil {
		return m.Cooperation
	}
	return ""
}

func (m *Item) GetAttribute() int32 {
	if m != nil {
		return m.Attribute
	}
	return 0
}

func (m *Item) GetFollower() int64 {
	if m != nil {
		return m.Follower
	}
	return 0
}

func (m *Item) GetOfficialVerify() *OfficialVerify {
	if m != nil {
		return m.OfficialVerify
	}
	return nil
}

func (m *Item) GetChildren() []*Item {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Item) GetRelation() *Relation {
	if m != nil {
		return m.Relation
	}
	return nil
}

// 认证信息
type OfficialVerify struct {
	// 认证类型
	// -1:无认证 0:个人认证 1:机构认证
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// 认证描述
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfficialVerify) Reset()         { *m = OfficialVerify{} }
func (m *OfficialVerify) String() string { return proto.CompactTextString(m) }
func (*OfficialVerify) ProtoMessage()    {}
func (*OfficialVerify) Descriptor() ([]byte, []int) {
	return fileDescriptor_025e0746947f3e29, []int{1}
}

func (m *OfficialVerify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfficialVerify.Unmarshal(m, b)
}
func (m *OfficialVerify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfficialVerify.Marshal(b, m, deterministic)
}
func (m *OfficialVerify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfficialVerify.Merge(m, src)
}
func (m *OfficialVerify) XXX_Size() int {
	return xxx_messageInfo_OfficialVerify.Size(m)
}
func (m *OfficialVerify) XXX_DiscardUnknown() {
	xxx_messageInfo_OfficialVerify.DiscardUnknown(m)
}

var xxx_messageInfo_OfficialVerify proto.InternalMessageInfo

func (m *OfficialVerify) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *OfficialVerify) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

// 全站排行榜-请求
type RankAllResultReq struct {
	// 必须为"all"
	Order string `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	// 页码
	// 默认1页
	Pn int32 `protobuf:"varint,2,opt,name=pn,proto3" json:"pn,omitempty"`
	// 每页项数
	// 默认100项，最大100
	Ps                   int32    `protobuf:"varint,3,opt,name=ps,proto3" json:"ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankAllResultReq) Reset()         { *m = RankAllResultReq{} }
func (m *RankAllResultReq) String() string { return proto.CompactTextString(m) }
func (*RankAllResultReq) ProtoMessage()    {}
func (*RankAllResultReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_025e0746947f3e29, []int{2}
}

func (m *RankAllResultReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankAllResultReq.Unmarshal(m, b)
}
func (m *RankAllResultReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankAllResultReq.Marshal(b, m, deterministic)
}
func (m *RankAllResultReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankAllResultReq.Merge(m, src)
}
func (m *RankAllResultReq) XXX_Size() int {
	return xxx_messageInfo_RankAllResultReq.Size(m)
}
func (m *RankAllResultReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankAllResultReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankAllResultReq proto.InternalMessageInfo

func (m *RankAllResultReq) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *RankAllResultReq) GetPn() int32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *RankAllResultReq) GetPs() int32 {
	if m != nil {
		return m.Ps
	}
	return 0
}

// 排行榜信息-响应
type RankListReply struct {
	// 排行榜列表
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankListReply) Reset()         { *m = RankListReply{} }
func (m *RankListReply) String() string { return proto.CompactTextString(m) }
func (*RankListReply) ProtoMessage()    {}
func (*RankListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_025e0746947f3e29, []int{3}
}

func (m *RankListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankListReply.Unmarshal(m, b)
}
func (m *RankListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankListReply.Marshal(b, m, deterministic)
}
func (m *RankListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankListReply.Merge(m, src)
}
func (m *RankListReply) XXX_Size() int {
	return xxx_messageInfo_RankListReply.Size(m)
}
func (m *RankListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RankListReply.DiscardUnknown(m)
}

var xxx_messageInfo_RankListReply proto.InternalMessageInfo

func (m *RankListReply) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

// 分区排行榜-请求
type RankRegionResultReq struct {
	// 一级分区tid(二级分区不可用)
	// 0:全站
	Rid int32 `protobuf:"varint,1,opt,name=rid,proto3" json:"rid,omitempty"`
	// 页码
	// 默认1页
	Pn int32 `protobuf:"varint,2,opt,name=pn,proto3" json:"pn,omitempty"`
	// 每页项数
	// 默认100项，最大100
	Ps                   int32    `protobuf:"varint,3,opt,name=ps,proto3" json:"ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankRegionResultReq) Reset()         { *m = RankRegionResultReq{} }
func (m *RankRegionResultReq) String() string { return proto.CompactTextString(m) }
func (*RankRegionResultReq) ProtoMessage()    {}
func (*RankRegionResultReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_025e0746947f3e29, []int{4}
}

func (m *RankRegionResultReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankRegionResultReq.Unmarshal(m, b)
}
func (m *RankRegionResultReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankRegionResultReq.Marshal(b, m, deterministic)
}
func (m *RankRegionResultReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankRegionResultReq.Merge(m, src)
}
func (m *RankRegionResultReq) XXX_Size() int {
	return xxx_messageInfo_RankRegionResultReq.Size(m)
}
func (m *RankRegionResultReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankRegionResultReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankRegionResultReq proto.InternalMessageInfo

func (m *RankRegionResultReq) GetRid() int32 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *RankRegionResultReq) GetPn() int32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *RankRegionResultReq) GetPs() int32 {
	if m != nil {
		return m.Ps
	}
	return 0
}

// 关系信息
type Relation struct {
	// 关系状态id
	// 1:未关注 2:已关注 3:被关注 4:互相关注
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// 是否关注
	IsFollow int32 `protobuf:"varint,2,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
	// 是否粉丝
	IsFollowed           int32    `protobuf:"varint,3,opt,name=is_followed,json=isFollowed,proto3" json:"is_followed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Relation) Reset()         { *m = Relation{} }
func (m *Relation) String() string { return proto.CompactTextString(m) }
func (*Relation) ProtoMessage()    {}
func (*Relation) Descriptor() ([]byte, []int) {
	return fileDescriptor_025e0746947f3e29, []int{5}
}

func (m *Relation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Relation.Unmarshal(m, b)
}
func (m *Relation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Relation.Marshal(b, m, deterministic)
}
func (m *Relation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relation.Merge(m, src)
}
func (m *Relation) XXX_Size() int {
	return xxx_messageInfo_Relation.Size(m)
}
func (m *Relation) XXX_DiscardUnknown() {
	xxx_messageInfo_Relation.DiscardUnknown(m)
}

var xxx_messageInfo_Relation proto.InternalMessageInfo

func (m *Relation) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Relation) GetIsFollow() int32 {
	if m != nil {
		return m.IsFollow
	}
	return 0
}

func (m *Relation) GetIsFollowed() int32 {
	if m != nil {
		return m.IsFollowed
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "bilibili.app.show.v1.Item")
	proto.RegisterType((*OfficialVerify)(nil), "bilibili.app.show.v1.OfficialVerify")
	proto.RegisterType((*RankAllResultReq)(nil), "bilibili.app.show.v1.RankAllResultReq")
	proto.RegisterType((*RankListReply)(nil), "bilibili.app.show.v1.RankListReply")
	proto.RegisterType((*RankRegionResultReq)(nil), "bilibili.app.show.v1.RankRegionResultReq")
	proto.RegisterType((*Relation)(nil), "bilibili.app.show.v1.Relation")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/app/show/rank/v1/rank.proto", fileDescriptor_025e0746947f3e29)
}

var fileDescriptor_025e0746947f3e29 = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0x13, 0x3b,
	0x10, 0xd5, 0x36, 0x49, 0x93, 0x4c, 0xda, 0xb4, 0xd7, 0xed, 0xed, 0x75, 0x73, 0xaf, 0x2e, 0x21,
	0x20, 0x14, 0x1e, 0x48, 0x68, 0x91, 0x10, 0xe2, 0xad, 0x12, 0xe2, 0x43, 0x02, 0x21, 0xad, 0x44,
	0x9f, 0x90, 0x82, 0xb3, 0xeb, 0xb4, 0x56, 0x9c, 0xb5, 0xf1, 0x7a, 0x53, 0xe5, 0xbf, 0xf1, 0x6f,
	0xf8, 0x23, 0x68, 0xc6, 0xbb, 0x09, 0xad, 0x5a, 0x95, 0x87, 0x28, 0x73, 0x8e, 0x67, 0x8e, 0x67,
	0xce, 0xce, 0x2e, 0x3c, 0xbb, 0x70, 0x36, 0x99, 0x08, 0xab, 0xc6, 0x53, 0xa5, 0x15, 0xfe, 0xc6,
	0xc2, 0xda, 0x71, 0x7e, 0x69, 0xae, 0xc6, 0x4e, 0x64, 0xf3, 0xf1, 0xf2, 0x84, 0xfe, 0x47, 0xd6,
	0x19, 0x6f, 0xd8, 0x61, 0x95, 0x35, 0x12, 0xd6, 0x8e, 0x30, 0x6b, 0xb4, 0x3c, 0x19, 0xfc, 0x6c,
	0x40, 0xfd, 0x83, 0x97, 0x0b, 0x76, 0x08, 0x0d, 0xaf, 0xbc, 0x96, 0x3c, 0xea, 0x47, 0xc3, 0x76,
	0x1c, 0x00, 0xb2, 0x89, 0x59, 0x4a, 0xc7, 0xb7, 0x02, 0x4b, 0x00, 0x59, 0x2b, 0x9c, 0x58, 0xf0,
	0x5a, 0x60, 0x09, 0xb0, 0x7d, 0xa8, 0x15, 0x4e, 0xf1, 0x3a, 0x71, 0x18, 0xb2, 0x87, 0xb0, 0xe3,
	0x64, 0xaa, 0x9c, 0x4c, 0xfc, 0xa4, 0x70, 0x9a, 0x37, 0xe8, 0xa8, 0x53, 0x71, 0x5f, 0x9c, 0x66,
	0x0c, 0xea, 0x17, 0xc6, 0x1b, 0xbe, 0x4d, 0x47, 0x14, 0x23, 0x67, 0xb5, 0x58, 0xf1, 0x66, 0x3f,
	0x1a, 0x36, 0x62, 0x8a, 0x19, 0x87, 0x66, 0x2a, 0xb2, 0x85, 0x98, 0x17, 0xbc, 0x45, 0x74, 0x05,
	0xf1, 0xda, 0x85, 0x4a, 0x79, 0xbb, 0x1f, 0x0d, 0x6b, 0x31, 0x86, 0x58, 0x9f, 0x89, 0x85, 0xe4,
	0x10, 0x34, 0x31, 0x46, 0x6e, 0x26, 0x12, 0xc9, 0x3b, 0x81, 0xc3, 0x18, 0xc7, 0x70, 0xd2, 0xea,
	0x15, 0xdf, 0x21, 0xc5, 0x00, 0xd8, 0x7f, 0xd0, 0x9e, 0x89, 0xa5, 0x29, 0x9c, 0xf2, 0x92, 0xef,
	0xd2, 0xc9, 0x86, 0x60, 0xc7, 0xd0, 0xb2, 0xc5, 0x74, 0x92, 0x0a, 0x2f, 0x79, 0x97, 0xae, 0x6c,
	0xda, 0x62, 0xfa, 0x46, 0x78, 0x89, 0x8d, 0x38, 0x95, 0xf2, 0x3d, 0x2a, 0xc1, 0x90, 0x2e, 0xa0,
	0x4e, 0xf6, 0x83, 0x4f, 0x04, 0x58, 0x0f, 0x5a, 0x69, 0xe1, 0x84, 0x57, 0x26, 0xe3, 0x7f, 0x91,
	0xc4, 0x1a, 0x63, 0x9b, 0x5a, 0xcd, 0x25, 0x67, 0x61, 0x74, 0x8c, 0x51, 0x37, 0x51, 0x29, 0x3f,
	0x08, 0x03, 0x26, 0x2a, 0x45, 0xc6, 0xfa, 0x9c, 0x1f, 0x06, 0xc6, 0xfa, 0x9c, 0xf5, 0xa1, 0x93,
	0x18, 0x63, 0x65, 0x29, 0xfb, 0x77, 0x30, 0xfa, 0x37, 0x0a, 0xc7, 0x12, 0xde, 0x3b, 0x35, 0x2d,
	0xbc, 0xe4, 0x47, 0x61, 0xac, 0x35, 0x81, 0x3d, 0xcd, 0x8c, 0xd6, 0xe6, 0x4a, 0x3a, 0xfe, 0x4f,
	0xe8, 0xa9, 0xc2, 0xec, 0x13, 0xec, 0x99, 0xd9, 0x4c, 0x25, 0x4a, 0xe8, 0xc9, 0x52, 0x3a, 0x35,
	0x5b, 0x71, 0xde, 0x8f, 0x86, 0x9d, 0xd3, 0xc7, 0xa3, 0xdb, 0x56, 0x6a, 0xf4, 0xb9, 0x4c, 0x3e,
	0xa7, 0xdc, 0xb8, 0x6b, 0xae, 0x61, 0xf6, 0x12, 0x5a, 0xc9, 0xa5, 0xd2, 0xa9, 0x93, 0x19, 0x3f,
	0xee, 0xd7, 0x86, 0x9d, 0xd3, 0xde, 0xed, 0x3a, 0xb8, 0x96, 0xf1, 0x3a, 0x97, 0xbd, 0x86, 0x96,
	0x93, 0x3a, 0xcc, 0xd7, 0xa3, 0xfb, 0xff, 0xbf, 0xbd, 0x2e, 0x2e, 0xb3, 0xe2, 0x75, 0xfe, 0xe0,
	0x15, 0x74, 0xaf, 0x77, 0x85, 0x46, 0xfb, 0x95, 0x0d, 0xdb, 0xde, 0x88, 0x29, 0x46, 0x2e, 0x95,
	0x79, 0x52, 0xee, 0x3a, 0xc5, 0x83, 0xf7, 0xb0, 0x1f, 0x8b, 0x6c, 0x7e, 0xa6, 0x75, 0x2c, 0xf3,
	0x42, 0xfb, 0x58, 0x7e, 0xc7, 0xc7, 0x6a, 0x5c, 0x2a, 0x5d, 0xf5, 0xaa, 0x10, 0x60, 0x5d, 0xd8,
	0xb2, 0x19, 0xd5, 0x36, 0xe2, 0x2d, 0x9b, 0x11, 0xce, 0xe9, 0x0d, 0x41, 0x9c, 0x0f, 0xce, 0x60,
	0x17, 0x95, 0x3e, 0xaa, 0xdc, 0xc7, 0xb4, 0x68, 0xcf, 0xa1, 0xa1, 0xbc, 0x5c, 0xe4, 0x3c, 0xba,
	0xd7, 0x85, 0x90, 0x38, 0x78, 0x07, 0x07, 0x28, 0x11, 0xcb, 0x0b, 0x1c, 0x6f, 0xdd, 0x4f, 0xb9,
	0x78, 0xd1, 0x66, 0xf1, 0xee, 0xeb, 0xe5, 0x1b, 0xb4, 0x2a, 0x97, 0xd8, 0x11, 0x6c, 0xe7, 0x5e,
	0xf8, 0x22, 0x2f, 0x05, 0x4a, 0xc4, 0xfe, 0x85, 0xb6, 0xca, 0x27, 0x61, 0x0b, 0x4a, 0xa9, 0x96,
	0xca, 0xdf, 0x12, 0x66, 0x0f, 0xa0, 0xb3, 0x3e, 0x94, 0x69, 0xa9, 0x0c, 0xd5, 0xb1, 0x4c, 0x4f,
	0x7f, 0x44, 0x50, 0xc7, 0x5e, 0xd9, 0x39, 0x34, 0x4b, 0x03, 0xd9, 0x93, 0x3b, 0x9e, 0xd7, 0x0d,
	0x7f, 0x7b, 0x8f, 0xee, 0xce, 0xdb, 0xb8, 0xf7, 0x15, 0x60, 0xe3, 0x05, 0x7b, 0x7a, 0x77, 0xc9,
	0x0d, 0xb7, 0xfe, 0x48, 0x7d, 0xba, 0x4d, 0xdf, 0xcc, 0x17, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x0d, 0xdf, 0xcc, 0x58, 0x64, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RankClient is the client API for Rank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RankClient interface {
	// 全站排行榜
	RankAll(ctx context.Context, in *RankAllResultReq, opts ...grpc.CallOption) (*RankListReply, error)
	// 分区排行榜
	RankRegion(ctx context.Context, in *RankRegionResultReq, opts ...grpc.CallOption) (*RankListReply, error)
}

type rankClient struct {
	cc *grpc.ClientConn
}

func NewRankClient(cc *grpc.ClientConn) RankClient {
	return &rankClient{cc}
}

func (c *rankClient) RankAll(ctx context.Context, in *RankAllResultReq, opts ...grpc.CallOption) (*RankListReply, error) {
	out := new(RankListReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.show.v1.Rank/RankAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankClient) RankRegion(ctx context.Context, in *RankRegionResultReq, opts ...grpc.CallOption) (*RankListReply, error) {
	out := new(RankListReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.show.v1.Rank/RankRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankServer is the server API for Rank service.
type RankServer interface {
	// 全站排行榜
	RankAll(context.Context, *RankAllResultReq) (*RankListReply, error)
	// 分区排行榜
	RankRegion(context.Context, *RankRegionResultReq) (*RankListReply, error)
}

// UnimplementedRankServer can be embedded to have forward compatible implementations.
type UnimplementedRankServer struct {
}

func (*UnimplementedRankServer) RankAll(ctx context.Context, req *RankAllResultReq) (*RankListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RankAll not implemented")
}
func (*UnimplementedRankServer) RankRegion(ctx context.Context, req *RankRegionResultReq) (*RankListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RankRegion not implemented")
}

func RegisterRankServer(s *grpc.Server, srv RankServer) {
	s.RegisterService(&_Rank_serviceDesc, srv)
}

func _Rank_RankAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankAllResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).RankAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.show.v1.Rank/RankAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).RankAll(ctx, req.(*RankAllResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rank_RankRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankRegionResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).RankRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.show.v1.Rank/RankRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).RankRegion(ctx, req.(*RankRegionResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rank_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.show.v1.Rank",
	HandlerType: (*RankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RankAll",
			Handler:    _Rank_RankAll_Handler,
		},
		{
			MethodName: "RankRegion",
			Handler:    _Rank_RankRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_api/bilibili/app/show/rank/v1/rank.proto",
}