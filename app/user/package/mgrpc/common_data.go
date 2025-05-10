package mgrpc

import (
	"context"
	"user/package/contxt"
	"user/package/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func SetCommonData() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log := logger.FromContext(ctx)

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			log.Warn("Get metadata from request got error")
			return handler(ctx, req)
		}

		ctxWrapper, err := contxt.GetWrapper(ctx)
		if err != nil {
			log.Warn("Get context wrapper failed", zap.Error(err))
			return handler(ctx, req)
		}

		if val := md.Get("accept-language"); len(val) > 0 {
			ctxWrapper.Set("accept-language", val[0])
		}

		if val := md.Get("ip-address"); len(val) > 0 {
			ctxWrapper.Set("ip-address", val[0])
		}

		if val := md.Get("uid"); len(val) > 0 {
			ctxWrapper.Set("uid", val[0])
		}

		if val := md.Get("token"); len(val) > 0 {
			ctxWrapper.Set("token", val[0])
		}

		if val := md.Get("x-request-id"); len(val) > 0 {
			ctxWrapper.Set("x-request-id", val[0])
		}

		if val := md.Get("user-email"); len(val) > 0 {
			ctxWrapper.Set("user-email", val[0])
		}

		ctx = contxt.ContextWithWrapper(ctx, ctxWrapper)

		return handler(ctx, req)
	}
}
