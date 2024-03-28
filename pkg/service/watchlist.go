package service

import (
	"context"

	productpb "github.com/anjush-bhargavan/go_trade_user/pkg/clients/product/pb"
	"github.com/anjush-bhargavan/go_trade_user/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

// AddWatchlistService method will help to add category items in the users watchlist
func (u *UserService) AddWatchlistService(p *pb.IDs) (*pb.Response, error) {
	item := model.Wathlist{
		CategoryID: uint(p.ID),
		UserID:     uint(p.User_ID),
	}

	err := u.Repo.CreateWatchlist(&item)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to add item in watchlist",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "item added to watchlist successfully",
	}, nil
}

// FetchWatchlistService handles the users to view their watchlist
func (u *UserService) FetchWatchlistService(p *pb.ID) (*pb.UserCategoryList, error) {
	ctx := context.Background()

	items, err := u.Repo.FetchWatchlist(uint(p.ID))
	if err != nil {
		return nil, err
	}

	var categories []*pb.UserCategory

	for _, v := range *items {
		result, err := u.ProductClient.FindCategoryByID(ctx, &productpb.PrID{ID: uint32(v.CategoryID)})
		if err != nil {
			return nil, err
		}
		categories = append(categories, &pb.UserCategory{
			Category_ID: result.Category_ID,
			Name:        result.Name,
			Description: result.Description,
		})
	}

	return &pb.UserCategoryList{
		Categories: categories,
	}, nil
}
