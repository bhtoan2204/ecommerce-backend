package grpc

import (
	"context"
	"user/proto/user"

	"google.golang.org/grpc"
)

func (g *grpcApp) Login(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*user.LoginResponse, error) {
	// Implement the login logic here
	// For example, validate the user credentials and generate tokens
	// Return the LoginResponse with the generated tokens
	return &user.LoginResponse{
		AccessToken:           "access_token",
		RefreshToken:          "refresh_token",
		AccessTokenExpiresIn:  3600,
		RefreshTokenExpiresIn: 7200,
	}, nil
}
