package monsters

import (
	"context"
	"mh-api/app/internal/domain/monsters"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/pkg"
)

type monsterRepository struct {
}

func NewMonsterRepository() *monsterRepository {
	return &monsterRepository{}
}

func (r *monsterRepository) Get(ctx context.Context, monsterId string) (monsters.Monsters, error) {
	monster := []mysql.Monster{}
	err := mysql.CtxFromDB(ctx).Find(&monster).Error
	if err != nil {
		return nil, err
	}

	res := monsters.Monsters{}
	for _, r := range monster {
		element := pkg.PtrToStr(r.Element)
		domainMonster := monsters.NewMonster(r.MonsterId, r.Name, r.Description, element)
		res = append(res, domainMonster)
	}

	return res, nil
}

func (r *monsterRepository) Save(ctx context.Context, m monsters.Monster) error {
	monster := mysql.Monster{
		MonsterId:   m.GetId(),
		Name:        m.GetName(),
		Description: m.GetDesc(),
		Element:     pkg.StrToPtr(m.GetElement()),
	}
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 0")
	err := mysql.CtxFromDB(ctx).Save(&monster).Error
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 1")
	if err != nil {
		return err
	}
	return nil
}

func (r *monsterRepository) Remove(ctx context.Context, monsterId string) error {
	monster := mysql.Monster{
		MonsterId: monsterId,
	}
	err := mysql.CtxFromDB(ctx).Delete(&monster).Error
	if err != nil {
		return err
	}
	return nil
}
