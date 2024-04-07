package mysql

import (
	"context"
	param "mh-api/app/internal/controller/monster"
	"mh-api/app/internal/service/monsters"
	"strings"

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
	var monsterIds []string
	var result *gorm.DB
	var p param.RequestParam
	var err error

	where_clade := ""
	sort := ""

	if id == "" {
		p = ctx.Value("param").(param.RequestParam)
	}

	limit := p.Limit
	offset := p.Offset

	if p.MonsterIds != "" {
		monsterIds = strings.Split(p.MonsterIds, ",")
		where_clade = "monster_id IN (?)"
	}

	if p.MonsterName != "" && p.MonsterIds != "" {
		where_clade += " and name LIKE '%" + p.MonsterName + "%' "
	} else if p.MonsterName != "" {
		where_clade += " name LIKE '%" + p.MonsterName + "%' "
	}

	if p.Sort == "1" {
		sort = "monster_id ASC"
	} else {
		sort = "monster_id DESC"
	}

	if id != "" {
		result = db.Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Where("monster_id = ? ", id).Find(&monster)
	} else if where_clade != "" && p.MonsterIds != "" {
		result = db.Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Where(where_clade, monsterIds).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else if where_clade != "" {
		result = db.Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Where(where_clade).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else {
		result = db.Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Limit(limit).Offset(offset).Order(sort).Find(&monster)
	}

	if result.Error != nil {
		return nil, err
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
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

func (s *monsterQueryService) FetchMonsterRanking(ctx context.Context) ([]*monsters.FetchMonsterRankingDto, error) {
	var monster []Monster
	var monsterIds []string
	var result *gorm.DB
	var p param.RequestParam
	var err error

	where_clade := ""
	sort := ""

	p = ctx.Value("param").(param.RequestParam)

	limit := p.Limit
	offset := p.Offset

	if p.MonsterIds != "" {
		monsterIds = strings.Split(p.MonsterIds, ",")
		where_clade = "monster_id IN (?)"
	}

	if p.MonsterName != "" && p.MonsterIds != "" {
		where_clade += " and name LIKE '%" + p.MonsterName + "%' "
	} else if p.MonsterName != "" {
		where_clade += " name LIKE '%" + p.MonsterName + "%' "
	}

	if p.Sort == "1" {
		sort = "monster_id ASC"
	} else {
		sort = "monster_id DESC"
	}

	if where_clade != "" && p.MonsterIds != "" {
		result = db.Model(&monster).Preload("Field").Preload("Tribe").Preload("Product").Where(where_clade, monsterIds).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else if where_clade != "" {
		result = db.Model(&monster).Preload("Field").Preload("Tribe").Preload("Product").Where(where_clade).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else {
		result = db.Model(&monster).Preload("Field").Preload("Tribe").Preload("Product").Limit(limit).Offset(offset).Order(sort).Find(&monster)
	}

	if result.Error != nil {
		return nil, err
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	res := []*monsters.FetchMonsterRankingDto{}
	for _, m := range monster {
		var l []string
		var t []string
		var rank []monsters.Ranking

		for _, f := range m.Field {
			l = append(l, f.Name)
		}
		for _, p := range m.Product {
			t = append(t, p.Name)
		}

		for _, r := range m.Ranking {
			rank = append(rank, monsters.Ranking{
				Ranking:  r.Ranking,
				VoteYear: r.VoteYear,
			})
		}

		r := monsters.FetchMonsterRankingDto{
			Id:          m.MonsterId,
			Name:        m.Name,
			Description: m.Description,
			Location:    l,
			Category:    m.Tribe.Name_ja,
			Title:       t,
			Ranking:     rank,
		}
		res = append(res, &r)
	}
	return res, err
}
