package main

import (
	"mh-api/app/internal/presenter"
	"mh-api/app/pkg"
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
	cfg, err := pkg.New()
	if err != nil {
		panic(err)
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.SentryDSN,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		panic(err)
	}
	defer sentry.Flush(2 * time.Second)

	s, err := presenter.NewServer()
	if err != nil {
		panic(err)
	}

	if err := s.Run(); err != nil {
		panic(err)
	}
}
