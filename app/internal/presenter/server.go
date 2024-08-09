package presenter

import (
	di "mh-api/app/internal/DI"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/pkg"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func NewServer() (*gin.Engine, error) {
	r := gin.New()
	cfg, _ := pkg.New()
	if cfg.Env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	// バリデーション
	var validId validator.Func = func(fieldLevel validator.FieldLevel) bool {
		Id, ok := fieldLevel.Field().Interface().(string)
		if !ok {
			return false
		}

		if Id == "" {
			// itemIdsは、requieredではないので、空文字で良い
			return true
		}

		Ids := strings.Split(Id, ",")
		for _, id := range Ids {
			if len(id) != 10 {
				return false
			}
		}
		return true
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validateId", validId)
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

	// アイテム検索
	item := v1.Group("/items")
	itemHandler := di.InitItemHaandler()
	{
		item.GET("", itemHandler.GetItems)
		item.GET("/:id", itemHandler.GetItem)
		item.GET("/monsters", itemHandler.GetItemByMonster)
		item.GET("/monsters/:id", itemHandler.GetItemByMonsterId)
	}

	return r, nil
}
