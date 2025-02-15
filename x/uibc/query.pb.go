// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/uibc/v1/query.proto

package uibc

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryParams defines the request structure for the Params gRPC service
// handler.
type QueryParams struct {
}

func (m *QueryParams) Reset()         { *m = QueryParams{} }
func (m *QueryParams) String() string { return proto.CompactTextString(m) }
func (*QueryParams) ProtoMessage()    {}
func (*QueryParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ca7e17b0958935d, []int{0}
}
func (m *QueryParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParams.Merge(m, src)
}
func (m *QueryParams) XXX_Size() int {
	return m.Size()
}
func (m *QueryParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParams.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParams proto.InternalMessageInfo

// QueryParamsResponse defines the response structure for the Params gRPC
// service handler.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ca7e17b0958935d, []int{1}
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

// QueryOutflow defines request type for query the quota of denoms
type QueryOutflows struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryOutflows) Reset()         { *m = QueryOutflows{} }
func (m *QueryOutflows) String() string { return proto.CompactTextString(m) }
func (*QueryOutflows) ProtoMessage()    {}
func (*QueryOutflows) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ca7e17b0958935d, []int{2}
}
func (m *QueryOutflows) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOutflows) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOutflows.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOutflows) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOutflows.Merge(m, src)
}
func (m *QueryOutflows) XXX_Size() int {
	return m.Size()
}
func (m *QueryOutflows) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOutflows.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOutflows proto.InternalMessageInfo

// QueryOutflowResponse defines response type of Query/Outflow
type QueryOutflowsResponse struct {
	Amount github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"amount"`
}

func (m *QueryOutflowsResponse) Reset()         { *m = QueryOutflowsResponse{} }
func (m *QueryOutflowsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryOutflowsResponse) ProtoMessage()    {}
func (*QueryOutflowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ca7e17b0958935d, []int{3}
}
func (m *QueryOutflowsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOutflowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOutflowsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOutflowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOutflowsResponse.Merge(m, src)
}
func (m *QueryOutflowsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryOutflowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOutflowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOutflowsResponse proto.InternalMessageInfo

// QueryOutflow defines request type for query the quota of denoms
type QueryAllOutflows struct {
}

func (m *QueryAllOutflows) Reset()         { *m = QueryAllOutflows{} }
func (m *QueryAllOutflows) String() string { return proto.CompactTextString(m) }
func (*QueryAllOutflows) ProtoMessage()    {}
func (*QueryAllOutflows) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ca7e17b0958935d, []int{4}
}
func (m *QueryAllOutflows) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllOutflows) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllOutflows.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllOutflows) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllOutflows.Merge(m, src)
}
func (m *QueryAllOutflows) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllOutflows) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllOutflows.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllOutflows proto.InternalMessageInfo

// QueryOutflowResponse defines response type of Query/Outflow
type QueryAllOutflowsResponse struct {
	Outflows github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=outflows,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"outflows"`
}

func (m *QueryAllOutflowsResponse) Reset()         { *m = QueryAllOutflowsResponse{} }
func (m *QueryAllOutflowsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllOutflowsResponse) ProtoMessage()    {}
func (*QueryAllOutflowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ca7e17b0958935d, []int{5}
}
func (m *QueryAllOutflowsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllOutflowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllOutflowsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllOutflowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllOutflowsResponse.Merge(m, src)
}
func (m *QueryAllOutflowsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllOutflowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllOutflowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllOutflowsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryParams)(nil), "umee.uibc.v1.QueryParams")
	proto.RegisterType((*QueryParamsResponse)(nil), "umee.uibc.v1.QueryParamsResponse")
	proto.RegisterType((*QueryOutflows)(nil), "umee.uibc.v1.QueryOutflows")
	proto.RegisterType((*QueryOutflowsResponse)(nil), "umee.uibc.v1.QueryOutflowsResponse")
	proto.RegisterType((*QueryAllOutflows)(nil), "umee.uibc.v1.QueryAllOutflows")
	proto.RegisterType((*QueryAllOutflowsResponse)(nil), "umee.uibc.v1.QueryAllOutflowsResponse")
}

func init() { proto.RegisterFile("umee/uibc/v1/query.proto", fileDescriptor_2ca7e17b0958935d) }

