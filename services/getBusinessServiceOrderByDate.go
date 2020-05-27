package services

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/order"
	db "github.com/AkezhanOb1/orders/repositories"
	"time"
)

//GetBusinessServiceOrderByDate is
func (*BusinessServiceOrder) GetBusinessServiceOrderByDate(ctx context.Context,  request *pb.GetBusinessServiceOrderByDateRequest) (*pb.GetBusinessServiceOrderByDateResponse, error) {
	date, err := time.Parse("2006-01-02", request.GetDate())
	if err != nil {
		return nil, err
	}

	nextDay := date.Add(time.Hour * 24)
	orders, err := db.GetBusinessServiceOrderByDateRepository(ctx, request.GetBusinessServiceID(),date.String(), nextDay.String())
	if err != nil {
		return nil, err
	}
	return orders, nil
}
