package mysql

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"mh-api/pkg/config"
	"sync"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type ctxKey string

const CtxKey ctxKey = "db"
const MAX_RETRY = 10

var (
	db   *gorm.DB
	once sync.Once
)

func New(ctx context.Context) context.Context {
	once.Do(func() {
		cfg, err := config.New()
		if err != nil {
			return
		}

		dialector := mysql.Open(cfg.Database_url)
		db, err = gorm.Open(dialector, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			ctx = connect(ctx, dialector)
		}

		// SQLDBインスタンスを取得
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}

		// コネクションプールの設定
		sqlDB.SetMaxIdleConns(10)           // アイドル状態の最大接続数
		sqlDB.SetMaxOpenConns(100)          // 最大接続数
		sqlDB.SetConnMaxLifetime(time.Hour) // 接続の最大生存期間
	})

	return context.WithValue(ctx, CtxKey, db)
}

func connect(ctx context.Context, dialector gorm.Dialector) context.Context {
	var err error
	for i := 0; i < MAX_RETRY; i++ {
		if db, err = gorm.Open(dialector, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
			},
			Logger: logger.Default.LogMode(logger.Info),
		}); err == nil {
			return context.WithValue(ctx, CtxKey, db)
		}
		time.Sleep(5 * time.Second)
		slog.InfoContext(ctx, fmt.Sprintf("Failed to connect to database. Retry after 5 seconds.: %d times", i))
	}
	return ctx
}

func SetTestDB(t *testing.T, db *gorm.DB) {
	t.Helper()

	tx := db.Begin() // トランザクションを開始

	// テスト終了後にロールバック
	defer tx.Rollback()
}

func CtxFromDB(ctx context.Context) *gorm.DB {
	return ctx.Value(CtxKey).(*gorm.DB)
}
