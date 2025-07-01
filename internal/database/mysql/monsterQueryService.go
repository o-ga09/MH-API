package mysql

import (
	"context"
	"fmt"

	param "mh-api/internal/controller/monster"
	"mh-api/internal/domain/music"
	"mh-api/internal/service/monsters"

	"strings"

	"gorm.io/gorm"
)

type monsterQueryService struct{}

func NewmonsterQueryService() *monsterQueryService {
	return &monsterQueryService{}
}

// MonsterToDTO は、Monster構造体からFetchMonsterListDtoに変換する関数です
func MonsterToDTO(m Monster) *monsters.FetchMonsterListDto {
	var locations []string
	var titles []string
	var weak_attack []monsters.Weakness_attack
	var weak_element []monsters.Weakness_element
	var firstWeaknessAttack, secondWeaknessAttack, firstWeaknesselement, secondWeaknessElement string
	var ranking []monsters.Ranking
	var bgm []music.Music

	for _, f := range m.Field {
		locations = append(locations, f.Name)
	}

	for _, p := range m.Product {
		titles = append(titles, p.Name)
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
		bgm = append(bgm, *music.NewMusic(b.MusicId, b.MonsterId, b.Name, b.Url))
	}

	if len(m.Weakness) > 0 {
		firstWeaknessAttack = m.Weakness[0].FirstWeakAttack
		secondWeaknessAttack = m.Weakness[0].SecondWeakAttack
		firstWeaknesselement = m.Weakness[0].FirstWeakElement
		secondWeaknessElement = m.Weakness[0].SecondWeakElement
	}

	return &monsters.FetchMonsterListDto{
		Id:                 m.MonsterId,
		Name:               m.Name,
		Description:        m.Description,
		AnotherName:        m.AnotherName,
		NameEn:             m.NameEn,
		Location:           locations,
		Category:           m.Tribe.Name_ja,
		Title:              titles,
		FirstWeak_Attack:   firstWeaknessAttack,
		SecondWeak_Attack:  secondWeaknessAttack,
		FirstWeak_Element:  firstWeaknesselement,
		SecondWeak_Element: secondWeaknessElement,
		Weakness_attack:    weak_attack,
		Weakness_element:   weak_element,
		Ranking:            ranking,
		BGM:                bgm,
		Element:            m.Element,
	}
}

func (s *monsterQueryService) FetchList(ctx context.Context, id string) ([]*monsters.FetchMonsterListDto, error) {
	var monster []Monster
	var monsterIds []string
	var result *gorm.DB
	var p param.RequestParam
	var err error

	where_clade := ""
	whereArgs := []interface{}{}
	sort := ""

	if id == "" {
		p = ctx.Value("param").(param.RequestParam)
	}

	limit := p.Limit
	offset := p.Offset

	if p.MonsterIds != "" {
		monsterIds = strings.Split(p.MonsterIds, ",")
		where_clade = "monster_id IN (?)"
		whereArgs = append(whereArgs, monsterIds)
	}

	if p.MonsterName != "" {
		if where_clade != "" {
			where_clade += " and name LIKE ?"
		} else {
			where_clade = "name LIKE ?"
		}
		whereArgs = append(whereArgs, "%"+p.MonsterName+"%")
	}

	// Add usage element search
	if p.UsageElement != "" {
		if where_clade != "" {
			where_clade += " and element = ?"
		} else {
			where_clade = "element = ?"
		}
		whereArgs = append(whereArgs, p.UsageElement)
	}

	// Add weakness element search - need to join with Weakness table
	if p.WeaknessElement != "" {
		weaknessJoin := " EXISTS (SELECT 1 FROM weakness w WHERE w.monster_id = monster.monster_id AND " +
			"(w.first_weak_element = ? OR " +
			"w.second_weak_element = ? OR " +
			"w.fire = ? OR " +
			"w.water = ? OR " +
			"w.lightning = ? OR " +
			"w.ice = ? OR " +
			"w.dragon = ?))"
		
		if where_clade != "" {
			where_clade += " and " + weaknessJoin
		} else {
			where_clade = weaknessJoin
		}
		// Add the same value 7 times for the 7 placeholders
		for i := 0; i < 7; i++ {
			whereArgs = append(whereArgs, p.WeaknessElement)
		}
	}

	if p.Sort == "1" {
		sort = "CAST(monster_id AS UNSIGNED) DESC"
	} else {
		sort = "CAST(monster_id AS UNSIGNED) ASC"
	}

	if id != "" {
		result = CtxFromDB(ctx).WithContext(ctx).Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Where("monster_id = ? ", id).Find(&monster)
	} else if where_clade != "" {
		result = CtxFromDB(ctx).WithContext(ctx).Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Where(where_clade, whereArgs...).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else {
		result = CtxFromDB(ctx).WithContext(ctx).Model(&monster).Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Limit(limit).Offset(offset).Order(sort).Find(&monster)
	}

	if result.Error != nil {
		return nil, err
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	res := []*monsters.FetchMonsterListDto{}
	for _, m := range monster {
		dto := MonsterToDTO(m)
		res = append(res, dto)
	}

	return res, err
}

func IsPreloadNotFound(monsters *Monster) bool {
	fmt.Println(monsters.Tribe)
	if len(monsters.Weakness) == 0 || monsters.Tribe == nil || len(monsters.Field) == 0 || len(monsters.Product) == 0 {
		return true
	}

	return false
}
