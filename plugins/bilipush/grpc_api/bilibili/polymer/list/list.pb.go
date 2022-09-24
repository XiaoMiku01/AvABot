// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/polymer/list/list.proto

package bilibili_polymer_list

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
type CheckAccountReply struct {
	//
	IsNew                bool     `protobuf:"varint,1,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAccountReply) Reset()         { *m = CheckAccountReply{} }
func (m *CheckAccountReply) String() string { return proto.CompactTextString(m) }
func (*CheckAccountReply) ProtoMessage()    {}
func (*CheckAccountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_98aa7b1f535fa257, []int{0}
}

func (m *CheckAccountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAccountReply.Unmarshal(m, b)
}
func (m *CheckAccountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAccountReply.Marshal(b, m, deterministic)
}
func (m *CheckAccountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAccountReply.Merge(m, src)
}
func (m *CheckAccountReply) XXX_Size() int {
	return xxx_messageInfo_CheckAccountReply.Size(m)
}
func (m *CheckAccountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAccountReply.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAccountReply proto.InternalMessageInfo

func (m *CheckAccountReply) GetIsNew() bool {
	if m != nil {
		return m.IsNew
	}
	return false
}

//
type CheckAccountReq struct {
	//
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	//
	Periods              string   `protobuf:"bytes,2,opt,name=periods,proto3" json:"periods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAccountReq) Reset()         { *m = CheckAccountReq{} }
func (m *CheckAccountReq) String() string { return proto.CompactTextString(m) }
func (*CheckAccountReq) ProtoMessage()    {}
func (*CheckAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_98aa7b1f535fa257, []int{1}
}

func (m *CheckAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAccountReq.Unmarshal(m, b)
}
func (m *CheckAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAccountReq.Marshal(b, m, deterministic)
}
func (m *CheckAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAccountReq.Merge(m, src)
}
func (m *CheckAccountReq) XXX_Size() int {
	return xxx_messageInfo_CheckAccountReq.Size(m)
}
func (m *CheckAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAccountReq proto.InternalMessageInfo

func (m *CheckAccountReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CheckAccountReq) GetPeriods() string {
	if m != nil {
		return m.Periods
	}
	return ""
}

//
type FavoriteTabItem struct {
	//
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	//
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FavoriteTabItem) Reset()         { *m = FavoriteTabItem{} }
func (m *FavoriteTabItem) String() string { return proto.CompactTextString(m) }
func (*FavoriteTabItem) ProtoMessage()    {}
func (*FavoriteTabItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_98aa7b1f535fa257, []int{2}
}

func (m *FavoriteTabItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FavoriteTabItem.Unmarshal(m, b)
}
func (m *FavoriteTabItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FavoriteTabItem.Marshal(b, m, deterministic)
}
func (m *FavoriteTabItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FavoriteTabItem.Merge(m, src)
}
func (m *FavoriteTabItem) XXX_Size() int {
	return xxx_messageInfo_FavoriteTabItem.Size(m)
}
func (m *FavoriteTabItem) XXX_DiscardUnknown() {
	xxx_messageInfo_FavoriteTabItem.DiscardUnknown(m)
}

var xxx_messageInfo_FavoriteTabItem proto.InternalMessageInfo

func (m *FavoriteTabItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FavoriteTabItem) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *FavoriteTabItem) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

//
type FavoriteTabReply struct {
	//
	Items                []*FavoriteTabItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FavoriteTabReply) Reset()         { *m = FavoriteTabReply{} }
func (m *FavoriteTabReply) String() string { return proto.CompactTextString(m) }
func (*FavoriteTabReply) ProtoMessage()    {}
func (*FavoriteTabReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_98aa7b1f535fa257, []int{3}
}

func (m *FavoriteTabReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FavoriteTabReply.Unmarshal(m, b)
}
func (m *FavoriteTabReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FavoriteTabReply.Marshal(b, m, deterministic)
}
func (m *FavoriteTabReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FavoriteTabReply.Merge(m, src)
}
func (m *FavoriteTabReply) XXX_Size() int {
	return xxx_messageInfo_FavoriteTabReply.Size(m)
}
func (m *FavoriteTabReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FavoriteTabReply.DiscardUnknown(m)
}

var xxx_messageInfo_FavoriteTabReply proto.InternalMessageInfo

func (m *FavoriteTabReply) GetItems() []*FavoriteTabItem {
	if m != nil {
		return m.Items
	}
	return nil
}

//
type FavoriteTabReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FavoriteTabReq) Reset()         { *m = FavoriteTabReq{} }
func (m *FavoriteTabReq) String() string { return proto.CompactTextString(m) }
func (*FavoriteTabReq) ProtoMessage()    {}
func (*FavoriteTabReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_98aa7b1f535fa257, []int{4}
}

