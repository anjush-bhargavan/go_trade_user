package repo

import "github.com/anjush-bhargavan/go_trade_user/pkg/model"

// CreateTransaction will create a Transaction for seller in the database.
func (u *UserRepository) CreateTransaction(Transaction *model.Transaction) error {
	if err := u.DB.Create(&Transaction).Error ; err != nil {
		return err
	}
	return nil
}

//FindTransactionByUserID will fetch the transactions by userID.
func (u *UserRepository) FindTransactionByUserID(userID uint) (*[]model.Transaction,error) {
	var Transaction []model.Transaction
	if err := u.DB.Where("user_id = ?",userID).Find(&Transaction).Error; err != nil {
		return nil,err
	}
	return &Transaction,nil
}