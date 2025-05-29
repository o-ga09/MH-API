package mysql

import (
	"context"
	"mh-api/internal/domain/armors"
)

type armorRepository struct {
	queryService *ArmorQueryService
}

func NewArmorRepository(qs *ArmorQueryService) armors.Repository {
	return &armorRepository{queryService: qs}
}

func (r *armorRepository) GetAll(ctx context.Context) (armors.Armors, error) {
	return r.queryService.GetAll(ctx)
}

func (r *armorRepository) GetByID(ctx context.Context, armorId string) (*armors.Armor, error) {
	return r.queryService.GetByID(ctx, armorId)
}
