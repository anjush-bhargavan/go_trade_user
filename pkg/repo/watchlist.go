package repo

import "github.com/anjush-bhargavan/go_trade_user/pkg/model"

// CreateWatchlist will create a watchlist field in the database.
func (u *UserRepository) CreateWatchlist(item *model.Wathlist) error {
	if err := u.DB.Create(&item).Error; err != nil {
		return err
	}
	return nil
}

// FetchWatchlistByUser will fetch the watchlist of the user from database.
func (u *UserRepository) FetchWatchlistByUser(userID uint) (*[]model.Wathlist,error) {
	var watchlist []model.Wathlist
	if err := u.DB.Where("user_id = ?",userID).Find(&watchlist).Error; err != nil {
		return nil,err
	}
	return &watchlist,nil
}

// FetchWatchlistByCategory will fetch the watchlist of the category from database.
func (u *UserRepository) FetchWatchlistByCategory(categoryID uint) (*[]model.Wathlist,error) {
	var watchlist []model.Wathlist
	if err := u.DB.Where("category_id = ?",categoryID).Find(&watchlist).Error; err != nil {
		return nil,err
	}
	return &watchlist,nil
}