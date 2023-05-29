package grpc

import (
	"practice1/order_service_go/config"
	"practice1/order_service_go/genproto/order_service"
	"practice1/order_service_go/grpc/client"
	"practice1/order_service_go/grpc/service"
	"practice1/order_service_go/pkg/logger"
	"practice1/order_service_go/storage"
	
	"google.golang.org/grpc"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	order_service.RegisterOrderServiceServer(grpcServer, service.NewOrderService(cfg, log, strg, svcs))

	return grpcServer
}
