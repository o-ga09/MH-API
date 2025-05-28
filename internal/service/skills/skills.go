package skills

import (
	"context"
	"mh-api/internal/domain/skills"
)

//go:generate moq -out skills_mock.go . ISkillService
type ISkillService interface {
	GetAllSkills(ctx context.Context) (*SkillListResponseDTO, error)
	GetSkillByID(ctx context.Context, skillID string) (*SkillDTO, error)
}

type SkillDTO struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Level       []SkillLevelDTO     `json:"level"`
}

type SkillLevelDTO struct {
	Level       int    `json:"level"`
	Description string `json:"description"`
}

type SkillListResponseDTO struct {
	Skills []SkillDTO `json:"skills"`
}

type Service struct {
	skillRepo skills.Repository
}

func NewService(skillRepo skills.Repository) *Service {
	return &Service{
		skillRepo: skillRepo,
	}
}

func (s *Service) GetAllSkills(ctx context.Context) (*SkillListResponseDTO, error) {
	domainSkills, err := s.skillRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var skillDTOs []SkillDTO
	for _, domainSkill := range domainSkills {
		var levelDTOs []SkillLevelDTO
		for _, level := range domainSkill.GetLevels() {
			levelDTOs = append(levelDTOs, SkillLevelDTO{
				Level:       level.GetLevel(),
				Description: level.GetDescription(),
			})
		}

		skillDTOs = append(skillDTOs, SkillDTO{
			ID:          domainSkill.GetId(),
			Name:        domainSkill.GetName(),
			Description: domainSkill.GetDescription(),
			Level:       levelDTOs,
		})
	}

	return &SkillListResponseDTO{Skills: skillDTOs}, nil
}

func (s *Service) GetSkillByID(ctx context.Context, skillID string) (*SkillDTO, error) {
	domainSkill, err := s.skillRepo.FindById(ctx, skillID)
	if err != nil {
		return nil, err
	}

	var levelDTOs []SkillLevelDTO
	for _, level := range domainSkill.GetLevels() {
		levelDTOs = append(levelDTOs, SkillLevelDTO{
			Level:       level.GetLevel(),
			Description: level.GetDescription(),
		})
	}

	return &SkillDTO{
		ID:          domainSkill.GetId(),
		Name:        domainSkill.GetName(),
		Description: domainSkill.GetDescription(),
		Level:       levelDTOs,
	}, nil
}