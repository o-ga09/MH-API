package main

import (
	"context"
	"fmt"
	"log/slog"
	"mh-api/app/internal/batch"
	"mh-api/app/internal/presenter/middleware"
	"os"
)

func main() {
	_ = middleware.New()
	ctx := context.Background()
	slog.InfoContext(ctx, "[Batch Started]")
	if len(os.Args) < 2 {
		fmt.Printf("number of batch argument more 2")
		os.Exit(0)
	}
	batchName := os.Args[1]
	err := batch.Exec(ctx, batchName)
	if err != nil {
		panic(err)
	}
	slog.InfoContext(ctx, "[Batch Ended]")
}
