package testutil

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// AssertGoldenJSON は、テスト結果とゴールデンファイルの内容を比較します
func AssertGoldenJSON(t *testing.T, goldenFile string, actual []byte) {
	t.Helper()

	// ゴールデンファイルのパスを作成
	goldenPath := filepath.Join("testdata", goldenFile)

	// UPDATE_GOLDEN環境変数が設定されていれば、ゴールデンファイルを更新
	if os.Getenv("UPDATE_GOLDEN") == "true" {
		err := os.MkdirAll(filepath.Dir(goldenPath), 0755)
		require.NoError(t, err)

		err = os.WriteFile(goldenPath, actual, 0644)
		require.NoError(t, err)
	}

	// ゴールデンファイルを読み込む
	expected, err := os.ReadFile(goldenPath)
	require.NoError(t, err)

	// JSON形式に整形して比較
	var expectedJSON, actualJSON any
	err = json.Unmarshal(expected, &expectedJSON)
	require.NoError(t, err)

	err = json.Unmarshal(actual, &actualJSON)
	require.NoError(t, err)

	// 期待値と実際の値を比較
	require.Equal(t, expectedJSON, actualJSON)
}
