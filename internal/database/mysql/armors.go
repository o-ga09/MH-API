package mysql

import (
	"context"
	"mh-api/internal/domain/armors"
)

type armorRepository struct{}

func NewArmorRepository() armors.Repository {
	return &armorRepository{}
}

func (r *armorRepository) GetAll(ctx context.Context) (armors.Armors, error) {
	qs := NewArmorQueryService()
	return qs.GetAll(ctx)
}

func (r *armorRepository) GetByID(ctx context.Context, armorId string) (*armors.Armor, error) {
	qs := NewArmorQueryService()
	return qs.GetByID(ctx, armorId)
}