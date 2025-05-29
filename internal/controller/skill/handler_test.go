package skill

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"mh-api/internal/service/skills"
	"mh-api/pkg/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockNewServer はテスト用のサーバーを初期化する関数
func MockNewServer(t *testing.T, skillHandler *SkillHandler) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	e := gin.New()

	return e
}

// SetupRouter はテスト用のルーターを設定する関数
func (h *SkillHandler) SetupRouter(r *gin.Engine) {
	r.GET("/v1/skills", h.GetSkills)
	r.GET("/v1/skills/:skillId", h.GetSkill)
}

func TestSkillHandler_GetSkills(t *testing.T) {
	// テストケースを定義
	tests := []struct {
		name       string
		input      map[string]any
		mock       func() *skills.ISkillServiceMock
		wantStatus int
		goldenFile string
	}{
		{
			name:  "正常系：スキルの一覧が取得できる",
			input: map[string]any{},
			mock: func() *skills.ISkillServiceMock {
				mock := &skills.ISkillServiceMock{
					GetAllSkillsFunc: func(ctx context.Context) (*skills.SkillListResponseDTO, error) {
						return &skills.SkillListResponseDTO{
							Skills: []skills.SkillDTO{
								{
									ID:          "0000000001",
									Name:        "攻撃",
									Description: "攻撃力が上昇する",
									Level: []skills.SkillLevelDTO{
										{Level: 1, Description: "攻撃力+3"},
										{Level: 2, Description: "攻撃力+6"},
									},
								},
								{
									ID:          "0000000002",
									Name:        "防御",
									Description: "防御力が上昇する",
									Level: []skills.SkillLevelDTO{
										{Level: 1, Description: "防御力+5"},
									},
								},
							},
						}, nil
					},
				}
				return mock
			},
			wantStatus: http.StatusOK,
			goldenFile: "skills/get_skills_success.json.golden",
		},
		{
			name:  "異常系：サービス層でエラーが発生する",
			input: map[string]any{},
			mock: func() *skills.ISkillServiceMock {
				mock := &skills.ISkillServiceMock{
					GetAllSkillsFunc: func(ctx context.Context) (*skills.SkillListResponseDTO, error) {
						return nil, errors.New("service error")
					},
				}
				return mock
			},
			wantStatus: http.StatusInternalServerError,
			goldenFile: "skills/get_skills_error.json.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックサービスの作成
			mockService := tt.mock()

			// ハンドラーの初期化
			handler := NewSkillHandler(mockService)

			// テスト用サーバーの作成
			router := MockNewServer(t, handler)
			handler.SetupRouter(router)

			// リクエストの作成
			req, _ := http.NewRequest("GET", "/v1/skills", nil)
			rec := httptest.NewRecorder()

			// テスト実行
			router.ServeHTTP(rec, req)

			// ステータスコードのアサーション
			assert.Equal(t, tt.wantStatus, rec.Code)

			// レスポンスボディのアサーション
			testutil.AssertGoldenJSON(t, tt.goldenFile, rec.Body.Bytes())
		})
	}
}

func TestSkillHandler_GetSkill(t *testing.T) {
	// テストケースを定義
	tests := []struct {
		name       string
		skillId    string
		mock       func() *skills.ISkillServiceMock
		wantStatus int
		goldenFile string
	}{
		{
			name:    "正常系：存在するスキルIDの場合",
			skillId: "0000000001",
			mock: func() *skills.ISkillServiceMock {
				mock := &skills.ISkillServiceMock{
					GetSkillByIDFunc: func(ctx context.Context, skillID string) (*skills.SkillDTO, error) {
						return &skills.SkillDTO{
							ID:          "0000000001",
							Name:        "攻撃",
							Description: "攻撃力が上昇する",
							Level: []skills.SkillLevelDTO{
								{Level: 1, Description: "攻撃力+3"},
								{Level: 2, Description: "攻撃力+6"},
							},
						}, nil
					},
				}
				return mock
			},
			wantStatus: http.StatusOK,
			goldenFile: "skills/get_skill_success.json.golden",
		},
		{
			name:    "準正常系：不正なスキルIDの場合",
			skillId: " ",
			mock: func() *skills.ISkillServiceMock {
				mock := &skills.ISkillServiceMock{}
				return mock
			},
			wantStatus: http.StatusBadRequest,
			goldenFile: "skills/get_skill_bad_request.json.golden",
		},
		{
			name:    "異常系：存在しないスキルIDの場合",
			skillId: "9999999999",
			mock: func() *skills.ISkillServiceMock {
				mock := &skills.ISkillServiceMock{
					GetSkillByIDFunc: func(ctx context.Context, skillID string) (*skills.SkillDTO, error) {
						return nil, errors.New("skill not found")
					},
				}
				return mock
			},
			wantStatus: http.StatusInternalServerError,
			goldenFile: "skills/get_skill_error.json.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックサービスの作成
			mockService := tt.mock()

			// ハンドラーの初期化
			handler := NewSkillHandler(mockService)

			// テスト用サーバーの作成
			router := MockNewServer(t, handler)
			handler.SetupRouter(router)

			// リクエストの作成
			url := "/v1/skills/" + tt.skillId
			req, _ := http.NewRequest("GET", url, nil)
			rec := httptest.NewRecorder()

			// テスト実行
			router.ServeHTTP(rec, req)

			// ステータスコードのアサーション
			assert.Equal(t, tt.wantStatus, rec.Code)

			// レスポンスボディのアサーション
			testutil.AssertGoldenJSON(t, tt.goldenFile, rec.Body.Bytes())
		})
	}
}
