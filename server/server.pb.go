// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/server.proto

package server

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	entity "pathwar.land/entity"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Void is an empty message
type Void struct {
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde5b5d7aefe7c04, []int{0}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Void.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return m.Size()
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

// AuthenticateInput contains everything (credentials) to authenticate a user and create a new session.
type AuthenticateInput struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *AuthenticateInput) Reset()         { *m = AuthenticateInput{} }
func (m *AuthenticateInput) String() string { return proto.CompactTextString(m) }
func (*AuthenticateInput) ProtoMessage()    {}
func (*AuthenticateInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde5b5d7aefe7c04, []int{1}
}
func (m *AuthenticateInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticateInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticateInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthenticateInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateInput.Merge(m, src)
}
func (m *AuthenticateInput) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticateInput) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateInput.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateInput proto.InternalMessageInfo

func (m *AuthenticateInput) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticateInput) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// AuthenticateOutput contains a session token (JWT).
type AuthenticateOutput struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *AuthenticateOutput) Reset()         { *m = AuthenticateOutput{} }
func (m *AuthenticateOutput) String() string { return proto.CompactTextString(m) }
func (*AuthenticateOutput) ProtoMessage()    {}
func (*AuthenticateOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde5b5d7aefe7c04, []int{2}
}
func (m *AuthenticateOutput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticateOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticateOutput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthenticateOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateOutput.Merge(m, src)
}
func (m *AuthenticateOutput) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticateOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateOutput.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateOutput proto.InternalMessageInfo

func (m *AuthenticateOutput) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Void)(nil), "pathwar.server.Void")
	proto.RegisterType((*AuthenticateInput)(nil), "pathwar.server.AuthenticateInput")
	proto.RegisterType((*AuthenticateOutput)(nil), "pathwar.server.AuthenticateOutput")
}

func init() { proto.RegisterFile("server/server.proto", fileDescriptor_fde5b5d7aefe7c04) }

