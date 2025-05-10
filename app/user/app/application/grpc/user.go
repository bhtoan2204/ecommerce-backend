package grpc

import (
	"context"
	"user/proto/user"

	"google.golang.org/grpc"
)

func (g *grpcApp) GetProfile(ctx context.Context, in *user.GetProfileRequest, opts ...grpc.CallOption) (*user.GetProfileResponse, error) {
	return nil, nil
}
