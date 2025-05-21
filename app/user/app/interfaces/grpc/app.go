package grpc

import (
	"context"
	command_bus "user/app/application/commands"
	"user/app/application/commands/handler"
	"user/app/domain/dto"
	"user/app/domain/services"
	"user/proto/user"
)

var _ user.UserServiceServer = (*grpcApp)(nil)

type GrpcApp interface {
	user.UserServiceServer
}

type grpcApp struct {
	commandBus *command_bus.CommandBus
	srvs       services.Service
}

func NewGrpcApp(srvs services.Service) (GrpcApp, error) {
	commandBus := command_bus.NewCommandBus()
	userService := srvs.UserService()
	// auth
	commandBus.RegisterHandler("LoginCommand", func(ctx context.Context, c command_bus.Command) (interface{}, error) {
		handler := handler.NewLoginCommandHandler(userService)
		return handler.Handle(ctx, c.(*dto.LoginRequest))
	})

	// user
	commandBus.RegisterHandler("CreateUserRequest", func(ctx context.Context, c command_bus.Command) (interface{}, error) {
		handler := handler.NewCreateUserCommandHandler(userService)
		return handler.Handle(ctx, c.(*dto.CreateUserRequest))
	})
	return &grpcApp{
		commandBus: commandBus,
		srvs:       srvs,
	}, nil
}
