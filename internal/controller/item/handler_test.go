package item

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"mh-api/internal/domain/items"
	"mh-api/internal/domain/monsters"
	Tribes "mh-api/internal/domain/tribes"
	"mh-api/pkg/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func MockNewServer(t *testing.T, handler *ItemHandler) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func (h *ItemHandler) SetupRouter(r *gin.Engine) {
	r.GET("/v1/items", h.GetItems)
	r.GET("/v1/items/:itemId", h.GetItem)
	r.GET("/v1/items/monster/:monsterId", h.GetItemByMonster)
}

func TestItemHandler_GetItems(t *testing.T) {
	tests := []struct {
		name       string
		mock       func() (items.Repository, monsters.Repository)
		wantStatus int
		goldenFile string
	}{
		{
			name: "正常系：アイテムの一覧が取得できる",
			mock: func() (items.Repository, monsters.Repository) {
				itemRepo := &items.RepositoryMock{
					FindFunc: func(ctx context.Context, params items.SearchParams) (*items.SearchResult, error) {
						return &items.SearchResult{
							Items: items.Items{
								{ItemId: "1", Name: "回復薬"},
								{ItemId: "2", Name: "回復薬グレート"},
							},
							Total: 2,
						}, nil
					},
					FindAllFunc: func(ctx context.Context) (items.Items, error) {
						return nil, nil
					},
					FindByIDFunc: func(ctx context.Context, itemID string) (*items.Item, error) {
						return nil, nil
					},
					FindByMonsterIDFunc: func(ctx context.Context, monsterID string) (items.Items, error) {
						return nil, nil
					},
				}
				monsterRepo := &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
					},
				}
				return itemRepo, monsterRepo
			},
			wantStatus: http.StatusOK,
			goldenFile: "items/get_items_success.json",
		},
		{
			name: "異常系：リポジトリでエラーが発生",
			mock: func() (items.Repository, monsters.Repository) {
				itemRepo := &items.RepositoryMock{
					FindFunc: func(ctx context.Context, params items.SearchParams) (*items.SearchResult, error) {
						return nil, errors.New("repository error")
					},
					FindAllFunc: func(ctx context.Context) (items.Items, error) {
						return nil, nil
					},
					FindByIDFunc: func(ctx context.Context, itemID string) (*items.Item, error) {
						return nil, nil
					},
					FindByMonsterIDFunc: func(ctx context.Context, monsterID string) (items.Items, error) {
						return nil, nil
					},
				}
				monsterRepo := &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
					},
				}
				return itemRepo, monsterRepo
			},
			wantStatus: http.StatusInternalServerError,
			goldenFile: "items/get_items_error.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itemRepo, monsterRepo := tt.mock()
			handler := NewItemHandler(itemRepo, monsterRepo)
			r := MockNewServer(t, handler)
			handler.SetupRouter(r)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/v1/items", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
		})
	}
}

func TestItemHandler_GetItem(t *testing.T) {
	tests := []struct {
		name       string
		pathParam  string
		mock       func() (items.Repository, monsters.Repository)
		wantStatus int
		goldenFile string
	}{
		{
			name:      "正常系：アイテムが取得できる",
			pathParam: "1",
			mock: func() (items.Repository, monsters.Repository) {
				itemRepo := &items.RepositoryMock{
					FindAllFunc: func(ctx context.Context) (items.Items, error) {
						return nil, nil
					},
					FindByIDFunc: func(ctx context.Context, itemID string) (*items.Item, error) {
						return &items.Item{ItemId: "1", Name: "回復薬"}, nil
					},
					FindByMonsterIDFunc: func(ctx context.Context, monsterID string) (items.Items, error) {
						return nil, nil
					},
				}
				monsterRepo := &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
					},
				}
				return itemRepo, monsterRepo
			},
			wantStatus: http.StatusOK,
			goldenFile: "items/get_item_success.json",
		},
		{
			name:      "異常系：アイテムが見つからない場合",
			pathParam: "999",
			mock: func() (items.Repository, monsters.Repository) {
				itemRepo := &items.RepositoryMock{
					FindAllFunc: func(ctx context.Context) (items.Items, error) {
						return nil, nil
					},
					FindByIDFunc: func(ctx context.Context, itemID string) (*items.Item, error) {
						return nil, nil
					},
					FindByMonsterIDFunc: func(ctx context.Context, monsterID string) (items.Items, error) {
						return nil, nil
					},
				}
				monsterRepo := &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
					},
				}
				return itemRepo, monsterRepo
			},
			wantStatus: http.StatusNotFound,
			goldenFile: "items/get_item_not_found.json",
		},
		{
			name:      "異常系：リポジトリでエラーが発生",
			pathParam: "1",
			mock: func() (items.Repository, monsters.Repository) {
				itemRepo := &items.RepositoryMock{
					FindAllFunc: func(ctx context.Context) (items.Items, error) {
						return nil, nil
					},
					FindByIDFunc: func(ctx context.Context, itemID string) (*items.Item, error) {
						return nil, errors.New("repository error")
					},
					FindByMonsterIDFunc: func(ctx context.Context, monsterID string) (items.Items, error) {
						return nil, nil
					},
				}
				monsterRepo := &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
					},
				}
				return itemRepo, monsterRepo
			},
			wantStatus: http.StatusInternalServerError,
			goldenFile: "items/get_item_error.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itemRepo, monsterRepo := tt.mock()
			handler := NewItemHandler(itemRepo, monsterRepo)
			r := MockNewServer(t, handler)
			handler.SetupRouter(r)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/items/%s", tt.pathParam), nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
		})
	}
}

