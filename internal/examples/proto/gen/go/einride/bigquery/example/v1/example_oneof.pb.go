// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: einride/bigquery/example/v1/example_oneof.proto

package examplev1

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

type ExampleOneof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OneofFields_1:
	//
	//	*ExampleOneof_OneofEmptyMessage_1
	//	*ExampleOneof_OneofBool_1
	OneofFields_1 isExampleOneof_OneofFields_1 `protobuf_oneof:"oneof_fields_1"`
	// Types that are assignable to OneofFields_2:
	//
	//	*ExampleOneof_OneofEmptyMessage_2
	//	*ExampleOneof_OneofMessage
	OneofFields_2 isExampleOneof_OneofFields_2 `protobuf_oneof:"oneof_fields_2"`
}

func (x *ExampleOneof) Reset() {
	*x = ExampleOneof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_example_v1_example_oneof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleOneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleOneof) ProtoMessage() {}

func (x *ExampleOneof) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_oneof_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleOneof.ProtoReflect.Descriptor instead.
func (*ExampleOneof) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_oneof_proto_rawDescGZIP(), []int{0}
}

func (m *ExampleOneof) GetOneofFields_1() isExampleOneof_OneofFields_1 {
	if m != nil {
		return m.OneofFields_1
	}
	return nil
}

func (x *ExampleOneof) GetOneofEmptyMessage_1() *ExampleOneof_EmptyMessage {
	if x, ok := x.GetOneofFields_1().(*ExampleOneof_OneofEmptyMessage_1); ok {
		return x.OneofEmptyMessage_1
	}
	return nil
}

func (x *ExampleOneof) GetOneofBool_1() bool {
	if x, ok := x.GetOneofFields_1().(*ExampleOneof_OneofBool_1); ok {
		return x.OneofBool_1
	}
	return false
}

func (m *ExampleOneof) GetOneofFields_2() isExampleOneof_OneofFields_2 {
	if m != nil {
		return m.OneofFields_2
	}
	return nil
}

func (x *ExampleOneof) GetOneofEmptyMessage_2() *ExampleOneof_EmptyMessage {
	if x, ok := x.GetOneofFields_2().(*ExampleOneof_OneofEmptyMessage_2); ok {
		return x.OneofEmptyMessage_2
	}
	return nil
}

func (x *ExampleOneof) GetOneofMessage() *ExampleOneof_Message {
	if x, ok := x.GetOneofFields_2().(*ExampleOneof_OneofMessage); ok {
		return x.OneofMessage
	}
	return nil
}

type isExampleOneof_OneofFields_1 interface {
	isExampleOneof_OneofFields_1()
}

type ExampleOneof_OneofEmptyMessage_1 struct {
	OneofEmptyMessage_1 *ExampleOneof_EmptyMessage `protobuf:"bytes,1,opt,name=oneof_empty_message_1,json=oneofEmptyMessage1,proto3,oneof"`
}

type ExampleOneof_OneofBool_1 struct {
	OneofBool_1 bool `protobuf:"varint,2,opt,name=oneof_bool_1,json=oneofBool1,proto3,oneof"`
}

func (*ExampleOneof_OneofEmptyMessage_1) isExampleOneof_OneofFields_1() {}

func (*ExampleOneof_OneofBool_1) isExampleOneof_OneofFields_1() {}

type isExampleOneof_OneofFields_2 interface {
	isExampleOneof_OneofFields_2()
}

type ExampleOneof_OneofEmptyMessage_2 struct {
	OneofEmptyMessage_2 *ExampleOneof_EmptyMessage `protobuf:"bytes,3,opt,name=oneof_empty_message_2,json=oneofEmptyMessage2,proto3,oneof"`
}

type ExampleOneof_OneofMessage struct {
	OneofMessage *ExampleOneof_Message `protobuf:"bytes,4,opt,name=oneof_message,json=oneofMessage,proto3,oneof"`
}

func (*ExampleOneof_OneofEmptyMessage_2) isExampleOneof_OneofFields_2() {}

func (*ExampleOneof_OneofMessage) isExampleOneof_OneofFields_2() {}

type ExampleOneof_EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExampleOneof_EmptyMessage) Reset() {
	*x = ExampleOneof_EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_example_v1_example_oneof_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleOneof_EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleOneof_EmptyMessage) ProtoMessage() {}

