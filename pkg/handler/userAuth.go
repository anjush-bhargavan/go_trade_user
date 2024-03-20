package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

//Handler methods acts as intermediatery between services and incoming request

// UserSignup deliver the data to service layer to signup user
func (u *UserHandler) UserSignup(ctx context.Context, p *pb.Signup) (*pb.Response, error) {
	response, err := u.SVC.SignupService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// VerfiyUser handles the verfication details to service layer to complete signup.
func (u *UserHandler) VerfiyUser(ctx context.Context, p *pb.OTP) (*pb.Response, error) {
	response, err := u.SVC.VerificationService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// UserLogin handles return token after logging in
func (u *UserHandler) UserLogin(ctx context.Context, p *pb.Login) (*pb.Response, error) {
	response, err := u.SVC.LoginService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
