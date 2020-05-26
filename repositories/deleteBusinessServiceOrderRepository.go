package repositories

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/orders"
	config "github.com/AkezhanOb1/orders/configs"
	"github.com/jackc/pgx/v4"
	"log"
)

//CreateBusinessServiceOrderRepository is
func DeleteBusinessServiceOrderRepository(ctx context.Context, request *pb.DeleteBusinessServiceOrderRequest) (*pb.DeleteBusinessServiceOrderResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())


	sqlQuery := `DELETE FROM business_company_service_order WHERE id=$1 RETURNING *;`


	row := conn.QueryRow(
		context.Background(),
		sqlQuery,
		request.GetOrderID(),
	)

	var order pb.BusinessServiceOrder

	err = row.Scan(
		&order.BusinessServiceOrderID,
		&order.ClientID,
		&order.BusinessServiceID,
		&order.StartAt,
		&order.PrePaid,
		&order.CreatedAt,
		&order.ClientFirstName,
		&order.ClientPhoneNumber,
		&order.ClientPhoneNumberPrefix,
		&order.ClientCommentary,
		&order.EndAt,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}


	return &pb.DeleteBusinessServiceOrderResponse{
		BusinessServiceOrder: &order,
	}, nil
}
