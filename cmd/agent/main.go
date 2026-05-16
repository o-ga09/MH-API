package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"mh-api/internal/agent"
	"mh-api/pkg/config"
	"mh-api/pkg/metrics"
	"mh-api/pkg/profiler"
	"mh-api/pkg/telemetry"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Prometheusメトリクスの初期化
	shutdownMeter, metricsHandler, err := telemetry.InitMeter(ctx, "mh-agent")
	if err != nil {
		log.Fatalf("Failed to init meter: %v", err)
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if meterErr := shutdownMeter(shutdownCtx); meterErr != nil {
			log.Printf("Failed to shutdown meter: %v", meterErr)
		}
	}()

	// metrics 専用サーバーを別ポートで起動
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
			log.Fatalf("Failed to start metrics server: %v", listenErr)
		}
	}()

	stopProfiler := profiler.StartPyroscope(cfg, "mh-agent")
	defer stopProfiler()

	if cfg.GeminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is required")
	}

	agentMetrics, err := metrics.NewAgentMetrics()
	if err != nil {
		log.Fatalf("Failed to init agent metrics: %v", err)
	}

	agentCfg := &agent.Config{
		GeminiAPIKey: cfg.GeminiAPIKey,
		GeminiModel:  cfg.GeminiModel,
		Port:         cfg.Port,
	}

	server, err := agent.NewServer(agentCfg, agentMetrics)
	if err != nil {
		log.Fatalf("Failed to create agent server: %v", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
