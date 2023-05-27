package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Skip()
	// テスト対象の関数を呼び出し
	gin.SetMode(gin.TestMode)
	server, err := NewServer()

	// エラーチェック
	assert.NoError(t, err, "NewServer should not return an error")

	// serverがnilでないことを確認
	assert.NotNil(t, server, "NewServer should return a non-nil *gin.Engine instance")
}