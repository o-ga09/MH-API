package weapons

import (
	"context"
	"fmt"
	"mh-api/app/internal/domain/weapons"
	"mh-api/app/internal/driver/mysql"
)

type weaponRepository struct {
}

func NewweaponRepository() *weaponRepository {
	return &weaponRepository{}
}

func (r *weaponRepository) Save(ctx context.Context, w weapons.Weapon) error {
	weapon := mysql.Weapon{
		WeaponId:      w.GetID(),
		Name:          w.GetName(),
		ImageUrl:      w.GetURL(),
		Rare:          w.GetRERATY(),
		Attack:        w.GetAttack(),
		ElemantAttaxk: w.GetElementAttack(),
		Critical:      w.GetCritical(),
		Shapness:      w.GetShapness(),
		Description:   w.GetDescription(),
	}
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 0")
	err := mysql.CtxFromDB(ctx).Save(&weapon).Error
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 1")
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (r *weaponRepository) Remove(ctx context.Context, weaponId string) error {
	weapon := mysql.Weapon{
		WeaponId: weaponId,
	}
	err := mysql.CtxFromDB(ctx).Delete(&weapon).Error
	if err != nil {
		return err
	}
	return nil
}
