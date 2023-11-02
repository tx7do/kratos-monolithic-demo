// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: admin/service/v1/i_dict.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-monolithic-demo/gen/api/go/common/pagination"
	v1 "kratos-monolithic-demo/gen/api/go/system/service/v1"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_admin_service_v1_i_dict_proto protoreflect.FileDescriptor

var file_admin_service_v1_i_dict_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbb, 0x05, 0x0a, 0x0b, 0x44, 0x69, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63,
	0x74, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x36, 0xba, 0x47, 0x1c, 0x5a, 0x1a, 0x0a, 0x18, 0x0a, 0x14, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x32, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x44, 0x69, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x63,
	0x74, 0x22, 0x3b, 0xba, 0x47, 0x1c, 0x5a, 0x1a, 0x0a, 0x18, 0x0a, 0x14, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x32, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x86,
	0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x12, 0x24, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x22, 0x39, 0xba, 0x47,
	0x1c, 0x5a, 0x1a, 0x0a, 0x18, 0x0a, 0x14, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x63, 0x74, 0x22, 0x41, 0xba, 0x47, 0x1c, 0x5a, 0x1a, 0x0a, 0x18, 0x0a, 0x14,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x04, 0x64, 0x69,
	0x63, 0x74, 0x1a, 0x14, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x87, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3b, 0xba, 0x47, 0x1c, 0x5a, 0x1a, 0x0a, 0x18, 0x0a, 0x14,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x42, 0xc2, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x49, 0x44, 0x69,
	0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x69, 0x63, 0x2d, 0x64, 0x65, 0x6d,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53, 0x58, 0xaa, 0x02, 0x10,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_admin_service_v1_i_dict_proto_goTypes = []interface{}{
	(*pagination.PagingRequest)(nil), // 0: pagination.PagingRequest
	(*v1.GetDictRequest)(nil),        // 1: system.service.v1.GetDictRequest
	(*v1.CreateDictRequest)(nil),     // 2: system.service.v1.CreateDictRequest
	(*v1.UpdateDictRequest)(nil),     // 3: system.service.v1.UpdateDictRequest
	(*v1.DeleteDictRequest)(nil),     // 4: system.service.v1.DeleteDictRequest
	(*v1.ListDictResponse)(nil),      // 5: system.service.v1.ListDictResponse
	(*v1.Dict)(nil),                  // 6: system.service.v1.Dict
	(*emptypb.Empty)(nil),            // 7: google.protobuf.Empty
}
var file_admin_service_v1_i_dict_proto_depIdxs = []int32{
	0, // 0: admin.service.v1.DictService.ListDict:input_type -> pagination.PagingRequest
	1, // 1: admin.service.v1.DictService.GetDict:input_type -> system.service.v1.GetDictRequest
	2, // 2: admin.service.v1.DictService.CreateDict:input_type -> system.service.v1.CreateDictRequest
	3, // 3: admin.service.v1.DictService.UpdateDict:input_type -> system.service.v1.UpdateDictRequest
	4, // 4: admin.service.v1.DictService.DeleteDict:input_type -> system.service.v1.DeleteDictRequest
	5, // 5: admin.service.v1.DictService.ListDict:output_type -> system.service.v1.ListDictResponse
	6, // 6: admin.service.v1.DictService.GetDict:output_type -> system.service.v1.Dict
	6, // 7: admin.service.v1.DictService.CreateDict:output_type -> system.service.v1.Dict
	6, // 8: admin.service.v1.DictService.UpdateDict:output_type -> system.service.v1.Dict
	7, // 9: admin.service.v1.DictService.DeleteDict:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_service_v1_i_dict_proto_init() }
func file_admin_service_v1_i_dict_proto_init() {
	if File_admin_service_v1_i_dict_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_service_v1_i_dict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_service_v1_i_dict_proto_goTypes,
		DependencyIndexes: file_admin_service_v1_i_dict_proto_depIdxs,
	}.Build()
	File_admin_service_v1_i_dict_proto = out.File
	file_admin_service_v1_i_dict_proto_rawDesc = nil
	file_admin_service_v1_i_dict_proto_goTypes = nil
	file_admin_service_v1_i_dict_proto_depIdxs = nil
}
