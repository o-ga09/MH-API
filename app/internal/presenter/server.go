package presenter

import (
	"context"
	di "mh-api/app/internal/DI"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/pkg"

	"github.com/gin-gonic/gin"
)

func NewServer(ctx context.Context) (*gin.Engine, error) {
	r := gin.New()
	cfg, _ := pkg.New()
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
		systemHandler := di.InitHealthService()
		v1.GET("/health", systemHandler.Health)
		v1.GET("/health/db", systemHandler.DBHealth)
	}

	monsters := v1.Group("/monsters")
	monsterHandler := di.ProvideMonsterHandler(ctx)
	{
		monsters.GET("", monsterHandler.GetAll)
		monsters.GET("/:id", monsterHandler.GetById)
		monsters.GET("/ranking", monsterHandler.GetRankingMonster)
	}

	items := v1.Group("/items")
	itemsHandler := di.ProvideItemHandler(ctx)
	{
		items.GET("", itemsHandler.GetItems)
		items.GET("/:id", itemsHandler.GetItem)
		items.GET("/:id/monsters", itemsHandler.GetItemByMonster)
	}

	return r, nil
}
