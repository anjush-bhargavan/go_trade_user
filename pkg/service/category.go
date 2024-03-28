package service

import (
	"context"

	productpb "github.com/anjush-bhargavan/go_trade_user/pkg/clients/product/pb"
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

func (u *UserService) FindCategoryService(p *pb.ID) (*pb.UserCategory, error) {
	ctx := context.Background()

	category := &productpb.PrID{
		ID: p.ID,
	}

	result, err := u.ProductClient.FindCategoryByID(ctx, category)
	if err != nil {
		return nil, err
	}

	return &pb.UserCategory{
		Category_ID: result.Category_ID,
		Name:        result.Name,
		Description: result.Description,
	}, nil
}

func (u *UserService) FindAllCategoryService(p *pb.NoParam) (*pb.UserCategoryList, error) {
	ctx := context.Background()

	result, err := u.ProductClient.FindAllCategories(ctx, &productpb.ProductNoParam{})
	if err != nil {
		return nil, err
	}

	var categories []*pb.UserCategory

	for _, category := range result.Categories {
		categories = append(categories, &pb.UserCategory{
			Category_ID: category.Category_ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.UserCategoryList{
		Categories: categories,
	}, nil
}
