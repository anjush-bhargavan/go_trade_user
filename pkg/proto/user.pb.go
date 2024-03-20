// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: user.proto

package __

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

type NoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoParam) Reset() {
	*x = NoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParam) ProtoMessage() {}

func (x *NoParam) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParam.ProtoReflect.Descriptor instead.
func (*NoParam) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Signup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User_Name     string `protobuf:"bytes,1,opt,name=User_Name,json=UserName,proto3" json:"User_Name,omitempty"`
	Email         string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password      string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Mobile        string `protobuf:"bytes,4,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Referral_Code string `protobuf:"bytes,5,opt,name=Referral_Code,json=ReferralCode,proto3" json:"Referral_Code,omitempty"`
}

func (x *Signup) Reset() {
	*x = Signup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signup) ProtoMessage() {}

func (x *Signup) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signup.ProtoReflect.Descriptor instead.
func (*Signup) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *Signup) GetUser_Name() string {
	if x != nil {
		return x.User_Name
	}
	return ""
}

func (x *Signup) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Signup) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Signup) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Signup) GetReferral_Code() string {
	if x != nil {
		return x.Referral_Code
	}
	return ""
}

type OTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	OTP   string `protobuf:"bytes,2,opt,name=OTP,proto3" json:"OTP,omitempty"`
}

func (x *OTP) Reset() {
	*x = OTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTP) ProtoMessage() {}

func (x *OTP) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTP.ProtoReflect.Descriptor instead.
func (*OTP) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *OTP) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OTP) GetOTP() string {
	if x != nil {
		return x.OTP
	}
	return ""
}

type Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *Login) Reset() {
	*x = Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login) ProtoMessage() {}

func (x *Login) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login.ProtoReflect.Descriptor instead.
func (*Login) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *Login) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Login) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *ID) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type IDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	User_ID uint32 `protobuf:"varint,2,opt,name=User_ID,json=UserID,proto3" json:"User_ID,omitempty"`
}

func (x *IDs) Reset() {
	*x = IDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDs) ProtoMessage() {}

func (x *IDs) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDs.ProtoReflect.Descriptor instead.
func (*IDs) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *IDs) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *IDs) GetUser_ID() uint32 {
	if x != nil {
		return x.User_ID
	}
	return 0
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User_ID       uint32  `protobuf:"varint,1,opt,name=User_ID,json=UserID,proto3" json:"User_ID,omitempty"`
	User_Name     string  `protobuf:"bytes,2,opt,name=User_Name,json=UserName,proto3" json:"User_Name,omitempty"`
	Email         string  `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Mobile        string  `protobuf:"bytes,4,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Referral_Code string  `protobuf:"bytes,5,opt,name=Referral_Code,json=ReferralCode,proto3" json:"Referral_Code,omitempty"`
	Wallet        float32 `protobuf:"fixed32,6,opt,name=Wallet,proto3" json:"Wallet,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
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
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *Profile) GetUser_ID() uint32 {
	if x != nil {
		return x.User_ID
	}
	return 0
}

func (x *Profile) GetUser_Name() string {
	if x != nil {
		return x.User_Name
	}
	return ""
}

func (x *Profile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Profile) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Profile) GetReferral_Code() string {
	if x != nil {
		return x.Referral_Code
	}
	return ""
}

