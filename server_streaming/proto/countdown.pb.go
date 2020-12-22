// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/countdown.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

type CountdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timer int32 `protobuf:"varint,1,opt,name=Timer,proto3" json:"Timer,omitempty"`
}

func (x *CountdownRequest) Reset() {
	*x = CountdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_countdown_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountdownRequest) ProtoMessage() {}

func (x *CountdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_countdown_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountdownRequest.ProtoReflect.Descriptor instead.
func (*CountdownRequest) Descriptor() ([]byte, []int) {
	return file_proto_countdown_proto_rawDescGZIP(), []int{0}
}

func (x *CountdownRequest) GetTimer() int32 {
	if x != nil {
		return x.Timer
	}
	return 0
}

type CountdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountdownResponse) Reset() {
	*x = CountdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_countdown_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountdownResponse) ProtoMessage() {}

func (x *CountdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_countdown_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountdownResponse.ProtoReflect.Descriptor instead.
func (*CountdownResponse) Descriptor() ([]byte, []int) {
	return file_proto_countdown_proto_rawDescGZIP(), []int{1}
}

func (x *CountdownResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_countdown_proto protoreflect.FileDescriptor

var file_proto_countdown_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x6f, 0x77,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x22, 0x28, 0x0a, 0x10, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x54, 0x69,
	0x6d, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x11, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x6f, 0x77, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x61,
	0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x54, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x6f, 0x77,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4a, 0x61, 0x69, 0x6d, 0x65, 0x52, 0x43, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_countdown_proto_rawDescOnce sync.Once
	file_proto_countdown_proto_rawDescData = file_proto_countdown_proto_rawDesc
)

func file_proto_countdown_proto_rawDescGZIP() []byte {
	file_proto_countdown_proto_rawDescOnce.Do(func() {
		file_proto_countdown_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_countdown_proto_rawDescData)
	})
	return file_proto_countdown_proto_rawDescData
}

var file_proto_countdown_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_countdown_proto_goTypes = []interface{}{
	(*CountdownRequest)(nil),  // 0: server_streaming.CountdownRequest
	(*CountdownResponse)(nil), // 1: server_streaming.CountdownResponse
}
var file_proto_countdown_proto_depIdxs = []int32{
	0, // 0: server_streaming.Countdown.Start:input_type -> server_streaming.CountdownRequest
	1, // 1: server_streaming.Countdown.Start:output_type -> server_streaming.CountdownResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_countdown_proto_init() }
func file_proto_countdown_proto_init() {
	if File_proto_countdown_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_countdown_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountdownRequest); i {
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
		file_proto_countdown_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountdownResponse); i {
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
			RawDescriptor: file_proto_countdown_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_countdown_proto_goTypes,
		DependencyIndexes: file_proto_countdown_proto_depIdxs,
		MessageInfos:      file_proto_countdown_proto_msgTypes,
	}.Build()
	File_proto_countdown_proto = out.File
	file_proto_countdown_proto_rawDesc = nil
	file_proto_countdown_proto_goTypes = nil
	file_proto_countdown_proto_depIdxs = nil
}
