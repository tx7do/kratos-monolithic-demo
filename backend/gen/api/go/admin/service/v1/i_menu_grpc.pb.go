// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: admin/service/v1/i_menu.proto

package servicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-monolithic-demo/gen/api/go/common/pagination"
	v1 "kratos-monolithic-demo/gen/api/go/system/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MenuService_ListMenu_FullMethodName   = "/admin.service.v1.MenuService/ListMenu"
	MenuService_GetMenu_FullMethodName    = "/admin.service.v1.MenuService/GetMenu"
	MenuService_CreateMenu_FullMethodName = "/admin.service.v1.MenuService/CreateMenu"
	MenuService_UpdateMenu_FullMethodName = "/admin.service.v1.MenuService/UpdateMenu"
	MenuService_DeleteMenu_FullMethodName = "/admin.service.v1.MenuService/DeleteMenu"
)

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuServiceClient interface {
	// 查询菜单列表
	ListMenu(ctx context.Context, in *pagination.PagingRequest, opts ...grpc.CallOption) (*v1.ListMenuResponse, error)
	// 查询菜单详情
	GetMenu(ctx context.Context, in *v1.GetMenuRequest, opts ...grpc.CallOption) (*v1.Menu, error)
	// 创建菜单
	CreateMenu(ctx context.Context, in *v1.CreateMenuRequest, opts ...grpc.CallOption) (*v1.Menu, error)
	// 更新菜单
	UpdateMenu(ctx context.Context, in *v1.UpdateMenuRequest, opts ...grpc.CallOption) (*v1.Menu, error)
	// 删除菜单
	DeleteMenu(ctx context.Context, in *v1.DeleteMenuRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type menuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuServiceClient(cc grpc.ClientConnInterface) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) ListMenu(ctx context.Context, in *pagination.PagingRequest, opts ...grpc.CallOption) (*v1.ListMenuResponse, error) {
	out := new(v1.ListMenuResponse)
	err := c.cc.Invoke(ctx, MenuService_ListMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) GetMenu(ctx context.Context, in *v1.GetMenuRequest, opts ...grpc.CallOption) (*v1.Menu, error) {
	out := new(v1.Menu)
	err := c.cc.Invoke(ctx, MenuService_GetMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) CreateMenu(ctx context.Context, in *v1.CreateMenuRequest, opts ...grpc.CallOption) (*v1.Menu, error) {
	out := new(v1.Menu)
	err := c.cc.Invoke(ctx, MenuService_CreateMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) UpdateMenu(ctx context.Context, in *v1.UpdateMenuRequest, opts ...grpc.CallOption) (*v1.Menu, error) {
	out := new(v1.Menu)
	err := c.cc.Invoke(ctx, MenuService_UpdateMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) DeleteMenu(ctx context.Context, in *v1.DeleteMenuRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MenuService_DeleteMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServiceServer is the server API for MenuService service.
// All implementations must embed UnimplementedMenuServiceServer
// for forward compatibility
type MenuServiceServer interface {
	// 查询菜单列表
	ListMenu(context.Context, *pagination.PagingRequest) (*v1.ListMenuResponse, error)
	// 查询菜单详情
	GetMenu(context.Context, *v1.GetMenuRequest) (*v1.Menu, error)
	// 创建菜单
	CreateMenu(context.Context, *v1.CreateMenuRequest) (*v1.Menu, error)
	// 更新菜单
	UpdateMenu(context.Context, *v1.UpdateMenuRequest) (*v1.Menu, error)
	// 删除菜单
	DeleteMenu(context.Context, *v1.DeleteMenuRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMenuServiceServer()
}

// UnimplementedMenuServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMenuServiceServer struct {
}

func (UnimplementedMenuServiceServer) ListMenu(context.Context, *pagination.PagingRequest) (*v1.ListMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMenu not implemented")
}
func (UnimplementedMenuServiceServer) GetMenu(context.Context, *v1.GetMenuRequest) (*v1.Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (UnimplementedMenuServiceServer) CreateMenu(context.Context, *v1.CreateMenuRequest) (*v1.Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenu not implemented")
}
func (UnimplementedMenuServiceServer) UpdateMenu(context.Context, *v1.UpdateMenuRequest) (*v1.Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenu not implemented")
}
func (UnimplementedMenuServiceServer) DeleteMenu(context.Context, *v1.DeleteMenuRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMenu not implemented")
}
func (UnimplementedMenuServiceServer) mustEmbedUnimplementedMenuServiceServer() {}

// UnsafeMenuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MenuServiceServer will
// result in compilation errors.
type UnsafeMenuServiceServer interface {
	mustEmbedUnimplementedMenuServiceServer()
}

func RegisterMenuServiceServer(s grpc.ServiceRegistrar, srv MenuServiceServer) {
	s.RegisterService(&MenuService_ServiceDesc, srv)
}

func _MenuService_ListMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).ListMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_ListMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).ListMenu(ctx, req.(*pagination.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_GetMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_GetMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetMenu(ctx, req.(*v1.GetMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_CreateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CreateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_CreateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreateMenu(ctx, req.(*v1.CreateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_UpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).UpdateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_UpdateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).UpdateMenu(ctx, req.(*v1.UpdateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_DeleteMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DeleteMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).DeleteMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_DeleteMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).DeleteMenu(ctx, req.(*v1.DeleteMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MenuService_ServiceDesc is the grpc.ServiceDesc for MenuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MenuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMenu",
			Handler:    _MenuService_ListMenu_Handler,
		},
		{
			MethodName: "GetMenu",
			Handler:    _MenuService_GetMenu_Handler,
		},
		{
			MethodName: "CreateMenu",
			Handler:    _MenuService_CreateMenu_Handler,
		},
		{
			MethodName: "UpdateMenu",
			Handler:    _MenuService_UpdateMenu_Handler,
		},
		{
			MethodName: "DeleteMenu",
			Handler:    _MenuService_DeleteMenu_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/service/v1/i_menu.proto",
}
