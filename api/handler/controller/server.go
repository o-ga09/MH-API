package controller

import (
	"context"
	"mh-api/api/config"
	"mh-api/api/handler/store"
	"mh-api/api/interface/monster"
	"mh-api/api/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	apiVersion = "/v1"
	langTagJa = "/ja"
	langTagEn = "/en"
)

func NewServer() (*gin.Engine, error) {
	r := gin.Default()

	//setting a CORS
	r.Use(cors.New(cors.Config{
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
	}))

	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	db, err := store.New(context.Background(),cfg)
	if err != nil {
		return nil, err
	}

	monsterRepo := store.NewMonsterRepository(db)
	monstetInterface := monster.NewMosterService(monsterRepo)
	monsterService := service.NewMonsterUsecase(monstetInterface)

	tagJa := r.Group(apiVersion + langTagJa)

	{
		systemHandler := NewSystemHandler()
		monsterHandler := NewMonsterHandler(monsterService)
		tagJa.GET("/system/health",systemHandler.Health)

		tagJa.GET("/monster/:id",monsterHandler.GetMonsterById)
		tagJa.GET("/monsters",monsterHandler.GetMonsterAll)
	}

	return r,nil
}