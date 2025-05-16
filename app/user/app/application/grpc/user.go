package grpc

import (
	"context"
	"user/proto/user"
)

func (g *grpcApp) GetProfile(context.Context, *user.GetProfileRequest) (*user.GetProfileResponse, error) {
	return nil, nil
}
