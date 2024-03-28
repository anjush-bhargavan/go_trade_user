package service

import (
	"context"
	"errors"

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
			Status:  pb.Response_ERROR,
			Message: "Failed to create seller",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Seller created successfully",
	}, nil
}

func (u *UserService) CreateProductService(p *pb.UserProduct) (*pb.Response, error) {
	ctx := context.Background()

	seller, err := u.Repo.FindSellerByUserID(uint(p.Seller_ID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "user is not registered as a seller",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	product := &productpb.Product{
		Name:      p.Name,
		Seller_ID: uint32(seller.ID),
		Category: &productpb.Category{
			Category_ID: p.Category.Category_ID,
		},
		Base_Price:         p.Base_Price,
		Description:        p.Description,
		Bidding_Type:       p.Bidding_Type,
		Bidding_Start_Time: p.Bidding_Start_Time,
		Bidding_End_Time:   p.Bidding_End_Time,
	}

	result, err := u.ProductClient.CreateProduct(ctx, product)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to create product",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_Status(result.Status),
		Message: result.Message,
	}, nil
}

func (u *UserService) EditProductService(p *pb.UserProduct) (*pb.UserProduct, error) {
	ctx := context.Background()

	seller, err := u.Repo.FindSellerByUserID(uint(p.Seller_ID))
	if err != nil {
		return nil, err
	}

	product, err := u.ProductClient.FindProductByID(ctx, &productpb.PrID{
		ID: p.Product_ID,
	})
	if err != nil {
		return nil, err
	}

	if seller.ID != uint(product.Seller_ID) {
		return nil, errors.New("invalid seller")
	}

	product.Base_Price = p.Base_Price
	product.Category.Category_ID = p.Category.Category_ID
	product.Name = p.Name
	product.Description = p.Description
	product.Bidding_Type = p.Bidding_Type
	product.Bidding_Start_Time = p.Bidding_Start_Time
	product.Bidding_End_Time = p.Bidding_End_Time

	_, err = u.ProductClient.EditProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u *UserService) RemoveProductService(p *pb.IDs) (*pb.Response, error) {
	ctx := context.Background()

	seller, err := u.Repo.FindSellerByUserID(uint(p.User_ID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to find seller",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	result, err := u.ProductClient.RemoveProduct(ctx, &productpb.PrIDs{
		ID:        p.ID,
		Seller_ID: uint32(seller.ID),
	})
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_Status(result.Status),
			Message: result.Message,
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: result.Message,
	}, nil
}
