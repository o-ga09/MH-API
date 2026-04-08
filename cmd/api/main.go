package main

import (
	"context"
	"fmt"
	"mh-api/internal/presenter"
	"mh-api/pkg/config"
	"mh-api/pkg/profiler"
	"mh-api/pkg/telemetry"
	"net/http"
	"time"
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
	shutdown, err := telemetry.InitTracer(ctx, "mh-api", cfg.OtelExporterOtlpEndpoint, cfg.OtelInsecure)
	if err != nil {
		panic(err)
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if shutdownErr := shutdown(shutdownCtx); shutdownErr != nil {
			panic(shutdownErr)
		}
	}()

	// Prometheusメトリクスの初期化
	shutdownMeter, metricsHandler, err := telemetry.InitMeter(ctx, "mh-api")
	if err != nil {
		panic(err)
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if meterErr := shutdownMeter(shutdownCtx); meterErr != nil {
			panic(meterErr)
		}
	}()

	// metrics 専用サーバーを別ポートで起動（APIポートからは /metrics にアクセス不可）
	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", metricsHandler)
	go func() {
		addr := fmt.Sprintf(":%s", cfg.MetricsPort)
		srv := &http.Server{
			Addr:         addr,
			Handler:      metricsMux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		if listenErr := srv.ListenAndServe(); listenErr != nil {
			panic(listenErr)
		}
	}()

	// Pyroscopeプロファイラの初期化
	stopProfiler := profiler.StartPyroscope(cfg, "mh-api")
	defer stopProfiler()

	s, err := presenter.NewServer()
	if err != nil {
		panic(err)
	}

	if err := s.Run(); err != nil {
		panic(err)
	}
}
