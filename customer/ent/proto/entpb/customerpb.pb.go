// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: entpb/customerpb.proto

package entpb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber      string               `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email            string               `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	LicenseId        string               `protobuf:"bytes,5,opt,name=license_id,json=licenseId,proto3" json:"license_id,omitempty"`
	Address          string               `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	MembershipNumber int64                `protobuf:"varint,7,opt,name=membership_number,json=membershipNumber,proto3" json:"membership_number,omitempty"`
	IsActive         bool                 `protobuf:"varint,8,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Password         string               `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	CreateAt         *timestamp.Timestamp `protobuf:"bytes,10,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt         *timestamp.Timestamp `protobuf:"bytes,11,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	Role             string               `protobuf:"bytes,12,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_customerpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_customerpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_entpb_customerpb_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetLicenseId() string {
	if x != nil {
		return x.LicenseId
	}
	return ""
}

func (x *Customer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Customer) GetMembershipNumber() int64 {
	if x != nil {
		return x.MembershipNumber
	}
	return 0
}

func (x *Customer) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Customer) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Customer) GetCreateAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *Customer) GetUpdateAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

func (x *Customer) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type RegisterInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber      string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email            string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	LicenseId        string `protobuf:"bytes,5,opt,name=license_id,json=licenseId,proto3" json:"license_id,omitempty"`
	Address          string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	MembershipNumber int64  `protobuf:"varint,7,opt,name=membership_number,json=membershipNumber,proto3" json:"membership_number,omitempty"`
	Password         string `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterInput) Reset() {
	*x = RegisterInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_customerpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterInput) ProtoMessage() {}

func (x *RegisterInput) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_customerpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterInput.ProtoReflect.Descriptor instead.
func (*RegisterInput) Descriptor() ([]byte, []int) {
	return file_entpb_customerpb_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterInput) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RegisterInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterInput) GetLicenseId() string {
	if x != nil {
		return x.LicenseId
	}
	return ""
}

func (x *RegisterInput) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegisterInput) GetMembershipNumber() int64 {
	if x != nil {
		return x.MembershipNumber
	}
	return 0
}

func (x *RegisterInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetCustomerByEmailInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetCustomerByEmailInput) Reset() {
	*x = GetCustomerByEmailInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_customerpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByEmailInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByEmailInput) ProtoMessage() {}

func (x *GetCustomerByEmailInput) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_customerpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByEmailInput.ProtoReflect.Descriptor instead.
func (*GetCustomerByEmailInput) Descriptor() ([]byte, []int) {
	return file_entpb_customerpb_proto_rawDescGZIP(), []int{2}
}

func (x *GetCustomerByEmailInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type LoginInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginInput) Reset() {
	*x = LoginInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_customerpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginInput) ProtoMessage() {}

func (x *LoginInput) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_customerpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginInput.ProtoReflect.Descriptor instead.
func (*LoginInput) Descriptor() ([]byte, []int) {
	return file_entpb_customerpb_proto_rawDescGZIP(), []int{3}
}

func (x *LoginInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateCustomerInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber      string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email            string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	LicenseId        string `protobuf:"bytes,5,opt,name=license_id,json=licenseId,proto3" json:"license_id,omitempty"`
	Address          string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	MembershipNumber int64  `protobuf:"varint,7,opt,name=membership_number,json=membershipNumber,proto3" json:"membership_number,omitempty"`
}

func (x *UpdateCustomerInput) Reset() {
	*x = UpdateCustomerInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_customerpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCustomerInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerInput) ProtoMessage() {}

func (x *UpdateCustomerInput) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_customerpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerInput.ProtoReflect.Descriptor instead.
func (*UpdateCustomerInput) Descriptor() ([]byte, []int) {
	return file_entpb_customerpb_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCustomerInput) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCustomerInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCustomerInput) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UpdateCustomerInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateCustomerInput) GetLicenseId() string {
	if x != nil {
		return x.LicenseId
	}
	return ""
}

func (x *UpdateCustomerInput) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateCustomerInput) GetMembershipNumber() int64 {
	if x != nil {
		return x.MembershipNumber
	}
	return 0
}

type ChangePasswordInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OldPassword     string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword     string `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	ConfirmPassword string `protobuf:"bytes,4,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
}

func (x *ChangePasswordInput) Reset() {
	*x = ChangePasswordInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_customerpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordInput) ProtoMessage() {}

func (x *ChangePasswordInput) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_customerpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordInput.ProtoReflect.Descriptor instead.
func (*ChangePasswordInput) Descriptor() ([]byte, []int) {
	return file_entpb_customerpb_proto_rawDescGZIP(), []int{5}
}

func (x *ChangePasswordInput) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChangePasswordInput) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangePasswordInput) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *ChangePasswordInput) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

type UpdateRoleInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UpdateRoleInput) Reset() {
	*x = UpdateRoleInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_customerpb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleInput) ProtoMessage() {}

func (x *UpdateRoleInput) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_customerpb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleInput.ProtoReflect.Descriptor instead.
func (*UpdateRoleInput) Descriptor() ([]byte, []int) {
	return file_entpb_customerpb_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRoleInput) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRoleInput) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_entpb_customerpb_proto protoreflect.FileDescriptor

var file_entpb_customerpb_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8c, 0x03, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22,
	0xde, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x2f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42,
	0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x3e, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0xd8, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x2b, 0x0a, 0x11, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x96, 0x01, 0x0a,
	0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x32, 0xe5, 0x02, 0x0a,
	0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x31, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x0f, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0f, 0x2e, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0f, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0f, 0x2e,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x3d,
	0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1a, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0f, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x35, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x0f, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x74, 0x65, 0x6d, 0x74, 0x61, 0x6d, 0x2f, 0x65, 0x6e, 0x74, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entpb_customerpb_proto_rawDescOnce sync.Once
	file_entpb_customerpb_proto_rawDescData = file_entpb_customerpb_proto_rawDesc
)

func file_entpb_customerpb_proto_rawDescGZIP() []byte {
	file_entpb_customerpb_proto_rawDescOnce.Do(func() {
		file_entpb_customerpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_entpb_customerpb_proto_rawDescData)
	})
	return file_entpb_customerpb_proto_rawDescData
}

var file_entpb_customerpb_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_entpb_customerpb_proto_goTypes = []interface{}{
	(*Customer)(nil),                // 0: entpb.Customer
	(*RegisterInput)(nil),           // 1: entpb.RegisterInput
	(*GetCustomerByEmailInput)(nil), // 2: entpb.GetCustomerByEmailInput
	(*LoginInput)(nil),              // 3: entpb.LoginInput
	(*UpdateCustomerInput)(nil),     // 4: entpb.UpdateCustomerInput
	(*ChangePasswordInput)(nil),     // 5: entpb.ChangePasswordInput
	(*UpdateRoleInput)(nil),         // 6: entpb.UpdateRoleInput
	(*timestamp.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_entpb_customerpb_proto_depIdxs = []int32{
	7, // 0: entpb.Customer.create_at:type_name -> google.protobuf.Timestamp
	7, // 1: entpb.Customer.update_at:type_name -> google.protobuf.Timestamp
	1, // 2: entpb.CustomerService.Register:input_type -> entpb.RegisterInput
	2, // 3: entpb.CustomerService.GetCustomerByEmail:input_type -> entpb.GetCustomerByEmailInput
	3, // 4: entpb.CustomerService.Login:input_type -> entpb.LoginInput
	4, // 5: entpb.CustomerService.Update:input_type -> entpb.UpdateCustomerInput
	5, // 6: entpb.CustomerService.ChangePassword:input_type -> entpb.ChangePasswordInput
	6, // 7: entpb.CustomerService.UpdateRole:input_type -> entpb.UpdateRoleInput
	0, // 8: entpb.CustomerService.Register:output_type -> entpb.Customer
	0, // 9: entpb.CustomerService.GetCustomerByEmail:output_type -> entpb.Customer
	0, // 10: entpb.CustomerService.Login:output_type -> entpb.Customer
	0, // 11: entpb.CustomerService.Update:output_type -> entpb.Customer
	0, // 12: entpb.CustomerService.ChangePassword:output_type -> entpb.Customer
	0, // 13: entpb.CustomerService.UpdateRole:output_type -> entpb.Customer
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_entpb_customerpb_proto_init() }
func file_entpb_customerpb_proto_init() {
	if File_entpb_customerpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entpb_customerpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_entpb_customerpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterInput); i {
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
		file_entpb_customerpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerByEmailInput); i {
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
		file_entpb_customerpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginInput); i {
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
		file_entpb_customerpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCustomerInput); i {
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
		file_entpb_customerpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordInput); i {
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
		file_entpb_customerpb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleInput); i {
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
			RawDescriptor: file_entpb_customerpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_entpb_customerpb_proto_goTypes,
		DependencyIndexes: file_entpb_customerpb_proto_depIdxs,
		MessageInfos:      file_entpb_customerpb_proto_msgTypes,
	}.Build()
	File_entpb_customerpb_proto = out.File
	file_entpb_customerpb_proto_rawDesc = nil
	file_entpb_customerpb_proto_goTypes = nil
	file_entpb_customerpb_proto_depIdxs = nil
}
