// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: miiverse_service.proto

package miiverse

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_miiverse_service_proto protoreflect.FileDescriptor

var file_miiverse_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x69, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x69, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x1a, 0x1d, 0x73, 0x6d, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x67, 0x0a, 0x08, 0x4d, 0x69, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x5b, 0x0a,
	0x10, 0x53, 0x4d, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x21, 0x2e, 0x6d, 0x69, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x4d, 0x4d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x69, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e,
	0x53, 0x4d, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f,
	0x3b, 0x6d, 0x69, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_miiverse_service_proto_goTypes = []interface{}{
	(*SMMRequestPostIdRequest)(nil),  // 0: miiverse.SMMRequestPostIdRequest
	(*SMMRequestPostIdResponse)(nil), // 1: miiverse.SMMRequestPostIdResponse
}
var file_miiverse_service_proto_depIdxs = []int32{
	0, // 0: miiverse.Miiverse.SMMRequestPostId:input_type -> miiverse.SMMRequestPostIdRequest
	1, // 1: miiverse.Miiverse.SMMRequestPostId:output_type -> miiverse.SMMRequestPostIdResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_miiverse_service_proto_init() }
func file_miiverse_service_proto_init() {
	if File_miiverse_service_proto != nil {
		return
	}
	file_smm_request_post_id_rpc_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_miiverse_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_miiverse_service_proto_goTypes,
		DependencyIndexes: file_miiverse_service_proto_depIdxs,
	}.Build()
	File_miiverse_service_proto = out.File
	file_miiverse_service_proto_rawDesc = nil
	file_miiverse_service_proto_goTypes = nil
	file_miiverse_service_proto_depIdxs = nil
}
