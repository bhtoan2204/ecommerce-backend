package grpc

import "user/proto/user"

var _ user.UserServiceClient = (*grpcApp)(nil)

type GrpcApp interface {
	user.UserServiceClient
}

type grpcApp struct {
}

func NewGrpcApp() (GrpcApp, error) {
	return &grpcApp{}, nil
}
