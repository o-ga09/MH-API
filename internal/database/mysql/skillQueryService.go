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
	db := CtxFromDB(ctx)

	var skillModels []Skill
	if err := db.Preload("Levels").Find(&skillModels).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch skills: %w", err)
	}

	var domainSkills skills.Skills
	for _, model := range skillModels {
		levels := make([]skills.SkillLevelDetail, 0, len(model.Levels))
		for _, level := range model.Levels {
			levelDetail := skills.NewSkillLevelDetail(
				level.SkillLevelId,
				level.SkillId,
				level.Level,
				level.Description,
			)
			levels = append(levels, levelDetail)
		}

		domainSkill := skills.NewSkill(model.SkillId, model.Name, model.Description, levels)
		domainSkills = append(domainSkills, domainSkill)
	}

	return domainSkills, nil
}

func (s *skillQueryService) FindById(ctx context.Context, skillId string) (skills.Skill, error) {
	db := CtxFromDB(ctx)

	var skillModel Skill
	if err := db.Preload("Levels").Where("skill_id = ?", skillId).First(&skillModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return skills.Skill{}, gorm.ErrRecordNotFound
		}
		return skills.Skill{}, fmt.Errorf("failed to fetch skill by ID: %w", err)
	}

	levels := make([]skills.SkillLevelDetail, 0, len(skillModel.Levels))
	for _, level := range skillModel.Levels {
		levelDetail := skills.NewSkillLevelDetail(
			level.SkillLevelId,
			level.SkillId,
			level.Level,
			level.Description,
		)
		levels = append(levels, levelDetail)
	}

	domainSkill := skills.NewSkill(skillModel.SkillId, skillModel.Name, skillModel.Description, levels)
	return domainSkill, nil
}

func (s *skillQueryService) Save(ctx context.Context, skill skills.Skill) error {
	db := CtxFromDB(ctx)

	skillModel := Skill{
		SkillId:     skill.GetId(),
		Name:        skill.GetName(),
		Description: skill.GetDescription(),
	}

	for _, level := range skill.GetLevels() {
		levelModel := SkillLevel{
			SkillLevelId: level.GetLevelId(),
			SkillId:      level.GetSkillId(),
			Level:        level.GetLevel(),
			Description:  level.GetDescription(),
		}
		skillModel.Levels = append(skillModel.Levels, levelModel)
	}

	if err := db.Create(&skillModel).Error; err != nil {
		return fmt.Errorf("failed to save skill: %w", err)
	}

	return nil
}

func (s *skillQueryService) Remove(ctx context.Context, skillId string) error {
	db := CtxFromDB(ctx)

	if err := db.Where("skill_id = ?", skillId).Delete(&Skill{}).Error; err != nil {
		return fmt.Errorf("failed to remove skill: %w", err)
	}

	return nil
}
