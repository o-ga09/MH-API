package skill

import (
	"mh-api/internal/domain/skills"
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

func ToSkillListResponse(skillList skills.Skills) Skills {
	res := make([]ResponseSkill, len(skillList))
	for i, skill := range skillList {
		res[i] = toResponseSkill(skill)
	}
	return Skills{Skills: res}
}

func ToSkillResponse(skill skills.Skill) Skill {
	levels := make([]ResponseSkillLevel, len(skill.Levels))
	for i, level := range skill.Levels {
		levels[i] = ResponseSkillLevel{
			strconv.Itoa(level.Level): level.Description,
		}
	}
	return Skill{
		ID:          skill.SkillId,
		Name:        skill.Name,
		Description: skill.Description,
		Level:       levels,
	}
}

func toResponseSkill(skill *skills.Skill) ResponseSkill {
	levels := make([]ResponseSkillLevel, len(skill.Levels))
	for i, level := range skill.Levels {
		levels[i] = ResponseSkillLevel{
			strconv.Itoa(level.Level): level.Description,
		}
	}
	return ResponseSkill{
		ID:          skill.SkillId,
		Name:        skill.Name,
		Description: skill.Description,
		Level:       levels,
	}
}
