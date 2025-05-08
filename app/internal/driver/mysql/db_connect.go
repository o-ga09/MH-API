package mysql

import (
	"context"
	"log/slog"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/pkg"

	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
)

var db *gorm.DB

func New(ctx context.Context) *gorm.DB {
	cfg, err := pkg.New()
	if err != nil {
		slog.Log(ctx, middleware.SeverityError, "environment variable error", "error", err)
		return nil
	}

	dialector := mysql.Open(cfg.Database_url)

	if db, err = gorm.Open(dialector, &gorm.Config{
		Logger: NewSentryLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// コールバックでスパンを作成するよう設定
		QueryFields: true,
	}); err != nil {
		connect(ctx, dialector, 100)
	}

	// SQLクエリをトレースできるようにOpenTelemetryプラグインを登録
	if err := db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		slog.Log(ctx, middleware.SeverityError, "failed to register otelgorm plugin", "error", err)
	}

	return db
}

func connect(ctx context.Context, dialector gorm.Dialector, count uint) {
	var err error
	if db, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			slog.Log(ctx, middleware.SeverityInfo, "db connection retry")
			connect(ctx, dialector, count)
			return
		}
		slog.Log(ctx, middleware.SeverityInfo, "db connection retry count 100")
		panic(err.Error())
	}

	// リトライ接続でもトレースを有効化
	cfg, _ := pkg.New()
	if cfg.Env == "PROD" || cfg.Env == "STAGE" {
		if err := db.Use(otelgorm.NewPlugin(
			otelgorm.WithDBName("mh-api"),
			otelgorm.WithAttributes(
				attribute.String("db.type", "mysql"),
				attribute.String("db.env", cfg.Env),
				attribute.Bool("db.retry_connection", true),
			),
		)); err != nil {
			slog.Log(ctx, middleware.SeverityError, "failed to register otelgorm plugin in retry", "error", err)
		}
	}
}
