package monsters

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestNewFetchMonsterRankingService(t *testing.T) {
	qs := MonsterQueryServiceMock{}
	type args struct {
		qs MonsterQueryService
	}
	tests := []struct {
		name string
		args args
		want *FetchMonsterRankingService
	}{
		{name: "TestNewFetchMonsterRankingService", args: args{qs: &qs}, want: NewFetchMonsterRankingService(&qs)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFetchMonsterRankingService(tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFetchMonsterRankingService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchMonsterRankingService_Run(t *testing.T) {
	qs := MonsterQueryServiceMock{}
	type fields struct {
		qs MonsterQueryService
	}
	type args struct {
		ctx context.Context
	}
	type mockValue struct {
		res []*FetchMonsterRankingDto
		err error
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		mockValue mockValue
		want      []*FetchMonsterRankingDto
		wantErr   bool
	}{
		{name: "TestFetchMonsterRankingService_Run", fields: fields{qs: &qs}, args: args{ctx: context.Background()}, mockValue: mockValue{res: []*FetchMonsterRankingDto{}, err: errors.New("Err")}, want: []*FetchMonsterRankingDto{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FetchMonsterRankingService{
				qs: tt.fields.qs,
			}
			qs.FetchRankFunc = func(ctx context.Context) ([]*FetchMonsterRankingDto, error) {
				return tt.mockValue.res, tt.mockValue.err
			}
			got, err := f.Run(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchMonsterRankingService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchMonsterRankingService.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
