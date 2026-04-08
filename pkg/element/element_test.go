package element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeToJapanese(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "Fire → 火", input: "Fire", want: "火"},
		{name: "fire（小文字） → 火", input: "fire", want: "火"},
		{name: "FIRE（大文字） → 火", input: "FIRE", want: "火"},
		{name: "Water → 水", input: "Water", want: "水"},
		{name: "Thunder → 雷", input: "Thunder", want: "雷"},
		{name: "Lightning → 雷", input: "Lightning", want: "雷"},
		{name: "Ice → 氷", input: "Ice", want: "氷"},
		{name: "Dragon → 龍", input: "Dragon", want: "龍"},
		{name: "すでに日本語「火」はそのまま", input: "火", want: "火"},
		{name: "すでに日本語「水」はそのまま", input: "水", want: "水"},
		{name: "未知の値はそのまま", input: "Unknown", want: "Unknown"},
		{name: "空文字はそのまま", input: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeToJapanese(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
