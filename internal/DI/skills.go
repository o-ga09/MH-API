package di

import (
	"context"

	skillController "mh-api/internal/controller/skill"
	"mh-api/internal/database/mysql"
	skillService "mh-api/internal/service/skills"
)

// InitSkillsHandler は SkillHandler とその依存関係を初期化し、SkillHandler のインスタンスを返します。
func InitSkillsHandler(ctx context.Context) *skillController.SkillHandler {
	// 1. リポジトリ層の初期化
	skillRepo := mysql.NewSkillQueryService()

	// 2. サービス層の初期化
	skillsSvc := skillService.NewService(skillRepo)

	// 3. コントローラー層（ハンドラー）の初期化
	skillCtrl := skillController.NewSkillHandler(skillsSvc)

	return skillCtrl
}
