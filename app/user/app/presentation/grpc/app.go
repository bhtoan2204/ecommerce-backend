package grpc

import (
	"context"
	command_bus "user/app/application/commands"
	"user/app/application/commands/command"
	"user/app/application/commands/handler"
	"user/app/domain/usecases"
	"user/proto/user"
)

var _ user.UserServiceServer = (*grpcApp)(nil)

type GrpcApp interface {
	user.UserServiceServer
}

type grpcApp struct {
	commandBus *command_bus.CommandBus
	ucs        usecases.Usecase
}

func NewGrpcApp(ucs usecases.Usecase) (GrpcApp, error) {
	commandBus := command_bus.NewCommandBus()

	registerAuthHandlers(commandBus, ucs)
	registerUserHandlers(commandBus, ucs)

	return &grpcApp{
		commandBus: commandBus,
		ucs:        ucs,
	}, nil
}

func registerAuthHandlers(bus *command_bus.CommandBus, ucs usecases.Usecase) {
	userService := ucs.UserUsecase()

	command_bus.RegisterHandler(
		bus,
		func(ctx context.Context, cmd *command.LoginCommand) (*command.LoginCommandResult, error) {
			handler := handler.NewLoginCommandHandler(userService)
			return handler.Handle(ctx, cmd)
		},
	)
}

func registerUserHandlers(bus *command_bus.CommandBus, ucs usecases.Usecase) {
	userService := ucs.UserUsecase()

	command_bus.RegisterHandler(
		bus,
		func(ctx context.Context, cmd *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
			handler := handler.NewCreateUserCommandHandler(userService)
			return handler.Handle(ctx, cmd)
		},
	)
}
