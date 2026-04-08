package testutil

import (
	"bytes"
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
		err := os.MkdirAll(filepath.Dir(goldenPath), 0750) //nolint:gosec // G301: test data directory, 0750 is appropriate
		require.NoError(t, err)

		var pretty bytes.Buffer
		err = json.Indent(&pretty, actual, "", "  ")
		require.NoError(t, err)

		err = os.WriteFile(goldenPath, append(pretty.Bytes(), '\n'), 0600) //nolint:gosec // G306: test data file
		require.NoError(t, err)
	}

	// ゴールデンファイルを読み込む
	expected, err := os.ReadFile(goldenPath) // #nosec G304 -- path is constructed from test-controlled input
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
