package mysql

import (
	"context"
	"errors"
	"mh-api/internal/domain/weapons"

	"gorm.io/gorm"
)

type weaponRepository struct{}

func NewWeaponRepository() weapons.Repository {
	return &weaponRepository{}
}

func (r *weaponRepository) Find(ctx context.Context, params weapons.SearchParams) (*weapons.SearchResult, error) {
	var weaponsList []*weapons.Weapon

	query := CtxFromDB(ctx)

	if params.WeaponID != nil {
		query = query.Where("weapon_id = ?", *params.WeaponID)
	}

	if params.Name != nil {
		query = query.Where("name LIKE ?", "%"+*params.Name+"%")
	}

	if params.Sort != nil {
		switch *params.Sort {
		case "asc":
			query = query.Order("id ASC")
		case "desc":
			query = query.Order("id DESC")
		}
	}

	limit := 20
	if params.Limit != nil && *params.Limit > 0 {
		limit = *params.Limit
	}
	query = query.Limit(limit)

	if params.Offset != nil {
		query = query.Offset(*params.Offset)
	}

	result := query.Find(&weaponsList)
	if result.Error != nil {
		return nil, result.Error
	}

	return &weapons.SearchResult{
		Weapons:    weaponsList,
		TotalCount: int(result.RowsAffected),
		Limit:      limit,
	}, nil
}

func (r *weaponRepository) FindByID(ctx context.Context, weaponID string) (*weapons.Weapon, error) {
	var weapon weapons.Weapon
	result := CtxFromDB(ctx).Where("weapon_id = ?", weaponID).First(&weapon)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &weapon, nil
}
