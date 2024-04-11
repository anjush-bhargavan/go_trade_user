package service

import (
	"github.com/anjush-bhargavan/go_trade_user/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

func (u *UserService) CreateTransactionService(p *pb.Transaction) (*pb.Response, error) {
	transaction := model.Transaction{
		UserID: uint(p.User_ID),
		Name: p.Name,
		Amount: float64(p.Amount),
	}

	err := u.Repo.CreateTransaction(&transaction)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to create transaction",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Transaction created successfully",
	}, nil
}

// FindTransactionsService handles the users to view their transactions
func (u *UserService) FindTransactionsService(p *pb.ID) (*pb.TransactionList, error) {

	transactions, err := u.Repo.FindTransactionByUserID(uint(p.ID))
	if err != nil {
		return nil, err
	}

	var transactionList []*pb.Transaction

	for _, v := range *transactions {
	
		transactionList = append(transactionList, &pb.Transaction{
		Transaction_ID: uint32(v.ID),
		User_ID: uint32(v.UserID),
		Name: v.Name,
		Amount: float32(v.Amount),
		})
	}

	return &pb.TransactionList{
		Transactions: transactionList,
	}, nil
}

// FindAllTransactionsService handles the admin to view  transactions
func (u *UserService) FindAllTransactionsService(p *pb.NoParam) (*pb.TransactionList, error) {

	transactions, err := u.Repo.FindAllTransactions()
	if err != nil {
		return nil, err
	}

	var transactionList []*pb.Transaction

	for _, v := range *transactions {
	
		transactionList = append(transactionList, &pb.Transaction{
		Transaction_ID: uint32(v.ID),
		User_ID: uint32(v.UserID),
		Name: v.Name,
		Amount: float32(v.Amount),
		})
	}

	return &pb.TransactionList{
		Transactions: transactionList,
	}, nil
}