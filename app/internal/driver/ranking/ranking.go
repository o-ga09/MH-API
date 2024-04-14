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

func (r *rankingRepository) Get(ctx context.Context, monsterId string) (ranking.Rankings, error) {
	rank := []mysql.Ranking{}
	err := r.conn.Find(&rank).Error
	if err != nil {
		return nil, err
	}

	res := ranking.Rankings{}
	for _, r := range rank {
		res = append(res, *ranking.NewRanking(r.MonsterId, r.Ranking, r.VoteYear))
	}

	return res, nil
}

func (r *rankingRepository) Save(ctx context.Context, rank ranking.Ranking) error {
	data := mysql.Ranking{
		MonsterId: rank.GetID(),
		Ranking:   rank.GetRank(),
		VoteYear:  rank.GetVoteYear(),
	}
	err := r.conn.Save(&data).Error
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
