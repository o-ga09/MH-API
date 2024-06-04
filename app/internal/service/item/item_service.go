package item

import (
	"mh-api/app/internal/domain/items"
)

type ItemService struct {
	repo items.Repository
	qs   ItemQueryService
}

func NewItemService(repo items.Repository, qs ItemQueryService) *ItemService {
	return &ItemService{
		repo: repo,
		qs:   qs,
	}
}
