package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

//BeASeller helps to create a user to be a seller.
func (u *UserHandler) BeASeller(ctx context.Context,p *pb.ID) (*pb.Response, error) {
	response, err := u.SVC.CreateSellerService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

//AddProduct helps to create a product for a seller.
func (u *UserHandler) AddProduct(ctx context.Context,p *pb.SellerProdcut) (*pb.Response, error) {
	response, err := u.SVC.CreateProductService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}