package postgres

import (
	"context"
	"errors"
	"practice1/order_service_go/genproto/order_service"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}
func (b *productRepo) Create(ctx context.Context, req *order_service.CreateProductRequest) (resp *order_service.Product, err error) {
	query := `insert into orders 
		  (id, 
		  name, 
		  description,
		  price
		  ) VALUES (
			$1, 
			$2, 
			$3,
			$4
		  )`

	id := uuid.New().String()
	_, err = b.db.Exec(ctx, query,
	  id,
	  req.Name,
	  req.Description,
	  req.Price,
	)
  
	if err != nil {
	  return resp, err
	}
  
	resp = &order_service.Product{
		Id : req.Id,
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
	}
  
	return resp, nil
  }
// func (o productRepo) Create(ctx context.Context, req *order_service.CreateProductRequest) (resp *order_service.Primarykey, err error) {
// 	query := `INSERT INTO products (id, name, description, price) VALUES ($1, $2, $3, $4)`

// 	// id := "03a4c6c6-3f26-4ae4-aa6f-c943659cf895"
// 	id := uuid.New().String()
// 	_, err = o.db.Exec(ctx, query, id, req.Name, req.Description, req.Price)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp = &order_service.Primarykey{Id: id}
// 	return resp, nil

// }

func (o *productRepo) GetById(ctx context.Context, req *order_service.Primarykey) (resp *order_service.Product, err error) {
	query := `SELECT id,name, description, price FROM products WHERE id = $1`

	resp = &order_service.Product{}
	err = o.db.QueryRow(ctx, query, req.Id).Scan(&resp.Id, &resp.Name, &resp.Description, &resp.Price)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (o *productRepo) GetAll(ctx context.Context, req *order_service.GetAllProductsRequest) (resp *order_service.GetAllProductsResponse, err error) {
	query := `SELECT id, name, description, price FROM products`

	rows, err := o.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	resp = &order_service.GetAllProductsResponse{}
	for rows.Next() {
		product := &order_service.Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		resp.Products = append(resp.Products, product)
	}

	return resp, nil

}

func (o *productRepo) Delete(ctx context.Context, req *order_service.Primarykey) (resp *empty.Empty, err error) {
	query := `DELETE FROM products WHERE id = $1`

	_, err = o.db.Exec(ctx, query, req.Id)
	if err != nil {
		return &empty.Empty{}, errors.New("error deleting product")
	}
	return &empty.Empty{}, nil

}

func (o *productRepo) Update(ctx context.Context, req *order_service.UpdateProductRequest) (resp *empty.Empty, err error) {
	query := `UPDATE products SET name = $2, description = $3, price = $4 WHERE id = $1`

	_, err = o.db.Exec(ctx, query, req.Name, req.Description, req.Price, req.Id)
	if err != nil {
		return nil, errors.New("error updating product")
	}
	
	return &empty.Empty{}, nil
}