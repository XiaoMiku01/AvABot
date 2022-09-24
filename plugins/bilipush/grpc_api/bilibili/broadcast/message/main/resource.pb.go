// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/broadcast/message/main/resource.proto

package bilibili_broadcast_message_main

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
type TopActivityReply struct {
	// 当前生效的资源
	Online *TopOnline `protobuf:"bytes,1,opt,name=online,proto3" json:"online,omitempty"`
	// 对online内容进行hash和上次结果一样则不重新加载
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopActivityReply) Reset()         { *m = TopActivityReply{} }
func (m *TopActivityReply) String() string { return proto.CompactTextString(m) }
func (*TopActivityReply) ProtoMessage()    {}
func (*TopActivityReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_498ddd592723f790, []int{0}
}

func (m *TopActivityReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopActivityReply.Unmarshal(m, b)
}
func (m *TopActivityReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopActivityReply.Marshal(b, m, deterministic)
}
func (m *TopActivityReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopActivityReply.Merge(m, src)
}
func (m *TopActivityReply) XXX_Size() int {
	return xxx_messageInfo_TopActivityReply.Size(m)
}
func (m *TopActivityReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TopActivityReply.DiscardUnknown(m)
}

var xxx_messageInfo_TopActivityReply proto.InternalMessageInfo

func (m *TopActivityReply) GetOnline() *TopOnline {
	if m != nil {
		return m.Online
	}
	return nil
}

func (m *TopActivityReply) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// 当前生效的资源
type TopOnline struct {
	// 活动类型
	// 1:七日活动 2:后台配置
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// 图标
	Icon string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	// 跳转链接
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// 资源状态标识(后台配置)
	UniqueId string `protobuf:"bytes,4,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	// 动画资源
	Animate *Animate `protobuf:"bytes,5,opt,name=animate,proto3" json:"animate,omitempty"`
	// 红点
	RedDot *RedDot `protobuf:"bytes,6,opt,name=red_dot,json=redDot,proto3" json:"red_dot,omitempty"`
	// 活动名称
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	// 轮询间隔 单位秒
	Interval             int64    `protobuf:"varint,8,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopOnline) Reset()         { *m = TopOnline{} }
func (m *TopOnline) String() string { return proto.CompactTextString(m) }
func (*TopOnline) ProtoMessage()    {}
func (*TopOnline) Descriptor() ([]byte, []int) {
	return fileDescriptor_498ddd592723f790, []int{1}
}

func (m *TopOnline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopOnline.Unmarshal(m, b)
}
func (m *TopOnline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopOnline.Marshal(b, m, deterministic)
}
func (m *TopOnline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopOnline.Merge(m, src)
}
func (m *TopOnline) XXX_Size() int {
	return xxx_messageInfo_TopOnline.Size(m)
}
func (m *TopOnline) XXX_DiscardUnknown() {
	xxx_messageInfo_TopOnline.DiscardUnknown(m)
}

var xxx_messageInfo_TopOnline proto.InternalMessageInfo

func (m *TopOnline) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *TopOnline) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *TopOnline) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *TopOnline) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *TopOnline) GetAnimate() *Animate {
	if m != nil {
		return m.Animate
	}
	return nil
}

func (m *TopOnline) GetRedDot() *RedDot {
	if m != nil {
		return m.RedDot
	}
	return nil
}

func (m *TopOnline) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TopOnline) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

// 动画资源
type Animate struct {
	// 动效结束展示icon
	Icon string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
	// 7日活动动画
	Json string `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
	// s10活动svg动画
	Svg string `protobuf:"bytes,3,opt,name=svg,proto3" json:"svg,omitempty"`
	// 循环次数(默认0不返回 表示无限循环)
	Loop                 int32    `protobuf:"varint,4,opt,name=loop,proto3" json:"loop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Animate) Reset()         { *m = Animate{} }
func (m *Animate) String() string { return proto.CompactTextString(m) }
func (*Animate) ProtoMessage()    {}
func (*Animate) Descriptor() ([]byte, []int) {
	return fileDescriptor_498ddd592723f790, []int{2}
}

func (m *Animate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Animate.Unmarshal(m, b)
}
func (m *Animate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Animate.Marshal(b, m, deterministic)
}
func (m *Animate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Animate.Merge(m, src)
}
func (m *Animate) XXX_Size() int {
	return xxx_messageInfo_Animate.Size(m)
}
func (m *Animate) XXX_DiscardUnknown() {
	xxx_messageInfo_Animate.DiscardUnknown(m)
}

var xxx_messageInfo_Animate proto.InternalMessageInfo

func (m *Animate) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Animate) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func (m *Animate) GetSvg() string {
	if m != nil {
		return m.Svg
	}
	return ""
}

