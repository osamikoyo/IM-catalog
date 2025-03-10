// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: pkg/proto/catalog.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CatalogService_Add_FullMethodName     = "/CatalogService/Add"
	CatalogService_Update_FullMethodName  = "/CatalogService/Update"
	CatalogService_GetOne_FullMethodName  = "/CatalogService/GetOne"
	CatalogService_GetMore_FullMethodName = "/CatalogService/GetMore"
	CatalogService_GetAll_FullMethodName  = "/CatalogService/GetAll"
	CatalogService_Delete_FullMethodName  = "/CatalogService/Delete"
)

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogServiceClient interface {
	Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Response, error)
	GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneResp, error)
	GetMore(ctx context.Context, in *GetMoreReq, opts ...grpc.CallOption) (*GetMoreResp, error)
	GetAll(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (*GetAllResp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Response, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, CatalogService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, CatalogService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOneResp)
	err := c.cc.Invoke(ctx, CatalogService_GetOne_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetMore(ctx context.Context, in *GetMoreReq, opts ...grpc.CallOption) (*GetMoreResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMoreResp)
	err := c.cc.Invoke(ctx, CatalogService_GetMore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetAll(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (*GetAllResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllResp)
	err := c.cc.Invoke(ctx, CatalogService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, CatalogService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
// All implementations must embed UnimplementedCatalogServiceServer
// for forward compatibility.
type CatalogServiceServer interface {
	Add(context.Context, *AddReq) (*Response, error)
	Update(context.Context, *UpdateReq) (*Response, error)
	GetOne(context.Context, *GetOneReq) (*GetOneResp, error)
	GetMore(context.Context, *GetMoreReq) (*GetMoreResp, error)
	GetAll(context.Context, *GetAllReq) (*GetAllResp, error)
	Delete(context.Context, *DeleteReq) (*Response, error)
	mustEmbedUnimplementedCatalogServiceServer()
}

// UnimplementedCatalogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCatalogServiceServer struct{}

func (UnimplementedCatalogServiceServer) Add(context.Context, *AddReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCatalogServiceServer) Update(context.Context, *UpdateReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCatalogServiceServer) GetOne(context.Context, *GetOneReq) (*GetOneResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedCatalogServiceServer) GetMore(context.Context, *GetMoreReq) (*GetMoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMore not implemented")
}
func (UnimplementedCatalogServiceServer) GetAll(context.Context, *GetAllReq) (*GetAllResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCatalogServiceServer) Delete(context.Context, *DeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCatalogServiceServer) mustEmbedUnimplementedCatalogServiceServer() {}
func (UnimplementedCatalogServiceServer) testEmbeddedByValue()                        {}

// UnsafeCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServiceServer will
// result in compilation errors.
type UnsafeCatalogServiceServer interface {
	mustEmbedUnimplementedCatalogServiceServer()
}

func RegisterCatalogServiceServer(s grpc.ServiceRegistrar, srv CatalogServiceServer) {
	// If the following call pancis, it indicates UnimplementedCatalogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CatalogService_ServiceDesc, srv)
}

func _CatalogService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Add(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetOne(ctx, req.(*GetOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetMore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetMore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetMore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetMore(ctx, req.(*GetMoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetAll(ctx, req.(*GetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogService_ServiceDesc is the grpc.ServiceDesc for CatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CatalogService_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CatalogService_Update_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _CatalogService_GetOne_Handler,
		},
		{
			MethodName: "GetMore",
			Handler:    _CatalogService_GetMore_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CatalogService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CatalogService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/catalog.proto",
}
