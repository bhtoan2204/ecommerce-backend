package tracer

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc/stats"
)

func GrpcStatsHandler(opts ...otelgrpc.Option) stats.Handler {
	return otelgrpc.NewClientHandler(opts...)
}
