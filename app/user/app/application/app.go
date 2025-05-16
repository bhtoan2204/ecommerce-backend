package application

import (
	"context"
	"fmt"
	"user/app/domain/services"
	"user/package/contxt"
	"user/package/logger"
	"user/package/mgrpc"
	"user/package/settings"
	"user/package/tracer"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ App = (*app)(nil)

type App interface {
	Start(ctx context.Context) error
}

type app struct {
	cfg *settings.Config
	// env *env.Env

	// repos repos.Repo
	srvs services.Service
}

func NewApp(cfg *settings.Config) (App, error) {
	services, err := services.NewService(cfg)

	if err != nil {
		fmt.Println("error occurs", err)
		return nil, err
	}
	return &app{
		cfg:  cfg,
		srvs: services,
	}, nil
}

func (a *app) Start(ctx context.Context) error {
	// GRPC
	panicHandler := func(p any) (err error) {
		logger.DefaultLogger().Error("recovered from panic", zap.Any("panic", p))
		return status.Errorf(codes.Internal, "%s", p)
	}
	var sopts []grpc.ServerOption
	sopts = append(sopts,
		grpc.ChainUnaryInterceptor(
			contxt.SetupContext(),
			mgrpc.SetCommonData(),
			mgrpc.PopulateRequestID(),
			mgrpc.SetLogger(),
			// mgrpc.MonitorRequestDuration(a.env.XMetric().GRPCReqDuration()),
			mgrpc.Recovery(panicHandler),
			mgrpc.Timeout(),
			mgrpc.HandleError(),
		),
		grpc.StatsHandler(tracer.GrpcServerHandler()),
	)
	return nil
}
