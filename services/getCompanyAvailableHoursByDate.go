package services

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api"
	db "github.com/AkezhanOb1/orders/repositories"
	"log"
	"time"
)

//GetCompanyAvailableHoursByDate is
func (*BusinessServiceOrder) GetCompanyAvailableHoursByDate(ctx context.Context,  request *pb.GetCompanyAvailableHoursByDateRequest) (*pb.GetCompanyAvailableHoursByDateResponse, error) {
	//getting the list of orders which were made in the provided date

	date, err := time.Parse("2006-01-02", request.GetDate())
	if err != nil {
		return nil, err
	}

	nextDay := date.Add(time.Hour * 24)
	orders, err := db.GetBusinessServiceOrderByDateRepository(ctx, request.BusinessServiceID, date.String(), nextDay.String())
	if err != nil {
		return nil, err
	}

	//determining the week day based on date
	var dayOfWeek int64
	switch date.Weekday().String() {
	case "Sunday":
		dayOfWeek = 0
	case "Monday":
		dayOfWeek = 1
	case "Tuesday":
		dayOfWeek = 2
	case "Wednesday":
		dayOfWeek = 3
	case "Thursday":
		dayOfWeek = 4
	case "Friday":
		dayOfWeek = 5
	case "Saturday":
		dayOfWeek = 6
	}

	//getting company service schedule
	info, err := db.GetCompanyServiceInfoByWeekDayRepository(ctx, request.GetBusinessServiceID(), dayOfWeek)
	if err != nil {
		return nil, err
	}



	//requestDate, _ := time.Parse(time.RFC3339, request.GetDate())
	//log.Println(requestDate)
	duration := time.Duration(info.GetCompanyServiceDuration())
	startDate, _ := time.Parse("15:04:05", info.GetOpenTime())
	startDate = startDate.AddDate(date.Year(),int(date.Month()-1) , date.Day()-1)
	endDate, _ := time.Parse("15:04:05", info.GetCloseTime())
	endDate = endDate.AddDate(date.Year(),int(date.Month()-1) , date.Day()-1)

	//determining all possible order hours
	var workTime []time.Time
	for startDate.Add(time.Minute * duration).Before(endDate.Add(time.Minute * 1)) {
		workTime = append(workTime, startDate)
		startDate = startDate.Add(time.Minute * duration)
	}

	//removing from possible hours already booked hours
	for i := range orders.BusinessServicesOrders {
		log.Println(orders.BusinessServicesOrders[i].StartAt)
		orderDate, _ := time.Parse(time.RFC3339, orders.BusinessServicesOrders[i].StartAt)

		for j := range workTime {
			if orderDate.Equal(workTime[j]) {
				workTime = append(workTime[:j], workTime[j+1:]...)
				break
			}
		}
	}

	//converting list of time.Time to string slice
	var availableTime []string
	for i := range workTime {
		availableTime = append(availableTime, workTime[i].String())
	}

	return &pb.GetCompanyAvailableHoursByDateResponse{
		AvailableHour:availableTime,
	}, nil
}
