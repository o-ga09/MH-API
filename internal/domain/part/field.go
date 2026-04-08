package part

import "gorm.io/gorm"

type Parts []*Part

type Part struct {
	gorm.Model
	PartId      string `gorm:"column:part_id;primaryKey;type:varchar(10);not null"`
	MonsterId   string `gorm:"column:monster_id;type:varchar(10);not null"`
	Description string `gorm:"column:decription;type:varchar(255)"`
}
