package service

import (
	"context"

	productpb "github.com/anjush-bhargavan/go_trade_user/pkg/clients/product/pb"
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
)

func (u *UserService) FindProductService(p *pb.ID) (*pb.UserProduct, error) {
	ctx := context.Background()

	Product := &productpb.PrID{
		ID: p.ID,
	}

	result, err := u.ProductClient.FindProductByID(ctx, Product)
	if err != nil {
		return nil, err
	}

	return &pb.UserProduct{
		Product_ID: result.Product_ID,
		Name:       result.Name,
		Seller_ID: result.Seller_ID,
		Category: &pb.UserCategory{
			Category_ID: result.Category.Category_ID,
			Name:        result.Category.Name,
			Description: result.Category.Description,
		},
		Base_Price:         result.Base_Price,
		Bidding_Type:       result.Bidding_Type,
		Bidding_Start_Time: result.Bidding_Start_Time,
		Bidding_End_Time:   result.Bidding_End_Time,
		Listed_On:          result.Listed_On,
		Status:             result.Status,
	}, nil
}

func (u *UserService) FindAllProductService(p *pb.NoParam) (*pb.UserProductList,error) {
	ctx := context.Background()

	

	result, err := u.ProductClient.FindAllProducts(ctx,&productpb.ProductNoParam{})
	if err != nil {
		return nil, err
	}

	var products []*pb.UserProduct

	for _,product := range result.Products {
		products = append(products, &pb.UserProduct{
			Product_ID: product.Product_ID,
			Name:       product.Name,
			Seller_ID:  product.Seller_ID,
			Category: &pb.UserCategory{
				Category_ID: product.Category.Category_ID,
				Name:        product.Category.Name,
				Description: product.Category.Description,
			},
			Base_Price:         product.Base_Price,
			Bidding_Type:       product.Bidding_Type,
			Bidding_Start_Time: product.Bidding_Start_Time,
			Bidding_End_Time:   product.Bidding_End_Time,
			Listed_On:          product.Listed_On,
			Status:             product.Status,
		})
	}

	return &pb.UserProductList{
		Products: products,
	}, nil
}

func (u *UserService) FindAllProductByCategoryService(p *pb.ID) (*pb.UserProductList,error) {
	ctx := context.Background()

	result, err := u.ProductClient.FindProductByCategory(ctx,&productpb.PrID{ID: p.ID})
	if err != nil {
		return nil, err
	}

	var products []*pb.UserProduct

	for _,product := range result.Products {
		products = append(products, &pb.UserProduct{
			Product_ID: product.Product_ID,
			Name:       product.Name,
			Seller_ID:  product.Seller_ID,
			Category: &pb.UserCategory{
				Category_ID: product.Category.Category_ID,
				Name:        product.Category.Name,
				Description: product.Category.Description,
			},
			Base_Price:         product.Base_Price,
			Bidding_Type:       product.Bidding_Type,
			Bidding_Start_Time: product.Bidding_Start_Time,
			Bidding_End_Time:   product.Bidding_End_Time,
			Listed_On:          product.Listed_On,
			Status:             product.Status,
		})
	}

	return &pb.UserProductList{
		Products: products,
	}, nil
}