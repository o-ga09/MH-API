package entity

type Monsters struct {
	Values []Monster
}

type MonsterId struct {
	Value string //モンスターID
}

type MonsterName struct {
	Value string //モンスターの名前
}

type MonsterDesc struct {
	Value string //モンスターの説明文
}

type MonsterLocation struct {
	Value string //モンスターの生息地
}

type MonsterCategory struct {
	Value string //モンスターの分類
}

type GameTitle struct {
	Value string // 登場したゲームタイトル
}

type MonsterWeakness_A struct {
	Value Weakness_attack //モンスターの弱点(物理肉質)
}

type MonsterWeakness_E struct {
	Value Weakness_element //モンスターの弱点(属性肉質)

}

type Monster struct {
	Id               MonsterId         `db:"id" json:"id,omitempty"`
	Name             MonsterName       `db:"name" json:"name,omitempty"`
	Desc             MonsterDesc       `db:"desc" json:"desc,omitempty"`
	Location         MonsterLocation   `db:"location" json:"location,omitempty"`
	Category         MonsterCategory   `db:"category" json:"category,omitempty"`
	Title            GameTitle         `db:"title,omitempty"`
	Weakness_attack  MonsterWeakness_A `db:"weakness_attack" json:"weakness___attack,omitempty"`
	Weakness_element MonsterWeakness_E `db:"weakness_element" json:"weakness___element,omitempty"`
}

type MonsterJson struct {
	Name             MonsterName       `json:"name,omitempty"`
	Desc             MonsterDesc       `json:"desc,omitempty"`
	Location         MonsterLocation   `json:"location,omitempty"`
	Category         MonsterCategory   `json:"category,omitempty"`
	Title            GameTitle         `json:"title,omitempty"`
	Weakness_attack  MonsterWeakness_A `json:"weakness_attack,omitempty"`
	Weakness_element MonsterWeakness_E `json:"weakness_element,omitempty"`
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
