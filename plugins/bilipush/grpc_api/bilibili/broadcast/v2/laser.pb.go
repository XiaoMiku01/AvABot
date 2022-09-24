// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/broadcast/v2/laser.proto

package bilibili_broadcast_v2

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

// 服务端下发Laser事件
type LaserEventResp struct {
	// 任务id
	Taskid int64 `protobuf:"varint,1,opt,name=taskid,proto3" json:"taskid,omitempty"`
	// 指令名
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	// 指令参数json字符串
	Params               string   `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaserEventResp) Reset()         { *m = LaserEventResp{} }
func (m *LaserEventResp) String() string { return proto.CompactTextString(m) }
func (*LaserEventResp) ProtoMessage()    {}
func (*LaserEventResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5c15f357d08624b, []int{0}
}

func (m *LaserEventResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaserEventResp.Unmarshal(m, b)
}
func (m *LaserEventResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaserEventResp.Marshal(b, m, deterministic)
}
func (m *LaserEventResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaserEventResp.Merge(m, src)
}
func (m *LaserEventResp) XXX_Size() int {
	return xxx_messageInfo_LaserEventResp.Size(m)
}
func (m *LaserEventResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LaserEventResp.DiscardUnknown(m)
}

var xxx_messageInfo_LaserEventResp proto.InternalMessageInfo

func (m *LaserEventResp) GetTaskid() int64 {
	if m != nil {
		return m.Taskid
	}
	return 0
}

func (m *LaserEventResp) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *LaserEventResp) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func init() {
	proto.RegisterType((*LaserEventResp)(nil), "bilibili.broadcast.v2.LaserEventResp")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/broadcast/v2/laser.proto", fileDescriptor_e5c15f357d08624b)
}

var fileDescriptor_e5c15f357d08624b = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0x31, 0x4b, 0xc7, 0x30,
	0x14, 0xc4, 0x89, 0x7f, 0x2c, 0x98, 0xc1, 0x21, 0x60, 0x29, 0x75, 0x29, 0x82, 0x50, 0x1c, 0x5e,
	0xa4, 0x7e, 0x86, 0x6e, 0xba, 0x74, 0x50, 0x37, 0x79, 0x49, 0x63, 0x0d, 0xb6, 0x4d, 0x48, 0x62,
	0xc1, 0x6f, 0x2f, 0x49, 0x5a, 0x41, 0x70, 0x78, 0xc3, 0x1d, 0x77, 0xf7, 0x7e, 0xf4, 0x6e, 0x72,
	0x56, 0xbe, 0xa1, 0xd5, 0x5c, 0xe8, 0x59, 0xc7, 0xe3, 0xc2, 0x19, 0x1c, 0x25, 0xfa, 0xc0, 0xb7,
	0x8e, 0xcf, 0xe8, 0x95, 0x03, 0xeb, 0x4c, 0x30, 0xec, 0xea, 0x88, 0xc0, 0x6f, 0x04, 0xb6, 0xae,
	0xbe, 0x9e, 0x8c, 0x99, 0x66, 0xc5, 0x53, 0x48, 0x7c, 0xbd, 0x73, 0xb5, 0xd8, 0xf0, 0x9d, 0x3b,
	0x37, 0xaf, 0xf4, 0xf2, 0x31, 0x4e, 0xf4, 0x9b, 0x5a, 0xc3, 0xa0, 0xbc, 0x65, 0x25, 0x2d, 0x02,
	0xfa, 0x4f, 0x3d, 0x56, 0xa4, 0x21, 0xed, 0x69, 0xd8, 0x55, 0xf4, 0x51, 0x06, 0x6d, 0xd6, 0xea,
	0xac, 0x21, 0xed, 0xc5, 0xb0, 0xab, 0xe8, 0x5b, 0x74, 0xb8, 0xf8, 0xea, 0x94, 0xfd, 0xac, 0xba,
	0x67, 0x7a, 0x9e, 0x96, 0xd9, 0x13, 0xa5, 0x2f, 0x18, 0xe4, 0x47, 0x7a, 0xc1, 0x4a, 0xc8, 0x38,
	0x70, 0xe0, 0x40, 0x1f, 0x71, 0xea, 0x5b, 0xf8, 0x97, 0x1e, 0xfe, 0xd2, 0xdd, 0x13, 0x51, 0xa4,
	0xe2, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x19, 0x67, 0xdf, 0x1a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LaserClient is the client API for Laser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaserClient interface {
	// 监听Laser事件
	WatchEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Laser_WatchEventClient, error)
}

type laserClient struct {
	cc *grpc.ClientConn
}

func NewLaserClient(cc *grpc.ClientConn) LaserClient {
	return &laserClient{cc}
}

func (c *laserClient) WatchEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Laser_WatchEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Laser_serviceDesc.Streams[0], "/bilibili.broadcast.v2.Laser/WatchEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &laserWatchEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Laser_WatchEventClient interface {
	Recv() (*LaserEventResp, error)
	grpc.ClientStream
}

type laserWatchEventClient struct {
	grpc.ClientStream
}

func (x *laserWatchEventClient) Recv() (*LaserEventResp, error) {
	m := new(LaserEventResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LaserServer is the server API for Laser service.
type LaserServer interface {
	// 监听Laser事件
	WatchEvent(*emptypb.Empty, Laser_WatchEventServer) error
}

// UnimplementedLaserServer can be embedded to have forward compatible implementations.
type UnimplementedLaserServer struct {
}

func (*UnimplementedLaserServer) WatchEvent(req *emptypb.Empty, srv Laser_WatchEventServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchEvent not implemented")
}

func RegisterLaserServer(s *grpc.Server, srv LaserServer) {
	s.RegisterService(&_Laser_serviceDesc, srv)
}

func _Laser_WatchEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LaserServer).WatchEvent(m, &laserWatchEventServer{stream})
}

type Laser_WatchEventServer interface {
	Send(*LaserEventResp) error
	grpc.ServerStream
}

type laserWatchEventServer struct {
	grpc.ServerStream
}

func (x *laserWatchEventServer) Send(m *LaserEventResp) error {
	return x.ServerStream.SendMsg(m)
}

var _Laser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.v2.Laser",
	HandlerType: (*LaserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchEvent",
			Handler:       _Laser_WatchEvent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc_api/bilibili/broadcast/v2/laser.proto",
}