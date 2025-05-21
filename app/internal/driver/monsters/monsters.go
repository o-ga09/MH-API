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
		domainMonster := monsters.NewMonster(r.MonsterId, r.Name, r.Description)
		domainMonster.Element = r.Element // Assign Element from db model to domain model
		res = append(res, domainMonster)
	}

	return res, nil
}

func (r *monsterRepository) Save(ctx context.Context, m monsters.Monster) error {
	monster := mysql.Monster{
		MonsterId:   m.GetId(),
		Name:        m.GetName(),
		Description: m.GetDesc(),
		Element:     m.Element, // Assign Element from domain model to db model
	}
	r.conn.Exec("SET foreign_key_checks = 0")
	err := r.conn.Save(&monster).Error
	r.conn.Exec("SET foreign_key_checks = 1")
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
