package presenter

import (
	di "mh-api/app/internal/DI"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/pkg"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func NewServer() (*gin.Engine, error) {
	r := gin.New()
	cfg, _ := pkg.New()
	if cfg.Env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	// ロガー設定
	logger := middleware.New()
	httpLogger := middleware.RequestLogger(logger)

	// CORS設定
	cors := middleware.CORS()

	// リクエストタイムアウト設定
	withCtx := middleware.WithTimeout()

	// リクエストID付与
	withReqId := middleware.AddID()

	// ミドルウェア設定
	r.Use(withReqId)
	r.Use(withCtx)
	r.Use(cors)
	r.Use(httpLogger)
	r.Use(sentrygin.New(sentrygin.Options{}))

	// ヘルスチェック
	v1 := r.Group("/v1")
	{
		systemHandler := di.InitHealthService()
		v1.GET("/health", systemHandler.Health)
		v1.GET("/health/db", systemHandler.DBHealth)
	}

	// モンスター検索
	monsters := v1.Group("/monsters")
	monsterHandler := di.InitMonstersHandler()
	{
		monsters.GET("", monsterHandler.GetAll)
		monsters.GET("/:id", monsterHandler.GetById)
		monsters.GET("/ranking", monsterHandler.GetRankingMonster)
	}

	// BGM検索
	bgm := v1.Group("/bgms")
	bgmHandler := di.InitBGMHandler()
	{
		bgm.GET("", bgmHandler.GetBGM)
		bgm.GET("/:id", bgmHandler.GetBGMById)
		bgm.GET("/ranking", bgmHandler.GetRankingBGM)
	}

	return r, nil
}
