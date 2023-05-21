package api

import (
	"github.com/MH-API/mh-api/api/handler"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	{
		systemHandler := handler.NewSystemHandler()
		v1.GET("/system/health",systemHandler.Health)
	}

	return r
}