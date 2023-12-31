// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: mygrpc/proto/simple_parser.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AddressParserServiceClient is the client API for AddressParserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressParserServiceClient interface {
	// 簡易地址解析
	SimpleParser(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error)
}

type addressParserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressParserServiceClient(cc grpc.ClientConnInterface) AddressParserServiceClient {
	return &addressParserServiceClient{cc}
}

func (c *addressParserServiceClient) SimpleParser(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error) {
	out := new(ParserResponse)
	err := c.cc.Invoke(ctx, "/address.AddressParserService/SimpleParser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressParserServiceServer is the server API for AddressParserService service.
// All implementations must embed UnimplementedAddressParserServiceServer
// for forward compatibility
type AddressParserServiceServer interface {
	// 簡易地址解析
	SimpleParser(context.Context, *ParserRequest) (*ParserResponse, error)
	mustEmbedUnimplementedAddressParserServiceServer()
}

// UnimplementedAddressParserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddressParserServiceServer struct {
}

func (UnimplementedAddressParserServiceServer) SimpleParser(context.Context, *ParserRequest) (*ParserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimpleParser not implemented")
}
func (UnimplementedAddressParserServiceServer) mustEmbedUnimplementedAddressParserServiceServer() {}

// UnsafeAddressParserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressParserServiceServer will
// result in compilation errors.
type UnsafeAddressParserServiceServer interface {
	mustEmbedUnimplementedAddressParserServiceServer()
}

func RegisterAddressParserServiceServer(s grpc.ServiceRegistrar, srv AddressParserServiceServer) {
	s.RegisterService(&AddressParserService_ServiceDesc, srv)
}

func _AddressParserService_SimpleParser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressParserServiceServer).SimpleParser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/address.AddressParserService/SimpleParser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressParserServiceServer).SimpleParser(ctx, req.(*ParserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressParserService_ServiceDesc is the grpc.ServiceDesc for AddressParserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressParserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "address.AddressParserService",
	HandlerType: (*AddressParserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SimpleParser",
			Handler:    _AddressParserService_SimpleParser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mygrpc/proto/simple_parser.proto",
}