func (m *FavoriteTabReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FavoriteTabReq.Unmarshal(m, b)
}
func (m *FavoriteTabReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FavoriteTabReq.Marshal(b, m, deterministic)
}
func (m *FavoriteTabReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FavoriteTabReq.Merge(m, src)
}
func (m *FavoriteTabReq) XXX_Size() int {
	return xxx_messageInfo_FavoriteTabReq.Size(m)
}
func (m *FavoriteTabReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FavoriteTabReq.DiscardUnknown(m)
}

var xxx_messageInfo_FavoriteTabReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CheckAccountReply)(nil), "bilibili.polymer.list.CheckAccountReply")
	proto.RegisterType((*CheckAccountReq)(nil), "bilibili.polymer.list.CheckAccountReq")
	proto.RegisterType((*FavoriteTabItem)(nil), "bilibili.polymer.list.FavoriteTabItem")
	proto.RegisterType((*FavoriteTabReply)(nil), "bilibili.polymer.list.FavoriteTabReply")
	proto.RegisterType((*FavoriteTabReq)(nil), "bilibili.polymer.list.FavoriteTabReq")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/polymer/list/list.proto", fileDescriptor_98aa7b1f535fa257)
}

var fileDescriptor_98aa7b1f535fa257 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xa9, 0xdd, 0xa6, 0xbb, 0x13, 0x37, 0x03, 0x83, 0xb2, 0xa7, 0x51, 0x50, 0xab, 0x0f,
	0x1d, 0xcc, 0x57, 0x7d, 0x10, 0x41, 0x10, 0x45, 0x24, 0xf8, 0x2a, 0xb3, 0xeb, 0x2e, 0x7a, 0xb1,
	0x5d, 0xd2, 0x24, 0x73, 0xf4, 0x17, 0xfa, 0xb7, 0x24, 0xe9, 0x06, 0xdd, 0x50, 0xdc, 0x43, 0xcb,
	0x09, 0xf7, 0x3b, 0x87, 0x93, 0x4b, 0xe0, 0xfc, 0x5d, 0xc9, 0x74, 0x92, 0x48, 0x1a, 0x4d, 0x29,
	0x23, 0xfb, 0x8d, 0xa4, 0xc8, 0xca, 0x1c, 0xd5, 0x28, 0x23, 0x6d, 0xdc, 0x2f, 0x96, 0x4a, 0x18,
	0xc1, 0xfa, 0x6b, 0x22, 0x5e, 0x11, 0xb1, 0x1d, 0x86, 0x17, 0x70, 0x7c, 0xfb, 0x81, 0xe9, 0xe7,
	0x4d, 0x9a, 0x8a, 0xc5, 0xdc, 0x70, 0x94, 0x59, 0xc9, 0xfa, 0xd0, 0x22, 0x3d, 0x99, 0xe3, 0x32,
	0xf0, 0x86, 0x5e, 0x74, 0xc0, 0x9b, 0xa4, 0x9f, 0x70, 0x19, 0x5e, 0x43, 0x77, 0x93, 0x2d, 0x58,
	0x0f, 0xfc, 0x05, 0xcd, 0x1c, 0xe6, 0x73, 0x2b, 0x59, 0x00, 0xfb, 0x12, 0x15, 0x89, 0x99, 0x0e,
	0xf6, 0x86, 0x5e, 0xd4, 0xe6, 0xeb, 0x63, 0xf8, 0x00, 0xdd, 0xbb, 0xe4, 0x4b, 0x28, 0x32, 0xf8,
	0x92, 0x4c, 0xef, 0x0d, 0xe6, 0x8c, 0x41, 0x63, 0x9e, 0xe4, 0xe8, 0xfc, 0x6d, 0xee, 0xb4, 0x8b,
	0x54, 0xb4, 0x32, 0x5b, 0x69, 0x29, 0x53, 0x4a, 0x0c, 0xfc, 0x8a, 0xb2, 0x3a, 0x7c, 0x86, 0x5e,
	0x2d, 0xac, 0xaa, 0x7d, 0x05, 0x4d, 0x32, 0x98, 0xeb, 0xc0, 0x1b, 0xfa, 0x51, 0x67, 0x7c, 0x1a,
	0xff, 0x7a, 0xe5, 0x78, 0xab, 0x04, 0xaf, 0x4c, 0x61, 0x0f, 0x8e, 0x36, 0x12, 0x8b, 0xf1, 0xb7,
	0x07, 0x8d, 0x47, 0xd2, 0x86, 0xbd, 0x42, 0xa7, 0x36, 0x62, 0x27, 0xff, 0x07, 0x73, 0x2c, 0x06,
	0x67, 0xbb, 0x60, 0xb6, 0xf7, 0x1b, 0x1c, 0xd6, 0xf7, 0xca, 0xfe, 0x2a, 0xbe, 0xb5, 0xfc, 0x41,
	0xb4, 0x13, 0x27, 0xb3, 0x72, 0xda, 0x72, 0x6f, 0xe0, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x25,
	0xc4, 0xad, 0x5f, 0x30, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ListClient is the client API for List service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListClient interface {
	//
	FavoriteTab(ctx context.Context, in *FavoriteTabReq, opts ...grpc.CallOption) (*FavoriteTabReply, error)
	//
	CheckAccount(ctx context.Context, in *CheckAccountReq, opts ...grpc.CallOption) (*CheckAccountReply, error)
}

