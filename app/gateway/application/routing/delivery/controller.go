package delivery

import (
	"gateway/package/settings"
	"net/http"
	"strconv"

	"oms-gateway/pkg/grpc"
	"oms-gateway/pkg/logging"
	"oms-gateway/proto/oms"
	"oms-gateway/proto/wty"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"oms-gateway/application/model"
	"oms-gateway/application/routing"
	"oms-gateway/constant"
	pkgRedis "oms-gateway/pkg/infras/redis"
	jwtPkg "oms-gateway/pkg/jwt"
	"oms-gateway/pkg/wrapper"

	"github.com/golang-jwt/jwt"

	"github.com/redis/go-redis/v9"
)

type RoutingHandler struct {
	config        *settings.Config
	routingUC     routing.RoutingUseCase
	redisClient   *redis.Client
	registry      map[string]routingConfig
	verifyAuth    jwtPkg.Verifier
	verifyInside  jwtPkg.Verifier
	internalRoute map[string]struct{}
}

type AuthClaims struct {
	jwt.StandardClaims
	UserID int64  `json:"Sub"`
	Email  string `json:"Email"`
}

type InsideClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"sub"`
}

type ResponseError struct {
	HttpCode  int    `json:"http_code"`
	GrpcCode  int    `json:"grpc_code"`
	Message   string `json:"message"`
	RootError string `json:"root_error"`
}

type AppError struct {
	Errors []ResponseError `json:"errors"`
}

func NewRoutingHandler(cfg *settings.Config, routingUC routing.RoutingUseCase) *RoutingHandler {
	redisClient, err := pkgRedis.New(cfg.Redis)
	if err != nil {
		panic(err)
	}

	verifyAuth, err := jwtPkg.NewVerifier(
		jwtPkg.WithPublicKeyFile(cfg.JwtKey.AuthPublicKey),
	)
	if err != nil {
		panic(err)
	}

	verifyInside, err := jwtPkg.NewVerifier(
		jwtPkg.WithSecretKey(cfg.JwtKey.InsideSecret),
	)
	if err != nil {
		panic(err)
	}

	return &RoutingHandler{
		config:       cfg,
		routingUC:    routingUC,
		redisClient:  redisClient,
		registry:     buildRegistry(cfg),
		verifyAuth:   verifyAuth,
		verifyInside: verifyInside,
		internalRoute: map[string]struct{}{
			"get:/api/v1/oms/shipment-store-summary": struct{}{},
		},
	}
}

func (h *RoutingHandler) handle() gin.HandlerFunc {
	return wrapper.WithContext(func(ctx *wrapper.Context) {
		log := logging.DefaultLogger()

		route := ctx.Request.Method + ":" + ctx.FullPath()
		routingCfg, found := h.registry[route]
		if !found {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error has occurred. Please retry your request later"})
			return
		}

		uid := ctx.GetInt64(constant.KeyUserLoginID)
		insideUid := ctx.GetInt64(constant.KeyInsideUserID)
		metadata := map[string]string{
			"ip-address":      ctx.ClientIP(),
			"accept-language": ctx.GetHeader("Accept-Language"),
			"token":           ctx.GetString(constant.KeyAuthToken),
			"uid":             strconv.FormatInt(uid, 10),
			"inside-token":    ctx.GetHeader("X-Inside-Token"),
			"inside-uid":      strconv.FormatInt(insideUid, 10),
			"user-email":      ctx.GetString(constant.KeyUserLoginEmail),
		}

		data, err := routingCfg.handler.Handle(ctx)
		if err != nil {
			appErr := &AppError{
				Errors: []ResponseError{{
					HttpCode:  http.StatusBadRequest,
					GrpcCode:  int(codes.InvalidArgument),
					Message:   err.Error(),
					RootError: err.Error(),
				}},
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, appErr)
			return
		}
		if len(routingCfg.remoteServiceName) == 0 {
			ctx.JSON(http.StatusOK, data)
			return
		}

		routing := &model.RoutingData{
			ServiceName:   routingCfg.remoteServiceName,
			ServiceMethod: routingCfg.remoteServiceMethod,
			Payload:       data,
			Metadata:      metadata,
		}

		res, err := h.routingUC.Forward(routing)
		if err != nil {
			appErr := &AppError{}
			log.Errorf("Forward request failed: err", err, ", host:", h.config.Service.OmsServiceUrl)
			errStatus, _ := status.FromError(err)
			for _, detail := range errStatus.Details() {
				switch info := detail.(type) {
				case *wty.ResponseError:
					appErr.Errors = append(appErr.Errors, ResponseError{
						HttpCode:  int(info.Httpcode),
						GrpcCode:  int(info.Grpccode),
						Message:   info.Message,
						RootError: info.RootError,
					})
				case *oms.ResponseError:
					appErr.Errors = append(appErr.Errors, ResponseError{
						HttpCode:  int(info.Httpcode),
						GrpcCode:  int(info.Grpccode),
						Message:   info.Message,
						RootError: info.RootError,
					})
				default:
					appErr.Errors = append(appErr.Errors, ResponseError{
						GrpcCode:  int(errStatus.Code()),
						Message:   errStatus.String(),
						RootError: "",
					})
				}
			}

			if len(appErr.Errors) == 0 {
				appErr.Errors = append(appErr.Errors, ResponseError{
					GrpcCode:  int(errStatus.Code()),
					Message:   errStatus.String(),
					RootError: "",
				})
			}
			code := grpc.MapGRPCErrCodeToHttpStatus(errStatus.Code())
			ctx.AbortWithStatusJSON(code, appErr)
			return
		}

		ctx.JSON(http.StatusOK, res)
	})
}

func (h *RoutingHandler) Authorization() gin.HandlerFunc {
	return wrapper.WithContext(func(ctx *wrapper.Context) {
		return
	})
}
