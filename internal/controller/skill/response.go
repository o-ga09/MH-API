package skill

import (
	"mh-api/internal/service/skills"
	"strconv"
)

type Skills struct {
	Skills []ResponseSkill `json:"skills,omitempty"`
}

type Skill struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Level       []ResponseSkillLevel `json:"level"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ResponseSkill struct {
	ID          string               `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	Level       []ResponseSkillLevel `json:"level,omitempty"`
}

type ResponseSkillLevel map[string]string

func ToSkillListResponse(skills skills.SkillListResponseDTO) Skills {
	res := make([]ResponseSkill, len(skills.Skills))
	for i, skill := range skills.Skills {
		levels := make([]ResponseSkillLevel, len(skill.Level))
		for j, level := range skill.Level {
			levels[j] = ResponseSkillLevel{
				strconv.Itoa(level.Level): level.Description,
			}
		}
		res[i] = ResponseSkill{
			ID:          skill.ID,
			Name:        skill.Name,
			Description: skill.Description,
			Level:       levels,
		}
	}
	return Skills{
		Skills: res,
	}
}

func ToSkillResponse(skill skills.SkillDTO) Skill {
	levels := make([]ResponseSkillLevel, len(skill.Level))
	for i, level := range skill.Level {
		levels[i] = ResponseSkillLevel{
			strconv.Itoa(level.Level): level.Description,
		}
	}
	return Skill{
		ID:          skill.ID,
		Name:        skill.Name,
		Description: skill.Description,
		Level:       levels,
	}
}
