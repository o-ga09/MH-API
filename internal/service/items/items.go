package items

import (
	"context"
	"mh-api/internal/domain/items"
)

//go:generate moq -out items_mock.go . IitemService
type IitemService interface {
	GetAllItems(ctx context.Context) (*ItemListResponseDTO, error)
	GetItemByID(ctx context.Context, itemID string) (*ItemDTO, error)
	GetItemByMonsterID(ctx context.Context, monsterID string) (*ItemListResponseDTO, error)
}

type ItemDTO struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
}
type ItemListResponseDTO struct {
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
	Items  []ItemDTO `json:"items"`
}

type ItemByMonster struct {
	MonsterID   string    `json:"monster_id"`
	MonsterName string    `json:"monster_name"`
	Item        []ItemDTO `json:"item"`
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

func (s *Service) GetItemByID(ctx context.Context, itemID string) (*ItemDTO, error) {
	domainItem, err := s.itemRepo.FindByID(ctx, itemID)
	if err != nil {
		return nil, err
	}

	return &ItemDTO{
		ItemID:   domainItem.GetID(),
		ItemName: domainItem.GetName(),
	}, nil
}

func (s *Service) GetItemByMonsterID(ctx context.Context, monsterID string) (*ItemListResponseDTO, error) {
	domainItems, err := s.itemRepo.FindByMonsterID(ctx, monsterID)
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
