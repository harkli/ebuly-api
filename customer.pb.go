// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: customer.proto

package customer

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/metaverse/truss/deftree/googlethirdparty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type RegisteRequest struct {
	Phone  string `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Source uint32 `protobuf:"fixed32,2,opt,name=Source,proto3" json:"Source,omitempty"`
}

func (m *RegisteRequest) Reset()         { *m = RegisteRequest{} }
func (m *RegisteRequest) String() string { return proto.CompactTextString(m) }
func (*RegisteRequest) ProtoMessage()    {}
func (*RegisteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}
func (m *RegisteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteRequest.Merge(m, src)
}
func (m *RegisteRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegisteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteRequest proto.InternalMessageInfo

func (m *RegisteRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegisteRequest) GetSource() uint32 {
	if m != nil {
		return m.Source
	}
	return 0
}

type RegisteResponse struct {
	Code    string       `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    *RegisteData `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *RegisteResponse) Reset()         { *m = RegisteResponse{} }
func (m *RegisteResponse) String() string { return proto.CompactTextString(m) }
func (*RegisteResponse) ProtoMessage()    {}
func (*RegisteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}
func (m *RegisteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteResponse.Merge(m, src)
}
func (m *RegisteResponse) XXX_Size() int {
	return m.Size()
}
func (m *RegisteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteResponse proto.InternalMessageInfo

func (m *RegisteResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RegisteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RegisteResponse) GetData() *RegisteData {
	if m != nil {
		return m.Data
	}
	return nil
}

type RegisteData struct {
	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (m *RegisteData) Reset()         { *m = RegisteData{} }
func (m *RegisteData) String() string { return proto.CompactTextString(m) }
func (*RegisteData) ProtoMessage()    {}
func (*RegisteData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
}
func (m *RegisteData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteData.Merge(m, src)
}
func (m *RegisteData) XXX_Size() int {
	return m.Size()
}
func (m *RegisteData) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteData.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteData proto.InternalMessageInfo

func (m *RegisteData) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type LoginRequest struct {
	Phone    string `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{3}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return m.Size()
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type LoginResponse struct {
	Code    string     `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    *LoginData `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{4}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return m.Size()
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *LoginResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LoginResponse) GetData() *LoginData {
	if m != nil {
		return m.Data
	}
	return nil
}

type LoginData struct {
	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (m *LoginData) Reset()         { *m = LoginData{} }
func (m *LoginData) String() string { return proto.CompactTextString(m) }
func (*LoginData) ProtoMessage()    {}
func (*LoginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{5}
}
func (m *LoginData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginData.Merge(m, src)
}
func (m *LoginData) XXX_Size() int {
	return m.Size()
}
func (m *LoginData) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginData.DiscardUnknown(m)
}

var xxx_messageInfo_LoginData proto.InternalMessageInfo

func (m *LoginData) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisteRequest)(nil), "customer.RegisteRequest")
	proto.RegisterType((*RegisteResponse)(nil), "customer.RegisteResponse")
	proto.RegisterType((*RegisteData)(nil), "customer.RegisteData")
	proto.RegisterType((*LoginRequest)(nil), "customer.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "customer.LoginResponse")
	proto.RegisterType((*LoginData)(nil), "customer.LoginData")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x8e, 0x12, 0x41,
	0x10, 0xc7, 0x69, 0xbe, 0x29, 0x14, 0x4d, 0x23, 0x38, 0x60, 0x32, 0xe2, 0x5c, 0x44, 0x0e, 0x4c,
	0x82, 0x37, 0x0e, 0x1e, 0x44, 0x6f, 0x9a, 0x90, 0xc1, 0x9b, 0xa7, 0x86, 0x29, 0x86, 0x31, 0x30,
	0x8d, 0xdd, 0x3d, 0x1a, 0xaf, 0x3e, 0x81, 0x89, 0xaf, 0xb3, 0x0f, 0xb0, 0x47, 0x92, 0xbd, 0xec,
	0x71, 0x03, 0xfb, 0x20, 0x1b, 0x66, 0x7a, 0x66, 0xd9, 0x85, 0xec, 0x61, 0x6f, 0xfd, 0xaf, 0x8f,
	0x5f, 0xd7, 0xbf, 0x52, 0x50, 0x9b, 0x85, 0x52, 0xf1, 0x15, 0x8a, 0xfe, 0x5a, 0x70, 0xc5, 0x69,
	0x39, 0xd1, 0xed, 0xcf, 0x9e, 0xaf, 0x16, 0xe1, 0xb4, 0x3f, 0xe3, 0x2b, 0x7b, 0x85, 0x8a, 0xfd,
	0x42, 0x21, 0xd1, 0x56, 0x22, 0x94, 0xd2, 0x76, 0x71, 0xae, 0x04, 0xa2, 0xed, 0x71, 0xee, 0x2d,
	0x51, 0x2d, 0x7c, 0xe1, 0xae, 0x99, 0x50, 0x7f, 0x6c, 0x16, 0x04, 0x5c, 0x31, 0xe5, 0xf3, 0x40,
	0xc6, 0x40, 0xeb, 0x03, 0xd4, 0x1c, 0xf4, 0x7c, 0xa9, 0xd0, 0xc1, 0x9f, 0x21, 0x4a, 0x45, 0x5f,
	0x40, 0x61, 0xbc, 0xe0, 0x01, 0x1a, 0xa4, 0x43, 0xba, 0x15, 0x27, 0x16, 0xb4, 0x09, 0xc5, 0x09,
	0x0f, 0xc5, 0x0c, 0x8d, 0x6c, 0x87, 0x74, 0x4b, 0x8e, 0x56, 0xd6, 0x0f, 0x78, 0x96, 0xf6, 0xcb,
	0x35, 0x0f, 0x24, 0x52, 0x0a, 0xf9, 0x11, 0x77, 0x93, 0xfe, 0xe8, 0x4d, 0x0d, 0x28, 0x7d, 0x45,
	0x29, 0x99, 0x17, 0xf7, 0x57, 0x9c, 0x44, 0xd2, 0x77, 0x90, 0xff, 0xc4, 0x14, 0x33, 0x72, 0x1d,
	0xd2, 0xad, 0x0e, 0x1a, 0xfd, 0xd4, 0xb0, 0xc6, 0xee, 0x93, 0x4e, 0x54, 0x62, 0xbd, 0x81, 0xea,
	0x41, 0xf0, 0xd4, 0x3f, 0xd6, 0x37, 0x78, 0xf2, 0x85, 0x7b, 0x7e, 0xf0, 0xb0, 0x99, 0x36, 0x94,
	0xc7, 0x4c, 0xca, 0xdf, 0x5c, 0xb8, 0x7a, 0x9c, 0x54, 0xa7, 0xd4, 0xdc, 0x01, 0x75, 0x0e, 0x4f,
	0x35, 0xf5, 0x51, 0x16, 0xdf, 0xde, 0xb1, 0x58, 0xbf, 0xb5, 0x18, 0x41, 0x0f, 0x0c, 0xbe, 0x86,
	0x4a, 0x1a, 0x3a, 0xf5, 0xc7, 0xe0, 0x8c, 0x40, 0x79, 0xa4, 0xbb, 0xe9, 0x77, 0x28, 0xe9, 0x75,
	0x50, 0xe3, 0x68, 0x6d, 0x7a, 0x01, 0xed, 0xd6, 0x89, 0x4c, 0x6c, 0xc2, 0x7a, 0xf5, 0xf7, 0xe2,
	0xfa, 0x7f, 0xb6, 0x61, 0x3d, 0x4f, 0x2a, 0x6c, 0x11, 0x57, 0x0c, 0x49, 0x8f, 0x4e, 0xa0, 0x10,
	0x8d, 0x42, 0x9b, 0xf7, 0xc6, 0x4d, 0xc0, 0x2f, 0x8f, 0xe2, 0x1a, 0xdb, 0x8a, 0xb0, 0x75, 0x2b,
	0x3d, 0x5d, 0x7b, 0xb9, 0xcf, 0x0f, 0x49, 0xef, 0xa3, 0x71, 0xbe, 0x35, 0xc9, 0x66, 0x6b, 0x92,
	0xab, 0xad, 0x49, 0xfe, 0xed, 0xcc, 0xcc, 0x66, 0x67, 0x66, 0x2e, 0x77, 0x66, 0x66, 0x5a, 0x8c,
	0xae, 0xf1, 0xfd, 0x4d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x4a, 0x99, 0x11, 0xf0, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClient interface {
	Registe(ctx context.Context, in *RegisteRequest, opts ...grpc.CallOption) (*RegisteResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) Registe(ctx context.Context, in *RegisteRequest, opts ...grpc.CallOption) (*RegisteResponse, error) {
	out := new(RegisteResponse)
	err := c.cc.Invoke(ctx, "/customer.Customer/Registe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/customer.Customer/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
type CustomerServer interface {
	Registe(context.Context, *RegisteRequest) (*RegisteResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedCustomerServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (*UnimplementedCustomerServer) Registe(ctx context.Context, req *RegisteRequest) (*RegisteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registe not implemented")
}
func (*UnimplementedCustomerServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_Registe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Registe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/Registe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Registe(ctx, req.(*RegisteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registe",
			Handler:    _Customer_Registe_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Customer_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}

func (m *RegisteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Phone) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Phone)))
		i += copy(dAtA[i:], m.Phone)
	}
	if m.Source != 0 {
		dAtA[i] = 0x15
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Source))
		i += 4
	}
	return i, nil
}

func (m *RegisteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Code) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Data != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(m.Data.Size()))
		n1, err1 := m.Data.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	return i, nil
}

