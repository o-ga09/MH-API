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

func (s *skillQueryService) Find(ctx context.Context, params skills.SearchParams) (*skills.SearchResult, error) {
	var skillList []*skills.Skill
	query := CtxFromDB(ctx).Preload("Levels")

	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if params.Description != "" {
		query = query.Where("description LIKE ?", "%"+params.Description+"%")
	}

	limit := params.Limit
	if limit <= 0 {
		limit = 100
	}

	var total int64
	countQuery := CtxFromDB(ctx).Model(&skills.Skill{})
	if params.Name != "" {
		countQuery = countQuery.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if params.Description != "" {
		countQuery = countQuery.Where("description LIKE ?", "%"+params.Description+"%")
	}
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count skills: %w", err)
	}

	if err := query.Limit(limit).Offset(params.Offset).Find(&skillList).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch skills: %w", err)
	}
	return &skills.SearchResult{Skills: skillList, Total: int(total)}, nil
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
