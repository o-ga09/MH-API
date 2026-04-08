package di

import (
	"context"
	skillController "mh-api/internal/controller/skill"
	"mh-api/internal/database/mysql"
)

func InitSkillsHandler(ctx context.Context) *skillController.SkillHandler {
	repo := mysql.NewSkillQueryService()
	return skillController.NewSkillHandler(repo)
}
