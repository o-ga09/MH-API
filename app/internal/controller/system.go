package controller

import (
	"mh-api/app/internal/service/health"

	"github.com/gin-gonic/gin"
)

type SystemHandler struct {
	service health.HealthService
}

func NewHealthService(service health.HealthService) SystemHandler {
	return SystemHandler{
		service: service,
	}
}

func (s *SystemHandler) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "ok",
	})
}

func (s *SystemHandler) DBHealth(c *gin.Context) {
	err := s.service.GetStatus()
	if err != nil {
		c.JSON(500, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "db ok",
	})
}
