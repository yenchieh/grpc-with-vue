// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vote.proto

package vote

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Topic struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	VoteCount            uint64   `protobuf:"varint,3,opt,name=voteCount,proto3" json:"voteCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_21d31c94b62a6ac7, []int{0}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Topic) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Topic) GetVoteCount() uint64 {
	if m != nil {
		return m.VoteCount
	}
	return 0
}

type Topics struct {
	Topics               []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topics) Reset()         { *m = Topics{} }
func (m *Topics) String() string { return proto.CompactTextString(m) }
func (*Topics) ProtoMessage()    {}
func (*Topics) Descriptor() ([]byte, []int) {
	return fileDescriptor_21d31c94b62a6ac7, []int{1}
}

func (m *Topics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topics.Unmarshal(m, b)
}
func (m *Topics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topics.Marshal(b, m, deterministic)
}
func (m *Topics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topics.Merge(m, src)
}
func (m *Topics) XXX_Size() int {
	return xxx_messageInfo_Topics.Size(m)
}
func (m *Topics) XXX_DiscardUnknown() {
	xxx_messageInfo_Topics.DiscardUnknown(m)
}

var xxx_messageInfo_Topics proto.InternalMessageInfo

func (m *Topics) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type RequestByID struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestByID) Reset()         { *m = RequestByID{} }
func (m *RequestByID) String() string { return proto.CompactTextString(m) }
func (*RequestByID) ProtoMessage()    {}
func (*RequestByID) Descriptor() ([]byte, []int) {
	return fileDescriptor_21d31c94b62a6ac7, []int{2}
}

func (m *RequestByID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestByID.Unmarshal(m, b)
}
func (m *RequestByID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestByID.Marshal(b, m, deterministic)
}
func (m *RequestByID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestByID.Merge(m, src)
}
func (m *RequestByID) XXX_Size() int {
	return xxx_messageInfo_RequestByID.Size(m)
}
func (m *RequestByID) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestByID.DiscardUnknown(m)
}

var xxx_messageInfo_RequestByID proto.InternalMessageInfo

func (m *RequestByID) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Topic)(nil), "vote.Topic")
	proto.RegisterType((*Topics)(nil), "vote.Topics")
	proto.RegisterType((*RequestByID)(nil), "vote.RequestByID")
}

func init() {
	proto.RegisterFile("vote.proto", fileDescriptor_21d31c94b62a6ac7)
}

var fileDescriptor_21d31c94b62a6ac7 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x9b, 0xed, 0xba, 0x92, 0x59, 0x11, 0xcc, 0x41, 0xc2, 0xaa, 0xb0, 0xc4, 0xcb, 0x5e,
	0x9a, 0x42, 0xfb, 0x0d, 0xfc, 0x83, 0x78, 0x12, 0x82, 0x78, 0xb7, 0x76, 0x5c, 0x56, 0xac, 0xb3,
	0x9a, 0x59, 0xa1, 0xdf, 0x5e, 0x92, 0xb4, 0xb8, 0xf4, 0x36, 0x99, 0x97, 0x79, 0xef, 0xfd, 0x00,
	0x7e, 0x89, 0xd1, 0xf6, 0x3f, 0xc4, 0xa4, 0xf2, 0x30, 0x57, 0x17, 0x2d, 0x51, 0xfb, 0x89, 0xf3,
	0xb8, 0x5b, 0x0d, 0xef, 0x73, 0xdc, 0xf4, 0xbc, 0x4d, 0x5f, 0xcc, 0x13, 0x1c, 0x3d, 0x53, 0xdf,
	0xbd, 0xa9, 0x53, 0xc8, 0xba, 0xb5, 0x16, 0xb5, 0x68, 0x72, 0x97, 0x75, 0x6b, 0xa5, 0xe1, 0x78,
	0x83, 0xde, 0xbf, 0xb6, 0xa8, 0xb3, 0x5a, 0x34, 0xd2, 0xed, 0x9f, 0xea, 0x12, 0x64, 0xf0, 0xbd,
	0xa5, 0xe1, 0x8b, 0xf5, 0x34, 0x1e, 0xfc, 0x2f, 0xcc, 0x0c, 0x8a, 0x68, 0xe8, 0xd5, 0x35, 0x14,
	0x1c, 0x27, 0x2d, 0xea, 0x69, 0x53, 0x2e, 0x4a, 0x1b, 0xab, 0x45, 0xd5, 0xed, 0x24, 0x73, 0x05,
	0xa5, 0xc3, 0xef, 0x01, 0x3d, 0xdf, 0x6c, 0x1f, 0xef, 0x0e, 0x5b, 0x2c, 0x3e, 0x20, 0x7f, 0x21,
	0x46, 0xb5, 0x04, 0xf9, 0x80, 0xbc, 0x33, 0x3e, 0xb7, 0x89, 0xc8, 0xee, 0x89, 0xec, 0x7d, 0x20,
	0xaa, 0x4e, 0x46, 0x01, 0xde, 0x4c, 0xd4, 0x0c, 0x64, 0x38, 0x4e, 0x7c, 0x67, 0x49, 0x1c, 0x85,
	0x55, 0xe3, 0x42, 0x66, 0xb2, 0x2a, 0xa2, 0xdd, 0xf2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x50, 0x32,
	0xbd, 0x15, 0x42, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VoteClient is the client API for Vote service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VoteClient interface {
	GetTopics(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Topics, error)
	VoteTopic(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*Topic, error)
}

type voteClient struct {
	cc grpc.ClientConnInterface
}

func NewVoteClient(cc grpc.ClientConnInterface) VoteClient {
	return &voteClient{cc}
}

func (c *voteClient) GetTopics(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Topics, error) {
	out := new(Topics)
	err := c.cc.Invoke(ctx, "/vote.Vote/GetTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteClient) VoteTopic(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/vote.Vote/VoteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoteServer is the server API for Vote service.
type VoteServer interface {
	GetTopics(context.Context, *empty.Empty) (*Topics, error)
	VoteTopic(context.Context, *RequestByID) (*Topic, error)
}

// UnimplementedVoteServer can be embedded to have forward compatible implementations.
type UnimplementedVoteServer struct {
}

func (*UnimplementedVoteServer) GetTopics(ctx context.Context, req *empty.Empty) (*Topics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopics not implemented")
}
func (*UnimplementedVoteServer) VoteTopic(ctx context.Context, req *RequestByID) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteTopic not implemented")
}

func RegisterVoteServer(s *grpc.Server, srv VoteServer) {
	s.RegisterService(&_Vote_serviceDesc, srv)
}

func _Vote_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vote.Vote/GetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServer).GetTopics(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vote_VoteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServer).VoteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vote.Vote/VoteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServer).VoteTopic(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Vote_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vote.Vote",
	HandlerType: (*VoteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopics",
			Handler:    _Vote_GetTopics_Handler,
		},
		{
			MethodName: "VoteTopic",
			Handler:    _Vote_VoteTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vote.proto",
}