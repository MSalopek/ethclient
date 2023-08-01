// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethclient/ethclient/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgCreateStorage struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Storage string `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	// ethereum block height as hex string (padded)
	Block int64 `protobuf:"varint,4,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *MsgCreateStorage) Reset()         { *m = MsgCreateStorage{} }
func (m *MsgCreateStorage) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStorage) ProtoMessage()    {}
func (*MsgCreateStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_492b3ce756be7624, []int{0}
}
func (m *MsgCreateStorage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStorage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStorage.Merge(m, src)
}
func (m *MsgCreateStorage) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStorage.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStorage proto.InternalMessageInfo

func (m *MsgCreateStorage) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateStorage) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgCreateStorage) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *MsgCreateStorage) GetBlock() int64 {
	if m != nil {
		return m.Block
	}
	return 0
}

type MsgCreateStorageResponse struct {
	Storage *Storage `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (m *MsgCreateStorageResponse) Reset()         { *m = MsgCreateStorageResponse{} }
func (m *MsgCreateStorageResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStorageResponse) ProtoMessage()    {}
func (*MsgCreateStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_492b3ce756be7624, []int{1}
}
func (m *MsgCreateStorageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStorageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStorageResponse.Merge(m, src)
}
func (m *MsgCreateStorageResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStorageResponse proto.InternalMessageInfo

func (m *MsgCreateStorageResponse) GetStorage() *Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

type MsgStoreProof struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Storage string `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	// ethereum block height as hex string (padded)
	Block int64 `protobuf:"varint,4,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *MsgStoreProof) Reset()         { *m = MsgStoreProof{} }
func (m *MsgStoreProof) String() string { return proto.CompactTextString(m) }
func (*MsgStoreProof) ProtoMessage()    {}
func (*MsgStoreProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_492b3ce756be7624, []int{2}
}
func (m *MsgStoreProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStoreProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStoreProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStoreProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStoreProof.Merge(m, src)
}
func (m *MsgStoreProof) XXX_Size() int {
	return m.Size()
}
func (m *MsgStoreProof) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStoreProof.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStoreProof proto.InternalMessageInfo

func (m *MsgStoreProof) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgStoreProof) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgStoreProof) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *MsgStoreProof) GetBlock() int64 {
	if m != nil {
		return m.Block
	}
	return 0
}

