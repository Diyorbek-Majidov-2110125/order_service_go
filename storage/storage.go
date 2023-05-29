package storage

import (
	"context"
	"practice1/order_service_go/genproto/order_service"

	"github.com/golang/protobuf/ptypes/empty"
)

type StorageI interface {
	CloseDB()
	Product() ProductRepoI
	Order() OrderRepoI
}

type ProductRepoI interface {
	Create(ctx context.Context, req *order_service.CreateProductRequest) (resp *order_service.Product, err error)
	GetById(ctx context.Context, req *order_service.Primarykey) (resp *order_service.Product, err error)
	GetAll(ctx context.Context, req *order_service.GetAllProductsRequest) (resp *order_service.GetAllProductsResponse, err error)
	Delete(ctx context.Context, req *order_service.Primarykey) (resp *empty.Empty, err error)
	Update(ctx context.Context, req *order_service.UpdateProductRequest) (resp *empty.Empty, err error)
}
type OrderRepoI interface {
	Create(ctx context.Context, req *order_service.CreateOrderRequest) (resp *order_service.Order, err error)
	GetById(ctx context.Context, req *order_service.Pkey) (resp *order_service.Order, err error)
	GetAll(ctx context.Context, req *order_service.GetAllOrdersRequest) (resp *order_service.GetAllOrdersResponse, err error)
	Delete(ctx context.Context, req *order_service.Pkey) (resp *empty.Empty, err error)
	Update(ctx context.Context, req *order_service.UpdateOrderRequest) (resp *empty.Empty, err error)
}
