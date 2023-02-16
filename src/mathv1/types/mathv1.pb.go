// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: proto/mathv1.proto

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

type ReqMathv1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *ReqMathv1) Reset() {
	*x = ReqMathv1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mathv1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqMathv1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMathv1) ProtoMessage() {}

func (x *ReqMathv1) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mathv1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMathv1.ProtoReflect.Descriptor instead.
func (*ReqMathv1) Descriptor() ([]byte, []int) {
	return file_proto_mathv1_proto_rawDescGZIP(), []int{0}
}

func (x *ReqMathv1) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *ReqMathv1) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

type ResMathv1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res int64 `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *ResMathv1) Reset() {
	*x = ResMathv1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mathv1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResMathv1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResMathv1) ProtoMessage() {}

func (x *ResMathv1) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mathv1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResMathv1.ProtoReflect.Descriptor instead.
func (*ResMathv1) Descriptor() ([]byte, []int) {
	return file_proto_mathv1_proto_rawDescGZIP(), []int{1}
}

func (x *ResMathv1) GetRes() int64 {
	if x != nil {
		return x.Res
	}
	return 0
}

var File_proto_mathv1_proto protoreflect.FileDescriptor

var file_proto_mathv1_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x2e, 0x6d, 0x61, 0x74, 0x68, 0x22, 0x27, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x4d, 0x61, 0x74, 0x68,
	0x76, 0x31, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x61,
	0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x22, 0x1d,
	0x0a, 0x09, 0x52, 0x65, 0x73, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0x84, 0x02,
	0x0a, 0x06, 0x4d, 0x61, 0x74, 0x68, 0x56, 0x31, 0x12, 0x3d, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74,
	0x68, 0x2e, 0x52, 0x65, 0x71, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x31, 0x1a, 0x1a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x52, 0x65,
	0x73, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x31, 0x12, 0x3d, 0x0a, 0x03, 0x53, 0x75, 0x62, 0x12, 0x1a,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68,
	0x2e, 0x52, 0x65, 0x71, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x31, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73,
	0x4d, 0x61, 0x74, 0x68, 0x76, 0x31, 0x12, 0x3d, 0x0a, 0x03, 0x4d, 0x75, 0x6c, 0x12, 0x1a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x71, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x31, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x4d,
	0x61, 0x74, 0x68, 0x76, 0x31, 0x12, 0x3d, 0x0a, 0x03, 0x44, 0x69, 0x76, 0x12, 0x1a, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x52,
	0x65, 0x71, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x31, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x4d, 0x61,
	0x74, 0x68, 0x76, 0x31, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x72, 0x79, 0x46, 0x72, 0x65, 0x65, 0x35, 0x36, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61,
	0x74, 0x68, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_mathv1_proto_rawDescOnce sync.Once
	file_proto_mathv1_proto_rawDescData = file_proto_mathv1_proto_rawDesc
)

func file_proto_mathv1_proto_rawDescGZIP() []byte {
	file_proto_mathv1_proto_rawDescOnce.Do(func() {
		file_proto_mathv1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mathv1_proto_rawDescData)
	})
	return file_proto_mathv1_proto_rawDescData
}

var file_proto_mathv1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_mathv1_proto_goTypes = []interface{}{
	(*ReqMathv1)(nil), // 0: grpc_study.math.ReqMathv1
	(*ResMathv1)(nil), // 1: grpc_study.math.ResMathv1
}
var file_proto_mathv1_proto_depIdxs = []int32{
	0, // 0: grpc_study.math.MathV1.Add:input_type -> grpc_study.math.ReqMathv1
	0, // 1: grpc_study.math.MathV1.Sub:input_type -> grpc_study.math.ReqMathv1
	0, // 2: grpc_study.math.MathV1.Mul:input_type -> grpc_study.math.ReqMathv1
	0, // 3: grpc_study.math.MathV1.Div:input_type -> grpc_study.math.ReqMathv1
	1, // 4: grpc_study.math.MathV1.Add:output_type -> grpc_study.math.ResMathv1
	1, // 5: grpc_study.math.MathV1.Sub:output_type -> grpc_study.math.ResMathv1
	1, // 6: grpc_study.math.MathV1.Mul:output_type -> grpc_study.math.ResMathv1
	1, // 7: grpc_study.math.MathV1.Div:output_type -> grpc_study.math.ResMathv1
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_mathv1_proto_init() }
func file_proto_mathv1_proto_init() {
	if File_proto_mathv1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mathv1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqMathv1); i {
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
		file_proto_mathv1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResMathv1); i {
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
			RawDescriptor: file_proto_mathv1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mathv1_proto_goTypes,
		DependencyIndexes: file_proto_mathv1_proto_depIdxs,
		MessageInfos:      file_proto_mathv1_proto_msgTypes,
	}.Build()
	File_proto_mathv1_proto = out.File
	file_proto_mathv1_proto_rawDesc = nil
	file_proto_mathv1_proto_goTypes = nil
	file_proto_mathv1_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MathV1Client is the client API for MathV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathV1Client interface {
	Add(ctx context.Context, in *ReqMathv1, opts ...grpc.CallOption) (*ResMathv1, error)
	Sub(ctx context.Context, in *ReqMathv1, opts ...grpc.CallOption) (*ResMathv1, error)
	Mul(ctx context.Context, in *ReqMathv1, opts ...grpc.CallOption) (*ResMathv1, error)
	Div(ctx context.Context, in *ReqMathv1, opts ...grpc.CallOption) (*ResMathv1, error)
}

type mathV1Client struct {
	cc grpc.ClientConnInterface
}

func NewMathV1Client(cc grpc.ClientConnInterface) MathV1Client {
	return &mathV1Client{cc}
}

func (c *mathV1Client) Add(ctx context.Context, in *ReqMathv1, opts ...grpc.CallOption) (*ResMathv1, error) {
	out := new(ResMathv1)
	err := c.cc.Invoke(ctx, "/grpc_study.math.MathV1/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathV1Client) Sub(ctx context.Context, in *ReqMathv1, opts ...grpc.CallOption) (*ResMathv1, error) {
	out := new(ResMathv1)
	err := c.cc.Invoke(ctx, "/grpc_study.math.MathV1/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathV1Client) Mul(ctx context.Context, in *ReqMathv1, opts ...grpc.CallOption) (*ResMathv1, error) {
	out := new(ResMathv1)
	err := c.cc.Invoke(ctx, "/grpc_study.math.MathV1/Mul", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathV1Client) Div(ctx context.Context, in *ReqMathv1, opts ...grpc.CallOption) (*ResMathv1, error) {
	out := new(ResMathv1)
	err := c.cc.Invoke(ctx, "/grpc_study.math.MathV1/Div", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathV1Server is the server API for MathV1 service.
type MathV1Server interface {
	Add(context.Context, *ReqMathv1) (*ResMathv1, error)
	Sub(context.Context, *ReqMathv1) (*ResMathv1, error)
	Mul(context.Context, *ReqMathv1) (*ResMathv1, error)
	Div(context.Context, *ReqMathv1) (*ResMathv1, error)
}

// UnimplementedMathV1Server can be embedded to have forward compatible implementations.
type UnimplementedMathV1Server struct {
}

func (*UnimplementedMathV1Server) Add(context.Context, *ReqMathv1) (*ResMathv1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedMathV1Server) Sub(context.Context, *ReqMathv1) (*ResMathv1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (*UnimplementedMathV1Server) Mul(context.Context, *ReqMathv1) (*ResMathv1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mul not implemented")
}
func (*UnimplementedMathV1Server) Div(context.Context, *ReqMathv1) (*ResMathv1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}

func RegisterMathV1Server(s *grpc.Server, srv MathV1Server) {
	s.RegisterService(&_MathV1_serviceDesc, srv)
}

func _MathV1_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMathv1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathV1Server).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_study.math.MathV1/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathV1Server).Add(ctx, req.(*ReqMathv1))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathV1_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMathv1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathV1Server).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_study.math.MathV1/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathV1Server).Sub(ctx, req.(*ReqMathv1))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathV1_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMathv1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathV1Server).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_study.math.MathV1/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathV1Server).Mul(ctx, req.(*ReqMathv1))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathV1_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMathv1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathV1Server).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_study.math.MathV1/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathV1Server).Div(ctx, req.(*ReqMathv1))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_study.math.MathV1",
	HandlerType: (*MathV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MathV1_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _MathV1_Sub_Handler,
		},
		{
			MethodName: "Mul",
			Handler:    _MathV1_Mul_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _MathV1_Div_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mathv1.proto",
}