package service

import (
	"errors"
	"fmt"

	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
	"github.com/anjush-bhargavan/go_trade_user/utils"
)

func (u *UserService) ViewProfileService(p *pb.ID) (*pb.Profile, error) {
	user, err := u.Repo.FindUserByID(uint(p.ID))
	if err != nil {
		return nil, err
	}
	userModel := &pb.Profile{
		User_Name:     user.UserName,
		Email:         user.Email,
		Mobile:        user.Mobile,
		Referral_Code: user.ReferralCode,
		Wallet:        float32(user.Wallet),
	}
	return userModel, nil
}

func (u *UserService) EditProfileService(p *pb.Profile) (*pb.Profile, error) {
	user, err := u.Repo.FindUserByID(uint(p.User_ID))
	if err != nil {
		return nil, err
	}

	user.UserName = p.User_Name
	user.Mobile = p.Mobile

	err = u.Repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (u *UserService) ChangePassword(p *pb.Password) (*pb.Response, error) {

	user, err := u.Repo.FindUserByID(uint(p.User_ID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting user from database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	fmt.Println(p.Old_Password, user.Password)
	if !utils.CheckPassword(p.Old_Password, user.Password) {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "old password is incorrect",
		}, errors.New("old password mismatch")
	}

	if p.New_Password != p.Confirm_Password {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "password is incorrect, passwords not matching",
		}, errors.New("new password mismatch")
	}

	newPassword, err := utils.HashPassword(p.New_Password)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error while hashing new password",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	user.Password = newPassword

	err = u.Repo.UpdateUser(user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error while updating password",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Password changed successfully",
	}, nil
}

func (u *UserService) BlockUserService(p *pb.ID) (*pb.Response, error) {
	user, err := u.Repo.FindUserByID(uint(p.ID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting user from database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	user.IsBlocked = true

	err = u.Repo.UpdateUser(user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in updating user in database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "user blocked successfully",
	}, nil
}
