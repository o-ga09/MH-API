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

type MockItemsService struct {
	GetAllItemsFunc func(ctx context.Context) (*itemService.ItemListResponseDTO, error)
}

func (m *MockItemsService) GetAllItems(ctx context.Context) (*itemService.ItemListResponseDTO, error) {
	if m.GetAllItemsFunc != nil {
		return m.GetAllItemsFunc(ctx)
	}
	return nil, errors.New("GetAllItemsFunc not implemented in mock service")
}

func setupRouter() *gin.Engine {
	r := gin.New()
	return r
}

func TestItemHandler_GetItems_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	expectedServiceResponse := &itemService.ItemListResponseDTO{
		Items: []itemService.ItemDTO{
			{ItemID: "1", ItemName: "Test Item 1"},
			{ItemID: "2", ItemName: "Test Item 2"},
		},
	}

	mockRepo := &MockItemRepository{
		FindAllFunc: func(ctx context.Context) (items.Items, error) {
			return items.Items{
				*items.NewItem("1", "Test Item 1", ""), // imageUrlは空文字で仮置き
				*items.NewItem("2", "Test Item 2", ""), // imageUrlは空文字で仮置き
			}, nil
		},
	}
	realService := itemService.NewService(mockRepo)
	itemCtrlWithMockedRepo := NewItemHandler(realService)

	r.GET("/v1/items/success", itemCtrlWithMockedRepo.GetItems)

	reqSuccess, _ := http.NewRequest(http.MethodGet, "/v1/items/success", nil)
	wSuccess := httptest.NewRecorder()
	r.ServeHTTP(wSuccess, reqSuccess)

	assert.Equal(t, http.StatusOK, wSuccess.Code)

	var actualResponse itemService.ItemListResponseDTO
	err := json.Unmarshal(wSuccess.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedServiceResponse, &actualResponse)
}

func TestItemHandler_GetItems_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	mockRepo := &MockItemRepository{
		FindAllFunc: func(ctx context.Context) (items.Items, error) {
			return nil, errors.New("service layer error") // サービス層（実際はリポジトリ）からのエラー
		},
	}
	realService := itemService.NewService(mockRepo)
	itemCtrlWithErrorService := NewItemHandler(realService)

	r.GET("/v1/items/error", itemCtrlWithErrorService.GetItems)

	reqError, _ := http.NewRequest(http.MethodGet, "/v1/items/error", nil)
	wError := httptest.NewRecorder()
	r.ServeHTTP(wError, reqError)

	assert.Equal(t, http.StatusInternalServerError, wError.Code)
	var errorResponse MessageResponse
	err := json.Unmarshal(wError.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Failed to get items", errorResponse.Message)
}

func TestItemHandler_GetItem_NotImplemented(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

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
