package client

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/customer"
	config "github.com/AkezhanOb1/orders/configs"
	"google.golang.org/grpc"
	"log"
)

//GetCompanyServiceInfo
func GetCustomerByEmail(email string) (*pb.GetCustomerByEmailResponse, error) {
	cc, err := grpc.Dial(config.CustomerServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cc.Close()


	c := pb.NewCustomerServiceClient(cc)
	e := pb.GetCustomerByEmailRequest{
		Email:	email,
	}

	customer, err := c.GetCustomerByEmail(context.Background(), &e)
	if err != nil {
		return nil, err
	}

	return customer, nil

}
