package main

import (
	"context"
	"gateway/package/config"
	"gateway/package/logger"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

//	@title			OMS API Gateway
//	@version		1.0
//	@description	REST -> GRPC API Gateway

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

// @host						localhost:8080
// @BasePath					/
// @query.collection.format	multi
func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	config, _ := config.InitLoadConfig()
	log := logger.NewLogger(config.LogConfig)
	ctx = logger.WithLogger(ctx, log)

	defer func() {
		done()
		if r := recover(); r != nil {
			log.Info("Application went wrong. Panic err:", zap.Error(r.(error)))
		}
	}()
	err := Initialize(ctx, config)
	done()
	if err != nil {
		log.Error("Initialize failed", zap.Error(err))
		return
	}

	log.Info("App shutdown successful")
}
