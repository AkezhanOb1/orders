package services

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api"
	db "github.com/AkezhanOb1/orders/repositories"
)

type BusinessServiceOrder struct {}

//CreateBusinessServiceOrder is
func (*BusinessServiceOrder) CreateBusinessServiceOrder(ctx context.Context, request *pb.CreateBusinessServiceOrderRequest) (*pb.CreateBusinessServiceOrderResponse, error) {
	order, err := db.CreateBusinessServiceOrderRepository(ctx, request)
	if err != nil {
		return nil, err
	}
	return order, nil
}
