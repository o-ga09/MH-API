package handler

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	{
		systemHandler := NewSystemHandler()
		v1.GET("/system/health",systemHandler.Health)
	}

	return r
}