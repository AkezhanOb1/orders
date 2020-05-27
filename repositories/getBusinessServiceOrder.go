package repositories

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/order"
	config "github.com/AkezhanOb1/orders/configs"
	"github.com/jackc/pgx/v4"
)

//GetBusinessServiceOrderRepository is a
func GetBusinessServiceOrderRepository(ctx context.Context, orderID int64) (*pb.GetBusinessServiceOrderResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())


	var order pb.BusinessServiceOrder
	sqlQuery := `SELECT id, client_id, business_service_id, start_at, end_at, pre_paid, created_at,
    				  client_first_name, client_phone_number, client_phone_number_prefix, client_commentary
			     FROM business_company_service_order 
	             WHERE id=$1 ORDER BY id;`

	err = conn.QueryRow(ctx, sqlQuery, orderID).Scan(
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


	return &pb.GetBusinessServiceOrderResponse{
		BusinessServiceOrder: &order,
	}, nil
}

