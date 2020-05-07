package repositories

import (
	pb "github.com/AkezhanOb1/orders/api"
	config "github.com/AkezhanOb1/orders/configs"
	"github.com/jackc/pgx/v4"
	"context"
)


//GetBusinessServiceOrdersRepository is a repository that responsible to all the requests to DB
//about business categories
func GetBusinessServiceOrdersRepository(ctx context.Context) (*pb.GetBusinessServiceOrdersResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)


	sqlQuery := `SELECT id, client_id, business_service_id, order_date, pre_paid, created_at,
     				client_first_name, client_phone_number, client_phone_number_prefix, client_commentary
	             FROM business_company_service_order`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	var orders pb.GetBusinessServiceOrdersResponse

	for rows.Next() {
		var order pb.BusinessServiceOrder

		err = rows.Scan(
			&order.BusinessServiceOrderID,
			&order.ClientID,
			&order.BusinessServiceID,
			&order.OrderDate,
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
