package mgrpc

import (
	"context"
	"fmt"
	"payment/package/xerror"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		response, err := handler(ctx, req)
		if err != nil {
			// requestID := RequestIDFromCtx(ctx)
			input := fmt.Sprint(req)
			if input == "" {
				input = "no input"
			}
			appErr, ok := err.(*xerror.XError)
			if !ok {
				return response, err
			}

			// handle error
			st := status.New(codes.Code(appErr.GrpcCode), err.Error())
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

			return nil, st.Err()
		}

		return response, nil
	}
}
