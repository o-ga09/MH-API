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

func (s *monsterQueryService) FetchList(ctx context.Context, id string) (*monsters.FetchMonsterListResult, error) {
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
		sort = "CAST(monster_id AS UNSIGNED) DESC"
	} else {
		sort = "CAST(monster_id AS UNSIGNED) ASC"
	}

	// 総件数を取得
	var totalCount int64
	db := CtxFromDB(ctx).WithContext(ctx).Model(&Monster{})
	
	if id != "" {
		result = db.Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Where("monster_id = ? ", id).Find(&monster)
		totalCount = result.RowsAffected
	} else if where_clade != "" && p.MonsterIds != "" {
		// カウントクエリ
		db.Where(where_clade, monsterIds).Count(&totalCount)
		// データ取得クエリ
		result = db.Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Where(where_clade, monsterIds).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else if where_clade != "" {
		// カウントクエリ
		db.Where(where_clade).Count(&totalCount)
		// データ取得クエリ
		result = db.Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Where(where_clade).Limit(limit).Offset(offset).Order(sort).Find(&monster)
	} else {
		// カウントクエリ
		db.Count(&totalCount)
		// データ取得クエリ
		result = db.Preload("Weakness").Preload("Field").Preload("Tribe").Preload("Product").Preload("Ranking").Preload("BGM").Limit(limit).Offset(offset).Order(sort).Find(&monster)
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

	return &monsters.FetchMonsterListResult{
		Monsters: res,
		Total:    int(totalCount),
	}, err
}

func IsPreloadNotFound(monsters *Monster) bool {
	fmt.Println(monsters.Tribe)
	if len(monsters.Weakness) == 0 || monsters.Tribe == nil || len(monsters.Field) == 0 || len(monsters.Product) == 0 {
		return true
	}

	return false
}
