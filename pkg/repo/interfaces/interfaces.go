package interfaces

import "github.com/anjush-bhargavan/go_trade_user/pkg/model"

// UserRepoInter represents an interface for interacting with user-related data in a repository.
type UserRepoInter interface {
	CreateUser(user *model.User) (uint,error)
	FindUserByEmail(email string) (*model.User, error)
	FindUserByID(userID uint) (*model.User, error)
	UpdateUser(user *model.User) error
	FindReferralCode(code string) (*model.User,error)

	CreateAddress(address *model.Address) error
	EditAddress(address *model.Address) error
	FindAddressByUser(userID uint) (*model.Address, error)
	RemoveAddress(addressID,userID uint) error

	// CreateSeller(seller *model.Seller) error
	// FindSeller(sellerID uint) (*model.Seller,error)
	// FindSellerByUserID(userID uint) (*model.Seller,error)
	
	CreateTransaction(Transaction *model.Transaction) error
	FindTransactionByUserID(userID uint) (*[]model.Transaction,error)
	FindAllTransactions() (*[]model.Transaction,error)

	CreateWatchlist(item *model.Wathlist) error
	FetchWatchlistByUser(userID uint) (*[]model.Wathlist,error)
	FetchWatchlistByCategory(categoryID uint) (*[]model.Wathlist,error)
}
