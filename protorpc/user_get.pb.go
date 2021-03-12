// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: user_get.proto

package protorpc

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

// 枚举类型第一个字段必须为 0
type UserSex int32

const (
	UserSex_MEN   UserSex = 0
	UserSex_WOMEN UserSex = 1
)

// Enum value maps for UserSex.
var (
	UserSex_name = map[int32]string{
		0: "MEN",
		1: "WOMEN",
	}
	UserSex_value = map[string]int32{
		"MEN":   0,
		"WOMEN": 1,
	}
)

func (x UserSex) Enum() *UserSex {
	p := new(UserSex)
	*p = x
	return p
}

func (x UserSex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserSex) Descriptor() protoreflect.EnumDescriptor {
	return file_user_get_proto_enumTypes[0].Descriptor()
}

func (UserSex) Type() protoreflect.EnumType {
	return &file_user_get_proto_enumTypes[0]
}

func (x UserSex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserSex.Descriptor instead.
func (UserSex) EnumDescriptor() ([]byte, []int) {
	return file_user_get_proto_rawDescGZIP(), []int{0}
}

// GetUser 请求结构
type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int64 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_get_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

// GetUser 响应结构
type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid   int64   `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Username string  `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Sex      UserSex `protobuf:"varint,3,opt,name=sex,proto3,enum=protorpc.UserSex" json:"sex,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_user_get_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *GetUserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetUserResponse) GetSex() UserSex {
	if x != nil {
		return x.Sex
	}
	return UserSex_MEN
}

// GetUserList 请求结构
type GetUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserListRequest) Reset() {
	*x = GetUserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListRequest) ProtoMessage() {}

func (x *GetUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListRequest.ProtoReflect.Descriptor instead.
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return file_user_get_proto_rawDescGZIP(), []int{2}
}

// 响应结构
type UserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repeated 重复(数组)
	List []*GetUserResponse `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_user_get_proto_rawDescGZIP(), []int{3}
}

func (x *UserListResponse) GetList() []*GetUserResponse {
	if x != nil {
		return x.List
	}
	return nil
}

var File_user_get_proto protoreflect.FileDescriptor

var file_user_get_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x70, 0x63, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x73,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78,
	0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x2a, 0x1d, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x78, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x57, 0x4f, 0x4d, 0x45, 0x4e, 0x10, 0x01, 0x32, 0x93, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_get_proto_rawDescOnce sync.Once
	file_user_get_proto_rawDescData = file_user_get_proto_rawDesc
)

func file_user_get_proto_rawDescGZIP() []byte {
	file_user_get_proto_rawDescOnce.Do(func() {
		file_user_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_get_proto_rawDescData)
	})
	return file_user_get_proto_rawDescData
}

var file_user_get_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_get_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_get_proto_goTypes = []interface{}{
	(UserSex)(0),               // 0: protorpc.UserSex
	(*GetUserRequest)(nil),     // 1: protorpc.GetUserRequest
	(*GetUserResponse)(nil),    // 2: protorpc.GetUserResponse
	(*GetUserListRequest)(nil), // 3: protorpc.GetUserListRequest
	(*UserListResponse)(nil),   // 4: protorpc.UserListResponse
}
var file_user_get_proto_depIdxs = []int32{
	0, // 0: protorpc.GetUserResponse.sex:type_name -> protorpc.UserSex
	2, // 1: protorpc.UserListResponse.list:type_name -> protorpc.GetUserResponse
	1, // 2: protorpc.User.GetUser:input_type -> protorpc.GetUserRequest
	3, // 3: protorpc.User.GetUserList:input_type -> protorpc.GetUserListRequest
	2, // 4: protorpc.User.GetUser:output_type -> protorpc.GetUserResponse
	4, // 5: protorpc.User.GetUserList:output_type -> protorpc.UserListResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_get_proto_init() }
func file_user_get_proto_init() {
	if File_user_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_user_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_user_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListRequest); i {
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
		file_user_get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResponse); i {
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
			RawDescriptor: file_user_get_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_get_proto_goTypes,
		DependencyIndexes: file_user_get_proto_depIdxs,
		EnumInfos:         file_user_get_proto_enumTypes,
		MessageInfos:      file_user_get_proto_msgTypes,
	}.Build()
	File_user_get_proto = out.File
	file_user_get_proto_rawDesc = nil
	file_user_get_proto_goTypes = nil
	file_user_get_proto_depIdxs = nil
}
