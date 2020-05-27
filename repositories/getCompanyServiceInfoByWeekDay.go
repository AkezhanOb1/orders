package repositories

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/order"
	"log"

	config "github.com/AkezhanOb1/orders/configs"
	"github.com/jackc/pgx/v4"
)

//GetCompanyServiceInfoByWeekDayRepository is
func GetCompanyServiceInfoByWeekDayRepository(ctx context.Context, serviceID int64, dayOfWeek int64) (*pb.GetCompanyServiceInfoByWeekDayResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())

	log.Println(serviceID)
	log.Println(dayOfWeek)

	sqlQuery := `SELECT bs.id, bs.name, bs.duration, bs.price, op.day_of_week, op.open_time, op.close_time
				 FROM business_company_service bs
	             INNER JOIN business_company_service_operation_hours op
					ON bs.id = op.business_service_id
	             WHERE  bs.id=$1 AND op.day_of_week=$2 ORDER BY bs.id;`

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

	return &info, nil
}
