package monster

type Monsters struct {
	Total    int            `json:"total,omitempty"`
	Limit    int            `json:"limit,omitempty"`
	Offset   int            `json:"offset,omitempty"`
	Monsters []ResponseJson `json:"monsters,omitempty"`
}

type Monster struct {
	Monster ResponseJson `json:"monster"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ResponseJson struct {
	Id                 string              `json:"monster_id,omitempty"`          // モンスターID
	Name               string              `json:"name,omitempty"`                // モンスター名
	AnotherName        *string             `json:"another_name,omitempty"`        // モンスター別名
	Location           []*string           `json:"location,omitempty"`            // モンスターの出現場所
	Category           string              `json:"category,omitempty"`            // モンスターのカテゴリ
	Title              []*string           `json:"title,omitempty"`               // 登場作品
	FirstWeak_Attack   *string             `json:"first_weak_attack,omitempty"`   // 最有効弱点
	SecondWeak_Attack  *string             `json:"second_weak_attack,omitempty"`  // 2番目に有効な弱点
	FirstWeak_Element  *string             `json:"first_weak_element,omitempty"`  // 最有効属性
	SecondWeak_Element *string             `json:"second_weak_element,omitempty"` // 2番目に有効な属性
	Weakness_attack    []*Weakness_attack  `json:"weakness_attack,omitempty"`     // 弱点肉質（物理）
	Weakness_element   []*Weakness_element `json:"weakness_element,omitempty"`    // 弱点肉質（属性）
	Ranking            []*Ranking          `json:"ranking,omitempty"`             // 人気投票ランキング
	ImageUrl           *string             `json:"image_url,omitempty"`           // モンスター画像URL
	BGM                []*Music            `json:"bgm,omitempty"`                 // BGM
}

type Weakness_attack struct {
	Slashing string `json:"slashing,omitempty"`
	Blow     string `json:"blow,omitempty"`
	Bullet   string `json:"bullet,omitempty"`
}

type Weakness_element struct {
	Fire    string `json:"fire,omitempty"`
	Ice     string `json:"ice,omitempty"`
	Water   string `json:"water,omitempty"`
	Thunder string `json:"thunder,omitempty"`
	Dragon  string `json:"dragon,omitempty"`
}

type Ranking struct {
	Ranking  string `json:"ranking,omitempty"`
	VoteYear string `json:"vote_year,omitempty"`
}

type Music struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}
