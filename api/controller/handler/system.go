package handler

import (
	"mh-api/api/controller/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const cookieMaxAge = 3600

type SystemHandler struct {}

type AuthHandler struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (s * SystemHandler) Health(c *gin.Context) {
	c.JSON(200,gin.H{
		"Message": "ok",
	})
}

func (h *AuthHandler) SignUpHandler(ctx *gin.Context) {
	err := ctx.ShouldBind(h)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	if h.Name != user || h.Password != password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	token, err := middleware.GenerateToken(h.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to sign up",
		})
		return
	}

    // Cookieにトークンをセット
	ctx.SetCookie("token", token, cookieMaxAge, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"user_id": h.Name,
		"message": "Successfully",
	})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}