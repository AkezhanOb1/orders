package services

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/order"
	db "github.com/AkezhanOb1/orders/repositories"
	"github.com/AkezhanOb1/orders/client"
)

//GetBusinessServiceOrdersByEmail is
func (*BusinessServiceOrder) GetBusinessServiceOrdersByEmail(ctx context.Context, request *pb.GetBusinessServiceOrdersByEmailRequest) (*pb.GetBusinessServiceOrdersByEmailResponse, error) {
	customer, err := client.GetCustomerByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	order, err := db.GetBusinessServiceOrdersByEmailRepository(ctx, request, customer.Customer.GetCustomerID())
	if err != nil {
		return nil, err
	}
	return order, nil
}
