// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethclient/ethclient/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eee5262ce6e2bd, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eee5262ce6e2bd, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

type QueryGetStorageRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Block   int64  `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
	Storage string `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (m *QueryGetStorageRequest) Reset()         { *m = QueryGetStorageRequest{} }
func (m *QueryGetStorageRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetStorageRequest) ProtoMessage()    {}
func (*QueryGetStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eee5262ce6e2bd, []int{2}
}
func (m *QueryGetStorageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetStorageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetStorageRequest.Merge(m, src)
}
func (m *QueryGetStorageRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetStorageRequest proto.InternalMessageInfo

func (m *QueryGetStorageRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueryGetStorageRequest) GetBlock() int64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *QueryGetStorageRequest) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

type QueryGetStorageResponse struct {
	Storage Storage `protobuf:"bytes,1,opt,name=Storage,proto3" json:"Storage"`
}

func (m *QueryGetStorageResponse) Reset()         { *m = QueryGetStorageResponse{} }
func (m *QueryGetStorageResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetStorageResponse) ProtoMessage()    {}
func (*QueryGetStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eee5262ce6e2bd, []int{3}
}
func (m *QueryGetStorageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetStorageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetStorageResponse.Merge(m, src)
}
func (m *QueryGetStorageResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetStorageResponse proto.InternalMessageInfo

func (m *QueryGetStorageResponse) GetStorage() Storage {
	if m != nil {
		return m.Storage
	}
	return Storage{}
}

type QueryProofRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Storage string `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`
	Block   int64  `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *QueryProofRequest) Reset()         { *m = QueryProofRequest{} }
func (m *QueryProofRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProofRequest) ProtoMessage()    {}
func (*QueryProofRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eee5262ce6e2bd, []int{4}
}
func (m *QueryProofRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProofRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProofRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProofRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProofRequest.Merge(m, src)
}
func (m *QueryProofRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProofRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProofRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProofRequest proto.InternalMessageInfo

func (m *QueryProofRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueryProofRequest) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *QueryProofRequest) GetBlock() int64 {
	if m != nil {
		return m.Block
	}
	return 0
}

type QueryProofResponse struct {
	Proof string `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *QueryProofResponse) Reset()         { *m = QueryProofResponse{} }
