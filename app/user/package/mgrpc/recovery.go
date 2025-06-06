package mgrpc

import (
	"context"
	"fmt"
	"runtime"

	"google.golang.org/grpc"
)

type RecoveryHandlerFunc func(p any) (err error)
type RecoveryHandlerFuncContext func(ctx context.Context, p any) (err error)

func Recovery(f RecoveryHandlerFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		defer func() {
			f := RecoveryHandlerFuncContext(func(ctx context.Context, p any) error {
				return f(p)
			})
			if r := recover(); r != nil {
				err = recoverFrom(ctx, r, f)
			}
		}()

		return handler(ctx, req)
	}
}

func recoverFrom(ctx context.Context, p any, r RecoveryHandlerFuncContext) error {
	if r != nil {
		return r(ctx, p)
	}
	stack := make([]byte, 64<<10)
	stack = stack[:runtime.Stack(stack, false)]
	return &PanicError{Panic: p, Stack: stack}
}

type PanicError struct {
	Panic any
	Stack []byte
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("panic caught: %v\n\n%s", e.Panic, e.Stack)
}
