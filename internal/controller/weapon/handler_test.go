package weapon

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/o-ga09/MH-API/internal/service/weapons" // サービスDTOを参照するため
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockWeaponService は IWeaponService インターフェースのモック実装です。
type MockWeaponService struct {
	mock.Mock
}

func (m *MockWeaponService) SearchWeapons(ctx context.Context, params weapons.SearchWeaponsParams) (*weapons.ListWeaponsResponse, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*weapons.ListWeaponsResponse), args.Error(1)
}

func TestSearchWeapons_Success_EmptyResult(t *testing.T) {
	// Ginのテストモード設定
	gin.SetMode(gin.TestMode)

	// モックサービスの準備
	mockService := new(MockWeaponService)
	expectedServiceResponse := &weapons.ListWeaponsResponse{
		Weapons:    []weapons.WeaponData{}, // 空のリスト
		TotalCount: 0,
		Offset:     0,
		Limit:      0, // ダミー実装のサービスはLimitをそのまま返す想定
	}
	// モックの期待動作設定: SearchWeapons が任意の値で呼ばれたら、expectedServiceResponse と nilエラーを返す
	mockService.On("SearchWeapons", mock.Anything, mock.AnythingOfType("weapons.SearchWeaponsParams")).Return(expectedServiceResponse, nil)

	// ハンドラとルーターの設定
	handler := NewWeaponHandler(mockService)
	router := gin.New()
	router.GET("/v1/weapons", handler.SearchWeapons)

	// HTTPリクエストの作成
	req, _ := http.NewRequest(http.MethodGet, "/v1/weapons", nil)
	rr := httptest.NewRecorder()

	// リクエストの実行
	router.ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusOK, rr.Code)

	// レスポンスボディの検証 (コントローラー層で定義した ListWeaponsResponse を期待)
	var responseBody ListWeaponsResponse // controller/weapon/response.go の型
	err := json.Unmarshal(rr.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, 0, responseBody.TotalCount)
	assert.Empty(t, responseBody.Weapons)
	assert.Equal(t, 0, responseBody.Offset)
	// Limitはリクエストで指定していない場合、サービスがデフォルト値を返すか、
	// コントローラーが設定した値(今回は0)になる。サービスのダミー実装に合わせる。
	assert.Equal(t, 0, responseBody.Limit)


	// モックの呼び出し検証
	mockService.AssertExpectations(t)
}

func TestSearchWeapons_WithQueryParams_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockWeaponService)
	// ダミーのレスポンス。limit=5, offset=10 を反映することを期待
	expectedServiceResponse := &weapons.ListWeaponsResponse{
		Weapons:    []weapons.WeaponData{},
		TotalCount: 0, // ダミーなので0件
		Offset:     10,
		Limit:      5,
	}
	// params の Limit と Offset がポインタなので、それらが適切に渡されることを期待
	// ここでは mock.AnythingOfType を使って型のみチェック
	mockService.On("SearchWeapons", mock.Anything, mock.MatchedBy(func(params weapons.SearchWeaponsParams) bool {
		return params.Limit != nil && *params.Limit == 5 && params.Offset != nil && *params.Offset == 10
	})).Return(expectedServiceResponse, nil)


	handler := NewWeaponHandler(mockService)
	router := gin.New()
	router.GET("/v1/weapons", handler.SearchWeapons)

	req, _ := http.NewRequest(http.MethodGet, "/v1/weapons?limit=5&offset=10", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var responseBody ListWeaponsResponse
	err := json.Unmarshal(rr.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, 0, responseBody.TotalCount)
	assert.Empty(t, responseBody.Weapons)
	assert.Equal(t, 10, responseBody.Offset) // レスポンスにはリクエストされたoffsetが反映される
	assert.Equal(t, 5, responseBody.Limit)   // レスポンスにはリクエストされたlimitが反映される

	mockService.AssertExpectations(t)
}

// TODO: サービスがエラーを返す場合のテストケース
// TODO: リクエストパラメータのバインディングエラーのテストケース (例: limitに文字列)
