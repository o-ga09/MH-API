package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSystemHandler_Health(t *testing.T) {
	// テスト用のContextとGinのContextを作成
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	ginCtx, _ := gin.CreateTestContext(response)
	ginCtx.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/system/healthcheck", nil)

	// テスト実行
	handler := NewSystemHandler()
	handler.Health(ginCtx)

	// 構造体をJSONに変換
	jsonData := struct {
		Message string
	}{
		Message: "ok",
	}
	resJson, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}

	// レスポンスを検証
	assert.Equal(t, http.StatusOK, ginCtx.Writer.Status())
	assert.Equal(t, string(resJson), response.Body.String())
}
