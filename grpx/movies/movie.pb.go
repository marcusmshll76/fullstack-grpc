// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: movie.proto

package movie

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Rating string `protobuf:"bytes,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Movie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Movie) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

type Search struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Search) Reset() {
	*x = Search{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Search) ProtoMessage() {}

func (x *Search) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Search.ProtoReflect.Descriptor instead.
func (*Search) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{1}
}

func (x *Search) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_movie_proto protoreflect.FileDescriptor

var file_movie_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x43, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x18, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x32, 0x48, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0d, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x0c, 0x2e, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x09, 0x12, 0x07, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x6f, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x3b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_proto_rawDescOnce sync.Once
	file_movie_proto_rawDescData = file_movie_proto_rawDesc
)

func file_movie_proto_rawDescGZIP() []byte {
	file_movie_proto_rawDescOnce.Do(func() {
		file_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_proto_rawDescData)
	})
	return file_movie_proto_rawDescData
}

var file_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_movie_proto_goTypes = []interface{}{
	(*Movie)(nil),  // 0: movie.Movie
	(*Search)(nil), // 1: movie.Search
}
var file_movie_proto_depIdxs = []int32{
	1, // 0: movie.MovieService.GetMovie:input_type -> movie.Search
	0, // 1: movie.MovieService.GetMovie:output_type -> movie.Movie
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_movie_proto_init() }
func file_movie_proto_init() {
	if File_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
		file_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Search); i {
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
			RawDescriptor: file_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_proto_goTypes,
		DependencyIndexes: file_movie_proto_depIdxs,
		MessageInfos:      file_movie_proto_msgTypes,
	}.Build()
	File_movie_proto = out.File
	file_movie_proto_rawDesc = nil
	file_movie_proto_goTypes = nil
	file_movie_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieServiceClient interface {
	GetMovie(ctx context.Context, in *Search, opts ...grpc.CallOption) (*Movie, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) GetMovie(ctx context.Context, in *Search, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/movie.MovieService/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
type MovieServiceServer interface {
	GetMovie(context.Context, *Search) (*Movie, error)
}

// UnimplementedMovieServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (*UnimplementedMovieServiceServer) GetMovie(context.Context, *Search) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}

func RegisterMovieServiceServer(s *grpc.Server, srv MovieServiceServer) {
	s.RegisterService(&_MovieService_serviceDesc, srv)
}

func _MovieService_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movie.MovieService/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovie(ctx, req.(*Search))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "movie.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _MovieService_GetMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}
