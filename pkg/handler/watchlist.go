package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)


// AddToWatchlist adds an item to the user's watchlist.
func (u *UserHandler) AddToWatchlist(ctx context.Context, p *pb.IDs) (*pb.Response, error) {
	response, err := u.SVC.AddWatchlistService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// ViewWatchlist retrieves the user's watchlist.
func (u *UserHandler) ViewWatchlist(ctx context.Context, p *pb.ID) (*pb.UserCategoryList, error) {
	response, err := u.SVC.FetchWatchlistService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// ViewWatchlist retrieves the categorie's watchlist.
func (u *UserHandler) ViewWatchlistUsers(ctx context.Context,p *pb.ID) (*pb.UserList, error) {
	response, err := u.SVC.FetchWatchlistByCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}