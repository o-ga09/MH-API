package mysql

import (
	"context"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	// テスト全体の前処理
	setupTestDB(ctx)

	// テストの実行
	code := m.Run()

	// テストの終了
	os.Exit(code)
}