func (x *Profile) GetWallet() float32 {
	if x != nil {
		return x.Wallet
	}
	return 0
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint32 `protobuf:"varint,7,opt,name=ID,proto3" json:"ID,omitempty"`
	User_ID uint32 `protobuf:"varint,1,opt,name=User_ID,json=UserID,proto3" json:"User_ID,omitempty"`
	House   string `protobuf:"bytes,2,opt,name=House,proto3" json:"House,omitempty"`
	Street  string `protobuf:"bytes,3,opt,name=Street,proto3" json:"Street,omitempty"`
	City    string `protobuf:"bytes,4,opt,name=City,proto3" json:"City,omitempty"`
	Zip     uint32 `protobuf:"varint,5,opt,name=Zip,proto3" json:"Zip,omitempty"`
	State   string `protobuf:"bytes,6,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *Address) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Address) GetUser_ID() uint32 {
	if x != nil {
		return x.User_ID
	}
	return 0
}

func (x *Address) GetHouse() string {
	if x != nil {
		return x.House
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetZip() uint32 {
	if x != nil {
		return x.Zip
	}
	return 0
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type AddressList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []*Address `protobuf:"bytes,1,rep,name=Addresses,proto3" json:"Addresses,omitempty"`
}

func (x *AddressList) Reset() {
	*x = AddressList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressList) ProtoMessage() {}

func (x *AddressList) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressList.ProtoReflect.Descriptor instead.
func (*AddressList) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *AddressList) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type Password struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User_ID          uint32 `protobuf:"varint,1,opt,name=User_ID,json=UserID,proto3" json:"User_ID,omitempty"`
	Old_Password     string `protobuf:"bytes,2,opt,name=Old_Password,json=OldPassword,proto3" json:"Old_Password,omitempty"`
	New_Password     string `protobuf:"bytes,3,opt,name=New_Password,json=NewPassword,proto3" json:"New_Password,omitempty"`
	Confirm_Password string `protobuf:"bytes,4,opt,name=Confirm_Password,json=ConfirmPassword,proto3" json:"Confirm_Password,omitempty"`
}

func (x *Password) Reset() {
	*x = Password{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Password) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Password) ProtoMessage() {}

func (x *Password) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Password.ProtoReflect.Descriptor instead.
func (*Password) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *Password) GetUser_ID() uint32 {
	if x != nil {
		return x.User_ID
	}
	return 0
}

func (x *Password) GetOld_Password() string {
	if x != nil {
		return x.Old_Password
	}
	return ""
}

func (x *Password) GetNew_Password() string {
	if x != nil {
		return x.New_Password
	}
	return ""
}

func (x *Password) GetConfirm_Password() string {
	if x != nil {
		return x.Confirm_Password
	}
	return ""
}

type SellerProdcut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Seller_ID          uint32 `protobuf:"varint,2,opt,name=Seller_ID,json=SellerID,proto3" json:"Seller_ID,omitempty"`
	Category           uint32 `protobuf:"varint,3,opt,name=Category,proto3" json:"Category,omitempty"`
	Base_Price         uint32 `protobuf:"varint,4,opt,name=Base_Price,json=BasePrice,proto3" json:"Base_Price,omitempty"`
	Description        string `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	Bidding_Type       bool   `protobuf:"varint,6,opt,name=Bidding_Type,json=BiddingType,proto3" json:"Bidding_Type,omitempty"`
	Bidding_Start_Time string `protobuf:"bytes,7,opt,name=Bidding_Start_Time,json=BiddingStartTime,proto3" json:"Bidding_Start_Time,omitempty"`
	Bidiing_End_Time   string `protobuf:"bytes,8,opt,name=Bidiing_End_Time,json=BidiingEndTime,proto3" json:"Bidiing_End_Time,omitempty"`
	Listed_On          string `protobuf:"bytes,9,opt,name=Listed_On,json=ListedOn,proto3" json:"Listed_On,omitempty"`
}

func (x *SellerProdcut) Reset() {
	*x = SellerProdcut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerProdcut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerProdcut) ProtoMessage() {}

func (x *SellerProdcut) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerProdcut.ProtoReflect.Descriptor instead.
func (*SellerProdcut) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *SellerProdcut) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SellerProdcut) GetSeller_ID() uint32 {
	if x != nil {
		return x.Seller_ID
	}
	return 0
}

func (x *SellerProdcut) GetCategory() uint32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *SellerProdcut) GetBase_Price() uint32 {
	if x != nil {
		return x.Base_Price
	}
	return 0
}

func (x *SellerProdcut) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SellerProdcut) GetBidding_Type() bool {
	if x != nil {
		return x.Bidding_Type
	}
	return false
}

func (x *SellerProdcut) GetBidding_Start_Time() string {
	if x != nil {
		return x.Bidding_Start_Time
	}
	return ""
}

func (x *SellerProdcut) GetBidiing_End_Time() string {
	if x != nil {
		return x.Bidiing_End_Time
	}
	return ""
}

