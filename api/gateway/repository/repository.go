package repository

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type MonsterDriver interface {
	GetAll() []Monster
	GetById(id int) Monster
	Create(MonsterJson) error
	Update(id int, monsterJson MonsterJson) error
	Delete(id int) error
}

type Monster struct {
	Id               int              `gorm:"column:id" json:"id,omitempty"`
	MonsterId        string           `gorm:"column:monster_id" json:"monster_id,omitempty"`
	Name             string           `gorm:"column:name" json:"name,omitempty"`
	Desc             string           `gorm:"column:desc" json:"desc,omitempty"`
	Location         string           `gorm:"column:location" json:"location,omitempty"`
	Category         string           `gorm:"column:category" json:"category,omitempty"`
	Title            string           `gorm:"column:title" json:"title,omitempty"`
	Weakness_attack  Weakness_attack  `gorm:"column:weakness_attack" json:"weakness_attack,omitempty"`
	Weakness_element Weakness_element `gorm:"column:weakness_element" json:"weakness_element,omitempty"`
	Created_at       time.Time        `gorm:"column:created_at" json:"created___at,omitempty"`
	Updated_at       time.Time        `gorm:"column:updated_at" json:"updated___at,omitempty"`
}

type MonsterJson struct {
	Id               string           `gorm:"column:monster_id" json:"monster_id,omitempty"`
	Name             string           `gorm:"column:name" json:"name,omitempty"`
	Desc             string           `gorm:"column:desc" json:"desc,omitempty"`
	Location         string           `gorm:"column:location" json:"location,omitempty"`
	Category         string           `gorm:"column:category" json:"category,omitempty"`
	Title            string           `gorm:"column:title" json:"title,omitempty"`
	Weakness_attack  Weakness_attack  `gorm:"column:weakness_attack" json:"weakness_attack,omitempty"`
	Weakness_element Weakness_element `gorm:"column:weakness_element" json:"weakness_element,omitempty"`
}

type Weakness_attack struct {
	FrontLegs AttackCatetgory `json:"front_legs,omitempty"`
	Tail      AttackCatetgory `json:"tail,omitempty"`
	HindLegs  AttackCatetgory `json:"hind_legs,omitempty"`
	Body      AttackCatetgory `json:"body,omitempty"`
	Head      AttackCatetgory `json:"head,omitempty"`
}

type Weakness_element struct {
	FrontLegs Elements `json:"front_legs,omitempty"`
	Tail      Elements `json:"tail,omitempty"`
	HindLegs  Elements `json:"hind_legs,omitempty"`
	Body      Elements `json:"body,omitempty"`
	Head      Elements `json:"head,omitempty"`
}

type AttackCatetgory struct {
	Slashing string `json:"slashing,omitempty"`
	Blow     string `json:"blow,omitempty"`
	Bullet   string `json:"bullet,omitempty"`
}

type Elements struct {
	Fire      string `json:"fire,omitempty"`
	Water     string `json:"water,omitempty"`
	Lightning string `json:"lightning,omitempty"`
	Ice       string `json:"ice,omitempty"`
	Dragon    string `json:"dragon,omitempty"`
}

func (w Weakness_attack) Value() (driver.Value, error) {
	bytes, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

func (w *Weakness_attack) Scan(input interface{}) error {
	switch v := input.(type) {
	case string:
		return json.Unmarshal([]byte(v), w)
	case []byte:
		return json.Unmarshal(v, w)
	default:
		return fmt.Errorf("unsupported Type %T", input)
	}
}

func (w Weakness_element) Value() (driver.Value, error) {
	bytes, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

func (w *Weakness_element) Scan(input interface{}) error {
	switch v := input.(type) {
	case string:
		return json.Unmarshal([]byte(v), w)
	case []byte:
		return json.Unmarshal(v, w)
	default:
		return fmt.Errorf("unsupported Type %T", input)
	}
}

func (MonsterJson) TableName() string {
	return "monsters"
}
