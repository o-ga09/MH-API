package weapons

import "context"

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	Find(ctx context.Context, params SearchParams) (*SearchResult, error)
}
