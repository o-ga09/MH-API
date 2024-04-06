package controller

import (
	di "mh-api/api/DI"
	"mh-api/api/config"
	"mh-api/api/controller/handler"
	"mh-api/api/middleware"

	"github.com/gin-gonic/gin"
)

func NewServer() (*gin.Engine, error) {
	r := gin.New()
	cfg, _ := config.New()
	if cfg.Env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	// setting logger
	logger := middleware.New()
	httpLogger := middleware.RequestLogger(logger)

	//setting a CORS
	cors := middleware.CORS()

	// With Context
	withCtx := middleware.WithTimeout()

	// With Request Id
	withReqId := middleware.AddID()

	r.Use(withReqId)
	r.Use(withCtx)
	r.Use(cors)
	r.Use(httpLogger)

	v1 := r.Group("/v1")
	{
		systemHandler := handler.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}

	monsters := v1.Group("/monsters")
	monsterHandler := di.InitMonstersHandler()
	{
		monsters.GET("", monsterHandler.GetAll)
		monsters.GET("/:id", monsterHandler.GetById)
	}

	return r, nil
}
