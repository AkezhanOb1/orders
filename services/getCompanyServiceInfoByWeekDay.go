package services

import (
"context"
pb "github.com/AkezhanOb1/orders/api"
db "github.com/AkezhanOb1/orders/repositories"
)

//GetCompanyServiceInfoByWeekDay is
func (*BusinessServiceOrder) GetCompanyServiceInfoByWeekDay(ctx context.Context,  request *pb.GetCompanyServiceInfoByWeekDayRequest) (*pb.GetCompanyServiceInfoByWeekDayResponse, error) {
	serviceInfo, err := db.GetCompanyServiceInfoByWeekDayRepository(ctx, request.GetBusinessServiceID(), request.GetDayOfWeek())
	if err != nil {
		return nil, err
	}
	return serviceInfo, nil
}
