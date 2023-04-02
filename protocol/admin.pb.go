// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protocol/admin/admin.proto

package jiancongtang

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// RegistrationReq 注册请求
type RegistrationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Passwd string `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	Avatar string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *RegistrationReq) Reset() {
	*x = RegistrationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_admin_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationReq) ProtoMessage() {}

func (x *RegistrationReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_admin_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationReq.ProtoReflect.Descriptor instead.
func (*RegistrationReq) Descriptor() ([]byte, []int) {
	return file_protocol_admin_admin_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrationReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegistrationReq) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

func (x *RegistrationReq) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

// RegistrationRsp 注册响应
type RegistrationRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrInfo *ErrMsg `protobuf:"bytes,1,opt,name=errInfo,proto3" json:"errInfo,omitempty"`
}

func (x *RegistrationRsp) Reset() {
	*x = RegistrationRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_admin_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRsp) ProtoMessage() {}

func (x *RegistrationRsp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_admin_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRsp.ProtoReflect.Descriptor instead.
func (*RegistrationRsp) Descriptor() ([]byte, []int) {
	return file_protocol_admin_admin_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrationRsp) GetErrInfo() *ErrMsg {
	if x != nil {
		return x.ErrInfo
	}
	return nil
}

type ErrMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err     string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	ErrCode int32  `protobuf:"varint,2,opt,name=errCode,proto3" json:"errCode,omitempty"` // 0成功 非零失败
}

func (x *ErrMsg) Reset() {
	*x = ErrMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_admin_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrMsg) ProtoMessage() {}

func (x *ErrMsg) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_admin_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrMsg.ProtoReflect.Descriptor instead.
func (*ErrMsg) Descriptor() ([]byte, []int) {
	return file_protocol_admin_admin_proto_rawDescGZIP(), []int{2}
}

func (x *ErrMsg) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *ErrMsg) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

var File_protocol_admin_admin_proto protoreflect.FileDescriptor

var file_protocol_admin_admin_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x55, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22,
	0x3f, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x73, 0x70, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x07, 0x65, 0x72, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x34, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65,
	0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x4a, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x6a, 0x69, 0x61, 0x6e, 0x63, 0x6f, 0x6e, 0x67, 0x74,
	0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_admin_admin_proto_rawDescOnce sync.Once
	file_protocol_admin_admin_proto_rawDescData = file_protocol_admin_admin_proto_rawDesc
)

func file_protocol_admin_admin_proto_rawDescGZIP() []byte {
	file_protocol_admin_admin_proto_rawDescOnce.Do(func() {
		file_protocol_admin_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_admin_admin_proto_rawDescData)
	})
	return file_protocol_admin_admin_proto_rawDescData
}

var file_protocol_admin_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocol_admin_admin_proto_goTypes = []interface{}{
	(*RegistrationReq)(nil), // 0: helloworld.RegistrationReq
	(*RegistrationRsp)(nil), // 1: helloworld.RegistrationRsp
	(*ErrMsg)(nil),          // 2: helloworld.errMsg
}
var file_protocol_admin_admin_proto_depIdxs = []int32{
	2, // 0: helloworld.RegistrationRsp.errInfo:type_name -> helloworld.errMsg
	0, // 1: helloworld.Admin.Registration:input_type -> helloworld.RegistrationReq
	0, // 2: helloworld.Admin.Registration:output_type -> helloworld.RegistrationReq
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocol_admin_admin_proto_init() }
func file_protocol_admin_admin_proto_init() {
	if File_protocol_admin_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_admin_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationReq); i {
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
		file_protocol_admin_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRsp); i {
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
		file_protocol_admin_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrMsg); i {
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
			RawDescriptor: file_protocol_admin_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_admin_admin_proto_goTypes,
		DependencyIndexes: file_protocol_admin_admin_proto_depIdxs,
		MessageInfos:      file_protocol_admin_admin_proto_msgTypes,
	}.Build()
	File_protocol_admin_admin_proto = out.File
	file_protocol_admin_admin_proto_rawDesc = nil
	file_protocol_admin_admin_proto_goTypes = nil
	file_protocol_admin_admin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminClient interface {
	// 注册接口
	Registration(ctx context.Context, in *RegistrationReq, opts ...grpc.CallOption) (*RegistrationReq, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) Registration(ctx context.Context, in *RegistrationReq, opts ...grpc.CallOption) (*RegistrationReq, error) {
	out := new(RegistrationReq)
	err := c.cc.Invoke(ctx, "/helloworld.Admin/Registration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	// 注册接口
	Registration(context.Context, *RegistrationReq) (*RegistrationReq, error)
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) Registration(context.Context, *RegistrationReq) (*RegistrationReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Admin/Registration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Registration(ctx, req.(*RegistrationReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registration",
			Handler:    _Admin_Registration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/admin/admin.proto",
}
