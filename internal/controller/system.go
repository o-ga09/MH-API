package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthRepository interface {
	GetStatus(ctx context.Context) error
}

type SystemHandler struct {
	healthRepo HealthRepository
}

func NewHealthService(repo HealthRepository) SystemHandler {
	return SystemHandler{healthRepo: repo}
}

func (s *SystemHandler) Health(c *gin.Context) {
	c.JSON(200, gin.H{"Message": "ok"})
}

func (s *SystemHandler) DBHealth(c *gin.Context) {
	ctx := c.Request.Context()
	err := s.healthRepo.GetStatus(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "db ok"})
}