var fileDescriptor_2ca7e17b0958935d = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0xc1, 0xaa, 0xe1, 0x32, 0x09, 0x99, 0x32, 0x75, 0xa1, 0x72, 0x47, 0x10, 0xd3,
	0x24, 0xd4, 0x58, 0xed, 0x38, 0x72, 0xa1, 0xf4, 0xc2, 0x09, 0xa8, 0x38, 0x71, 0x41, 0x6e, 0x6a,
	0x4a, 0xb4, 0x24, 0x6f, 0xa8, 0x9d, 0x8e, 0x5e, 0xb9, 0x71, 0x43, 0xe2, 0x5b, 0x70, 0xe6, 0x43,
	0xf4, 0x38, 0x81, 0x84, 0x10, 0x87, 0x01, 0x2d, 0x1f, 0x04, 0xf9, 0x4f, 0xa2, 0x64, 0x1a, 0xd3,
	0x4e, 0xb5, 0xdf, 0xe7, 0xed, 0xf3, 0x7b, 0xfd, 0xd8, 0x41, 0xad, 0x2c, 0xe6, 0x9c, 0x66, 0xe1,
	0x38, 0xa0, 0xf3, 0x1e, 0x7d, 0x9b, 0xf1, 0xd9, 0xc2, 0x4f, 0x67, 0x20, 0x01, 0x5f, 0x57, 0x8a,
	0xaf, 0x14, 0x7f, 0xde, 0x73, 0xdb, 0x53, 0x80, 0x69, 0xc4, 0x29, 0x4b, 0x43, 0xca, 0x92, 0x04,
	0x24, 0x93, 0x21, 0x24, 0xc2, 0xf4, 0xba, 0xcd, 0x29, 0x4c, 0x41, 0x2f, 0xa9, 0x5a, 0xd9, 0xea,
	0x6e, 0x00, 0x22, 0x06, 0xf1, 0xca, 0x08, 0x66, 0x63, 0xa5, 0xb3, 0x58, 0x90, 0xcc, 0x2a, 0xc4,
	0xf4, 0xd1, 0x31, 0x13, 0x9c, 0xce, 0x7b, 0x63, 0x2e, 0x59, 0x8f, 0x06, 0x10, 0x26, 0x46, 0xf7,
	0xb6, 0x51, 0xe3, 0xb9, 0x9a, 0xf2, 0x19, 0x9b, 0xb1, 0x58, 0x78, 0x4f, 0xd0, 0xcd, 0xd2, 0x76,
	0xc4, 0x45, 0x0a, 0x89, 0xe0, 0xb8, 0x8f, 0xea, 0xa9, 0xae, 0xb4, 0x9c, 0x3d, 0xe7, 0xa0, 0xd1,
	0x6f, 0xfa, 0xe5, 0xd3, 0xf8, 0xa6, 0x7b, 0x70, 0x75, 0x79, 0xda, 0xa9, 0x8d, 0x6c, 0xa7, 0x77,
	0x0f, 0x6d, 0x6b, 0xab, 0xa7, 0x99, 0x7c, 0x1d, 0xc1, 0xb1, 0xc0, 0x4d, 0xb4, 0x39, 0xe1, 0x09,
	0xc4, 0xda, 0xe3, 0xda, 0xc8, 0x6c, 0xbc, 0x18, 0xdd, 0xaa, 0xb4, 0x15, 0xcc, 0x17, 0xa8, 0xce,
	0x62, 0xc8, 0x12, 0x69, 0xfa, 0x07, 0x0f, 0x95, 0xfb, 0xcf, 0xd3, 0xce, 0xfe, 0x34, 0x94, 0x6f,
	0xb2, 0xb1, 0x1f, 0x40, 0x6c, 0x43, 0xb0, 0x3f, 0x5d, 0x31, 0x39, 0xa2, 0x72, 0x91, 0x72, 0xe1,
	0x0f, 0x79, 0xf0, 0xf5, 0x4b, 0x17, 0xd9, 0x8c, 0x86, 0x3c, 0x18, 0x59, 0x2f, 0x0f, 0xa3, 0x1b,
	0x1a, 0xf7, 0x28, 0x8a, 0x72, 0xa2, 0xf7, 0xc1, 0x41, 0xad, 0xb3, 0xc5, 0x62, 0x8c, 0x18, 0x6d,
	0x81, 0xad, 0xb5, 0x9c, 0xbd, 0x2b, 0x07, 0x8d, 0x7e, 0xdb, 0xb7, 0xbe, 0x2a, 0x53, 0xdf, 0x66,
	0xaa, 0x20, 0x8f, 0x21, 0x4c, 0x06, 0x87, 0x6a, 0xcc, 0xcf, 0xbf, 0x3a, 0xf7, 0x2f, 0x37, 0xa6,
	0xfa, 0x8f, 0x18, 0x15, 0x88, 0xfe, 0xf7, 0x0d, 0xb4, 0xa9, 0x67, 0xc1, 0x13, 0x54, 0x37, 0xb9,
	0xe2, 0xdd, 0x6a, 0xda, 0xa5, 0x0b, 0x72, 0xef, 0xfc, 0x57, 0xca, 0x0f, 0xe0, 0xb5, 0xdf, 0x7f,
	0xfb, 0xfb, 0x69, 0x63, 0x07, 0x37, 0x69, 0xe5, 0x91, 0x98, 0x5b, 0xc2, 0x11, 0xda, 0x2a, 0x2e,
	0xe8, 0xf6, 0x39, 0x66, 0xb9, 0xe8, 0xde, 0xbd, 0x40, 0x2c, 0x58, 0x44, 0xb3, 0x5a, 0x78, 0xa7,
	0xca, 0xca, 0x4f, 0x87, 0x17, 0xa8, 0x51, 0xca, 0x18, 0x93, 0x73, 0x3c, 0x4b, 0xba, 0xbb, 0x7f,
	0xb1, 0x5e, 0x60, 0x3d, 0x8d, 0x6d, 0x63, 0xb7, 0x8a, 0x65, 0x51, 0xd4, 0xcd, 0xd1, 0x83, 0xe1,
	0xf2, 0x0f, 0xa9, 0x2d, 0x57, 0xc4, 0x39, 0x59, 0x11, 0xe7, 0xf7, 0x8a, 0x38, 0x1f, 0xd7, 0xa4,
	0x76, 0xb2, 0x26, 0xb5, 0x1f, 0x6b, 0x52, 0x7b, 0x59, 0x7e, 0x54, 0xca, 0xa3, 0x9b, 0x70, 0x79,
	0x0c, 0xb3, 0x23, 0x63, 0x38, 0x7f, 0x40, 0xdf, 0x69, 0xd7, 0x71, 0x5d, 0x7f, 0x35, 0x87, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x38, 0xf4, 0x22, 0xe8, 0x03, 0x00, 0x00,
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
	// Params queries the parameters of the x/uibc module.
	Params(ctx context.Context, in *QueryParams, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Outflow returns IBC denom outflows in the current quota period.
	// If denom is not specified, returns sum of all registered outflows.
	Outflows(ctx context.Context, in *QueryOutflows, opts ...grpc.CallOption) (*QueryOutflowsResponse, error)
	// AllOutflow returns outflows for each denom in the current quota period.
	AllOutflows(ctx context.Context, in *QueryAllOutflows, opts ...grpc.CallOption) (*QueryAllOutflowsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParams, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/umee.uibc.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Outflows(ctx context.Context, in *QueryOutflows, opts ...grpc.CallOption) (*QueryOutflowsResponse, error) {
	out := new(QueryOutflowsResponse)
	err := c.cc.Invoke(ctx, "/umee.uibc.v1.Query/Outflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllOutflows(ctx context.Context, in *QueryAllOutflows, opts ...grpc.CallOption) (*QueryAllOutflowsResponse, error) {
	out := new(QueryAllOutflowsResponse)
	err := c.cc.Invoke(ctx, "/umee.uibc.v1.Query/AllOutflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries the parameters of the x/uibc module.
	Params(context.Context, *QueryParams) (*QueryParamsResponse, error)
	// Outflow returns IBC denom outflows in the current quota period.
	// If denom is not specified, returns sum of all registered outflows.
	Outflows(context.Context, *QueryOutflows) (*QueryOutflowsResponse, error)
	// AllOutflow returns outflows for each denom in the current quota period.
	AllOutflows(context.Context, *QueryAllOutflows) (*QueryAllOutflowsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParams) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Outflows(ctx context.Context, req *QueryOutflows) (*QueryOutflowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Outflows not implemented")
}
func (*UnimplementedQueryServer) AllOutflows(ctx context.Context, req *QueryAllOutflows) (*QueryAllOutflowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllOutflows not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.uibc.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Outflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOutflows)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Outflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.uibc.v1.Query/Outflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Outflows(ctx, req.(*QueryOutflows))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllOutflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllOutflows)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllOutflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.uibc.v1.Query/AllOutflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllOutflows(ctx, req.(*QueryAllOutflows))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "umee.uibc.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Outflows",
			Handler:    _Query_Outflows_Handler,
		},
		{
			MethodName: "AllOutflows",
			Handler:    _Query_AllOutflows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umee/uibc/v1/query.proto",
}

func (m *QueryParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryOutflows) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOutflows) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOutflows) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryOutflowsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOutflowsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOutflowsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllOutflows) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllOutflows) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllOutflows) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAllOutflowsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllOutflowsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllOutflowsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Outflows) > 0 {
		for iNdEx := len(m.Outflows) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Outflows[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryParams) Size() (n int) {
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

func (m *QueryOutflows) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryOutflowsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllOutflows) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAllOutflowsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Outflows) > 0 {
		for _, e := range m.Outflows {
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
func (m *QueryParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParams: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryOutflows) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryOutflows: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOutflows: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
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
func (m *QueryOutflowsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryOutflowsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOutflowsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllOutflows) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllOutflows: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllOutflows: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllOutflowsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllOutflowsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllOutflowsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outflows", wireType)
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
			m.Outflows = append(m.Outflows, types.DecCoin{})
			if err := m.Outflows[len(m.Outflows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
