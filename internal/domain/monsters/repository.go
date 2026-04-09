package monsters

import "context"

// SearchParams はモンスター検索の条件を表す
type SearchParams struct {
	MonsterIds      string
	MonsterName     string
	UsageElement    string
	WeaknessElement string
	TribeName       string
	FieldName       string
	ProductName     string
	Limit           int
	Offset          int
	Sort            string
}

// SearchResult はモンスター検索の結果を表す
type SearchResult struct {
	Monsters []*Monster
	Total    int
}

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	FindAll(ctx context.Context, params SearchParams) (*SearchResult, error)
	FindById(ctx context.Context, id string) (*Monster, error)
}
