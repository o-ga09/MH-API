package presenter

import (
	"context"
	di "mh-api/app/internal/DI"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/pkg"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/propagators/autoprop"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewServer() (*gin.Engine, error) {
	ctx := context.Background()
	r := gin.New()
	cfg, _ := pkg.New()
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
	r.Use(sentryMiddleware)

	// ヘルスチェック
	v1 := r.Group("/v1")
	{
		systemHandler := di.InitHealthService(ctx)
		v1.GET("/health", systemHandler.Health)
		v1.GET("/health/db", systemHandler.DBHealth)
	}

	// モンスター検索
	monsters := v1.Group("/monsters")
	monsterHandler := di.InitMonstersHandler(ctx)
	{
		monsters.GET("", monsterHandler.GetAll)
		monsters.GET("/:id", monsterHandler.GetById)
	}

	return r, nil
}
