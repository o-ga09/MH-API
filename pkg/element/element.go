package element

import "strings"

// elementMap は英語の属性名を日本語に変換するマッピング
var elementMap = map[string]string{
	"fire":     "火",
	"water":    "水",
	"thunder":  "雷",
	"lightning": "雷",
	"ice":      "氷",
	"dragon":   "龍",
}

// NormalizeToJapanese は英語の属性名（Fire, Water 等）を日本語（火, 水 等）に正規化する。
// すでに日本語など未知の値はそのまま返す。
func NormalizeToJapanese(s string) string {
	if ja, ok := elementMap[strings.ToLower(s)]; ok {
		return ja
	}
	return s
}
