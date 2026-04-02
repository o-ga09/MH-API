package main

import (
	"context"
	"mh-api/internal/presenter"
	"mh-api/pkg/config"
	"mh-api/pkg/telemetry"
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
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	// OpenTelemetryの初期化
	shutdown, err := telemetry.InitTracer(ctx, "mh-api", cfg.OtelExporterOtlpEndpoint)
	if err != nil {
		panic(err)
	}
	defer shutdown(ctx)

	s, err := presenter.NewServer()
	if err != nil {
		panic(err)
	}

	if err := s.Run(); err != nil {
		panic(err)
	}
}
