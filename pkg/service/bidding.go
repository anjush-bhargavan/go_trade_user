package service

import (
	"context"

	productpb "github.com/anjush-bhargavan/go_trade_user/pkg/clients/product/pb"
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

func (u *UserService) AddBidService(p *pb.UserBid) (*pb.Response, error) {
	ctx := context.Background()

	bid := productpb.Bid{
		User_ID:    p.User_ID,
		Product_ID: p.Product_ID,
		Amount:     p.Amount,
	}

	_, err := u.ProductClient.CreateBid(ctx, &bid)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to create bid",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "bid added successfully",
	}, nil
}


func (u *UserService) GetBidsService(p *pb.ID) (*pb.UserBidList, error) {
	ctx := context.Background()


	result, err := u.ProductClient.FetchBids(ctx,&productpb.PrID{ID: p.ID})
	if err != nil {
		return nil, err
	}
	var bids []*pb.UserBid

	for _,bid := range result.Bids {
		bids =append(bids, &pb.UserBid{
			User_ID: bid.User_ID,
			Product_ID: bid.Product_ID,
			Amount: bid.Amount,
		})
	}

	return &pb.UserBidList{
		Bids: bids,
	}, nil
}