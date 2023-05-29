package postgres

import (
	"context"
	"errors"
	"fmt"
	"practice1/order_service_go/genproto/order_service"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}
func (b *orderRepo) Create(ctx context.Context, req *order_service.CreateOrderRequest) (resp *order_service.Order, err error) {
	query := `insert into orders 
		  (id, 
		  product_id, 
		  user_id,
		  quantity
		  ) VALUES (
			$1, 
			$2, 
			$3,
			$4
		  )`

	id := uuid.New().String()
	_, err = b.db.Exec(ctx, query,
		id,
		req.ProductId,
		req.UserId,
		req.Quantity,
	)

	if err != nil {
		return resp, err
	}

	resp = &order_service.Order{
		Id: id,
	}

	return resp, nil
}

func (o *orderRepo) GetById(ctx context.Context, req *order_service.Pkey) (resp *order_service.Order, err error) {
	fmt.Println("GetById function")
	query := `SELECT id, user_id, product_id, quantity FROM orders WHERE id = $1`

	resp = &order_service.Order{}
	err = o.db.QueryRow(ctx, query, req.Id).Scan(&resp.Id, &resp.UserId, &resp.ProductId, &resp.Quantity)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (o *orderRepo) GetAll(ctx context.Context, req *order_service.GetAllOrdersRequest) (resp *order_service.GetAllOrdersResponse, err error) {
	fmt.Println("GetAll function")
	query := `SELECT id, user_id, product_id, quantity FROM orders`

	rows, err := o.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	resp = &order_service.GetAllOrdersResponse{}
	for rows.Next() {
		order := &order_service.Order{}
		err = rows.Scan(&order.Id, &order.UserId, &order.ProductId, &order.Quantity)
		if err != nil {
			return nil, err
		}
		resp.Orders = append(resp.Orders, order)
	}

	return resp, nil
}

func (o *orderRepo) Delete(ctx context.Context, req *order_service.Pkey) (resp *empty.Empty, err error) {
	fmt.Println("Delete function")
	query := `DELETE FROM orders WHERE id = $1`

	_, err = o.db.Exec(ctx, query, req.Id)
	if err != nil {
		return resp, errors.New("Order not found")
	}

	return &empty.Empty{}, nil
}

func (o *orderRepo) Update(ctx context.Context, req *order_service.UpdateOrderRequest) (resp *empty.Empty, err error) {
	
	query := `UPDATE orders SET user_id = $2, product_id = $3, quantity = $4 WHERE id = $1`

	
	_, err = o.db.Exec(ctx, query,  req.Id,req.UserId, req.ProductId, req.Quantity)
	if err != nil {
		return resp, errors.New("Order not found")
	}

	return &empty.Empty{}, nil
}