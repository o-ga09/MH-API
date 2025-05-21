package middleware

import (
	"context"
	"log/slog"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/pkg"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type dbKey string

const dbCtxKey dbKey = "db"

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

func GetDB(ctx context.Context) *gorm.DB {
	value := ctx.Value(dbCtxKey)
	if db, ok := value.(*gorm.DB); ok {
		return db
	}
	return nil
}

func WithDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := mysql.New(c.Request.Context())
		if db == nil {
			slog.Log(c.Request.Context(), pkg.SeverityError, "failed to connect to database for request")
			c.Next()
			return
		}

		ctxWithDB := context.WithValue(c.Request.Context(), dbCtxKey, db)
		c.Request = c.Request.WithContext(ctxWithDB)

		c.Next()
	}
}
