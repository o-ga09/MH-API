package ranking

import (
	"context"
	"mh-api/app/internal/domain/ranking"
	"mh-api/app/internal/driver/mysql"
	"strconv"

	"gorm.io/gorm"
)

type rankingRepository struct {
	conn *gorm.DB
}

func NewMonsterRepository(conn *gorm.DB) *rankingRepository {
	return &rankingRepository{
		conn: conn,
	}
}

func (r *rankingRepository) Save(ctx context.Context, rank ranking.Ranking) error {
	data := mysql.Ranking{
		MonsterId: rank.GetID(),
		Ranking:   rank.GetRank(),
		VoteYear:  rank.GetVoteYear(),
	}
	r.conn.Exec("SET foreign_key_checks = 0")
	err := r.conn.Save(&data).Error
	r.conn.Exec("SET foreign_key_checks = 1")
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
	err := r.conn.Debug().Delete(&data).Error
	if err != nil {
		return err
	}
	return nil
}