var fileDescriptor_fde5b5d7aefe7c04 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x63, 0x37, 0x4d, 0x9b, 0xd7, 0xa6, 0xd9, 0x4e, 0x0b, 0x74, 0x03, 0x64, 0x8d, 0x4f,
	0x28, 0xc2, 0x9e, 0x92, 0xad, 0x10, 0x2a, 0x17, 0xb2, 0xb4, 0x44, 0xa1, 0x05, 0xa2, 0x76, 0xf9,
	0xb5, 0x17, 0x34, 0x49, 0x5e, 0x1c, 0xd3, 0x64, 0xc6, 0x78, 0xc6, 0x2d, 0x15, 0xe2, 0xc2, 0x81,
	0x33, 0x5a, 0x6e, 0x5c, 0xf8, 0x03, 0xf8, 0x47, 0x38, 0xae, 0xc4, 0x85, 0x23, 0x6a, 0xf9, 0x2b,
	0x38, 0xa1, 0x19, 0x3b, 0xbb, 0xde, 0x64, 0xc3, 0x76, 0x4f, 0xe3, 0xf7, 0xe6, 0xfb, 0x3e, 0xf3,
	0xd5, 0x9b, 0x67, 0x1b, 0xb6, 0x24, 0xc6, 0xe7, 0x18, 0xd3, 0x74, 0xf1, 0xa3, 0x58, 0x28, 0x41,
	0x36, 0x22, 0xa6, 0x46, 0x17, 0x2c, 0xf6, 0xd3, 0x6c, 0xed, 0xb5, 0x40, 0x88, 0x60, 0x8c, 0x94,
	0x45, 0x21, 0x65, 0x9c, 0x0b, 0xc5, 0x54, 0x28, 0xb8, 0x4c, 0xd5, 0xb5, 0x2d, 0xe4, 0x2a, 0x54,
	0x97, 0x34, 0x5d, 0xb2, 0xa4, 0x17, 0x84, 0x6a, 0x94, 0xf4, 0xfc, 0xbe, 0x98, 0xd0, 0x40, 0x04,
	0x82, 0x9a, 0x74, 0x2f, 0x19, 0x9a, 0xc8, 0x04, 0xe6, 0x29, 0x93, 0xbf, 0x65, 0x96, 0xbe, 0x17,
	0x20, 0xf7, 0xe4, 0x05, 0x0b, 0x02, 0x8c, 0xa9, 0x88, 0xcc, 0x29, 0xf3, 0x27, 0xba, 0x9b, 0x50,
	0xfc, 0x5c, 0x84, 0x83, 0xfd, 0xf2, 0xc3, 0x56, 0xa9, 0x59, 0x24, 0xf6, 0xf7, 0x3f, 0xb8, 0x47,
	0xb0, 0xd9, 0x4a, 0xd4, 0x48, 0x7b, 0xe8, 0x33, 0x85, 0x1d, 0x1e, 0x25, 0x8a, 0xd4, 0x60, 0x35,
	0x91, 0x18, 0x73, 0x36, 0xc1, 0x1d, 0xcb, 0xb1, 0xde, 0x2c, 0x9f, 0x3c, 0x8e, 0xf5, 0x5e, 0xc4,
	0xa4, 0xbc, 0x10, 0xf1, 0x60, 0xc7, 0x4e, 0xf7, 0xa6, 0xb1, 0xdb, 0x00, 0x92, 0x87, 0x7d, 0x9a,
	0x28, 0x4d, 0xdb, 0x86, 0x65, 0x25, 0xce, 0x90, 0x67, 0xa8, 0x34, 0x68, 0xfe, 0xb4, 0x02, 0xa5,
	0x53, 0xd3, 0x26, 0x22, 0x60, 0x3d, 0x5f, 0x46, 0xde, 0xf0, 0x9f, 0xee, 0xa3, 0x3f, 0xe7, 0xb0,
	0xe6, 0xfe, 0x9f, 0x24, 0x3d, 0xd7, 0xdd, 0xf9, 0xf1, 0xcf, 0x7f, 0x7e, 0xb1, 0x89, 0x5b, 0xa1,
	0x2c, 0xb7, 0xb9, 0x6f, 0x35, 0x48, 0x0b, 0x8a, 0xdd, 0x90, 0x07, 0x64, 0x7b, 0x96, 0xa2, 0xbb,
	0x53, 0x7b, 0x66, 0xd6, 0xad, 0x18, 0xda, 0x0a, 0x59, 0xa6, 0x91, 0x2e, 0xfd, 0x02, 0xd6, 0x3e,
	0x93, 0x18, 0x9f, 0xa2, 0x94, 0xa1, 0xe0, 0x0b, 0x48, 0xaf, 0x3e, 0xce, 0x66, 0x77, 0x9c, 0x2b,
	0x71, 0x5f, 0x32, 0xc0, 0x2a, 0xa9, 0x50, 0xdd, 0x5b, 0x4f, 0x66, 0xa4, 0xaf, 0x60, 0xed, 0xbe,
	0x48, 0x4c, 0xaf, 0xb9, 0x92, 0x0b, 0xc0, 0xf5, 0x59, 0xf0, 0x93, 0x92, 0xe3, 0x50, 0x2a, 0x77,
	0xdb, 0xb0, 0x37, 0xc8, 0x3a, 0x55, 0x39, 0x56, 0x1b, 0x96, 0xb5, 0x81, 0x45, 0xd0, 0x9d, 0x67,
	0xb9, 0x35, 0xb8, 0x0d, 0x83, 0x5b, 0x25, 0x25, 0x63, 0x55, 0x92, 0x8f, 0xa0, 0x74, 0x8c, 0xe7,
	0x38, 0x5e, 0x44, 0xba, 0x3d, 0x4b, 0x32, 0x6a, 0x83, 0xaa, 0x1a, 0x54, 0x99, 0xac, 0xd0, 0x71,
	0x4a, 0x68, 0xc3, 0xf2, 0x7d, 0x64, 0x93, 0x1b, 0x9b, 0xd2, 0xe2, 0x19, 0x53, 0xca, 0xd4, 0xb7,
	0xa1, 0x74, 0xaa, 0x98, 0x4a, 0x16, 0x91, 0x5e, 0x9e, 0x25, 0xa5, 0xea, 0x9c, 0x23, 0x99, 0x96,
	0xb7, 0xa0, 0xd8, 0xe1, 0x43, 0xf1, 0xdc, 0xe9, 0xc8, 0x30, 0x5a, 0x9b, 0x9b, 0x8e, 0x50, 0x97,
	0x7e, 0x0d, 0xb7, 0xda, 0xc8, 0x31, 0x66, 0x0a, 0x3f, 0x64, 0x67, 0x78, 0xc0, 0x14, 0x7b, 0xa1,
	0x61, 0xbb, 0x63, 0x70, 0xb7, 0xdd, 0x57, 0xe8, 0x00, 0xcf, 0x69, 0x90, 0xa1, 0xbc, 0x21, 0x3b,
	0x43, 0x6f, 0xa0, 0x61, 0x1d, 0x28, 0x1e, 0x24, 0x93, 0xe8, 0xa6, 0x1e, 0xb5, 0x76, 0x3a, 0x70,
	0x6e, 0xc5, 0x40, 0xe5, 0xb7, 0x63, 0x6f, 0x90, 0x4c, 0xa2, 0x7b, 0xbf, 0x17, 0x1f, 0xb6, 0xfe,
	0x5d, 0x22, 0xbf, 0x59, 0xb0, 0xd6, 0x4d, 0xab, 0x9c, 0x56, 0xb7, 0xe3, 0x1e, 0xc2, 0xfa, 0x34,
	0xd4, 0x2d, 0x27, 0xee, 0x48, 0xa9, 0x48, 0xee, 0x53, 0x9a, 0xfb, 0x40, 0x65, 0xa7, 0x4c, 0xd7,
	0xda, 0xa6, 0x54, 0x6c, 0x38, 0x7c, 0x7f, 0x7a, 0x38, 0x47, 0xd5, 0x38, 0x00, 0x68, 0x45, 0xac,
	0x3f, 0x42, 0xaf, 0xe9, 0xef, 0x92, 0x77, 0x9e, 0x0f, 0xa1, 0xbd, 0xb1, 0xe8, 0xd1, 0x09, 0x93,
	0x0a, 0x63, 0x7a, 0xdc, 0xf9, 0xe0, 0xf0, 0x93, 0xd3, 0xc3, 0xe6, 0xd2, 0xdb, 0xfe, 0x6e, 0xad,
	0xca, 0xa2, 0xd0, 0xcf, 0xb1, 0x5d, 0x8b, 0x36, 0x6c, 0xbb, 0xd8, 0xbc, 0xc5, 0xa2, 0x68, 0xac,
	0xdf, 0xed, 0x50, 0x70, 0xfa, 0x8d, 0x14, 0x7c, 0x7f, 0x2e, 0x73, 0xd2, 0x85, 0xa5, 0xbd, 0xdd,
	0xbb, 0xa4, 0x03, 0xed, 0x13, 0x54, 0x49, 0xcc, 0x71, 0xe0, 0x5c, 0x8c, 0x90, 0x3b, 0x6a, 0x84,
	0x8e, 0x9e, 0x69, 0x67, 0x20, 0x50, 0x3a, 0x5c, 0x28, 0x67, 0xc4, 0xce, 0xd1, 0x89, 0x30, 0x9e,
	0x84, 0xe6, 0x7d, 0x74, 0x94, 0x70, 0x58, 0xbf, 0x8f, 0x52, 0x1a, 0x6d, 0x8c, 0x52, 0x24, 0x71,
	0x1f, 0xfd, 0x93, 0xf7, 0x34, 0x71, 0x8f, 0xec, 0x41, 0x63, 0x9e, 0x38, 0x55, 0x3d, 0xa1, 0xe2,
	0x77, 0xa1, 0x54, 0x3e, 0x29, 0x41, 0xf1, 0x57, 0xdb, 0x5a, 0x79, 0xb0, 0x0b, 0x55, 0x28, 0xdf,
	0x63, 0x32, 0xec, 0xeb, 0xef, 0x16, 0xb1, 0x57, 0x2d, 0x78, 0x5d, 0x37, 0x2a, 0x3c, 0xc2, 0x4b,
	0x93, 0xa9, 0xae, 0xda, 0xb5, 0xf2, 0x97, 0x5e, 0xab, 0xdb, 0xf1, 0x8e, 0xf0, 0xd2, 0xb1, 0x7b,
	0x77, 0xa0, 0x92, 0xaf, 0x28, 0xc0, 0xc6, 0x53, 0xfa, 0x42, 0xfc, 0x2e, 0x90, 0x8f, 0x45, 0x8c,
	0x0e, 0xeb, 0x89, 0x44, 0x39, 0xd9, 0xd5, 0xdd, 0xe4, 0xd6, 0xfe, 0xb8, 0xaa, 0x5b, 0x8f, 0xae,
	0xea, 0xd6, 0xdf, 0x57, 0x75, 0xeb, 0xe7, 0xeb, 0x7a, 0xe1, 0xd1, 0x75, 0xbd, 0xf0, 0xd7, 0x75,
	0xbd, 0xf0, 0x60, 0x6b, 0xda, 0xeb, 0x31, 0xe3, 0x83, 0xec, 0x47, 0xd7, 0x2b, 0x99, 0x3f, 0xc9,
	0xdd, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x23, 0x2e, 0x73, 0x00, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	// Create a session based on credentials
	Authenticate(ctx context.Context, in *AuthenticateInput, opts ...grpc.CallOption) (*AuthenticateOutput, error)
	// Verify the service is up and running
	Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	// Get current user's session
	UserSession(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.UserSession, error)
	// List tournaments
	Tournaments(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.TournamentList, error)
	// List users
	Users(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.UserList, error)
	// List levels
	Levels(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.LevelList, error)
	// List teams
	Teams(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.TeamList, error)
	Status(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.Status, error)
	Info(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.Info, error)
	GenerateFakeData(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	Dump(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.Dump, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Authenticate(ctx context.Context, in *AuthenticateInput, opts ...grpc.CallOption) (*AuthenticateOutput, error) {
	out := new(AuthenticateOutput)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) UserSession(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.UserSession, error) {
	out := new(entity.UserSession)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/UserSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Tournaments(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.TournamentList, error) {
	out := new(entity.TournamentList)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Tournaments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Users(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.UserList, error) {
	out := new(entity.UserList)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Users", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Levels(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.LevelList, error) {
	out := new(entity.LevelList)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Levels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Teams(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.TeamList, error) {
	out := new(entity.TeamList)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Teams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Status(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.Status, error) {
	out := new(entity.Status)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Info(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.Info, error) {
	out := new(entity.Info)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) GenerateFakeData(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/GenerateFakeData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Dump(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.Dump, error) {
	out := new(entity.Dump)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Dump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	// Create a session based on credentials
	Authenticate(context.Context, *AuthenticateInput) (*AuthenticateOutput, error)
	// Verify the service is up and running
	Ping(context.Context, *Void) (*Void, error)
	// Get current user's session
	UserSession(context.Context, *Void) (*entity.UserSession, error)
	// List tournaments
	Tournaments(context.Context, *Void) (*entity.TournamentList, error)
	// List users
	Users(context.Context, *Void) (*entity.UserList, error)
	// List levels
	Levels(context.Context, *Void) (*entity.LevelList, error)
	// List teams
	Teams(context.Context, *Void) (*entity.TeamList, error)
	Status(context.Context, *Void) (*entity.Status, error)
	Info(context.Context, *Void) (*entity.Info, error)
	GenerateFakeData(context.Context, *Void) (*Void, error)
	Dump(context.Context, *Void) (*entity.Dump, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Authenticate(ctx, req.(*AuthenticateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Ping(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_UserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).UserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/UserSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).UserSession(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Tournaments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Tournaments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Tournaments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Tournaments(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Users_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Users(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Users",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Users(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Levels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Levels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Levels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Levels(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Teams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Teams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Teams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Teams(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Status(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Info(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_GenerateFakeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GenerateFakeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/GenerateFakeData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GenerateFakeData(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Dump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Dump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Dump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Dump(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pathwar.server.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Server_Authenticate_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Server_Ping_Handler,
		},
		{
			MethodName: "UserSession",
			Handler:    _Server_UserSession_Handler,
		},
		{
			MethodName: "Tournaments",
			Handler:    _Server_Tournaments_Handler,
		},
		{
			MethodName: "Users",
			Handler:    _Server_Users_Handler,
		},
		{
			MethodName: "Levels",
			Handler:    _Server_Levels_Handler,
		},
		{
			MethodName: "Teams",
			Handler:    _Server_Teams_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Server_Status_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Server_Info_Handler,
		},
		{
			MethodName: "GenerateFakeData",
			Handler:    _Server_GenerateFakeData_Handler,
		},
		{
			MethodName: "Dump",
			Handler:    _Server_Dump_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/server.proto",
}

func (m *Void) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Void) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *AuthenticateInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateInput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintServer(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintServer(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	return i, nil
}

func (m *AuthenticateOutput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateOutput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintServer(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func encodeVarintServer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Void) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *AuthenticateInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	return n
}

func (m *AuthenticateOutput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	return n
}

func sovServer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozServer(x uint64) (n int) {
	return sovServer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Void) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Void: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Void: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthenticateInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthenticateInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthenticateOutput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthenticateOutput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateOutput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipServer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthServer
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthServer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowServer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipServer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthServer
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthServer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServer   = fmt.Errorf("proto: integer overflow")
)
