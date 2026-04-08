package fields

import "gorm.io/gorm"

type Fields []*Field

type Field struct {
	gorm.Model
	FieldId   string `gorm:"column:field_id;primaryKey;type:varchar(10);not null"`
	MonsterId string `gorm:"column:monster_id;type:varchar(10);not null"`
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	ImageUrl  string `gorm:"column:image_url;type:varchar(255)"`
}
