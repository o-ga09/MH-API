package ranking

import (
	"context"
	"mh-api/app/internal/domain/ranking"
	"mh-api/app/internal/driver/mysql"
	"strconv"

	"gorm.io/gorm"
)

type rankingRepository struct {
}

func NewMonsterRepository() *rankingRepository {
	return &rankingRepository{}
}

func (r *rankingRepository) Save(ctx context.Context, rank ranking.Ranking) error {
	data := mysql.Ranking{
		MonsterId: rank.GetID(),
		Ranking:   rank.GetRank(),
		VoteYear:  rank.GetVoteYear(),
	}
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 0")
	err := mysql.CtxFromDB(ctx).Save(&data).Error
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 1")
	if err != nil {
		return err
	}
	return nil
}

func (r *rankingRepository) Remove(ctx context.Context, Id string) error {
	i, _ := strconv.Atoi(Id)
	data := mysql.Ranking{
		Model: gorm.Model{ID: uint(i)},
	}
	err := mysql.CtxFromDB(ctx).Debug().Delete(&data).Error
	if err != nil {
		return err
	}
	return nil
}
