package middleware

import (
	"context"
	"mh-api/app/internal/driver/mysql"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestId string

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

func WithDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = mysql.New(ctx)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
