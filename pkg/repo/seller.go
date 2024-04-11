package repo

// import "github.com/anjush-bhargavan/go_trade_user/pkg/model"

// // CreateSeller will create a seller in the database.
// func (u *UserRepository) CreateSeller(seller *model.Seller) error {
// 	if err := u.DB.Create(&seller).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// //FindSeller method finds the user from database using primary key
// func (u *UserRepository) FindSeller(sellerID uint) (*model.Seller,error) {
// 	var seller model.Seller

// 	if err := u.DB.First(&seller,sellerID).Error ; err != nil {
// 		return nil,err
// 	}
// 	return &seller,nil
// }


// //FindSellerByUserID method finds the seller from database using user id.
// func (u *UserRepository) FindSellerByUserID(userID uint) (*model.Seller,error) {
// 	var seller model.Seller
// 	if err := u.DB.Where("user_id = ?",userID).First(&seller).Error ; err != nil {
// 		return nil,err
// 	}
// 	return &seller,nil
// }