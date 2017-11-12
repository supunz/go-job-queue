// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mail.proto

/*
Package mail is a generated protocol buffer package.

It is generated from these files:
	mail.proto

It has these top-level messages:
	EmailRequest
	EmailResponse
*/
package mail

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// EmailRequest - email request definition
type EmailRequest struct {
	Sender   string `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Receiver string `protobuf:"bytes,2,opt,name=receiver" json:"receiver,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
}

func (m *EmailRequest) Reset()                    { *m = EmailRequest{} }
func (m *EmailRequest) String() string            { return proto.CompactTextString(m) }
func (*EmailRequest) ProtoMessage()               {}
func (*EmailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EmailRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EmailRequest) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *EmailRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EmailRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// EmailResponse - email response definition
type EmailResponse struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *EmailResponse) Reset()                    { *m = EmailResponse{} }
func (m *EmailResponse) String() string            { return proto.CompactTextString(m) }
func (*EmailResponse) ProtoMessage()               {}
func (*EmailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EmailResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*EmailRequest)(nil), "mail.EmailRequest")
	proto.RegisterType((*EmailResponse)(nil), "mail.EmailResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Mail service

type MailClient interface {
	// put an email request an email
	PutEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
}

type mailClient struct {
	cc *grpc.ClientConn
}

func NewMailClient(cc *grpc.ClientConn) MailClient {
	return &mailClient{cc}
}

func (c *mailClient) PutEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := grpc.Invoke(ctx, "/mail.Mail/PutEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mail service

type MailServer interface {
	// put an email request an email
	PutEmail(context.Context, *EmailRequest) (*EmailResponse, error)
}

func RegisterMailServer(s *grpc.Server, srv MailServer) {
	s.RegisterService(&_Mail_serviceDesc, srv)
}

func _Mail_PutEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).PutEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mail.Mail/PutEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).PutEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mail_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mail.Mail",
	HandlerType: (*MailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutEmail",
			Handler:    _Mail_PutEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mail.proto",
}

func init() { proto.RegisterFile("mail.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0xce, 0x82, 0x40,
	0x10, 0x84, 0x7f, 0x7e, 0x11, 0x71, 0xa3, 0xcd, 0x6a, 0xcc, 0x85, 0xca, 0xd0, 0x68, 0x45, 0xa1,
	0xb1, 0xb4, 0xb4, 0x34, 0x31, 0xbc, 0x01, 0xe2, 0x16, 0x97, 0xe0, 0x1d, 0xde, 0x2d, 0x3e, 0xbf,
	0x61, 0x39, 0x8d, 0x76, 0xf3, 0xcd, 0xe4, 0x6e, 0x76, 0x00, 0xee, 0x95, 0x6e, 0x8a, 0xd6, 0x59,
	0xb6, 0x18, 0xf7, 0x3a, 0x77, 0x30, 0x3b, 0xf5, 0xa2, 0xa4, 0x47, 0x47, 0x9e, 0x71, 0x05, 0x89,
	0x27, 0x73, 0x23, 0xa7, 0xa2, 0x75, 0xb4, 0x9d, 0x96, 0x81, 0x30, 0x83, 0xd4, 0x51, 0x4d, 0xfa,
	0x49, 0x4e, 0xfd, 0x4b, 0xf2, 0x61, 0x5c, 0xc2, 0x98, 0x35, 0x37, 0xa4, 0x46, 0x12, 0x0c, 0x80,
	0x0a, 0x26, 0xb5, 0x35, 0x4c, 0x86, 0x55, 0x2c, 0xfe, 0x1b, 0xf3, 0x0d, 0xcc, 0x43, 0xa7, 0x6f,
	0xad, 0xf1, 0x24, 0xa5, 0x5c, 0x71, 0xe7, 0xa5, 0x34, 0x2d, 0x03, 0xed, 0x8e, 0x10, 0x9f, 0x2b,
	0xdd, 0xe0, 0x01, 0xd2, 0x4b, 0xc7, 0xf2, 0x06, 0xb1, 0x90, 0x0d, 0xdf, 0x47, 0x67, 0x8b, 0x1f,
	0x6f, 0xf8, 0x34, 0xff, 0xbb, 0x26, 0x32, 0x74, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x59, 0xf9,
	0xa9, 0x60, 0xf6, 0x00, 0x00, 0x00,
}
