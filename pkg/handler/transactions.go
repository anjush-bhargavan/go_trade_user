package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)


func (u *UserHandler) CreateTransaction(ctx context.Context,p *pb.Transaction) (*pb.Response, error) {
	response, err := u.SVC.CreateTransactionService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) FindTransactionByUser(ctx context.Context,p *pb.ID) (*pb.TransactionList, error) {
	response, err := u.SVC.FindTransactionsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) FindAllTransactions(ctx context.Context,p *pb.NoParam) (*pb.TransactionList, error) {
	response, err := u.SVC.FindAllTransactionsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

