package item

import "context"

//go:generate moq -out queryService_mock.go . MonsterQueryService
type ItemQueryService interface {
	FetchList(ctx context.Context, id string) ([]*FetchItemListDto, error)
	FetchListWithMonster(ctx context.Context) ([]*FetchItemListWithMonsterDto, error)
	FetchListByMonster(ctx context.Context, id string) ([]*FetchItemListByMonsterDto, error)
}

type FetchItemListDto struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type FetchItemListWithMonsterDto struct {
	Id      string     `json:"id,omitempty"`
	Name    string     `json:"name,omitempty"`
	Url     string     `json:"url,omitempty"`
	Monster []*Monster `json:"ranking,omitempty"`
}

type FetchItemListByMonsterDto struct {
	MonsterId   string              `json:"id,omitempty"`
	MonsterName string              `json:"name,omitempty"`
	Item        []*FetchItemListDto `json:"item,omitempty"`
}

type Monster struct {
	MonsterId   string `json:"monster_id,omitempty"`
	MonsterName string `json:"monster_name,omitempty"`
}

type Ranking struct {
	Rank     string `json:"rank,omitempty"`
	VoteYear string `json:"vote_year,omitempty"`
}
