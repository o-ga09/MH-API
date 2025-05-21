package middleware

import (
	"context"
	"log/slog" // Add slog for logging
	"mh-api/app/internal/driver/mysql" // Add import for mysql package
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// dbKey is the type for the database context key.
type dbKey string

// dbCtxKey is the key for storing the GORM DB instance in the context.
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

// GetDB retrieves the GORM DB instance from the context.
// It's placed here as it's related to context and DB,
// and will be used by the new WithDB middleware and services.
func GetDB(ctx context.Context) *gorm.DB {
	value := ctx.Value(dbCtxKey)
	if db, ok := value.(*gorm.DB); ok {
		return db
	}
	return nil
}

// WithDB creates a new DB session and stores it in the context.
func WithDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtain a DB connection using the existing New function
		// We use c.Request.Context() so that if it's already timed out or cancelled,
		// the DB connection attempt can respect that.
		db := mysql.New(c.Request.Context())
		if db == nil {
			// If DB connection fails, we might want to return an error response
			// For now, log it and proceed. Depending on application requirements,
			// this could be c.AbortWithError(http.StatusInternalServerError, errors.New("failed to connect to database"))
			slog.Log(c.Request.Context(), SeverityError, "failed to connect to database for request")
			c.Next() // Or c.Abort() if requests cannot proceed without DB
			return
		}

		// Store the DB session in the context
		ctxWithDB := context.WithValue(c.Request.Context(), dbCtxKey, db)
		c.Request = c.Request.WithContext(ctxWithDB)

		c.Next()

		// Note: GORM typically manages a connection pool, so explicit closing of the 'db' instance
		// per request might not be necessary, as it's drawn from and returned to the pool.
		// If specific session/transaction management is added later, this would be the place
		// to handle rollback/commit and closing.
	}
}
