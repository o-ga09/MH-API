package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestInfo struct {
	status                                            int
	contents_length                                   int64
	method, path, sourceIP, query, user_agent, errors string
	elapsed                                           time.Duration
}

func RequestLogger(l *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		slog.Log(c, SeverityInfo, "処理開始", "request Id", GetRequestID(c.Request.Context()))
		c.Next()

		r := &RequestInfo{
			status:          c.Writer.Status(),
			contents_length: c.Request.ContentLength,
			path:            c.Request.URL.Path,
			sourceIP:        c.ClientIP(),
			query:           c.Request.URL.RawQuery,
			user_agent:      c.Request.UserAgent(),
			errors:          c.Errors.ByType(gin.ErrorTypePrivate).String(),
			elapsed:         time.Since(start),
		}
		slog.Log(c, SeverityInfo, "処理終了", "Request", r.LogValue(), "requestId", GetRequestID(c.Request.Context()))
	}
}

func (r *RequestInfo) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("status", r.status),
		slog.Int64("Content-length", r.contents_length),
		slog.String("method", r.method),
		slog.String("path", r.path),
		slog.String("sourceIP", r.sourceIP),
		slog.String("query", r.query),
		slog.String("user_agent", r.user_agent),
		slog.String("errors", r.errors),
		slog.String("elapsed", r.elapsed.String()),
	)
}
