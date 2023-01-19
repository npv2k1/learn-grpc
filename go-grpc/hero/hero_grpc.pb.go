// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: hero/hero.proto

package hero

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

// HeroesServiceClient is the client API for HeroesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeroesServiceClient interface {
	FindOne(ctx context.Context, in *HeroById, opts ...grpc.CallOption) (*Hero, error)
}

type heroesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeroesServiceClient(cc grpc.ClientConnInterface) HeroesServiceClient {
	return &heroesServiceClient{cc}
}

func (c *heroesServiceClient) FindOne(ctx context.Context, in *HeroById, opts ...grpc.CallOption) (*Hero, error) {
	out := new(Hero)
	err := c.cc.Invoke(ctx, "/hero.HeroesService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeroesServiceServer is the server API for HeroesService service.
// All implementations must embed UnimplementedHeroesServiceServer
// for forward compatibility
type HeroesServiceServer interface {
	FindOne(context.Context, *HeroById) (*Hero, error)
	mustEmbedUnimplementedHeroesServiceServer()
}

// UnimplementedHeroesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeroesServiceServer struct {
}

func (UnimplementedHeroesServiceServer) FindOne(context.Context, *HeroById) (*Hero, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedHeroesServiceServer) mustEmbedUnimplementedHeroesServiceServer() {}

// UnsafeHeroesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeroesServiceServer will
// result in compilation errors.
type UnsafeHeroesServiceServer interface {
	mustEmbedUnimplementedHeroesServiceServer()
}

func RegisterHeroesServiceServer(s grpc.ServiceRegistrar, srv HeroesServiceServer) {
	s.RegisterService(&HeroesService_ServiceDesc, srv)
}

func _HeroesService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeroById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroesServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hero.HeroesService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroesServiceServer).FindOne(ctx, req.(*HeroById))
	}
	return interceptor(ctx, in, info, handler)
}

// HeroesService_ServiceDesc is the grpc.ServiceDesc for HeroesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeroesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hero.HeroesService",
	HandlerType: (*HeroesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOne",
			Handler:    _HeroesService_FindOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hero/hero.proto",
}