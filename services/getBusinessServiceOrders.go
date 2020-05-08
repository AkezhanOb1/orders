package services

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api"
	db "github.com/AkezhanOb1/orders/repositories"
)

//GetBusinessServiceOrders is
func (*BusinessServiceOrder) GetBusinessServiceOrders(ctx context.Context,  request *pb.GetBusinessServiceOrdersRequest) (*pb.GetBusinessServiceOrdersResponse, error) {
	orders, err := db.GetBusinessServiceOrdersRepository(ctx, request.GetBusinessServiceID())
	if err != nil {
		return nil, err
	}
	return orders, nil
}
