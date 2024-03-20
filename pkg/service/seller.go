package service

import (
	"context"
	"fmt"

	productpb "github.com/anjush-bhargavan/go_trade_user/pkg/clients/product/pb"
	"github.com/anjush-bhargavan/go_trade_user/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

func (u *UserService) CreateSellerService(p *pb.ID) (*pb.Response, error) {
	seller := model.Seller{
		UserID: uint(p.ID),
	}

	err := u.Repo.CreateSeller(&seller)
	if err != nil {
		return &pb.Response{
			Status:  "Failed",
			Message: "Failed to create seller",
			Error:   err.Error(),
		}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: "Seller created successfully",
	}, nil
}

func (u *UserService) CreateProductService(p *pb.SellerProdcut) (*pb.Response, error) {
	ctx := context.Background()

	seller, err := u.Repo.FindSellerByUserID(uint(p.Seller_ID))
	if err != nil {
		return &pb.Response{
			Status:  "Failed",
			Message: "user is not registered as a seller",
			Error:   err.Error(),
		}, err
	}

	product := &productpb.Prodcut{
		Name:               p.Name,
		Seller_ID:          uint32(seller.ID),
		Category:           p.Category,
		Base_Price:         p.Base_Price,
		Description:        p.Description,
		Bidding_Type:       p.Bidding_Type,
		Bidding_Start_Time: p.Bidding_Start_Time,
		Bidiing_End_Time:   p.Bidiing_End_Time,
	}

	result, err := u.ProductClient.CreateProduct(ctx, product)
	if err != nil {
		return &pb.Response{
			Status:  "Failed",
			Message: "Failed to create product",
			Error:   err.Error(),
		}, err
	}
	fmt.Println("hiiii")
	return &pb.Response{
		Status:  result.Status,
		Message: result.Message,
		Error:   result.Error,
	}, nil
}
