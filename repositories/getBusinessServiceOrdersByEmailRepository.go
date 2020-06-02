package repositories

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api/order"
	config "github.com/AkezhanOb1/orders/configs"
	"github.com/jackc/pgx/v4"
)


//GetBusinessServiceOrdersByEmailRepository is a repository that responsible to all the requests to the database
//about business categories
func GetBusinessServiceOrdersByEmailRepository(ctx context.Context, request *pb.GetBusinessServiceOrdersByEmailRequest, clientID int64) (*pb.GetBusinessServiceOrdersByEmailResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT COUNT(1) from business_company_service_order;`

	err = conn.QueryRow(context.Background(), sqlQuery).Scan(
		&request.Pagination.Count,
	)
	if err != nil {
		return nil, err
	}

	offset := request.Pagination.Offset
	limit := request.Pagination.Limit
	sqlQuery = `SELECT bc.id, bc.client_id, b.name, c.name, b.price, bc.start_at, bc.pre_paid, bc.created_at,
					   bc.client_first_name, bc.client_phone_number, bc.client_phone_number_prefix,
					   bc.client_commentary, bc.end_at, bc.business_service_id
				FROM business_company_service_order bc 
				INNER JOIN business_company_service b
					on bc.business_service_id = b.id 
	            INNER JOIN   business_company c 
				    on b.company_id = c.id 
			    WHERE bc.client_id = $1 
				ORDER BY bc.id
				OFFSET $2 LIMIT $3;`

	rows, err := conn.Query(ctx, sqlQuery, clientID, offset, limit)
	if err != nil {
		return nil, err
	}

	var orders pb.GetBusinessServiceOrdersByEmailResponse
	for rows.Next() {
		var order pb.BusinessServiceOrder

		err = rows.Scan(
			&order.BusinessServiceOrderID,
			&order.ClientID,
			&order.BusinessServiceName,
			&order.BusinessCompanyName,
			&order.Price,
			&order.StartAt,
			&order.PrePaid,
			&order.CreatedAt,
			&order.ClientFirstName,
			&order.ClientPhoneNumber,
			&order.ClientPhoneNumberPrefix,
			&order.ClientCommentary,
			&order.EndAt,
			&order.BusinessServiceID,
		)

		if err != nil {
			return nil, err
		}

		orders.BusinessServicesOrders = append(orders.BusinessServicesOrders, &order)
	}

	orders.Pagination = request.Pagination
	orders.Pagination.Count =  orders.Pagination.Count - (orders.Pagination.Offset + int64(len(orders.BusinessServicesOrders)))
	if orders.Pagination.Count < 0 {
		orders.Pagination.Count = 0
	}

	return &orders, nil
}
