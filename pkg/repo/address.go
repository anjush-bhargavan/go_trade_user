package repo

import "github.com/anjush-bhargavan/go_trade_user/pkg/model"

//CreateAddress method for creating a new address for a user in db.
func (u *UserRepository) CreateAddress(address *model.Address) error {
	if err := u.DB.Create(&address).Error; err != nil {
		return err
	}
	return nil
}

//EditAddress method update the details of address in database.
func (u *UserRepository) EditAddress(address *model.Address) error {
	if err := u.DB.Save(&address).Error; err != nil {
		return err
	}
	return nil
}

//FindAddressByUser method will find the address of user by userID.
func (u *UserRepository) FindAddressByUser(userID uint) (*model.Address, error) {
	var address model.Address
	if err := u.DB.Where("userid = ?", userID).First(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

//RemoveAddress will remove the address from database.
func (u *UserRepository) RemoveAddress(addressID,userID uint) error {
	if err := u.DB.Where("id = ? AND user_id = ?",addressID,userID).Delete(&model.Address{}).Error; err != nil {
		return err
	}
	return nil
}
