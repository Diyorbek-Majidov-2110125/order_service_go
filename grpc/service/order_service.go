package service

import (
	"context"
	"practice1/order_service_go/config"
	"practice1/order_service_go/genproto/order_service"
	"practice1/order_service_go/genproto/user_service"
	"practice1/order_service_go/grpc/client"
	"practice1/order_service_go/pkg/logger"
	"practice1/order_service_go/storage"

	// "google/protobuf/empty.proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type orderService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	order_service.UnimplementedOrderServiceServer
}

func NewOrderService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *orderService {
	return &orderService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (b *orderService) Create(ctx context.Context, req *order_service.CreateOrderRequest) (resp *order_service.Order, err error) {
	b.log.Info("---CreateOrder--->", logger.Any("req", req))

	userId, err := b.services.UserService().GetById(ctx, &user_service.Pkey{Id: req.UserId})
	if err != nil {
		b.log.Error("!!!CreateOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	productId, err := b.strg.Order().GetById(ctx, &order_service.Pkey{Id: req.ProductId})

	if err != nil {
		b.log.Error("!!!CreateOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = b.strg.Order().Create(ctx, req)
	if err != nil {
		b.log.Error("!!!CreateOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	
	resp = &order_service.Order{
		Id:          resp.Id,
		UserId:   	userId.Id,
		ProductId: 	productId.Id,
		Quantity:   resp.Quantity,
	}
	return resp, nil

}

func (b *orderService) GetById(ctx context.Context, req *order_service.Pkey) (resp *order_service.Order, err error) {
	b.log.Info("---GetOrderById--->", logger.Any("req", req))

	resp, err = b.strg.Order().GetById(ctx, req)
	if err != nil {
		b.log.Error("!!!GetOrderById--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (b *orderService) GetAll(ctx context.Context, req *order_service.GetAllOrdersRequest) (resp *order_service.GetAllOrdersResponse, err error) {
	b.log.Info("---GetAllOrders--->", logger.Any("req", req))

	resp, err = b.strg.Order().GetAll(ctx, req)
	if err != nil {
		b.log.Error("!!!GetAllOrders--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (b *orderService) Update(ctx context.Context, req *order_service.UpdateOrderRequest) (resp *empty.Empty, err error) {
	b.log.Info("---UpdateOrder--->", logger.Any("req", req))

	resp, err = b.strg.Order().Update(ctx, req)
	if err != nil {
		b.log.Error("!!!UpdateOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (b *orderService) Delete(ctx context.Context, req *order_service.Pkey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteOrder--->", logger.Any("req", req))

	resp, err = b.strg.Order().Delete(ctx, req)
	if err != nil {
		b.log.Error("!!!DeleteOrder--->", logger.Error(err))
		return &empty.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (b *orderService) CreateProduct(ctx context.Context, req *order_service.CreateProductRequest) (resp *order_service.Product, err error) {
	b.log.Info("---CreateProduct--->", logger.Any("req", req))

	resp = &order_service.Product{}
	resp, err = b.strg.Product().Create(ctx, req)
	if err != nil {
		b.log.Error("!!!CreateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (b *orderService) GetProductById(ctx context.Context, req *order_service.Primarykey) (resp *order_service.Product, err error) {
	b.log.Info("---GetProductById--->", logger.Any("req", req))

	resp, err = b.strg.Product().GetById(ctx, req)
	if err != nil {
		b.log.Error("!!!GetProductById--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (b *orderService) GetAllProducts(ctx context.Context, req *order_service.GetAllProductsRequest) (resp *order_service.GetAllProductsResponse, err error) {
	b.log.Info("---GetAllProducts--->", logger.Any("req", req))

	resp, err = b.strg.Product().GetAll(ctx, req)
	if err != nil {
		b.log.Error("!!!GetAllProducts--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (b *orderService) DeleteProduct(ctx context.Context, req *order_service.Primarykey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteProduct--->", logger.Any("req", req))

	resp, err = b.strg.Product().Delete(ctx, req)
	if err != nil {
		b.log.Error("!!!DeleteProduct--->", logger.Error(err))
		return &empty.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (b *orderService) UpdateProduct(ctx context.Context, req *order_service.UpdateProductRequest) (resp *empty.Empty, err error) {
	b.log.Info("---UpdateProduct--->", logger.Any("req", req))

	resp, err = b.strg.Product().Update(ctx, req)
	if err != nil {
		b.log.Error("!!!UpdateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil

}
