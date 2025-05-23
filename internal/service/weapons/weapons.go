package weapons

import (
	"context"
	"mh-api/internal/domain/weapons"
)

//go:generate moq -out mock_weapon_query_service_test.go . IWeaponQueryService
type IWeaponQueryService interface {
	FindWeapons(ctx context.Context, params SearchWeaponsParams) ([]*weapons.Weapon, int, error)
}

type WeaponService struct {
	queryService IWeaponQueryService
}

func NewWeaponService(qs IWeaponQueryService) *WeaponService {
	return &WeaponService{queryService: qs}
}

func (s *WeaponService) SearchWeapons(ctx context.Context, params SearchWeaponsParams) (*ListWeaponsResponse, error) {

	domainWeapons, totalCount, err := s.queryService.FindWeapons(ctx, params)
	if err != nil {
		return nil, err
	}

	weaponDataList := ToWeaponDataList(domainWeapons)

	currentOffset := 0
	if params.Offset != nil {
		currentOffset = *params.Offset
	}
	currentLimit := 0

	if params.Limit != nil {
		currentLimit = *params.Limit
	}

	return &ListWeaponsResponse{
		Weapons:    weaponDataList,
		TotalCount: totalCount,
		Offset:     currentOffset,
		Limit:      currentLimit,
	}, nil
}
