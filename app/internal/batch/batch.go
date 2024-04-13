package batch

import (
	"context"
	"fmt"
	"log/slog"
	di "mh-api/app/internal/DI"
	"mh-api/app/internal/service/monsters"
)

type BatchService struct {
	Service monsters.MonsterService
}

func NewBatchService(s monsters.MonsterService) *BatchService {
	return &BatchService{Service: s}
}

func Exec(ctx context.Context, batchName string) error {
	service := di.BatchDI()
	batchService := NewBatchService(*service)

	var saveData monsters.MonsterDto
	var id string

	switch batchName {
	case "createMonster":
		slog.InfoContext(ctx, fmt.Sprintf("[Batch Name]: %s", batchName))
		err := batchService.Service.SaveMonsters(ctx, saveData)
		if err != nil {
			return err
		}
	case "removeMonster":
		slog.InfoContext(ctx, fmt.Sprintf("[Batch Name]: %s", batchName))
		err := batchService.Service.RemoveMonsters(ctx, id)
		if err != nil {
			return err
		}
	case "createItem":
		slog.InfoContext(ctx, fmt.Sprintf("[Batch Name]: %s", batchName))
	case "removeItem":
		slog.InfoContext(ctx, fmt.Sprintf("[Batch Name]: %s", batchName))
	case "createWeapon":
		slog.InfoContext(ctx, fmt.Sprintf("[Batch Name]: %s", batchName))
	case "removeWeapon":
		slog.InfoContext(ctx, fmt.Sprintf("[Batch Name]: %s", batchName))
	}

	return nil
}
