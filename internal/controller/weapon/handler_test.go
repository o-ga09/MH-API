package weapon

import (
	"context"
	"errors"
	"mh-api/internal/service/weapons"
	"mh-api/pkg/testutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// サーバー初期化
func setupTestRouter(t *testing.T, weaponService IWeaponService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	handler := NewWeaponHandler(weaponService)

	// ルーティング設定
	weaponsGroup := r.Group("/weapons")
	weaponsGroup.GET("", handler.SearchWeapons)

	return r
}

// TestWeaponHandler_SearchWeapons は、SearchWeapons関数のテストを実行します
func TestWeaponHandler_SearchWeapons(t *testing.T) {
	// デフォルト値を設定
	defaultLimit := 20
	defaultOffset := 0

	tests := []struct {
		name       string
		mockSetup  func() IWeaponService
		query      string
		wantStatus int
		goldenFile string
	}{
		{
			name: "正常系: 武器一覧が取得できる",
			mockSetup: func() IWeaponService {
				return &MockWeaponService{
					SearchWeaponsFunc: func(ctx context.Context, params weapons.SearchWeaponsParams) (*weapons.ListWeaponsResponse, error) {
						// パラメータの検証
						assert.Equal(t, defaultLimit, *params.Limit)
						assert.Equal(t, defaultOffset, *params.Offset)

						return createTestWeaponsResponse(), nil
					},
				}
			},
			query:      "?limit=20&offset=0",
			wantStatus: http.StatusOK,
			goldenFile: "weapon/weapon_search_200.json",
		},
		{
			name: "異常系: バリデーションエラー",
			mockSetup: func() IWeaponService {
				return &MockWeaponService{
					SearchWeaponsFunc: func(ctx context.Context, params weapons.SearchWeaponsParams) (*weapons.ListWeaponsResponse, error) {
						// このモックは呼び出されないはず
						return nil, nil
					},
				}
			},
			// limitに負の値を指定してバリデーションエラーを発生させる
			query:      "?limit=-1",
			wantStatus: http.StatusBadRequest,
			goldenFile: "weapon/weapon_search_400.json",
		},
		{
			name: "異常系: 内部エラー",
			mockSetup: func() IWeaponService {
				return &MockWeaponService{
					SearchWeaponsFunc: func(ctx context.Context, params weapons.SearchWeaponsParams) (*weapons.ListWeaponsResponse, error) {
						return nil, errors.New("database error")
					},
				}
			},
			query:      "?limit=20&offset=0",
			wantStatus: http.StatusInternalServerError,
			goldenFile: "weapon/weapon_search_500.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.mockSetup()
			router := setupTestRouter(t, mock)

			req, err := http.NewRequest(http.MethodGet, "/weapons"+tt.query, nil)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
		})
	}
}

// MockWeaponService はIWeaponServiceのモック実装です
type MockWeaponService struct {
	SearchWeaponsFunc func(ctx context.Context, params weapons.SearchWeaponsParams) (*weapons.ListWeaponsResponse, error)
}

// SearchWeapons はモックのSearchWeapons実装です
func (m *MockWeaponService) SearchWeapons(ctx context.Context, params weapons.SearchWeaponsParams) (*weapons.ListWeaponsResponse, error) {
	return m.SearchWeaponsFunc(ctx, params)
}

// createTestWeaponsResponse はテスト用の武器レスポンスを作成します
func createTestWeaponsResponse() *weapons.ListWeaponsResponse {
	return &weapons.ListWeaponsResponse{
		Weapons: []weapons.WeaponData{
			{
				WeaponID:      "1",
				Name:          "リオレウス剣",
				ImageURL:      "http://example.com/weapon1.jpg",
				Rare:          "5",
				Attack:        "800",
				ElementAttack: "火 200",
				Sharpness:     "青",
				Critical:      "0%",
				Description:   "リオレウスの素材から作られた剣",
			},
			{
				WeaponID:      "2",
				Name:          "ジンオウガ太刀",
				ImageURL:      "http://example.com/weapon2.jpg",
				Rare:          "6",
				Attack:        "750",
				ElementAttack: "雷 300",
				Sharpness:     "白",
				Critical:      "10%",
				Description:   "ジンオウガの素材から作られた太刀",
			},
		},
		TotalCount: 2,
		Offset:     0,
		Limit:      20,
	}
}
