// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/broadcast/message/main/native.proto

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
type NativePageEvent struct {
	// Native页ID
	PageID int64 `protobuf:"varint,1,opt,name=PageID,proto3" json:"PageID,omitempty"`
	//
	Items                []*EventItem `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NativePageEvent) Reset()         { *m = NativePageEvent{} }
func (m *NativePageEvent) String() string { return proto.CompactTextString(m) }
func (*NativePageEvent) ProtoMessage()    {}
func (*NativePageEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_80fa7161b92b1a74, []int{0}
}

func (m *NativePageEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NativePageEvent.Unmarshal(m, b)
}
func (m *NativePageEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NativePageEvent.Marshal(b, m, deterministic)
}
func (m *NativePageEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NativePageEvent.Merge(m, src)
}
func (m *NativePageEvent) XXX_Size() int {
	return xxx_messageInfo_NativePageEvent.Size(m)
}
func (m *NativePageEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_NativePageEvent.DiscardUnknown(m)
}

var xxx_messageInfo_NativePageEvent proto.InternalMessageInfo

func (m *NativePageEvent) GetPageID() int64 {
	if m != nil {
		return m.PageID
	}
	return 0
}

func (m *NativePageEvent) GetItems() []*EventItem {
	if m != nil {
		return m.Items
	}
	return nil
}

//
type EventItem struct {
	// 组件标识
	ItemID int64 `protobuf:"varint,1,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
	// 组件类型
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	// 进度条数值
	Num int64 `protobuf:"varint,3,opt,name=Num,proto3" json:"Num,omitempty"`
	// 进度条展示数值
	DisplayNum string `protobuf:"bytes,4,opt,name=DisplayNum,proto3" json:"DisplayNum,omitempty"`
	// h5的组件标识
	WebKey string `protobuf:"bytes,5,opt,name=WebKey,proto3" json:"WebKey,omitempty"`
	// 活动统计维度
	// 0:用户维度 1:规则维度
	Dimension            int64    `protobuf:"varint,6,opt,name=dimension,proto3" json:"dimension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventItem) Reset()         { *m = EventItem{} }
func (m *EventItem) String() string { return proto.CompactTextString(m) }
func (*EventItem) ProtoMessage()    {}
func (*EventItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_80fa7161b92b1a74, []int{1}
}

func (m *EventItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventItem.Unmarshal(m, b)
}
func (m *EventItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventItem.Marshal(b, m, deterministic)
}
func (m *EventItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventItem.Merge(m, src)
}
func (m *EventItem) XXX_Size() int {
	return xxx_messageInfo_EventItem.Size(m)
}
func (m *EventItem) XXX_DiscardUnknown() {
	xxx_messageInfo_EventItem.DiscardUnknown(m)
}

var xxx_messageInfo_EventItem proto.InternalMessageInfo

func (m *EventItem) GetItemID() int64 {
	if m != nil {
		return m.ItemID
	}
	return 0
}

func (m *EventItem) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EventItem) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *EventItem) GetDisplayNum() string {
	if m != nil {
		return m.DisplayNum
	}
	return ""
}

func (m *EventItem) GetWebKey() string {
	if m != nil {
		return m.WebKey
	}
	return ""
}

func (m *EventItem) GetDimension() int64 {
	if m != nil {
		return m.Dimension
	}
	return 0
}

func init() {
	proto.RegisterType((*NativePageEvent)(nil), "bilibili.broadcast.message.main.NativePageEvent")
	proto.RegisterType((*EventItem)(nil), "bilibili.broadcast.message.main.EventItem")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/broadcast/message/main/native.proto", fileDescriptor_80fa7161b92b1a74)
}

var fileDescriptor_80fa7161b92b1a74 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x6b, 0x3a, 0x31,
	0x10, 0xc5, 0x59, 0x57, 0x05, 0xc7, 0xc3, 0xff, 0x4f, 0x0e, 0x12, 0x6c, 0x69, 0xc5, 0x93, 0xf4,
	0x90, 0x88, 0xa5, 0xf7, 0x1e, 0xf4, 0x20, 0x05, 0x29, 0x52, 0x90, 0x9e, 0x4a, 0x56, 0xc7, 0x34,
	0xd4, 0x6c, 0xc2, 0x26, 0x0a, 0xfb, 0x69, 0xfa, 0x55, 0x4b, 0xb2, 0xab, 0x5b, 0x7a, 0xf1, 0x10,
	0x32, 0x6f, 0xe6, 0x37, 0x6f, 0xe0, 0xc1, 0x93, 0x2c, 0xec, 0xf6, 0x43, 0x58, 0xc5, 0x33, 0x75,
	0x50, 0xe1, 0xf1, 0xac, 0x30, 0x62, 0xb7, 0x15, 0xce, 0x73, 0x8d, 0xce, 0x09, 0x89, 0x5c, 0x0b,
	0x95, 0xf3, 0x5c, 0x78, 0x75, 0x42, 0x66, 0x0b, 0xe3, 0x0d, 0xb9, 0x3f, 0xd3, 0xec, 0x42, 0xb3,
	0x9a, 0x66, 0x81, 0x1e, 0xde, 0x48, 0x63, 0xe4, 0x01, 0x79, 0xc4, 0xb3, 0xe3, 0x9e, 0xa3, 0xb6,
	0xbe, 0xac, 0xb6, 0xc7, 0x5f, 0xf0, 0x6f, 0x15, 0xdd, 0x5e, 0x85, 0xc4, 0xc5, 0x09, 0x73, 0x4f,
	0x06, 0xd0, 0x0d, 0x62, 0x39, 0xa7, 0xc9, 0x28, 0x99, 0xa4, 0xeb, 0x5a, 0x91, 0x67, 0xe8, 0x2c,
	0x3d, 0x6a, 0x47, 0x5b, 0xa3, 0x74, 0xd2, 0x9f, 0x3d, 0xb0, 0x2b, 0x87, 0x59, 0xb4, 0x0b, 0x2b,
	0xeb, 0x6a, 0x71, 0xfc, 0x9d, 0x40, 0xef, 0xd2, 0x0c, 0x77, 0xc2, 0xdf, 0xdc, 0xa9, 0x14, 0x21,
	0xd0, 0x7e, 0x2b, 0x2d, 0xd2, 0xd6, 0x28, 0x99, 0xf4, 0xd6, 0xb1, 0x26, 0xff, 0x21, 0x5d, 0x1d,
	0x35, 0x4d, 0x23, 0x18, 0x4a, 0x72, 0x07, 0x30, 0x57, 0xce, 0x1e, 0x44, 0x19, 0x06, 0xed, 0xc8,
	0xfe, 0xea, 0x04, 0xf7, 0x0d, 0x66, 0x2f, 0x58, 0xd2, 0x4e, 0x9c, 0xd5, 0x8a, 0xdc, 0x42, 0x6f,
	0xa7, 0x34, 0xe6, 0x4e, 0x99, 0x9c, 0x76, 0xa3, 0x5f, 0xd3, 0x98, 0x49, 0x80, 0x26, 0x0e, 0xf2,
	0x0e, 0xfd, 0x8d, 0xf0, 0xdb, 0xcf, 0x95, 0xf1, 0x6a, 0x5f, 0x92, 0x01, 0xab, 0x92, 0x64, 0xe7,
	0x24, 0xd9, 0x22, 0x24, 0x39, 0x9c, 0x5e, 0x4d, 0xe2, 0x4f, 0xc4, 0xd3, 0x24, 0xeb, 0x46, 0x8f,
	0xc7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x83, 0x0b, 0x53, 0xf5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NativePageClient is the client API for NativePage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NativePageClient interface {
	//
	WatchNotify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (NativePage_WatchNotifyClient, error)
}

type nativePageClient struct {
	cc *grpc.ClientConn
}

func NewNativePageClient(cc *grpc.ClientConn) NativePageClient {
	return &nativePageClient{cc}
}

func (c *nativePageClient) WatchNotify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (NativePage_WatchNotifyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NativePage_serviceDesc.Streams[0], "/bilibili.broadcast.message.main.NativePage/WatchNotify", opts...)
	if err != nil {
		return nil, err
	}
	x := &nativePageWatchNotifyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NativePage_WatchNotifyClient interface {
	Recv() (*NativePageEvent, error)
	grpc.ClientStream
}

type nativePageWatchNotifyClient struct {
	grpc.ClientStream
}

func (x *nativePageWatchNotifyClient) Recv() (*NativePageEvent, error) {
	m := new(NativePageEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NativePageServer is the server API for NativePage service.
type NativePageServer interface {
	//
	WatchNotify(*emptypb.Empty, NativePage_WatchNotifyServer) error
}

// UnimplementedNativePageServer can be embedded to have forward compatible implementations.
type UnimplementedNativePageServer struct {
}

func (*UnimplementedNativePageServer) WatchNotify(req *emptypb.Empty, srv NativePage_WatchNotifyServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchNotify not implemented")
}

func RegisterNativePageServer(s *grpc.Server, srv NativePageServer) {
	s.RegisterService(&_NativePage_serviceDesc, srv)
}

func _NativePage_WatchNotify_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NativePageServer).WatchNotify(m, &nativePageWatchNotifyServer{stream})
}

type NativePage_WatchNotifyServer interface {
	Send(*NativePageEvent) error
	grpc.ServerStream
}

type nativePageWatchNotifyServer struct {
	grpc.ServerStream
}

func (x *nativePageWatchNotifyServer) Send(m *NativePageEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _NativePage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.message.main.NativePage",
	HandlerType: (*NativePageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchNotify",
			Handler:       _NativePage_WatchNotify_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc_api/bilibili/broadcast/message/main/native.proto",
}
