package main

import (
	pb "github.com/AkezhanOb1/orders/api/order"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/AkezhanOb1/orders/services"
)

func main() {
	var address = "0.0.0.0:50053"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	pb.RegisterBusinessServiceOrdersServer(s, &services.BusinessServiceOrder{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
//
//package main
//
//import (
//	"fmt"
//	"log"
//	"time"
//)
//
//func main() {
//	a, _ := time.Parse(time.RFC3339, "2011-10-05T14:48:00.000Z")
//	fmt.Println(a)
//	log.Println(a.Year())
//}