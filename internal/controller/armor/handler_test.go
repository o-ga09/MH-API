package armor

import (
	"context"
	"errors"
	"mh-api/internal/service/armors"
	"mh-api/pkg/testutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupTestRouter(t *testing.T, armorService IArmorService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	handler := NewArmorHandler(armorService)

	skillsGroup := r.Group("/skills")
	skillsGroup.GET("", handler.GetAllArmors)
	skillsGroup.GET("/:id", handler.GetArmorByID)

	return r
}

func TestArmorHandler_GetAllArmors(t *testing.T) {
	tests := []struct {
		name       string
		mockSetup  func() IArmorService
		wantStatus int
		goldenFile string
	}{
		{
			name: "正常系: 防具一覧が取得できる",
			mockSetup: func() IArmorService {
				return &MockArmorService{
					GetAllArmorsFunc: func(ctx context.Context) (*armors.ListArmorsResponse, error) {
						return createTestArmorsResponse(), nil
					},
				}
			},
			wantStatus: http.StatusOK,
			goldenFile: "armor/armor_get_all_200.json",
		},
		{
			name: "正常系: 防具一覧が空の場合",
			mockSetup: func() IArmorService {
				return &MockArmorService{
					GetAllArmorsFunc: func(ctx context.Context) (*armors.ListArmorsResponse, error) {
						return &armors.ListArmorsResponse{
							Armors: []armors.ArmorData{},
						}, nil
					},
				}
			},
			wantStatus: http.StatusOK,
			goldenFile: "armor/armor_get_all_empty.json",
		},
		{
			name: "異常系: 内部エラー",
			mockSetup: func() IArmorService {
				return &MockArmorService{
					GetAllArmorsFunc: func(ctx context.Context) (*armors.ListArmorsResponse, error) {
						return nil, errors.New("database error")
					},
				}
			},
			wantStatus: http.StatusInternalServerError,
			goldenFile: "armor/armor_get_all_500.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.mockSetup()
			router := setupTestRouter(t, mock)

			req, err := http.NewRequest(http.MethodGet, "/skills", nil)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
		})
	}
}

func TestArmorHandler_GetArmorByID(t *testing.T) {
	tests := []struct {
		name       string
		mockSetup  func() IArmorService
		armorID    string
		wantStatus int
		goldenFile string
	}{
		{
			name: "正常系: 防具詳細が取得できる",
			mockSetup: func() IArmorService {
				return &MockArmorService{
					GetArmorByIDFunc: func(ctx context.Context, armorId string) (*armors.ArmorData, error) {
						if armorId == "1" {
							return createTestArmorDetailResponse(), nil
						}
						return nil, gorm.ErrRecordNotFound
					},
				}
			},
			armorID:    "1",
			wantStatus: http.StatusOK,
			goldenFile: "armor/armor_get_by_id_200.json",
		},
		{
			name: "異常系: バリデーションエラー（IDが空）",
			mockSetup: func() IArmorService {
				return &MockArmorService{
					GetArmorByIDFunc: func(ctx context.Context, armorId string) (*armors.ArmorData, error) {
						return nil, nil
					},
				}
			},
			armorID:    "",
			wantStatus: http.StatusNotFound,
			goldenFile: "armor/armor_get_by_id_400.json",
		},
		{
			name: "異常系: 防具が見つからない",
			mockSetup: func() IArmorService {
				return &MockArmorService{
					GetArmorByIDFunc: func(ctx context.Context, armorId string) (*armors.ArmorData, error) {
						return nil, gorm.ErrRecordNotFound
					},
				}
			},
			armorID:    "999",
			wantStatus: http.StatusNotFound,
			goldenFile: "armor/armor_get_by_id_404.json",
		},
		{
			name: "異常系: 内部エラー",
			mockSetup: func() IArmorService {
				return &MockArmorService{
					GetArmorByIDFunc: func(ctx context.Context, armorId string) (*armors.ArmorData, error) {
						return nil, errors.New("database error")
					},
				}
			},
			armorID:    "1",
			wantStatus: http.StatusInternalServerError,
			goldenFile: "armor/armor_get_by_id_500.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.mockSetup()
			router := setupTestRouter(t, mock)

			url := "/skills/" + tt.armorID
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
		})
	}
}

type MockArmorService struct {
	GetAllArmorsFunc func(ctx context.Context) (*armors.ListArmorsResponse, error)
	GetArmorByIDFunc func(ctx context.Context, armorId string) (*armors.ArmorData, error)
}

func (m *MockArmorService) GetAllArmors(ctx context.Context) (*armors.ListArmorsResponse, error) {
	return m.GetAllArmorsFunc(ctx)
}

func (m *MockArmorService) GetArmorByID(ctx context.Context, armorId string) (*armors.ArmorData, error) {
	return m.GetArmorByIDFunc(ctx, armorId)
}

func createTestArmorsResponse() *armors.ListArmorsResponse {
	return &armors.ListArmorsResponse{
		Armors: []armors.ArmorData{
			{
				ID:   "1",
				Name: "レウスヘルム",
				Skill: []armors.SkillData{
					{ID: "1", Name: "攻撃LV1"},
					{ID: "2", Name: "火属性攻撃強化LV1"},
				},
				Slot:    "①②③",
				Defense: 100,
				Resistance: armors.ResistanceData{
					Fire:      10,
					Water:     5,
					Lightning: -10,
					Ice:       5,
					Dragon:    15,
				},
				Required: []armors.RequiredItemData{
					{ID: "ITM0019", Name: "リオレウスの鱗"},
					{ID: "ITM0016", Name: "ドラグライト鉱石"},
				},
			},
			{
				ID:   "2",
				Name: "レウスメイル",
				Skill: []armors.SkillData{
					{ID: "1", Name: "攻撃LV2"},
					{ID: "3", Name: "体力増強LV1"},
				},
				Slot:    "①①②",
				Defense: 120,
				Resistance: armors.ResistanceData{
					Fire:      15,
					Water:     0,
					Lightning: -5,
					Ice:       0,
					Dragon:    20,
				},
				Required: []armors.RequiredItemData{
					{ID: "ITM0019", Name: "リオレウスの鱗"},
					{ID: "ITM0017", Name: "大地の結晶"},
				},
			},
		},
	}
}

func createTestArmorDetailResponse() *armors.ArmorData {
	return &armors.ArmorData{
		ID:   "1",
		Name: "レウスヘルム",
		Skill: []armors.SkillData{
			{ID: "1", Name: "攻撃LV1"},
			{ID: "2", Name: "火属性攻撃強化LV1"},
		},
		Slot:    "①②③",
		Defense: 100,
		Resistance: armors.ResistanceData{
			Fire:      10,
			Water:     5,
			Lightning: -10,
			Ice:       5,
			Dragon:    15,
		},
		Required: []armors.RequiredItemData{
			{ID: "ITM0019", Name: "リオレウスの鱗"},
			{ID: "ITM0016", Name: "ドラグライト鉱石"},
		},
	}
}