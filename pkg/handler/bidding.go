package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

func (u *UserHandler) AddBid(ctx context.Context, p *pb.UserBid) (*pb.Response, error) {
	response, err := u.SVC.AddBidService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) GetBids(ctx context.Context, p *pb.ID) (*pb.UserBidList, error) {
	response, err := u.SVC.GetBidsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
