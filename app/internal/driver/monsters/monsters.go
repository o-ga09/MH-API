package monsters

import (
	"context"
	"mh-api/app/internal/domain/monsters"
	"mh-api/app/internal/driver/mysql"

	"gorm.io/gorm"
)

type monsterRepository struct {
	conn *gorm.DB
}

func NewMonsterRepository(conn *gorm.DB) *monsterRepository {
	return &monsterRepository{
		conn: conn,
	}
}

func (r *monsterRepository) Get(ctx context.Context, monsterId string) (monsters.Monsters, error) {
	monster := []mysql.Monster{}
	err := r.conn.Find(&monster).Error
	if err != nil {
		return nil, err
	}

	res := monsters.Monsters{}
	for _, r := range monster {
		res = append(res, monsters.NewMonster(r.MonsterId, r.Name, r.Description))
	}

	return res, nil
}

func (r *monsterRepository) Save(ctx context.Context, m monsters.Monster) error {
	monster := mysql.Monster{
		MonsterId:   m.GetId(),
		Name:        m.GetName(),
		Description: m.GetDesc(),
	}
	err := r.conn.Save(&monster).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *monsterRepository) Remove(ctx context.Context, monsterId string) error {
	monster := mysql.Monster{
		MonsterId: monsterId,
	}
	err := r.conn.Delete(&monster).Error
	if err != nil {
		return err
	}
	return nil
}
