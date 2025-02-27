// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/api/probe/v1/probe.proto

package bilibili_api_probe_v1

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
type CodeReq struct {
	//
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeReq) Reset()         { *m = CodeReq{} }
func (m *CodeReq) String() string { return proto.CompactTextString(m) }
func (*CodeReq) ProtoMessage()    {}
func (*CodeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f95f8b4708a0a9, []int{0}
}

func (m *CodeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeReq.Unmarshal(m, b)
}
func (m *CodeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeReq.Marshal(b, m, deterministic)
}
func (m *CodeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeReq.Merge(m, src)
}
func (m *CodeReq) XXX_Size() int {
	return xxx_messageInfo_CodeReq.Size(m)
}
func (m *CodeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeReq.DiscardUnknown(m)
}

var xxx_messageInfo_CodeReq proto.InternalMessageInfo

func (m *CodeReq) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

//
type CodeReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeReply) Reset()         { *m = CodeReply{} }
func (m *CodeReply) String() string { return proto.CompactTextString(m) }
func (*CodeReply) ProtoMessage()    {}
func (*CodeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f95f8b4708a0a9, []int{1}
}

func (m *CodeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeReply.Unmarshal(m, b)
}
func (m *CodeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeReply.Marshal(b, m, deterministic)
}
func (m *CodeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeReply.Merge(m, src)
}
func (m *CodeReply) XXX_Size() int {
	return xxx_messageInfo_CodeReply.Size(m)
}
func (m *CodeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeReply.DiscardUnknown(m)
}

var xxx_messageInfo_CodeReply proto.InternalMessageInfo

//
type ProbeReq struct {
	//
	Mid int64 `protobuf:"varint,1,opt,name=mid,proto3" json:"mid,omitempty"`
	//
	Buvid                string   `protobuf:"bytes,2,opt,name=buvid,proto3" json:"buvid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeReq) Reset()         { *m = ProbeReq{} }
func (m *ProbeReq) String() string { return proto.CompactTextString(m) }
func (*ProbeReq) ProtoMessage()    {}
func (*ProbeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f95f8b4708a0a9, []int{2}
}

func (m *ProbeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeReq.Unmarshal(m, b)
}
func (m *ProbeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeReq.Marshal(b, m, deterministic)
}
func (m *ProbeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeReq.Merge(m, src)
}
func (m *ProbeReq) XXX_Size() int {
	return xxx_messageInfo_ProbeReq.Size(m)
}
func (m *ProbeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeReq proto.InternalMessageInfo

func (m *ProbeReq) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *ProbeReq) GetBuvid() string {
	if m != nil {
		return m.Buvid
	}
	return ""
}

//
type ProbeReply struct {
	//
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	//
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeReply) Reset()         { *m = ProbeReply{} }
func (m *ProbeReply) String() string { return proto.CompactTextString(m) }
func (*ProbeReply) ProtoMessage()    {}
func (*ProbeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f95f8b4708a0a9, []int{3}
}

func (m *ProbeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeReply.Unmarshal(m, b)
}
func (m *ProbeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeReply.Marshal(b, m, deterministic)
}
func (m *ProbeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeReply.Merge(m, src)
}
func (m *ProbeReply) XXX_Size() int {
	return xxx_messageInfo_ProbeReply.Size(m)
}
func (m *ProbeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeReply proto.InternalMessageInfo

func (m *ProbeReply) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ProbeReply) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

//
type ProbeStreamReq struct {
	//
	Mid int64 `protobuf:"varint,1,opt,name=mid,proto3" json:"mid,omitempty"`
	//
	Sequence             int64    `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeStreamReq) Reset()         { *m = ProbeStreamReq{} }
func (m *ProbeStreamReq) String() string { return proto.CompactTextString(m) }
func (*ProbeStreamReq) ProtoMessage()    {}
func (*ProbeStreamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f95f8b4708a0a9, []int{4}
}

func (m *ProbeStreamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeStreamReq.Unmarshal(m, b)
}
func (m *ProbeStreamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeStreamReq.Marshal(b, m, deterministic)
}
func (m *ProbeStreamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeStreamReq.Merge(m, src)
}
func (m *ProbeStreamReq) XXX_Size() int {
	return xxx_messageInfo_ProbeStreamReq.Size(m)
}
func (m *ProbeStreamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeStreamReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeStreamReq proto.InternalMessageInfo

func (m *ProbeStreamReq) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *ProbeStreamReq) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