type listClient struct {
	cc *grpc.ClientConn
}

func NewListClient(cc *grpc.ClientConn) ListClient {
	return &listClient{cc}
}

func (c *listClient) FavoriteTab(ctx context.Context, in *FavoriteTabReq, opts ...grpc.CallOption) (*FavoriteTabReply, error) {
	out := new(FavoriteTabReply)
	err := c.cc.Invoke(ctx, "/bilibili.polymer.list.List/FavoriteTab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) CheckAccount(ctx context.Context, in *CheckAccountReq, opts ...grpc.CallOption) (*CheckAccountReply, error) {
	out := new(CheckAccountReply)
	err := c.cc.Invoke(ctx, "/bilibili.polymer.list.List/CheckAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListServer is the server API for List service.
type ListServer interface {
	//
	FavoriteTab(context.Context, *FavoriteTabReq) (*FavoriteTabReply, error)
	//
	CheckAccount(context.Context, *CheckAccountReq) (*CheckAccountReply, error)
}

// UnimplementedListServer can be embedded to have forward compatible implementations.
type UnimplementedListServer struct {
}

func (*UnimplementedListServer) FavoriteTab(ctx context.Context, req *FavoriteTabReq) (*FavoriteTabReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteTab not implemented")
}
func (*UnimplementedListServer) CheckAccount(ctx context.Context, req *CheckAccountReq) (*CheckAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccount not implemented")
}

func RegisterListServer(s *grpc.Server, srv ListServer) {
	s.RegisterService(&_List_serviceDesc, srv)
}

func _List_FavoriteTab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteTabReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).FavoriteTab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.polymer.list.List/FavoriteTab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).FavoriteTab(ctx, req.(*FavoriteTabReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_CheckAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).CheckAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.polymer.list.List/CheckAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).CheckAccount(ctx, req.(*CheckAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _List_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.polymer.list.List",
	HandlerType: (*ListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteTab",
			Handler:    _List_FavoriteTab_Handler,
		},
		{
			MethodName: "CheckAccount",
			Handler:    _List_CheckAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_api/bilibili/polymer/list/list.proto",
}