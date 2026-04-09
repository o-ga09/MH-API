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

// SearchParams はスキル検索の条件を表す
type SearchParams struct {
	Name        string
	Description string
	Limit       int
	Offset      int
}

// SearchResult はスキル検索の結果を表す
type SearchResult struct {
	Skills Skills
	Total  int
}
