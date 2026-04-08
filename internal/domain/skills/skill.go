package skills

import "gorm.io/gorm"

type Skills []*Skill

type Skill struct {
	gorm.Model
	SkillId     string       `gorm:"column:skill_id;primaryKey;type:varchar(10);not null;index"`
	Name        string       `gorm:"column:name;type:varchar(255);not null"`
	Description string       `gorm:"column:description;type:varchar(500)"`
	Levels      []SkillLevel `gorm:"foreignKey:skill_id;references:skill_id"`
}

type SkillLevel struct {
	gorm.Model
	SkillLevelId string `gorm:"column:skill_level_id;primaryKey;type:varchar(10);not null;index"`
	SkillId      string `gorm:"column:skill_id;type:varchar(10);not null"`
	Level        int    `gorm:"column:level;type:int;not null"`
	Description  string `gorm:"column:description;type:varchar(500);not null"`
}
