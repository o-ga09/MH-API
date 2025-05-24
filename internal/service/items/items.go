package items

import (
	"context"
	"mh-api/internal/domain/items"
)

//go:generate moq -out items_mock.go . IitemService
type IitemService interface {
	GetAllItems(ctx context.Context) (*ItemListResponseDTO, error)
}

type ItemDTO struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
}
type ItemListResponseDTO struct {
	Items []ItemDTO `json:"items"`
}
type Service struct {
	itemRepo items.Repository
}

func NewService(itemRepo items.Repository) *Service {
	return &Service{
		itemRepo: itemRepo,
	}
}

func (s *Service) GetAllItems(ctx context.Context) (*ItemListResponseDTO, error) {
	domainItems, err := s.itemRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var itemDTOs []ItemDTO
	for _, domainItem := range domainItems {
		itemDTOs = append(itemDTOs, ItemDTO{
			ItemID:   domainItem.GetID(),
			ItemName: domainItem.GetName(),
		})
	}

	return &ItemListResponseDTO{Items: itemDTOs}, nil
}
