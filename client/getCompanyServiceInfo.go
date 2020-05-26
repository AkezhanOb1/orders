package client

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/companyService"
	config "github.com/AkezhanOb1/orders/configs"
	"google.golang.org/grpc"
	"log"
)

//GetCompanyServiceInfo
func GetCompanyServiceInfo(serviceID int64) (*pb.GetCompanyServiceResponse, error) {
	cc, err := grpc.Dial(config.MainService, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cc.Close()


	c := pb.NewCompanyServicesClient(cc)
	e := pb.GetCompanyServiceRequest{
		CompanyServiceID:	serviceID,
	}

	service, err := c.GetCompanyService(context.Background(), &e)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return service, nil

}
