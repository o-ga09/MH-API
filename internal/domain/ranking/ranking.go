package ranking

import "gorm.io/gorm"

type Rankings []*Ranking

type Ranking struct {
	gorm.Model
	MonsterId string `gorm:"column:monster_id;type:varchar(10); not null"`
	Ranking   string `gorm:"column:ranking;type:varchar(10)"`
	VoteYear  string `gorm:"column:vote_year;type:varchar(20)"`
}
