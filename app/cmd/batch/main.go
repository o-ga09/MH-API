package main

import (
	"context"
	"log/slog"
	"mh-api/app/internal/batch"
	"mh-api/app/internal/presenter/middleware"
	"os"
	"github.com/getsentry/sentry-go"
	"time"
	"mh-api/app/pkg"
)

func main() {
	cfg, err := pkg.New()
	if err != nil {
		panic(err)
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: cfg.SentryDSN,
	}); err != nil {
		panic(err)
	}
	defer sentry.Flush(2 * time.Second)

	_ = middleware.New()
	ctx := context.Background()
	slog.InfoContext(ctx, "[Batch Started]")
	if len(os.Args) < 2 {
		slog.Log(ctx, middleware.SeverityInfo, "[Batch]: number of batch argument more 2")
		os.Exit(0)
	}
	batchName := os.Args[1]
	err = batch.Exec(ctx, batchName)
	if err != nil {
		slog.Log(ctx, middleware.SeverityInfo, "[Batch Accident]")
	}
	slog.Log(ctx, middleware.SeverityInfo, "[Batch Ended]")
}
