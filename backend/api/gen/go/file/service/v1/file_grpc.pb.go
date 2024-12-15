// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: file/service/v1/file.proto

package servicev1

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
	FileService_OssUploadUrl_FullMethodName   = "/file.service.v1.FileService/OssUploadUrl"
	FileService_GetDownloadUrl_FullMethodName = "/file.service.v1.FileService/GetDownloadUrl"
	FileService_ListFile_FullMethodName       = "/file.service.v1.FileService/ListFile"
	FileService_DeleteFile_FullMethodName     = "/file.service.v1.FileService/DeleteFile"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 文件服务
type FileServiceClient interface {
	// 获取对象存储（OSS）上传链接
	OssUploadUrl(ctx context.Context, in *OssUploadUrlRequest, opts ...grpc.CallOption) (*OssUploadUrlResponse, error)
	// 获取对象存储（OSS）下载链接
	GetDownloadUrl(ctx context.Context, in *GetDownloadUrlRequest, opts ...grpc.CallOption) (*GetDownloadUrlResponse, error)
	// 获取文件夹下面的文件列表
	ListFile(ctx context.Context, in *ListFileRequest, opts ...grpc.CallOption) (*ListFileResponse, error)
	// 删除一个文件
	DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) OssUploadUrl(ctx context.Context, in *OssUploadUrlRequest, opts ...grpc.CallOption) (*OssUploadUrlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OssUploadUrlResponse)
	err := c.cc.Invoke(ctx, FileService_OssUploadUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetDownloadUrl(ctx context.Context, in *GetDownloadUrlRequest, opts ...grpc.CallOption) (*GetDownloadUrlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDownloadUrlResponse)
	err := c.cc.Invoke(ctx, FileService_GetDownloadUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) ListFile(ctx context.Context, in *ListFileRequest, opts ...grpc.CallOption) (*ListFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFileResponse)
	err := c.cc.Invoke(ctx, FileService_ListFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteFileResponse)
	err := c.cc.Invoke(ctx, FileService_DeleteFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility.
//
// 文件服务
type FileServiceServer interface {
	// 获取对象存储（OSS）上传链接
	OssUploadUrl(context.Context, *OssUploadUrlRequest) (*OssUploadUrlResponse, error)
	// 获取对象存储（OSS）下载链接
	GetDownloadUrl(context.Context, *GetDownloadUrlRequest) (*GetDownloadUrlResponse, error)
	// 获取文件夹下面的文件列表
	ListFile(context.Context, *ListFileRequest) (*ListFileResponse, error)
	// 删除一个文件
	DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServiceServer struct{}

func (UnimplementedFileServiceServer) OssUploadUrl(context.Context, *OssUploadUrlRequest) (*OssUploadUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OssUploadUrl not implemented")
}
func (UnimplementedFileServiceServer) GetDownloadUrl(context.Context, *GetDownloadUrlRequest) (*GetDownloadUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadUrl not implemented")
}
func (UnimplementedFileServiceServer) ListFile(context.Context, *ListFileRequest) (*ListFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFile not implemented")
}
func (UnimplementedFileServiceServer) DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}
func (UnimplementedFileServiceServer) testEmbeddedByValue()                     {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	// If the following call pancis, it indicates UnimplementedFileServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_OssUploadUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OssUploadUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).OssUploadUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_OssUploadUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).OssUploadUrl(ctx, req.(*OssUploadUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetDownloadUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDownloadUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetDownloadUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_GetDownloadUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetDownloadUrl(ctx, req.(*GetDownloadUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_ListFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).ListFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_ListFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).ListFile(ctx, req.(*ListFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DeleteFile(ctx, req.(*DeleteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.service.v1.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OssUploadUrl",
			Handler:    _FileService_OssUploadUrl_Handler,
		},
		{
			MethodName: "GetDownloadUrl",
			Handler:    _FileService_GetDownloadUrl_Handler,
		},
		{
			MethodName: "ListFile",
			Handler:    _FileService_ListFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _FileService_DeleteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file/service/v1/file.proto",
}
