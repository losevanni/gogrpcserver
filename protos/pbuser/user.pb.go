// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.6.1
// source: protos/pbuser/user.proto

package protos

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

type ReqJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pw string `protobuf:"bytes,2,opt,name=pw,proto3" json:"pw,omitempty"`
}

func (x *ReqJoin) Reset() {
	*x = ReqJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pbuser_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqJoin) ProtoMessage() {}

func (x *ReqJoin) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pbuser_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqJoin.ProtoReflect.Descriptor instead.
func (*ReqJoin) Descriptor() ([]byte, []int) {
	return file_protos_pbuser_user_proto_rawDescGZIP(), []int{0}
}

func (x *ReqJoin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReqJoin) GetPw() string {
	if x != nil {
		return x.Pw
	}
	return ""
}

type ResJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State bool `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ResJoin) Reset() {
	*x = ResJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pbuser_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResJoin) ProtoMessage() {}

func (x *ResJoin) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pbuser_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResJoin.ProtoReflect.Descriptor instead.
func (*ResJoin) Descriptor() ([]byte, []int) {
	return file_protos_pbuser_user_proto_rawDescGZIP(), []int{1}
}

func (x *ResJoin) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

type ReqLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pw string `protobuf:"bytes,2,opt,name=pw,proto3" json:"pw,omitempty"`
}

func (x *ReqLogin) Reset() {
	*x = ReqLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pbuser_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqLogin) ProtoMessage() {}

func (x *ReqLogin) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pbuser_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqLogin.ProtoReflect.Descriptor instead.
func (*ReqLogin) Descriptor() ([]byte, []int) {
	return file_protos_pbuser_user_proto_rawDescGZIP(), []int{2}
}

func (x *ReqLogin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReqLogin) GetPw() string {
	if x != nil {
		return x.Pw
	}
	return ""
}

type ResLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State bool `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ResLogin) Reset() {
	*x = ResLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pbuser_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResLogin) ProtoMessage() {}

func (x *ResLogin) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pbuser_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResLogin.ProtoReflect.Descriptor instead.
func (*ResLogin) Descriptor() ([]byte, []int) {
	return file_protos_pbuser_user_proto_rawDescGZIP(), []int{3}
}

func (x *ResLogin) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

type ReqInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReqInfo) Reset() {
	*x = ReqInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pbuser_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqInfo) ProtoMessage() {}

func (x *ReqInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pbuser_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqInfo.ProtoReflect.Descriptor instead.
func (*ReqInfo) Descriptor() ([]byte, []int) {
	return file_protos_pbuser_user_proto_rawDescGZIP(), []int{4}
}

func (x *ReqInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ResInfo) Reset() {
	*x = ResInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pbuser_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResInfo) ProtoMessage() {}

func (x *ResInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pbuser_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResInfo.ProtoReflect.Descriptor instead.
func (*ResInfo) Descriptor() ([]byte, []int) {
	return file_protos_pbuser_user_proto_rawDescGZIP(), []int{5}
}

func (x *ResInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_protos_pbuser_user_proto protoreflect.FileDescriptor

var file_protos_pbuser_user_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x62, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x70, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x77, 0x22, 0x1f, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x2a,
	0x0a, 0x08, 0x72, 0x65, 0x71, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x77,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x77, 0x22, 0x20, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x19, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x19, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0x37, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4a, 0x6f, 0x69,
	0x6e, 0x12, 0x28, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x72, 0x65, 0x71, 0x4a, 0x6f, 0x69, 0x6e, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x4a, 0x6f, 0x69, 0x6e, 0x32, 0x3b, 0x0a, 0x0c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65,
	0x71, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x72, 0x65, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x32, 0x37, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0f, 0x2e, 0x70, 0x62, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_pbuser_user_proto_rawDescOnce sync.Once
	file_protos_pbuser_user_proto_rawDescData = file_protos_pbuser_user_proto_rawDesc
)

func file_protos_pbuser_user_proto_rawDescGZIP() []byte {
	file_protos_pbuser_user_proto_rawDescOnce.Do(func() {
		file_protos_pbuser_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_pbuser_user_proto_rawDescData)
	})
	return file_protos_pbuser_user_proto_rawDescData
}

var file_protos_pbuser_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_pbuser_user_proto_goTypes = []interface{}{
	(*ReqJoin)(nil),  // 0: pbuser.reqJoin
	(*ResJoin)(nil),  // 1: pbuser.resJoin
	(*ReqLogin)(nil), // 2: pbuser.reqLogin
	(*ResLogin)(nil), // 3: pbuser.resLogin
	(*ReqInfo)(nil),  // 4: pbuser.reqInfo
	(*ResInfo)(nil),  // 5: pbuser.resInfo
}
var file_protos_pbuser_user_proto_depIdxs = []int32{
	0, // 0: pbuser.ServiceJoin.Join:input_type -> pbuser.reqJoin
	2, // 1: pbuser.ServiceLogin.Login:input_type -> pbuser.reqLogin
	4, // 2: pbuser.ServiceInfo.Info:input_type -> pbuser.reqInfo
	1, // 3: pbuser.ServiceJoin.Join:output_type -> pbuser.resJoin
	3, // 4: pbuser.ServiceLogin.Login:output_type -> pbuser.resLogin
	5, // 5: pbuser.ServiceInfo.Info:output_type -> pbuser.resInfo
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_pbuser_user_proto_init() }
func file_protos_pbuser_user_proto_init() {
	if File_protos_pbuser_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_pbuser_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqJoin); i {
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
		file_protos_pbuser_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResJoin); i {
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
		file_protos_pbuser_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqLogin); i {
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
		file_protos_pbuser_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResLogin); i {
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
		file_protos_pbuser_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqInfo); i {
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
		file_protos_pbuser_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResInfo); i {
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
			RawDescriptor: file_protos_pbuser_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_protos_pbuser_user_proto_goTypes,
		DependencyIndexes: file_protos_pbuser_user_proto_depIdxs,
		MessageInfos:      file_protos_pbuser_user_proto_msgTypes,
	}.Build()
	File_protos_pbuser_user_proto = out.File
	file_protos_pbuser_user_proto_rawDesc = nil
	file_protos_pbuser_user_proto_goTypes = nil
	file_protos_pbuser_user_proto_depIdxs = nil
}
