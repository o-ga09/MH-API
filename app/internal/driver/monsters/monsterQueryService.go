package monsters

import (
	"context"
	"fmt"
	param "mh-api/app/internal/controller/monster"
	"mh-api/app/internal/domain/music"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/monsters"
	"strings"

	"gorm.io/gorm"
)

type monsterQueryService struct{}

func NewmonsterQueryService() *monsterQueryService {
	return &monsterQueryService{}
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
		sort = "CAST(monster_id AS UNSIGNED) DESC"
	} else {
		sort = "CAST(monster_id AS UNSIGNED) ASC"
	}

	if id != "" {
		result = mysql.CtxFromDB(ctx).WithContext(ctx).Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Where("monster_id = ? ", id).Find(&monster)
	} else if where_clade != "" && p.MonsterIds != "" {
		result = mysql.CtxFromDB(ctx).WithContext(ctx).Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Where(where_clade, monsterIds).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else if where_clade != "" {
		result = mysql.CtxFromDB(ctx).WithContext(ctx).Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Where(where_clade).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else {
		result = mysql.CtxFromDB(ctx).WithContext(ctx).Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Limit(limit).Offset(offset).Order(sort).Find(&monster)
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
		var firstWeaknessAttack, secondWeaknessAttack, firstWeaknesselement, secondWeaknessElement string
		var ranking []monsters.Ranking
		var bgm []music.Music

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

		for _, r := range m.Ranking {
			ranking = append(ranking, monsters.Ranking{
				Ranking:  r.Ranking,
				VoteYear: r.VoteYear,
			})
		}

		for _, b := range m.BGM {
			bgm = append(bgm, *music.NewMusic(b.MonsterId, b.MonsterId, b.Name, b.Url))
		}

		if len(m.Weakness) > 0 {
			firstWeaknessAttack = m.Weakness[0].FirstWeakAttack
			secondWeaknessAttack = m.Weakness[0].SecondWeakAttack
			firstWeaknesselement = m.Weakness[0].FirstWeakElement
			secondWeaknessElement = m.Weakness[0].SecondWeakElement
		}

		r := monsters.FetchMonsterListDto{
			Id:                 m.MonsterId,
			Name:               m.Name,
			Description:        m.Description,
			AnotherName:        m.AnotherName,
			NameEn:             m.NameEn,
			Location:           l,
			Category:           m.Tribe.Name_ja,
			Title:              t,
			FirstWeak_Attack:   firstWeaknessAttack,
			SecondWeak_Attack:  secondWeaknessAttack,
			FirstWeak_Element:  firstWeaknesselement,
			SecondWeak_Element: secondWeaknessElement,
			Weakness_attack:    weak_attack,
			Weakness_element:   weak_element,
			Ranking:            ranking,
			BGM:                bgm,
		}
		res = append(res, &r)
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
