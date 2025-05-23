package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"user/app/presentation"
	"user/config"
	"user/package/logger"
	"user/package/settings"

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
	start(ctx, &config)
}

func start(ctx context.Context, config *settings.Config) error {
	log := logger.FromContext(ctx)
	app, err := presentation.NewApp(ctx, config)
	if err != nil {
		log.Error("NewApp failed", zap.Error(err))
		return fmt.Errorf("new app got err=%w", err)
	}

	if app == nil {
		log.Error("NewApp returned nil app without error")
		return fmt.Errorf("NewApp returned nil app without error")
	}
	log.Info("Starting application")
	return app.Start(ctx)
}
