package service

import (
	"github.com/anjush-bhargavan/go_trade_user/config"
	inter "github.com/anjush-bhargavan/go_trade_user/pkg/repo/interfaces"
	"github.com/anjush-bhargavan/go_trade_user/pkg/service/interfaces"
	productpb "github.com/anjush-bhargavan/go_trade_user/pkg/clients/product/pb"
)

// UserService represents a service for user-related operations.
type UserService struct {
	Repo inter.UserRepoInter
	twilio *config.TwilioService
	redis *config.RedisService
	ProductClient productpb.ProductServiceClient
}

// NewUserService creates a new instance of UserService. It returns a UserServiceInter interface representing the created user service.
func NewUserService(repo inter.UserRepoInter,redis *config.RedisService,twilio *config.TwilioService,productClient productpb.ProductServiceClient) interfaces.UserServiceInter {
	return &UserService{
		Repo: repo,
		twilio: twilio,
		redis: redis,
		ProductClient : productClient,
	}
}