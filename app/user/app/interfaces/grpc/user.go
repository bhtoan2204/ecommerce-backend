package grpc

import (
	"context"
	"time"
	"user/app/domain/dto"
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

	response, err := g.commandBus.Dispatch(ctx, request)
	if err != nil {
		log.Error("Failed to create user", zap.Error(err))
		return nil, err
	}

	return response.(*dto.CreateUserResponse).ToPb(), nil
}
