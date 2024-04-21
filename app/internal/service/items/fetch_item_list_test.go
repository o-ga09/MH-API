package items

import (
	"reflect"
	"testing"
)

func TestNewFetchItemList(t *testing.T) {
	qs := ItemQueryServiceMock{}
	type args struct {
		qs ItemQueryService
	}
	tests := []struct {
		name string
		args args
		want *FetchItemList
	}{
		{name: "TestNewFetchItemList", args: args{qs: &qs}, want: NewFetchItemList(&qs)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFetchItemList(tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFetchItemList() = %v, want %v", got, tt.want)
			}
		})
	}
}
