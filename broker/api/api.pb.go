// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type Message struct {
	//metadata
	Meta []byte `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	//message body
	Body                 map[string]string `protobuf:"bytes,2,rep,name=body,proto3" json:"body,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Message) GetBody() map[string]string {
	if m != nil {
		return m.Body
	}
	return nil
}

//RegisterReq register an agent
type RegisterReq struct {
	Id                   int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Servername           string            `protobuf:"bytes,2,opt,name=servername,proto3" json:"servername,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,3,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegisterReq) GetServername() string {
	if m != nil {
		return m.Servername
	}
	return ""
}

func (m *RegisterReq) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

type RegisterResp struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Accepted             bool     `protobuf:"varint,2,opt,name=accepted,proto3" json:"accepted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

func (m *RegisterResp) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegisterResp) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

//UnRegisterReq unregister an agent
type UnRegisterReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnRegisterReq) Reset()         { *m = UnRegisterReq{} }
func (m *UnRegisterReq) String() string { return proto.CompactTextString(m) }
func (*UnRegisterReq) ProtoMessage()    {}
func (*UnRegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *UnRegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnRegisterReq.Unmarshal(m, b)
}
func (m *UnRegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnRegisterReq.Marshal(b, m, deterministic)
}
func (m *UnRegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnRegisterReq.Merge(m, src)
}
func (m *UnRegisterReq) XXX_Size() int {
	return xxx_messageInfo_UnRegisterReq.Size(m)
}
func (m *UnRegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UnRegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_UnRegisterReq proto.InternalMessageInfo

func (m *UnRegisterReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

//UnRegisterResp
type UnRegisterResp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnRegisterResp) Reset()         { *m = UnRegisterResp{} }
func (m *UnRegisterResp) String() string { return proto.CompactTextString(m) }
func (*UnRegisterResp) ProtoMessage()    {}
func (*UnRegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *UnRegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnRegisterResp.Unmarshal(m, b)
}
func (m *UnRegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnRegisterResp.Marshal(b, m, deterministic)
}
func (m *UnRegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnRegisterResp.Merge(m, src)
}
func (m *UnRegisterResp) XXX_Size() int {
	return xxx_messageInfo_UnRegisterResp.Size(m)
}
func (m *UnRegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UnRegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_UnRegisterResp proto.InternalMessageInfo

func (m *UnRegisterResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//HeartbeatReq heartbeat request
type HeartbeatReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatReq) Reset()         { *m = HeartbeatReq{} }
func (m *HeartbeatReq) String() string { return proto.CompactTextString(m) }
func (*HeartbeatReq) ProtoMessage()    {}
func (*HeartbeatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *HeartbeatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatReq.Unmarshal(m, b)
}
func (m *HeartbeatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatReq.Marshal(b, m, deterministic)
}
func (m *HeartbeatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatReq.Merge(m, src)
}
func (m *HeartbeatReq) XXX_Size() int {
	return xxx_messageInfo_HeartbeatReq.Size(m)
}
func (m *HeartbeatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatReq proto.InternalMessageInfo

func (m *HeartbeatReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

//HeartbeatResp heartbeat response
type HeartbeatResp struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatResp) Reset()         { *m = HeartbeatResp{} }
func (m *HeartbeatResp) String() string { return proto.CompactTextString(m) }
func (*HeartbeatResp) ProtoMessage()    {}
func (*HeartbeatResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *HeartbeatResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatResp.Unmarshal(m, b)
}
func (m *HeartbeatResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatResp.Marshal(b, m, deterministic)
}
func (m *HeartbeatResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatResp.Merge(m, src)
}
func (m *HeartbeatResp) XXX_Size() int {
	return xxx_messageInfo_HeartbeatResp.Size(m)
}
func (m *HeartbeatResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatResp.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatResp proto.InternalMessageInfo

func (m *HeartbeatResp) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

//HeartbeatResp commit request
type CommitReq struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitReq) Reset()         { *m = CommitReq{} }
func (m *CommitReq) String() string { return proto.CompactTextString(m) }
func (*CommitReq) ProtoMessage()    {}
func (*CommitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *CommitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitReq.Unmarshal(m, b)
}
func (m *CommitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitReq.Marshal(b, m, deterministic)
}
func (m *CommitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitReq.Merge(m, src)
}
func (m *CommitReq) XXX_Size() int {
	return xxx_messageInfo_CommitReq.Size(m)
}
func (m *CommitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitReq.DiscardUnknown(m)
}

var xxx_messageInfo_CommitReq proto.InternalMessageInfo

func (m *CommitReq) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

//CommitResp commit response
type CommitResp struct {
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitResp) Reset()         { *m = CommitResp{} }
func (m *CommitResp) String() string { return proto.CompactTextString(m) }
func (*CommitResp) ProtoMessage()    {}
func (*CommitResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *CommitResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitResp.Unmarshal(m, b)
}
func (m *CommitResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitResp.Marshal(b, m, deterministic)
}
func (m *CommitResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitResp.Merge(m, src)
}
func (m *CommitResp) XXX_Size() int {
	return xxx_messageInfo_CommitResp.Size(m)
}
func (m *CommitResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitResp.DiscardUnknown(m)
}

var xxx_messageInfo_CommitResp proto.InternalMessageInfo

func (m *CommitResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Message)(nil), "api.Message")
	proto.RegisterMapType((map[string]string)(nil), "api.Message.BodyEntry")
	proto.RegisterType((*RegisterReq)(nil), "api.RegisterReq")
	proto.RegisterMapType((map[string]string)(nil), "api.RegisterReq.MetaEntry")
	proto.RegisterType((*RegisterResp)(nil), "api.RegisterResp")
	proto.RegisterType((*UnRegisterReq)(nil), "api.UnRegisterReq")
	proto.RegisterType((*UnRegisterResp)(nil), "api.UnRegisterResp")
	proto.RegisterType((*HeartbeatReq)(nil), "api.HeartbeatReq")
	proto.RegisterType((*HeartbeatResp)(nil), "api.HeartbeatResp")
	proto.RegisterType((*CommitReq)(nil), "api.CommitReq")
	proto.RegisterType((*CommitResp)(nil), "api.CommitResp")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0xab, 0xd3, 0x40,
	0x10, 0xc7, 0x49, 0xf2, 0xda, 0xd7, 0x4c, 0xf3, 0xea, 0x73, 0x14, 0x09, 0x39, 0xbc, 0x96, 0x1c,
	0x4a, 0xed, 0x21, 0x42, 0x8a, 0x28, 0x3d, 0x56, 0x04, 0x2f, 0xbd, 0x04, 0xbc, 0x78, 0xdb, 0x26,
	0x43, 0x0d, 0x35, 0xc9, 0xba, 0xbb, 0x2d, 0xe4, 0xea, 0x1f, 0xe3, 0xbf, 0xe3, 0xbf, 0x24, 0xd9,
	0xfc, 0x68, 0xda, 0x8a, 0xf0, 0x6e, 0x3b, 0xdf, 0x9d, 0xef, 0xcc, 0x67, 0x66, 0x59, 0xb0, 0x19,
	0x4f, 0x03, 0x2e, 0x0a, 0x55, 0xa0, 0xc5, 0x78, 0xea, 0xff, 0x32, 0xe0, 0x7e, 0x4b, 0x52, 0xb2,
	0x3d, 0x21, 0xc2, 0x5d, 0x46, 0x8a, 0xb9, 0xc6, 0xcc, 0x58, 0x38, 0x91, 0x3e, 0xe3, 0x12, 0xee,
	0x76, 0x45, 0x52, 0xba, 0xe6, 0xcc, 0x5a, 0x8c, 0xc3, 0x37, 0x41, 0x65, 0x6f, 0xf2, 0x83, 0x4d,
	0x91, 0x94, 0x9f, 0x73, 0x25, 0xca, 0x48, 0xe7, 0x78, 0x1f, 0xc0, 0xee, 0x24, 0x7c, 0x04, 0xeb,
	0x40, 0xa5, 0xae, 0x65, 0x47, 0xd5, 0x11, 0x5f, 0xc3, 0xe0, 0xc4, 0x7e, 0x1c, 0xc9, 0x35, 0xb5,
	0x56, 0x07, 0x6b, 0xf3, 0xa3, 0xe1, 0xff, 0x36, 0x60, 0x1c, 0xd1, 0x3e, 0x95, 0x8a, 0x44, 0x44,
	0x3f, 0x71, 0x02, 0x66, 0x9a, 0x68, 0xeb, 0x20, 0x32, 0xd3, 0x04, 0x9f, 0x00, 0x24, 0x89, 0x13,
	0x89, 0x9c, 0x65, 0xad, 0xbd, 0xa7, 0x60, 0xd0, 0x80, 0x5b, 0x1a, 0xd2, 0xd3, 0x90, 0xbd, 0x7a,
	0xc1, 0x96, 0x14, 0x6b, 0x40, 0xab, 0xbc, 0x0a, 0xb4, 0x93, 0x9e, 0x05, 0xba, 0x06, 0xe7, 0x5c,
	0x57, 0xf2, 0x1b, 0x50, 0x0f, 0x46, 0x2c, 0x8e, 0x89, 0x2b, 0x4a, 0xb4, 0x79, 0x14, 0x75, 0xb1,
	0x3f, 0x85, 0x87, 0xaf, 0xf9, 0x7f, 0xa6, 0xf4, 0x97, 0x30, 0xe9, 0x27, 0x48, 0x8e, 0x2e, 0xdc,
	0x67, 0xf5, 0xae, 0x1b, 0xbc, 0x36, 0xf4, 0x9f, 0xc0, 0xf9, 0x42, 0x4c, 0xa8, 0x1d, 0x31, 0xf5,
	0xaf, 0x5a, 0x53, 0x78, 0xe8, 0xdd, 0xdf, 0x92, 0xfa, 0x2b, 0xb0, 0x3f, 0x15, 0x59, 0x96, 0x6a,
	0xf7, 0xfc, 0xb2, 0xcf, 0x38, 0x74, 0xfa, 0xef, 0x7c, 0xee, 0x3a, 0x07, 0x68, 0x4d, 0x35, 0x9d,
	0x3c, 0xc6, 0x31, 0x49, 0xd9, 0xcc, 0xda, 0x86, 0xe1, 0x1f, 0x03, 0x86, 0x1b, 0x51, 0x1c, 0x48,
	0xe0, 0x3b, 0x18, 0x89, 0x66, 0x24, 0x7c, 0xbc, 0x7e, 0x18, 0xef, 0xe5, 0x95, 0x22, 0x39, 0x86,
	0x60, 0x7f, 0x6f, 0xc9, 0xb1, 0xbe, 0xef, 0x4f, 0xea, 0xe1, 0xb5, 0x24, 0x39, 0xbe, 0x07, 0x38,
	0xe6, 0x5d, 0x9b, 0x3a, 0xe3, 0x62, 0xd7, 0xde, 0xab, 0x1b, 0x4d, 0x72, 0x7c, 0x0b, 0xc3, 0x58,
	0x8f, 0x83, 0x13, 0x7d, 0xdd, 0x2d, 0xc4, 0x7b, 0x71, 0x11, 0x4b, 0xbe, 0x19, 0x7c, 0xab, 0x7e,
	0xcb, 0x6e, 0xa8, 0x7f, 0xce, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xcc, 0x7c, 0x12,
	0x46, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrokerClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	//Heartbeat
	//agent should be heartbeat with interval 30s~720s
	Heartbeat(ctx context.Context, in *HeartbeatReq, opts ...grpc.CallOption) (*HeartbeatResp, error)
	//UnRegister agent unregister from broker cluster
	Unregister(ctx context.Context, in *UnRegisterReq, opts ...grpc.CallOption) (*UnRegisterResp, error)
	//commit message into broker
	Commit(ctx context.Context, in *CommitReq, opts ...grpc.CallOption) (*CommitResp, error)
}

type brokerClient struct {
	cc *grpc.ClientConn
}

func NewBrokerClient(cc *grpc.ClientConn) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/api.Broker/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) Heartbeat(ctx context.Context, in *HeartbeatReq, opts ...grpc.CallOption) (*HeartbeatResp, error) {
	out := new(HeartbeatResp)
	err := c.cc.Invoke(ctx, "/api.Broker/heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) Unregister(ctx context.Context, in *UnRegisterReq, opts ...grpc.CallOption) (*UnRegisterResp, error) {
	out := new(UnRegisterResp)
	err := c.cc.Invoke(ctx, "/api.Broker/unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) Commit(ctx context.Context, in *CommitReq, opts ...grpc.CallOption) (*CommitResp, error) {
	out := new(CommitResp)
	err := c.cc.Invoke(ctx, "/api.Broker/commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServer is the server API for Broker service.
type BrokerServer interface {
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	//Heartbeat
	//agent should be heartbeat with interval 30s~720s
	Heartbeat(context.Context, *HeartbeatReq) (*HeartbeatResp, error)
	//UnRegister agent unregister from broker cluster
	Unregister(context.Context, *UnRegisterReq) (*UnRegisterResp, error)
	//commit message into broker
	Commit(context.Context, *CommitReq) (*CommitResp, error)
}

// UnimplementedBrokerServer can be embedded to have forward compatible implementations.
type UnimplementedBrokerServer struct {
}

func (*UnimplementedBrokerServer) Register(ctx context.Context, req *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedBrokerServer) Heartbeat(ctx context.Context, req *HeartbeatReq) (*HeartbeatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (*UnimplementedBrokerServer) Unregister(ctx context.Context, req *UnRegisterReq) (*UnRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (*UnimplementedBrokerServer) Commit(ctx context.Context, req *CommitReq) (*CommitResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}

func RegisterBrokerServer(s *grpc.Server, srv BrokerServer) {
	s.RegisterService(&_Broker_serviceDesc, srv)
}

func _Broker_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Broker/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Broker/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Heartbeat(ctx, req.(*HeartbeatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Broker/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Unregister(ctx, req.(*UnRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Broker/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Commit(ctx, req.(*CommitReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Broker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _Broker_Register_Handler,
		},
		{
			MethodName: "heartbeat",
			Handler:    _Broker_Heartbeat_Handler,
		},
		{
			MethodName: "unregister",
			Handler:    _Broker_Unregister_Handler,
		},
		{
			MethodName: "commit",
			Handler:    _Broker_Commit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
