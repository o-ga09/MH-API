package driver

import (
	"context"
	"fmt"

	"mh-api/api/gateway/repository"
	"mh-api/api/middleware"

	"log/slog"

	"gorm.io/gorm"
)

type MonsterDriverimpl struct {
	conn *gorm.DB
}

func (d MonsterDriverimpl) GetAll() []repository.Monster {
	monster := []repository.Monster{}
	d.conn.Find(&monster)
	return monster
}

func (d MonsterDriverimpl) GetById(id int) repository.Monster {
	monster := repository.Monster{}
	err := d.conn.Where("`monster_id` = ?", id).First(&monster).Error
	if err != nil {
		slog.Log(context.Background(), middleware.SeverityError, "Driver Error", "error", err)
	}
	return monster
}

func (d MonsterDriverimpl) Create(driverJson repository.MonsterJson) error {
	err := d.conn.Create(&driverJson).Error
	if err != nil {
		slog.Log(context.Background(), middleware.SeverityError, "Driver Error", "error", err)
		return fmt.Errorf(" Record Create Error : %v", err)
	}
	return nil
}

func (d MonsterDriverimpl) Update(id int, driverJson repository.MonsterJson) error {
	err := d.conn.Model(&repository.Monster{}).Where("id = ?", id).Updates(&driverJson).Error
	if err != nil {
		slog.Log(context.Background(), middleware.SeverityError, "Driver Error", "error", err)
		return fmt.Errorf(" Record Update Error : %v", err)
	}
	return nil
}

func (d MonsterDriverimpl) Delete(id int) error {
	err := d.conn.Delete(&repository.Monster{}, id).Error
	if err != nil {
		slog.Log(context.Background(), middleware.SeverityError, "Driver Error", "error", err)
		return fmt.Errorf(" Record Delete Error : %v", err)
	}
	return nil
}

func ProvideMonsterDriver(conn *gorm.DB) repository.MonsterDriver {
	return &MonsterDriverimpl{conn: conn}
}
