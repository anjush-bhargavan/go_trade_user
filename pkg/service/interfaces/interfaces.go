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
	BlockUserService(p *pb.ID) (*pb.Response, error)

	AddAddressService(p *pb.Address) (*pb.Response, error)
	EditAddressService(p *pb.Address) (*pb.Address, error)
	RemoveAddressService(p *pb.IDs) (*pb.Response, error)

	CreateSellerService(p *pb.ID) (*pb.Response, error)
	CreateProductService(p *pb.UserProduct) (*pb.Response,error)
	EditProductService(p *pb.UserProduct) (*pb.UserProduct, error) 
	RemoveProductService(p *pb.IDs) (*pb.Response, error)

	FindProductService(p *pb.ID) (*pb.UserProduct, error)
	FindAllProductService(p *pb.NoParam) (*pb.UserProductList,error)
	FindAllProductByCategoryService(p *pb.ID) (*pb.UserProductList,error)

	FindCategoryService(p *pb.ID) (*pb.UserCategory, error)
	FindAllCategoryService(p *pb.NoParam) (*pb.UserCategoryList, error)

	AddWatchlistService(p *pb.IDs) (*pb.Response, error)
	FetchWatchlistService(p *pb.ID) (*pb.UserCategoryList, error)

	AddBidService(p *pb.UserBid) (*pb.Response, error)
	GetBidsService(p *pb.ID) (*pb.UserBidList, error)
}
