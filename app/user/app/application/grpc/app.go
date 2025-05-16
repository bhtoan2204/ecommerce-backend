package grpc

import "user/proto/user"

var _ user.UserServiceServer = (*grpcApp)(nil)

type GrpcApp interface {
	user.UserServiceServer
}

type grpcApp struct {
}

func NewGrpcApp() (GrpcApp, error) {
	return &grpcApp{}, nil
}
