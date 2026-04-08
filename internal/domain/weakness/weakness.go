package weakness

import "gorm.io/gorm"

type Weaknesses []*Weakness

type Weakness struct {
	gorm.Model
	MonsterId         string `gorm:"column:monster_id;type:varchar(10);not null"`
	PartId            string `gorm:"column:aprt_id;type:varchar(255);not null"`
	Fire              string `gorm:"column:fire;type:varchar(255)"`
	Water             string `gorm:"column:water;type:varchar(255)"`
	Lightning         string `gorm:"column:lightning;type:varchar(255)"`
	Ice               string `gorm:"column:ice;type:varchar(255)"`
	Dragon            string `gorm:"column:dragon;type:varchar(255)"`
	Slashing          string `gorm:"column:slashing;type:varchar(255)"`
	Blow              string `gorm:"column:blow;type:varchar(255)"`
	Bullet            string `gorm:"column:bullet;type:varchar(255)"`
	FirstWeakAttack   string `gorm:"column:first_weak_attack;type:varchar(255)"`
	SecondWeakAttack  string `gorm:"column:second_weak_attack;type:varchar(255)"`
	FirstWeakElement  string `gorm:"column:first_weak_element;type:varchar(255)"`
	SecondWeakElement string `gorm:"column:second_weak_element;type:varchar(255)"`
}
