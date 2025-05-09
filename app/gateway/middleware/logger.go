package middleware

import (
	"gateway/package/contxt"
	"gateway/package/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// grab whatever logger was already on the context
		log := logger.FromContext(ctx)

		// if there's a requestID in the context, clone+attach it
		if reqID := contxt.RequestIDFromCtx(ctx); reqID != "" {
			log = log.WithFields(
				zap.String("request_id", reqID),
			)
		}

		// shove the new logger back into the context
		ctx = logger.WithLogger(ctx, log)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
