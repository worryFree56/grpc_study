// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: mathv3.proto

package types

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ReqMathv3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	//customname 修改编译后字段名称
	Oper []Operation `protobuf:"varint,3,rep,packed,name=oper,proto3,enum=grpc_study.math.Operation" json:"oper,omitempty"`
}

func (x *ReqMathv3) Reset() {
	*x = ReqMathv3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mathv3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqMathv3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMathv3) ProtoMessage() {}

func (x *ReqMathv3) ProtoReflect() protoreflect.Message {
	mi := &file_mathv3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMathv3.ProtoReflect.Descriptor instead.
func (*ReqMathv3) Descriptor() ([]byte, []int) {
	return file_mathv3_proto_rawDescGZIP(), []int{0}
}

func (x *ReqMathv3) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *ReqMathv3) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

func (x *ReqMathv3) GetOper() []Operation {
	if x != nil {
		return x.Oper
	}
	return nil
}

type ResMathv3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Result []*Res `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ResMathv3) Reset() {
	*x = ResMathv3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mathv3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResMathv3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResMathv3) ProtoMessage() {}

func (x *ResMathv3) ProtoReflect() protoreflect.Message {
	mi := &file_mathv3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResMathv3.ProtoReflect.Descriptor instead.
func (*ResMathv3) Descriptor() ([]byte, []int) {
	return file_mathv3_proto_rawDescGZIP(), []int{1}
}

func (x *ResMathv3) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResMathv3) GetResult() []*Res {
	if x != nil {
		return x.Result
	}
	return nil
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mathv3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_mathv3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_mathv3_proto_rawDescGZIP(), []int{2}
}

func (x *Res) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Res) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_mathv3_proto protoreflect.FileDescriptor

var file_mathv3_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x68, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6d,
	0x61, 0x74, 0x68, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x09, 0x52,
	0x65, 0x71, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x33, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x62, 0x12, 0x32, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e,
	0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x72, 0x22, 0x4d, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x4d,
	0x61, 0x74, 0x68, 0x76, 0x33, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2f, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xa2, 0x01, 0x0a, 0x06, 0x4d, 0x61, 0x74,
	0x68, 0x56, 0x33, 0x12, 0x97, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d,
	0x61, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x33, 0x1a, 0x1a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x73, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x33, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x4c, 0x12, 0x1f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x6d,
	0x61, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x7b, 0x6f, 0x70, 0x65,
	0x72, 0x7d, 0x5a, 0x29, 0x12, 0x27, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x7b,
	0x61, 0x7d, 0x2f, 0x7b, 0x6f, 0x70, 0x65, 0x72, 0x7d, 0x2f, 0x7b, 0x62, 0x7d, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mathv3_proto_rawDescOnce sync.Once
	file_mathv3_proto_rawDescData = file_mathv3_proto_rawDesc
)

func file_mathv3_proto_rawDescGZIP() []byte {
	file_mathv3_proto_rawDescOnce.Do(func() {
		file_mathv3_proto_rawDescData = protoimpl.X.CompressGZIP(file_mathv3_proto_rawDescData)
	})
	return file_mathv3_proto_rawDescData
}

var file_mathv3_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mathv3_proto_goTypes = []interface{}{
	(*ReqMathv3)(nil), // 0: grpc_study.math.ReqMathv3
	(*ResMathv3)(nil), // 1: grpc_study.math.ResMathv3
	(*Res)(nil),       // 2: grpc_study.math.Res
	(Operation)(0),    // 3: grpc_study.math.Operation
}
var file_mathv3_proto_depIdxs = []int32{
	3, // 0: grpc_study.math.ReqMathv3.oper:type_name -> grpc_study.math.Operation
	2, // 1: grpc_study.math.ResMathv3.result:type_name -> grpc_study.math.Res
	0, // 2: grpc_study.math.MathV3.Operation:input_type -> grpc_study.math.ReqMathv3
	1, // 3: grpc_study.math.MathV3.Operation:output_type -> grpc_study.math.ResMathv3
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mathv3_proto_init() }
func file_mathv3_proto_init() {
	if File_mathv3_proto != nil {
		return
	}
	file_mathv2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mathv3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqMathv3); i {
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
		file_mathv3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResMathv3); i {
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
		file_mathv3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
			RawDescriptor: file_mathv3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mathv3_proto_goTypes,
		DependencyIndexes: file_mathv3_proto_depIdxs,
		MessageInfos:      file_mathv3_proto_msgTypes,
	}.Build()
	File_mathv3_proto = out.File
	file_mathv3_proto_rawDesc = nil
	file_mathv3_proto_goTypes = nil
	file_mathv3_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MathV3Client is the client API for MathV3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathV3Client interface {
	Operation(ctx context.Context, in *ReqMathv3, opts ...grpc.CallOption) (*ResMathv3, error)
}

type mathV3Client struct {
	cc grpc.ClientConnInterface
}

func NewMathV3Client(cc grpc.ClientConnInterface) MathV3Client {
	return &mathV3Client{cc}
}

func (c *mathV3Client) Operation(ctx context.Context, in *ReqMathv3, opts ...grpc.CallOption) (*ResMathv3, error) {
	out := new(ResMathv3)
	err := c.cc.Invoke(ctx, "/grpc_study.math.MathV3/Operation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathV3Server is the server API for MathV3 service.
type MathV3Server interface {
	Operation(context.Context, *ReqMathv3) (*ResMathv3, error)
}

// UnimplementedMathV3Server can be embedded to have forward compatible implementations.
type UnimplementedMathV3Server struct {
}

func (*UnimplementedMathV3Server) Operation(context.Context, *ReqMathv3) (*ResMathv3, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Operation not implemented")
}

func RegisterMathV3Server(s *grpc.Server, srv MathV3Server) {
	s.RegisterService(&_MathV3_serviceDesc, srv)
}

func _MathV3_Operation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMathv3)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathV3Server).Operation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_study.math.MathV3/Operation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathV3Server).Operation(ctx, req.(*ReqMathv3))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathV3_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_study.math.MathV3",
	HandlerType: (*MathV3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Operation",
			Handler:    _MathV3_Operation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mathv3.proto",
}
