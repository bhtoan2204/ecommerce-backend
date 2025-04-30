package main

import (
	"context"
	"os/signal"
	"syscall"
	"user/config"
	"user/package/logger"

	"go.uber.org/zap"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	config, _ := config.InitLoadConfig()
	logger := logger.NewLogger(config.LogConfig)
	defer func() {
		done()
		if r := recover(); r != nil {
			logger.Error("application went wrong. Panic err=%v", zap.Error(r.(error)))
		}
	}()
	start(ctx)
}

func start(ctx context.Context) error {
	// Initialize your application here
	// For example, you can set up a database connection, start a web server, etc.
	return nil
}
