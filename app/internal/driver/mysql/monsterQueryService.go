package mysql

import (
	"context"
	"mh-api/app/internal/service/monsters"

	"gorm.io/gorm"
)

type monsterQueryService struct {
	conn *gorm.DB
}

func NewmonsterQueryService(conn *gorm.DB) *monsterQueryService {
	return &monsterQueryService{
		conn: conn,
	}
}

func (s *monsterQueryService) FetchMonsterList(ctx context.Context, id string) ([]*monsters.FetchMonsterListDto, error) {
	var monster []Monster
	err := db.Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Find(&monster).Error
	if err != nil {
		return nil, err
	}

	res := []*monsters.FetchMonsterListDto{}
	for _, m := range monster {

		var l []string
		var t []string
		var weak_attack []monsters.Weakness_attack
		var weak_element []monsters.Weakness_element
		for _, f := range m.Field {
			l = append(l, f.Name)
		}
		for _, p := range m.Product {
			t = append(t, p.Name)
		}
		for _, wa := range m.Weakness {
			weak_attack = append(weak_attack, monsters.Weakness_attack{
				PartId:   wa.PartId,
				Slashing: wa.Slashing,
				Blow:     wa.Blow,
				Bullet:   wa.Bullet,
			})
		}
		for _, we := range m.Weakness {
			weak_element = append(weak_element, monsters.Weakness_element{
				PartId:  we.PartId,
				Fire:    we.Fire,
				Water:   we.Water,
				Thunder: we.Lightning,
				Ice:     we.Ice,
				Dragon:  we.Dragon,
			})
		}

		r := monsters.FetchMonsterListDto{
			Id:               m.MonsterId,
			Name:             m.Name,
			Description:      m.Description,
			Location:         l,
			Category:         m.Tribe.Name_ja,
			Title:            t,
			Weakness_attack:  weak_attack,
			Weakness_element: weak_element,
		}

		res = append(res, &r)
	}
	return res, err
}
