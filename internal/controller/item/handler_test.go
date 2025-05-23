package item

import (
	"context" // contextパッケージのインポートを追加
	"encoding/json"
	"errors"

	"mh-api/internal/domain/items"
	itemService "mh-api/internal/service/items"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockItemsService は items.Service のメソッドをモックするための構造体です。
// items.Service がインターフェースでないため、具体的なメソッドの振る舞いを差し替えます。
type MockItemsService struct {
	GetAllItemsFunc func(ctx context.Context) (*itemService.ItemListResponseDTO, error)
	// 他に items.Service にメソッドがあればここに追加
}

// items.Service のメソッドシグネチャに合わせてモックメソッドを実装
func (m *MockItemsService) GetAllItems(ctx context.Context) (*itemService.ItemListResponseDTO, error) {
	if m.GetAllItemsFunc != nil {
		return m.GetAllItemsFunc(ctx)
	}
	return nil, errors.New("GetAllItemsFunc not implemented in mock service")
}

func setupRouter() *gin.Engine {
	// テスト用にGinルーターをセットアップ
	// middleware.WithDB() など、実際のDB接続を伴うミドルウェアは
	// テストの際にはモック化するか、テストに影響しないように注意が必要
	r := gin.New() // gin.Default() はロガーなどを含むため、テストでは gin.New() が適することも
	return r
}

func TestItemHandler_GetItems_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()
	// mockService := &MockItemsService{} // このモックは直接使用しない

	// ダミーのサービスレスポンス (期待値)
	expectedServiceResponse := &itemService.ItemListResponseDTO{
		Items: []itemService.ItemDTO{
			{ItemID: "1", ItemName: "Test Item 1"},
			{ItemID: "2", ItemName: "Test Item 2"},
		},
	}

	// モックリポジトリを設定 (サービス層のテストで使ったものと同様の構造)
	mockRepo := &MockItemRepository{
		FindAllFunc: func(ctx context.Context) (items.Items, error) {
			// ドメインオブジェクトのリストを返す
			return items.Items{
				*items.NewItem("1", "Test Item 1", ""), // imageUrlは空文字で仮置き
				*items.NewItem("2", "Test Item 2", ""), // imageUrlは空文字で仮置き
			}, nil
		},
	}
	// 実際の itemService.Service をモックリポジトリを使って初期化
	realService := itemService.NewService(mockRepo)
	itemCtrlWithMockedRepo := NewItemHandler(realService)

	// ルート登録
	r.GET("/v1/items/success", itemCtrlWithMockedRepo.GetItems)

	// リクエスト作成と実行
	reqSuccess, _ := http.NewRequest(http.MethodGet, "/v1/items/success", nil)
	wSuccess := httptest.NewRecorder()
	r.ServeHTTP(wSuccess, reqSuccess)

	// アサーション
	assert.Equal(t, http.StatusOK, wSuccess.Code)

	var actualResponse itemService.ItemListResponseDTO
	err := json.Unmarshal(wSuccess.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedServiceResponse, &actualResponse)
}

func TestItemHandler_GetItems_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	// エラーを返すモックリポジトリ
	mockRepo := &MockItemRepository{
		FindAllFunc: func(ctx context.Context) (items.Items, error) {
			return nil, errors.New("service layer error") // サービス層（実際はリポジトリ）からのエラー
		},
	}
	// 実際の itemService.Service をエラーを返すモックリポジトリを使って初期化
	realService := itemService.NewService(mockRepo)
	itemCtrlWithErrorService := NewItemHandler(realService)

	// ルート登録
	r.GET("/v1/items/error", itemCtrlWithErrorService.GetItems)

	// リクエスト作成と実行
	reqError, _ := http.NewRequest(http.MethodGet, "/v1/items/error", nil)
	wError := httptest.NewRecorder()
	r.ServeHTTP(wError, reqError)

	// アサーション
	assert.Equal(t, http.StatusInternalServerError, wError.Code)
	var errorResponse MessageResponse
	err := json.Unmarshal(wError.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Failed to get items", errorResponse.Message)
}

func TestItemHandler_GetItem_NotImplemented(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()
	// NewItemHandler には itemService.Service のインスタンスが必要。
	// GetItem は現時点ではサービスコールを行わない想定だが、nil安全のためダミーを渡す。
	// ただし、itemService.NewService(nil) のようにリポジトリがnilだと panic する可能性があるので注意。
	// ここでは、リポジトリをモックしたダミーサービスを渡すのが無難。
	dummyRepo := &MockItemRepository{}
	dummyService := itemService.NewService(dummyRepo)
	itemCtrl := NewItemHandler(dummyService)

	r.GET("/items/:itemId", itemCtrl.GetItem)

	req, _ := http.NewRequest(http.MethodGet, "/items/some-id", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
	var actualResponse MessageResponse
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Not Implemented", actualResponse.Message)
}

func TestItemHandler_GetItemByMonster_NotImplemented(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	dummyRepo := &MockItemRepository{}
	dummyService := itemService.NewService(dummyRepo)
	itemCtrl := NewItemHandler(dummyService)
	r.GET("/items/monsters", itemCtrl.GetItemByMonster)

	req, _ := http.NewRequest(http.MethodGet, "/items/monsters", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
	var actualResponse MessageResponse
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Not Implemented", actualResponse.Message)
}

// MockItemRepository は items.Repository インターフェースのモック実装です。
// (サービス層のテストで定義したものを再利用、または共通化するのが望ましい)
type MockItemRepository struct {
	FindAllFunc func(ctx context.Context) (items.Items, error)
	SaveFunc    func(ctx context.Context, m items.Item) error
	RemoveFunc  func(ctx context.Context, itemId string) error
}

func (m *MockItemRepository) FindAll(ctx context.Context) (items.Items, error) {
	if m.FindAllFunc != nil {
		return m.FindAllFunc(ctx)
	}
	return nil, errors.New("FindAllFunc not implemented in mock")
}

// Save と Remove はこのテストでは直接使われないが、インターフェースを満たすために定義
func (m *MockItemRepository) Save(ctx context.Context, i items.Item) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(ctx, i)
	}
	return errors.New("SaveFunc not implemented in mock")
}

func (m *MockItemRepository) Remove(ctx context.Context, itemId string) error {
	if m.RemoveFunc != nil {
		return m.RemoveFunc(ctx, itemId)
	}
	return errors.New("RemoveFunc not implemented in mock")
}
