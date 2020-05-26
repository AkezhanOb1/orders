package services

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/orders"
	db "github.com/AkezhanOb1/orders/repositories"
)

//GetBusinessServiceOrder is
func (*BusinessServiceOrder) GetBusinessServiceOrder(ctx context.Context, request *pb.GetBusinessServiceOrderRequest) (*pb.GetBusinessServiceOrderResponse, error) {
	order, err := db.GetBusinessServiceOrderRepository(ctx, request.BusinessServiceOrderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}
