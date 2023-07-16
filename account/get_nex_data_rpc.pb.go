// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_nex_data_rpc.proto

package grpc_account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetNEXDataRequest struct {
	Pid                  uint32   `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNEXDataRequest) Reset()         { *m = GetNEXDataRequest{} }
func (m *GetNEXDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetNEXDataRequest) ProtoMessage()    {}
func (*GetNEXDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de988723fd5a0207, []int{0}
}

func (m *GetNEXDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNEXDataRequest.Unmarshal(m, b)
}
func (m *GetNEXDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNEXDataRequest.Marshal(b, m, deterministic)
}
func (m *GetNEXDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNEXDataRequest.Merge(m, src)
}
func (m *GetNEXDataRequest) XXX_Size() int {
	return xxx_messageInfo_GetNEXDataRequest.Size(m)
}
func (m *GetNEXDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNEXDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNEXDataRequest proto.InternalMessageInfo

func (m *GetNEXDataRequest) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

type GetNEXDataResponse struct {
	Pid                  uint32   `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	OwningPid            uint32   `protobuf:"varint,3,opt,name=owning_pid,json=owningPid,proto3" json:"owning_pid,omitempty"`
	AccessLevel          uint32   `protobuf:"varint,4,opt,name=access_level,json=accessLevel,proto3" json:"access_level,omitempty"`
	ServerAccessLevel    string   `protobuf:"bytes,5,opt,name=server_access_level,json=serverAccessLevel,proto3" json:"server_access_level,omitempty"`
	FriendCode           string   `protobuf:"bytes,6,opt,name=friend_code,json=friendCode,proto3" json:"friend_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNEXDataResponse) Reset()         { *m = GetNEXDataResponse{} }
func (m *GetNEXDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetNEXDataResponse) ProtoMessage()    {}
func (*GetNEXDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de988723fd5a0207, []int{1}
}

func (m *GetNEXDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNEXDataResponse.Unmarshal(m, b)
}
func (m *GetNEXDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNEXDataResponse.Marshal(b, m, deterministic)
}
func (m *GetNEXDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNEXDataResponse.Merge(m, src)
}
func (m *GetNEXDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetNEXDataResponse.Size(m)
}
func (m *GetNEXDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNEXDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNEXDataResponse proto.InternalMessageInfo

func (m *GetNEXDataResponse) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *GetNEXDataResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *GetNEXDataResponse) GetOwningPid() uint32 {
	if m != nil {
		return m.OwningPid
	}
	return 0
}

func (m *GetNEXDataResponse) GetAccessLevel() uint32 {
	if m != nil {
		return m.AccessLevel
	}
	return 0
}

func (m *GetNEXDataResponse) GetServerAccessLevel() string {
	if m != nil {
		return m.ServerAccessLevel
	}
	return ""
}

func (m *GetNEXDataResponse) GetFriendCode() string {
	if m != nil {
		return m.FriendCode
	}
	return ""
}

func init() {
	proto.RegisterType((*GetNEXDataRequest)(nil), "account.GetNEXDataRequest")
	proto.RegisterType((*GetNEXDataResponse)(nil), "account.GetNEXDataResponse")
}

func init() {
	proto.RegisterFile("get_nex_data_rpc.proto", fileDescriptor_de988723fd5a0207)
}

var fileDescriptor_de988723fd5a0207 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0x59, 0xab, 0xd5, 0x4e, 0x15, 0xdd, 0x08, 0xb2, 0x08, 0x62, 0x2d, 0x08, 0x3d, 0xad,
	0x07, 0x8f, 0x9e, 0xfc, 0xc2, 0x8b, 0x88, 0xec, 0x49, 0xbc, 0x0c, 0x31, 0x19, 0x97, 0x85, 0x92,
	0x89, 0x99, 0xb4, 0xf5, 0x8f, 0xfa, 0x7f, 0x64, 0x13, 0x91, 0x8a, 0xb7, 0x99, 0x67, 0x9e, 0x97,
	0x81, 0x17, 0x8e, 0x5a, 0x8a, 0xe8, 0xe8, 0x13, 0xad, 0x8e, 0x1a, 0x83, 0x37, 0xb5, 0x0f, 0x1c,
	0x59, 0x6d, 0x6b, 0x63, 0x78, 0xe1, 0xe2, 0xf4, 0x1c, 0xca, 0x07, 0x8a, 0x4f, 0xf7, 0x2f, 0x77,
	0x3a, 0xea, 0x86, 0x3e, 0x16, 0x24, 0x51, 0x1d, 0xc0, 0xc0, 0x77, 0xb6, 0x2a, 0x26, 0xc5, 0x6c,
	0xaf, 0xe9, 0xc7, 0xe9, 0x57, 0x01, 0x6a, 0xdd, 0x13, 0xcf, 0x4e, 0xe8, 0xbf, 0xa8, 0x8e, 0x61,
	0xc7, 0x6b, 0x91, 0x15, 0x07, 0x5b, 0x6d, 0x4c, 0x8a, 0xd9, 0xa8, 0xf9, 0xdd, 0xd5, 0x09, 0x00,
	0xaf, 0x5c, 0xe7, 0x5a, 0xec, 0x43, 0x83, 0x14, 0x1a, 0x65, 0xf2, 0xdc, 0x59, 0x75, 0x06, 0xbb,
	0xda, 0x18, 0x12, 0xc1, 0x39, 0x2d, 0x69, 0x5e, 0x6d, 0x26, 0x61, 0x9c, 0xd9, 0x63, 0x8f, 0x54,
	0x0d, 0x87, 0x42, 0x61, 0x49, 0x01, 0xff, 0x98, 0x5b, 0xe9, 0x51, 0x99, 0x4f, 0xd7, 0x6b, 0xfe,
	0x29, 0x8c, 0xdf, 0x43, 0x47, 0xce, 0xa2, 0x61, 0x4b, 0xd5, 0x30, 0x79, 0x90, 0xd1, 0x2d, 0x5b,
	0xba, 0x29, 0x5f, 0xf7, 0xeb, 0x8b, 0xab, 0x36, 0x78, 0x83, 0x3f, 0x8d, 0xbc, 0x0d, 0x53, 0x43,
	0x97, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0x6e, 0x71, 0x32, 0x3b, 0x01, 0x00, 0x00,
}
