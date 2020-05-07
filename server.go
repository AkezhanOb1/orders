package main

import (
	"log"
	"net"
	"google.golang.org/grpc"

	"github.com/AkezhanOb1/orders/services"
	pb "github.com/AkezhanOb1/orders/api"

)

func main() {
	address := "0.0.0.0:50053"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	pb.RegisterBusinessServicesServer(s, &services.BusinessServiceOrder{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}