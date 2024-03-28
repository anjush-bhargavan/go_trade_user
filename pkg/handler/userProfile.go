package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

// ViewProfile fetch the user profile from database 
func (u *UserHandler)ViewProfile(ctx context.Context,p *pb.ID) (*pb.Profile, error) {
	response, err := u.SVC.ViewProfileService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// EditProfile update the user profile in database 
func (u *UserHandler) EditProfile(ctx context.Context,p *pb.Profile) (*pb.Profile, error) {
	response, err := u.SVC.EditProfileService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}


// ChangePassword update the user profile in database 
func (u *UserHandler) ChangePassword(ctx context.Context,p *pb.Password) (*pb.Response, error) {
	response, err := u.SVC.ChangePassword(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// BlockUser update the user as blocked in database 
func (u *UserHandler) BlockUser(ctx context.Context, p *pb.ID) (*pb.Response, error) {
	response, err := u.SVC.BlockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}