func (m *RegisteData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Code) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	return i, nil
}

func (m *LoginRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Phone) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Phone)))
		i += copy(dAtA[i:], m.Phone)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if len(m.Code) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	return i, nil
}

func (m *LoginResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Code) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Data != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(m.Data.Size()))
		n2, err2 := m.Data.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	return i, nil
}

func (m *LoginData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Code) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	return i, nil
}

func encodeVarintCustomer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RegisteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Phone)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	if m.Source != 0 {
		n += 5
	}
	return n
}

func (m *RegisteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovCustomer(uint64(l))
	}
	return n
}

func (m *RegisteData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	return n
}

func (m *LoginRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Phone)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	return n
}

func (m *LoginResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovCustomer(uint64(l))
	}
	return n
}

func (m *LoginData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	return n
}

func sovCustomer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCustomer(x uint64) (n int) {
	return sovCustomer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomer
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
			return fmt.Errorf("proto: RegisteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			m.Source = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		default:
			iNdEx = preIndex
			skippy, err := skipCustomer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCustomer
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
func (m *RegisteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomer
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
			return fmt.Errorf("proto: RegisteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &RegisteData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCustomer
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
func (m *RegisteData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomer
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
			return fmt.Errorf("proto: RegisteData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCustomer
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
func (m *LoginRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomer
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
			return fmt.Errorf("proto: LoginRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCustomer
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
func (m *LoginResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomer
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
			return fmt.Errorf("proto: LoginResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &LoginData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCustomer
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
func (m *LoginData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomer
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
			return fmt.Errorf("proto: LoginData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCustomer
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
func skipCustomer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCustomer
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
					return 0, ErrIntOverflowCustomer
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
					return 0, ErrIntOverflowCustomer
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
				return 0, ErrInvalidLengthCustomer
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCustomer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCustomer
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
				next, err := skipCustomer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCustomer
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
	ErrInvalidLengthCustomer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCustomer   = fmt.Errorf("proto: integer overflow")
)
