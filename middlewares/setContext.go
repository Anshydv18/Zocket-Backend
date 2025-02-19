package middlewares

import (
	"backend/constants"
	"backend/utils"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func SetContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
		defer cancel()

		requestId := c.GetHeader("X-Request-ID")
		if requestId == "" {
			requestId = utils.GenerateRandomRequestId()
		}

		ctx = context.WithValue(ctx, constants.REQUESTIDKEY, requestId)
		c.Set("context", ctx)
		c.Next()
	}
}
