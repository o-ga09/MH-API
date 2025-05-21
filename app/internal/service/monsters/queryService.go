package monsters

import (
	"context"
	"mh-api/app/internal/domain/music"
)

//go:generate moq -out queryService_mock.go . MonsterQueryService
type MonsterQueryService interface {
	FetchList(ctx context.Context, id string) ([]*FetchMonsterListDto, error)
}

type FetchMonsterListDto struct {
	Id                 string             `json:"id,omitempty"`
	Name               string             `json:"name,omitempty"`
	NameEnglish        *string            `json:"name_english,omitempty"`
	Description        string             `json:"description,omitempty"`
	AnotherName        string             `json:"another_name,omitempty"`
	NameEn             string             `json:"name_en,omitempty"`
	Element            *string            `json:"element,omitempty"`
	Location           []string           `json:"location,omitempty"`
	Category           string             `json:"category,omitempty"`
	CategoryEnglish    *string            `json:"category_english,omitempty"`
	Title              []string           `json:"title,omitempty"`
	FirstWeak_Attack   string             `json:"first_weak_attack,omitempty"`
	SecondWeak_Attack  string             `json:"second_weak_attack,omitempty"`
	FirstWeak_Element  string             `json:"first_weak_element,omitempty"`
	SecondWeak_Element string             `json:"second_weak_element,omitempty"`
	Weakness_attack    []Weakness_attack  `json:"weakness___attack,omitempty"`
	Weakness_element   []Weakness_element `json:"weakness___element,omitempty"`
	Ranking            []Ranking          `json:"ranking,omitempty"`
	BGM                []music.Music      `json:"bgm,omitempty"`
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
