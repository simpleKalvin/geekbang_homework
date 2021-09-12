// Code generated by protocol-gen-go. DO NOT EDIT.
// versions:
// 	protocol-gen-go v1.26.0
// 	protocol        v3.15.8
// source: api/order/v1/web.proto

package v1

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

// 进入
type ComeIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`   // 名称
	Trait string `protobuf:"bytes,2,opt,name=trait,proto3" json:"trait,omitempty"` // ip
}

func (x *ComeIn) Reset() {
	*x = ComeIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_web_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComeIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComeIn) ProtoMessage() {}

func (x *ComeIn) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_web_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComeIn.ProtoReflect.Descriptor instead.
func (*ComeIn) Descriptor() ([]byte, []int) {
	return file_api_order_v1_web_proto_rawDescGZIP(), []int{0}
}

func (x *ComeIn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComeIn) GetTrait() string {
	if x != nil {
		return x.Trait
	}
	return ""
}

// 离开
type GetOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 名称
	Ip   string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`     // ip
}

func (x *GetOut) Reset() {
	*x = GetOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_web_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOut) ProtoMessage() {}

func (x *GetOut) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_web_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOut.ProtoReflect.Descriptor instead.
func (*GetOut) Descriptor() ([]byte, []int) {
	return file_api_order_v1_web_proto_rawDescGZIP(), []int{1}
}

func (x *GetOut) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetOut) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

var File_api_order_v1_web_proto protoreflect.FileDescriptor

var file_api_order_v1_web_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x65, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x65,
	0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x72, 0x61, 0x69, 0x74, 0x22, 0x2c, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x32, 0x25, 0x0a, 0x03, 0x57, 0x65,
	0x62, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x65, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x07, 0x2e, 0x43, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x75,
	0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_order_v1_web_proto_rawDescOnce sync.Once
	file_api_order_v1_web_proto_rawDescData = file_api_order_v1_web_proto_rawDesc
)

func file_api_order_v1_web_proto_rawDescGZIP() []byte {
	file_api_order_v1_web_proto_rawDescOnce.Do(func() {
		file_api_order_v1_web_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_order_v1_web_proto_rawDescData)
	})
	return file_api_order_v1_web_proto_rawDescData
}

var file_api_order_v1_web_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_order_v1_web_proto_goTypes = []interface{}{
	(*ComeIn)(nil), // 0: ComeIn
	(*GetOut)(nil), // 1: GetOut
}
var file_api_order_v1_web_proto_depIdxs = []int32{
	0, // 0: Web.WebService:input_type -> ComeIn
	1, // 1: Web.WebService:output_type -> GetOut
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_order_v1_web_proto_init() }
func file_api_order_v1_web_proto_init() {
	if File_api_order_v1_web_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_order_v1_web_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComeIn); i {
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
		file_api_order_v1_web_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOut); i {
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
			RawDescriptor: file_api_order_v1_web_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_order_v1_web_proto_goTypes,
		DependencyIndexes: file_api_order_v1_web_proto_depIdxs,
		MessageInfos:      file_api_order_v1_web_proto_msgTypes,
	}.Build()
	File_api_order_v1_web_proto = out.File
	file_api_order_v1_web_proto_rawDesc = nil
	file_api_order_v1_web_proto_goTypes = nil
	file_api_order_v1_web_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WebClient is the client API for Web service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebClient interface {
	WebService(ctx context.Context, in *ComeIn, opts ...grpc.CallOption) (*GetOut, error)
}

type webClient struct {
	cc grpc.ClientConnInterface
}

func NewWebClient(cc grpc.ClientConnInterface) WebClient {
	return &webClient{cc}
}

func (c *webClient) WebService(ctx context.Context, in *ComeIn, opts ...grpc.CallOption) (*GetOut, error) {
	out := new(GetOut)
	err := c.cc.Invoke(ctx, "/Web/WebService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebServer is the server API for Web service.
type WebServer interface {
	WebService(context.Context, *ComeIn) (*GetOut, error)
}

// UnimplementedWebServer can be embedded to have forward compatible implementations.
type UnimplementedWebServer struct {
}

func (*UnimplementedWebServer) WebService(context.Context, *ComeIn) (*GetOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebService not implemented")
}

func RegisterWebServer(s *grpc.Server, srv WebServer) {
	s.RegisterService(&_Web_serviceDesc, srv)
}

func _Web_WebService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComeIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).WebService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Web/WebService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).WebService(ctx, req.(*ComeIn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Web_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Web",
	HandlerType: (*WebServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WebService",
			Handler:    _Web_WebService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/order/v1/web.proto",
}
