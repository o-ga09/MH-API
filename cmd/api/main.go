package main

import (
	"mh-api/internal/presenter"
	"mh-api/pkg/config"
	"time"

	"github.com/getsentry/sentry-go"
)

//		@title			MH-API
//		@version		v0.1.0
//		@description	モンスターハンターAPI
//		@host			https://mh-api-v2-8aznfogc.an.gateway.dev/
//	 @BasePath  /v1
//	 @license.name  Apache 2.0
//	 @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
//	 @externalDocs.description  OpenAPI
//	 @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	if cfg.Env == "PROD" || cfg.Env == "STAGE" {
		// Sentryの初期化設定を強化
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:              cfg.SentryDSN,
			EnableTracing:    true,
			TracesSampleRate: 1.0, // 本番環境では適切な値に調整することを推奨 (0.0-1.0)
			TracesSampler: func(ctx sentry.SamplingContext) float64 {
				// SQLクエリを含むトレースは優先的に収集
				if ctx.Span != nil && (ctx.Span.Op == "db.sql.query" || ctx.Span.Op == "db.sql.exec") {
					return 1.0
				}
				return 1.0 // デフォルトは全トレース取得
			},
			Debug:            cfg.Env == "STAGE", // STAGEのみデバッグログを有効化
			AttachStacktrace: true,               // スタックトレース情報を付加
			// サービス名を設定
			ServerName: "mh-api",
			// 環境名を設定
			Environment: cfg.Env,
		}); err != nil {
			panic(err)
		}
		defer sentry.Flush(2 * time.Second)
	}

	s, err := presenter.NewServer()
	if err != nil {
		panic(err)
	}

	if err := s.Run(); err != nil {
		panic(err)
	}
}
