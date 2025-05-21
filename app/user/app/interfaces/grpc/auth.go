package grpc

import (
	"context"
	"user/app/domain/dto"
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

	response, err := g.commandBus.Dispatch(ctx, request)
	if err != nil {
		log.Error("Failed to Login", zap.Error(err))
		return nil, err
	}

	return response.(*dto.LoginResponse).ToPb(), nil
}
