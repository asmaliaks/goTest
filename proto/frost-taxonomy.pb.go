// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: frost-taxonomy.proto

package app_taxonomy

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

type TaxonomyGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id"`
}

func (x *TaxonomyGetRequest) Reset() {
	*x = TaxonomyGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frost_taxonomy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaxonomyGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxonomyGetRequest) ProtoMessage() {}

func (x *TaxonomyGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_frost_taxonomy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxonomyGetRequest.ProtoReflect.Descriptor instead.
func (*TaxonomyGetRequest) Descriptor() ([]byte, []int) {
	return file_frost_taxonomy_proto_rawDescGZIP(), []int{0}
}

func (x *TaxonomyGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TaxonomyPutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,10,opt,name=success,proto3" json:"success"`
	Message string `protobuf:"bytes,20,opt,name=message,proto3" json:"message"`
}

func (x *TaxonomyPutResponse) Reset() {
	*x = TaxonomyPutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frost_taxonomy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaxonomyPutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxonomyPutResponse) ProtoMessage() {}

func (x *TaxonomyPutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_frost_taxonomy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxonomyPutResponse.ProtoReflect.Descriptor instead.
func (*TaxonomyPutResponse) Descriptor() ([]byte, []int) {
	return file_frost_taxonomy_proto_rawDescGZIP(), []int{1}
}

func (x *TaxonomyPutResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TaxonomyPutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TaxonomyItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,10,opt,name=id,proto3" json:"id"`
	State          string   `protobuf:"bytes,20,opt,name=state,proto3" json:"state"`
	Grade          string   `protobuf:"bytes,30,opt,name=grade,proto3" json:"grade"`
	Name           string   `protobuf:"bytes,40,opt,name=name,proto3" json:"name"`
	SourceDocument string   `protobuf:"bytes,50,opt,name=sourceDocument,proto3" json:"sourceDocument"`
	Parent         string   `protobuf:"bytes,60,opt,name=parent,proto3" json:"parent"`
	ItemsForward   []string `protobuf:"bytes,70,rep,name=itemsForward,proto3" json:"itemsForward"`
	ItemsBack      []string `protobuf:"bytes,80,rep,name=itemsBack,proto3" json:"itemsBack"`
	Description    string   `protobuf:"bytes,90,opt,name=description,proto3" json:"description"`
}

func (x *TaxonomyItem) Reset() {
	*x = TaxonomyItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frost_taxonomy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaxonomyItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxonomyItem) ProtoMessage() {}

func (x *TaxonomyItem) ProtoReflect() protoreflect.Message {
	mi := &file_frost_taxonomy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxonomyItem.ProtoReflect.Descriptor instead.
func (*TaxonomyItem) Descriptor() ([]byte, []int) {
	return file_frost_taxonomy_proto_rawDescGZIP(), []int{2}
}

func (x *TaxonomyItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TaxonomyItem) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *TaxonomyItem) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *TaxonomyItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaxonomyItem) GetSourceDocument() string {
	if x != nil {
		return x.SourceDocument
	}
	return ""
}

func (x *TaxonomyItem) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *TaxonomyItem) GetItemsForward() []string {
	if x != nil {
		return x.ItemsForward
	}
	return nil
}

func (x *TaxonomyItem) GetItemsBack() []string {
	if x != nil {
		return x.ItemsBack
	}
	return nil
}

func (x *TaxonomyItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_frost_taxonomy_proto protoreflect.FileDescriptor

var file_frost_taxonomy_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2d, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x22, 0x24, 0x0a, 0x12, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x82, 0x02, 0x0a, 0x0c, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x18, 0x46, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42,
	0x61, 0x63, 0x6b, 0x18, 0x50, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x61, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x89, 0x01, 0x0a, 0x08, 0x54, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x12, 0x3e, 0x0a, 0x03, 0x70, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x74, 0x61, 0x78,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x1a, 0x1d, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x54, 0x61,
	0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x74, 0x61, 0x78,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x49, 0x74, 0x65, 0x6d,
	0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_frost_taxonomy_proto_rawDescOnce sync.Once
	file_frost_taxonomy_proto_rawDescData = file_frost_taxonomy_proto_rawDesc
)

func file_frost_taxonomy_proto_rawDescGZIP() []byte {
	file_frost_taxonomy_proto_rawDescOnce.Do(func() {
		file_frost_taxonomy_proto_rawDescData = protoimpl.X.CompressGZIP(file_frost_taxonomy_proto_rawDescData)
	})
	return file_frost_taxonomy_proto_rawDescData
}

var file_frost_taxonomy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_frost_taxonomy_proto_goTypes = []interface{}{
	(*TaxonomyGetRequest)(nil),  // 0: taxonomy.TaxonomyGetRequest
	(*TaxonomyPutResponse)(nil), // 1: taxonomy.TaxonomyPutResponse
	(*TaxonomyItem)(nil),        // 2: taxonomy.TaxonomyItem
}
var file_frost_taxonomy_proto_depIdxs = []int32{
	2, // 0: taxonomy.Taxonomy.put:input_type -> taxonomy.TaxonomyItem
	0, // 1: taxonomy.Taxonomy.get:input_type -> taxonomy.TaxonomyGetRequest
	1, // 2: taxonomy.Taxonomy.put:output_type -> taxonomy.TaxonomyPutResponse
	2, // 3: taxonomy.Taxonomy.get:output_type -> taxonomy.TaxonomyItem
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_frost_taxonomy_proto_init() }
func file_frost_taxonomy_proto_init() {
	if File_frost_taxonomy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frost_taxonomy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaxonomyGetRequest); i {
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
		file_frost_taxonomy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaxonomyPutResponse); i {
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
		file_frost_taxonomy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaxonomyItem); i {
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
			RawDescriptor: file_frost_taxonomy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_frost_taxonomy_proto_goTypes,
		DependencyIndexes: file_frost_taxonomy_proto_depIdxs,
		MessageInfos:      file_frost_taxonomy_proto_msgTypes,
	}.Build()
	File_frost_taxonomy_proto = out.File
	file_frost_taxonomy_proto_rawDesc = nil
	file_frost_taxonomy_proto_goTypes = nil
	file_frost_taxonomy_proto_depIdxs = nil
}
