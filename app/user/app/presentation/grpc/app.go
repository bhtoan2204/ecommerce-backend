package grpc

import (
	"context"
	command_bus "user/app/application/commands"
	"user/app/application/commands/command"
	"user/app/application/commands/handler"
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

	registerAuthHandlers(commandBus, srvs)
	registerUserHandlers(commandBus, srvs)

	return &grpcApp{
		commandBus: commandBus,
		srvs:       srvs,
	}, nil
}

func registerAuthHandlers(bus *command_bus.CommandBus, srvs services.Service) {
	userService := srvs.UserService()

	command_bus.RegisterHandler[*command.LoginCommand, *command.LoginCommandResult](
		bus,
		func(ctx context.Context, cmd *command.LoginCommand) (*command.LoginCommandResult, error) {
			handler := handler.NewLoginCommandHandler(userService)
			return handler.Handle(ctx, cmd)
		},
	)
}

func registerUserHandlers(bus *command_bus.CommandBus, srvs services.Service) {
	userService := srvs.UserService()

	command_bus.RegisterHandler[*command.CreateUserCommand, *command.CreateUserCommandResult](
		bus,
		func(ctx context.Context, cmd *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
			handler := handler.NewCreateUserCommandHandler(userService)
			return handler.Handle(ctx, cmd)
		},
	)
}
