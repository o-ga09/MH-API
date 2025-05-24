package mysql

import (
	"context"
	"errors"
	"mh-api/internal/domain/weapons"
	weaponService "mh-api/internal/service/weapons"

	"gorm.io/gorm"
)

type WeaponQueryService struct{}

func NewWeaponQueryService() *WeaponQueryService {
	return &WeaponQueryService{}
}

func (qs *WeaponQueryService) FindWeapons(ctx context.Context, params weaponService.SearchWeaponsParams) ([]*weapons.Weapon, int, error) {
	var weaponsList []*Weapon
	var total int64

	// コンテキストからDBを取得
	db := CtxFromDB(ctx)

	if params.WeaponID != nil {
		db = db.Where("weapon_id = ?", *params.WeaponID)
	}

	if params.Name != nil {
		db = db.Where("name LIKE ?", "%"+*params.Name+"%")
	}

	if params.Sort != nil {
		if *params.Sort == "asc" {
			db = db.Order("id ASC")
		} else if *params.Sort == "desc" {
			db = db.Order("id DESC")
		}
	}

	if params.Order != nil {
		if *params.Order == 1 {
			db = db.Order("id ASC")
		} else if *params.Order == 2 {
			db = db.Order("id DESC")
		}
	}
	if params.Limit != nil && *params.Limit > 0 {
		db = db.Limit(*params.Limit)
	} else {
		defaultLimit := 20
		db = db.Limit(defaultLimit)
	}

	if params.Offset != nil {
		db = db.Offset(*params.Offset)
	}

	result := db.Find(&weaponsList)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	total = result.RowsAffected
	weaponDomainList := ToWeaponList(weaponsList)

	return weaponDomainList, int(total), nil
}

func (qs *WeaponQueryService) FindWeaponByID(ctx context.Context, weaponID string) (*weapons.Weapon, error) {
	var weapon Weapon
	result := CtxFromDB(ctx).Where("weapon_id = ?", weaponID).First(&weapon)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	res := ToWeapon(&weapon)
	return res, nil
}

func ToWeapon(weapon *Weapon) *weapons.Weapon {
	return weapons.NewWeapon(
		weapon.WeaponID,
		weapon.Name,
		weapon.ImageUrl,
		weapon.Rarerity,
		weapon.Attack,
		weapon.ElementAttack,
		weapon.Shapness,
		weapon.Critical,
		weapon.Description,
	)
}
func ToWeaponList(weaponsList []*Weapon) []*weapons.Weapon {
	weaponDomainList := make([]*weapons.Weapon, len(weaponsList))
	for i, w := range weaponsList {
		weaponDomainList[i] = ToWeapon(w)
	}
	return weaponDomainList
}
