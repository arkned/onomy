// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onomyprotocol/dao/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct{}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d6940140b83f0f7, []int{0}
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
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d6940140b83f0f7, []int{1}
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

// QueryTreasuryRequest is request type for the Query/Treasury RPC method.
type QueryTreasuryRequest struct{}

func (m *QueryTreasuryRequest) Reset()         { *m = QueryTreasuryRequest{} }
func (m *QueryTreasuryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTreasuryRequest) ProtoMessage()    {}
func (*QueryTreasuryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d6940140b83f0f7, []int{2}
}

func (m *QueryTreasuryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryTreasuryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTreasuryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryTreasuryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTreasuryRequest.Merge(m, src)
}

func (m *QueryTreasuryRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryTreasuryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTreasuryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTreasuryRequest proto.InternalMessageInfo

// QueryTreasuryResponse is response type for the Query/Treasury RPC method.
type QueryTreasuryResponse struct {
	TreasuryBalance github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=treasury_balance,json=treasuryBalance,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"treasury_balance" yaml:"treasury_balance"`
}

func (m *QueryTreasuryResponse) Reset()         { *m = QueryTreasuryResponse{} }
func (m *QueryTreasuryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTreasuryResponse) ProtoMessage()    {}
func (*QueryTreasuryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d6940140b83f0f7, []int{3}
}

func (m *QueryTreasuryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryTreasuryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTreasuryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryTreasuryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTreasuryResponse.Merge(m, src)
}

func (m *QueryTreasuryResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryTreasuryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTreasuryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTreasuryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "onomyprotocol.dao.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "onomyprotocol.dao.v1.QueryParamsResponse")
	proto.RegisterType((*QueryTreasuryRequest)(nil), "onomyprotocol.dao.v1.QueryTreasuryRequest")
	proto.RegisterType((*QueryTreasuryResponse)(nil), "onomyprotocol.dao.v1.QueryTreasuryResponse")
}

func init() { proto.RegisterFile("onomyprotocol/dao/v1/query.proto", fileDescriptor_1d6940140b83f0f7) }

var fileDescriptor_1d6940140b83f0f7 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0x03, 0x2a, 0xe4, 0x1b, 0x40, 0xa6, 0xc0, 0x51, 0x55, 0x6e, 0x89, 0x10, 0x2a,
	0x87, 0xb0, 0xd5, 0xb2, 0xdd, 0x58, 0x46, 0x18, 0xb8, 0x8a, 0x89, 0x05, 0x39, 0xa9, 0x15, 0x22,
	0x1a, 0xbf, 0x5c, 0xec, 0x54, 0x64, 0x85, 0x91, 0x05, 0xc4, 0xc8, 0x37, 0x80, 0x2f, 0xd2, 0xf1,
	0x24, 0x96, 0x9b, 0x0e, 0xae, 0xe5, 0x13, 0xf0, 0x09, 0x50, 0x6c, 0x17, 0xd1, 0x12, 0x9d, 0x6e,
	0x6a, 0xf3, 0xde, 0xfb, 0xff, 0xdf, 0xff, 0xfd, 0x8c, 0xfb, 0xa0, 0x20, 0xab, 0xf2, 0x02, 0x0c,
	0xc4, 0x30, 0xe3, 0x53, 0x01, 0x7c, 0x3e, 0xe4, 0x47, 0xa5, 0x2c, 0x2a, 0x66, 0xab, 0xa4, 0xbd,
	0x31, 0xc1, 0xa6, 0x02, 0xd8, 0x7c, 0xd8, 0x69, 0x27, 0x90, 0x80, 0x2d, 0xf2, 0xfa, 0x9f, 0x9b,
	0xed, 0xd0, 0x18, 0x74, 0x06, 0x9a, 0x47, 0x42, 0x4b, 0x3e, 0x1f, 0x46, 0xd2, 0x88, 0x21, 0x8f,
	0x21, 0x55, 0xbe, 0xdf, 0x4d, 0x00, 0x92, 0x99, 0xe4, 0x22, 0x4f, 0xb9, 0x50, 0x0a, 0x8c, 0x30,
	0x29, 0x28, 0xed, 0xbb, 0x77, 0x1b, 0xb3, 0xe4, 0xa2, 0x10, 0x99, 0x1f, 0x09, 0xdb, 0x98, 0x1c,
	0xd6, 0xd9, 0x9e, 0xdb, 0xe2, 0x44, 0x1e, 0x95, 0x52, 0x9b, 0xf0, 0x10, 0xdf, 0xd8, 0xa8, 0xea,
	0x1c, 0x94, 0x96, 0xe4, 0x00, 0xb7, 0x9c, 0x78, 0x0f, 0xf5, 0xd1, 0x60, 0x77, 0xd4, 0x65, 0x4d,
	0xa7, 0x30, 0xa7, 0x1a, 0x5f, 0x5e, 0x9c, 0xf6, 0x82, 0x89, 0x57, 0x84, 0xb7, 0x70, 0xdb, 0x5a,
	0xbe, 0x28, 0xa4, 0xd0, 0x65, 0x51, 0xad, 0x57, 0x7d, 0x43, 0xf8, 0xe6, 0x56, 0xc3, 0x6f, 0xfb,
	0x84, 0xf0, 0x75, 0xe3, 0x8b, 0xaf, 0x22, 0x31, 0x13, 0x2a, 0x96, 0x7b, 0xa8, 0x7f, 0x69, 0xb0,
	0x3b, 0xba, 0xc3, 0x1c, 0x17, 0x56, 0x73, 0x61, 0x9e, 0x0b, 0x7b, 0x02, 0xa9, 0x1a, 0x3f, 0xad,
	0xb7, 0xfe, 0x3e, 0xed, 0xdd, 0xae, 0x44, 0x36, 0x3b, 0x08, 0xb7, 0x0d, 0xc2, 0xaf, 0x3f, 0x7a,
	0x83, 0x24, 0x35, 0xaf, 0xcb, 0x88, 0xc5, 0x90, 0x71, 0xcf, 0xd7, 0xfd, 0x3c, 0xd2, 0xd3, 0x37,
	0xdc, 0x54, 0xb9, 0xd4, 0xd6, 0x4b, 0x4f, 0xae, 0xad, 0xe5, 0x63, 0xa7, 0x1e, 0x7d, 0xd9, 0xc1,
	0x57, 0x6c, 0x5a, 0xf2, 0x1e, 0xe1, 0x96, 0x3b, 0x94, 0x0c, 0x9a, 0x31, 0xfc, 0xcf, 0xb5, 0xf3,
	0xe0, 0x02, 0x93, 0xee, 0xfa, 0xf0, 0xde, 0xbb, 0xef, 0xbf, 0x3e, 0xef, 0x50, 0xd2, 0xe5, 0xe7,
	0x3c, 0x22, 0xf9, 0x80, 0xf0, 0xd5, 0x35, 0x38, 0xb2, 0x7f, 0x8e, 0xfb, 0x16, 0xf6, 0xce, 0xc3,
	0x0b, 0xcd, 0xfa, 0x2c, 0xf7, 0x6d, 0x96, 0x3e, 0xa1, 0xcd, 0x59, 0xfe, 0x42, 0x7a, 0xb6, 0x38,
	0xa3, 0xc1, 0xc9, 0x19, 0x45, 0x8b, 0x25, 0x45, 0xc7, 0x4b, 0x8a, 0x7e, 0x2e, 0x29, 0xfa, 0xb8,
	0xa2, 0xc1, 0xf1, 0x8a, 0x06, 0x27, 0x2b, 0x1a, 0xbc, 0xdc, 0xff, 0x07, 0xfd, 0xa6, 0x97, 0xfd,
	0xe2, 0x6f, 0xad, 0xa7, 0x7d, 0x82, 0xa8, 0x65, 0x7b, 0x8f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x35, 0x78, 0x27, 0xe5, 0x52, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Treasury queries the dao treasury.
	Treasury(ctx context.Context, in *QueryTreasuryRequest, opts ...grpc.CallOption) (*QueryTreasuryResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/onomyprotocol.dao.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Treasury(ctx context.Context, in *QueryTreasuryRequest, opts ...grpc.CallOption) (*QueryTreasuryResponse, error) {
	out := new(QueryTreasuryResponse)
	err := c.cc.Invoke(ctx, "/onomyprotocol.dao.v1.Query/Treasury", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Treasury queries the dao treasury.
	Treasury(context.Context, *QueryTreasuryRequest) (*QueryTreasuryResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct{}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func (*UnimplementedQueryServer) Treasury(ctx context.Context, req *QueryTreasuryRequest) (*QueryTreasuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Treasury not implemented")
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
		FullMethod: "/onomyprotocol.dao.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Treasury_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTreasuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Treasury(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onomyprotocol.dao.v1.Query/Treasury",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Treasury(ctx, req.(*QueryTreasuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onomyprotocol.dao.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Treasury",
			Handler:    _Query_Treasury_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onomyprotocol/dao/v1/query.proto",
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
	return len(dAtA) - i, nil
}

func (m *QueryTreasuryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTreasuryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTreasuryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryTreasuryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTreasuryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTreasuryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TreasuryBalance) > 0 {
		for iNdEx := len(m.TreasuryBalance) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TreasuryBalance[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryTreasuryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryTreasuryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TreasuryBalance) > 0 {
		for _, e := range m.TreasuryBalance {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
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

func (m *QueryTreasuryRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTreasuryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTreasuryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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

func (m *QueryTreasuryResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTreasuryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTreasuryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreasuryBalance", wireType)
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
			m.TreasuryBalance = append(m.TreasuryBalance, types.Coin{})
			if err := m.TreasuryBalance[len(m.TreasuryBalance)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
