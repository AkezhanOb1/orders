package services

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/orders"
	db "github.com/AkezhanOb1/orders/repositories"
)

//DeleteBusinessServiceOrder is
func (*BusinessServiceOrder) DeleteBusinessServiceOrder(ctx context.Context,  request *pb.DeleteBusinessServiceOrderRequest) (*pb.DeleteBusinessServiceOrderResponse, error) {
	order, err := db.DeleteBusinessServiceOrderRepository(ctx, request)
	if err != nil {
		return nil, err
	}
	return order, nil
}
