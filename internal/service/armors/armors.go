package armors

import (
	"context"
	"mh-api/internal/domain/armors"
)

//go:generate moq -out armors_mock.go . IArmorQueryService
type IArmorQueryService interface {
	GetAll(ctx context.Context) (armors.Armors, error)
	GetByID(ctx context.Context, armorId string) (*armors.Armor, error)
}

type ArmorService struct {
	queryService IArmorQueryService
}

func NewArmorService(qs IArmorQueryService) *ArmorService {
	return &ArmorService{queryService: qs}
}

func (s *ArmorService) GetAllArmors(ctx context.Context) (*ListArmorsResponse, error) {
	domainArmors, err := s.queryService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	armorDataList := ToArmorDataList(domainArmors)

	return &ListArmorsResponse{
		Armors: armorDataList,
	}, nil
}

func (s *ArmorService) GetArmorByID(ctx context.Context, armorId string) (*ArmorData, error) {
	domainArmor, err := s.queryService.GetByID(ctx, armorId)
	if err != nil {
		return nil, err
	}

	armorData := ToArmorData(domainArmor)
	return &armorData, nil
}
