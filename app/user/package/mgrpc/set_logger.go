package mgrpc

import (
	"context"
	"user/package/contxt"
	"user/package/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func SetLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log := logger.FromContext(ctx)
		if reqID := contxt.RequestIDFromCtx(ctx); reqID != "" {
			log = log.WithFields(
				zap.String("request_id", reqID),
			)
		}

		ctx = logger.WithLogger(ctx, log)

		return handler(ctx, req)
	}
}
