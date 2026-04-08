package monsters

import (
	"gorm.io/gorm"
	"mh-api/internal/domain/fields"
	"mh-api/internal/domain/music"
	Products "mh-api/internal/domain/products"
	"mh-api/internal/domain/ranking"
	Tribes "mh-api/internal/domain/tribes"
	"mh-api/internal/domain/weakness"
)

type Monsters []*Monster

type Monster struct {
	gorm.Model
	MonsterId   string               `gorm:"column:monster_id;primaryKey;type:varchar(10);not null;index"`
	Name        string               `gorm:"column:name;type:varchar(255)"`
	Description string               `gorm:"column:description;type:varchar(255)"`
	Element     *string              `gorm:"column:element;type:varchar(255)"`
	AnotherName string               `gorm:"column:another_name;type:varchar(255)"`
	NameEn      string               `gorm:"column:name_en;type:varchar(255)"`
	Weakness    []*weakness.Weakness  `gorm:"foreignKey:monster_id;references:monster_id"`
	Tribe       *Tribes.Tribe         `gorm:"foreignKey:monster_id;references:monster_id"`
	Product     []*Products.Product   `gorm:"foreignKey:monster_id;references:monster_id"`
	Field       []*fields.Field       `gorm:"foreignKey:monster_id;references:monster_id"`
	Ranking     []*ranking.Ranking    `gorm:"foreignKey:monster_id;references:monster_id"`
	BGM         []*music.Music        `gorm:"foreignKey:monster_id;references:monster_id"`
}
