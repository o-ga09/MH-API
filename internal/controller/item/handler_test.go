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
			name: "正常系：NotImplemented が返される",
			input: map[string]any{
				"itemId": "1",
			},
			mock: func() *items.IitemServiceMock {
				return &items.IitemServiceMock{}
			},
			wantStatus: http.StatusNotImplemented,
			goldenFile: "items/get_item_not_implemented.json",
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
			req, _ := http.NewRequest(http.MethodGet, "/v1/items/"+tt.input["itemId"].(string), nil)

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
			name:  "正常系：NotImplemented が返される",
			input: map[string]any{},
			mock: func() *items.IitemServiceMock {
				return &items.IitemServiceMock{}
			},
			wantStatus: http.StatusNotImplemented,
			goldenFile: "items/get_item_by_monster_not_implemented.json",
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
