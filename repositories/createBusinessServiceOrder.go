package repositories

import (
	"context"
	config "github.com/AkezhanOb1/orders/configs"
	"github.com/jackc/pgx/v4"
	pb "github.com/AkezhanOb1/orders/api"
	"time"
)

//CreateBusinessServiceOrderRepository is
func CreateBusinessServiceOrderRepository(ctx context.Context, request *pb.CreateBusinessServiceOrderRequest) (*pb.CreateBusinessServiceOrderResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())


	sqlQuery := `INSERT INTO business_company_service_order (client_id, business_service_id, order_date, pre_paid, created_at, 
     					client_first_name, client_phone_number, client_phone_number_prefix, client_commentary) 
				 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
				 RETURNING id;`


	var businessServiceOrderID int64
	clientID := request.GetClientID()
	businessServiceID := request.GetBusinessServiceID()
	orderDate := request.GetOrderDate()
	prePaid := request.GetPrePaid()
	createdAt := time.Now().Format(time.RFC3339)
	clientFirstName := request.GetClientFirstName()
	clientPhoneNumber := request.GetClientPhoneNumber()
	clientPhoneNumberPrefix := request.GetClientPhoneNumberPrefix()
	clientCommentary := request.GetClientCommentary()

	row := conn.QueryRow(
		context.Background(),
		sqlQuery,
		clientID,
		businessServiceID,
		orderDate,
		prePaid,
		createdAt,
		clientFirstName,
		clientPhoneNumber,
		clientPhoneNumberPrefix,
		clientCommentary,
		)


	err = row.Scan(&businessServiceOrderID)
	if err != nil {
		return nil, err
	}


	return &pb.CreateBusinessServiceOrderResponse{
		BusinessServiceOrder: &pb.BusinessServiceOrder{
			BusinessServiceOrderID: businessServiceOrderID,
			ClientID:               clientID,
			BusinessServiceID:      businessServiceID,
			OrderDate:              orderDate,
			CreatedAt:              createdAt,
			PrePaid:                prePaid,
			ClientFirstName:		clientFirstName,
			ClientPhoneNumber:		clientPhoneNumber,
			ClientPhoneNumberPrefix: clientPhoneNumberPrefix,
			ClientCommentary:        clientCommentary,
		},
	}, nil
}
