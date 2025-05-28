package mysql

import (
	"context"
	"errors"
	"mh-api/internal/domain/skills"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestSkillQueryService_FindAll(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	tx := db.Begin()
	defer tx.Rollback()

	// テストデータを準備
	testSkills := createSkillData(t, ctx)

	// テストケースを定義
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{
			name:    "正常系: スキルを全件取得できる",
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &skillQueryService{}

			got, err := service.FindAll(ctx)

			assert.True(t, (err != nil) == tt.wantErr)
			if !tt.wantErr {
				assert.Len(t, got, tt.want)
				for i, skill := range got {
					assert.Equal(t, testSkills[i].GetId(), skill.GetId())
					assert.Equal(t, testSkills[i].GetName(), skill.GetName())
					assert.Equal(t, testSkills[i].GetDescription(), skill.GetDescription())
					assert.Len(t, skill.GetLevels(), len(testSkills[i].GetLevels()))
				}
			}
		})
	}
}

func TestSkillQueryService_FindById(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	tx := db.Begin()
	defer tx.Rollback()

	// テストデータを準備
	testSkills := createSkillData(t, ctx)

	// テストケースを定義
	tests := []struct {
		name    string
		skillId string
		want    *skills.Skill
		wantErr bool
		errType error
	}{
		{
			name:    "正常系: 存在するIDの場合",
			skillId: testSkills[0].GetId(),
			want:    &testSkills[0],
			wantErr: false,
			errType: nil,
		},
		{
			name:    "異常系: 存在しないIDの場合",
			skillId: "non-existent-id",
			want:    nil,
			wantErr: true,
			errType: gorm.ErrRecordNotFound,
		},
		{
			name:    "異常系: 空のIDの場合",
			skillId: "",
			want:    nil,
			wantErr: true,
			errType: gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// クエリサービスの初期化
			service := &skillQueryService{}

			// テスト対象メソッド実行
			got, err := service.FindById(ctx, tt.skillId)

			// アサーション
			if tt.wantErr {
				require.Error(t, err)
				if tt.errType != nil {
					assert.True(t, errors.Is(err, tt.errType), "expected error type: %v, got: %v", tt.errType, err)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want.GetId(), got.GetId())
				assert.Equal(t, tt.want.GetName(), got.GetName())
				assert.Equal(t, tt.want.GetDescription(), got.GetDescription())
				assert.Len(t, got.GetLevels(), len(tt.want.GetLevels()))
			}
		})
	}
}

func TestSkillQueryService_Save(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	tx := db.Begin()
	defer tx.Rollback()

	// テストケースを定義
	tests := []struct {
		name    string
		skill   skills.Skill
		wantErr bool
	}{
		{
			name: "正常系: スキルを保存できる",
			skill: skills.NewSkill(
				"0000000003",
				"新しいスキル",
				"新しいスキルの説明",
				[]skills.SkillLevelDetail{
					skills.NewSkillLevelDetail("0000000005", "0000000003", 1, "レベル1の説明"),
				},
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &skillQueryService{}

			err := service.Save(ctx, tt.skill)

			assert.True(t, (err != nil) == tt.wantErr)
			if !tt.wantErr {
				// 保存されたデータを検証
				var skillModel Skill
				err := CtxFromTestDB(ctx).Preload("Levels").Where("skill_id = ?", tt.skill.GetId()).First(&skillModel).Error
				require.NoError(t, err)
				assert.Equal(t, tt.skill.GetId(), skillModel.SkillId)
				assert.Equal(t, tt.skill.GetName(), skillModel.Name)
				assert.Equal(t, tt.skill.GetDescription(), skillModel.Description)
				assert.Len(t, skillModel.Levels, len(tt.skill.GetLevels()))
			}
		})
	}
}

func TestSkillQueryService_Remove(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	tx := db.Begin()
	defer tx.Rollback()

	// テストデータを準備
	testSkills := createSkillData(t, ctx)

	// テストケースを定義
	tests := []struct {
		name    string
		skillId string
		wantErr bool
	}{
		{
			name:    "正常系: 存在するIDの場合",
			skillId: testSkills[0].GetId(),
			wantErr: false,
		},
		{
			name:    "正常系: 存在しないIDの場合（エラーにならない）",
			skillId: "non-existent-id",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &skillQueryService{}

			err := service.Remove(ctx, tt.skillId)

			assert.True(t, (err != nil) == tt.wantErr)
			if !tt.wantErr && tt.skillId == testSkills[0].GetId() {
				// 削除されたことを検証
				var count int64
				CtxFromTestDB(ctx).Model(&Skill{}).Where("skill_id = ?", tt.skillId).Count(&count)
				assert.Equal(t, int64(0), count)
			}
		})
	}
}

func createSkillData(t *testing.T, ctx context.Context) skills.Skills {
	t.Helper()

	skillModels := []Skill{
		{SkillId: "0000000001", Name: "攻撃", Description: "攻撃力が上昇する"},
		{SkillId: "0000000002", Name: "防御", Description: "防御力が上昇する"},
	}

	skillLevelModels := []SkillLevel{
		{SkillLevelId: "0000000001", SkillId: "0000000001", Level: 1, Description: "攻撃力+3"},
		{SkillLevelId: "0000000002", SkillId: "0000000001", Level: 2, Description: "攻撃力+6"},
		{SkillLevelId: "0000000003", SkillId: "0000000002", Level: 1, Description: "防御力+5"},
		{SkillLevelId: "0000000004", SkillId: "0000000002", Level: 2, Description: "防御力+10"},
	}

	err := CtxFromTestDB(ctx).Create(skillModels).Error
	require.NoError(t, err)

	err = CtxFromTestDB(ctx).Create(skillLevelModels).Error
	require.NoError(t, err)

	var domainSkills skills.Skills
	for _, model := range skillModels {
		var levels []skills.SkillLevelDetail
		for _, level := range skillLevelModels {
			if level.SkillId == model.SkillId {
				levelDetail := skills.NewSkillLevelDetail(
					level.SkillLevelId,
					level.SkillId,
					level.Level,
					level.Description,
				)
				levels = append(levels, levelDetail)
			}
		}
		domainSkill := skills.NewSkill(model.SkillId, model.Name, model.Description, levels)
		domainSkills = append(domainSkills, domainSkill)
	}

	return domainSkills
}