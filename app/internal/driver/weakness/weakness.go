package weakness

import (
	"context"
	"mh-api/app/internal/domain/weakness"
	"mh-api/app/internal/driver/mysql"
	"strconv"

	"gorm.io/gorm"
)

type weakRepository struct {
	conn *gorm.DB
}

func NewweakRepository(conn *gorm.DB) *weakRepository {
	return &weakRepository{
		conn: conn,
	}
}

func (r *weakRepository) Get(ctx context.Context, monsterId string) (weakness.Weaknesses, error) {
	weak := []mysql.Weakness{}
	err := r.conn.Find(&weak).Error
	if err != nil {
		return nil, err
	}

	res := weakness.Weaknesses{}
	for _, r := range weak {
		res = append(res, *weakness.NewWeakness(
			r.MonsterId,
			r.PartId,
			r.Fire,
			r.Water,
			r.Lightning,
			r.Ice,
			r.Dragon,
			r.Slashing,
			r.Blow,
			r.Bullet,
			r.FirstWeakAttack,
			r.SecondWeakAttack,
			r.FirstWeakElement,
			r.SecondWeakElement,
		))
	}

	return res, nil
}

func (r *weakRepository) Save(ctx context.Context, w weakness.Weakness) error {
	weak := mysql.Weakness{
		MonsterId:         w.GetMonsterID(),
		PartId:            w.GetPartID(),
		Fire:              w.GetFire(),
		Water:             w.GetWater(),
		Lightning:         w.GetLightning(),
		Ice:               w.GetIce(),
		Dragon:            w.GetDragon(),
		Slashing:          w.GetSlashing(),
		Blow:              w.GetBlow(),
		Bullet:            w.GetBullet(),
		FirstWeakAttack:   w.GetFirstWeakAttack(),
		SecondWeakAttack:  w.GetSecondWeakAttack(),
		FirstWeakElement:  w.GetFirstWeakElement(),
		SecondWeakElement: w.GetSecondWeakElement(),
	}
	r.conn.Exec("SET foreign_key_checks = 0")
	err := r.conn.Save(&weak).Error
	r.conn.Exec("SET foreign_key_checks = 1")
	if err != nil {
		return err
	}
	return nil
}

func (r *weakRepository) Remove(ctx context.Context, Id string) error {
	i, _ := strconv.Atoi(Id)
	weak := mysql.Weakness{
		Model: gorm.Model{ID: uint(i)},
	}
	err := r.conn.Delete(&weak).Error
	if err != nil {
		return err
	}
	return nil
}
