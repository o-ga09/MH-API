package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mh-api/api/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockMonsterService struct {
	MockFindMonsterById    func(ctx context.Context, id int) (entity.Monster, error)
	MockFindAllMonsters    func(ctx context.Context) (entity.Monsters, error)
}

func (m *mockMonsterService) FindMonsterById(ctx context.Context, id int) (entity.Monster, error) {
	return m.MockFindMonsterById(ctx, id)
}

func (m *mockMonsterService) FindAllMonsters(ctx context.Context) (entity.Monsters, error) {
	return m.MockFindAllMonsters(ctx)
}

func TestMonsterHandler_GetMonsterById(t *testing.T) {
	// モックのMonsterServiceを作成
	mockData1 := entity.Monster{
		Id: 1,
		Name: "ジンオウガ",
		Desc: "ジンオウガかっこいい",
		Location: "大社跡",
		Specify: "牙竜種",
		Weakness_attack: "10 10 10 10 10",
		Weakness_element: "10 10 10 10 10",
	}

	mockData2 := entity.Monster{
		Id: 2,
		Name: "タマミツネ",
		Desc: "男の娘",
		Location: "大社跡",
		Specify: "海竜種",
		Weakness_attack: "10 10 10 10 10",
		Weakness_element: "10 10 10 10 10",
	}
	mockService := &mockMonsterService{
		MockFindMonsterById: func(ctx context.Context, id int) (entity.Monster, error) {
			// テストケースごとに返すデータを設定
			if id == 1 {
				return mockData1, nil
			} else if id == 2 {
				return mockData2, nil
			} else {
				return entity.Monster{}, errors.New("Monster not found")
			}
		},
	}

	// MonsterHandlerを作成
	handler := NewMonsterHandler(mockService)

	//パラメータ生成
	param := gin.Param{Key: "id",Value: "1"}
	params := gin.Params{param}

	// テスト用のContextとGinのContextを作成
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	ginCtx, _ := gin.CreateTestContext(response)
	ginCtx.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/monster/1", nil)
	ginCtx.Params = params

	// テスト実行
	handler.GetMonsterById(ginCtx)

	// 構造体をJSONに変換
	resJson, err := json.Marshal(mockData1)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}

	// レスポンスを検証
	assert.Equal(t, http.StatusOK, ginCtx.Writer.Status())
	assert.Equal(t, string(resJson), response.Body.String())
}

func TestMonsterHandler_GetMonsterAll(t *testing.T) {
	// モックのMonsterServiceを作成
	mockData1 := entity.Monster{
		Id: 1,
		Name: "ジンオウガ",
		Desc: "ジンオウガかっこいい",
		Location: "大社跡",
		Specify: "牙竜種",
		Weakness_attack: "10 10 10 10 10",
		Weakness_element: "10 10 10 10 10",
	}

	mockData2 := entity.Monster{
		Id: 2,
		Name: "タマミツネ",
		Desc: "男の娘",
		Location: "大社跡",
		Specify: "海竜種",
		Weakness_attack: "10 10 10 10 10",
		Weakness_element: "10 10 10 10 10",
	}
	mockService := &mockMonsterService{
		MockFindAllMonsters: func(ctx context.Context) (entity.Monsters, error) {
			// テストケースごとに返すデータを設定
			return entity.Monsters{mockData1, mockData2}, nil
		},
	}

	// MonsterHandlerを作成
	handler := NewMonsterHandler(mockService)

	// テスト用のContextとGinのContextを作成
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	ginCtx, _ := gin.CreateTestContext(response)
	ginCtx.Request, _ = http.NewRequest(http.MethodGet, "api/v1/monsters", nil)

	// テスト実行
	handler.GetMonsterAll(ginCtx)

	// 構造体をJSONに変換
	resJson, err := json.Marshal(entity.Monsters{mockData1,mockData2})
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}

	// レスポンスを検証
	assert.Equal(t, http.StatusOK, ginCtx.Writer.Status())
	assert.Equal(t, string(resJson), response.Body.String())
}
