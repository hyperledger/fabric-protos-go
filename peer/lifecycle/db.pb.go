// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.1
// source: peer/lifecycle/db.proto

package lifecycle

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

// StateMetadata describes the keys in a namespace.  It is necessary because
// in collections, range scans are not possible during transactions which
// write.  Therefore we must track the keys in our namespace ourselves.
type StateMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datatype string   `protobuf:"bytes,1,opt,name=datatype,proto3" json:"datatype,omitempty"`
	Fields   []string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *StateMetadata) Reset() {
	*x = StateMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_lifecycle_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateMetadata) ProtoMessage() {}

func (x *StateMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_peer_lifecycle_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateMetadata.ProtoReflect.Descriptor instead.
func (*StateMetadata) Descriptor() ([]byte, []int) {
	return file_peer_lifecycle_db_proto_rawDescGZIP(), []int{0}
}

func (x *StateMetadata) GetDatatype() string {
	if x != nil {
		return x.Datatype
	}
	return ""
}

func (x *StateMetadata) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

// StateData encodes a particular field of a datatype
type StateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*StateData_Int64
	//	*StateData_Bytes
	//	*StateData_String_
	Type isStateData_Type `protobuf_oneof:"Type"`
}

func (x *StateData) Reset() {
	*x = StateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_lifecycle_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateData) ProtoMessage() {}

func (x *StateData) ProtoReflect() protoreflect.Message {
	mi := &file_peer_lifecycle_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateData.ProtoReflect.Descriptor instead.
func (*StateData) Descriptor() ([]byte, []int) {
	return file_peer_lifecycle_db_proto_rawDescGZIP(), []int{1}
}

func (m *StateData) GetType() isStateData_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *StateData) GetInt64() int64 {
	if x, ok := x.GetType().(*StateData_Int64); ok {
		return x.Int64
	}
	return 0
}

func (x *StateData) GetBytes() []byte {
	if x, ok := x.GetType().(*StateData_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (x *StateData) GetString_() string {
	if x, ok := x.GetType().(*StateData_String_); ok {
		return x.String_
	}
	return ""
}

type isStateData_Type interface {
	isStateData_Type()
}

type StateData_Int64 struct {
	Int64 int64 `protobuf:"varint,1,opt,name=Int64,proto3,oneof"`
}

type StateData_Bytes struct {
	Bytes []byte `protobuf:"bytes,2,opt,name=Bytes,proto3,oneof"`
}

type StateData_String_ struct {
	String_ string `protobuf:"bytes,3,opt,name=String,proto3,oneof"`
}

func (*StateData_Int64) isStateData_Type() {}

func (*StateData_Bytes) isStateData_Type() {}

func (*StateData_String_) isStateData_Type() {}

var File_peer_lifecycle_db_proto protoreflect.FileDescriptor

var file_peer_lifecycle_db_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x2f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x22, 0x43, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x5d, 0x0a, 0x09, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x16,
	0x0a, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x42, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x42, 0x66, 0x0a, 0x2c, 0x6f, 0x72, 0x67, 0x2e,
	0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x6c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67,
	0x6f, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_lifecycle_db_proto_rawDescOnce sync.Once
	file_peer_lifecycle_db_proto_rawDescData = file_peer_lifecycle_db_proto_rawDesc
)

func file_peer_lifecycle_db_proto_rawDescGZIP() []byte {
	file_peer_lifecycle_db_proto_rawDescOnce.Do(func() {
		file_peer_lifecycle_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_lifecycle_db_proto_rawDescData)
	})
	return file_peer_lifecycle_db_proto_rawDescData
}

var file_peer_lifecycle_db_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_peer_lifecycle_db_proto_goTypes = []interface{}{
	(*StateMetadata)(nil), // 0: lifecycle.StateMetadata
	(*StateData)(nil),     // 1: lifecycle.StateData
}
var file_peer_lifecycle_db_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_peer_lifecycle_db_proto_init() }
func file_peer_lifecycle_db_proto_init() {
	if File_peer_lifecycle_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_lifecycle_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateMetadata); i {
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
		file_peer_lifecycle_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateData); i {
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
	file_peer_lifecycle_db_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StateData_Int64)(nil),
		(*StateData_Bytes)(nil),
		(*StateData_String_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_peer_lifecycle_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_peer_lifecycle_db_proto_goTypes,
		DependencyIndexes: file_peer_lifecycle_db_proto_depIdxs,
		MessageInfos:      file_peer_lifecycle_db_proto_msgTypes,
	}.Build()
	File_peer_lifecycle_db_proto = out.File
	file_peer_lifecycle_db_proto_rawDesc = nil
	file_peer_lifecycle_db_proto_goTypes = nil
	file_peer_lifecycle_db_proto_depIdxs = nil
}
