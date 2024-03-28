package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

// FindCategory retrieves a category by its ID.
func (u *UserHandler) FindCategory(ctx context.Context, p *pb.ID) (*pb.UserCategory, error) {
	response, err := u.SVC.FindCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindCategories retrieves all categories.
func (u *UserHandler) FindCategories(ctx context.Context, p *pb.NoParam) (*pb.UserCategoryList, error) {
	response, err := u.SVC.FindAllCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
