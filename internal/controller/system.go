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

// Health godoc
// @Summary ヘルスチェック
// @Description システムが稼働しているか確認する
// @Tags System
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (s *SystemHandler) Health(c *gin.Context) {
	c.JSON(200, gin.H{"Message": "ok"})
}

// DBHealth godoc
// @Summary DBヘルスチェック
// @Description データベースへの接続を確認する
// @Tags System
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /health/db [get]
func (s *SystemHandler) DBHealth(c *gin.Context) {
	ctx := c.Request.Context()
	err := s.healthRepo.GetStatus(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "db ok"})
}
