package services

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api"
	db "github.com/AkezhanOb1/orders/repositories"
)

//GetBusinessServiceOrderByDate is
func (*BusinessServiceOrder) GetBusinessServiceOrderByDate(ctx context.Context,  request *pb.GetBusinessServiceOrderByDateRequest) (*pb.GetBusinessServiceOrderByDateResponse, error) {
	orders, err := db.GetBusinessServiceOrderByDateRepository(ctx, request.GetBusinessServiceID(), request.GetDate())
	if err != nil {
		return nil, err
	}
	return orders, nil
}
