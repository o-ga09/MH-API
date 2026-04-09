package weapons

import "gorm.io/gorm"

type Weapons []*Weapon

type Weapon struct {
	gorm.Model
	WeaponID      string `gorm:"column:weapon_id;primaryKey;type:varchar(255);not null"`
	Name          string `gorm:"column:name;type:varchar(255);not null"`
	ImageUrl      string `gorm:"column:image_url;type:varchar(255)"`
	Rarerity      string `gorm:"column:rarerity;type:varchar(255)"`
	Attack        string `gorm:"column:attack;type:varchar(255)"`
	ElementAttack string `gorm:"column:element_attack;type:varchar(255)"`
	Shapness      string `gorm:"column:shapness;type:varchar(255)"`
	Critical      string `gorm:"column:critical;type:varchar(255)"`
	Description   string `gorm:"column:description;type:varchar(255)"`
}

// SearchParams は武器検索の条件を表す
type SearchParams struct {
	WeaponID      *string
	Name          *string
	Rarity        *string
	ElementAttack *string
	Limit         *int
	Offset        *int
	Sort          *string
	Order         *int
}

// SearchResult は武器検索の結果を表す
type SearchResult struct {
	Weapons    []*Weapon
	TotalCount int
	Offset     int
	Limit      int
}
