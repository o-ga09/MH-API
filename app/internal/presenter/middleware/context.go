package middleware

import (
	"context"
	"mh-api/app/pkg"

	"time"

	"github.com/gin-gonic/gin"
)

type RequestId string

func AddID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), RequestId("requestId"), pkg.GenerateID())
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func WithTimeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetRequestID(ctx context.Context) string {
	return ctx.Value(RequestId("requestId")).(string)
}
