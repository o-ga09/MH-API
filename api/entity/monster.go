package entity

type Monsters []Monster

type Monster struct {
	Id               int    `db:"id" json:"id,omitempty"`
	Name             string `db:"name" json:"name,omitempty"`
	Desc             string `db:"desc" json:"desc,omitempty"`
	Location         string `db:"location" json:"location,omitempty"`
	Specify          string `db:"specify" json:"specify,omitempty"`
	Weakness_attack  string `db:"weakness_attack" json:"weakness___attack,omitempty"`
	Weakness_element string `db:"weakness_element" json:"weakness___element,omitempty"`
}
