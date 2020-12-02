// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: profile.proto

package codegen

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID uint64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Login  string `protobuf:"bytes,2,opt,name=Login,proto3" json:"Login,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0}
}

func (x *UserProfile) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserProfile) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCredentials *UserProfile `protobuf:"bytes,1,opt,name=UserCredentials,proto3" json:"UserCredentials,omitempty"`
	Name            string       `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Surname         string       `protobuf:"bytes,3,opt,name=Surname,proto3" json:"Surname,omitempty"`
	AvatarPath      string       `protobuf:"bytes,4,opt,name=AvatarPath,proto3" json:"AvatarPath,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetUserCredentials() *UserProfile {
	if x != nil {
		return x.UserCredentials
	}
	return nil
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Profile) GetAvatarPath() string {
	if x != nil {
		return x.AvatarPath
	}
	return ""
}

type Nil struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nil) Reset() {
	*x = Nil{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nil) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nil) ProtoMessage() {}

func (x *Nil) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nil.ProtoReflect.Descriptor instead.
func (*Nil) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{2}
}

type CreateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *CreateProfileRequest) Reset() {
	*x = CreateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileRequest) ProtoMessage() {}

func (x *CreateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProfileRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type GetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{4}
}

func (x *GetProfileRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type GetProfileByUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetProfileByUserIDRequest) Reset() {
	*x = GetProfileByUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileByUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileByUserIDRequest) ProtoMessage() {}

func (x *GetProfileByUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileByUserIDRequest.ProtoReflect.Descriptor instead.
func (*GetProfileByUserIDRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{5}
}

func (x *GetProfileByUserIDRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type UpdateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *UpdateProfileRequest) Reset() {
	*x = UpdateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileRequest) ProtoMessage() {}

func (x *UpdateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateProfileRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

var File_profile_proto protoreflect.FileDescriptor

var file_profile_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x3b, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x97, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x05, 0x0a, 0x03, 0x4e, 0x69, 0x6c, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x33, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x42, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x90,
	0x02, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4e, 0x69, 0x6c, 0x12,
	0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4e, 0x69,
	0x6c, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_proto_rawDescOnce sync.Once
	file_profile_proto_rawDescData = file_profile_proto_rawDesc
)

func file_profile_proto_rawDescGZIP() []byte {
	file_profile_proto_rawDescOnce.Do(func() {
		file_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_proto_rawDescData)
	})
	return file_profile_proto_rawDescData
}

var file_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_profile_proto_goTypes = []interface{}{
	(*UserProfile)(nil),               // 0: profile.UserProfile
	(*Profile)(nil),                   // 1: profile.Profile
	(*Nil)(nil),                       // 2: profile.Nil
	(*CreateProfileRequest)(nil),      // 3: profile.CreateProfileRequest
	(*GetProfileRequest)(nil),         // 4: profile.GetProfileRequest
	(*GetProfileByUserIDRequest)(nil), // 5: profile.GetProfileByUserIDRequest
	(*UpdateProfileRequest)(nil),      // 6: profile.UpdateProfileRequest
}
var file_profile_proto_depIdxs = []int32{
	0, // 0: profile.Profile.UserCredentials:type_name -> profile.UserProfile
	1, // 1: profile.CreateProfileRequest.profile:type_name -> profile.Profile
	1, // 2: profile.UpdateProfileRequest.profile:type_name -> profile.Profile
	3, // 3: profile.profileService.CreateProfile:input_type -> profile.CreateProfileRequest
	4, // 4: profile.profileService.GetProfile:input_type -> profile.GetProfileRequest
	5, // 5: profile.profileService.GetProfileByID:input_type -> profile.GetProfileByUserIDRequest
	6, // 6: profile.profileService.UpdateProfile:input_type -> profile.UpdateProfileRequest
	2, // 7: profile.profileService.CreateProfile:output_type -> profile.Nil
	1, // 8: profile.profileService.GetProfile:output_type -> profile.Profile
	1, // 9: profile.profileService.GetProfileByID:output_type -> profile.Profile
	2, // 10: profile.profileService.UpdateProfile:output_type -> profile.Nil
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_profile_proto_init() }
func file_profile_proto_init() {
	if File_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
		file_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nil); i {
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
		file_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileRequest); i {
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
		file_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileRequest); i {
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
		file_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileByUserIDRequest); i {
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
		file_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileRequest); i {
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
			RawDescriptor: file_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_proto_goTypes,
		DependencyIndexes: file_profile_proto_depIdxs,
		MessageInfos:      file_profile_proto_msgTypes,
	}.Build()
	File_profile_proto = out.File
	file_profile_proto_rawDesc = nil
	file_profile_proto_goTypes = nil
	file_profile_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileServiceClient interface {
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*Nil, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*Profile, error)
	GetProfileByID(ctx context.Context, in *GetProfileByUserIDRequest, opts ...grpc.CallOption) (*Profile, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*Nil, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := c.cc.Invoke(ctx, "/profile.profileService/CreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/profile.profileService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfileByID(ctx context.Context, in *GetProfileByUserIDRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/profile.profileService/GetProfileByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := c.cc.Invoke(ctx, "/profile.profileService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
type ProfileServiceServer interface {
	CreateProfile(context.Context, *CreateProfileRequest) (*Nil, error)
	GetProfile(context.Context, *GetProfileRequest) (*Profile, error)
	GetProfileByID(context.Context, *GetProfileByUserIDRequest) (*Profile, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*Nil, error)
}

// UnimplementedProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (*UnimplementedProfileServiceServer) CreateProfile(context.Context, *CreateProfileRequest) (*Nil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (*UnimplementedProfileServiceServer) GetProfile(context.Context, *GetProfileRequest) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (*UnimplementedProfileServiceServer) GetProfileByID(context.Context, *GetProfileByUserIDRequest) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileByID not implemented")
}
func (*UnimplementedProfileServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*Nil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.profileService/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.profileService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfileByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfileByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.profileService/GetProfileByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfileByID(ctx, req.(*GetProfileByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.profileService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.profileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProfile",
			Handler:    _ProfileService_CreateProfile_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _ProfileService_GetProfile_Handler,
		},
		{
			MethodName: "GetProfileByID",
			Handler:    _ProfileService_GetProfileByID_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _ProfileService_UpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
