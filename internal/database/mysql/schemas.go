package mysql

import "gorm.io/gorm"

type Monster struct {
	gorm.Model
	MonsterId   string      `gorm:"column:monster_id;primaryKey;type:varchar(10);not null;index"`
	Name        string      `gorm:"column:name;type:varchar(255)"`
	Description string      `gorm:"column:description;type:varchar(255)"`
	Element     *string     `gorm:"column:element;type:varchar(255)"`
	AnotherName string      `gorm:"column:another_name;type:varchar(255)"`
	NameEn      string      `gorm:"column:name_en;type:varchar(255)"`
	Weakness    []*Weakness `gorm:"foreignKey:monster_id;references:monster_id"`
	Tribe       *Tribe      `gorm:"foreignKey:monster_id;references:monster_id"`
	Product     []*Product  `gorm:"foreignKey:monster_id;references:monster_id"`
	Field       []*Field    `gorm:"foreignKey:monster_id;references:monster_id"`
	Ranking     []*Ranking  `gorm:"foreignKey:monster_id;references:monster_id"`
	BGM         []*Music    `gorm:"foreignKey:monster_id;references:monster_id"`
}

type Field struct {
	gorm.Model
	FieldId   string `gorm:"column:field_id;primaryKey;type:varchar(10);not null"`
	MonsterId string `gorm:"column:monster_id;type:varchar(10);not null"`
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	ImageUrl  string `gorm:"column:image_url;type:varchar(255)"`
}

type Item struct {
	gorm.Model
	ItemId   string `gorm:"column:item_id;primaryKey;type:varchar(10);not null"`
	Name     string `gorm:"column:name;type:varchar(255);not null"`
	ImageUrl string `gorm:"column:image_url;type:varchar(255)"`
}

type Music struct {
	gorm.Model
	MusicId    string        `gorm:"column:music_id;primaryKey;type:varchar(10);not null;index"`
	MonsterId  string        `gorm:"column:monster_id;type:varchar(10);not null"`
	Name       string        `gorm:"column:name;type:varchar(255);not null"`
	Url        string        `gorm:"column:url;type:varchar(255)"`
	BgmRanking []*BgmRanking `gorm:"foreignKey:music_id;references:music_id"`
}

type Part struct {
	gorm.Model
	PartId      string `gorm:"column:part_id;primaryKey;type:varchar(10);not null"`
	MonsterId   string `gorm:"column:monster_id;type:varchar(10);not null"`
	Description string `gorm:"column:decription;type:varchar(255)"`
}

type Product struct {
	gorm.Model
	ProductId   string `gorm:"column:product_id;primaryKey;type:varchar(255);not null"`
	MonsterId   string `gorm:"column:monster_id;type:varchar(10);not null"`
	Name        string `gorm:"column:name;type:varchar(255);not null"`
	PublishYear string `gorm:"column:publish_year;type:varchar(20)"`
	TotalSales  string `gorm:"column:total_sales;type:varchar(255)"`
}

type Tribe struct {
	gorm.Model
	TribeId     string `gorm:"column:tribe_id;primaryKey;type:varchar(10);not null"`
	MonsterId   string `gorm:"column:monster_id;type:varchar(10);not null"`
	Name_ja     string `gorm:"column:name_ja;type:varchar(255);not null"`
	Name_en     string `gorm:"column:name_en;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:varchar(255)"`
}

type Weakness struct {
	gorm.Model
	MonsterId         string `gorm:"column:monster_id;type:varchar(10);not null"`
	PartId            string `gorm:"column:aprt_id;type:varchar(255);not null"`
	Fire              string `gorm:"column:fire;type:varchar(255)"`
	Water             string `gorm:"column:water;type:varchar(255)"`
	Lightning         string `gorm:"column:lightning;type:varchar(255)"`
	Ice               string `gorm:"column:ice;type:varchar(255)"`
	Dragon            string `gorm:"column:dragon;type:varchar(255)"`
	Slashing          string `gorm:"column:slashing;type:varchar(255)"`
	Blow              string `gorm:"column:blow;type:varchar(255)"`
	Bullet            string `gorm:"column:bullet;type:varchar(255)"`
	FirstWeakAttack   string `gorm:"column:first_weak_attack;type:varchar(255)"`
	SecondWeakAttack  string `gorm:"column:second_weak_attack;type:varchar(255)"`
	FirstWeakElement  string `gorm:"column:first_weak_element;type:varchar(255)"`
	SecondWeakElement string `gorm:"column:second_weak_element;type:varchar(255)"`
}

type Weapon struct {
	gorm.Model
	WeaponId      string `gorm:"column:weapon_id;primaryKey;type:varchar(255);not null"`
	Name          string `gorm:"column:name;type:varchar(255);not null"`
	ImageUrl      string `gorm:"column:image_url;type:varchar(255)"`
	Rare          string `gorm:"column:rarerity;type:varchar(255)"`
	Attack        string `gorm:"column:attack;type:varchar(255)"`
	ElemantAttaxk string `gorm:"column:element_attack;type:varchar(255)"`
	Shapness      string `gorm:"column:shapness;type:varchar(255)"`
	Critical      string `gorm:"column:critical;type:varchar(255)"`
	Description   string `gorm:"column:description;type:varchar(255)"`
}

type Ranking struct {
	gorm.Model
	MonsterId string `gorm:"column:monster_id;type:varchar(10); not null"`
	Ranking   string `gorm:"column:ranking;type:varchar(10)"`
	VoteYear  string `gorm:"column:vote_year;type:varchar(20)"`
}

type BgmRanking struct {
	gorm.Model
	MusicId  string `gorm:"column:music_id;type:varchar(10); not null"`
	Ranking  string `gorm:"column:ranking;type:varchar(10)"`
	VoteYear string `gorm:"column:vote_year;type:varchar(20)"`
}
