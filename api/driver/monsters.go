package driver

import (
	"context"
	"mh-api/api/middleware"

	"log/slog"

	"gorm.io/gorm"
)

type MonsterDriver interface {
	GetAll() ([]Monster)
	GetById(id int) (Monster)
	Create(MonsterJson) error
	Update(id int, monsterJson MonsterJson) error
	Delete(id int) error
}

type MonsterDriverimpl struct {
	conn *gorm.DB
}

func (d MonsterDriverimpl) GetAll() []Monster {
	monster := []Monster{}
	d.conn.Find(&monster)
	return monster
}

func (d MonsterDriverimpl) GetById(id int) Monster {
	monster := Monster{}
	err := d.conn.First(&monster,id).Error
	if err != nil {
		slog.Log(context.Background(),middleware.SeverityError,"Driver Error","error",err)
	}
	return monster
}

func (d MonsterDriverimpl) Create(driverJson MonsterJson) error {
	err := d.conn.Create(&driverJson)
	if err != nil {
		slog.Log(context.Background(),middleware.SeverityError,"Driver Error","error",err.Error)
	}
	return err.Error
}

func (d MonsterDriverimpl) Update(id int, driverJson MonsterJson) error {
	err := d.conn.Model(&Monster{}).Where("id = ?",id).Updates(&driverJson)
	if err != nil {
		slog.Log(context.Background(),middleware.SeverityError,"Driver Error","error",err.Error)
	}
	return err.Error
}

func (d MonsterDriverimpl) Delete(id int) error {
	err := d.conn.Delete(&Monster{},id)
	if err != nil {
		slog.Log(context.Background(),middleware.SeverityError,"Driver Error","error",err.Error)
	}
	return err.Error
}

func ProvideMonsterDriver(conn *gorm.DB) MonsterDriver {
	return &MonsterDriverimpl{conn: conn}
}

type Monster struct {
	Id               int    `db:"id" json:"id,omitempty"`
	Name             string `db:"name" json:"name,omitempty"`
	Desc             string `db:"desc" json:"desc,omitempty"`
	Location         string `db:"location" json:"location,omitempty"`
	Specify          string `db:"specify" json:"specify,omitempty"`
	Weakness_attack  string `db:"weakness_attack" json:"weakness___attack,omitempty"`
	Weakness_element string `db:"weakness_element" json:"weakness___element,omitempty"`
}

type MonsterJson struct {
	Name             string `db:"name" json:"name,omitempty"`
	Desc             string `db:"desc" json:"desc,omitempty"`
	Location         string `db:"location" json:"location,omitempty"`
	Specify          string `db:"specify" json:"specify,omitempty"`
	Weakness_attack  string `db:"weakness_attack" json:"weakness___attack,omitempty"`
	Weakness_element string `db:"weakness_element" json:"weakness___element,omitempty"`
}

func (MonsterJson) TableName() string {
	return "monster"
}