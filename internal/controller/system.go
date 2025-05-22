package controller

import (
	"mh-api/internal/service/health"
	"net/http"

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
	ctx := c.Request.Context()

	err := s.service.GetStatus(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "db ok",
	})
}
