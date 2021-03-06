// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: ports.proto

package gorpc

import (
	"context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City        string    `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country     string    `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Alias       []string  `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions     []string  `protobuf:"bytes,6,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates []float32 `protobuf:"fixed32,7,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province    string    `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Timezone    string    `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs      []string  `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code        string    `protobuf:"bytes,11,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_ports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_ports_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Data) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Data) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Data) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *Data) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Data) GetCoordinates() []float32 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Data) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Data) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Data) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *Data) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type PortArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports []*Data `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *PortArray) Reset() {
	*x = PortArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortArray) ProtoMessage() {}

func (x *PortArray) ProtoReflect() protoreflect.Message {
	mi := &file_ports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortArray.ProtoReflect.Descriptor instead.
func (*PortArray) Descriptor() ([]byte, []int) {
	return file_ports_proto_rawDescGZIP(), []int{1}
}

func (x *PortArray) GetPorts() []*Data {
	if x != nil {
		return x.Ports
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_ports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_ports_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Result) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PortKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PortKey) Reset() {
	*x = PortKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortKey) ProtoMessage() {}

func (x *PortKey) ProtoReflect() protoreflect.Message {
	mi := &file_ports_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortKey.ProtoReflect.Descriptor instead.
func (*PortKey) Descriptor() ([]byte, []int) {
	return file_ports_proto_rawDescGZIP(), []int{3}
}

func (x *PortKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Size  int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_ports_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_ports_proto_rawDescGZIP(), []int{4}
}

func (x *Page) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Page) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

// PortDomainClient is the client API for PortDomain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortDomainClient interface {
	Save(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Result, error)
	GetPort(ctx context.Context, in *PortKey, opts ...grpc.CallOption) (*Data, error)
	GetPorts(ctx context.Context, in *Page, opts ...grpc.CallOption) (*PortArray, error)
}

type portDomainClient struct {
	cc grpc.ClientConnInterface
}

func NewPortDomainClient(cc grpc.ClientConnInterface) PortDomainClient {
	return &portDomainClient{cc}
}

func (c *portDomainClient) Save(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/gorpc.PortDomain/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDomainClient) GetPort(ctx context.Context, in *PortKey, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/gorpc.PortDomain/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDomainClient) GetPorts(ctx context.Context, in *Page, opts ...grpc.CallOption) (*PortArray, error) {
	out := new(PortArray)
	err := c.cc.Invoke(ctx, "/gorpc.PortDomain/GetPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortDomainServer is the server API for PortDomain service.
// All implementations must embed UnimplementedPortDomainServer
// for forward compatibility
type PortDomainServer interface {
	Save(context.Context, *Data) (*Result, error)
	GetPort(context.Context, *PortKey) (*Data, error)
	GetPorts(context.Context, *Page) (*PortArray, error)
	mustEmbedUnimplementedPortDomainServer()
}

// UnimplementedPortDomainServer must be embedded to have forward compatible implementations.
type UnimplementedPortDomainServer struct {
}

func (UnimplementedPortDomainServer) Save(context.Context, *Data) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedPortDomainServer) GetPort(context.Context, *PortKey) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (UnimplementedPortDomainServer) GetPorts(context.Context, *Page) (*PortArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPorts not implemented")
}
func (UnimplementedPortDomainServer) mustEmbedUnimplementedPortDomainServer() {}

// UnsafePortDomainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortDomainServer will
// result in compilation errors.
type UnsafePortDomainServer interface {
	mustEmbedUnimplementedPortDomainServer()
}

func RegisterPortDomainServer(s grpc.ServiceRegistrar, srv PortDomainServer) {
	s.RegisterService(&PortDomain_ServiceDesc, srv)
}

func _PortDomain_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gorpc.PortDomain/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServer).Save(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDomain_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gorpc.PortDomain/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServer).GetPort(ctx, req.(*PortKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDomain_GetPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServer).GetPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gorpc.PortDomain/GetPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServer).GetPorts(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

// PortDomain_ServiceDesc is the grpc.ServiceDesc for PortDomain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortDomain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gorpc.PortDomain",
	HandlerType: (*PortDomainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _PortDomain_Save_Handler,
		},
		{
			MethodName: "GetPort",
			Handler:    _PortDomain_GetPort_Handler,
		},
		{
			MethodName: "GetPorts",
			Handler:    _PortDomain_GetPorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ports.proto",
}

var File_ports_proto protoreflect.FileDescriptor

var file_ports_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x6f, 0x72, 0x70, 0x63, 0x22, 0x90, 0x02, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2e, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x1b, 0x0a, 0x07, 0x50, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x30, 0x0a, 0x04,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x32, 0x89,
	0x01, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x24, 0x0a,
	0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x0b, 0x2e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e,
	0x2e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x1a, 0x0b,
	0x2e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x2b, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x0b, 0x2e, 0x67, 0x6f, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x67, 0x6f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ports_proto_rawDescOnce sync.Once
	file_ports_proto_rawDescData = file_ports_proto_rawDesc
)

func file_ports_proto_rawDescGZIP() []byte {
	file_ports_proto_rawDescOnce.Do(func() {
		file_ports_proto_rawDescData = protoimpl.X.CompressGZIP(file_ports_proto_rawDescData)
	})
	return file_ports_proto_rawDescData
}

var file_ports_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ports_proto_goTypes = []interface{}{
	(*Data)(nil),      // 0: gorpc.Data
	(*PortArray)(nil), // 1: gorpc.PortArray
	(*Result)(nil),    // 2: gorpc.Result
	(*PortKey)(nil),   // 3: gorpc.PortKey
	(*Page)(nil),      // 4: gorpc.Page
}
var file_ports_proto_depIdxs = []int32{
	0, // 0: gorpc.PortArray.ports:type_name -> gorpc.Data
	0, // 1: gorpc.PortDomain.Save:input_type -> gorpc.Data
	3, // 2: gorpc.PortDomain.GetPort:input_type -> gorpc.PortKey
	4, // 3: gorpc.PortDomain.GetPorts:input_type -> gorpc.Page
	2, // 4: gorpc.PortDomain.Save:output_type -> gorpc.Result
	0, // 5: gorpc.PortDomain.GetPort:output_type -> gorpc.Data
	1, // 6: gorpc.PortDomain.GetPorts:output_type -> gorpc.PortArray
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ports_proto_init() }
func file_ports_proto_init() {
	if File_ports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_ports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortArray); i {
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
		file_ports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_ports_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortKey); i {
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
		file_ports_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
			RawDescriptor: file_ports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ports_proto_goTypes,
		DependencyIndexes: file_ports_proto_depIdxs,
		MessageInfos:      file_ports_proto_msgTypes,
	}.Build()
	File_ports_proto = out.File
	file_ports_proto_rawDesc = nil
	file_ports_proto_goTypes = nil
	file_ports_proto_depIdxs = nil
}
