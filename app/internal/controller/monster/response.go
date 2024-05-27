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

type MonsterRanking struct {
	Ranking []ResponseRankingJson `json:"ranking,omitempty"`
}

type MessageResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseJson struct {
	Id                 string             `json:"monster_id,omitempty"`
	Name               string             `json:"name,omitempty"`
	Desc               string             `json:"desc,omitempty"`
	Location           []string           `json:"location,omitempty"`
	Tribe              string             `json:"tribe,omitempty"`
	Title              []string           `json:"title,omitempty"`
	FirstWeak_Attack   string             `json:"first_weak_attack,omitempty"`
	SecondWeak_Attack  string             `json:"second_weak_attack,omitempty"`
	FirstWeak_Element  string             `json:"first_weak_element,omitempty"`
	SecondWeak_Element string             `json:"second_weak_element,omitempty"`
	Weakness_attack    []Weakness_attack  `json:"weakness_attack,omitempty"`
	Weakness_element   []Weakness_element `json:"weakness_element,omitempty"`
}

type ResponseRankingJson struct {
	Id       string    `json:"monster_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Desc     string    `json:"desc,omitempty"`
	Location []string  `json:"location,omitempty"`
	Tribe    string    `json:"tribe,omitempty"`
	Title    []string  `json:"title,omitempty"`
	Ranking  []Ranking `json:"ranking,omitempty"`
}

type Weakness_attack struct {
	Part     string `json:"part,omitempty"`
	Slashing string `json:"slashing,omitempty"`
	Blow     string `json:"blow,omitempty"`
	Bullet   string `json:"bullet,omitempty"`
}

type Weakness_element struct {
	Part    string `json:"part,omitempty"`
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
