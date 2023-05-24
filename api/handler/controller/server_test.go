package controller

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		{"テストケース1",&gin.Engine{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got,_ := NewServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
