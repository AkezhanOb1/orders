package services

import (
	pb "github.com/AkezhanOb1/orders/api"
	db "github.com/AkezhanOb1/orders/repositories"
	"context"
)

//GetBusinessServiceOrder is
func (*BusinessServiceOrder) GetBusinessServiceOrder(ctx context.Context, request *pb.GetBusinessServiceOrderRequest) (*pb.GetBusinessServiceOrderResponse, error) {
	order, err := db.GetBusinessServiceOrderRepository(ctx, request.BusinessServiceOrderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}