func (m *Animate) GetLoop() int32 {
	if m != nil {
		return m.Loop
	}
	return 0
}

// 红点
type RedDot struct {
	// 红点类型
	// 1:纯红点 2:数字红点
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// 如果是数字红点 显示的数字
	Number               int32    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedDot) Reset()         { *m = RedDot{} }
func (m *RedDot) String() string { return proto.CompactTextString(m) }
func (*RedDot) ProtoMessage()    {}
func (*RedDot) Descriptor() ([]byte, []int) {
	return fileDescriptor_498ddd592723f790, []int{3}
}

func (m *RedDot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedDot.Unmarshal(m, b)
}
func (m *RedDot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedDot.Marshal(b, m, deterministic)
}
func (m *RedDot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedDot.Merge(m, src)
}
func (m *RedDot) XXX_Size() int {
	return xxx_messageInfo_RedDot.Size(m)
}
func (m *RedDot) XXX_DiscardUnknown() {
	xxx_messageInfo_RedDot.DiscardUnknown(m)
}

var xxx_messageInfo_RedDot proto.InternalMessageInfo

func (m *RedDot) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RedDot) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*TopActivityReply)(nil), "bilibili.broadcast.message.main.TopActivityReply")
	proto.RegisterType((*TopOnline)(nil), "bilibili.broadcast.message.main.TopOnline")
	proto.RegisterType((*Animate)(nil), "bilibili.broadcast.message.main.Animate")
	proto.RegisterType((*RedDot)(nil), "bilibili.broadcast.message.main.RedDot")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/broadcast/message/main/resource.proto", fileDescriptor_498ddd592723f790)
}

var fileDescriptor_498ddd592723f790 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6b, 0xdc, 0x30,
	0x10, 0xc5, 0x71, 0x76, 0xd7, 0xf6, 0x4e, 0x2e, 0x41, 0x87, 0x20, 0x36, 0x87, 0x2e, 0xbe, 0xd4,
	0xf4, 0x20, 0xb7, 0x69, 0xa1, 0xd7, 0x26, 0xb4, 0x87, 0x9e, 0x0a, 0x22, 0x50, 0xe8, 0x65, 0x91,
	0x6d, 0xc5, 0x51, 0xb0, 0x25, 0x55, 0x92, 0x0d, 0xfe, 0x26, 0xfd, 0xb8, 0x45, 0xf2, 0x9f, 0xa6,
	0xf4, 0xb0, 0x3d, 0x18, 0x3f, 0x0d, 0xbf, 0x19, 0x8d, 0x1e, 0x0f, 0x3e, 0x36, 0x46, 0x57, 0x27,
	0xa6, 0x45, 0x51, 0x8a, 0x56, 0xf8, 0xaf, 0x28, 0x8d, 0x62, 0x75, 0xc5, 0xac, 0x2b, 0x3a, 0x6e,
	0x2d, 0x6b, 0x78, 0xd1, 0x31, 0x21, 0x0b, 0xc3, 0xad, 0xea, 0x4d, 0xc5, 0x89, 0x36, 0xca, 0x29,
	0xf4, 0x6a, 0xe1, 0xc9, 0xca, 0x93, 0x99, 0x27, 0x9e, 0x3f, 0xdc, 0x34, 0x4a, 0x35, 0x2d, 0x2f,
	0x02, 0x5e, 0xf6, 0x8f, 0x05, 0xef, 0xb4, 0x1b, 0xa7, 0xee, 0xec, 0x19, 0xae, 0x1e, 0x94, 0xbe,
	0xab, 0x9c, 0x18, 0x84, 0x1b, 0x29, 0xd7, 0xed, 0x88, 0xee, 0x21, 0x56, 0xb2, 0x15, 0x92, 0xe3,
	0xe8, 0x18, 0xe5, 0x97, 0xb7, 0x6f, 0xc8, 0x99, 0x2b, 0xc8, 0x83, 0xd2, 0xdf, 0x42, 0x07, 0x9d,
	0x3b, 0x11, 0x82, 0xed, 0x13, 0xb3, 0x4f, 0xf8, 0xe2, 0x18, 0xe5, 0x7b, 0x1a, 0x74, 0xf6, 0xeb,
	0x02, 0xf6, 0x2b, 0xe9, 0x09, 0x37, 0xea, 0xe9, 0x8e, 0x1d, 0x0d, 0xda, 0xd7, 0x44, 0xa5, 0xe4,
	0xd2, 0xe5, 0x35, 0xba, 0x82, 0x4d, 0x6f, 0x04, 0xde, 0x84, 0x92, 0x97, 0xe8, 0x06, 0xf6, 0xbd,
	0x14, 0x3f, 0x7b, 0x7e, 0x12, 0x35, 0xde, 0x86, 0x7a, 0x3a, 0x15, 0xbe, 0xd6, 0xe8, 0x1e, 0x12,
	0x26, 0x45, 0xc7, 0x1c, 0xc7, 0xbb, 0xb0, 0x7d, 0x7e, 0x76, 0xfb, 0xbb, 0x89, 0xa7, 0x4b, 0x23,
	0xfa, 0x04, 0x89, 0xe1, 0xf5, 0xa9, 0x56, 0x0e, 0xc7, 0x61, 0xc6, 0xeb, 0xb3, 0x33, 0x28, 0xaf,
	0x3f, 0x2b, 0x47, 0x63, 0x13, 0xfe, 0xfe, 0x21, 0x92, 0x75, 0x1c, 0x27, 0xd3, 0x43, 0xbc, 0x46,
	0x07, 0x48, 0x85, 0x74, 0xdc, 0x0c, 0xac, 0xc5, 0xe9, 0x31, 0xca, 0x37, 0x74, 0x3d, 0x67, 0xdf,
	0x21, 0x99, 0xb7, 0x58, 0x3d, 0x88, 0x5e, 0x78, 0x80, 0x60, 0xfb, 0x6c, 0xff, 0xf8, 0xe2, 0xb5,
	0xf7, 0xc5, 0x0e, 0xcd, 0xe2, 0x8b, 0x1d, 0x1a, 0x4f, 0xb5, 0x4a, 0xe9, 0x60, 0xc9, 0x8e, 0x06,
	0x9d, 0x7d, 0x80, 0x98, 0xae, 0x2b, 0xfd, 0xe3, 0xf7, 0x35, 0xc4, 0xb2, 0xef, 0x4a, 0x6e, 0xc2,
	0xe4, 0x1d, 0x9d, 0x4f, 0xb7, 0x8f, 0x90, 0xd2, 0x39, 0x65, 0xe8, 0x07, 0x5c, 0xbe, 0x48, 0x08,
	0xba, 0x26, 0x53, 0x9c, 0xc8, 0x12, 0x27, 0xf2, 0xc5, 0xc7, 0xe9, 0xf0, 0xee, 0x7f, 0x42, 0xf2,
	0x57, 0xce, 0xde, 0x46, 0x65, 0x1c, 0x86, 0xbc, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x3c,
	0x7c, 0xc0, 0xfd, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceClient is the client API for Resource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceClient interface {
	//
	TopActivity(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Resource_TopActivityClient, error)
}

