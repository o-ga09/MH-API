package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		latency := time.Now()
		slog.Log(ctx, SeverityInfo, "処理開始", "method", c.Request.Method, "path", c.Request.URL.Path)
		// request body
		reqBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			slog.Log(ctx, SeverityError, "failed to read request body", "error", err)
		}
		c.Next()
		go func() {
			if len(reqBody) == 0 {
				slog.Log(ctx, SeverityInfo, "処理終了", "method", c.Request.Method, "path", c.Request.URL.Path, "処理時間", fmt.Sprintf("%v", time.Since(latency)))
				return
			}
			buf := bytes.NewBuffer(nil)
			err := json.Compact(buf, reqBody)
			if err != nil {
				slog.Log(ctx, SeverityInfo, "処理終了", "method", c.Request.Method, "path", c.Request.URL.Path, "処理時間", fmt.Sprintf("%v", time.Since(latency)))
				return
			}
			slog.Log(ctx, SeverityInfo, buf.String(), "method", c.Request.Method, "path", c.Request.URL.Path, "処理時間", fmt.Sprintf("%v", time.Since(latency)))
		}()
	}
}
