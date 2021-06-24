// Copyright 2019 Google LLC.
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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/api/expr/v1beta1/source.proto

package expr

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// Source information collected at parse time.
type SourceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The location name. All position information attached to an expression is
	// relative to this location.
	//
	// The location could be a file, UI element, or similar. For example,
	// `acme/app/AnvilPolicy.cel`.
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// Monotonically increasing list of character offsets where newlines appear.
	//
	// The line number of a given position is the index `i` where for a given
	// `id` the `line_offsets[i] < id_positions[id] < line_offsets[i+1]`. The
	// column may be derivd from `id_positions[id] - line_offsets[i]`.
	LineOffsets []int32 `protobuf:"varint,3,rep,packed,name=line_offsets,json=lineOffsets,proto3" json:"line_offsets,omitempty"`
	// A map from the parse node id (e.g. `Expr.id`) to the character offset
	// within source.
	Positions map[int32]int32 `protobuf:"bytes,4,rep,name=positions,proto3" json:"positions,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *SourceInfo) Reset() {
	*x = SourceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_expr_v1beta1_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceInfo) ProtoMessage() {}

func (x *SourceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_expr_v1beta1_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceInfo.ProtoReflect.Descriptor instead.
func (*SourceInfo) Descriptor() ([]byte, []int) {
	return file_google_api_expr_v1beta1_source_proto_rawDescGZIP(), []int{0}
}

func (x *SourceInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *SourceInfo) GetLineOffsets() []int32 {
	if x != nil {
		return x.LineOffsets
	}
	return nil
}

func (x *SourceInfo) GetPositions() map[int32]int32 {
	if x != nil {
		return x.Positions
	}
	return nil
}

// A specific position in source.
type SourcePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The soucre location name (e.g. file name).
	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// The character offset.
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// The 1-based index of the starting line in the source text
	// where the issue occurs, or 0 if unknown.
	Line int32 `protobuf:"varint,3,opt,name=line,proto3" json:"line,omitempty"`
	// The 0-based index of the starting position within the line of source text
	// where the issue occurs.  Only meaningful if line is nonzer..
	Column int32 `protobuf:"varint,4,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *SourcePosition) Reset() {
	*x = SourcePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_expr_v1beta1_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourcePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourcePosition) ProtoMessage() {}

func (x *SourcePosition) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_expr_v1beta1_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourcePosition.ProtoReflect.Descriptor instead.
func (*SourcePosition) Descriptor() ([]byte, []int) {
	return file_google_api_expr_v1beta1_source_proto_rawDescGZIP(), []int{1}
}

func (x *SourcePosition) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *SourcePosition) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SourcePosition) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *SourcePosition) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

var File_google_api_expr_v1beta1_source_proto protoreflect.FileDescriptor

var file_google_api_expr_v1beta1_source_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x70,
	0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22,
	0xdb, 0x01, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x73, 0x12, 0x50, 0x0a,
	0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78,
	0x70, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x3c, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x70, 0x0a,
	0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x42,
	0x6c, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0b,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x65, 0x78, 0x70, 0x72, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_expr_v1beta1_source_proto_rawDescOnce sync.Once
	file_google_api_expr_v1beta1_source_proto_rawDescData = file_google_api_expr_v1beta1_source_proto_rawDesc
)

func file_google_api_expr_v1beta1_source_proto_rawDescGZIP() []byte {
	file_google_api_expr_v1beta1_source_proto_rawDescOnce.Do(func() {
		file_google_api_expr_v1beta1_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_expr_v1beta1_source_proto_rawDescData)
	})
	return file_google_api_expr_v1beta1_source_proto_rawDescData
}

var file_google_api_expr_v1beta1_source_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_api_expr_v1beta1_source_proto_goTypes = []interface{}{
	(*SourceInfo)(nil),     // 0: google.api.expr.v1beta1.SourceInfo
	(*SourcePosition)(nil), // 1: google.api.expr.v1beta1.SourcePosition
	nil,                    // 2: google.api.expr.v1beta1.SourceInfo.PositionsEntry
}
var file_google_api_expr_v1beta1_source_proto_depIdxs = []int32{
	2, // 0: google.api.expr.v1beta1.SourceInfo.positions:type_name -> google.api.expr.v1beta1.SourceInfo.PositionsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_api_expr_v1beta1_source_proto_init() }
func file_google_api_expr_v1beta1_source_proto_init() {
	if File_google_api_expr_v1beta1_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_expr_v1beta1_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceInfo); i {
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
		file_google_api_expr_v1beta1_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourcePosition); i {
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
			RawDescriptor: file_google_api_expr_v1beta1_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_expr_v1beta1_source_proto_goTypes,
		DependencyIndexes: file_google_api_expr_v1beta1_source_proto_depIdxs,
		MessageInfos:      file_google_api_expr_v1beta1_source_proto_msgTypes,
	}.Build()
	File_google_api_expr_v1beta1_source_proto = out.File
	file_google_api_expr_v1beta1_source_proto_rawDesc = nil
	file_google_api_expr_v1beta1_source_proto_goTypes = nil
	file_google_api_expr_v1beta1_source_proto_depIdxs = nil
}
