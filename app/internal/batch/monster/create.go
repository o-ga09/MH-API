package monster

import (
	"context"
	"log/slog"
	"mh-api/app/internal/batch"
	"mh-api/app/internal/service/monsters"

	"mh-api/app/pkg"
)

func Create(ctx context.Context, batchService *batch.BatchService) error {
	data, err := pkg.GetCSV(ctx, "monster/monsters.csv")
	if err != nil {
		slog.InfoContext(ctx, "[Batch Error Occurred]: csv from GCS", "error message", err)
		return err
	}

	for _, r := range *data {
		m := monsters.MonsterDto{
			ID:          r[0],
			Name:        r[1],
			Description: r[2],
		}
		err := batchService.Service.SaveMonsters(ctx, m)
		if err != nil {
			slog.InfoContext(ctx, "[Batch Error Occurred]: data Insert to DB", "error message", err)
			return err
		}
		slog.InfoContext(ctx, "[Batch]: Insert to DB", "data", m)
	}
	return nil
}
