package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)


func (u *UserHandler) UserCreatePayment(ctx context.Context, p *pb.UserBid) (*pb.UserPaymentResponse, error) {
	response, err := u.SVC.UserPaymentService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) UserPaymentSuccess(ctx context.Context, p *pb.UserPayment) (*pb.Response, error) {
	response, err := u.SVC.UserPaymentSuccessService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
