package presentation

import (
	"context"
	"user/app/domain/usecases"
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
	appHttp "user/app/presentation/http"

	pbHealth "google.golang.org/grpc/health/grpc_health_v1"
)

var _ App = (*app)(nil)

type App interface {
	Start(ctx context.Context) error
}

type app struct {
	cfg *settings.Config
	ucs usecases.Usecase
}

func NewApp(ctx context.Context, cfg *settings.Config) (App, error) {
	postgresDB := postgresql.NewPostgresDB(ctx, cfg)
	postgresRepository := repository.NewRepository(postgresDB.GetDB())
	ucs, err := usecases.NewUsecase(cfg, &postgresRepository)

	if err != nil {
		return nil, err
	}
	return &app{
		cfg: cfg,
		ucs: ucs,
	}, nil
}

func (a *app) Start(ctx context.Context) error {
	// Metrics
	log := logger.FromContext(ctx)

	go func() {
		srvMetric, err := server.New(a.cfg.Server.MetricPort)
		if err != nil {
			log.Warn("New server metric", zap.Error(err))
			return
		}

		log.Info("Metric running on PORT", zap.String("port", srvMetric.Port()))

		metric, err := appHttp.NewMetric()
		if err != nil {
			log.Warn("New app metric", zap.Error(err))
			return
		}
		if err := srvMetric.ServeHTTPHandler(ctx, metric.Handler()); err != nil {
			log.Warn("Serve metric handler", zap.Error(err))
		}
	}()

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

	gApp, err := appGrpc.NewGrpcApp(a.ucs)
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
