package music

import "gorm.io/gorm"

type Musics []*Music

type Music struct {
	gorm.Model
	MusicId   string `gorm:"column:music_id;primaryKey;type:varchar(10);not null;index"`
	MonsterId string `gorm:"column:monster_id;type:varchar(10);not null"`
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	Url       string `gorm:"column:url;type:varchar(255)"`
}
