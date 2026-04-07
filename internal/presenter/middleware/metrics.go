package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const meterName = "mh-api"

// MetricsMiddleware は HTTP リクエストの RED メトリクス（Rate/Error/Duration）を計測する Gin ミドルウェアです。
// 計測するメトリクス:
//   - http.server.request.duration: リクエストのレイテンシ（ヒストグラム, 秒単位）
//   - http.server.request.count:    リクエスト数（カウンター）
//
// 各メトリクスには以下の属性が付与されます:
//   - http.route:        リクエストにマッチした Gin ルートパターン（例: /v1/monsters/:id）
//   - http.method:       HTTP メソッド（GET, POST, ...）
//   - http.status_code:  レスポンスステータスコード
func MetricsMiddleware() gin.HandlerFunc {
	meter := otel.Meter(meterName)

	requestDuration, _ := meter.Float64Histogram(
		"http.server.request.duration",
		metric.WithDescription("HTTP server request duration in seconds"),
		metric.WithUnit("s"),
	)
	requestCount, _ := meter.Int64Counter(
		"http.server.request.count",
		metric.WithDescription("Total number of HTTP server requests"),
	)

	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		route := c.FullPath()
		if route == "" {
			route = "unknown"
		}

		attrs := []attribute.KeyValue{
			attribute.String("http.route", route),
			attribute.String("http.method", c.Request.Method),
			attribute.String("http.status_code", strconv.Itoa(c.Writer.Status())),
		}
		opt := metric.WithAttributes(attrs...)

		elapsed := time.Since(start).Seconds()
		requestDuration.Record(c.Request.Context(), elapsed, opt)
		requestCount.Add(c.Request.Context(), 1, opt)
	}
}
