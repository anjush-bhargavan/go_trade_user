package interfaces

import (
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

// UserServiceInter represents an interface for user-related operations.
type UserServiceInter interface {
	SignupService(p *pb.Signup) (*pb.Response, error)
	LoginService(p *pb.Login) (*pb.Response, error)
	VerificationService(p *pb.OTP) (*pb.Response, error)

	ViewProfileService(p *pb.ID) (*pb.Profile, error)
	EditProfileService(p *pb.Profile) (*pb.Profile, error)
	ChangePassword(p *pb.Password) (*pb.Response,error)

	AddAddressService(p *pb.Address) (*pb.Response, error)
	EditAddressService(p *pb.Address) (*pb.Address, error)
	RemoveAddressService(p *pb.IDs) (*pb.Response, error)

	CreateSellerService(p *pb.ID) (*pb.Response, error)
	CreateProductService(p *pb.SellerProdcut) (*pb.Response,error)
}
