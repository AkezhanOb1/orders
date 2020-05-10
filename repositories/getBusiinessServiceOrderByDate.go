package repositories

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api"
	config "github.com/AkezhanOb1/orders/configs"
	"github.com/jackc/pgx/v4"
)


//GetBusinessServiceOrderByDateRepository is a repository that responsible to all the requests to DB
//about business categories
func GetBusinessServiceOrderByDateRepository(ctx context.Context, businessServiceID int64, today string, nextDay string) (*pb.GetBusinessServiceOrderByDateResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)


	sqlQuery := `SELECT id, client_id, business_service_id, start_at, end_at, 
					pre_paid, created_at,client_first_name, client_phone_number, 
					client_phone_number_prefix, client_commentary
	             FROM business_company_service_order 
				 WHERE business_service_id=$1 
				 AND start_at>$2 AND start_at<$3`

	rows, err := conn.Query(
		ctx,
		sqlQuery,
		businessServiceID,
		today,
		nextDay)

	if err != nil {
		return nil, err
	}

	var orders pb.GetBusinessServiceOrderByDateResponse

	for rows.Next() {
		var order pb.BusinessServiceOrder

		err = rows.Scan(
			&order.BusinessServiceOrderID,
			&order.ClientID,
			&order.BusinessServiceID,
			&order.StartAt,
			&order.EndAt,
			&order.PrePaid,
			&order.CreatedAt,
			&order.ClientFirstName,
			&order.ClientPhoneNumber,
			&order.ClientPhoneNumberPrefix,
			&order.ClientCommentary,
		)

		if err != nil {
			return nil, err
		}

		orders.BusinessServicesOrders = append(orders.BusinessServicesOrders, &order)
	}

	return &orders, nil
}
