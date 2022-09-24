// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/metadata/metadata.proto

package bilibili_metadata

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 请求元数据
// gRPC头部:x-bili-metadata-bin
type Metadata struct {
	// 登录Token
	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// 包类型
	MobiApp string `protobuf:"bytes,2,opt,name=mobi_app,json=mobiApp,proto3" json:"mobi_app,omitempty"`
	// 运行设备
	Device string `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	// 构建id
	Build int32 `protobuf:"varint,4,opt,name=build,proto3" json:"build,omitempty"`
	// 渠道
	Channel string `protobuf:"bytes,5,opt,name=channel,proto3" json:"channel,omitempty"`
	// 设备buvid
	Buvid string `protobuf:"bytes,6,opt,name=buvid,proto3" json:"buvid,omitempty"`
	// 平台类型
	Platform             string   `protobuf:"bytes,7,opt,name=platform,proto3" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_29587c0a3a1bd0c0, []int{0}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *Metadata) GetMobiApp() string {
	if m != nil {
		return m.MobiApp
	}
	return ""
}

func (m *Metadata) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *Metadata) GetBuild() int32 {
	if m != nil {
		return m.Build
	}
	return 0
}

func (m *Metadata) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Metadata) GetBuvid() string {
	if m != nil {
		return m.Buvid
	}
	return ""
}

func (m *Metadata) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadata)(nil), "bilibili.metadata.Metadata")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/metadata/metadata.proto", fileDescriptor_29587c0a3a1bd0c0)
}

var fileDescriptor_29587c0a3a1bd0c0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xcb, 0x4a, 0xc5, 0x30,
	0x10, 0x86, 0x89, 0xda, 0xcb, 0x99, 0x9d, 0x41, 0x64, 0x14, 0x84, 0x83, 0xab, 0xe3, 0xe6, 0x9c,
	0x85, 0x4f, 0xe0, 0x5a, 0xdc, 0xf4, 0x05, 0xca, 0xe4, 0xa2, 0x06, 0xd3, 0x66, 0x68, 0x63, 0xa1,
	0x2f, 0xe7, 0xb3, 0x49, 0x93, 0xb6, 0x8b, 0x81, 0xf9, 0xfe, 0xef, 0x67, 0x60, 0xe0, 0xe5, 0x6b,
	0x60, 0xdd, 0x12, 0xbb, 0x8b, 0x72, 0xde, 0x2d, 0x73, 0xe9, 0x6c, 0x24, 0x43, 0x91, 0xf6, 0xe5,
	0xcc, 0x43, 0x88, 0x41, 0xde, 0x6e, 0x8d, 0xf3, 0x26, 0x9e, 0xff, 0x04, 0xd4, 0x1f, 0x2b, 0xc8,
	0x27, 0x00, 0xd2, 0xda, 0x8e, 0x63, 0xfb, 0x63, 0x67, 0x14, 0x47, 0x71, 0x3a, 0x34, 0x87, 0x9c,
	0xbc, 0xdb, 0x59, 0x3e, 0x40, 0xdd, 0x05, 0xe5, 0x5a, 0x62, 0xc6, 0xab, 0x24, 0xab, 0x85, 0xdf,
	0x98, 0xe5, 0x3d, 0x94, 0xc6, 0x4e, 0x4e, 0x5b, 0xbc, 0x4e, 0x62, 0x25, 0x79, 0x07, 0x85, 0xfa,
	0x75, 0xde, 0xe0, 0xcd, 0x51, 0x9c, 0x8a, 0x26, 0x83, 0x44, 0xa8, 0xf4, 0x37, 0xf5, 0xbd, 0xf5,
	0x58, 0xe4, 0x3b, 0x2b, 0xe6, 0xfe, 0xe4, 0x0c, 0x96, 0x29, 0xcf, 0x20, 0x1f, 0xa1, 0x66, 0x4f,
	0xf1, 0x33, 0x0c, 0x1d, 0x56, 0x49, 0xec, 0xac, 0xca, 0xf4, 0xda, 0xeb, 0x7f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0x80, 0x77, 0x8f, 0x07, 0x01, 0x00, 0x00,
}