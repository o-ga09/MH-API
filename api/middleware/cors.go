package middleware

import (
	"mh-api/api/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	cfg, _ := config.New()
	return cors.New(cors.Config{
		AllowOrigins: []string{
			cfg.ALLOW_URL,
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	})
}
