// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: pkg/order/pb/order.proto

package pb

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

type AddOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ProductId int64 `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Quantity  int64 `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (x *AddOrderRequest) Reset() {
	*x = AddOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_order_pb_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderRequest) ProtoMessage() {}

func (x *AddOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_pb_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderRequest.ProtoReflect.Descriptor instead.
func (*AddOrderRequest) Descriptor() ([]byte, []int) {
	return file_pkg_order_pb_order_proto_rawDescGZIP(), []int{0}
}

func (x *AddOrderRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddOrderRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *AddOrderRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddOrderResponse) Reset() {
	*x = AddOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_order_pb_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderResponse) ProtoMessage() {}

func (x *AddOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_pb_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderResponse.ProtoReflect.Descriptor instead.
func (*AddOrderResponse) Descriptor() ([]byte, []int) {
	return file_pkg_order_pb_order_proto_rawDescGZIP(), []int{1}
}

func (x *AddOrderResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddOrderResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *AddOrderResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_pkg_order_pb_order_proto protoreflect.FileDescriptor

var file_pkg_order_pb_order_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x22, 0x63, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x50, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x4c, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_order_pb_order_proto_rawDescOnce sync.Once
	file_pkg_order_pb_order_proto_rawDescData = file_pkg_order_pb_order_proto_rawDesc
)

func file_pkg_order_pb_order_proto_rawDescGZIP() []byte {
	file_pkg_order_pb_order_proto_rawDescOnce.Do(func() {
		file_pkg_order_pb_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_order_pb_order_proto_rawDescData)
	})
	return file_pkg_order_pb_order_proto_rawDescData
}

var file_pkg_order_pb_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_order_pb_order_proto_goTypes = []interface{}{
	(*AddOrderRequest)(nil),  // 0: order.AddOrderRequest
	(*AddOrderResponse)(nil), // 1: order.AddOrderResponse
}
var file_pkg_order_pb_order_proto_depIdxs = []int32{
	0, // 0: order.CartService.AddOrder:input_type -> order.AddOrderRequest
	1, // 1: order.CartService.AddOrder:output_type -> order.AddOrderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_order_pb_order_proto_init() }
func file_pkg_order_pb_order_proto_init() {
	if File_pkg_order_pb_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_order_pb_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrderRequest); i {
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
		file_pkg_order_pb_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrderResponse); i {
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
			RawDescriptor: file_pkg_order_pb_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_order_pb_order_proto_goTypes,
		DependencyIndexes: file_pkg_order_pb_order_proto_depIdxs,
		MessageInfos:      file_pkg_order_pb_order_proto_msgTypes,
	}.Build()
	File_pkg_order_pb_order_proto = out.File
	file_pkg_order_pb_order_proto_rawDesc = nil
	file_pkg_order_pb_order_proto_goTypes = nil
	file_pkg_order_pb_order_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartServiceClient interface {
	AddOrder(ctx context.Context, in *AddOrderRequest, opts ...grpc.CallOption) (*AddOrderResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) AddOrder(ctx context.Context, in *AddOrderRequest, opts ...grpc.CallOption) (*AddOrderResponse, error) {
	out := new(AddOrderResponse)
	err := c.cc.Invoke(ctx, "/order.CartService/AddOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
type CartServiceServer interface {
	AddOrder(context.Context, *AddOrderRequest) (*AddOrderResponse, error)
}

// UnimplementedCartServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (*UnimplementedCartServiceServer) AddOrder(context.Context, *AddOrderRequest) (*AddOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}

func RegisterCartServiceServer(s *grpc.Server, srv CartServiceServer) {
	s.RegisterService(&_CartService_serviceDesc, srv)
}

func _CartService_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.CartService/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddOrder(ctx, req.(*AddOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrder",
			Handler:    _CartService_AddOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/order/pb/order.proto",
}