type MsgStoreProofResponse struct {
	// JSON serialized proof
	Proof *StorageProof `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *MsgStoreProofResponse) Reset()         { *m = MsgStoreProofResponse{} }
func (m *MsgStoreProofResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStoreProofResponse) ProtoMessage()    {}
func (*MsgStoreProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_492b3ce756be7624, []int{3}
}
func (m *MsgStoreProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStoreProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStoreProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStoreProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStoreProofResponse.Merge(m, src)
}
func (m *MsgStoreProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStoreProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStoreProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStoreProofResponse proto.InternalMessageInfo

func (m *MsgStoreProofResponse) GetProof() *StorageProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgCreateStorage)(nil), "ethclient.ethclient.MsgCreateStorage")
	proto.RegisterType((*MsgCreateStorageResponse)(nil), "ethclient.ethclient.MsgCreateStorageResponse")
	proto.RegisterType((*MsgStoreProof)(nil), "ethclient.ethclient.MsgStoreProof")
	proto.RegisterType((*MsgStoreProofResponse)(nil), "ethclient.ethclient.MsgStoreProofResponse")
}

func init() { proto.RegisterFile("ethclient/ethclient/tx.proto", fileDescriptor_492b3ce756be7624) }

var fileDescriptor_492b3ce756be7624 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x2d, 0xc9, 0x48,
	0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x47, 0xb0, 0x4a, 0x2a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x84, 0xe1, 0x62, 0x7a, 0x70, 0x96, 0x94, 0x22, 0x36, 0x2d, 0xc5, 0x25, 0xf9, 0x45, 0x89,
	0xe9, 0xa9, 0x10, 0x7d, 0x4a, 0x65, 0x5c, 0x02, 0xbe, 0xc5, 0xe9, 0xce, 0x45, 0xa9, 0x89, 0x25,
	0xa9, 0xc1, 0x10, 0x19, 0x21, 0x09, 0x2e, 0xf6, 0x64, 0x90, 0x40, 0x7e, 0x91, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x8c, 0x0b, 0x92, 0x49, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60,
	0x82, 0xc8, 0x40, 0xb9, 0x20, 0x19, 0xa8, 0xc1, 0x12, 0xcc, 0x10, 0x19, 0x28, 0x57, 0x48, 0x84,
	0x8b, 0x35, 0x29, 0x27, 0x3f, 0x39, 0x5b, 0x82, 0x45, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc2, 0x51,
	0x0a, 0xe2, 0x92, 0x40, 0xb7, 0x37, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x0c,
	0x61, 0x16, 0xc8, 0x7e, 0x6e, 0x23, 0x19, 0x3d, 0x2c, 0xbe, 0xd3, 0x83, 0x69, 0x83, 0x29, 0x56,
	0x2a, 0xe6, 0xe2, 0xf5, 0x2d, 0x4e, 0x07, 0x09, 0xa7, 0x06, 0x14, 0xe5, 0xe7, 0xa7, 0xd1, 0xc5,
	0x23, 0x01, 0x5c, 0xa2, 0x28, 0x96, 0xc2, 0x7d, 0x61, 0xce, 0xc5, 0x5a, 0x00, 0x12, 0x80, 0xfa,
	0x41, 0x11, 0x9f, 0x1f, 0x20, 0x3a, 0x21, 0xea, 0x8d, 0x4e, 0x31, 0x72, 0x31, 0xfb, 0x16, 0xa7,
	0x0b, 0xa5, 0x72, 0xf1, 0xa2, 0xc6, 0x8b, 0x2a, 0x56, 0x23, 0xd0, 0x83, 0x51, 0x4a, 0x97, 0x28,
	0x65, 0x70, 0x77, 0xc6, 0x70, 0x71, 0x21, 0x05, 0x99, 0x12, 0x2e, 0xcd, 0x08, 0x35, 0x52, 0x5a,
	0x84, 0xd5, 0xc0, 0x4c, 0x77, 0x32, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0x69, 0x44, 0x92, 0xac, 0x40, 0x4e, 0xd1, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xd4,
	0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x84, 0xb4, 0x9b, 0x9a, 0xf5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// calls eth_getStorageAt and saves the result
	CreateStorage(ctx context.Context, in *MsgCreateStorage, opts ...grpc.CallOption) (*MsgCreateStorageResponse, error)
	// calls eth_getProof and saves the result
	StoreProof(ctx context.Context, in *MsgStoreProof, opts ...grpc.CallOption) (*MsgStoreProofResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateStorage(ctx context.Context, in *MsgCreateStorage, opts ...grpc.CallOption) (*MsgCreateStorageResponse, error) {
	out := new(MsgCreateStorageResponse)
	err := c.cc.Invoke(ctx, "/ethclient.ethclient.Msg/CreateStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StoreProof(ctx context.Context, in *MsgStoreProof, opts ...grpc.CallOption) (*MsgStoreProofResponse, error) {
	out := new(MsgStoreProofResponse)
	err := c.cc.Invoke(ctx, "/ethclient.ethclient.Msg/StoreProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// calls eth_getStorageAt and saves the result
	CreateStorage(context.Context, *MsgCreateStorage) (*MsgCreateStorageResponse, error)
	// calls eth_getProof and saves the result
	StoreProof(context.Context, *MsgStoreProof) (*MsgStoreProofResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateStorage(ctx context.Context, req *MsgCreateStorage) (*MsgCreateStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorage not implemented")
}
func (*UnimplementedMsgServer) StoreProof(ctx context.Context, req *MsgStoreProof) (*MsgStoreProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreProof not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateStorage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethclient.ethclient.Msg/CreateStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateStorage(ctx, req.(*MsgCreateStorage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StoreProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStoreProof)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StoreProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethclient.ethclient.Msg/StoreProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StoreProof(ctx, req.(*MsgStoreProof))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethclient.ethclient.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStorage",
			Handler:    _Msg_CreateStorage_Handler,
		},
		{
			MethodName: "StoreProof",
			Handler:    _Msg_StoreProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ethclient/ethclient/tx.proto",
}

func (m *MsgCreateStorage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStorage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStorage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Storage) > 0 {
		i -= len(m.Storage)
		copy(dAtA[i:], m.Storage)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Storage)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateStorageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStorageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStorageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Storage != nil {
		{
			size, err := m.Storage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStoreProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStoreProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStoreProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Storage) > 0 {
		i -= len(m.Storage)
		copy(dAtA[i:], m.Storage)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Storage)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStoreProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStoreProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStoreProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateStorage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Storage)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovTx(uint64(m.Block))
	}
	return n
}

func (m *MsgCreateStorageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Storage != nil {
		l = m.Storage.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgStoreProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Storage)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovTx(uint64(m.Block))
	}
	return n
}

func (m *MsgStoreProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateStorage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateStorage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStorage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateStorageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateStorageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStorageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Storage == nil {
				m.Storage = &Storage{}
			}
			if err := m.Storage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStoreProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStoreProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStoreProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStoreProofResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStoreProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStoreProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &StorageProof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
