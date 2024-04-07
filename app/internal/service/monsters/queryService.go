package monsters

import "context"

type MonsterQueryService interface {
	FetchMonsterList(ctx context.Context, id string) ([]*FetchMonsterListDto, error)
	FetchMonsterRank(ctx context.Context) ([]*FetchMonsterRankingDto, error)
}

type FetchMonsterListDto struct {
	Id                 string             `json:"id,omitempty"`
	Name               string             `json:"name,omitempty"`
	Description        string             `json:"description,omitempty"`
	Location           []string           `json:"location,omitempty"`
	Category           string             `json:"category,omitempty"`
	Title              []string           `json:"title,omitempty"`
	FirstWeak_Attack   string             `json:"first_weak_attack,omitempty"`
	SecondWeak_Attack  string             `json:"second_weak_attack,omitempty"`
	FirstWeak_Element  string             `json:"first_weak_element,omitempty"`
	SecondWeak_Element string             `json:"second_weak_element,omitempty"`
	Weakness_attack    []Weakness_attack  `json:"weakness___attack,omitempty"`
	Weakness_element   []Weakness_element `json:"weakness___element,omitempty"`
}

type FetchMonsterRankingDto struct {
	Id          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Location    []string  `json:"location,omitempty"`
	Category    string    `json:"category,omitempty"`
	Title       []string  `json:"title,omitempty"`
	Ranking     []Ranking `json:"ranking,omitempty"`
}

type Weakness_attack struct {
	PartId   string `json:"part_id,omitempty"`
	Slashing string `json:"slashing,omitempty"`
	Blow     string `json:"blow,omitempty"`
	Bullet   string `json:"bullet,omitempty"`
}

type Weakness_element struct {
	PartId  string `json:"part_id,omitempty"`
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
