// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: mathv2/proto/mathv2.proto

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

type Operation int32

const (
	Operation_Add Operation = 0
	Operation_Sub Operation = 1
	Operation_Mul Operation = 2
	Operation_Div Operation = 3
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "Add",
		1: "Sub",
		2: "Mul",
		3: "Div",
	}
	Operation_value = map[string]int32{
		"Add": 0,
		"Sub": 1,
		"Mul": 2,
		"Div": 3,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_mathv2_proto_mathv2_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_mathv2_proto_mathv2_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_mathv2_proto_mathv2_proto_rawDescGZIP(), []int{0}
}

type ReqMathv2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A    int64       `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B    int64       `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	Oper []Operation `protobuf:"varint,11,rep,packed,name=oper,proto3,enum=grpc_study.math.Operation" json:"oper,omitempty"` //repeated 支持重复此参数
}

func (x *ReqMathv2) Reset() {
	*x = ReqMathv2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mathv2_proto_mathv2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqMathv2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMathv2) ProtoMessage() {}

func (x *ReqMathv2) ProtoReflect() protoreflect.Message {
	mi := &file_mathv2_proto_mathv2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMathv2.ProtoReflect.Descriptor instead.
func (*ReqMathv2) Descriptor() ([]byte, []int) {
	return file_mathv2_proto_mathv2_proto_rawDescGZIP(), []int{0}
}

func (x *ReqMathv2) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *ReqMathv2) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *ReqMathv2) GetOper() []Operation {
	if x != nil {
		return x.Oper
	}
	return nil
}

