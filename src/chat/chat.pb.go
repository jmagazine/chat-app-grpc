// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.1
// source: chat/chat.proto

package chat_app_grpc

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName string `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName string `protobuf:"bytes,1,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserMessage) Reset() {
	*x = CreateUserMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserMessage) ProtoMessage() {}

func (x *CreateUserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserMessage.ProtoReflect.Descriptor instead.
func (*CreateUserMessage) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserMessage) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateUserMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DeleteUserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserMessage) Reset() {
	*x = DeleteUserMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserMessage) ProtoMessage() {}

func (x *DeleteUserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserMessage.ProtoReflect.Descriptor instead.
func (*DeleteUserMessage) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteUserMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateUserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatedFields map[string]string `protobuf:"bytes,1,rep,name=updatedFields,proto3" json:"updatedFields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpdateUserMessage) Reset() {
	*x = UpdateUserMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserMessage) ProtoMessage() {}

func (x *UpdateUserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserMessage.ProtoReflect.Descriptor instead.
func (*UpdateUserMessage) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserMessage) GetUpdatedFields() map[string]string {
	if x != nil {
		return x.UpdatedFields
	}
	return nil
}

var File_chat_chat_proto protoreflect.FileDescriptor

var file_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x6a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x67, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x23, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xa7, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xa9, 0x01, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x31, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0a, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_chat_proto_rawDescOnce sync.Once
	file_chat_chat_proto_rawDescData = file_chat_chat_proto_rawDesc
)

func file_chat_chat_proto_rawDescGZIP() []byte {
	file_chat_chat_proto_rawDescOnce.Do(func() {
		file_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_chat_proto_rawDescData)
	})
	return file_chat_chat_proto_rawDescData
}

var file_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chat_chat_proto_goTypes = []interface{}{
	(*User)(nil),              // 0: chat.User
	(*CreateUserMessage)(nil), // 1: chat.CreateUserMessage
	(*DeleteUserMessage)(nil), // 2: chat.DeleteUserMessage
	(*UpdateUserMessage)(nil), // 3: chat.UpdateUserMessage
	nil,                       // 4: chat.UpdateUserMessage.UpdatedFieldsEntry
}
var file_chat_chat_proto_depIdxs = []int32{
	4, // 0: chat.UpdateUserMessage.updatedFields:type_name -> chat.UpdateUserMessage.UpdatedFieldsEntry
	1, // 1: chat.ChatService.CreateNewUser:input_type -> chat.CreateUserMessage
	2, // 2: chat.ChatService.DeleteUser:input_type -> chat.DeleteUserMessage
	3, // 3: chat.ChatService.UpdateUser:input_type -> chat.UpdateUserMessage
	0, // 4: chat.ChatService.CreateNewUser:output_type -> chat.User
	0, // 5: chat.ChatService.DeleteUser:output_type -> chat.User
	0, // 6: chat.ChatService.UpdateUser:output_type -> chat.User
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_chat_proto_init() }
func file_chat_chat_proto_init() {
	if File_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserMessage); i {
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
		file_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserMessage); i {
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
		file_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserMessage); i {
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
			RawDescriptor: file_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_chat_proto_goTypes,
		DependencyIndexes: file_chat_chat_proto_depIdxs,
		MessageInfos:      file_chat_chat_proto_msgTypes,
	}.Build()
	File_chat_chat_proto = out.File
	file_chat_chat_proto_rawDesc = nil
	file_chat_chat_proto_goTypes = nil
	file_chat_chat_proto_depIdxs = nil
}
