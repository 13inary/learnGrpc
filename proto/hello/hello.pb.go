// auto big case all key word

// proto3 相对proto2，支持更多的语言（Ruby、C#等）、删除了一些复杂的语法和特性、引入了更多的约定等

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.2
// source: hello.proto

package hello

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

type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 1 is index
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *HelloRsp) Reset() {
	*x = HelloRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRsp) ProtoMessage() {}

func (x *HelloRsp) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRsp.ProtoReflect.Descriptor instead.
func (*HelloRsp) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRsp) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_hello_proto protoreflect.FileDescriptor

var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a,
	0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a,
	0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x26, 0x0a,
	0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1d, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x09, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_proto_rawDescOnce sync.Once
	file_hello_proto_rawDescData = file_hello_proto_rawDesc
)

func file_hello_proto_rawDescGZIP() []byte {
	file_hello_proto_rawDescOnce.Do(func() {
		file_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_proto_rawDescData)
	})
	return file_hello_proto_rawDescData
}

var file_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hello_proto_goTypes = []interface{}{
	(*HelloReq)(nil), // 0: HelloReq
	(*HelloRsp)(nil), // 1: HelloRsp
}
var file_hello_proto_depIdxs = []int32{
	0, // 0: Hello.Hello:input_type -> HelloReq
	1, // 1: Hello.Hello:output_type -> HelloRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq); i {
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
		file_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRsp); i {
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
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
		MessageInfos:      file_hello_proto_msgTypes,
	}.Build()
	File_hello_proto = out.File
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error) {
	out := new(HelloRsp)
	err := c.cc.Invoke(ctx, "/Hello/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	Hello(context.Context, *HelloReq) (*HelloRsp, error)
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) Hello(context.Context, *HelloReq) (*HelloRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Hello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Hello_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
