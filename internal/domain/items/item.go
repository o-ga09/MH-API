package items

import "gorm.io/gorm"

type Items []*Item

type Item struct {
	gorm.Model
	ItemId    string `gorm:"column:item_id;primaryKey;type:varchar(10);not null"`
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	ImageUrl  string `gorm:"column:image_url;type:varchar(255)"`
	MonsterId string `gorm:"column:monster_id;type:varchar(10)"`
}
