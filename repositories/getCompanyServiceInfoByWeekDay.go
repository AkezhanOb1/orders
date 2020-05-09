
package repositories

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api"
	config "github.com/AkezhanOb1/orders/configs"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

//GetCompanyServiceInfoByWeekDayRepository is
func GetCompanyServiceInfoByWeekDayRepository(ctx context.Context, serviceID int64, dayOfWeek int64) (*pb.GetCompanyServiceInfoByWeekDayResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())

	sqlQuery := `SELECT bs.id, bs.name, bs.duration, bs.price, op.day_of_week, op.open_time, op.close_time
				 FROM business_company_service bs
	             INNER JOIN business_company_service_operation_hours op
					ON bs.id = op.business_service_id
	             WHERE  bs.id=$1 AND op.day_of_week=$2;`



	var info pb.GetCompanyServiceInfoByWeekDayResponse
	row := conn.QueryRow(
		context.Background(),
		sqlQuery,
		serviceID,
		dayOfWeek,
	)


	err = row.Scan(
		&info.CompanyServiceID,
		&info.CompanyServiceName,
		&info.CompanyServiceDuration,
		&info.CompanyServicePrice,
		&info.DayOfWeek,
		&info.OpenTime,
		&info.CloseTime,
	)
	if err != nil {
		return nil, err
	}

	c, _ := time.Parse("2006-01-02","2020-04-08")
	log.Println(c.Weekday())

	return &info, nil
}

