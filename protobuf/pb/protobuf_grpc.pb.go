// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: protobuf.proto

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

// RendererClient is the client API for Renderer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RendererClient interface {
	CreatePolygons(ctx context.Context, opts ...grpc.CallOption) (Renderer_CreatePolygonsClient, error)
	CreateTriangle(ctx context.Context, opts ...grpc.CallOption) (Renderer_CreateTriangleClient, error)
	SearchPoint(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Triangle, error)
	GetPolyTriangles(ctx context.Context, in *Polygon, opts ...grpc.CallOption) (Renderer_GetPolyTrianglesClient, error)
	ListOfTriangles(ctx context.Context, in *Void, opts ...grpc.CallOption) (Renderer_ListOfTrianglesClient, error)
	ListOfPolygons(ctx context.Context, in *Void, opts ...grpc.CallOption) (Renderer_ListOfPolygonsClient, error)
}

type rendererClient struct {
	cc grpc.ClientConnInterface
}

func NewRendererClient(cc grpc.ClientConnInterface) RendererClient {
	return &rendererClient{cc}
}

func (c *rendererClient) CreatePolygons(ctx context.Context, opts ...grpc.CallOption) (Renderer_CreatePolygonsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Renderer_ServiceDesc.Streams[0], "/Renderer/CreatePolygons", opts...)
	if err != nil {
		return nil, err
	}
	x := &rendererCreatePolygonsClient{stream}
	return x, nil
}

type Renderer_CreatePolygonsClient interface {
	Send(*Triangle) error
	Recv() (*Polygon, error)
	grpc.ClientStream
}

type rendererCreatePolygonsClient struct {
	grpc.ClientStream
}

func (x *rendererCreatePolygonsClient) Send(m *Triangle) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rendererCreatePolygonsClient) Recv() (*Polygon, error) {
	m := new(Polygon)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rendererClient) CreateTriangle(ctx context.Context, opts ...grpc.CallOption) (Renderer_CreateTriangleClient, error) {
	stream, err := c.cc.NewStream(ctx, &Renderer_ServiceDesc.Streams[1], "/Renderer/CreateTriangle", opts...)
	if err != nil {
		return nil, err
	}
	x := &rendererCreateTriangleClient{stream}
	return x, nil
}

type Renderer_CreateTriangleClient interface {
	Send(*Point) error
	CloseAndRecv() (*Triangle, error)
	grpc.ClientStream
}

type rendererCreateTriangleClient struct {
	grpc.ClientStream
}

