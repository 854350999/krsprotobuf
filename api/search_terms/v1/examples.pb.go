// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: search_terms/v1/examples.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetExamplesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *GetExamplesReq) Reset() {
	*x = GetExamplesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_terms_v1_examples_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExamplesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExamplesReq) ProtoMessage() {}

func (x *GetExamplesReq) ProtoReflect() protoreflect.Message {
	mi := &file_search_terms_v1_examples_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExamplesReq.ProtoReflect.Descriptor instead.
func (*GetExamplesReq) Descriptor() ([]byte, []int) {
	return file_search_terms_v1_examples_proto_rawDescGZIP(), []int{0}
}

func (x *GetExamplesReq) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type GetExamplesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchTerms []string        `protobuf:"bytes,1,rep,name=searchTerms,proto3" json:"searchTerms,omitempty"`
	Test        *structpb.Value `protobuf:"bytes,2,opt,name=test,proto3" json:"test,omitempty"`
}

func (x *GetExamplesReply) Reset() {
	*x = GetExamplesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_terms_v1_examples_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExamplesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExamplesReply) ProtoMessage() {}

func (x *GetExamplesReply) ProtoReflect() protoreflect.Message {
	mi := &file_search_terms_v1_examples_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExamplesReply.ProtoReflect.Descriptor instead.
func (*GetExamplesReply) Descriptor() ([]byte, []int) {
	return file_search_terms_v1_examples_proto_rawDescGZIP(), []int{1}
}

func (x *GetExamplesReply) GetSearchTerms() []string {
	if x != nil {
		return x.SearchTerms
	}
	return nil
}

func (x *GetExamplesReply) GetTest() *structpb.Value {
	if x != nil {
		return x.Test
	}
	return nil
}

var File_search_terms_v1_examples_proto protoreflect.FileDescriptor

var file_search_terms_v1_examples_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x60, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x73,
	0x12, 0x2a, 0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x32, 0x3d, 0x0a, 0x08,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x18, 0x5a, 0x16, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_terms_v1_examples_proto_rawDescOnce sync.Once
	file_search_terms_v1_examples_proto_rawDescData = file_search_terms_v1_examples_proto_rawDesc
)

func file_search_terms_v1_examples_proto_rawDescGZIP() []byte {
	file_search_terms_v1_examples_proto_rawDescOnce.Do(func() {
		file_search_terms_v1_examples_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_terms_v1_examples_proto_rawDescData)
	})
	return file_search_terms_v1_examples_proto_rawDescData
}

var file_search_terms_v1_examples_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_search_terms_v1_examples_proto_goTypes = []interface{}{
	(*GetExamplesReq)(nil),   // 0: GetExamplesReq
	(*GetExamplesReply)(nil), // 1: GetExamplesReply
	(*structpb.Value)(nil),   // 2: google.protobuf.Value
}
var file_search_terms_v1_examples_proto_depIdxs = []int32{
	2, // 0: GetExamplesReply.test:type_name -> google.protobuf.Value
	0, // 1: Examples.GetExamples:input_type -> GetExamplesReq
	1, // 2: Examples.GetExamples:output_type -> GetExamplesReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_search_terms_v1_examples_proto_init() }
func file_search_terms_v1_examples_proto_init() {
	if File_search_terms_v1_examples_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_terms_v1_examples_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExamplesReq); i {
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
		file_search_terms_v1_examples_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExamplesReply); i {
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
			RawDescriptor: file_search_terms_v1_examples_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_terms_v1_examples_proto_goTypes,
		DependencyIndexes: file_search_terms_v1_examples_proto_depIdxs,
		MessageInfos:      file_search_terms_v1_examples_proto_msgTypes,
	}.Build()
	File_search_terms_v1_examples_proto = out.File
	file_search_terms_v1_examples_proto_rawDesc = nil
	file_search_terms_v1_examples_proto_goTypes = nil
	file_search_terms_v1_examples_proto_depIdxs = nil
}
