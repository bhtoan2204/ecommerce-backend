package application

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"gateway/middleware"
	"gateway/package/logger"
	"gateway/package/server"
	"gateway/package/tracer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"honnef.co/go/tools/config"
)

var _ App = (*Server)(nil)

type App interface {
	Start(ctx context.Context) error
}

type Server struct {
	cfg        *config.Config
	router     *gin.Engine
	httpServer *http.Server
	// handler    *delivery.RoutingHandler
}

func New(cfg *config.Config) (App, error) {
	return &Server{
		cfg: cfg,
	}, nil
}

func (s *Server) Routes(ctx context.Context) http.Handler {
	r := gin.New()
	r.MaxMultipartMemory = 50 << 20
	r.RedirectTrailingSlash = false

	r.Use(middleware.ErrorHandler())
	r.Use(middleware.SetRequestID())
	r.Use(middleware.SetLogger())
	r.Use(tracer.GinMiddleware("gateway"))
	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		fmt.Println("panic: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"errors": gin.H{"error": "something went wrong"}}) // not return detail error to client when panic
	}))

	if os.Getenv("ENVIRONMENT") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"X-Inside-Token",
	}
	r.Use(cors.New(corsConfig))

	// // health-check
	pingHandler := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"version":  "1.0.0",
				"clientIP": ctx.ClientIP(),
			},
		})
	}
	r.GET("/health-check", pingHandler)
	r.HEAD("/health-check", pingHandler)

	// // swagger
	// if os.Getenv("ENVIRONMENT") != "prod" {
	// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// }

	// s.router = r
	// s.handler = delivery.NewRoutingHandler(
	// 	s.cfg,
	// 	usecase.NewRoutingUseCase(service.NewServiceClient(s.cfg)))

	return r
}

func (s *Server) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)

	srv, err := server.New(8080)
	if err != nil {
		return err
	}

	log.Info("HTTP Server running on PORT", zap.Int("port", 8080))

	return srv.ServeHTTPHandler(ctx, s.Routes(ctx))
}
