package grpc

import (
	"context"
	command_bus "user/app/application/commands"
	"user/app/application/commands/command"
	"user/app/presentation/dto"
	"user/package/logger"
	"user/proto/user"

	"go.uber.org/zap"
)

func (g *grpcApp) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	log := logger.FromContext(ctx)
	request := &dto.LoginRequest{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	response, err := command_bus.Dispatch[*command.LoginCommand, *command.LoginCommandResult](
		g.commandBus, ctx, &command.LoginCommand{
			Email:    request.Email,
			Password: request.Password,
		},
	)
	if err != nil {
		log.Error("Failed to Login", zap.Error(err))
		return nil, err
	}

	return response.ToDto().ToPb(), nil
}
