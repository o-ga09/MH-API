package mysql

import (
	"context"
	"errors"
	"sort"
	"testing"

	"mh-api/internal/domain/skills"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestSkillQueryService_FindAll(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	db.Begin()
	defer db.Rollback()

	testSkills := createSkillData(t, ctx)

	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{name: "正常系: スキルを全件取得できる", want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &skillQueryService{}
			got, err := service.FindAll(ctx)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Len(t, got, tt.want)

			sortedGot := make([]*skills.Skill, len(got))
			copy(sortedGot, got)
			sortedExpected := make([]*skills.Skill, len(testSkills))
			copy(sortedExpected, testSkills)

			sort.Slice(sortedGot, func(i, j int) bool { return sortedGot[i].SkillId < sortedGot[j].SkillId })
			sort.Slice(sortedExpected, func(i, j int) bool { return sortedExpected[i].SkillId < sortedExpected[j].SkillId })

			for i := range sortedGot {
				assert.Equal(t, sortedExpected[i].SkillId, sortedGot[i].SkillId)
				assert.Equal(t, sortedExpected[i].Name, sortedGot[i].Name)
				assert.Equal(t, sortedExpected[i].Description, sortedGot[i].Description)
				assert.Len(t, sortedGot[i].Levels, len(sortedExpected[i].Levels))
			}
		})
	}
}

func TestSkillQueryService_FindById(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	db.Begin()
	defer db.Rollback()

	testSkills := createSkillData(t, ctx)

	tests := []struct {
		name    string
		skillId string
		want    *skills.Skill
		wantErr bool
		errType error
	}{
		{
			name:    "正常系: 存在するIDの場合",
			skillId: testSkills[0].SkillId,
			want:    testSkills[0],
		},
		{
			name:    "異常系: 存在しないIDの場合",
			skillId: "non-existent-id",
			wantErr: true,
			errType: gorm.ErrRecordNotFound,
		},
		{
			name:    "異常系: 空のIDの場合",
			skillId: "",
			wantErr: true,
			errType: gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &skillQueryService{}
			got, err := service.FindById(ctx, tt.skillId)

			if tt.wantErr {
				require.Error(t, err)
				if tt.errType != nil {
					assert.True(t, errors.Is(err, tt.errType))
				}
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want.SkillId, got.SkillId)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.Description, got.Description)
			assert.Len(t, got.Levels, len(tt.want.Levels))
		})
	}
}

func createSkillData(t *testing.T, ctx context.Context) []*skills.Skill {
	t.Helper()

	require.NoError(t, CtxFromTestDB(ctx).Exec("DELETE FROM skill_level").Error)
	require.NoError(t, CtxFromTestDB(ctx).Exec("DELETE FROM skill").Error)

	skillList := []*skills.Skill{
		{SkillId: "0000000001", Name: "攻撃", Description: "攻撃力が上昇する"},
		{SkillId: "0000000002", Name: "防御", Description: "防御力が上昇する"},
	}
	require.NoError(t, CtxFromTestDB(ctx).Create(skillList).Error)

	levelList := []*skills.SkillLevel{
		{SkillLevelId: "0000000001", SkillId: "0000000001", Level: 1, Description: "攻撃力+3"},
		{SkillLevelId: "0000000002", SkillId: "0000000001", Level: 2, Description: "攻撃力+6"},
		{SkillLevelId: "0000000003", SkillId: "0000000002", Level: 1, Description: "防御力+5"},
		{SkillLevelId: "0000000004", SkillId: "0000000002", Level: 2, Description: "防御力+10"},
	}
	require.NoError(t, CtxFromTestDB(ctx).Create(levelList).Error)

	// Levelsをセット（DBから取得せずにメモリ上で構築）
	skillList[0].Levels = []skills.SkillLevel{*levelList[0], *levelList[1]}
	skillList[1].Levels = []skills.SkillLevel{*levelList[2], *levelList[3]}

	return skillList
}
