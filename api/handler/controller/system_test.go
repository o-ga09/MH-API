package controller

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSystemHandler_Health(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		s    *SystemHandler
		args args
	}{
		{"ケース1",&SystemHandler{},args{c: &gin.Context{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SystemHandler{}
			s.Health(tt.args.c)
		})
	}
}

func TestNewSystemHandler(t *testing.T) {
	tests := []struct {
		name string
		want *SystemHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSystemHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSystemHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
