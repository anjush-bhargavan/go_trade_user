package di

import (
	"log"

	"github.com/anjush-bhargavan/go_trade_user/config"
	"github.com/anjush-bhargavan/go_trade_user/pkg/clients/product"
	"github.com/anjush-bhargavan/go_trade_user/pkg/db"
	"github.com/anjush-bhargavan/go_trade_user/pkg/handler"
	"github.com/anjush-bhargavan/go_trade_user/pkg/repo"
	"github.com/anjush-bhargavan/go_trade_user/pkg/server"
	"github.com/anjush-bhargavan/go_trade_user/pkg/service"
)

// Init initializes the application by setting up its dependencies.
func Init() {
	cnfg := config.LoadConfig()

	redis, err := config.SetupRedis(cnfg)
	if err != nil {
		log.Fatalf("failed to connect to redis")
	}

	twilio := config.SetupTwilio(cnfg)

	db := db.ConnectDB(cnfg)

	productClient,err := product.ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("failed to connect to product client")
	}

	userRepo := repo.NewUserRepository(db)

	userService := service.NewUserService(userRepo, redis, twilio, productClient)

	userHandler := handler.NewUserHandler(userService)

	err = server.NewGrpcUserServer(cnfg.GrpcPort, userHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}

}
