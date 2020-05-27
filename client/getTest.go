package client

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/order"
	"google.golang.org/grpc"
	"log"
)

//GetCompanyServiceInfo
func Test()  {
	cc, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer cc.Close()


	c := pb.NewBusinessServiceOrdersClient(cc)
	e := pb.GetBusinessServiceOrdersRequest{
		BusinessServiceID: 11,
	}

	service, err := c.GetBusinessServiceOrders(context.Background(), &e)
	if err != nil {
		log.Println(err)

	}

	log.Println(service)

}
