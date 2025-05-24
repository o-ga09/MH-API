package item

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"mh-api/internal/service/items"
	"mh-api/pkg/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockNewServer はテスト用のサーバーを初期化する関数
func MockNewServer(t *testing.T, itemHandler *ItemHandler) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	e := gin.New()

	return e
}

// SetupRouter はテスト用のルーターを設定する関数
func (h *ItemHandler) SetupRouter(r *gin.Engine) {
	r.GET("/v1/items", h.GetItems)
	r.GET("/v1/items/:itemId", h.GetItem)
	r.GET("/v1/items/monsters", h.GetItemByMonster)
}

func TestItemHandler_GetItems(t *testing.T) {
	// テストケースを定義
	tests := []struct {
		name       string
		input      map[string]any
		mock       func() *items.IitemServiceMock
		wantStatus int
		goldenFile string
	}{
		{
			name:  "正常系：アイテムの一覧が取得できる",
			input: map[string]any{},
			mock: func() *items.IitemServiceMock {
				mock := &items.IitemServiceMock{
					GetAllItemsFunc: func(ctx context.Context) (*items.ItemListResponseDTO, error) {
						return &items.ItemListResponseDTO{
							Items: []items.ItemDTO{
								{
									ItemID:   "1",
									ItemName: "回復薬",
								},
								{
									ItemID:   "2",
									ItemName: "回復薬グレート",
								},
							},
						}, nil
					},
				}
				return mock
			},
			wantStatus: http.StatusOK,
			goldenFile: "items/get_items_success.json",
		},
		{
			name:  "異常系：サービスからのエラーが発生",
			input: map[string]any{},
			mock: func() *items.IitemServiceMock {
				mock := &items.IitemServiceMock{
					GetAllItemsFunc: func(ctx context.Context) (*items.ItemListResponseDTO, error) {
						return nil, errors.New("service error")
					},
				}
				return mock
			},
			wantStatus: http.StatusInternalServerError,
			goldenFile: "items/get_items_error.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックの初期化
			mock := tt.mock()

			// ハンドラーの初期化
			itemHandler := NewItemHandler(mock)

			// ルーターの設定
			r := MockNewServer(t, itemHandler)
			itemHandler.SetupRouter(r)

			// リクエストの作成
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/v1/items", nil)

			// リクエストの実行
			r.ServeHTTP(w, req)

			// レスポンスの検証
			assert.Equal(t, tt.wantStatus, w.Code)

			// ゴールデンファイルとの比較
			if tt.goldenFile != "" {
				testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
			}
		})
	}
}

func TestItemHandler_GetItem(t *testing.T) {
	// テストケースを定義
	tests := []struct {
		name       string
		input      map[string]any
		mock       func() *items.IitemServiceMock
		wantStatus int
		goldenFile string
	}{
		{
			name: "正常系：アイテムが取得できる",
			input: map[string]any{
				"itemId": "1",
			},
			mock: func() *items.IitemServiceMock {
				mock := &items.IitemServiceMock{
					GetItemByIDFunc: func(ctx context.Context, itemID string) (*items.ItemDTO, error) {
						return &items.ItemDTO{
							ItemID:   "1",
							ItemName: "回復薬",
						}, nil
					},
				}
				return mock
			},
			wantStatus: http.StatusOK,
			goldenFile: "items/get_item_success.json",
		},
		{
			name: "異常系：アイテムが見つからない場合（404エラー）",
			input: map[string]any{
				"itemId": "invalid_path",
			},
			mock: func() *items.IitemServiceMock {
				mock := &items.IitemServiceMock{
					GetItemByIDFunc: func(ctx context.Context, itemID string) (*items.ItemDTO, error) {
						return nil, nil
					},
				}
				return mock
			},
			wantStatus: http.StatusNotFound,
			goldenFile: "items/get_item_not_found.json",
		},
		{
			name: "異常系：アイテムが見つからない場合",
			input: map[string]any{
				"itemId": "999",
			},
			mock: func() *items.IitemServiceMock {
				mock := &items.IitemServiceMock{
					GetItemByIDFunc: func(ctx context.Context, itemID string) (*items.ItemDTO, error) {
						return nil, nil
					},
				}
				return mock
			},
			wantStatus: http.StatusNotFound,
			goldenFile: "items/get_item_not_found.json",
		},
		{
			name: "異常系：サービスからのエラーが発生",
			input: map[string]any{
				"itemId": "1",
			},
			mock: func() *items.IitemServiceMock {
				mock := &items.IitemServiceMock{
					GetItemByIDFunc: func(ctx context.Context, itemID string) (*items.ItemDTO, error) {
						return nil, errors.New("service error")
					},
				}
				return mock
			},
			wantStatus: http.StatusInternalServerError,
			goldenFile: "items/get_item_error.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックの初期化
			mock := tt.mock()

			// ハンドラーの初期化
			itemHandler := NewItemHandler(mock)

			// ルーターの設定
			r := MockNewServer(t, itemHandler)
			itemHandler.SetupRouter(r)

			// リクエストの作成
			w := httptest.NewRecorder()
			var req *http.Request
			if itemId, ok := tt.input["itemId"].(string); ok {
				req, _ = http.NewRequest(http.MethodGet, "/v1/items/"+itemId, nil)
			} else {
				// アイテムIDがない場合は、ルーティングが合わないようにする
				req, _ = http.NewRequest(http.MethodGet, "/v1/items/invalid_path", nil)
			}

			// リクエストの実行
			r.ServeHTTP(w, req)

			// レスポンスの検証
			assert.Equal(t, tt.wantStatus, w.Code)

			// ゴールデンファイルとの比較
			if tt.goldenFile != "" {
				testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
			}
		})
	}
}