func TestItemHandler_GetItemByMonster(t *testing.T) {
	tests := []struct {
		name       string
		pathParam  string
		mock       func() (items.Repository, monsters.Repository)
		wantStatus int
		goldenFile string
	}{
		{
			name:      "正常系：モンスターIDに紐づくアイテムが取得できる",
			pathParam: "1",
			mock: func() (items.Repository, monsters.Repository) {
				itemRepo := &items.RepositoryMock{
					FindAllFunc: func(ctx context.Context) (items.Items, error) {
						return nil, nil
					},
					FindByIDFunc: func(ctx context.Context, itemID string) (*items.Item, error) {
						return nil, nil
					},
					FindByMonsterIDFunc: func(ctx context.Context, monsterID string) (items.Items, error) {
						return items.Items{
							{ItemId: "1", Name: "イャンクックの羽"},
							{ItemId: "2", Name: "イャンクックの鱗"},
						}, nil
					},
				}
				monsterRepo := &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return &monsters.Monster{
							MonsterId: "1",
							Name:      "イャンクック",
							Tribe:     &Tribes.Tribe{Name_ja: "鳥竜種"},
						}, nil
					},
				}
				return itemRepo, monsterRepo
			},
			wantStatus: http.StatusOK,
			goldenFile: "items/get_item_by_monster_success.json",
		},
		{
			name:      "異常系：アイテムが存在しない場合",
			pathParam: "999",
			mock: func() (items.Repository, monsters.Repository) {
				itemRepo := &items.RepositoryMock{
					FindAllFunc: func(ctx context.Context) (items.Items, error) {
						return nil, nil
					},
					FindByIDFunc: func(ctx context.Context, itemID string) (*items.Item, error) {
						return nil, nil
					},
					FindByMonsterIDFunc: func(ctx context.Context, monsterID string) (items.Items, error) {
						return nil, nil
					},
				}
				monsterRepo := &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
					},
				}
				return itemRepo, monsterRepo
			},
			wantStatus: http.StatusNotFound,
			goldenFile: "items/get_item_by_monster_empty.json",
		},
		{
			name:      "異常系：リポジトリでエラーが発生",
			pathParam: "1",
			mock: func() (items.Repository, monsters.Repository) {
				itemRepo := &items.RepositoryMock{
					FindAllFunc: func(ctx context.Context) (items.Items, error) {
						return nil, nil
					},
					FindByIDFunc: func(ctx context.Context, itemID string) (*items.Item, error) {
						return nil, nil
					},
					FindByMonsterIDFunc: func(ctx context.Context, monsterID string) (items.Items, error) {
						return nil, errors.New("repository error")
					},
				}
				monsterRepo := &monsters.RepositoryMock{
					FindAllFunc: func(ctx context.Context, params monsters.SearchParams) (*monsters.SearchResult, error) {
						return nil, nil
					},
					FindByIdFunc: func(ctx context.Context, id string) (*monsters.Monster, error) {
						return nil, nil
					},
				}
				return itemRepo, monsterRepo
			},
			wantStatus: http.StatusInternalServerError,
			goldenFile: "items/get_item_by_monster_error.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itemRepo, monsterRepo := tt.mock()
			handler := NewItemHandler(itemRepo, monsterRepo)
			r := MockNewServer(t, handler)
			handler.SetupRouter(r)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/items/monster/%s", tt.pathParam), nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
		})
	}
}
