// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: spaces.proto

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

type Vec3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vec3) Reset() {
	*x = Vec3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vec3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vec3) ProtoMessage() {}

func (x *Vec3) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vec3.ProtoReflect.Descriptor instead.
func (*Vec3) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{0}
}

func (x *Vec3) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vec3) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vec3) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Geometry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size     *Vec3 `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	Position *Vec3 `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Geometry) Reset() {
	*x = Geometry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Geometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Geometry) ProtoMessage() {}

func (x *Geometry) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Geometry.ProtoReflect.Descriptor instead.
func (*Geometry) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{1}
}

func (x *Geometry) GetSize() *Vec3 {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *Geometry) GetPosition() *Vec3 {
	if x != nil {
		return x.Position
	}
	return nil
}

type Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color string `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Material) Reset() {
	*x = Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Material) ProtoMessage() {}

func (x *Material) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Material.ProtoReflect.Descriptor instead.
func (*Material) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{2}
}

func (x *Material) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type Mesh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Geometry *Geometry `protobuf:"bytes,1,opt,name=geometry,proto3" json:"geometry,omitempty"`
	Material *Material `protobuf:"bytes,2,opt,name=material,proto3" json:"material,omitempty"`
}

func (x *Mesh) Reset() {
	*x = Mesh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mesh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mesh) ProtoMessage() {}

func (x *Mesh) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mesh.ProtoReflect.Descriptor instead.
func (*Mesh) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{3}
}

func (x *Mesh) GetGeometry() *Geometry {
	if x != nil {
		return x.Geometry
	}
	return nil
}

func (x *Mesh) GetMaterial() *Material {
	if x != nil {
		return x.Material
	}
	return nil
}

type MeshSequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meshes []*Mesh `protobuf:"bytes,1,rep,name=meshes,proto3" json:"meshes,omitempty"`
	Id     string  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MeshSequence) Reset() {
	*x = MeshSequence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshSequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshSequence) ProtoMessage() {}

func (x *MeshSequence) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshSequence.ProtoReflect.Descriptor instead.
func (*MeshSequence) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{4}
}

func (x *MeshSequence) GetMeshes() []*Mesh {
	if x != nil {
		return x.Meshes
	}
	return nil
}

func (x *MeshSequence) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_spaces_proto protoreflect.FileDescriptor

var file_spaces_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x30,
	0x0a, 0x04, 0x56, 0x65, 0x63, 0x33, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a,
	0x22, 0x66, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x33,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x72, 0x0a, 0x04, 0x4d, 0x65,
	0x73, 0x68, 0x12, 0x34, 0x0a, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x4c,
	0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x68, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x06, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x4d, 0x65, 0x73, 0x68, 0x52, 0x06, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaces_proto_rawDescOnce sync.Once
	file_spaces_proto_rawDescData = file_spaces_proto_rawDesc
)

func file_spaces_proto_rawDescGZIP() []byte {
	file_spaces_proto_rawDescOnce.Do(func() {
		file_spaces_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaces_proto_rawDescData)
	})
	return file_spaces_proto_rawDescData
}

var file_spaces_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaces_proto_goTypes = []interface{}{
	(*Vec3)(nil),         // 0: spaces_package.Vec3
	(*Geometry)(nil),     // 1: spaces_package.Geometry
	(*Material)(nil),     // 2: spaces_package.Material
	(*Mesh)(nil),         // 3: spaces_package.Mesh
	(*MeshSequence)(nil), // 4: spaces_package.MeshSequence
}
var file_spaces_proto_depIdxs = []int32{
	0, // 0: spaces_package.Geometry.size:type_name -> spaces_package.Vec3
	0, // 1: spaces_package.Geometry.position:type_name -> spaces_package.Vec3
	1, // 2: spaces_package.Mesh.geometry:type_name -> spaces_package.Geometry
	2, // 3: spaces_package.Mesh.material:type_name -> spaces_package.Material
	3, // 4: spaces_package.MeshSequence.meshes:type_name -> spaces_package.Mesh
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaces_proto_init() }
func file_spaces_proto_init() {
	if File_spaces_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vec3); i {
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
		file_spaces_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Geometry); i {
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
		file_spaces_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Material); i {
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
		file_spaces_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mesh); i {
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
		file_spaces_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshSequence); i {
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
			RawDescriptor: file_spaces_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spaces_proto_goTypes,
		DependencyIndexes: file_spaces_proto_depIdxs,
		MessageInfos:      file_spaces_proto_msgTypes,
	}.Build()
	File_spaces_proto = out.File
	file_spaces_proto_rawDesc = nil
	file_spaces_proto_goTypes = nil
	file_spaces_proto_depIdxs = nil
}