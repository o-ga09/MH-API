package mysql

import (
	"context"
	"errors"
	"fmt"
	"mh-api/internal/domain/skills"

	"gorm.io/gorm"
)

type skillQueryService struct{}

func NewSkillQueryService() skills.Repository {
	return &skillQueryService{}
}

func (s *skillQueryService) FindAll(ctx context.Context) (skills.Skills, error) {
	var skillList []*skills.Skill
	if err := CtxFromDB(ctx).Preload("Levels").Find(&skillList).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch skills: %w", err)
	}
	return skillList, nil
}

func (s *skillQueryService) FindById(ctx context.Context, skillId string) (*skills.Skill, error) {
	var skill skills.Skill
	if err := CtxFromDB(ctx).Preload("Levels").Where("skill_id = ?", skillId).First(&skill).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to fetch skill by ID: %w", err)
	}
	return &skill, nil
}
