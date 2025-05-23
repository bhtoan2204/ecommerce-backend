package presentation

import (
	"context"
	"user/app/domain/services"
	"user/app/infrastructure/persistent/postgresql"
	"user/app/infrastructure/persistent/postgresql/repository"
	"user/package/contxt"
	"user/package/ghealth"
	"user/package/logger"
	"user/package/mgrpc"
	"user/package/server"
	"user/package/settings"
	"user/package/tracer"
	"user/proto/user"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appGrpc "user/app/presentation/grpc"

	pbHealth "google.golang.org/grpc/health/grpc_health_v1"
)

var _ App = (*app)(nil)

type App interface {
	Start(ctx context.Context) error
}

type app struct {
	cfg  *settings.Config
	srvs services.Service
}

func NewApp(ctx context.Context, cfg *settings.Config) (App, error) {
	postgresDB := postgresql.NewPostgresDB(ctx, cfg)
	postgresRepository := repository.NewRepository(postgresDB.GetDB())
	services, err := services.NewService(cfg, &postgresRepository)

	if err != nil {
		return nil, err
	}
	return &app{
		cfg:  cfg,
		srvs: services,
	}, nil
}

func (a *app) Start(ctx context.Context) error {
	// HTTP
	log := logger.FromContext(ctx)

	// app, err := appHttp.New(a.cfg, a.ucs)
	// if err != nil {
	// 	return fmt.Errorf("new application failed err=%w", err)
	// }

	// srv, err := server.New()
	// if err != nil {
	// 	return err
	// }
	// log.Info("HTTP Server running on PORT: %s", zap.String("port", srv.Port()))

	// go func() {
	// 	err = srv.ServeHTTPHandler(ctx, app.Routes(ctx))
	// 	if err != nil {
	// 		log.Warn("Serve HTTP Handler failed err=%w", zap.Error(err))
	// 	}
	// }()

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

	rpcServer := grpc.NewServer(sopts...)
	healthCheck := ghealth.NewHealthService()
	pbHealth.RegisterHealthServer(rpcServer, healthCheck)

	gApp, err := appGrpc.NewGrpcApp(a.srvs)
	if err != nil {
		return err
	}
	user.RegisterUserServiceServer(rpcServer, gApp)

	// TODO: service discovery in future
	grpcSrv, err := server.New(a.cfg.Server.GRPCPort)
	if err != nil {
		log.Error("Error occurs", zap.Error(err))
		return err
	}
	log.Info("Starting GRPC Server", zap.String("ip", grpcSrv.IP()), zap.String("port", grpcSrv.Port()))

	return grpcSrv.ServeGRPC(ctx, rpcServer)
}
