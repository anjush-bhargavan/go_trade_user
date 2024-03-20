package handler

import (
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
	inter "github.com/anjush-bhargavan/go_trade_user/pkg/service/interfaces"
)

// UserHandler represents the service methods and gRPC server for user-related operations.
type UserHandler struct {
	SVC inter.UserServiceInter
	pb.UserServiceServer
}

// NewUserHandler creates a new instance of UserHandler with the provided user service interface.
func NewUserHandler(svc inter.UserServiceInter) *UserHandler {
	return &UserHandler{
		SVC: svc,
	}
}
