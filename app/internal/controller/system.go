package controller

import (
	"log/slog" // Added
	"mh-api/app/internal/presenter/middleware" // Added
	"mh-api/app/internal/service/health"
	"net/http" // Added

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
	// ADDED: Retrieve DB
	db := middleware.GetDB(c.Request.Context())
	if db == nil {
		slog.Log(c, middleware.SeverityError, "database connection not found in context")
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "INTERNAL SERVER ERROR"})
		return
	}
	// END ADDED

	err := s.service.GetStatus(db) // MODIFIED: pass db
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ // Changed to 500 to be consistent
			"Message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{ // Changed to 200 to be consistent
		"Message": "db ok",
	})
}
