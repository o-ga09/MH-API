package queryservice

import (
	"context"
	"fmt"
	"mh-api/app/internal/driver/schema"
	"mh-api/app/internal/service/monsters"

	"gorm.io/gorm"
)

type MonsterQueryService struct {
	qs *gorm.DB
}

func NewMonsterQueryService(qs *gorm.DB) MonsterQueryService {
	return MonsterQueryService{qs: qs}
}

func (s MonsterQueryService) FetchList(ctx context.Context, id string) ([]*monsters.FetchMonsterListDto, error) {
	var result *gorm.DB
	monster := []struct {
		MonsterID           string
		MonsterName         string
		MonsterDescription  string
		FieldName           string
		TribeName           string
		ProductName         string
		PartId              string
		Fire                string
		Water               string
		Lightning           string
		Ice                 string
		Dragon              string
		Slashing            string
		Blow                string
		Bullet              string
		First_weak_attack   string
		Second_weak_attack  string
		First_weak_element  string
		Second_weak_element string
	}{}

	var part []schema.Part
	result = s.qs.Find(&part)
	if result.Error != nil {
		return nil, result.Error
	}

	if id == "" {
		result = s.qs.Table("monster").
			Select("monster.monster_id, monster.name as monster_name, monster.description as monster_description, tribe.name_ja as tribe_name, product.name as product_name, field.name as field_name, weakness.fire, weakness.water, weakness.ice, weakness.lightning, weakness.dragon, weakness.slashing, weakness.blow, weakness.bullet, weakness.part_id").
			Joins("JOIN tribe ON monster.monster_id = tribe.monster_id").
			Joins("JOIN product ON monster.monster_id = product.monster_id").
			Joins("JOIN field ON monster.monster_id = field.monster_id").
			Joins("JOIN weakness ON monster.monster_id = weakness.monster_id").
			Scan(&monster)

	} else {
		result = s.qs.Table("monster").
			Select("monster.monster_id, monster.name as monster_name, monster.description as monster_description, tribe.name_ja as tribe_name, product.name as product_name, field.name as field_name, weakness.fire, weakness.water, weakness.ice, weakness.lightning, weakness.dragon, weakness.slashing, weakness.blow, weakness.bullet, weakness.part_id").
			Joins("JOIN tribe ON monster.monster_id = tribe.monster_id").
			Joins("JOIN product ON monster.monster_id = product.monster_id").
			Joins("JOIN field ON monster.monster_id = field.monster_id").
			Joins("JOIN weakness ON monster.monster_id = weakness.monster_id").
			Where("monster.monster_id = ?", id).
			Scan(&monster)
	}

	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, fmt.Errorf("NO DATA")
	}

	var res []*monsters.FetchMonsterListDto

	fields := make(map[string][]string)
	products := make(map[string][]string)
	weak_attack := make(map[string][]monsters.Weakness_attack)
	weak_element := make(map[string][]monsters.Weakness_element)

	for i := 0; i < len(monster); i++ {
		partName := ""
		for j := 0; j < len(part); j++ {
			if monster[i].PartId == part[j].PartId {
				partName = part[j].Name
				break
			}
		}
		fields[monster[i].MonsterID] = append(fields[monster[i].MonsterID], monster[i].FieldName)
		products[monster[i].MonsterID] = append(products[monster[i].MonsterID], monster[i].ProductName)
		weak_attack[monster[i].MonsterID] = append(weak_attack[monster[i].MonsterID], monsters.Weakness_attack{
			PartId:   monster[i].PartId,
			PartName: partName,
			Slashing: monster[i].Slashing,
			Blow:     monster[i].Blow,
			Bullet:   monster[i].Bullet,
		})
		weak_element[monster[i].MonsterID] = append(weak_element[monster[i].MonsterID], monsters.Weakness_element{
			PartId:   monster[i].PartId,
			PartName: partName,
			Fire:     monster[i].Fire,
			Water:    monster[i].Water,
			Thunder:  monster[i].Lightning,
			Ice:      monster[i].Ice,
			Dragon:   monster[i].Dragon,
		})
	}

	for i, m := range monster {
		if i != 0 && m.MonsterID == monster[i-1].MonsterID {
			continue
		}
		res = append(res, &monsters.FetchMonsterListDto{
			Id:                 m.MonsterID,
			Name:               m.MonsterName,
			Description:        m.MonsterDescription,
			Category:           m.TribeName,
			Title:              products[m.MonsterID],
			Location:           fields[m.MonsterID],
			FirstWeak_Attack:   m.First_weak_attack,
			FirstWeak_Element:  m.First_weak_element,
			SecondWeak_Attack:  m.Second_weak_attack,
			SecondWeak_Element: m.Second_weak_element,
			Weakness_attack:    weak_attack[m.MonsterID],
			Weakness_element:   weak_element[m.MonsterID],
		})
	}

	return res, nil
}

func (s MonsterQueryService) FetchRank(ctx context.Context) ([]*monsters.FetchMonsterRankingDto, error) {
	var res []*monsters.FetchMonsterRankingDto
	monster := []struct {
		MonsterID          string
		MonsterName        string
		MonsterDescription string
		TribeName          string
		ProductName        string
		FieldName          string
		Ranking            string
		VoteYear           string
	}{}

	result := s.qs.Table("monster").
		Select("monster.monster_id, monster.name as monster_name, monster.description as monster_description, tribe.name_ja as tribe_name, product.name as product_name, field.name as field_name, ranking.ranking as ranking, ranking.vote_year as vote_year").
		Joins("JOIN tribe ON monster.monster_id = tribe.monster_id").
		Joins("JOIN product ON monster.monster_id = product.monster_id").
		Joins("JOIN field ON monster.monster_id = field.monster_id").
		Joins("JOIN ranking ON monster.monster_id = ranking.monster_id").
		Scan(&monster)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, fmt.Errorf("NO DATA")
	}

	ranking := make(map[string][]monsters.Ranking)
	products := make(map[string][]string)
	fields := make(map[string][]string)

	for i := 0; i < len(monster); i++ {
		fields[monster[i].MonsterID] = append(fields[monster[i].MonsterID], monster[i].FieldName)
		products[monster[i].MonsterID] = append(products[monster[i].MonsterID], monster[i].ProductName)
		ranking[monster[i].MonsterID] = append(ranking[monster[i].MonsterID], monsters.Ranking{
			Ranking:  monster[i].Ranking,
			VoteYear: monster[i].VoteYear,
		})
	}

	for i, m := range monster {
		if i != 0 && m.MonsterID == monster[i-1].MonsterID {
			continue
		}
		res = append(res, &monsters.FetchMonsterRankingDto{
			Id:          m.MonsterID,
			Name:        m.MonsterName,
			Description: m.MonsterDescription,
			Category:    m.TribeName,
			Title:       products[m.MonsterID],
			Location:    fields[m.MonsterID],
			Ranking:     ranking[m.MonsterID],
		})
	}

	return res, nil
}
