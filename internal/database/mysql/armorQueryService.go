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
