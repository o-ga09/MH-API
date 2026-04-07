package middleware

import (
	"net/http"

	"mh-api/internal/database/mysql"
)

// DBSessionMiddleware ensures a *gorm.DB is available in the request context.
// It initializes the DB (once) via mysql.New and injects the resulting
// context into the request so repository/query service code can call
// mysql.CtxFromDB(ctx).
func DBSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := mysql.New(r.Context())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