//
type ProbeStreamReply struct {
	//
	Sequence int64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	//
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	//
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeStreamReply) Reset()         { *m = ProbeStreamReply{} }
func (m *ProbeStreamReply) String() string { return proto.CompactTextString(m) }
func (*ProbeStreamReply) ProtoMessage()    {}
func (*ProbeStreamReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f95f8b4708a0a9, []int{5}
}

func (m *ProbeStreamReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeStreamReply.Unmarshal(m, b)
}
func (m *ProbeStreamReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeStreamReply.Marshal(b, m, deterministic)
}
func (m *ProbeStreamReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeStreamReply.Merge(m, src)
}
func (m *ProbeStreamReply) XXX_Size() int {
	return xxx_messageInfo_ProbeStreamReply.Size(m)
}
func (m *ProbeStreamReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeStreamReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeStreamReply proto.InternalMessageInfo

func (m *ProbeStreamReply) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *ProbeStreamReply) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ProbeStreamReply) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

//
type ProbeSubReq struct {
	//
	Buvid                int64    `protobuf:"varint,1,opt,name=buvid,proto3" json:"buvid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeSubReq) Reset()         { *m = ProbeSubReq{} }
func (m *ProbeSubReq) String() string { return proto.CompactTextString(m) }
func (*ProbeSubReq) ProtoMessage()    {}
func (*ProbeSubReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f95f8b4708a0a9, []int{6}
}

func (m *ProbeSubReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeSubReq.Unmarshal(m, b)
}
func (m *ProbeSubReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeSubReq.Marshal(b, m, deterministic)
}
func (m *ProbeSubReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeSubReq.Merge(m, src)
}
func (m *ProbeSubReq) XXX_Size() int {
	return xxx_messageInfo_ProbeSubReq.Size(m)
}
func (m *ProbeSubReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeSubReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeSubReq proto.InternalMessageInfo

func (m *ProbeSubReq) GetBuvid() int64 {
	if m != nil {
		return m.Buvid
	}
	return 0
}

//
type ProbeSubReply struct {
	//
	MessageId            int64    `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeSubReply) Reset()         { *m = ProbeSubReply{} }
func (m *ProbeSubReply) String() string { return proto.CompactTextString(m) }
func (*ProbeSubReply) ProtoMessage()    {}
func (*ProbeSubReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f95f8b4708a0a9, []int{7}
}

