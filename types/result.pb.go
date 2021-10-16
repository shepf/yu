// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.13.0
// source: result.proto

package types

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caller     []byte `protobuf:"bytes,1,opt,name=Caller,proto3" json:"Caller,omitempty"`
	BlockStage string `protobuf:"bytes,2,opt,name=BlockStage,proto3" json:"BlockStage,omitempty"`
	BlockHash  []byte `protobuf:"bytes,3,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	Height     uint64 `protobuf:"varint,4,opt,name=Height,proto3" json:"Height,omitempty"`
	TripodName string `protobuf:"bytes,5,opt,name=TripodName,proto3" json:"TripodName,omitempty"`
	ExecName   string `protobuf:"bytes,6,opt,name=ExecName,proto3" json:"ExecName,omitempty"`
	Value      string `protobuf:"bytes,7,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetCaller() []byte {
	if x != nil {
		return x.Caller
	}
	return nil
}

func (x *Event) GetBlockStage() string {
	if x != nil {
		return x.BlockStage
	}
	return ""
}

func (x *Event) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *Event) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Event) GetTripodName() string {
	if x != nil {
		return x.TripodName
	}
	return ""
}

func (x *Event) GetExecName() string {
	if x != nil {
		return x.ExecName
	}
	return ""
}

func (x *Event) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caller     []byte `protobuf:"bytes,1,opt,name=Caller,proto3" json:"Caller,omitempty"`
	BlockStage string `protobuf:"bytes,2,opt,name=BlockStage,proto3" json:"BlockStage,omitempty"`
	BlockHash  []byte `protobuf:"bytes,3,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	Height     uint64 `protobuf:"varint,4,opt,name=Height,proto3" json:"Height,omitempty"`
	TripodName string `protobuf:"bytes,5,opt,name=TripodName,proto3" json:"TripodName,omitempty"`
	ExecName   string `protobuf:"bytes,6,opt,name=ExecName,proto3" json:"ExecName,omitempty"`
	Err        string `protobuf:"bytes,7,opt,name=Err,proto3" json:"Err,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCaller() []byte {
	if x != nil {
		return x.Caller
	}
	return nil
}

func (x *Error) GetBlockStage() string {
	if x != nil {
		return x.BlockStage
	}
	return ""
}

func (x *Error) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *Error) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Error) GetTripodName() string {
	if x != nil {
		return x.TripodName
	}
	return ""
}

func (x *Error) GetExecName() string {
	if x != nil {
		return x.ExecName
	}
	return ""
}

func (x *Error) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_result_proto protoreflect.FileDescriptor

var file_result_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7,
	0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x72, 0x69, 0x70, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x72, 0x69, 0x70,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x63, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x78, 0x65, 0x63, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x72, 0x69, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x72, 0x69, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x45, 0x78, 0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x45, 0x72, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72, 0x72, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_result_proto_rawDescOnce sync.Once
	file_result_proto_rawDescData = file_result_proto_rawDesc
)

func file_result_proto_rawDescGZIP() []byte {
	file_result_proto_rawDescOnce.Do(func() {
		file_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_result_proto_rawDescData)
	})
	return file_result_proto_rawDescData
}

var file_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_result_proto_goTypes = []interface{}{
	(*Event)(nil), // 0: Event
	(*Error)(nil), // 1: Error
}
var file_result_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_result_proto_init() }
func file_result_proto_init() {
	if File_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_result_proto_goTypes,
		DependencyIndexes: file_result_proto_depIdxs,
		MessageInfos:      file_result_proto_msgTypes,
	}.Build()
	File_result_proto = out.File
	file_result_proto_rawDesc = nil
	file_result_proto_goTypes = nil
	file_result_proto_depIdxs = nil
}
