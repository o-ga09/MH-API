package skills

import (
	"context"
	"errors"
	"mh-api/internal/domain/skills"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_GetAllSkills(t *testing.T) {
	tests := []struct {
		name       string
		mockFunc   func() (skills.Skills, error)
		want       *SkillListResponseDTO
		wantErr    bool
		checkCalls bool
	}{
		{
			name: "正常系: スキル一覧を取得できる",
			mockFunc: func() (skills.Skills, error) {
				return skills.Skills{
					skills.NewSkill("0000000001", "攻撃", "攻撃力が上昇する", []skills.SkillLevelDetail{
						skills.NewSkillLevelDetail("0000000001", "0000000001", 1, "攻撃力+3"),
						skills.NewSkillLevelDetail("0000000002", "0000000001", 2, "攻撃力+6"),
					}),
					skills.NewSkill("0000000002", "防御", "防御力が上昇する", []skills.SkillLevelDetail{
						skills.NewSkillLevelDetail("0000000003", "0000000002", 1, "防御力+5"),
					}),
				}, nil
			},
			want: &SkillListResponseDTO{
				Skills: []SkillDTO{
					{
						ID:          "0000000001",
						Name:        "攻撃",
						Description: "攻撃力が上昇する",
						Level: []SkillLevelDTO{
							{Level: 1, Description: "攻撃力+3"},
							{Level: 2, Description: "攻撃力+6"},
						},
					},
					{
						ID:          "0000000002",
						Name:        "防御",
						Description: "防御力が上昇する",
						Level: []SkillLevelDTO{
							{Level: 1, Description: "防御力+5"},
						},
					},
				},
			},
			wantErr:    false,
			checkCalls: true,
		},
		{
			name: "正常系: スキルが存在しない場合は空のリストが返る",
			mockFunc: func() (skills.Skills, error) {
				return skills.Skills{}, nil
			},
			want: &SkillListResponseDTO{
				Skills: []SkillDTO{},
			},
			wantErr: false,
		},
		{
			name: "異常系: リポジトリでエラーが発生する",
			mockFunc: func() (skills.Skills, error) {
				return nil, errors.New("repository error")
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックリポジトリの作成
			mockRepo := &skills.RepositoryMock{
				FindAllFunc: func(ctx context.Context) (skills.Skills, error) {
					return tt.mockFunc()
				},
			}

			// テスト対象のサービスを初期化
			service := NewService(mockRepo)

			// テスト実行
			got, err := service.GetAllSkills(context.Background())

			// アサーション
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantErr {
				return
			}

			require.NotNil(t, got)
			assert.Equal(t, len(tt.want.Skills), len(got.Skills))
			for i, skill := range got.Skills {
				assert.Equal(t, tt.want.Skills[i].ID, skill.ID)
				assert.Equal(t, tt.want.Skills[i].Name, skill.Name)
				assert.Equal(t, tt.want.Skills[i].Description, skill.Description)
				assert.Equal(t, len(tt.want.Skills[i].Level), len(skill.Level))
				for j, level := range skill.Level {
					assert.Equal(t, tt.want.Skills[i].Level[j].Level, level.Level)
					assert.Equal(t, tt.want.Skills[i].Level[j].Description, level.Description)
				}
			}

			// モックの呼び出しを確認
			if tt.checkCalls {
				assert.Equal(t, 1, len(mockRepo.FindAllCalls()))
			}
		})
	}
}

func TestService_GetSkillByID(t *testing.T) {
	tests := []struct {
		name       string
		skillID    string
		mockFunc   func(skillID string) (skills.Skill, error)
		want       *SkillDTO
		wantErr    bool
		checkCalls bool
	}{
		{
			name:    "正常系: 存在するIDの場合",
			skillID: "0000000001",
			mockFunc: func(skillID string) (skills.Skill, error) {
				return skills.NewSkill("0000000001", "攻撃", "攻撃力が上昇する", []skills.SkillLevelDetail{
					skills.NewSkillLevelDetail("0000000001", "0000000001", 1, "攻撃力+3"),
					skills.NewSkillLevelDetail("0000000002", "0000000001", 2, "攻撃力+6"),
				}), nil
			},
			want: &SkillDTO{
				ID:          "0000000001",
				Name:        "攻撃",
				Description: "攻撃力が上昇する",
				Level: []SkillLevelDTO{
					{Level: 1, Description: "攻撃力+3"},
					{Level: 2, Description: "攻撃力+6"},
				},
			},
			wantErr:    false,
			checkCalls: true,
		},
		{
			name:    "異常系: 存在しないIDの場合",
			skillID: "non-existent-id",
			mockFunc: func(skillID string) (skills.Skill, error) {
				return skills.Skill{}, errors.New("skill not found")
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "異常系: リポジトリでエラーが発生する",
			skillID: "0000000001",
			mockFunc: func(skillID string) (skills.Skill, error) {
				return skills.Skill{}, errors.New("repository error")
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックリポジトリの作成
			mockRepo := &skills.RepositoryMock{
				FindByIdFunc: func(ctx context.Context, skillId string) (skills.Skill, error) {
					return tt.mockFunc(skillId)
				},
			}

			// テスト対象のサービスを初期化
			service := NewService(mockRepo)

			// テスト実行
			got, err := service.GetSkillByID(context.Background(), tt.skillID)

			// アサーション
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantErr {
				return
			}

			require.NotNil(t, got)
			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.Description, got.Description)
			assert.Equal(t, len(tt.want.Level), len(got.Level))
			for i, level := range got.Level {
				assert.Equal(t, tt.want.Level[i].Level, level.Level)
				assert.Equal(t, tt.want.Level[i].Description, level.Description)
			}

			// モックの呼び出しを確認
			if tt.checkCalls {
				calls := mockRepo.FindByIdCalls()
				assert.Equal(t, 1, len(calls))
				assert.Equal(t, tt.skillID, calls[0].SkillId)
			}
		})
	}
}