type ResMathv2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result map[string]string `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //编译之后，返回值为map[string]string类型
}

func (x *ResMathv2) Reset() {
	*x = ResMathv2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mathv2_proto_mathv2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResMathv2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResMathv2) ProtoMessage() {}

func (x *ResMathv2) ProtoReflect() protoreflect.Message {
	mi := &file_mathv2_proto_mathv2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResMathv2.ProtoReflect.Descriptor instead.
func (*ResMathv2) Descriptor() ([]byte, []int) {
	return file_mathv2_proto_mathv2_proto_rawDescGZIP(), []int{1}
}

func (x *ResMathv2) GetResult() map[string]string {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_mathv2_proto_mathv2_proto protoreflect.FileDescriptor

var file_mathv2_proto_mathv2_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x61, 0x74, 0x68, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d,
	0x61, 0x74, 0x68, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x09, 0x52, 0x65,
	0x71, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x32, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x62, 0x12, 0x2e, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d,
	0x61, 0x74, 0x68, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6f,
	0x70, 0x65, 0x72, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a,
	0x04, 0x08, 0x05, 0x10, 0x0b, 0x52, 0x01, 0x63, 0x52, 0x01, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x09,
	0x52, 0x65, 0x73, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x32, 0x12, 0x3e, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x4d,
	0x61, 0x74, 0x68, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x2a, 0x2f, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x75,
	0x62, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x75, 0x6c, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03,
	0x44, 0x69, 0x76, 0x10, 0x03, 0x32, 0xa2, 0x01, 0x0a, 0x06, 0x4d, 0x61, 0x74, 0x68, 0x56, 0x32,
	0x12, 0x97, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68,
	0x2e, 0x52, 0x65, 0x71, 0x4d, 0x61, 0x74, 0x68, 0x76, 0x32, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73,
	0x4d, 0x61, 0x74, 0x68, 0x76, 0x32, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4c, 0x12, 0x1f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x6d, 0x61, 0x74, 0x68,
	0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x7b, 0x6f, 0x70, 0x65, 0x72, 0x7d, 0x5a,
	0x29, 0x12, 0x27, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x6d,
	0x61, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x7b, 0x61, 0x7d, 0x2f,
	0x7b, 0x6f, 0x70, 0x65, 0x72, 0x7d, 0x2f, 0x7b, 0x62, 0x7d, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x72, 0x79, 0x46, 0x72,
	0x65, 0x65, 0x35, 0x36, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mathv2_proto_mathv2_proto_rawDescOnce sync.Once
	file_mathv2_proto_mathv2_proto_rawDescData = file_mathv2_proto_mathv2_proto_rawDesc
)

func file_mathv2_proto_mathv2_proto_rawDescGZIP() []byte {
	file_mathv2_proto_mathv2_proto_rawDescOnce.Do(func() {
		file_mathv2_proto_mathv2_proto_rawDescData = protoimpl.X.CompressGZIP(file_mathv2_proto_mathv2_proto_rawDescData)
	})
	return file_mathv2_proto_mathv2_proto_rawDescData
}

var file_mathv2_proto_mathv2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mathv2_proto_mathv2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mathv2_proto_mathv2_proto_goTypes = []interface{}{
	(Operation)(0),    // 0: grpc_study.math.Operation
	(*ReqMathv2)(nil), // 1: grpc_study.math.ReqMathv2
	(*ResMathv2)(nil), // 2: grpc_study.math.ResMathv2
	nil,               // 3: grpc_study.math.ResMathv2.ResultEntry
}
var file_mathv2_proto_mathv2_proto_depIdxs = []int32{
	0, // 0: grpc_study.math.ReqMathv2.oper:type_name -> grpc_study.math.Operation
	3, // 1: grpc_study.math.ResMathv2.result:type_name -> grpc_study.math.ResMathv2.ResultEntry
	1, // 2: grpc_study.math.MathV2.Operation:input_type -> grpc_study.math.ReqMathv2
	2, // 3: grpc_study.math.MathV2.Operation:output_type -> grpc_study.math.ResMathv2
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mathv2_proto_mathv2_proto_init() }
func file_mathv2_proto_mathv2_proto_init() {
	if File_mathv2_proto_mathv2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mathv2_proto_mathv2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqMathv2); i {
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
		file_mathv2_proto_mathv2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResMathv2); i {
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
			RawDescriptor: file_mathv2_proto_mathv2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mathv2_proto_mathv2_proto_goTypes,
		DependencyIndexes: file_mathv2_proto_mathv2_proto_depIdxs,
		EnumInfos:         file_mathv2_proto_mathv2_proto_enumTypes,
		MessageInfos:      file_mathv2_proto_mathv2_proto_msgTypes,
	}.Build()
	File_mathv2_proto_mathv2_proto = out.File
	file_mathv2_proto_mathv2_proto_rawDesc = nil
	file_mathv2_proto_mathv2_proto_goTypes = nil
	file_mathv2_proto_mathv2_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MathV2Client is the client API for MathV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathV2Client interface {
	Operation(ctx context.Context, in *ReqMathv2, opts ...grpc.CallOption) (*ResMathv2, error)
}

type mathV2Client struct {
	cc grpc.ClientConnInterface
}

func NewMathV2Client(cc grpc.ClientConnInterface) MathV2Client {
	return &mathV2Client{cc}
}

func (c *mathV2Client) Operation(ctx context.Context, in *ReqMathv2, opts ...grpc.CallOption) (*ResMathv2, error) {
	out := new(ResMathv2)
	err := c.cc.Invoke(ctx, "/grpc_study.math.MathV2/Operation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathV2Server is the server API for MathV2 service.
type MathV2Server interface {
	Operation(context.Context, *ReqMathv2) (*ResMathv2, error)
}

// UnimplementedMathV2Server can be embedded to have forward compatible implementations.
type UnimplementedMathV2Server struct {
}

func (*UnimplementedMathV2Server) Operation(context.Context, *ReqMathv2) (*ResMathv2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Operation not implemented")
}

func RegisterMathV2Server(s *grpc.Server, srv MathV2Server) {
	s.RegisterService(&_MathV2_serviceDesc, srv)
}

func _MathV2_Operation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMathv2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathV2Server).Operation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_study.math.MathV2/Operation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathV2Server).Operation(ctx, req.(*ReqMathv2))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_study.math.MathV2",
	HandlerType: (*MathV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Operation",
			Handler:    _MathV2_Operation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mathv2/proto/mathv2.proto",
}