func (m *ProbeSubReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeSubReply.Unmarshal(m, b)
}
func (m *ProbeSubReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeSubReply.Marshal(b, m, deterministic)
}
func (m *ProbeSubReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeSubReply.Merge(m, src)
}
func (m *ProbeSubReply) XXX_Size() int {
	return xxx_messageInfo_ProbeSubReply.Size(m)
}
func (m *ProbeSubReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeSubReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeSubReply proto.InternalMessageInfo

func (m *ProbeSubReply) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func init() {
	proto.RegisterType((*CodeReq)(nil), "bilibili.api.probe.v1.CodeReq")
	proto.RegisterType((*CodeReply)(nil), "bilibili.api.probe.v1.CodeReply")
	proto.RegisterType((*ProbeReq)(nil), "bilibili.api.probe.v1.ProbeReq")
	proto.RegisterType((*ProbeReply)(nil), "bilibili.api.probe.v1.ProbeReply")
	proto.RegisterType((*ProbeStreamReq)(nil), "bilibili.api.probe.v1.ProbeStreamReq")
	proto.RegisterType((*ProbeStreamReply)(nil), "bilibili.api.probe.v1.ProbeStreamReply")
	proto.RegisterType((*ProbeSubReq)(nil), "bilibili.api.probe.v1.ProbeSubReq")
	proto.RegisterType((*ProbeSubReply)(nil), "bilibili.api.probe.v1.ProbeSubReply")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/api/probe/v1/probe.proto", fileDescriptor_e7f95f8b4708a0a9)
}

var fileDescriptor_e7f95f8b4708a0a9 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x49, 0xf3, 0xf6, 0x6d, 0x33, 0x45, 0x29, 0x8b, 0x42, 0x09, 0x56, 0xeb, 0xaa, 0x58,
	0x3c, 0xa4, 0xb4, 0xde, 0xbd, 0xe8, 0x45, 0x50, 0x90, 0xd4, 0xa3, 0x50, 0xf2, 0x67, 0x2c, 0x81,
	0xa4, 0xd9, 0x26, 0x9b, 0x42, 0x3f, 0xa9, 0x5f, 0x47, 0x66, 0x37, 0x4d, 0x5b, 0xa8, 0xd1, 0x43,
	0x60, 0x67, 0xf7, 0xf7, 0x3c, 0xfb, 0xec, 0x0c, 0x81, 0xbb, 0x79, 0x26, 0x82, 0x99, 0x27, 0xa2,
	0x91, 0x1f, 0xc5, 0x11, 0x7d, 0x23, 0x2a, 0x44, 0x96, 0xfa, 0x38, 0x5a, 0x8d, 0xf5, 0xc2, 0x11,
	0x59, 0x2a, 0x53, 0x76, 0xba, 0x41, 0x1c, 0x4f, 0x44, 0x8e, 0x3e, 0x59, 0x8d, 0x79, 0x1f, 0x5a,
	0x8f, 0x69, 0x88, 0x2e, 0x2e, 0x19, 0x83, 0x7f, 0x41, 0x1a, 0x62, 0xcf, 0x18, 0x18, 0x43, 0xd3,
	0x55, 0x6b, 0xde, 0x01, 0x4b, 0x1f, 0x8b, 0x78, 0xcd, 0x27, 0xd0, 0x7e, 0x23, 0x1d, 0xc1, 0x5d,
	0x30, 0x93, 0x28, 0x2c, 0x59, 0x5a, 0xb2, 0x13, 0x68, 0xfa, 0xc5, 0x2a, 0x0a, 0x7b, 0x8d, 0x81,
	0x31, 0xb4, 0x5c, 0x5d, 0xf0, 0x27, 0x80, 0x52, 0x23, 0xe2, 0x35, 0xeb, 0x41, 0x2b, 0x48, 0x17,
	0x12, 0x17, 0x52, 0x29, 0x2d, 0x77, 0x53, 0xb2, 0x33, 0xb0, 0x64, 0x94, 0x60, 0x2e, 0xbd, 0x44,
	0x28, 0x07, 0xd3, 0xdd, 0x6e, 0xf0, 0x07, 0x38, 0x56, 0x2e, 0x53, 0x99, 0xa1, 0x97, 0x1c, 0xbe,
	0xdf, 0x86, 0x76, 0x8e, 0xcb, 0x02, 0x17, 0x01, 0x96, 0x06, 0x55, 0xcd, 0x3f, 0xa1, 0xbb, 0xa7,
	0xa7, 0x2c, 0xbb, 0xbc, 0xb1, 0xcf, 0xd7, 0xa7, 0xd9, 0x7d, 0x85, 0xb9, 0xf7, 0x0a, 0x7e, 0x05,
	0x1d, 0x7d, 0x4f, 0xe1, 0x53, 0xc8, 0xaa, 0x25, 0xda, 0xbf, 0x6c, 0x89, 0x03, 0x47, 0x5b, 0x88,
	0x92, 0xf4, 0x01, 0x12, 0xcc, 0x73, 0x6f, 0x8e, 0xb3, 0x8a, 0xb5, 0xca, 0x9d, 0xe7, 0x70, 0xf2,
	0xd5, 0x80, 0xa6, 0x12, 0xb0, 0x17, 0x68, 0xbf, 0x63, 0x2e, 0x69, 0x22, 0xec, 0xdc, 0x39, 0x38,
	0x50, 0xa7, 0x9c, 0xa6, 0x3d, 0xa8, 0x3d, 0xa7, 0x6b, 0x5f, 0xa1, 0x45, 0x6e, 0x14, 0xf4, 0xe2,
	0x07, 0x78, 0x33, 0x6e, 0xfb, 0xb2, 0x1e, 0x20, 0xbb, 0x0f, 0x00, 0xb2, 0xd3, 0x2d, 0x66, 0x37,
	0x75, 0x82, 0x6a, 0x8c, 0xf6, 0xed, 0x5f, 0x30, 0x72, 0x9f, 0xea, 0xb0, 0xd3, 0xc2, 0x67, 0xbc,
	0x56, 0xa3, 0x3a, 0x6f, 0x5f, 0xff, 0xca, 0x88, 0x78, 0xed, 0xff, 0x57, 0xbf, 0xc6, 0xfd, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0x70, 0xfa, 0x95, 0x48, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProbeClient is the client API for Probe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProbeClient interface {
	//
	TestCode(ctx context.Context, in *CodeReq, opts ...grpc.CallOption) (*CodeReply, error)
	//
	TestReq(ctx context.Context, in *ProbeReq, opts ...grpc.CallOption) (*ProbeReply, error)
	//
	TestStream(ctx context.Context, in *ProbeStreamReq, opts ...grpc.CallOption) (*ProbeStreamReply, error)
	//
	TestSub(ctx context.Context, in *ProbeSubReq, opts ...grpc.CallOption) (*ProbeSubReply, error)
}

type probeClient struct {
	cc *grpc.ClientConn
}

func NewProbeClient(cc *grpc.ClientConn) ProbeClient {
	return &probeClient{cc}
}

func (c *probeClient) TestCode(ctx context.Context, in *CodeReq, opts ...grpc.CallOption) (*CodeReply, error) {
	out := new(CodeReply)
	err := c.cc.Invoke(ctx, "/bilibili.api.probe.v1.Probe/TestCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *probeClient) TestReq(ctx context.Context, in *ProbeReq, opts ...grpc.CallOption) (*ProbeReply, error) {
	out := new(ProbeReply)
	err := c.cc.Invoke(ctx, "/bilibili.api.probe.v1.Probe/TestReq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *probeClient) TestStream(ctx context.Context, in *ProbeStreamReq, opts ...grpc.CallOption) (*ProbeStreamReply, error) {
	out := new(ProbeStreamReply)
	err := c.cc.Invoke(ctx, "/bilibili.api.probe.v1.Probe/TestStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *probeClient) TestSub(ctx context.Context, in *ProbeSubReq, opts ...grpc.CallOption) (*ProbeSubReply, error) {
	out := new(ProbeSubReply)
	err := c.cc.Invoke(ctx, "/bilibili.api.probe.v1.Probe/TestSub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProbeServer is the server API for Probe service.
type ProbeServer interface {
	//
	TestCode(context.Context, *CodeReq) (*CodeReply, error)
	//
	TestReq(context.Context, *ProbeReq) (*ProbeReply, error)
	//
	TestStream(context.Context, *ProbeStreamReq) (*ProbeStreamReply, error)
	//
	TestSub(context.Context, *ProbeSubReq) (*ProbeSubReply, error)
}

// UnimplementedProbeServer can be embedded to have forward compatible implementations.
type UnimplementedProbeServer struct {
}

func (*UnimplementedProbeServer) TestCode(ctx context.Context, req *CodeReq) (*CodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestCode not implemented")
}
func (*UnimplementedProbeServer) TestReq(ctx context.Context, req *ProbeReq) (*ProbeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestReq not implemented")
}
func (*UnimplementedProbeServer) TestStream(ctx context.Context, req *ProbeStreamReq) (*ProbeStreamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestStream not implemented")
}
func (*UnimplementedProbeServer) TestSub(ctx context.Context, req *ProbeSubReq) (*ProbeSubReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestSub not implemented")
}

func RegisterProbeServer(s *grpc.Server, srv ProbeServer) {
	s.RegisterService(&_Probe_serviceDesc, srv)
}

func _Probe_TestCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProbeServer).TestCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.api.probe.v1.Probe/TestCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProbeServer).TestCode(ctx, req.(*CodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Probe_TestReq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProbeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProbeServer).TestReq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.api.probe.v1.Probe/TestReq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProbeServer).TestReq(ctx, req.(*ProbeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Probe_TestStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProbeStreamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProbeServer).TestStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.api.probe.v1.Probe/TestStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProbeServer).TestStream(ctx, req.(*ProbeStreamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Probe_TestSub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProbeSubReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProbeServer).TestSub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.api.probe.v1.Probe/TestSub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProbeServer).TestSub(ctx, req.(*ProbeSubReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Probe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.api.probe.v1.Probe",
	HandlerType: (*ProbeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestCode",
			Handler:    _Probe_TestCode_Handler,
		},
		{
			MethodName: "TestReq",
			Handler:    _Probe_TestReq_Handler,
		},
		{
			MethodName: "TestStream",
			Handler:    _Probe_TestStream_Handler,
		},
		{
			MethodName: "TestSub",
			Handler:    _Probe_TestSub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_api/bilibili/api/probe/v1/probe.proto",
}
