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
	gormDB := CtxFromDB(ctx)

	if params.WeaponID != nil {
		gormDB = gormDB.Where("weapon_id = ?", *params.WeaponID)
	}

	if params.Name != nil {
		gormDB = gormDB.Where("name LIKE ?", "%"+*params.Name+"%")
	}

	if params.Sort != nil {
		switch *params.Sort {
		case "asc":
			gormDB = gormDB.Order("id ASC")
		case "desc":
			gormDB = gormDB.Order("id DESC")
		}
	}

	if params.Order != nil {
		switch *params.Order {
		case 1:
			gormDB = gormDB.Order("id ASC")
		case 2:
			gormDB = gormDB.Order("id DESC")
		}
	}
	if params.Limit != nil && *params.Limit > 0 {
		gormDB = gormDB.Limit(*params.Limit)
	} else {
		defaultLimit := 20
		gormDB = gormDB.Limit(defaultLimit)
	}

	if params.Offset != nil {
		gormDB = gormDB.Offset(*params.Offset)
	}

	result := gormDB.Find(&weaponsList)
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