func TestItemHandler_GetItemByMonster(t *testing.T) {
	// テストケースを定義
	tests := []struct {
		name       string
		input      map[string]any
		mock       func() *items.IitemServiceMock
		wantStatus int
		goldenFile string
	}{
		{
			name: "正常系：モンスターIDに紐づくアイテムが取得できる",
			input: map[string]any{
				"monster_id": "1",
			},
			mock: func() *items.IitemServiceMock {
				mock := &items.IitemServiceMock{
					GetItemByMonsterIDFunc: func(ctx context.Context, monsterID string) (*items.ItemListResponseDTO, error) {
						return &items.ItemListResponseDTO{
							Items: []items.ItemDTO{
								{
									ItemID:   "1",
									ItemName: "回復薬",
								},
							},
						}, nil
					},
				}
				return mock
			},
			wantStatus: http.StatusOK,
			goldenFile: "items/get_item_by_monster_success.json",
		},
		{
			name:  "異常系：モンスターIDが空の場合",
			input: map[string]any{},
			mock: func() *items.IitemServiceMock {
				return &items.IitemServiceMock{}
			},
			wantStatus: http.StatusBadRequest,
			goldenFile: "items/get_item_by_monster_bad_request.json",
		},
		{
			name: "正常系：アイテムが空の場合",
			input: map[string]any{
				"monster_id": "999",
			},
			mock: func() *items.IitemServiceMock {
				mock := &items.IitemServiceMock{
					GetItemByMonsterIDFunc: func(ctx context.Context, monsterID string) (*items.ItemListResponseDTO, error) {
						return &items.ItemListResponseDTO{
							Items: []items.ItemDTO{},
						}, nil
					},
				}
				return mock
			},
			wantStatus: http.StatusOK,
			goldenFile: "items/get_item_by_monster_empty.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックの初期化
			mock := tt.mock()

			// ハンドラーの初期化
			itemHandler := NewItemHandler(mock)

			// ルーターの設定
			r := MockNewServer(t, itemHandler)
			itemHandler.SetupRouter(r)

			// リクエストの作成
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/v1/items/monsters", nil)

			// クエリパラメータの追加
			if monsterID, ok := tt.input["monster_id"]; ok {
				q := req.URL.Query()
				q.Add("monster_id", monsterID.(string))
				req.URL.RawQuery = q.Encode()
			}

			// リクエストの実行
			r.ServeHTTP(w, req)

			// レスポンスの検証
			assert.Equal(t, tt.wantStatus, w.Code)

			// ゴールデンファイルとの比較
			if tt.goldenFile != "" {
				testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
			}
		})
	}
}
