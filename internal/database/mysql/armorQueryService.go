package mysql

import (
	"context"
	"errors"
	"mh-api/internal/domain/armors"

	"gorm.io/gorm"
)

type armorRepository struct{}

func NewArmorRepository() armors.Repository {
	return &armorRepository{}
}

func (r *armorRepository) Find(ctx context.Context, params armors.SearchParams) (*armors.SearchResult, error) {
	var armorsList []*armors.Armor
	query := CtxFromDB(ctx).Preload("Skills").Preload("RequiredItems")

	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if params.Slot != "" {
		query = query.Where("slot = ?", params.Slot)
	}
	if params.SkillName != "" {
		query = query.Where("EXISTS (SELECT 1 FROM armor_skill ars WHERE ars.armor_id = armor.armor_id AND ars.skill_name LIKE ?)", "%"+params.SkillName+"%")
	}

	limit := params.Limit
	if limit <= 0 {
		limit = 100
	}

	sort := "armor_id ASC"
	if params.Sort == "desc" {
		sort = "armor_id DESC"
	}

	var total int64
	countQuery := CtxFromDB(ctx).Model(&armors.Armor{})
	if params.Name != "" {
		countQuery = countQuery.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if params.Slot != "" {
		countQuery = countQuery.Where("slot = ?", params.Slot)
	}
	if params.SkillName != "" {
		countQuery = countQuery.Where("EXISTS (SELECT 1 FROM armor_skill ars WHERE ars.armor_id = armor.armor_id AND ars.skill_name LIKE ?)", "%"+params.SkillName+"%")
	}
	countQuery.Count(&total)

	result := query.Limit(limit).Offset(params.Offset).Order(sort).Find(&armorsList)
	if result.Error != nil {
		return nil, result.Error
	}
	return &armors.SearchResult{Armors: armorsList, Total: int(total)}, nil
}

func (r *armorRepository) GetAll(ctx context.Context) (armors.Armors, error) {
	var armorsList []*armors.Armor
	result := CtxFromDB(ctx).Preload("Skills").Preload("RequiredItems").Find(&armorsList)
	if result.Error != nil {
		return nil, result.Error
	}
	return armorsList, nil
}

func (r *armorRepository) GetByID(ctx context.Context, armorId string) (*armors.Armor, error) {
	var armor armors.Armor
	result := CtxFromDB(ctx).Preload("Skills").Preload("RequiredItems").Where("armor_id = ?", armorId).First(&armor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &armor, nil
}
