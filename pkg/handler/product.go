package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

// FindProductByIDUser retrieves a product by its ID for the user by delegating the operation to the FindProductService method of the User service.
func (u *UserHandler) FindProductByIDUser(ctx context.Context, p *pb.ID) (*pb.UserProduct, error) {
	response, err := u.SVC.FindProductService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindProductByCategoryUser retrieves products by category for the user by delegating the operation to the FindAllProductByCategoryService method of the User service.
func (u *UserHandler) FindProductByCategoryUser(ctx context.Context, p *pb.ID) (*pb.UserProductList, error) {
	response, err := u.SVC.FindAllProductByCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindAllProductsUser retrieves all products for the user by delegating the operation to the FindAllProductService method of the User service.
func (u *UserHandler) FindAllProductsUser(ctx context.Context, p *pb.NoParam) (*pb.UserProductList, error) {
	response, err := u.SVC.FindAllProductService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
