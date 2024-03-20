package repo

import ("gorm.io/gorm"
		inter "github.com/anjush-bhargavan/go_trade_user/pkg/repo/interfaces"
)

// UserRepository for connecting db to UserRepoInter methods
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates an instance of user repo
func NewUserRepository(db *gorm.DB) inter.UserRepoInter {
	return &UserRepository{
		DB: db,
	}
}  