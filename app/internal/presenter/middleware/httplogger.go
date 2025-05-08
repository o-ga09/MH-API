package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		slog.Log(ctx, SeverityInfo, "処理開始", "method", c.Request.Method, "path", c.Request.URL.Path)
		// request body
		reqBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			slog.Log(ctx, SeverityError, "failed to read request body", "error", err)
		}
		c.Next()
		go func() {
			if len(reqBody) == 0 {
				slog.Log(ctx, SeverityInfo, "処理終了", "method", c.Request.Method, "path", c.Request.URL.Path)
				return
			}
			buf := bytes.NewBuffer(nil)
			err := json.Compact(buf, reqBody)
			if err != nil {
				slog.Log(ctx, SeverityError, "failed to compact request body", "error", err)
			}
			slog.Log(ctx, SeverityInfo, buf.String(), "method", c.Request.Method, "path", c.Request.URL.Path)
		}()
	}
}
