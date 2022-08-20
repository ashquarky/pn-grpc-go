// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: mii.proto

package account

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Mii struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Mii) Reset() {
	*x = Mii{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mii_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mii) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mii) ProtoMessage() {}

func (x *Mii) ProtoReflect() protoreflect.Message {
	mi := &file_mii_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mii.ProtoReflect.Descriptor instead.
func (*Mii) Descriptor() ([]byte, []int) {
	return file_mii_proto_rawDescGZIP(), []int{0}
}

func (x *Mii) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Mii) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Mii) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_mii_proto protoreflect.FileDescriptor

var file_mii_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x69, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3f, 0x0a, 0x03, 0x4d, 0x69, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mii_proto_rawDescOnce sync.Once
	file_mii_proto_rawDescData = file_mii_proto_rawDesc
)

func file_mii_proto_rawDescGZIP() []byte {
	file_mii_proto_rawDescOnce.Do(func() {
		file_mii_proto_rawDescData = protoimpl.X.CompressGZIP(file_mii_proto_rawDescData)
	})
	return file_mii_proto_rawDescData
}

var file_mii_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mii_proto_goTypes = []interface{}{
	(*Mii)(nil), // 0: account.Mii
}
var file_mii_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mii_proto_init() }
func file_mii_proto_init() {
	if File_mii_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mii_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mii); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mii_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mii_proto_goTypes,
		DependencyIndexes: file_mii_proto_depIdxs,
		MessageInfos:      file_mii_proto_msgTypes,
	}.Build()
	File_mii_proto = out.File
	file_mii_proto_rawDesc = nil
	file_mii_proto_goTypes = nil
	file_mii_proto_depIdxs = nil
}
