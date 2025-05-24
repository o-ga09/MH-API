package presenter

import (
	"context"
	di "mh-api/internal/DI"
	"mh-api/internal/presenter/middleware"
	"mh-api/pkg/config"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/propagators/autoprop"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewServer() (*gin.Engine, error) {
	ctx := context.Background()
	r := gin.New()
	cfg, _ := config.New()
	if cfg.Env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	tp := trace.NewTracerProvider()
	defer tp.Shutdown(ctx)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(autoprop.NewTextMapPropagator())

	// ロガー設定
	middleware.New()

	// CORS設定
	cors := middleware.CORS()

	// リクエストタイムアウト設定
	withCtx := middleware.WithTimeout()

	// Sentry設定
	sentryMiddleware := middleware.SentryTracingMiddleware(gin.Logger())

	// ミドルウェア設定
	r.Use(otelgin.Middleware("mh-api"))
	r.Use(withCtx)
	r.Use(cors)
	r.Use(middleware.RequestLogger())
	r.Use(middleware.WithDB())
	r.Use(sentryMiddleware)

	// ヘルスチェック
	v1 := r.Group("/v1")
	{
		systemHandler := di.InitHealthService(ctx)
		v1.GET("/health", systemHandler.Health)
		v1.GET("/health/db", systemHandler.DBHealth)
	}

	// モンスター検索
	monsterHandler := di.InitMonstersHandler(ctx)
	monsters := v1.Group("/monsters")
	{
		monsters.GET("", monsterHandler.GetAll)
		monsters.GET("/:id", monsterHandler.GetById)
	}

	// アイテム検索
	itemHandler := di.InitItemsHandler(ctx) // di.InitItemsHandler を呼び出して ItemHandler を初期化
	items := v1.Group("/items")
	{
		items.GET("", itemHandler.GetItems) // ItemHandler の GetItems メソッドをルーティングに設定
	}

	// 武器検索
	weaponHandler := di.InitWeaponHandler(ctx)
	weapons := v1.Group("/weapons")
	{
		weapons.GET("", weaponHandler.SearchWeapons)
	}
	return r, nil
}