type resourceClient struct {
	cc *grpc.ClientConn
}

func NewResourceClient(cc *grpc.ClientConn) ResourceClient {
	return &resourceClient{cc}
}

func (c *resourceClient) TopActivity(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Resource_TopActivityClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Resource_serviceDesc.Streams[0], "/bilibili.broadcast.message.main.Resource/TopActivity", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceTopActivityClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Resource_TopActivityClient interface {
	Recv() (*TopActivityReply, error)
	grpc.ClientStream
}

type resourceTopActivityClient struct {
	grpc.ClientStream
}

func (x *resourceTopActivityClient) Recv() (*TopActivityReply, error) {
	m := new(TopActivityReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ResourceServer is the server API for Resource service.
type ResourceServer interface {
	//
	TopActivity(*emptypb.Empty, Resource_TopActivityServer) error
}

// UnimplementedResourceServer can be embedded to have forward compatible implementations.
type UnimplementedResourceServer struct {
}

func (*UnimplementedResourceServer) TopActivity(req *emptypb.Empty, srv Resource_TopActivityServer) error {
	return status.Errorf(codes.Unimplemented, "method TopActivity not implemented")
}

func RegisterResourceServer(s *grpc.Server, srv ResourceServer) {
	s.RegisterService(&_Resource_serviceDesc, srv)
}

func _Resource_TopActivity_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourceServer).TopActivity(m, &resourceTopActivityServer{stream})
}

type Resource_TopActivityServer interface {
	Send(*TopActivityReply) error
	grpc.ServerStream
}

type resourceTopActivityServer struct {
	grpc.ServerStream
}

func (x *resourceTopActivityServer) Send(m *TopActivityReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Resource_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.message.main.Resource",
	HandlerType: (*ResourceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TopActivity",
			Handler:       _Resource_TopActivity_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc_api/bilibili/broadcast/message/main/resource.proto",
}