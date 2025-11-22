package monster

import (
	"context"
	"errors"
	"fmt"
	"mh-api/internal/domain/music"
	"mh-api/internal/service/monsters"
	"mh-api/pkg/testutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// サーバー初期化
func setupTestRouter(t *testing.T, monsterService monsters.IMonsterService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	handler := NewMonsterHandler(monsterService)

	// ルーティング設定
	monsters := r.Group("/monsters")
	monsters.GET("", handler.GetAll)
	monsters.GET("/:id", handler.GetById)

	return r
}

// TestMonsterHandler_GetAll は、GetAll関数のテストを実行します
func TestMonsterHandler_GetAll(t *testing.T) {
	tests := []struct {
		name       string
		mockSetup  func() monsters.IMonsterService
		query      string
		wantStatus int
		goldenFile string
	}{
		{
			name: "正常系: モンスター一覧が取得できる",
			mockSetup: func() monsters.IMonsterService {
				return &monsters.IMonsterServiceMock{
					FetchMonsterDetailFunc: func(ctx context.Context, id string) (*monsters.FetchMonsterListResult, error) {
						return &monsters.FetchMonsterListResult{
							Monsters: createTestMonsters(),
							Total:    100, // 総件数を返す（テストデータの2件ではなく、全体の100件）
						}, nil
					},
				}
			},
			query:      "",
			wantStatus: http.StatusOK,
			goldenFile: "monster/monster_get_all_200.json",
		},
		{
			name: "異常系: バリデーションエラー",
			mockSetup: func() monsters.IMonsterService {
				return &monsters.IMonsterServiceMock{
					FetchMonsterDetailFunc: func(ctx context.Context, id string) (*monsters.FetchMonsterListResult, error) {
						// このモックは呼び出されないはず
						return nil, nil
					},
				}
			},
			// limitに負の値を指定してバリデーションエラーを発生させる
			query:      "?limit=-1",
			wantStatus: http.StatusBadRequest,
			goldenFile: "monster/monster_get_all_400.json",
		},
		{
			name: "異常系: レコードが存在しない",
			mockSetup: func() monsters.IMonsterService {
				return &monsters.IMonsterServiceMock{
					FetchMonsterDetailFunc: func(ctx context.Context, id string) (*monsters.FetchMonsterListResult, error) {
						return nil, gorm.ErrRecordNotFound
					},
				}
			},
			query:      "",
			wantStatus: http.StatusNotFound,
			goldenFile: "monster/monster_get_all_404.json",
		},
		{
			name: "異常系: 内部エラー",
			mockSetup: func() monsters.IMonsterService {
				return &monsters.IMonsterServiceMock{
					FetchMonsterDetailFunc: func(ctx context.Context, id string) (*monsters.FetchMonsterListResult, error) {
						return nil, errors.New("database error")
					},
				}
			},
			query:      "",
			wantStatus: http.StatusInternalServerError,
			goldenFile: "monster/monster_get_all_500.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.mockSetup()
			router := setupTestRouter(t, mock)

			req, err := http.NewRequest(http.MethodGet, "/monsters"+tt.query, nil)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
		})
	}
}

