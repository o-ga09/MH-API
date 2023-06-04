package controller

import (
	di "mh-api/api/DI"
	"mh-api/api/controller/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer() (*gin.Engine, error) {
	r := gin.Default()

	//setting a CORS
	cors := CORS()
	r.Use(cors)

	v1 := r.Group("/v1")
	{
		systemHandler := handler.NewSystemHandler()
		v1.GET("/health",systemHandler.Health)
	}

	monsters := v1.Group("/monsters")
	{
		monsterHandler := di.InitMonstersHandler()
		monsters.GET("",monsterHandler.GetrAll)
		monsters.GET("/:id",monsterHandler.GetById)
		monsters.POST("",monsterHandler.Create)
		monsters.PATCH("/:id",monsterHandler.Update)
		// monsters.DELETE("/:id",monsterHandler.Delete)
	}

	return r,nil
}

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowCredentials: false,
		MaxAge: 24 * time.Hour,
	})
}