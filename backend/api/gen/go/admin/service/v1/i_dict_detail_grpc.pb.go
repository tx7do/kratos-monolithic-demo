// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: admin/service/v1/i_dict_detail.proto

package servicev1

import (
	context "context"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-monolithic-demo/api/gen/go/system/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DictDetailService_ListDictDetail_FullMethodName   = "/admin.service.v1.DictDetailService/ListDictDetail"
	DictDetailService_GetDictDetail_FullMethodName    = "/admin.service.v1.DictDetailService/GetDictDetail"
	DictDetailService_CreateDictDetail_FullMethodName = "/admin.service.v1.DictDetailService/CreateDictDetail"
	DictDetailService_UpdateDictDetail_FullMethodName = "/admin.service.v1.DictDetailService/UpdateDictDetail"
	DictDetailService_DeleteDictDetail_FullMethodName = "/admin.service.v1.DictDetailService/DeleteDictDetail"
)

// DictDetailServiceClient is the client API for DictDetailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DictDetailServiceClient interface {
	// 查询字典详情列表
	ListDictDetail(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListDictDetailResponse, error)
	// 查询字典详情
	GetDictDetail(ctx context.Context, in *v11.GetDictDetailRequest, opts ...grpc.CallOption) (*v11.DictDetail, error)
	// 创建字典详情
	CreateDictDetail(ctx context.Context, in *v11.CreateDictDetailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新字典详情
	UpdateDictDetail(ctx context.Context, in *v11.UpdateDictDetailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除字典详情
	DeleteDictDetail(ctx context.Context, in *v11.DeleteDictDetailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dictDetailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDictDetailServiceClient(cc grpc.ClientConnInterface) DictDetailServiceClient {
	return &dictDetailServiceClient{cc}
}

func (c *dictDetailServiceClient) ListDictDetail(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListDictDetailResponse, error) {
	out := new(v11.ListDictDetailResponse)
	err := c.cc.Invoke(ctx, DictDetailService_ListDictDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictDetailServiceClient) GetDictDetail(ctx context.Context, in *v11.GetDictDetailRequest, opts ...grpc.CallOption) (*v11.DictDetail, error) {
	out := new(v11.DictDetail)
	err := c.cc.Invoke(ctx, DictDetailService_GetDictDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictDetailServiceClient) CreateDictDetail(ctx context.Context, in *v11.CreateDictDetailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DictDetailService_CreateDictDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictDetailServiceClient) UpdateDictDetail(ctx context.Context, in *v11.UpdateDictDetailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DictDetailService_UpdateDictDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictDetailServiceClient) DeleteDictDetail(ctx context.Context, in *v11.DeleteDictDetailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DictDetailService_DeleteDictDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictDetailServiceServer is the server API for DictDetailService service.
// All implementations must embed UnimplementedDictDetailServiceServer
// for forward compatibility
type DictDetailServiceServer interface {
	// 查询字典详情列表
	ListDictDetail(context.Context, *v1.PagingRequest) (*v11.ListDictDetailResponse, error)
	// 查询字典详情
	GetDictDetail(context.Context, *v11.GetDictDetailRequest) (*v11.DictDetail, error)
	// 创建字典详情
	CreateDictDetail(context.Context, *v11.CreateDictDetailRequest) (*emptypb.Empty, error)
	// 更新字典详情
	UpdateDictDetail(context.Context, *v11.UpdateDictDetailRequest) (*emptypb.Empty, error)
	// 删除字典详情
	DeleteDictDetail(context.Context, *v11.DeleteDictDetailRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDictDetailServiceServer()
}

// UnimplementedDictDetailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDictDetailServiceServer struct {
}

func (UnimplementedDictDetailServiceServer) ListDictDetail(context.Context, *v1.PagingRequest) (*v11.ListDictDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDictDetail not implemented")
}
func (UnimplementedDictDetailServiceServer) GetDictDetail(context.Context, *v11.GetDictDetailRequest) (*v11.DictDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDictDetail not implemented")
}
func (UnimplementedDictDetailServiceServer) CreateDictDetail(context.Context, *v11.CreateDictDetailRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDictDetail not implemented")
}
func (UnimplementedDictDetailServiceServer) UpdateDictDetail(context.Context, *v11.UpdateDictDetailRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDictDetail not implemented")
}
func (UnimplementedDictDetailServiceServer) DeleteDictDetail(context.Context, *v11.DeleteDictDetailRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDictDetail not implemented")
}
func (UnimplementedDictDetailServiceServer) mustEmbedUnimplementedDictDetailServiceServer() {}

// UnsafeDictDetailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictDetailServiceServer will
// result in compilation errors.
type UnsafeDictDetailServiceServer interface {
	mustEmbedUnimplementedDictDetailServiceServer()
}

func RegisterDictDetailServiceServer(s grpc.ServiceRegistrar, srv DictDetailServiceServer) {
	s.RegisterService(&DictDetailService_ServiceDesc, srv)
}

func _DictDetailService_ListDictDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictDetailServiceServer).ListDictDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictDetailService_ListDictDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictDetailServiceServer).ListDictDetail(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictDetailService_GetDictDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.GetDictDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictDetailServiceServer).GetDictDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictDetailService_GetDictDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictDetailServiceServer).GetDictDetail(ctx, req.(*v11.GetDictDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictDetailService_CreateDictDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.CreateDictDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictDetailServiceServer).CreateDictDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictDetailService_CreateDictDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictDetailServiceServer).CreateDictDetail(ctx, req.(*v11.CreateDictDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictDetailService_UpdateDictDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.UpdateDictDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictDetailServiceServer).UpdateDictDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictDetailService_UpdateDictDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictDetailServiceServer).UpdateDictDetail(ctx, req.(*v11.UpdateDictDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictDetailService_DeleteDictDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.DeleteDictDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictDetailServiceServer).DeleteDictDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictDetailService_DeleteDictDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictDetailServiceServer).DeleteDictDetail(ctx, req.(*v11.DeleteDictDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DictDetailService_ServiceDesc is the grpc.ServiceDesc for DictDetailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DictDetailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.DictDetailService",
	HandlerType: (*DictDetailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDictDetail",
			Handler:    _DictDetailService_ListDictDetail_Handler,
		},
		{
			MethodName: "GetDictDetail",
			Handler:    _DictDetailService_GetDictDetail_Handler,
		},
		{
			MethodName: "CreateDictDetail",
			Handler:    _DictDetailService_CreateDictDetail_Handler,
		},
		{
			MethodName: "UpdateDictDetail",
			Handler:    _DictDetailService_UpdateDictDetail_Handler,
		},
		{
			MethodName: "DeleteDictDetail",
			Handler:    _DictDetailService_DeleteDictDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/service/v1/i_dict_detail.proto",
}