// TestMonsterHandler_GetById は、GetById関数のテストを実行します
func TestMonsterHandler_GetById(t *testing.T) {
	tests := []struct {
		name       string
		mockSetup  func() monsters.IMonsterService
		pathParam  string
		wantStatus int
		goldenFile string
	}{
		{
			name: "正常系: モンスター詳細が取得できる",
			mockSetup: func() monsters.IMonsterService {
				return &monsters.IMonsterServiceMock{
					FetchMonsterDetailFunc: func(ctx context.Context, id string) (*monsters.FetchMonsterListResult, error) {
						// IDが指定されていれば1件だけ返す
						return &monsters.FetchMonsterListResult{
							Monsters: createTestMonsters()[:1],
							Total:    1,
						}, nil
					},
				}
			},
			pathParam:  "1",
			wantStatus: http.StatusOK,
			goldenFile: "monster/monster_get_by_id_200.json",
		},
		{
			name: "異常系: 不正なID形式",
			mockSetup: func() monsters.IMonsterService {
				return &monsters.IMonsterServiceMock{
					FetchMonsterDetailFunc: func(ctx context.Context, id string) (*monsters.FetchMonsterListResult, error) {
						// 無効なIDのため、レコードが見つからないエラーを返す
						return nil, gorm.ErrRecordNotFound
					},
				}
			},
			pathParam:  "invalid-id",
			wantStatus: http.StatusNotFound,
			goldenFile: "monster/monster_get_by_id_404.json",
		},
		{
			name: "異常系: レコードが存在しない",
			mockSetup: func() monsters.IMonsterService {
				return &monsters.IMonsterServiceMock{
					FetchMonsterDetailFunc: func(ctx context.Context, id string) (*monsters.FetchMonsterListResult, error) {
						return nil, gorm.ErrRecordNotFound
					},
				}
			},
			pathParam:  "999",
			wantStatus: http.StatusNotFound,
			goldenFile: "monster/monster_get_by_id_404.json",
		},
		{
			name: "異常系: 内部エラー",
			mockSetup: func() monsters.IMonsterService {
				return &monsters.IMonsterServiceMock{
					FetchMonsterDetailFunc: func(ctx context.Context, id string) (*monsters.FetchMonsterListResult, error) {
						return nil, errors.New("database error")
					},
				}
			},
			pathParam:  "1",
			wantStatus: http.StatusInternalServerError,
			goldenFile: "monster/monster_get_by_id_500.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.mockSetup()
			router := setupTestRouter(t, mock)

			path := "/monsters"
			if tt.pathParam != "" {
				path = fmt.Sprintf("%s/%s", path, tt.pathParam)
			}

			req, err := http.NewRequest(http.MethodGet, path, nil)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
		})
	}
}

// createTestMonsters はテスト用のモンスターリストを作成します
func createTestMonsters() []*monsters.FetchMonsterListDto {
	fire := "火"
	thunder := "雷"

	bgm1 := music.NewMusic("1", "1", "森と火の戦い", "abcdefg")
	bgm2 := music.NewMusic("2", "2", "雷鳴の如く", "hijklmn")

	return []*monsters.FetchMonsterListDto{
		{
			Id:                 "1",
			Name:               "リオレウス",
			Description:        "天空を司る王者",
			Element:            &fire,
			Location:           []string{"森林", "砂漠"},
			Category:           "飛竜種",
			Title:              []string{"MH1", "MH2"},
			FirstWeak_Attack:   "頭部",
			FirstWeak_Element:  "雷",
			SecondWeak_Attack:  "翼",
			SecondWeak_Element: "龍",
			Weakness_attack: []monsters.Weakness_attack{
				{
					Slashing: "30",
					Blow:     "25",
					Bullet:   "20",
				},
			},
			Weakness_element: []monsters.Weakness_element{
				{
					Fire:    "0",
					Water:   "10",
					Thunder: "20",
					Ice:     "15",
					Dragon:  "15",
				},
			},
			Ranking: []monsters.Ranking{
				{
					Ranking:  "1",
					VoteYear: "2020",
				},
			},
			BGM: []music.Music{
				*bgm1,
			},
		},
		{
			Id:                 "2",
			Name:               "ジンオウガ",
			Description:        "雷を纏いし獣",
			Element:            &thunder,
			Location:           []string{"高地", "森林"},
			Category:           "牙獣種",
			Title:              []string{"MH3", "MH4"},
			FirstWeak_Attack:   "頭部",
			FirstWeak_Element:  "水",
			SecondWeak_Attack:  "尻尾",
			SecondWeak_Element: "氷",
			Weakness_attack: []monsters.Weakness_attack{
				{
					Slashing: "25",
					Blow:     "30",
					Bullet:   "20",
				},
			},
			Weakness_element: []monsters.Weakness_element{
				{
					Fire:    "5",
					Water:   "20",
					Thunder: "0",
					Ice:     "15",
					Dragon:  "10",
				},
			},
			Ranking: []monsters.Ranking{
				{
					Ranking:  "3",
					VoteYear: "2020",
				},
			},
			BGM: []music.Music{
				*bgm2,
			},
		},
	}
}
