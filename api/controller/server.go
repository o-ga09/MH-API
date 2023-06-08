package controller

import (
	di "mh-api/api/DI"
	"mh-api/api/controller/handler"
	"mh-api/api/controller/middleware"

	"github.com/gin-gonic/gin"
)

func NewServer() (*gin.Engine, error) {
	r := gin.Default()

	//setting a CORS
	cors := middleware.CORS()
	r.Use(cors)

	v1 := r.Group("/v1")
	{
		systemHandler := handler.NewSystemHandler()
		authHandler := handler.NewAuthHandler()
		v1.GET("/health",systemHandler.Health)
		v1.POST("/auth",authHandler.SignUpHandler)
	}

	monsters := v1.Group("/monsters")
	monsterHandler := di.InitMonstersHandler()
	{
		monsters.GET("",monsterHandler.GetAll)
		monsters.GET("/:id",monsterHandler.GetById)
	}
	
	auth := v1.Group("/auth/monsters")
	auth.Use(middleware.AuthMiddleware)
	{
		auth.POST("",monsterHandler.Create)
		auth.POST("/json",monsterHandler.CreateJson)
		auth.PATCH("/:id",monsterHandler.Update)
		auth.DELETE("/:id",monsterHandler.Delete)
	}

	return r,nil
}