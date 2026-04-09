package armors

import "gorm.io/gorm"

type Armors []*Armor

type Armor struct {
	gorm.Model
	ArmorId             string               `gorm:"column:armor_id;type:varchar(255);not null"`
	Name                string               `gorm:"column:name;type:varchar(255);not null"`
	Slot                string               `gorm:"column:slot;type:varchar(255)"`
	Defense             int                  `gorm:"column:defense;type:int"`
	FireResistance      int                  `gorm:"column:fire_resistance;type:int"`
	WaterResistance     int                  `gorm:"column:water_resistance;type:int"`
	LightningResistance int                  `gorm:"column:lightning_resistance;type:int"`
	IceResistance       int                  `gorm:"column:ice_resistance;type:int"`
	DragonResistance    int                  `gorm:"column:dragon_resistance;type:int"`
	Skills              []*ArmorSkill        `gorm:"foreignKey:armor_id;references:armor_id"`
	RequiredItems       []*ArmorRequiredItem `gorm:"foreignKey:armor_id;references:armor_id"`
}

type ArmorSkill struct {
	gorm.Model
	ArmorId   string `gorm:"column:armor_id;type:varchar(255);not null"`
	SkillId   string `gorm:"column:skill_id;type:varchar(255);not null"`
	SkillName string `gorm:"column:skill_name;type:varchar(255);not null"`
}

type ArmorRequiredItem struct {
	gorm.Model
	ArmorId  string `gorm:"column:armor_id;type:varchar(255);not null"`
	ItemId   string `gorm:"column:item_id;type:varchar(255);not null"`
	ItemName string `gorm:"column:item_name;type:varchar(255);not null"`
}

// SearchParams は防具検索の条件を表す
type SearchParams struct {
	Name      string
	SkillName string
	Slot      string
	Limit     int
	Offset    int
	Sort      string
}

// SearchResult は防具検索の結果を表す
type SearchResult struct {
	Armors Armors
	Total  int
}
