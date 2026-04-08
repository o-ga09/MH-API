package monster

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"mh-api/internal/domain/fields"
	"mh-api/internal/domain/monsters"
	"mh-api/internal/domain/music"
	Products "mh-api/internal/domain/products"
	"mh-api/internal/domain/ranking"
	Tribes "mh-api/internal/domain/tribes"
	"mh-api/internal/domain/weakness"
	"mh-api/pkg/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupTestRouter(t *testing.T, repo monsters.Repository) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	handler := NewMonsterHandler(repo)
	g := r.Group("/monsters")
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	return r
}

func TestMonsterHandler_GetAll(t *testing.T) {
	tests := []struct {
		name       string
		mockSetup  func() monsters.Repository
		query      string
		wantStatus int
		goldenFile string
	}{
		{
			name: "正常系: モンスター一覧が取得できる",
			mockSetup: func() monsters.Repository {
				return &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return &monsters.SearchResult{
							Monsters: createTestMonsters(),
							Total:    100,
						}, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
					},
				}
			},
			query:      "",
			wantStatus: http.StatusOK,
			goldenFile: "monster/monster_get_all_200.json",
		},
		{
			name: "異常系: バリデーションエラー",
			mockSetup: func() monsters.Repository {
				return &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
					},
				}
			},
			query:      "?limit=-1",
			wantStatus: http.StatusBadRequest,
			goldenFile: "monster/monster_get_all_400.json",
		},
		{
			name: "異常系: レコードが存在しない",
			mockSetup: func() monsters.Repository {
				return &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, gorm.ErrRecordNotFound
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
					},
				}
			},
			query:      "",
			wantStatus: http.StatusNotFound,
			goldenFile: "monster/monster_get_all_404.json",
		},
		{
			name: "異常系: 内部エラー",
			mockSetup: func() monsters.Repository {
				return &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, errors.New("database error")
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
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
			router := setupTestRouter(t, tt.mockSetup())
			req, err := http.NewRequest(http.MethodGet, "/monsters"+tt.query, nil)
			require.NoError(t, err)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
			testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
		})
	}
}

func TestMonsterHandler_GetById(t *testing.T) {
	tests := []struct {
		name       string
		mockSetup  func() monsters.Repository
		pathParam  string
		wantStatus int
		goldenFile string
	}{
		{
			name: "正常系: モンスター詳細が取得できる",
			mockSetup: func() monsters.Repository {
				return &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return createTestMonsters()[0], nil
					},
				}
			},
			pathParam:  "1",
			wantStatus: http.StatusOK,
			goldenFile: "monster/monster_get_by_id_200.json",
		},
		{
			name: "異常系: レコードが存在しない",
			mockSetup: func() monsters.Repository {
				return &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, gorm.ErrRecordNotFound
					},
				}
			},
			pathParam:  "invalid-id",
			wantStatus: http.StatusNotFound,
			goldenFile: "monster/monster_get_by_id_404.json",
		},
		{
			name: "異常系: 内部エラー",
			mockSetup: func() monsters.Repository {
				return &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
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
			router := setupTestRouter(t, tt.mockSetup())
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

func createTestMonsters() []*monsters.Monster {
	fire := "火"
	thunder := "雷"
	return []*monsters.Monster{
		{
			MonsterId:   "1",
			Name:        "リオレウス",
			Description: "天空を司る王者",
			Element:     &fire,
			Weakness: []*weakness.Weakness{
				{
					Slashing:          "30",
					Blow:              "25",
					Bullet:            "20",
					Fire:              "0",
					Water:             "10",
					Lightning:         "20",
					Ice:               "15",
					Dragon:            "15",
					FirstWeakAttack:   "頭部",
					SecondWeakAttack:  "翼",
					FirstWeakElement:  "雷",
					SecondWeakElement: "龍",
				},
			},
			Tribe: &Tribes.Tribe{Name_ja: "飛竜種"},
			Product: []*Products.Product{
				{Name: "MH1"},
				{Name: "MH2"},
			},
			Field: []*fields.Field{
				{Name: "森林"},
				{Name: "砂漠"},
			},
			Ranking: []*ranking.Ranking{
				{Ranking: "1", VoteYear: "2020"},
			},
			BGM: []*music.Music{
				{Name: "森と火の戦い", Url: "abcdefg"},
			},
		},
		{
			MonsterId:   "2",
			Name:        "ジンオウガ",
			Description: "雷を纏いし獣",
			Element:     &thunder,
			Weakness: []*weakness.Weakness{
				{
					Slashing:          "25",
					Blow:              "30",
					Bullet:            "20",
					Fire:              "5",
					Water:             "20",
					Lightning:         "0",
					Ice:               "15",
					Dragon:            "10",
					FirstWeakAttack:   "頭部",
					SecondWeakAttack:  "尻尾",
					FirstWeakElement:  "水",
					SecondWeakElement: "氷",
				},
			},
			Tribe: &Tribes.Tribe{Name_ja: "牙獣種"},
			Product: []*Products.Product{
				{Name: "MH3"},
				{Name: "MH4"},
			},
			Field: []*fields.Field{
				{Name: "高地"},
				{Name: "森林"},
			},
			Ranking: []*ranking.Ranking{
				{Ranking: "3", VoteYear: "2020"},
			},
			BGM: []*music.Music{
				{Name: "雷鳴の如く", Url: "hijklmn"},
			},
		},
	}
}
