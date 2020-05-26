package repositories

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/orders"
	config "github.com/AkezhanOb1/orders/configs"
	"github.com/jackc/pgx/v4"
	"log"
)

//UpdateBusinessServiceOrderRepository is
func UpdateBusinessServiceOrderRepository(ctx context.Context, request *pb.UpdateBusinessServiceOrderRequest, endAt string) (*pb.UpdateBusinessServiceOrderResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())


	sqlQuery := `UPDATE business_company_service_order
				 SET start_at = $1,
					end_at = $2,
					pre_paid = $3,
					client_first_name = $4,
					client_phone_number = $5,
					client_phone_number_prefix = $6,
				 	client_commentary = $7
				 WHERE id = $8 RETURNING created_at, client_id;`


	var createdAt string
	var clientID int64
	businessServiceID := request.GetBusinessServiceID()
	orderID := request.GetOrderID()
	startAt := request.GetStartAt()
	prePaid := request.GetPrePaid()
	clientFirstName := request.GetClientFirstName()
	clientPhoneNumber := request.GetClientPhoneNumber()
	clientPhoneNumberPrefix := request.GetClientPhoneNumberPrefix()
	clientCommentary := request.GetClientCommentary()

	row := conn.QueryRow(
		context.Background(),
		sqlQuery,
		startAt,
		endAt,
		prePaid,
		clientFirstName,
		clientPhoneNumber,
		clientPhoneNumberPrefix,
		clientCommentary,
		orderID,
	)


	err = row.Scan(&createdAt, &clientID)
	if err != nil {
		log.Println(err)
		return nil, err
	}


	return &pb.UpdateBusinessServiceOrderResponse{
		BusinessServiceOrder: &pb.BusinessServiceOrder{
			BusinessServiceOrderID: orderID,
			ClientID:               clientID,
			BusinessServiceID:      businessServiceID,
			StartAt:             	startAt,
			EndAt:                  endAt,
			CreatedAt:              createdAt,
			PrePaid:                prePaid,
			ClientFirstName:		clientFirstName,
			ClientPhoneNumber:		clientPhoneNumber,
			ClientPhoneNumberPrefix: clientPhoneNumberPrefix,
			ClientCommentary:        clientCommentary,
		},
	}, nil
}
