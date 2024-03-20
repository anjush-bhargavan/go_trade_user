package repo

import "github.com/anjush-bhargavan/go_trade_user/pkg/model"

//CreateUser method for creating a new user in db
func (u *UserRepository) CreateUser(user *model.User) (uint ,error) {
	if err := u.DB.Create(&user).Error ; err != nil {
		return 0,err
	}
	return user.ID,nil
}

//FindUserByEmail method finds the user from database using email
func (u *UserRepository) FindUserByEmail(email string) (*model.User,error) {
	var user model.User

	if err := u.DB.Model(&model.User{}).Where("email = ?",email).First(&user).Error ; err != nil {
		return nil,err
	} 
	return &user,nil
}

//FindUserByID method finds the user from database using primary key
func (u *UserRepository) FindUserByID(userID uint) (*model.User,error) {
	var user model.User

	if err := u.DB.First(&user,userID).Error ; err != nil {
		return nil,err
	}
	return &user,nil
}

//UpdateUser method update the details of user in databse
func (u *UserRepository) UpdateUser(user *model.User) error {
	if err := u.DB.Save(&user).Error ; err != nil {
		return err
	}
	return nil
}

// FindReferralCode finds the referralcode user details
func (u *UserRepository) FindReferralCode(code string) (*model.User,error) {
	var user model.User

	if err := u.DB.Model(&model.User{}).Where("referral_code = ?",code).First(&user).Error ; err != nil {
		return nil,err
	} 
	return &user,nil
}