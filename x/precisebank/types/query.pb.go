// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/precisebank/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// QueryTotalFractionalBalancesRequest defines the request type for Query/TotalFractionalBalances method.
type QueryTotalFractionalBalancesRequest struct {
}

func (m *QueryTotalFractionalBalancesRequest) Reset()         { *m = QueryTotalFractionalBalancesRequest{} }
func (m *QueryTotalFractionalBalancesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTotalFractionalBalancesRequest) ProtoMessage()    {}
func (*QueryTotalFractionalBalancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77e9b73389f3988, []int{0}
}
func (m *QueryTotalFractionalBalancesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalFractionalBalancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalFractionalBalancesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalFractionalBalancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalFractionalBalancesRequest.Merge(m, src)
}
func (m *QueryTotalFractionalBalancesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalFractionalBalancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalFractionalBalancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalFractionalBalancesRequest proto.InternalMessageInfo

// QueryTotalFractionalBalancesResponse defines the response type for Query/TotalFractionalBalances method.
type QueryTotalFractionalBalancesResponse struct {
	// total is the total sum of all fractional balances managed by the precisebank
	// module.
	Total types.Coin `protobuf:"bytes,1,opt,name=total,proto3" json:"total"`
}

func (m *QueryTotalFractionalBalancesResponse) Reset()         { *m = QueryTotalFractionalBalancesResponse{} }
func (m *QueryTotalFractionalBalancesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTotalFractionalBalancesResponse) ProtoMessage()    {}
func (*QueryTotalFractionalBalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77e9b73389f3988, []int{1}
}
func (m *QueryTotalFractionalBalancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalFractionalBalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalFractionalBalancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalFractionalBalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalFractionalBalancesResponse.Merge(m, src)
}
func (m *QueryTotalFractionalBalancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalFractionalBalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalFractionalBalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalFractionalBalancesResponse proto.InternalMessageInfo

// QueryRemainderRequest defines the request type for Query/Remainder method.
type QueryRemainderRequest struct {
}

func (m *QueryRemainderRequest) Reset()         { *m = QueryRemainderRequest{} }
func (m *QueryRemainderRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRemainderRequest) ProtoMessage()    {}
func (*QueryRemainderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77e9b73389f3988, []int{2}
}
func (m *QueryRemainderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRemainderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRemainderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRemainderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRemainderRequest.Merge(m, src)
}
func (m *QueryRemainderRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRemainderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRemainderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRemainderRequest proto.InternalMessageInfo

// QueryRemainderResponse defines the response type for Query/Remainder method.
type QueryRemainderResponse struct {
	// remainder is the amount backed by the reserve, but not yet owned by any
	// account, i.e. not in circulation.
	Remainder types.Coin `protobuf:"bytes,1,opt,name=remainder,proto3" json:"remainder"`
}

func (m *QueryRemainderResponse) Reset()         { *m = QueryRemainderResponse{} }
func (m *QueryRemainderResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRemainderResponse) ProtoMessage()    {}
func (*QueryRemainderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77e9b73389f3988, []int{3}
}
func (m *QueryRemainderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRemainderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRemainderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRemainderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRemainderResponse.Merge(m, src)
}
func (m *QueryRemainderResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRemainderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRemainderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRemainderResponse proto.InternalMessageInfo

// QueryFractionalBalanceRequest defines the request type for Query/FractionalBalance method.
type QueryFractionalBalanceRequest struct {
	// address is the account address to query  fractional balance for.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryFractionalBalanceRequest) Reset()         { *m = QueryFractionalBalanceRequest{} }
func (m *QueryFractionalBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFractionalBalanceRequest) ProtoMessage()    {}
func (*QueryFractionalBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77e9b73389f3988, []int{4}
}
func (m *QueryFractionalBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFractionalBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFractionalBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFractionalBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFractionalBalanceRequest.Merge(m, src)
}
func (m *QueryFractionalBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFractionalBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFractionalBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFractionalBalanceRequest proto.InternalMessageInfo

// QueryFractionalBalanceResponse defines the response type for Query/FractionalBalance method.
type QueryFractionalBalanceResponse struct {
	// fractional_balance is the fractional balance of the address.
	FractionalBalance types.Coin `protobuf:"bytes,1,opt,name=fractional_balance,json=fractionalBalance,proto3" json:"fractional_balance"`
}

func (m *QueryFractionalBalanceResponse) Reset()         { *m = QueryFractionalBalanceResponse{} }
func (m *QueryFractionalBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFractionalBalanceResponse) ProtoMessage()    {}
func (*QueryFractionalBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77e9b73389f3988, []int{5}
}
func (m *QueryFractionalBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFractionalBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFractionalBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFractionalBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFractionalBalanceResponse.Merge(m, src)
}
func (m *QueryFractionalBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFractionalBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFractionalBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFractionalBalanceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryTotalFractionalBalancesRequest)(nil), "zgc.precisebank.v1.QueryTotalFractionalBalancesRequest")
	proto.RegisterType((*QueryTotalFractionalBalancesResponse)(nil), "zgc.precisebank.v1.QueryTotalFractionalBalancesResponse")
	proto.RegisterType((*QueryRemainderRequest)(nil), "zgc.precisebank.v1.QueryRemainderRequest")
	proto.RegisterType((*QueryRemainderResponse)(nil), "zgc.precisebank.v1.QueryRemainderResponse")
	proto.RegisterType((*QueryFractionalBalanceRequest)(nil), "zgc.precisebank.v1.QueryFractionalBalanceRequest")
	proto.RegisterType((*QueryFractionalBalanceResponse)(nil), "zgc.precisebank.v1.QueryFractionalBalanceResponse")
}

func init() { proto.RegisterFile("zgc/precisebank/v1/query.proto", fileDescriptor_c77e9b73389f3988) }

var fileDescriptor_c77e9b73389f3988 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6b, 0x14, 0x31,
	0x18, 0xc6, 0x67, 0xa4, 0x55, 0x36, 0x9e, 0x1a, 0xd4, 0xd6, 0xa1, 0x46, 0x19, 0x2b, 0xa8, 0x68,
	0xb2, 0xb3, 0x2a, 0xd5, 0x83, 0x97, 0x15, 0x3c, 0x89, 0xe0, 0x22, 0x08, 0x82, 0x94, 0x4c, 0x36,
	0x4d, 0x07, 0x67, 0x93, 0xe9, 0x24, 0xbb, 0xd8, 0x8a, 0x17, 0x4f, 0x5e, 0x04, 0xc1, 0x8f, 0xe2,
	0x37, 0xf0, 0xb4, 0xc7, 0x82, 0x17, 0x4f, 0xa2, 0xbb, 0x7e, 0x10, 0x99, 0x4c, 0x76, 0xd5, 0x9d,
	0x4e, 0x99, 0xbd, 0x65, 0xf2, 0xfe, 0x79, 0x7e, 0x6f, 0xde, 0x87, 0x01, 0xe8, 0x50, 0x30, 0x92,
	0xe5, 0x9c, 0x25, 0x9a, 0xc7, 0x54, 0xbe, 0x26, 0xa3, 0x88, 0xec, 0x0f, 0x79, 0x7e, 0x80, 0xb3,
	0x5c, 0x19, 0x05, 0xe1, 0xa1, 0x60, 0xf8, 0x9f, 0x38, 0x1e, 0x45, 0x01, 0x62, 0x4a, 0x0f, 0x94,
	0x26, 0x31, 0xd5, 0x9c, 0x8c, 0xa2, 0x98, 0x1b, 0x1a, 0x11, 0xa6, 0x12, 0x59, 0xd6, 0x04, 0xe7,
	0x84, 0x12, 0xca, 0x1e, 0x49, 0x71, 0x72, 0xb7, 0x9b, 0x42, 0x29, 0x91, 0x72, 0x42, 0xb3, 0x84,
	0x50, 0x29, 0x95, 0xa1, 0x26, 0x51, 0x52, 0x97, 0xd1, 0xf0, 0x1a, 0xb8, 0xfa, 0xac, 0x90, 0x7d,
	0xae, 0x0c, 0x4d, 0x1f, 0xe7, 0x94, 0x15, 0x41, 0x9a, 0x76, 0x69, 0x4a, 0x25, 0xe3, 0xba, 0xc7,
	0xf7, 0x87, 0x5c, 0x9b, 0xf0, 0x15, 0xd8, 0x3a, 0x39, 0x4d, 0x67, 0x4a, 0x6a, 0x0e, 0xef, 0x81,
	0x55, 0x53, 0xa4, 0x6c, 0xf8, 0x57, 0xfc, 0xeb, 0x67, 0x3b, 0x17, 0x71, 0x89, 0x8c, 0x0b, 0x64,
	0xec, 0x90, 0xf1, 0x23, 0x95, 0xc8, 0xee, 0xca, 0xf8, 0xc7, 0x65, 0xaf, 0x57, 0x66, 0x87, 0xeb,
	0xe0, 0xbc, 0x6d, 0xdf, 0xe3, 0x03, 0x9a, 0xc8, 0x3e, 0xcf, 0x67, 0xba, 0x2f, 0xc0, 0x85, 0xc5,
	0x80, 0x53, 0x7a, 0x08, 0x5a, 0xf9, 0xec, 0xb2, 0xa9, 0xda, 0xdf, 0x8a, 0xf0, 0x01, 0xb8, 0x64,
	0x1b, 0x57, 0x66, 0x71, 0xca, 0x70, 0x03, 0x9c, 0xa1, 0xfd, 0x7e, 0xce, 0xb5, 0xb6, 0xdd, 0x5b,
	0xbd, 0xd9, 0x67, 0x98, 0x01, 0x54, 0x57, 0xea, 0xd8, 0x9e, 0x02, 0xb8, 0x3b, 0x0f, 0xee, 0xc4,
	0x65, 0xb4, 0x29, 0xe4, 0xda, 0xee, 0x62, 0xdf, 0xce, 0xc7, 0x15, 0xb0, 0x6a, 0x25, 0xe1, 0x57,
	0x1f, 0xac, 0xd7, 0xec, 0x00, 0x6e, 0xe3, 0xaa, 0x67, 0x70, 0x83, 0xe5, 0x06, 0xf7, 0x97, 0x2f,
	0x2c, 0x07, 0x0d, 0xef, 0xbe, 0xff, 0xf6, 0xfb, 0xf3, 0x29, 0x0c, 0x6f, 0x91, 0xb6, 0x58, 0x74,
	0xb3, 0xdd, 0xec, 0x4e, 0xf5, 0x1d, 0x34, 0xfc, 0xe0, 0x83, 0xd6, 0x7c, 0xa1, 0xf0, 0x46, 0xad,
	0xfa, 0xa2, 0x1b, 0x82, 0x9b, 0x4d, 0x52, 0x1d, 0xda, 0x96, 0x45, 0x43, 0x70, 0xf3, 0x18, 0xb4,
	0xb9, 0x0d, 0xe0, 0x17, 0x1f, 0xac, 0x55, 0xe6, 0x83, 0x51, 0xad, 0x4e, 0x9d, 0x5d, 0x82, 0xce,
	0x32, 0x25, 0x0e, 0x71, 0xdb, 0x22, 0x46, 0x90, 0x1c, 0x83, 0x58, 0x7d, 0x37, 0xf2, 0xd6, 0x19,
	0xf0, 0x5d, 0xf7, 0xc9, 0xf8, 0x17, 0xf2, 0xc6, 0x13, 0xe4, 0x1f, 0x4d, 0x90, 0xff, 0x73, 0x82,
	0xfc, 0x4f, 0x53, 0xe4, 0x1d, 0x4d, 0x91, 0xf7, 0x7d, 0x8a, 0xbc, 0x97, 0x58, 0x24, 0x66, 0x6f,
	0x18, 0x63, 0xa6, 0x06, 0xa4, 0x2d, 0x52, 0x1a, 0x6b, 0xd2, 0x16, 0xb7, 0xd9, 0x1e, 0x4d, 0x24,
	0x79, 0xf3, 0x9f, 0x8e, 0x39, 0xc8, 0xb8, 0x8e, 0x4f, 0xdb, 0x3f, 0xc1, 0x9d, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xb3, 0xb3, 0x06, 0x02, 0x93, 0x04, 0x00, 0x00,
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
	// TotalFractionalBalances returns the total sum of all fractional balances
	// managed by the precisebank module.
	TotalFractionalBalances(ctx context.Context, in *QueryTotalFractionalBalancesRequest, opts ...grpc.CallOption) (*QueryTotalFractionalBalancesResponse, error)
	// Remainder returns the amount backed by the reserve, but not yet owned by
	// any account, i.e. not in circulation.
	Remainder(ctx context.Context, in *QueryRemainderRequest, opts ...grpc.CallOption) (*QueryRemainderResponse, error)
	// FractionalBalance returns only the fractional balance of an address. This
	// does not include any integer balance.
	FractionalBalance(ctx context.Context, in *QueryFractionalBalanceRequest, opts ...grpc.CallOption) (*QueryFractionalBalanceResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) TotalFractionalBalances(ctx context.Context, in *QueryTotalFractionalBalancesRequest, opts ...grpc.CallOption) (*QueryTotalFractionalBalancesResponse, error) {
	out := new(QueryTotalFractionalBalancesResponse)
	err := c.cc.Invoke(ctx, "/zgc.precisebank.v1.Query/TotalFractionalBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Remainder(ctx context.Context, in *QueryRemainderRequest, opts ...grpc.CallOption) (*QueryRemainderResponse, error) {
	out := new(QueryRemainderResponse)
	err := c.cc.Invoke(ctx, "/zgc.precisebank.v1.Query/Remainder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FractionalBalance(ctx context.Context, in *QueryFractionalBalanceRequest, opts ...grpc.CallOption) (*QueryFractionalBalanceResponse, error) {
	out := new(QueryFractionalBalanceResponse)
	err := c.cc.Invoke(ctx, "/zgc.precisebank.v1.Query/FractionalBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// TotalFractionalBalances returns the total sum of all fractional balances
	// managed by the precisebank module.
	TotalFractionalBalances(context.Context, *QueryTotalFractionalBalancesRequest) (*QueryTotalFractionalBalancesResponse, error)
	// Remainder returns the amount backed by the reserve, but not yet owned by
	// any account, i.e. not in circulation.
	Remainder(context.Context, *QueryRemainderRequest) (*QueryRemainderResponse, error)
	// FractionalBalance returns only the fractional balance of an address. This
	// does not include any integer balance.
	FractionalBalance(context.Context, *QueryFractionalBalanceRequest) (*QueryFractionalBalanceResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) TotalFractionalBalances(ctx context.Context, req *QueryTotalFractionalBalancesRequest) (*QueryTotalFractionalBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalFractionalBalances not implemented")
}
func (*UnimplementedQueryServer) Remainder(ctx context.Context, req *QueryRemainderRequest) (*QueryRemainderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remainder not implemented")
}
func (*UnimplementedQueryServer) FractionalBalance(ctx context.Context, req *QueryFractionalBalanceRequest) (*QueryFractionalBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FractionalBalance not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_TotalFractionalBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalFractionalBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalFractionalBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zgc.precisebank.v1.Query/TotalFractionalBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalFractionalBalances(ctx, req.(*QueryTotalFractionalBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Remainder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRemainderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Remainder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zgc.precisebank.v1.Query/Remainder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Remainder(ctx, req.(*QueryRemainderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FractionalBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFractionalBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FractionalBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zgc.precisebank.v1.Query/FractionalBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FractionalBalance(ctx, req.(*QueryFractionalBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zgc.precisebank.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TotalFractionalBalances",
			Handler:    _Query_TotalFractionalBalances_Handler,
		},
		{
			MethodName: "Remainder",
			Handler:    _Query_Remainder_Handler,
		},
		{
			MethodName: "FractionalBalance",
			Handler:    _Query_FractionalBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zgc/precisebank/v1/query.proto",
}

func (m *QueryTotalFractionalBalancesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalFractionalBalancesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalFractionalBalancesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryTotalFractionalBalancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalFractionalBalancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalFractionalBalancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Total.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryRemainderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRemainderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRemainderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryRemainderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRemainderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRemainderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Remainder.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryFractionalBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFractionalBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFractionalBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFractionalBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFractionalBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFractionalBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.FractionalBalance.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryTotalFractionalBalancesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryTotalFractionalBalancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Total.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryRemainderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryRemainderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Remainder.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryFractionalBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFractionalBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.FractionalBalance.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryTotalFractionalBalancesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTotalFractionalBalancesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalFractionalBalancesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryTotalFractionalBalancesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTotalFractionalBalancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalFractionalBalancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
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
			if err := m.Total.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryRemainderRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRemainderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRemainderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryRemainderResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRemainderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRemainderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Remainder", wireType)
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
			if err := m.Remainder.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryFractionalBalanceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFractionalBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFractionalBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryFractionalBalanceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFractionalBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFractionalBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FractionalBalance", wireType)
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
			if err := m.FractionalBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
