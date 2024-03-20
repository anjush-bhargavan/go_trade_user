// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: user.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_UserSignup_FullMethodName     = "/pb.UserService/UserSignup"
	UserService_VerfiyUser_FullMethodName     = "/pb.UserService/VerfiyUser"
	UserService_UserLogin_FullMethodName      = "/pb.UserService/UserLogin"
	UserService_ViewProfile_FullMethodName    = "/pb.UserService/ViewProfile"
	UserService_EditProfile_FullMethodName    = "/pb.UserService/EditProfile"
	UserService_ChangePassword_FullMethodName = "/pb.UserService/ChangePassword"
	UserService_AddAddress_FullMethodName     = "/pb.UserService/AddAddress"
	UserService_ViewAllAddress_FullMethodName = "/pb.UserService/ViewAllAddress"
	UserService_EditAddress_FullMethodName    = "/pb.UserService/EditAddress"
	UserService_RemoveAddress_FullMethodName  = "/pb.UserService/RemoveAddress"
	UserService_BeASeller_FullMethodName      = "/pb.UserService/BeASeller"
	UserService_AddProduct_FullMethodName     = "/pb.UserService/AddProduct"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserSignup(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*Response, error)
	VerfiyUser(ctx context.Context, in *OTP, opts ...grpc.CallOption) (*Response, error)
	UserLogin(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Response, error)
	ViewProfile(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Profile, error)
	EditProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Profile, error)
	ChangePassword(ctx context.Context, in *Password, opts ...grpc.CallOption) (*Response, error)
	AddAddress(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Response, error)
	ViewAllAddress(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*AddressList, error)
	EditAddress(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error)
	RemoveAddress(ctx context.Context, in *IDs, opts ...grpc.CallOption) (*Response, error)
	BeASeller(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error)
	AddProduct(ctx context.Context, in *SellerProdcut, opts ...grpc.CallOption) (*Response, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserSignup(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_UserSignup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerfiyUser(ctx context.Context, in *OTP, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_VerfiyUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ViewProfile(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, UserService_ViewProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, UserService_EditProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *Password, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_ChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddAddress(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_AddAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ViewAllAddress(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*AddressList, error) {
	out := new(AddressList)
	err := c.cc.Invoke(ctx, UserService_ViewAllAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditAddress(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, UserService_EditAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveAddress(ctx context.Context, in *IDs, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_RemoveAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BeASeller(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_BeASeller_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddProduct(ctx context.Context, in *SellerProdcut, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_AddProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserSignup(context.Context, *Signup) (*Response, error)
	VerfiyUser(context.Context, *OTP) (*Response, error)
	UserLogin(context.Context, *Login) (*Response, error)
	ViewProfile(context.Context, *ID) (*Profile, error)
	EditProfile(context.Context, *Profile) (*Profile, error)
	ChangePassword(context.Context, *Password) (*Response, error)
	AddAddress(context.Context, *Address) (*Response, error)
	ViewAllAddress(context.Context, *NoParam) (*AddressList, error)
	EditAddress(context.Context, *Address) (*Address, error)
	RemoveAddress(context.Context, *IDs) (*Response, error)
	BeASeller(context.Context, *ID) (*Response, error)
	AddProduct(context.Context, *SellerProdcut) (*Response, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserSignup(context.Context, *Signup) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignup not implemented")
}
func (UnimplementedUserServiceServer) VerfiyUser(context.Context, *OTP) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerfiyUser not implemented")
}
func (UnimplementedUserServiceServer) UserLogin(context.Context, *Login) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServiceServer) ViewProfile(context.Context, *ID) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewProfile not implemented")
}
func (UnimplementedUserServiceServer) EditProfile(context.Context, *Profile) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProfile not implemented")
}
func (UnimplementedUserServiceServer) ChangePassword(context.Context, *Password) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserServiceServer) AddAddress(context.Context, *Address) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAddress not implemented")
}
func (UnimplementedUserServiceServer) ViewAllAddress(context.Context, *NoParam) (*AddressList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAllAddress not implemented")
}
func (UnimplementedUserServiceServer) EditAddress(context.Context, *Address) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAddress not implemented")
}
func (UnimplementedUserServiceServer) RemoveAddress(context.Context, *IDs) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAddress not implemented")
}
func (UnimplementedUserServiceServer) BeASeller(context.Context, *ID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeASeller not implemented")
}
func (UnimplementedUserServiceServer) AddProduct(context.Context, *SellerProdcut) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserSignup(ctx, req.(*Signup))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerfiyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerfiyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_VerfiyUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerfiyUser(ctx, req.(*OTP))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Login)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*Login))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ViewProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ViewProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ViewProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ViewProfile(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_EditProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditProfile(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Password)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*Password))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddAddress(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ViewAllAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ViewAllAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ViewAllAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ViewAllAddress(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_EditAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditAddress(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RemoveAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveAddress(ctx, req.(*IDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BeASeller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BeASeller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_BeASeller_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BeASeller(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerProdcut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddProduct(ctx, req.(*SellerProdcut))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSignup",
			Handler:    _UserService_UserSignup_Handler,
		},
		{
			MethodName: "VerfiyUser",
			Handler:    _UserService_VerfiyUser_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "ViewProfile",
			Handler:    _UserService_ViewProfile_Handler,
		},
		{
			MethodName: "EditProfile",
			Handler:    _UserService_EditProfile_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
		{
			MethodName: "AddAddress",
			Handler:    _UserService_AddAddress_Handler,
		},
		{
			MethodName: "ViewAllAddress",
			Handler:    _UserService_ViewAllAddress_Handler,
		},
		{
			MethodName: "EditAddress",
			Handler:    _UserService_EditAddress_Handler,
		},
		{
			MethodName: "RemoveAddress",
			Handler:    _UserService_RemoveAddress_Handler,
		},
		{
			MethodName: "BeASeller",
			Handler:    _UserService_BeASeller_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _UserService_AddProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}