// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: libnettop.proto

package nettop

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

type Kernel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Release      string `protobuf:"bytes,1,opt,name=release,proto3" json:"release,omitempty"`
	Version      string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Architecture string `protobuf:"bytes,3,opt,name=architecture,proto3" json:"architecture,omitempty"`
}

func (x *Kernel) Reset() {
	*x = Kernel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libnettop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kernel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kernel) ProtoMessage() {}

func (x *Kernel) ProtoReflect() protoreflect.Message {
	mi := &file_libnettop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kernel.ProtoReflect.Descriptor instead.
func (*Kernel) Descriptor() ([]byte, []int) {
	return file_libnettop_proto_rawDescGZIP(), []int{0}
}

func (x *Kernel) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *Kernel) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Kernel) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

type CriMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version        string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	RuntimeName    string `protobuf:"bytes,2,opt,name=runtime_name,json=runtimeName,proto3" json:"runtime_name,omitempty"`
	RuntimeVersion string `protobuf:"bytes,3,opt,name=runtime_version,json=runtimeVersion,proto3" json:"runtime_version,omitempty"`
	RuntimeSock    string `protobuf:"bytes,4,opt,name=runtime_sock,json=runtimeSock,proto3" json:"runtime_sock,omitempty"`
}

func (x *CriMeta) Reset() {
	*x = CriMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libnettop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CriMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CriMeta) ProtoMessage() {}

func (x *CriMeta) ProtoReflect() protoreflect.Message {
	mi := &file_libnettop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CriMeta.ProtoReflect.Descriptor instead.
func (*CriMeta) Descriptor() ([]byte, []int) {
	return file_libnettop_proto_rawDescGZIP(), []int{1}
}

func (x *CriMeta) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CriMeta) GetRuntimeName() string {
	if x != nil {
		return x.RuntimeName
	}
	return ""
}

func (x *CriMeta) GetRuntimeVersion() string {
	if x != nil {
		return x.RuntimeVersion
	}
	return ""
}

func (x *CriMeta) GetRuntimeSock() string {
	if x != nil {
		return x.RuntimeSock
	}
	return ""
}

type NodeMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Kernel   *Kernel  `protobuf:"bytes,2,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Crimeta  *CriMeta `protobuf:"bytes,3,opt,name=crimeta,proto3" json:"crimeta,omitempty"`
}

func (x *NodeMeta) Reset() {
	*x = NodeMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libnettop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMeta) ProtoMessage() {}

func (x *NodeMeta) ProtoReflect() protoreflect.Message {
	mi := &file_libnettop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMeta.ProtoReflect.Descriptor instead.
func (*NodeMeta) Descriptor() ([]byte, []int) {
	return file_libnettop_proto_rawDescGZIP(), []int{2}
}

func (x *NodeMeta) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *NodeMeta) GetKernel() *Kernel {
	if x != nil {
		return x.Kernel
	}
	return nil
}

func (x *NodeMeta) GetCrimeta() *CriMeta {
	if x != nil {
		return x.Crimeta
	}
	return nil
}

var File_libnettop_proto protoreflect.FileDescriptor

var file_libnettop_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x69, 0x62, 0x6e, 0x65, 0x74, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x6e, 0x65, 0x74, 0x74, 0x6f, 0x70, 0x22, 0x60, 0x0a, 0x06, 0x4b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x07,
	0x43, 0x72, 0x69, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x6f, 0x63, 0x6b,
	0x22, 0x7a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x6b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x65, 0x74, 0x74,
	0x6f, 0x70, 0x2e, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x52, 0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x65, 0x74, 0x74, 0x6f, 0x70, 0x2e, 0x43, 0x72, 0x69, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x07, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x3b, 0x6e, 0x65, 0x74, 0x74, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_libnettop_proto_rawDescOnce sync.Once
	file_libnettop_proto_rawDescData = file_libnettop_proto_rawDesc
)

func file_libnettop_proto_rawDescGZIP() []byte {
	file_libnettop_proto_rawDescOnce.Do(func() {
		file_libnettop_proto_rawDescData = protoimpl.X.CompressGZIP(file_libnettop_proto_rawDescData)
	})
	return file_libnettop_proto_rawDescData
}

var file_libnettop_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_libnettop_proto_goTypes = []interface{}{
	(*Kernel)(nil),   // 0: nettop.Kernel
	(*CriMeta)(nil),  // 1: nettop.CriMeta
	(*NodeMeta)(nil), // 2: nettop.NodeMeta
}
var file_libnettop_proto_depIdxs = []int32{
	0, // 0: nettop.NodeMeta.kernel:type_name -> nettop.Kernel
	1, // 1: nettop.NodeMeta.crimeta:type_name -> nettop.CriMeta
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_libnettop_proto_init() }
func file_libnettop_proto_init() {
	if File_libnettop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_libnettop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kernel); i {
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
		file_libnettop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CriMeta); i {
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
		file_libnettop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMeta); i {
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
			RawDescriptor: file_libnettop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_libnettop_proto_goTypes,
		DependencyIndexes: file_libnettop_proto_depIdxs,
		MessageInfos:      file_libnettop_proto_msgTypes,
	}.Build()
	File_libnettop_proto = out.File
	file_libnettop_proto_rawDesc = nil
	file_libnettop_proto_goTypes = nil
	file_libnettop_proto_depIdxs = nil
}
