package services

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api"
	"github.com/AkezhanOb1/orders/client"
	db "github.com/AkezhanOb1/orders/repositories"
	"time"
)

type BusinessServiceOrder struct {}

//CreateBusinessServiceOrder is
func (*BusinessServiceOrder) CreateBusinessServiceOrder(ctx context.Context, request *pb.CreateBusinessServiceOrderRequest) (*pb.CreateBusinessServiceOrderResponse, error) {
	service, err := client.GetCompanyServiceInfo(request.GetBusinessServiceID())
	if err != nil {
		return nil, err
	}

	duration := time.Duration(service.CompanyService.CompanyServiceDuration)
	startAt, err := time.Parse(time.RFC3339, request.GetStartAt())
	if err != nil {
		return nil, err
	}
	endAt := (startAt.Add(time.Minute * duration)).Format(time.RFC3339)
	request.StartAt = startAt.Format(time.RFC3339)

	order, err := db.CreateBusinessServiceOrderRepository(ctx, request, endAt)
	if err != nil {
		return nil, err
	}
	return order, nil
}
