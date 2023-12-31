// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: g4.proto

package g4

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

type Mota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ID   int32  `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Mota) Reset() {
	*x = Mota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_g4_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mota) ProtoMessage() {}

func (x *Mota) ProtoReflect() protoreflect.Message {
	mi := &file_g4_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mota.ProtoReflect.Descriptor instead.
func (*Mota) Descriptor() ([]byte, []int) {
	return file_g4_proto_rawDescGZIP(), []int{0}
}

func (x *Mota) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Mota) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_g4_proto protoreflect.FileDescriptor

var file_g4_proto_rawDesc = []byte{
	0x0a, 0x08, 0x67, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x67, 0x34, 0x22, 0x2a,
	0x0a, 0x04, 0x4d, 0x6f, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x32, 0x25, 0x0a, 0x07, 0x44, 0x65,
	0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x1a, 0x0a, 0x04, 0x5a, 0x65, 0x6e, 0x6f, 0x12, 0x08, 0x2e,
	0x67, 0x34, 0x2e, 0x4d, 0x6f, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x67, 0x34, 0x2e, 0x4d, 0x6f, 0x74,
	0x61, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x67, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_g4_proto_rawDescOnce sync.Once
	file_g4_proto_rawDescData = file_g4_proto_rawDesc
)

func file_g4_proto_rawDescGZIP() []byte {
	file_g4_proto_rawDescOnce.Do(func() {
		file_g4_proto_rawDescData = protoimpl.X.CompressGZIP(file_g4_proto_rawDescData)
	})
	return file_g4_proto_rawDescData
}

var file_g4_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_g4_proto_goTypes = []interface{}{
	(*Mota)(nil), // 0: g4.Mota
}
var file_g4_proto_depIdxs = []int32{
	0, // 0: g4.Destroy.Zeno:input_type -> g4.Mota
	0, // 1: g4.Destroy.Zeno:output_type -> g4.Mota
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_g4_proto_init() }
func file_g4_proto_init() {
	if File_g4_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_g4_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mota); i {
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
			RawDescriptor: file_g4_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_g4_proto_goTypes,
		DependencyIndexes: file_g4_proto_depIdxs,
		MessageInfos:      file_g4_proto_msgTypes,
	}.Build()
	File_g4_proto = out.File
	file_g4_proto_rawDesc = nil
	file_g4_proto_goTypes = nil
	file_g4_proto_depIdxs = nil
}
