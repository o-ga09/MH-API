package skills

import "context"

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	Find(ctx context.Context, params SearchParams) (*SearchResult, error)
	FindAll(ctx context.Context) (Skills, error)
	FindById(ctx context.Context, skillId string) (*Skill, error)
}
