package skills

import "context"

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	Save(ctx context.Context, s Skill) error
	Remove(ctx context.Context, skillId string) error
	FindAll(ctx context.Context) (Skills, error)
	FindById(ctx context.Context, skillId string) (Skill, error)
}