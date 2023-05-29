package client

import (
	"practice1/order_service_go/config"
	"practice1/order_service_go/genproto/user_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	UserService() user_service.UserServiceClient
}

type grpcClients struct {
	userService user_service.UserServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connUserService, err := grpc.Dial(
		cfg.UserServiceHost+cfg.UserServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &grpcClients{
		userService: user_service.NewUserServiceClient(connUserService),
	}, nil
}

func (c *grpcClients) UserService() user_service.UserServiceClient {
	return c.userService
}