func (x *SellerProdcut) GetListed_On() string {
	if x != nil {
		return x.Listed_On
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x09, 0x0a, 0x07, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x52, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x94, 0x01, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x5f, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2d, 0x0a, 0x03, 0x4f, 0x54, 0x50, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x4f, 0x54, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4f, 0x54, 0x50, 0x22, 0x39, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x22, 0x2e, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x17, 0x0a,
	0x07, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0xaa, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x61, 0x6c, 0x5f, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x17, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x48, 0x6f, 0x75, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x5a, 0x69,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x5a, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x38, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x94, 0x01, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x21, 0x0a, 0x0c, 0x4f, 0x6c, 0x64, 0x5f, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x6c, 0x64, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x5f, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x5f, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0xb5, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x63, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x53, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x53, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x5f, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x69, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x42, 0x69, 0x64, 0x69, 0x69, 0x6e, 0x67, 0x5f,
	0x45, 0x6e, 0x64, 0x5f, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x42, 0x69, 0x64, 0x69, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x4f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x32, 0xf7, 0x03, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0a, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x66, 0x69, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x54, 0x50, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x06, 0x2e,
	0x70, 0x62, 0x2e, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x0e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0e, 0x56, 0x69, 0x65, 0x77, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0d, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x07, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x44, 0x73, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x42, 0x65, 0x41, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x64, 0x63, 0x75, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_user_proto_goTypes = []interface{}{
	(*NoParam)(nil),       // 0: pb.NoParam
	(*Response)(nil),      // 1: pb.Response
	(*Signup)(nil),        // 2: pb.Signup
	(*OTP)(nil),           // 3: pb.OTP
	(*Login)(nil),         // 4: pb.Login
	(*ID)(nil),            // 5: pb.ID
	(*IDs)(nil),           // 6: pb.IDs
	(*Profile)(nil),       // 7: pb.Profile
	(*Address)(nil),       // 8: pb.Address
	(*AddressList)(nil),   // 9: pb.AddressList
	(*Password)(nil),      // 10: pb.Password
	(*SellerProdcut)(nil), // 11: pb.SellerProdcut
}
var file_user_proto_depIdxs = []int32{
	8,  // 0: pb.AddressList.Addresses:type_name -> pb.Address
	2,  // 1: pb.UserService.UserSignup:input_type -> pb.Signup
	3,  // 2: pb.UserService.VerfiyUser:input_type -> pb.OTP
	4,  // 3: pb.UserService.UserLogin:input_type -> pb.Login
	5,  // 4: pb.UserService.ViewProfile:input_type -> pb.ID
	7,  // 5: pb.UserService.EditProfile:input_type -> pb.Profile
	10, // 6: pb.UserService.ChangePassword:input_type -> pb.Password
	8,  // 7: pb.UserService.AddAddress:input_type -> pb.Address
	0,  // 8: pb.UserService.ViewAllAddress:input_type -> pb.NoParam
	8,  // 9: pb.UserService.EditAddress:input_type -> pb.Address
	6,  // 10: pb.UserService.RemoveAddress:input_type -> pb.IDs
	5,  // 11: pb.UserService.BeASeller:input_type -> pb.ID
	11, // 12: pb.UserService.AddProduct:input_type -> pb.SellerProdcut
	1,  // 13: pb.UserService.UserSignup:output_type -> pb.Response
	1,  // 14: pb.UserService.VerfiyUser:output_type -> pb.Response
	1,  // 15: pb.UserService.UserLogin:output_type -> pb.Response
	7,  // 16: pb.UserService.ViewProfile:output_type -> pb.Profile
	7,  // 17: pb.UserService.EditProfile:output_type -> pb.Profile
	1,  // 18: pb.UserService.ChangePassword:output_type -> pb.Response
	1,  // 19: pb.UserService.AddAddress:output_type -> pb.Response
	9,  // 20: pb.UserService.ViewAllAddress:output_type -> pb.AddressList
	8,  // 21: pb.UserService.EditAddress:output_type -> pb.Address
	1,  // 22: pb.UserService.RemoveAddress:output_type -> pb.Response
	1,  // 23: pb.UserService.BeASeller:output_type -> pb.Response
	1,  // 24: pb.UserService.AddProduct:output_type -> pb.Response
	13, // [13:25] is the sub-list for method output_type
	1,  // [1:13] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoParam); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signup); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTP); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Login); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDs); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressList); i {
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
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Password); i {
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
		file_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerProdcut); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}