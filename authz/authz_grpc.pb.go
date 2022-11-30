// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: github.com/openconfig/gnsi/authz/authz.proto

package authz

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

// AuthzClient is the client API for Authz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthzClient interface {
	Rotate(ctx context.Context, opts ...grpc.CallOption) (Authz_RotateClient, error)
	Probe(ctx context.Context, in *ProbeRequest, opts ...grpc.CallOption) (*ProbeResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type authzClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthzClient(cc grpc.ClientConnInterface) AuthzClient {
	return &authzClient{cc}
}

func (c *authzClient) Rotate(ctx context.Context, opts ...grpc.CallOption) (Authz_RotateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Authz_ServiceDesc.Streams[0], "/gnsi.authz.Authz/Rotate", opts...)
	if err != nil {
		return nil, err
	}
	x := &authzRotateClient{stream}
	return x, nil
}

type Authz_RotateClient interface {
	Send(*RotateAuthzRequest) error
	Recv() (*RotateAuthzResponse, error)
	grpc.ClientStream
}

type authzRotateClient struct {
	grpc.ClientStream
}

func (x *authzRotateClient) Send(m *RotateAuthzRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *authzRotateClient) Recv() (*RotateAuthzResponse, error) {
	m := new(RotateAuthzResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authzClient) Probe(ctx context.Context, in *ProbeRequest, opts ...grpc.CallOption) (*ProbeResponse, error) {
	out := new(ProbeResponse)
	err := c.cc.Invoke(ctx, "/gnsi.authz.Authz/Probe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authzClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/gnsi.authz.Authz/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzServer is the server API for Authz service.
// All implementations must embed UnimplementedAuthzServer
// for forward compatibility
type AuthzServer interface {
	Rotate(Authz_RotateServer) error
	Probe(context.Context, *ProbeRequest) (*ProbeResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedAuthzServer()
}

// UnimplementedAuthzServer must be embedded to have forward compatible implementations.
type UnimplementedAuthzServer struct {
}

func (UnimplementedAuthzServer) Rotate(Authz_RotateServer) error {
	return status.Errorf(codes.Unimplemented, "method Rotate not implemented")
}
func (UnimplementedAuthzServer) Probe(context.Context, *ProbeRequest) (*ProbeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Probe not implemented")
}
func (UnimplementedAuthzServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAuthzServer) mustEmbedUnimplementedAuthzServer() {}

// UnsafeAuthzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthzServer will
// result in compilation errors.
type UnsafeAuthzServer interface {
	mustEmbedUnimplementedAuthzServer()
}

func RegisterAuthzServer(s grpc.ServiceRegistrar, srv AuthzServer) {
	s.RegisterService(&Authz_ServiceDesc, srv)
}

func _Authz_Rotate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AuthzServer).Rotate(&authzRotateServer{stream})
}

type Authz_RotateServer interface {
	Send(*RotateAuthzResponse) error
	Recv() (*RotateAuthzRequest, error)
	grpc.ServerStream
}

type authzRotateServer struct {
	grpc.ServerStream
}

func (x *authzRotateServer) Send(m *RotateAuthzResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *authzRotateServer) Recv() (*RotateAuthzRequest, error) {
	m := new(RotateAuthzRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Authz_Probe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProbeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServer).Probe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnsi.authz.Authz/Probe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServer).Probe(ctx, req.(*ProbeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authz_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnsi.authz.Authz/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authz_ServiceDesc is the grpc.ServiceDesc for Authz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnsi.authz.Authz",
	HandlerType: (*AuthzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Probe",
			Handler:    _Authz_Probe_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Authz_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Rotate",
			Handler:       _Authz_Rotate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnsi/authz/authz.proto",
}