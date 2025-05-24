package grpc

import (
	command_bus "payment/app/application/commands"
	"payment/app/domain/usecases"
	"payment/proto/payment"
)

var _ payment.PaymentServiceServer = (*grpcApp)(nil)

type GrpcApp interface {
	payment.PaymentServiceServer
}

type grpcApp struct {
	commandBus *command_bus.CommandBus
	ucs        usecases.Usecase
}

func NewGrpcApp(ucs usecases.Usecase) (GrpcApp, error) {
	commandBus := command_bus.NewCommandBus()

	return &grpcApp{
		commandBus: commandBus,
		ucs:        ucs,
	}, nil
}
