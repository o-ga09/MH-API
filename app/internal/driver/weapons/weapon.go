package weapons

import (
	"context"
	"fmt"
	"mh-api/app/internal/domain/weapons"
	"mh-api/app/internal/driver/mysql"

	"gorm.io/gorm"
)

type weaponRepository struct {
	conn *gorm.DB
}

func NewweaponRepository(conn *gorm.DB) *weaponRepository {
	return &weaponRepository{
		conn: conn,
	}
}

func (r *weaponRepository) Get(ctx context.Context, monsterId string) (weapons.Weapons, error) {
	weapon := []mysql.Weapon{}
	err := r.conn.Find(&weapon).Error
	if err != nil {
		return nil, err
	}

	res := weapons.Weapons{}
	for _, w := range weapon {
		res = append(res, *weapons.NewWeapon(
			w.WeaponId,
			w.Name,
			w.ImageUrl,
			w.Rare,
			w.Attack,
			w.ElemantAttaxk,
			w.Shapness,
			w.Critical,
			w.Description,
		))
	}

	return res, nil
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
	err := r.conn.Save(&weapon).Error
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (r *weaponRepository) Remove(ctx context.Context, weaponId string) error {
	weapon := mysql.Weapon{
		WeaponId: weaponId,
	}
	err := r.conn.Delete(&weapon).Error
	if err != nil {
		return err
	}
	return nil
}
