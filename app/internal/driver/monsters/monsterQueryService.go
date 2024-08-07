package monsters

import (
	"context"
	"fmt"
	param "mh-api/app/internal/controller/monster"
	"mh-api/app/internal/driver/mysql"
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

func (s *monsterQueryService) FetchList(ctx context.Context, id string) ([]*monsters.FetchMonsterListDto, error) {
	var monster []mysql.Monster
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
		result = s.conn.Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Where("monster_id = ? ", id).Find(&monster)
	} else if where_clade != "" && p.MonsterIds != "" {
		result = s.conn.Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Where(where_clade, monsterIds).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else if where_clade != "" {
		result = s.conn.Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Where(where_clade).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else {
		result = s.conn.Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Limit(limit).Offset(offset).Order(sort).Find(&monster)
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
			Id:                 m.MonsterId,
			Name:               m.Name,
			Description:        m.Description,
			Location:           l,
			Category:           m.Tribe.Name_ja,
			Title:              t,
			FirstWeak_Attack:   m.Weakness[0].FirstWeakAttack,
			SecondWeak_Attack:  m.Weakness[0].SecondWeakAttack,
			FirstWeak_Element:  m.Weakness[0].FirstWeakElement,
			SecondWeak_Element: m.Weakness[0].SecondWeakElement,
			Weakness_attack:    weak_attack,
			Weakness_element:   weak_element,
		}
		res = append(res, &r)
	}
	return res, err
}

func (s *monsterQueryService) FetchRank(ctx context.Context) ([]*monsters.FetchMonsterRankingDto, error) {
	var monster []mysql.Monster
	var monsterIds []string
	var result *gorm.DB
	var p param.RequestRankingParam
	var err error

	where_clade := ""

	where_clade_field := ""
	where_clade_tribe := ""
	where_clade_title := ""

	sort := ""

	p = ctx.Value("param").(param.RequestRankingParam)

	limit := p.Limit
	offset := p.Offset

	if p.MonsterIds != "" {
		monsterIds = strings.Split(p.MonsterIds, ",")
		where_clade = "monster_id IN (?)"
	}

	if p.MonsterName != "" && where_clade != "" {
		where_clade += " and name LIKE '%" + p.MonsterName + "%' "
	} else if p.MonsterName != "" {
		where_clade += " name LIKE '%" + p.MonsterName + "%' "
	}

	if p.LocationName != "" {
		where_clade_field += " name LIKE '%" + p.LocationName + "%' "
	}

	if p.TribeName != "" {
		where_clade_tribe += " name_ja LIKE '%" + p.TribeName + "%' "
	}

	if p.Title != "" {
		where_clade_title += " name LIKE '%" + p.Title + "%' "
	}

	if p.Sort == "2" {
		sort = "monster_id DESC"
	} else {
		sort = "monster_id ASC"
	}

	if where_clade != "" && p.MonsterIds != "" {
		result = s.conn.Model(&monster).Preload("Field", where_clade_field).Preload("Tribe", where_clade_tribe).Preload("Product", where_clade_title).Preload("Ranking").Preload("Weakness").Where(where_clade, monsterIds).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else if where_clade != "" {
		result = s.conn.Model(&monster).Preload("Field", where_clade_field).Preload("Tribe", where_clade_tribe).Preload("Product", where_clade_title).Preload("Ranking").Preload("Weakness").Where(where_clade).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else {
		result = s.conn.Model(&monster).Preload("Ranking").Preload("Field", where_clade_field).Preload("Tribe", where_clade_tribe).Preload("Product", where_clade_title).Preload("Weakness").Limit(limit).Offset(offset).Order(sort).Find(&monster)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("error: %s", result.Statement.Error)
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	res := []*monsters.FetchMonsterRankingDto{}
	for _, m := range monster {
		if IsPreloadNotFound(&m) {
			continue
		}
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

	if len(res) == 0 || res == nil {
		return nil, gorm.ErrRecordNotFound
	}
	return res, err
}

func IsPreloadNotFound(monsters *mysql.Monster) bool {
	fmt.Println(monsters.Tribe)
	if len(monsters.Weakness) == 0 || monsters.Tribe == nil || len(monsters.Field) == 0 || len(monsters.Product) == 0 {
		return true
	}

	return false
}
