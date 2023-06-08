package entity

type Monsters struct {
	Values []Monster
}

type MonsterId struct {
	Value int           //モンスターID
}

type MonsterName struct {
	Value string       //モンスターの名前
}

type MonsterDesc struct {
	Value string       //モンスターの説明文
} 

type MonsterLocation struct {
	Value string       //モンスターの生息地
} 

type MonsterSpecify struct {
	Value string       //モンスターの分類
} 

type MonsterWeakness_A struct {
	Value map[string]string       //モンスターの弱点(物理肉質)
} 

type MonsterWeakness_E struct {
	Value map[string]string       //モンスターの弱点(属性肉質)

} 

type Monster struct {
	Id               MonsterId         `db:"id" json:"id,omitempty"`
	Name             MonsterName       `db:"name" json:"name,omitempty"`
	Desc             MonsterDesc       `db:"desc" json:"desc,omitempty"`
	Location         MonsterLocation   `db:"location" json:"location,omitempty"`
	Specify          MonsterSpecify    `db:"specify" json:"specify,omitempty"`
	Weakness_attack  MonsterWeakness_A `db:"weakness_attack" json:"weakness___attack,omitempty"`
	Weakness_element MonsterWeakness_E `db:"weakness_element" json:"weakness___element,omitempty"`
}

type MonsterJson struct {
	Name             MonsterName       `json:"name"`
	Desc             MonsterDesc       `json:"desc"`
	Location         MonsterLocation   `json:"location"`
	Specify          MonsterSpecify    `json:"specify"`
	Weakness_attack  MonsterWeakness_A `json:"weakness___attack"`
	Weakness_element MonsterWeakness_E `json:"weakness___element"`
}