func (x *ExampleOneof_EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_oneof_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleOneof_EmptyMessage.ProtoReflect.Descriptor instead.
func (*ExampleOneof_EmptyMessage) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_oneof_proto_rawDescGZIP(), []int{0, 0}
}

type ExampleOneof_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
}

func (x *ExampleOneof_Message) Reset() {
	*x = ExampleOneof_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_example_v1_example_oneof_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleOneof_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleOneof_Message) ProtoMessage() {}

func (x *ExampleOneof_Message) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_oneof_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleOneof_Message.ProtoReflect.Descriptor instead.
func (*ExampleOneof_Message) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_oneof_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ExampleOneof_Message) GetStringValue() string {
	if x != nil {
		return x.StringValue
	}
	return ""
}

var File_einride_bigquery_example_v1_example_oneof_proto protoreflect.FileDescriptor

var file_einride_bigquery_example_v1_example_oneof_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xc8,
	0x03, 0x0a, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12,
	0x6b, 0x0a, 0x15, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x12, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x22, 0x0a, 0x0c,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x31, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x6f, 0x6f, 0x6c, 0x31,
	0x12, 0x6b, 0x0a, 0x15, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x01, 0x52, 0x12, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x12, 0x58, 0x0a,
	0x0d, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x01, 0x52, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2c, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x31, 0x42, 0x10, 0x0a, 0x0e, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x32, 0x42, 0xa8, 0x02, 0x0a, 0x1f, 0x63, 0x6f,
	0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x63, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74,
	0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x42, 0x45, 0xaa, 0x02, 0x1b,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27, 0x45, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x42,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_example_v1_example_oneof_proto_rawDescOnce sync.Once
	file_einride_bigquery_example_v1_example_oneof_proto_rawDescData = file_einride_bigquery_example_v1_example_oneof_proto_rawDesc
)

func file_einride_bigquery_example_v1_example_oneof_proto_rawDescGZIP() []byte {
	file_einride_bigquery_example_v1_example_oneof_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_example_v1_example_oneof_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_example_v1_example_oneof_proto_rawDescData)
	})
	return file_einride_bigquery_example_v1_example_oneof_proto_rawDescData
}

var file_einride_bigquery_example_v1_example_oneof_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_einride_bigquery_example_v1_example_oneof_proto_goTypes = []interface{}{
	(*ExampleOneof)(nil),              // 0: einride.bigquery.example.v1.ExampleOneof
	(*ExampleOneof_EmptyMessage)(nil), // 1: einride.bigquery.example.v1.ExampleOneof.EmptyMessage
	(*ExampleOneof_Message)(nil),      // 2: einride.bigquery.example.v1.ExampleOneof.Message
}
var file_einride_bigquery_example_v1_example_oneof_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.example.v1.ExampleOneof.oneof_empty_message_1:type_name -> einride.bigquery.example.v1.ExampleOneof.EmptyMessage
	1, // 1: einride.bigquery.example.v1.ExampleOneof.oneof_empty_message_2:type_name -> einride.bigquery.example.v1.ExampleOneof.EmptyMessage
	2, // 2: einride.bigquery.example.v1.ExampleOneof.oneof_message:type_name -> einride.bigquery.example.v1.ExampleOneof.Message
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_einride_bigquery_example_v1_example_oneof_proto_init() }
func file_einride_bigquery_example_v1_example_oneof_proto_init() {
	if File_einride_bigquery_example_v1_example_oneof_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_bigquery_example_v1_example_oneof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleOneof); i {
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
		file_einride_bigquery_example_v1_example_oneof_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleOneof_EmptyMessage); i {
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
		file_einride_bigquery_example_v1_example_oneof_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleOneof_Message); i {
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
	file_einride_bigquery_example_v1_example_oneof_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ExampleOneof_OneofEmptyMessage_1)(nil),
		(*ExampleOneof_OneofBool_1)(nil),
		(*ExampleOneof_OneofEmptyMessage_2)(nil),
		(*ExampleOneof_OneofMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_einride_bigquery_example_v1_example_oneof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_example_v1_example_oneof_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_example_v1_example_oneof_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_example_v1_example_oneof_proto_msgTypes,
	}.Build()
	File_einride_bigquery_example_v1_example_oneof_proto = out.File
	file_einride_bigquery_example_v1_example_oneof_proto_rawDesc = nil
	file_einride_bigquery_example_v1_example_oneof_proto_goTypes = nil
	file_einride_bigquery_example_v1_example_oneof_proto_depIdxs = nil
}
