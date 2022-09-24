// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/main/common/arch/doll/v1/doll.proto

package bilibili_main_common_arch_doll_v1

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

//
type ErrorRequest struct {
	//
	Error int32 `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	//
	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	//
	Delay                int64    `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorRequest) Reset()         { *m = ErrorRequest{} }
func (m *ErrorRequest) String() string { return proto.CompactTextString(m) }
func (*ErrorRequest) ProtoMessage()    {}
func (*ErrorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7695403d02236372, []int{0}
}

func (m *ErrorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorRequest.Unmarshal(m, b)
}
func (m *ErrorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorRequest.Marshal(b, m, deterministic)
}
func (m *ErrorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorRequest.Merge(m, src)
}
func (m *ErrorRequest) XXX_Size() int {
	return xxx_messageInfo_ErrorRequest.Size(m)
}
func (m *ErrorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorRequest proto.InternalMessageInfo

func (m *ErrorRequest) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *ErrorRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ErrorRequest) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

//
type ErrorResponse struct {
	//
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	//
	Time                 int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7695403d02236372, []int{1}
}

func (m *ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorResponse.Unmarshal(m, b)
}
func (m *ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorResponse.Marshal(b, m, deterministic)
}
func (m *ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorResponse.Merge(m, src)
}
func (m *ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_ErrorResponse.Size(m)
}
func (m *ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorResponse proto.InternalMessageInfo

func (m *ErrorResponse) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ErrorResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

//
type PingRequest struct {
	//
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7695403d02236372, []int{2}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

//
type PingResponse struct {
	//
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	//
	Time                 int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7695403d02236372, []int{3}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PingResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

//
type SayRequest struct {
	//
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayRequest) Reset()         { *m = SayRequest{} }
func (m *SayRequest) String() string { return proto.CompactTextString(m) }
func (*SayRequest) ProtoMessage()    {}
func (*SayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7695403d02236372, []int{4}
}

func (m *SayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayRequest.Unmarshal(m, b)
}
func (m *SayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayRequest.Marshal(b, m, deterministic)
}
func (m *SayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayRequest.Merge(m, src)
}
func (m *SayRequest) XXX_Size() int {
	return xxx_messageInfo_SayRequest.Size(m)
}
func (m *SayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayRequest proto.InternalMessageInfo

func (m *SayRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

//
type SayResponse struct {
	//
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	//
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	//
	Time                 int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayResponse) Reset()         { *m = SayResponse{} }
func (m *SayResponse) String() string { return proto.CompactTextString(m) }
func (*SayResponse) ProtoMessage()    {}
func (*SayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7695403d02236372, []int{5}
}

func (m *SayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayResponse.Unmarshal(m, b)
}
func (m *SayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayResponse.Marshal(b, m, deterministic)
}
func (m *SayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayResponse.Merge(m, src)
}
func (m *SayResponse) XXX_Size() int {
	return xxx_messageInfo_SayResponse.Size(m)
}
func (m *SayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayResponse proto.InternalMessageInfo

func (m *SayResponse) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SayResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SayResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*ErrorRequest)(nil), "bilibili.main.common.arch.doll.v1.ErrorRequest")
	proto.RegisterType((*ErrorResponse)(nil), "bilibili.main.common.arch.doll.v1.ErrorResponse")
	proto.RegisterType((*PingRequest)(nil), "bilibili.main.common.arch.doll.v1.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "bilibili.main.common.arch.doll.v1.PingResponse")
	proto.RegisterType((*SayRequest)(nil), "bilibili.main.common.arch.doll.v1.SayRequest")
	proto.RegisterType((*SayResponse)(nil), "bilibili.main.common.arch.doll.v1.SayResponse")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/main/common/arch/doll/v1/doll.proto", fileDescriptor_7695403d02236372)
}

var fileDescriptor_7695403d02236372 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4e, 0x83, 0x30,
	0x1c, 0xc7, 0xc3, 0x3f, 0x8d, 0xbf, 0xcd, 0x4b, 0xe3, 0x81, 0x70, 0xda, 0x38, 0x98, 0x5d, 0x6c,
	0x9d, 0x46, 0x7d, 0x82, 0x5d, 0xd5, 0xb0, 0x07, 0x30, 0x1d, 0x34, 0x50, 0x03, 0x2d, 0x16, 0x5c,
	0xc2, 0xe3, 0xf9, 0x66, 0xa6, 0x05, 0x1c, 0x26, 0x8b, 0x76, 0x07, 0x42, 0x7f, 0xcd, 0xe7, 0xfb,
	0x27, 0x6d, 0xe1, 0x21, 0x57, 0x75, 0xfa, 0x46, 0x6b, 0x4e, 0x76, 0xbc, 0xe4, 0xfa, 0x23, 0x15,
	0xe5, 0x82, 0xa4, 0xb2, 0xaa, 0xa4, 0x20, 0x54, 0xa5, 0x05, 0xc9, 0x64, 0x59, 0x92, 0xfd, 0xda,
	0xfc, 0x71, 0xad, 0x64, 0x2b, 0xd1, 0x72, 0xa4, 0xb1, 0xa6, 0x71, 0x4f, 0x63, 0x4d, 0x63, 0x43,
	0xed, 0xd7, 0xf1, 0x33, 0xcc, 0x37, 0x4a, 0x49, 0x95, 0xb0, 0x8f, 0x4f, 0xd6, 0xb4, 0xe8, 0x0a,
	0x02, 0xa6, 0xe7, 0xd0, 0x5d, 0x38, 0xab, 0x20, 0xe9, 0x07, 0x84, 0xc0, 0x6f, 0x79, 0xc5, 0x42,
	0x67, 0xe1, 0xac, 0xbc, 0xc4, 0xac, 0x35, 0x99, 0xb1, 0x92, 0x76, 0xa1, 0x67, 0x36, 0xfb, 0x21,
	0x7e, 0x82, 0xcb, 0xc1, 0xaf, 0xa9, 0xa5, 0x68, 0x98, 0x96, 0x16, 0xb2, 0x69, 0x8d, 0xf4, 0x22,
	0x31, 0xeb, 0x1f, 0x3b, 0xef, 0x60, 0x17, 0x2f, 0x61, 0xf6, 0xca, 0x45, 0x3e, 0xf6, 0x38, 0x92,
	0x18, 0x3f, 0xc2, 0xbc, 0x47, 0x4e, 0xb4, 0xbe, 0x06, 0xd8, 0xd2, 0x6e, 0x74, 0x0e, 0xe1, 0x3c,
	0x95, 0xa2, 0x65, 0x62, 0x14, 0x8e, 0x63, 0xfc, 0x02, 0x33, 0xc3, 0xfd, 0x61, 0x3f, 0x11, 0xbb,
	0xbf, 0xc4, 0xc7, 0x82, 0xef, 0xbe, 0x5c, 0xf0, 0x37, 0x69, 0x21, 0x51, 0x0e, 0xbe, 0x6e, 0x8e,
	0x30, 0xfe, 0xf7, 0x46, 0xf0, 0xe4, 0x14, 0x22, 0x62, 0xcd, 0x0f, 0x9d, 0x33, 0xf0, 0xb6, 0xb4,
	0x43, 0x37, 0x16, 0xba, 0xc3, 0x91, 0x44, 0xd8, 0x16, 0x1f, 0x52, 0xde, 0x21, 0x30, 0x97, 0x8c,
	0x6c, 0xfa, 0x4d, 0x9f, 0x57, 0x74, 0x6b, 0x2f, 0xe8, 0xb3, 0x76, 0x67, 0xe6, 0x29, 0xdf, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x62, 0x94, 0x9c, 0x03, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	//
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	//
	Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error)
	//
	Error(ctx context.Context, in *ErrorRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/bilibili.main.common.arch.doll.v1.Echo/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := c.cc.Invoke(ctx, "/bilibili.main.common.arch.doll.v1.Echo/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) Error(ctx context.Context, in *ErrorRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/bilibili.main.common.arch.doll.v1.Echo/Error", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	//
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	//
	Say(context.Context, *SayRequest) (*SayResponse, error)
	//
	Error(context.Context, *ErrorRequest) (*ErrorResponse, error)
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedEchoServer) Say(ctx context.Context, req *SayRequest) (*SayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}
func (*UnimplementedEchoServer) Error(ctx context.Context, req *ErrorRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Error not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.common.arch.doll.v1.Echo/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.common.arch.doll.v1.Echo/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Say(ctx, req.(*SayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_Error_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Error(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.common.arch.doll.v1.Echo/Error",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Error(ctx, req.(*ErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.main.common.arch.doll.v1.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Echo_Ping_Handler,
		},
		{
			MethodName: "Say",
			Handler:    _Echo_Say_Handler,
		},
		{
			MethodName: "Error",
			Handler:    _Echo_Error_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_api/bilibili/main/common/arch/doll/v1/doll.proto",
}