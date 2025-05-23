package grpc

import (
	"context"
	"time"
	command_bus "user/app/application/commands"
	"user/app/application/commands/command"
	"user/app/presentation/dto"
	"user/package/logger"
	"user/proto/user"

	"go.uber.org/zap"
)

func (g *grpcApp) GetProfile(context.Context, *user.GetProfileRequest) (*user.GetProfileResponse, error) {
	return nil, nil
}

func (g *grpcApp) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	log := logger.FromContext(ctx)
	birthDate, err := time.Parse(time.RFC3339, in.GetBirthDate())
	if err != nil {
		log.Error("Parse birthdate failed", zap.Error(err))
		return nil, err
	}
	request := &dto.CreateUserRequest{
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
		Password:    in.GetPassword(),
		FirstName:   in.GetFirstName(),
		LastName:    in.GetLastName(),
		Avatar:      in.GetAvatar(),
		BirthDate:   birthDate,
	}

	if err := request.Validate(); err != nil {
		log.Error("Validate request failed", zap.Error(err))
		return nil, err
	}

	response, err := command_bus.Dispatch[*command.CreateUserCommand, *command.CreateUserCommandResult](
		g.commandBus, ctx, &command.CreateUserCommand{
			Email:       request.Email,
			Password:    request.Password,
			FirstName:   request.FirstName,
			LastName:    request.LastName,
			PhoneNumber: request.PhoneNumber,
			Avatar:      request.Avatar,
			BirthDate:   request.BirthDate,
		})

	if err != nil {
		log.Error("Failed to create user", zap.Error(err))
		return nil, err
	}

	return response.ToDto().ToPb(), nil
}
