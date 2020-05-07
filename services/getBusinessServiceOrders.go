package services

import (
	pb "github.com/AkezhanOb1/orders/api"
	db "github.com/AkezhanOb1/orders/repositories"
	"context"
	"github.com/golang/protobuf/ptypes/empty"

)

//GetBusinessServiceOrders is
func (*BusinessServiceOrder) GetBusinessServiceOrders(ctx context.Context, emp *empty.Empty) (*pb.GetBusinessServiceOrdersResponse, error) {
	orders, err := db.GetBusinessServiceOrdersRepository(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