func (x *rendererCreateTriangleClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rendererCreateTriangleClient) CloseAndRecv() (*Triangle, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Triangle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rendererClient) SearchPoint(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Triangle, error) {
	out := new(Triangle)
	err := c.cc.Invoke(ctx, "/Renderer/SearchPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rendererClient) GetPolyTriangles(ctx context.Context, in *Polygon, opts ...grpc.CallOption) (Renderer_GetPolyTrianglesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Renderer_ServiceDesc.Streams[2], "/Renderer/GetPolyTriangles", opts...)
	if err != nil {
		return nil, err
	}
	x := &rendererGetPolyTrianglesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Renderer_GetPolyTrianglesClient interface {
	Recv() (*Triangle, error)
	grpc.ClientStream
}

type rendererGetPolyTrianglesClient struct {
	grpc.ClientStream
}

func (x *rendererGetPolyTrianglesClient) Recv() (*Triangle, error) {
	m := new(Triangle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rendererClient) ListOfTriangles(ctx context.Context, in *Void, opts ...grpc.CallOption) (Renderer_ListOfTrianglesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Renderer_ServiceDesc.Streams[3], "/Renderer/ListOfTriangles", opts...)
	if err != nil {
		return nil, err
	}
	x := &rendererListOfTrianglesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Renderer_ListOfTrianglesClient interface {
	Recv() (*Triangle, error)
	grpc.ClientStream
}

type rendererListOfTrianglesClient struct {
	grpc.ClientStream
}

func (x *rendererListOfTrianglesClient) Recv() (*Triangle, error) {
	m := new(Triangle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rendererClient) ListOfPolygons(ctx context.Context, in *Void, opts ...grpc.CallOption) (Renderer_ListOfPolygonsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Renderer_ServiceDesc.Streams[4], "/Renderer/ListOfPolygons", opts...)
	if err != nil {
		return nil, err
	}
	x := &rendererListOfPolygonsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Renderer_ListOfPolygonsClient interface {
	Recv() (*Polygon, error)
	grpc.ClientStream
}

type rendererListOfPolygonsClient struct {
	grpc.ClientStream
}

func (x *rendererListOfPolygonsClient) Recv() (*Polygon, error) {
	m := new(Polygon)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RendererServer is the server API for Renderer service.
// All implementations must embed UnimplementedRendererServer
// for forward compatibility
type RendererServer interface {
	CreatePolygons(Renderer_CreatePolygonsServer) error
	CreateTriangle(Renderer_CreateTriangleServer) error
	SearchPoint(context.Context, *Point) (*Triangle, error)
	GetPolyTriangles(*Polygon, Renderer_GetPolyTrianglesServer) error
	ListOfTriangles(*Void, Renderer_ListOfTrianglesServer) error
	ListOfPolygons(*Void, Renderer_ListOfPolygonsServer) error
	mustEmbedUnimplementedRendererServer()
}

// UnimplementedRendererServer must be embedded to have forward compatible implementations.
type UnimplementedRendererServer struct {
}

func (UnimplementedRendererServer) CreatePolygons(Renderer_CreatePolygonsServer) error {
	return status.Errorf(codes.Unimplemented, "method CreatePolygons not implemented")
}
func (UnimplementedRendererServer) CreateTriangle(Renderer_CreateTriangleServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateTriangle not implemented")
}
func (UnimplementedRendererServer) SearchPoint(context.Context, *Point) (*Triangle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPoint not implemented")
}
func (UnimplementedRendererServer) GetPolyTriangles(*Polygon, Renderer_GetPolyTrianglesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPolyTriangles not implemented")
}
func (UnimplementedRendererServer) ListOfTriangles(*Void, Renderer_ListOfTrianglesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListOfTriangles not implemented")
}
func (UnimplementedRendererServer) ListOfPolygons(*Void, Renderer_ListOfPolygonsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListOfPolygons not implemented")
}
func (UnimplementedRendererServer) mustEmbedUnimplementedRendererServer() {}

// UnsafeRendererServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RendererServer will
// result in compilation errors.
type UnsafeRendererServer interface {
	mustEmbedUnimplementedRendererServer()
}

func RegisterRendererServer(s grpc.ServiceRegistrar, srv RendererServer) {
	s.RegisterService(&Renderer_ServiceDesc, srv)
}

func _Renderer_CreatePolygons_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RendererServer).CreatePolygons(&rendererCreatePolygonsServer{stream})
}

type Renderer_CreatePolygonsServer interface {
	Send(*Polygon) error
	Recv() (*Triangle, error)
	grpc.ServerStream
}

type rendererCreatePolygonsServer struct {
	grpc.ServerStream
}

func (x *rendererCreatePolygonsServer) Send(m *Polygon) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rendererCreatePolygonsServer) Recv() (*Triangle, error) {
	m := new(Triangle)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Renderer_CreateTriangle_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RendererServer).CreateTriangle(&rendererCreateTriangleServer{stream})
}

type Renderer_CreateTriangleServer interface {
	SendAndClose(*Triangle) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type rendererCreateTriangleServer struct {
	grpc.ServerStream
}

func (x *rendererCreateTriangleServer) SendAndClose(m *Triangle) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rendererCreateTriangleServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Renderer_SearchPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RendererServer).SearchPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Renderer/SearchPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RendererServer).SearchPoint(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _Renderer_GetPolyTriangles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Polygon)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RendererServer).GetPolyTriangles(m, &rendererGetPolyTrianglesServer{stream})
}

type Renderer_GetPolyTrianglesServer interface {
	Send(*Triangle) error
	grpc.ServerStream
}

type rendererGetPolyTrianglesServer struct {
	grpc.ServerStream
}

func (x *rendererGetPolyTrianglesServer) Send(m *Triangle) error {
	return x.ServerStream.SendMsg(m)
}

func _Renderer_ListOfTriangles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RendererServer).ListOfTriangles(m, &rendererListOfTrianglesServer{stream})
}

type Renderer_ListOfTrianglesServer interface {
	Send(*Triangle) error
	grpc.ServerStream
}

type rendererListOfTrianglesServer struct {
	grpc.ServerStream
}

func (x *rendererListOfTrianglesServer) Send(m *Triangle) error {
	return x.ServerStream.SendMsg(m)
}

func _Renderer_ListOfPolygons_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RendererServer).ListOfPolygons(m, &rendererListOfPolygonsServer{stream})
}

type Renderer_ListOfPolygonsServer interface {
	Send(*Polygon) error
	grpc.ServerStream
}

type rendererListOfPolygonsServer struct {
	grpc.ServerStream
}

func (x *rendererListOfPolygonsServer) Send(m *Polygon) error {
	return x.ServerStream.SendMsg(m)
}

// Renderer_ServiceDesc is the grpc.ServiceDesc for Renderer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Renderer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Renderer",
	HandlerType: (*RendererServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchPoint",
			Handler:    _Renderer_SearchPoint_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreatePolygons",
			Handler:       _Renderer_CreatePolygons_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateTriangle",
			Handler:       _Renderer_CreateTriangle_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetPolyTriangles",
			Handler:       _Renderer_GetPolyTriangles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListOfTriangles",
			Handler:       _Renderer_ListOfTriangles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListOfPolygons",
			Handler:       _Renderer_ListOfPolygons_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobuf.proto",
}