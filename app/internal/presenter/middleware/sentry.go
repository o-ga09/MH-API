package middleware

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func SentryTracingMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		hub := sentry.GetHubFromContext(ctx)
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
			ctx = sentry.SetHubOnContext(ctx, hub)
		}
		options := []sentry.SpanOption{
			sentry.WithOpName("http.server"),
			sentry.ContinueFromRequest(c.Request),
			sentry.WithTransactionSource(sentry.SourceURL),
		}
		transaction := sentry.StartTransaction(ctx,
			fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path),
			options...,
		)
		defer transaction.Finish()
		c.Request = c.Request.WithContext(transaction.Context())

		// 次のハンドラを実行
		next(c)
	}
}
