package Tribes

import "gorm.io/gorm"

type Tribes []*Tribe

type Tribe struct {
	gorm.Model
	TribeId     string `gorm:"column:tribe_id;primaryKey;type:varchar(10);not null"`
	MonsterId   string `gorm:"column:monster_id;type:varchar(10);not null"`
	Name_ja     string `gorm:"column:name_ja;type:varchar(255);not null"`
	Name_en     string `gorm:"column:name_en;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:varchar(255)"`
}
