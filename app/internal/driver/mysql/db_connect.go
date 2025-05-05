package mysql

import (
	"context"
	"log/slog"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/pkg"

	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var tracer = otel.Tracer("mh-api/mysql")

func New(ctx context.Context) *gorm.DB {
	// データベース接続のスパンを作成
	ctx, span := tracer.Start(ctx, "mysql.Connect", trace.WithAttributes(
		attribute.String("db.system", "mysql"),
	))
	defer span.End()

	cfg, err := pkg.New()
	if err != nil {
		span.RecordError(err)
		slog.Log(context.Background(), middleware.SeverityError, "environment variable error", "error", err)
	}

	dialector := mysql.Open(cfg.Database_url)
	span.SetAttributes(attribute.String("db.connection_string", cfg.Database_url))

	if db, err = gorm.Open(dialector, &gorm.Config{
		Logger: NewSentryLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// コールバックでスパンを作成するよう設定
		QueryFields: true,
	}); err != nil {
		span.RecordError(err)
		connect(ctx, dialector, 100)
	}

	// SQLクエリをトレースできるようにOpenTelemetryプラグインを登録
	if cfg.Env == "PROD" || cfg.Env == "STAGE" {
		if err := db.Use(otelgorm.NewPlugin(
			otelgorm.WithDBName("mh-api"),
			otelgorm.WithQueryFormatter(func(query string) string {
				return query // 必要に応じてクエリをフォーマットできます
			}),
			otelgorm.WithAttributes(
				attribute.String("db.type", "mysql"),
				attribute.String("db.env", cfg.Env),
			),
			otelgorm.WithoutQueryVariables(), // 機密情報のロギングを防ぐ
		)); err != nil {
			span.RecordError(err)
			slog.Log(ctx, middleware.SeverityError, "failed to register otelgorm plugin", "error", err)
		} else {
			slog.Log(ctx, middleware.SeverityInfo, "otelgorm plugin registered for SQL tracing")
		}
	} else {
		// 開発環境ではSQLログを出力
		db.Logger = db.Logger.LogMode(logger.Info)
	}

	// 接続情報のトレーシングを追加
	sqlDB, err := db.DB()
	if err == nil {
		// 接続プールの統計情報を記録
		stats := sqlDB.Stats()
		span.SetAttributes(
			attribute.Int("db.pool.max_open_connections", stats.MaxOpenConnections),
			attribute.Int("db.pool.open_connections", stats.OpenConnections),
			attribute.Int("db.pool.in_use", stats.InUse),
			attribute.Int("db.pool.idle", stats.Idle),
		)
	}

	slog.Log(ctx, middleware.SeverityInfo, "db connected")
	return db
}

func connect(ctx context.Context, dialector gorm.Dialector, count uint) {
	ctx, span := tracer.Start(ctx, "mysql.RetryConnect", trace.WithAttributes(
		attribute.Int("retry.count", int(count)),
	))
	defer span.End()

	var err error
	if db, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		span.RecordError(err)
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
			span.RecordError(err)
			slog.Log(context.Background(), middleware.SeverityError, "failed to register otelgorm plugin in retry", "error", err)
		}
	}
}
