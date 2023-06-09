// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: grpc_challenge/server.proto

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

// The GenericRequest message contains a JSON-encoded request that the server should forward to Osmosis public RPC
type GenericRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GenericRequest) Reset() {
	*x = GenericRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_challenge_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericRequest) ProtoMessage() {}

func (x *GenericRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_challenge_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericRequest.ProtoReflect.Descriptor instead.
func (*GenericRequest) Descriptor() ([]byte, []int) {
	return file_grpc_challenge_server_proto_rawDescGZIP(), []int{0}
}

func (x *GenericRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *GenericRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The GenericResponse message contains a JSON-encoded response from Osmosis public RPC
type GenericResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GenericResponse) Reset() {
	*x = GenericResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_challenge_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericResponse) ProtoMessage() {}

func (x *GenericResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_challenge_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericResponse.ProtoReflect.Descriptor instead.
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return file_grpc_challenge_server_proto_rawDescGZIP(), []int{1}
}

func (x *GenericResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_challenge_server_proto protoreflect.FileDescriptor

var file_grpc_challenge_server_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x22, 0x42, 0x0a,
	0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x63,
	0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x51, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x68, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_challenge_server_proto_rawDescOnce sync.Once
	file_grpc_challenge_server_proto_rawDescData = file_grpc_challenge_server_proto_rawDesc
)

func file_grpc_challenge_server_proto_rawDescGZIP() []byte {
	file_grpc_challenge_server_proto_rawDescOnce.Do(func() {
		file_grpc_challenge_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_challenge_server_proto_rawDescData)
	})
	return file_grpc_challenge_server_proto_rawDescData
}

var file_grpc_challenge_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_challenge_server_proto_goTypes = []interface{}{
	(*GenericRequest)(nil),  // 0: grpc_challenge.GenericRequest
	(*GenericResponse)(nil), // 1: grpc_challenge.GenericResponse
}
var file_grpc_challenge_server_proto_depIdxs = []int32{
	0, // 0: grpc_challenge.GenericService.ForwardRequest:input_type -> grpc_challenge.GenericRequest
	1, // 1: grpc_challenge.GenericService.ForwardRequest:output_type -> grpc_challenge.GenericResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_challenge_server_proto_init() }
func file_grpc_challenge_server_proto_init() {
	if File_grpc_challenge_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_challenge_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericRequest); i {
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
		file_grpc_challenge_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericResponse); i {
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
			RawDescriptor: file_grpc_challenge_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_challenge_server_proto_goTypes,
		DependencyIndexes: file_grpc_challenge_server_proto_depIdxs,
		MessageInfos:      file_grpc_challenge_server_proto_msgTypes,
	}.Build()
	File_grpc_challenge_server_proto = out.File
	file_grpc_challenge_server_proto_rawDesc = nil
	file_grpc_challenge_server_proto_goTypes = nil
	file_grpc_challenge_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GenericServiceClient is the client API for GenericService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GenericServiceClient interface {
	ForwardRequest(ctx context.Context, in *GenericRequest, opts ...grpc.CallOption) (*GenericResponse, error)
}

type genericServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenericServiceClient(cc grpc.ClientConnInterface) GenericServiceClient {
	return &genericServiceClient{cc}
}

func (c *genericServiceClient) ForwardRequest(ctx context.Context, in *GenericRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/grpc_challenge.GenericService/ForwardRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenericServiceServer is the server API for GenericService service.
type GenericServiceServer interface {
	ForwardRequest(context.Context, *GenericRequest) (*GenericResponse, error)
}

// UnimplementedGenericServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGenericServiceServer struct {
}

func (*UnimplementedGenericServiceServer) ForwardRequest(context.Context, *GenericRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardRequest not implemented")
}

func RegisterGenericServiceServer(s *grpc.Server, srv GenericServiceServer) {
	s.RegisterService(&_GenericService_serviceDesc, srv)
}

func _GenericService_ForwardRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServiceServer).ForwardRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_challenge.GenericService/ForwardRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServiceServer).ForwardRequest(ctx, req.(*GenericRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GenericService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_challenge.GenericService",
	HandlerType: (*GenericServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForwardRequest",
			Handler:    _GenericService_ForwardRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_challenge/server.proto",
}
