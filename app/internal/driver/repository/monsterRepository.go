package repository

import (
	"context"
	monsterdomain "mh-api/app/internal/domain/monsterDomain"
	"mh-api/app/internal/driver/schema"

	"gorm.io/gorm"
)

type MonsterRepository struct {
	repo *gorm.DB
}

func NewMonsterRepository(repo *gorm.DB) *MonsterRepository {
	return &MonsterRepository{repo: repo}
}

func (r *MonsterRepository) Save(ctx context.Context, m *monsterdomain.Monster) error {
	monster := schema.Monster{
		MonsterId:   m.GetID(),
		Name:        m.GetName(),
		Description: m.GetDesc(),
	}
	if err := r.repo.Table("monster").Save(&monster).Error; err != nil {
		return err
	}
	return nil
}

func (r *MonsterRepository) Remove(ctx context.Context, monsterId string) error {
	if err := r.repo.Where("id = ?", monsterId).Delete(&monsterdomain.Monster{}).Error; err != nil {
		return err
	}
	return nil
}
