package controller

import (
	"github.com/gin-gonic/gin"
)

type SystemHandler struct{}

func (s *SystemHandler) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "ok",
	})
}

func (s *SystemHandler) DBHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "db ok",
	})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}
