package mysql

import (
	"context"

	"mh-api/internal/domain/monsters"
	"mh-api/pkg/ptr"
)

type monsterRepository struct {
}

func NewMonsterRepository() *monsterRepository {
	return &monsterRepository{}
}

func (r *monsterRepository) Get(ctx context.Context, monsterId string) (monsters.Monsters, error) {
	monster := []Monster{}
	err := CtxFromDB(ctx).Find(&monster).Error
	if err != nil {
		return nil, err
	}

	res := monsters.Monsters{}
	for _, r := range monster {
		element := ptr.PtrToStr(r.Element)
		domainMonster := monsters.NewMonster(r.MonsterId, r.Name, r.Description, element)
		res = append(res, domainMonster)
	}

	return res, nil
}

func (r *monsterRepository) Save(ctx context.Context, m monsters.Monster) error {
	monster := Monster{
		MonsterId:   m.GetId(),
		Name:        m.GetName(),
		Description: m.GetDesc(),
		Element:     ptr.StrToPtr(m.GetElement()),
	}
	CtxFromDB(ctx).Exec("SET foreign_key_checks = 0")
	err := CtxFromDB(ctx).Save(&monster).Error
	CtxFromDB(ctx).Exec("SET foreign_key_checks = 1")
	if err != nil {
		return err
	}
	return nil
}

func (r *monsterRepository) Remove(ctx context.Context, monsterId string) error {
	monster := Monster{
		MonsterId: monsterId,
	}
	err := CtxFromDB(ctx).Delete(&monster).Error
	if err != nil {
		return err
	}
	return nil
}
