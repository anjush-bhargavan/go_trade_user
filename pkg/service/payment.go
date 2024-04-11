package service

import (
	"context"
	"fmt"

	productpb "github.com/anjush-bhargavan/go_trade_user/pkg/clients/product/pb"
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

func (u *UserService) UserPaymentService(p *pb.UserBid) (*pb.UserPaymentResponse, error) {
	ctx := context.Background()

	bid := &productpb.Bid{
		User_ID:    p.User_ID,
		Product_ID: p.Product_ID,
		Amount:     p.Amount,
	}

	result, err := u.ProductClient.CreatePayment(ctx, bid)
	if err != nil {
		return nil, err
	}

	return &pb.UserPaymentResponse{
		Product_ID: result.Product_ID,
		User_Name:  result.User_Name,
		User_Email: result.User_Email,
		Amount:     result.Amount,
		Order_ID:   result.Order_ID,
	}, nil
}

func (u *UserService) UserPaymentSuccessService(p *pb.UserPayment) (*pb.Response, error) {
	ctx := context.Background()
	fmt.Println(p.Product_ID)
	payment := &productpb.Payment{
		User_ID:    p.User_ID,
		Amount:     p.Amount,
		Payment_ID: p.Payment_ID,
		Product_ID: p.Product_ID,
	}

	result, err := u.ProductClient.PaymentSuccess(ctx, payment)
	if err != nil {
		return nil, err
	}

	// data := result.Payload
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: result.Message,
		// Payload: &pb.Response_Data{Data: },
	}, nil
}
