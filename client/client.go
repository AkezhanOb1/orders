package main

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer cc.Close()


	c := pb.NewBusinessServiceOrdersClient(cc)
	e := pb.GetCompanyAvailableHoursByDateRequest{
		BusinessServiceID: 6,
		Date:              "2020-03-11",
	}

	orders, err := c.GetCompanyAvailableHoursByDate(context.Background(), &e)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(orders.AvailableHour)

}
