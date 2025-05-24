package contxt

import (
	"context"
	"fmt"
	"sync"

	"payment/package/utils"

	"google.golang.org/grpc"
)

type ctxKey int

const (
	CtxKey ctxKey = iota
)

type Context struct {
	parrentCtx context.Context
	mu         sync.RWMutex

	// Keys is a key/value pair exclusively for the context of each request.
	keys map[string]any
}

func SetupContext() grpc.UnaryServerInterceptor {
	return func(c context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx := context.WithValue(c, CtxKey, &Context{
			parrentCtx: c,
		})
		return handler(ctx, req)
	}
}

func GetWrapper(ctx context.Context) (*Context, error) {
	value := ctx.Value(CtxKey)
	if value == nil {
		return nil, fmt.Errorf("could not get contxt.Context from context")
	}

	wrapperCtx, ok := value.(*Context)
	if !ok {
		return nil, fmt.Errorf("could not get contxt.Context from context")
	}

	return wrapperCtx, nil
}

func ContextWithWrapper(ctx context.Context, wrapperCtx *Context) context.Context {
	// nolint
	return context.WithValue(ctx, CtxKey, wrapperCtx)
}

func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.keys == nil {
		c.keys = make(map[string]any)
	}

	c.keys[key] = value
}

func (c *Context) Get(key string) (value any, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists = c.keys[key]
	return
}

func (c *Context) GetInt64(key string) (int64, error) {
	if val, ok := c.Get(key); ok && val != nil {
		return utils.GetInt(val)
	}
	return 0, fmt.Errorf("key: %s not found", key)
}

func (c *Context) GetString(key string) (string, error) {
	if val, ok := c.Get(key); ok && val != nil {
		s, ok := val.(string)
		if !ok {
			return "", fmt.Errorf("parse value: %v to string error", val)
		}
		return s, nil
	}

	return "", fmt.Errorf("key: %s not found", key)
}

func GetAcceptLanguage(ctx context.Context) (string, error) {
	wrapperCtx, err := GetWrapper(ctx)
	if err != nil {
		return "", fmt.Errorf("get wrapper context err=%w", err)
	}

	return wrapperCtx.GetString("accept-language")
}

func RequestIDFromCtx(ctx context.Context) string {
	v := ctx.Value(CtxKey)
	if v == nil {
		return ""
	}

	if val, ok := v.(string); ok {
		return val
	}

	return ""
}
