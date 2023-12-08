// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: mygrpc/proto/simple_parser.proto

package pb

import (
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

type ParserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ParserRequest) Reset() {
	*x = ParserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_proto_simple_parser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParserRequest) ProtoMessage() {}

func (x *ParserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_proto_simple_parser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParserRequest.ProtoReflect.Descriptor instead.
func (*ParserRequest) Descriptor() ([]byte, []int) {
	return file_mygrpc_proto_simple_parser_proto_rawDescGZIP(), []int{0}
}

func (x *ParserRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ParserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	State   string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ParserResponse) Reset() {
	*x = ParserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_proto_simple_parser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParserResponse) ProtoMessage() {}

func (x *ParserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_proto_simple_parser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParserResponse.ProtoReflect.Descriptor instead.
func (*ParserResponse) Descriptor() ([]byte, []int) {
	return file_mygrpc_proto_simple_parser_proto_rawDescGZIP(), []int{1}
}

func (x *ParserResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ParserResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ParserResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ParserResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_mygrpc_proto_simple_parser_proto protoreflect.FileDescriptor

var file_mygrpc_proto_simple_parser_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x29, 0x0a, 0x0d, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x68, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x32, 0x59, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x6d,
	0x79, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mygrpc_proto_simple_parser_proto_rawDescOnce sync.Once
	file_mygrpc_proto_simple_parser_proto_rawDescData = file_mygrpc_proto_simple_parser_proto_rawDesc
)

func file_mygrpc_proto_simple_parser_proto_rawDescGZIP() []byte {
	file_mygrpc_proto_simple_parser_proto_rawDescOnce.Do(func() {
		file_mygrpc_proto_simple_parser_proto_rawDescData = protoimpl.X.CompressGZIP(file_mygrpc_proto_simple_parser_proto_rawDescData)
	})
	return file_mygrpc_proto_simple_parser_proto_rawDescData
}

var file_mygrpc_proto_simple_parser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mygrpc_proto_simple_parser_proto_goTypes = []interface{}{
	(*ParserRequest)(nil),  // 0: address.ParserRequest
	(*ParserResponse)(nil), // 1: address.ParserResponse
}
var file_mygrpc_proto_simple_parser_proto_depIdxs = []int32{
	0, // 0: address.AddressParserService.SimpleParser:input_type -> address.ParserRequest
	1, // 1: address.AddressParserService.SimpleParser:output_type -> address.ParserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mygrpc_proto_simple_parser_proto_init() }
func file_mygrpc_proto_simple_parser_proto_init() {
	if File_mygrpc_proto_simple_parser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mygrpc_proto_simple_parser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParserRequest); i {
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
		file_mygrpc_proto_simple_parser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParserResponse); i {
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
			RawDescriptor: file_mygrpc_proto_simple_parser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mygrpc_proto_simple_parser_proto_goTypes,
		DependencyIndexes: file_mygrpc_proto_simple_parser_proto_depIdxs,
		MessageInfos:      file_mygrpc_proto_simple_parser_proto_msgTypes,
	}.Build()
	File_mygrpc_proto_simple_parser_proto = out.File
	file_mygrpc_proto_simple_parser_proto_rawDesc = nil
	file_mygrpc_proto_simple_parser_proto_goTypes = nil
	file_mygrpc_proto_simple_parser_proto_depIdxs = nil
}