func (m *QueryProofResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProofResponse) ProtoMessage()    {}
func (*QueryProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eee5262ce6e2bd, []int{5}
}
func (m *QueryProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProofResponse.Merge(m, src)
}
func (m *QueryProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProofResponse proto.InternalMessageInfo

func (m *QueryProofResponse) GetProof() string {
	if m != nil {
		return m.Proof
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "ethclient.ethclient.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "ethclient.ethclient.QueryParamsResponse")
	proto.RegisterType((*QueryGetStorageRequest)(nil), "ethclient.ethclient.QueryGetStorageRequest")
	proto.RegisterType((*QueryGetStorageResponse)(nil), "ethclient.ethclient.QueryGetStorageResponse")
	proto.RegisterType((*QueryProofRequest)(nil), "ethclient.ethclient.QueryProofRequest")
	proto.RegisterType((*QueryProofResponse)(nil), "ethclient.ethclient.QueryProofResponse")
}

func init() { proto.RegisterFile("ethclient/ethclient/query.proto", fileDescriptor_d9eee5262ce6e2bd) }

var fileDescriptor_d9eee5262ce6e2bd = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x09, 0x49, 0xc5, 0xc0, 0x85, 0x6d, 0x04, 0x91, 0x13, 0xdc, 0x60, 0x10, 0x8d,
	0x0a, 0xca, 0xaa, 0xad, 0x40, 0x1c, 0x10, 0x87, 0x5e, 0x90, 0x38, 0x81, 0x39, 0x20, 0x21, 0x71,
	0x58, 0xa7, 0x8b, 0xb1, 0x48, 0xbd, 0xae, 0x77, 0x8b, 0xa8, 0xa2, 0x48, 0x88, 0x17, 0x00, 0x89,
	0x33, 0x6f, 0xc2, 0x03, 0xf4, 0x58, 0x89, 0x0b, 0x27, 0x84, 0x12, 0x1e, 0x04, 0x79, 0x77, 0x5c,
	0xbb, 0x8a, 0x9b, 0x70, 0xdb, 0x1d, 0xff, 0xf3, 0xcf, 0x37, 0x33, 0x6b, 0xd8, 0x10, 0xfa, 0xdd,
	0x68, 0x1c, 0x89, 0x58, 0xb3, 0xe2, 0x74, 0x78, 0x24, 0xd2, 0xe3, 0x61, 0x92, 0x4a, 0x2d, 0xe9,
	0xfa, 0x59, 0x78, 0x78, 0x76, 0x72, 0xda, 0xa1, 0x0c, 0xa5, 0xf9, 0xce, 0xb2, 0x93, 0x95, 0x3a,
	0xbd, 0x50, 0xca, 0x70, 0x2c, 0x18, 0x4f, 0x22, 0xc6, 0xe3, 0x58, 0x6a, 0xae, 0x23, 0x19, 0x2b,
	0xfc, 0xba, 0x35, 0x92, 0xea, 0x40, 0x2a, 0x16, 0x70, 0x25, 0x6c, 0x05, 0xf6, 0x61, 0x3b, 0x10,
	0x9a, 0x6f, 0xb3, 0x84, 0x87, 0x51, 0x6c, 0xc4, 0xa8, 0xed, 0x57, 0x51, 0x25, 0x3c, 0xe5, 0x07,
	0xb9, 0xdb, 0xad, 0x2a, 0x85, 0xd2, 0x32, 0xe5, 0xa1, 0xb0, 0x12, 0xaf, 0x0d, 0xf4, 0x45, 0x56,
	0xe6, 0xb9, 0xc9, 0xf3, 0xc5, 0xe1, 0x91, 0x50, 0xda, 0x7b, 0x06, 0xeb, 0xe7, 0xa2, 0x2a, 0x91,
	0xb1, 0x12, 0x74, 0x17, 0x5a, 0xd6, 0xbf, 0x43, 0xfa, 0x64, 0x70, 0x65, 0xa7, 0x3b, 0xac, 0xe8,
	0x7b, 0x88, 0x49, 0x28, 0xf5, 0x02, 0xb8, 0x6e, 0xbc, 0x9e, 0x0a, 0xfd, 0xd2, 0x96, 0xc6, 0x2a,
	0xb4, 0x03, 0x6b, 0x7c, 0x7f, 0x3f, 0x15, 0xca, 0xfa, 0x5d, 0xf6, 0xf3, 0x2b, 0x6d, 0x43, 0x33,
	0x18, 0xcb, 0xd1, 0xfb, 0x4e, 0xbd, 0x4f, 0x06, 0x0d, 0xdf, 0x5e, 0x32, 0x3d, 0xc2, 0x77, 0x1a,
	0x56, 0x8f, 0x57, 0xef, 0x15, 0xdc, 0x58, 0xa8, 0x81, 0xcc, 0x8f, 0x61, 0x0d, 0x43, 0x08, 0xdd,
	0xab, 0x84, 0x46, 0xcd, 0xde, 0xa5, 0x93, 0xdf, 0x1b, 0x35, 0x3f, 0x4f, 0xf1, 0xde, 0xc0, 0x35,
	0x3b, 0x88, 0x54, 0xca, 0xb7, 0xab, 0xb9, 0x4b, 0x84, 0xf5, 0x73, 0x84, 0x45, 0x47, 0x8d, 0x52,
	0x47, 0xde, 0x56, 0x3e, 0x7d, 0x6b, 0x8f, 0xc8, 0x6d, 0x68, 0x26, 0x59, 0x00, 0xdd, 0xed, 0x65,
	0xe7, 0x47, 0x03, 0x9a, 0x46, 0x4c, 0x3f, 0x11, 0x68, 0xd9, 0x21, 0xd3, 0xcd, 0xca, 0x66, 0x16,
	0x37, 0xea, 0x0c, 0x56, 0x0b, 0x6d, 0x75, 0xef, 0xf6, 0xe7, 0x9f, 0x7f, 0xbf, 0xd5, 0x6f, 0xd2,
	0x2e, 0xbb, 0xf8, 0x7d, 0xd1, 0x2f, 0x04, 0xae, 0x9a, 0x64, 0x1c, 0x14, 0xbd, 0x77, 0xb1, 0xff,
	0xc2, 0xe2, 0x9d, 0xfb, 0xff, 0x27, 0x46, 0xa0, 0x3b, 0x06, 0xc8, 0xa5, 0x3d, 0xb6, 0xe4, 0x39,
	0xd3, 0xef, 0x04, 0xa0, 0x98, 0x25, 0xbd, 0xbb, 0xa4, 0xdf, 0xd2, 0x2e, 0x9d, 0xcd, 0x95, 0x3a,
	0xa4, 0x78, 0x62, 0x28, 0x1e, 0xd1, 0x87, 0xd5, 0x63, 0xc9, 0xb4, 0x6c, 0x82, 0xef, 0x60, 0xca,
	0x26, 0x88, 0x35, 0x65, 0x13, 0xb3, 0xe9, 0xe9, 0xde, 0x83, 0x93, 0x99, 0x4b, 0x4e, 0x67, 0x2e,
	0xf9, 0x33, 0x73, 0xc9, 0xd7, 0xb9, 0x5b, 0x3b, 0x9d, 0xbb, 0xb5, 0x5f, 0x73, 0xb7, 0xf6, 0xba,
	0x5b, 0xd8, 0x7c, 0x2c, 0x59, 0xea, 0xe3, 0x44, 0xa8, 0xa0, 0x65, 0x7e, 0xd3, 0xdd, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xe0, 0x47, 0x3c, 0x73, 0x83, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a Storage by index.
	QueryStorage(ctx context.Context, in *QueryGetStorageRequest, opts ...grpc.CallOption) (*QueryGetStorageResponse, error)
	// Queries a list of Proof items.
	QueryProof(ctx context.Context, in *QueryProofRequest, opts ...grpc.CallOption) (*QueryProofResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/ethclient.ethclient.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryStorage(ctx context.Context, in *QueryGetStorageRequest, opts ...grpc.CallOption) (*QueryGetStorageResponse, error) {
	out := new(QueryGetStorageResponse)
	err := c.cc.Invoke(ctx, "/ethclient.ethclient.Query/QueryStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryProof(ctx context.Context, in *QueryProofRequest, opts ...grpc.CallOption) (*QueryProofResponse, error) {
	out := new(QueryProofResponse)
	err := c.cc.Invoke(ctx, "/ethclient.ethclient.Query/QueryProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a Storage by index.
	QueryStorage(context.Context, *QueryGetStorageRequest) (*QueryGetStorageResponse, error)
	// Queries a list of Proof items.
	QueryProof(context.Context, *QueryProofRequest) (*QueryProofResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) QueryStorage(ctx context.Context, req *QueryGetStorageRequest) (*QueryGetStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStorage not implemented")
}
func (*UnimplementedQueryServer) QueryProof(ctx context.Context, req *QueryProofRequest) (*QueryProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProof not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethclient.ethclient.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethclient.ethclient.Query/QueryStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryStorage(ctx, req.(*QueryGetStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethclient.ethclient.Query/QueryProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryProof(ctx, req.(*QueryProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethclient.ethclient.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "QueryStorage",
			Handler:    _Query_QueryStorage_Handler,
		},
		{
			MethodName: "QueryProof",
			Handler:    _Query_QueryProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ethclient/ethclient/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetStorageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetStorageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetStorageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Storage) > 0 {
		i -= len(m.Storage)
		copy(dAtA[i:], m.Storage)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Storage)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Block != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetStorageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetStorageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetStorageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Storage.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryProofRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProofRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProofRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Storage) > 0 {
		i -= len(m.Storage)
		copy(dAtA[i:], m.Storage)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Storage)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetStorageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovQuery(uint64(m.Block))
	}
	l = len(m.Storage)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetStorageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Storage.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryProofRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Storage)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovQuery(uint64(m.Block))
	}
	return n
}

func (m *QueryProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetStorageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetStorageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetStorageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetStorageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetStorageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetStorageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Storage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryProofRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryProofRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProofRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryProofResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
