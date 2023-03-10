// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: auth/proto/auth.proto

package types

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

type SayMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Export string `protobuf:"bytes,1,opt,name=export,proto3" json:"export,omitempty"`
}

func (x *SayMessage) Reset() {
	*x = SayMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayMessage) ProtoMessage() {}

func (x *SayMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayMessage.ProtoReflect.Descriptor instead.
func (*SayMessage) Descriptor() ([]byte, []int) {
	return file_auth_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SayMessage) GetExport() string {
	if x != nil {
		return x.Export
	}
	return ""
}

var File_auth_proto_auth_proto protoreflect.FileDescriptor

var file_auth_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74,
	0x75, 0x64, 0x79, 0x2e, 0x73, 0x72, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x22, 0x24, 0x0a, 0x0a,
	0x53, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x32, 0x57, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x4a, 0x0a, 0x06, 0x53, 0x61, 0x79, 0x4d, 0x73, 0x67, 0x12, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x73, 0x72, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x53, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x73, 0x72, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x53, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x72, 0x79, 0x46,
	0x72, 0x65, 0x65, 0x35, 0x36, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_auth_proto_rawDescOnce sync.Once
	file_auth_proto_auth_proto_rawDescData = file_auth_proto_auth_proto_rawDesc
)

func file_auth_proto_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_auth_proto_rawDescData)
	})
	return file_auth_proto_auth_proto_rawDescData
}

var file_auth_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_auth_proto_auth_proto_goTypes = []interface{}{
	(*SayMessage)(nil), // 0: grpc_study.src.auth.SayMessage
}
var file_auth_proto_auth_proto_depIdxs = []int32{
	0, // 0: grpc_study.src.auth.AuthToken.SayMsg:input_type -> grpc_study.src.auth.SayMessage
	0, // 1: grpc_study.src.auth.AuthToken.SayMsg:output_type -> grpc_study.src.auth.SayMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_auth_proto_init() }
func file_auth_proto_auth_proto_init() {
	if File_auth_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayMessage); i {
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
			RawDescriptor: file_auth_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_auth_proto_msgTypes,
	}.Build()
	File_auth_proto_auth_proto = out.File
	file_auth_proto_auth_proto_rawDesc = nil
	file_auth_proto_auth_proto_goTypes = nil
	file_auth_proto_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthTokenClient is the client API for AuthToken service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthTokenClient interface {
	SayMsg(ctx context.Context, in *SayMessage, opts ...grpc.CallOption) (*SayMessage, error)
}

type authTokenClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthTokenClient(cc grpc.ClientConnInterface) AuthTokenClient {
	return &authTokenClient{cc}
}

func (c *authTokenClient) SayMsg(ctx context.Context, in *SayMessage, opts ...grpc.CallOption) (*SayMessage, error) {
	out := new(SayMessage)
	err := c.cc.Invoke(ctx, "/grpc_study.src.auth.AuthToken/SayMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthTokenServer is the server API for AuthToken service.
type AuthTokenServer interface {
	SayMsg(context.Context, *SayMessage) (*SayMessage, error)
}

// UnimplementedAuthTokenServer can be embedded to have forward compatible implementations.
type UnimplementedAuthTokenServer struct {
}

func (*UnimplementedAuthTokenServer) SayMsg(context.Context, *SayMessage) (*SayMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayMsg not implemented")
}

func RegisterAuthTokenServer(s *grpc.Server, srv AuthTokenServer) {
	s.RegisterService(&_AuthToken_serviceDesc, srv)
}

func _AuthToken_SayMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthTokenServer).SayMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_study.src.auth.AuthToken/SayMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthTokenServer).SayMsg(ctx, req.(*SayMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthToken_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_study.src.auth.AuthToken",
	HandlerType: (*AuthTokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayMsg",
			Handler:    _AuthToken_SayMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/proto/auth.proto",
}
