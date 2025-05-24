package mgrpc

import (
	"context"
	"payment/package/contxt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ctxKey string

const ctxKeyRequestID = ctxKey("request_id")

func PopulateRequestID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		reqID := RequestIDFromCtx(ctx)
		if reqID == "" {
			uid, err := uuid.NewRandom()
			if err == nil {
				reqID = uid.String()
			}
		}

		ctx = withRequestID(ctx, reqID)

		return handler(ctx, req)
	}
}

func withRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxKeyRequestID, id)
}

func RequestIDFromCtx(ctx context.Context) string {
	wapprer, err := contxt.GetWrapper(ctx)
	if err == nil {
		reqID, err := wapprer.GetString("x-request-id")
		if err == nil {
			return reqID
		}
	}

	v := ctx.Value(ctxKeyRequestID)
	if v == nil {
		return ""
	}

	if val, ok := v.(string); ok {
		return val
	}

	return ""
}
