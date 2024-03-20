package service

import (
	"github.com/anjush-bhargavan/go_trade_user/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
	"gorm.io/gorm"
)

// AddAddressService manages to add address to the database.
func (u *UserService) AddAddressService(p *pb.Address) (*pb.Response, error) {

	address := &model.Address{
		House:  p.House,
		Street: p.Street,
		City:   p.City,
		ZIP:    uint(p.Zip),
		State:  p.State,
		UserID: uint(p.User_ID),
	}

	err := u.Repo.CreateAddress(address)
	if err != nil {
		return &pb.Response{
			Status:  "Failed",
			Message: "Error adding address",
		}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: "Address added successfully",
	}, nil

}

// EditAddressService updates any change in the address to database.
func (u *UserService) EditAddressService(p *pb.Address) (*pb.Address, error) {

	address := &model.Address{
		Model: gorm.Model{
			ID: uint(p.ID),
		},
		House:  p.House,
		Street: p.Street,
		City:   p.City,
		ZIP:    uint(p.Zip),
		State:  p.State,
		UserID: uint(p.User_ID),
	}

	err := u.Repo.EditAddress(address)
	if err != nil {
		return nil, err
	}

	return p, nil

}

// RemoveAddressService removes an address from database.
func (u *UserService) RemoveAddressService(p *pb.IDs) (*pb.Response, error) {

	err := u.Repo.RemoveAddress(uint(p.ID),uint(p.User_ID))
	if err != nil {
		return &pb.Response{
			Status:  "Failed",
			Message: "Error deleting address",
		}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: "Address removed successfully",
	}, nil

}
