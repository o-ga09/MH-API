package batch

import (
	"context"
	"fmt"
	"log/slog"

	"mh-api/app/internal/domain/fields"
	"mh-api/app/internal/domain/items"
	"mh-api/app/internal/domain/monsters"
	"mh-api/app/internal/domain/music"
	"mh-api/app/internal/domain/part"
	Products "mh-api/app/internal/domain/products"
	"mh-api/app/internal/domain/ranking"
	Tribes "mh-api/app/internal/domain/tribes"
	"mh-api/app/internal/domain/weakness"
	"mh-api/app/internal/domain/weapons"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/pkg"
)

func Create(ctx context.Context, batchService *BatchService) error {
	// データを作成するテーブル
	tables := []string{"monster", "item", "field", "product", "part", "music", "tribe", "weakness", "weapon", "ranking"}

	for _, table := range tables {
		data, err := pkg.GetCSV(ctx, fmt.Sprintf("monster/%s.csv", table))
		if err != nil {
			slog.Log(ctx, middleware.SeverityWarn, "[Batch]: Insert DB Skipped", "table name", table)
			continue
		}

		err = InsertDB(ctx, batchService, data, table)
		if err != nil {
			slog.Log(ctx, middleware.SeverityWarn, "[Batch Error Occurred]: Failed to insert to DB", "error message", err)
			return err
		}
		slog.Log(ctx, middleware.SeverityInfo, "[Batch]: Insert to DB", "table", table)
	}

	return nil
}

func InsertDB(ctx context.Context, batchService *BatchService, data *[][]string, tableName string) error {
	switch tableName {
	case "monster":
		for _, r := range *data {
			m := monsters.NewMonster(r[0], r[1], r[2])
			err := batchService.monsterService.Save(ctx, m)
			if err != nil {
				slog.Log(ctx, middleware.SeverityError, "[Batch Error Occurred]: data Insert to DB", "error message", err)
				return err
			}
		}
	case "item":
		for _, r := range *data {
			m := items.NewItem(r[0], r[1], r[2], r[3])
			err := batchService.itemService.Save(ctx, *m)
			if err != nil {
				slog.Log(ctx, middleware.SeverityError, "[Batch Error Occurred]: data Insert to DB", "error message", err)
				return err
			}
		}
	case "field":
		for _, r := range *data {
			m := fields.NewField(r[0], r[1], r[2], r[3])
			err := batchService.fieldService.Save(ctx, *m)
			if err != nil {
				slog.Log(ctx, middleware.SeverityError, "[Batch Error Occurred]: data Insert to DB", "error message", err)
				return err
			}
		}
	case "product":
		for _, r := range *data {
			m := Products.NewProduct(r[0], r[1], r[2], r[3])
			err := batchService.productService.Save(ctx, *m)
			if err != nil {
				slog.Log(ctx, middleware.SeverityError, "[Batch Error Occurred]: data Insert to DB", "error message", err)
				return err
			}
		}
	case "ranking":
		for _, r := range *data {
			m := ranking.NewRanking(r[0], r[1], r[2])
			err := batchService.rankingService.Save(ctx, *m)
			if err != nil {
				slog.Log(ctx, middleware.SeverityError, "[Batch Error Occurred]: data Insert to DB", "error message", err)
				return err
			}
		}
	case "tribe":
		for _, r := range *data {
			m := Tribes.NewTribe(r[0], r[1], r[2], r[3], r[4])
			err := batchService.tribeService.Save(ctx, *m)
			if err != nil {
				slog.Log(ctx, middleware.SeverityError, "[Batch Error Occurred]: data Insert to DB", "error message", err)
				return err
			}
		}
	case "part":
		for _, r := range *data {
			m := part.NewPart(r[0], r[1], r[2], r[3])
			err := batchService.partService.Save(ctx, *m)
			if err != nil {
				slog.Log(ctx, middleware.SeverityError, "[Batch Error Occurred]: data Insert to DB", "error message", err)
				return err
			}
		}
	case "music":
		for _, r := range *data {
			m := music.NewMusic(r[0], r[1], r[2], r[3])
			err := batchService.musicService.Save(ctx, *m)
			if err != nil {
				slog.Log(ctx, middleware.SeverityError, "[Batch Error Occurred]: data Insert to DB", "error message", err)
				return err
			}
		}
	case "weakness":
		for _, r := range *data {
			m := weakness.NewWeakness(r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8], r[9], r[10], r[11], r[11], r[12])
			err := batchService.weaknessService.Save(ctx, *m)
			if err != nil {
				slog.Log(ctx, middleware.SeverityError, "[Batch Error Occurred]: data Insert to DB", "error message", err)
				return err
			}
		}
	case "weapon":
		for _, r := range *data {
			m := weapons.NewWeapon(r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8])
			err := batchService.weaponService.Save(ctx, *m)
			if err != nil {
				slog.Log(ctx, middleware.SeverityError, "[Batch Error Occurred]: data Insert to DB", "error message", err)
				return err
			}
		}
	}

	return nil
}
