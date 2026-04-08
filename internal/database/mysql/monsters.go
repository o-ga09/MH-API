package mysql

import (
	"context"
	"strings"

	"mh-api/internal/domain/monsters"
	"mh-api/pkg/element"
)

type monsterRepository struct{}

func NewMonsterRepository() monsters.Repository {
	return &monsterRepository{}
}

func (r *monsterRepository) FindAll(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
	var monsterList []*monsters.Monster
	var whereClauses []string
	var whereArgs []interface{}

	if params.MonsterIds != "" {
		monsterIds := strings.Split(params.MonsterIds, ",")
		whereClauses = append(whereClauses, "monster_id IN (?)")
		whereArgs = append(whereArgs, monsterIds)
	}

	if params.MonsterName != "" {
		whereClauses = append(whereClauses, "name LIKE ?")
		whereArgs = append(whereArgs, "%"+params.MonsterName+"%")
	}

	if params.UsageElement != "" {
		whereClauses = append(whereClauses, "element = ?")
		whereArgs = append(whereArgs, element.NormalizeToJapanese(params.UsageElement))
	}

	if params.WeaknessElement != "" {
		weaknessJoin := "EXISTS (SELECT 1 FROM weakness w WHERE w.monster_id = monster.monster_id AND " +
			"(w.first_weak_element = ? OR " +
			"w.second_weak_element = ? OR " +
			"w.fire = ? OR " +
			"w.water = ? OR " +
			"w.lightning = ? OR " +
			"w.ice = ? OR " +
			"w.dragon = ?))"
		whereClauses = append(whereClauses, weaknessJoin)
		for i := 0; i < 7; i++ {
			whereArgs = append(whereArgs, params.WeaknessElement)
		}
	}

	limit := params.Limit
	if limit <= 0 {
		limit = 100
	}

	sort := "CAST(monster_id AS UNSIGNED) ASC"
	if params.Sort == "desc" {
		sort = "CAST(monster_id AS UNSIGNED) DESC"
	}

	query := CtxFromDB(ctx).
		Preload("Weakness").
		Preload("Field").
		Preload("Tribe").
		Preload("Product").
		Preload("Ranking").
		Preload("BGM")

	whereStr := strings.Join(whereClauses, " AND ")
	if whereStr != "" {
		query = query.Where(whereStr, whereArgs...)
	}

	result := query.Limit(limit).Offset(params.Offset).Order(sort).Find(&monsterList)
	if result.Error != nil {
		return nil, result.Error
	}

	var total int64
	countQuery := CtxFromDB(ctx).Model(&monsters.Monster{})
	if whereStr != "" {
		countQuery = countQuery.Where(whereStr, whereArgs...)
	}
	countQuery.Count(&total)

	return &monsters.SearchResult{
		Monsters: monsterList,
		Total:    int(total),
	}, nil
}

func (r *monsterRepository) FindById(ctx context.Context, id string) (*monsters.Monster, error) {
	var monster monsters.Monster
	result := CtxFromDB(ctx).
		Preload("Weakness").
		Preload("Field").
		Preload("Tribe").
		Preload("Product").
		Preload("Ranking").
		Preload("BGM").
		Where("monster_id = ?", id).
		First(&monster)

	if result.Error != nil {
		return nil, result.Error
	}

	return &monster, nil
}
