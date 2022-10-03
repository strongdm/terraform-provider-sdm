// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: tags.proto

package v1

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

// Tags defines a custom message that will translate from a slice of
// Pairs into a map[] in each supported language.  See options.proto for
// documentation on the custom message options.
type Tags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs []*Tags_Pair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *Tags) Reset() {
	*x = Tags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tags) ProtoMessage() {}

func (x *Tags) ProtoReflect() protoreflect.Message {
	mi := &file_tags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tags.ProtoReflect.Descriptor instead.
func (*Tags) Descriptor() ([]byte, []int) {
	return file_tags_proto_rawDescGZIP(), []int{0}
}

func (x *Tags) GetPairs() []*Tags_Pair {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tags_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_tags_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_tags_proto_rawDescGZIP(), []int{1}
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Tags_Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tags_Pair) Reset() {
	*x = Tags_Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tags_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tags_Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tags_Pair) ProtoMessage() {}

func (x *Tags_Pair) ProtoReflect() protoreflect.Message {
	mi := &file_tags_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tags_Pair.ProtoReflect.Descriptor instead.
func (*Tags_Pair) Descriptor() ([]byte, []int) {
	return file_tags_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Tags_Pair) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tags_Pair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_tags_proto protoreflect.FileDescriptor

var file_tags_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xee, 0x02, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67,
	0x73, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x1a, 0x4f, 0x0a,
	0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x1d, 0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xd2, 0xf3, 0xb3, 0x07,
	0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3a, 0xef,
	0x01, 0xfa, 0xf8, 0xb3, 0x07, 0xe9, 0x01, 0xca, 0xf3, 0xb3, 0x07, 0xcb, 0x01, 0xea, 0xf3, 0xb3,
	0x07, 0x04, 0x74, 0x61, 0x67, 0x73, 0xf2, 0xf3, 0xb3, 0x07, 0x0a, 0x0a, 0x02, 0x67, 0x6f, 0x12,
	0x04, 0x54, 0x61, 0x67, 0x73, 0xf2, 0xf3, 0xb3, 0x07, 0x12, 0x0a, 0x0a, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x04, 0x54, 0x61, 0x67, 0x73, 0xf2, 0xf3, 0xb3, 0x07,
	0x14, 0x0a, 0x0c, 0x67, 0x6f, 0x5f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x04, 0x54, 0x61, 0x67, 0x73, 0xf2, 0xf3, 0xb3, 0x07, 0x25, 0x0a, 0x04, 0x6a, 0x61, 0x76, 0x61,
	0x12, 0x1d, 0x6a, 0x61, 0x76, 0x61, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x4d, 0x61, 0x70, 0x3c,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2c, 0x20, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3e, 0xf2,
	0xf3, 0xb3, 0x07, 0x1d, 0x0a, 0x12, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x07, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61,
	0x70, 0xf2, 0xf3, 0xb3, 0x07, 0x21, 0x0a, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x12, 0x11, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xfa, 0xf3, 0xb3, 0x07, 0x0c, 0x74, 0x61, 0x67, 0x73,
	0x45, 0x6c, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xd2, 0xf3,
	0xb3, 0x07, 0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x22, 0x53, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3,
	0x07, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x0a, 0xfa, 0xf8, 0xb3, 0x07, 0x05,
	0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x42, 0x60, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72,
	0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x69,
	0x6e, 0x67, 0x42, 0x0c, 0x54, 0x61, 0x67, 0x73, 0x50, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72,
	0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2d, 0x73,
	0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tags_proto_rawDescOnce sync.Once
	file_tags_proto_rawDescData = file_tags_proto_rawDesc
)

func file_tags_proto_rawDescGZIP() []byte {
	file_tags_proto_rawDescOnce.Do(func() {
		file_tags_proto_rawDescData = protoimpl.X.CompressGZIP(file_tags_proto_rawDescData)
	})
	return file_tags_proto_rawDescData
}

var file_tags_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tags_proto_goTypes = []interface{}{
	(*Tags)(nil),      // 0: v1.Tags
	(*Tag)(nil),       // 1: v1.Tag
	(*Tags_Pair)(nil), // 2: v1.Tags.Pair
}
var file_tags_proto_depIdxs = []int32{
	2, // 0: v1.Tags.pairs:type_name -> v1.Tags.Pair
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tags_proto_init() }
func file_tags_proto_init() {
	if File_tags_proto != nil {
		return
	}
	file_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tags); i {
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
		file_tags_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_tags_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tags_Pair); i {
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
			RawDescriptor: file_tags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tags_proto_goTypes,
		DependencyIndexes: file_tags_proto_depIdxs,
		MessageInfos:      file_tags_proto_msgTypes,
	}.Build()
	File_tags_proto = out.File
	file_tags_proto_rawDesc = nil
	file_tags_proto_goTypes = nil
	file_tags_proto_depIdxs = nil
}