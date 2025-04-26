package music

import (
	"context"
	param "mh-api/app/internal/controller/music"
	"mh-api/app/internal/domain/music"
	"mh-api/app/internal/driver/mysql"
	musicService "mh-api/app/internal/service/music"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm"
)

type contextKey string

const paramKey contextKey = "param"

func TestNewmusicRepository(t *testing.T) {
	t.Skip()
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *musicRepository
	}{
		{name: "TestNewmusicRepository", args: args{conn: conn}, want: &musicRepository{conn: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMusicRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewmusicRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_musicRepository_Save(t *testing.T) {
	t.Skip()
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
		m   music.Music
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Save music successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), m: *music.NewMusic("0000000001", "0000000001", "test", "test")}, wantErr: false},
		// Test case 2
		{name: "Save music with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), m: *music.NewMusic("@$%&^#%$&&*%*&)(*)()", "", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &musicRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("musicRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_musicRepository_Remove(t *testing.T) {
	t.Skip()
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx   context.Context
		bgmId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Remove music successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), bgmId: "0000000001"}, wantErr: false},
		// Test case 2
		{name: "Remove music with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), bgmId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &musicRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.bgmId); (err != nil) != tt.wantErr {
				t.Errorf("musicRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFetchList(t *testing.T) {
	t.Skip()
	t.Helper()
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	ctx := context.Background()
	conn := mysql.New(ctx)

	param1 := param.RequestParam{BgmIds: "", BgmName: "", Limit: 100, Offset: 0, Sort: "1"}
	param2 := param.RequestParam{BgmIds: "0000000001,0000000002", BgmName: "", Limit: 100, Offset: 0, Sort: "1"}
	param3 := param.RequestParam{BgmIds: "", BgmName: "リオレウス", Limit: 100, Offset: 0, Sort: "1"}
	param4 := param.RequestParam{BgmIds: "", BgmName: "", Limit: 100, Offset: 0, Sort: "1"}
	param5 := param.RequestParam{BgmIds: "", BgmName: "", Limit: 100, Offset: 0, Sort: "2"}
	param6 := param.RequestParam{BgmIds: "0000000001", BgmName: "リオレウス", Limit: 100, Offset: 0, Sort: "1"}

	wantBgm1 := []*musicService.FetchMusicListDto{
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1"},
		{Id: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2"},
		{Id: "0000000003", Name: "ティガレックスのテーマ", Url: "https://www.youtube.com/watch?v=3"},
	}
	wantBgm2 := []*musicService.FetchMusicListDto{
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1"},
		{Id: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2"},
	}
	wantBgm3 := []*musicService.FetchMusicListDto{
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1"},
	}
	wantBgm4 := []*musicService.FetchMusicListDto{
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1"},
		{Id: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2"},
		{Id: "0000000003", Name: "ティガレックスのテーマ", Url: "https://www.youtube.com/watch?v=3"},
	}
	wantBgm5 := []*musicService.FetchMusicListDto{
		{Id: "0000000003", Name: "ティガレックスのテーマ", Url: "https://www.youtube.com/watch?v=3"},
		{Id: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2"},
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1"},
	}
	wantBgm6 := []*musicService.FetchMusicListDto{
		{Id: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2"},
	}
	wantBgm7 := []*musicService.FetchMusicListDto{
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1"},
	}

	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		id    string
		param param.RequestParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*musicService.FetchMusicListDto
		wantErr bool
	}{
		{name: "DBからモンスターデータを複数件取得できる", fields: fields{conn: conn}, args: args{id: "", param: param1}, want: wantBgm1, wantErr: false},
		{name: "DBからモンスターデータをbgmIdを複数件指定して取得できる", fields: fields{conn: conn}, args: args{id: "", param: param2}, want: wantBgm2, wantErr: false},
		{name: "DBからモンスターの名前を部分一致検索で指定して取得できる", fields: fields{conn: conn}, args: args{id: "", param: param3}, want: wantBgm3, wantErr: false},
		{name: "DBからモンスターデータをbgmIdでソート（昇順）して取得できる", fields: fields{conn: conn}, args: args{id: "", param: param4}, want: wantBgm4, wantErr: false},
		{name: "DBからモンスターデータをbgmIdでソート（降順）して取得できる", fields: fields{conn: conn}, args: args{id: "", param: param5}, want: wantBgm5, wantErr: false},
		{name: "DBからモンスターデータをid指定で1件取得できる", fields: fields{conn: conn}, args: args{id: "0000000002", param: param.RequestParam{}}, want: wantBgm6, wantErr: false},
		{name: "DBからモンスターデータを取得できない場合、NotFoundErrorで返す", fields: fields{conn: conn}, args: args{id: "", param: param.RequestParam{}}, want: nil, wantErr: true},
		{name: "DBからモンスターデータをmusicIdとmusicNameを指定して取得できる", fields: fields{conn: conn}, args: args{id: "", param: param6}, want: wantBgm7, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx = context.WithValue(context.Background(), paramKey, tt.args.param)
			s := &musicQueryService{
				conn: tt.fields.conn,
			}
			got, err := s.FetchList(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("monsterQueryService.FetchList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); !tt.wantErr && diff != "" {
				t.Errorf("monsterQueryService.FetchList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchRank(t *testing.T) {
	t.Skip()
	t.Helper()
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	ctx := context.Background()
	conn := mysql.New(ctx)

	param1 := param.RequestParam{BgmIds: "", BgmName: "", Limit: 100, Offset: 0, Sort: "1"}
	param2 := param.RequestParam{BgmIds: "0000000001,0000000002", BgmName: "", Limit: 100, Offset: 0, Sort: "1"}
	param3 := param.RequestParam{BgmIds: "", BgmName: "リオレウス", Limit: 100, Offset: 0, Sort: "1"}
	param4 := param.RequestParam{BgmIds: "", BgmName: "", Limit: 100, Offset: 0, Sort: "1"}
	param5 := param.RequestParam{BgmIds: "", BgmName: "", Limit: 100, Offset: 0, Sort: "2"}
	param6 := param.RequestParam{BgmIds: "0000000001", BgmName: "リオレウス", Limit: 100, Offset: 0, Sort: "1"}

	wantBgm1 := []*musicService.FetchMusicRankingDto{
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1", Ranking: []*musicService.Ranking{{Rank: "1", VoteYear: "2024/3/12"}}},
		{Id: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2", Ranking: []*musicService.Ranking{{Rank: "2", VoteYear: "2024/3/12"}}},
		{Id: "0000000003", Name: "ティガレックスのテーマ", Url: "https://www.youtube.com/watch?v=3", Ranking: []*musicService.Ranking{{Rank: "3", VoteYear: "2024/3/12"}}},
	}
	wantBgm2 := []*musicService.FetchMusicRankingDto{
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1", Ranking: []*musicService.Ranking{{Rank: "1", VoteYear: "2024/3/12"}}},
		{Id: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2", Ranking: []*musicService.Ranking{{Rank: "2", VoteYear: "2024/3/12"}}},
	}
	wantBgm3 := []*musicService.FetchMusicRankingDto{
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1", Ranking: []*musicService.Ranking{{Rank: "1", VoteYear: "2024/3/12"}}},
	}
	wantBgm4 := []*musicService.FetchMusicRankingDto{
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1", Ranking: []*musicService.Ranking{{Rank: "1", VoteYear: "2024/3/12"}}},
		{Id: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2", Ranking: []*musicService.Ranking{{Rank: "2", VoteYear: "2024/3/12"}}},
		{Id: "0000000003", Name: "ティガレックスのテーマ", Url: "https://www.youtube.com/watch?v=3", Ranking: []*musicService.Ranking{{Rank: "3", VoteYear: "2024/3/12"}}},
	}
	wantBgm5 := []*musicService.FetchMusicRankingDto{
		{Id: "0000000003", Name: "ティガレックスのテーマ", Url: "https://www.youtube.com/watch?v=3", Ranking: []*musicService.Ranking{{Rank: "3", VoteYear: "2024/3/12"}}},
		{Id: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2", Ranking: []*musicService.Ranking{{Rank: "2", VoteYear: "2024/3/12"}}},
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1", Ranking: []*musicService.Ranking{{Rank: "1", VoteYear: "2024/3/12"}}},
	}
	wantBgm7 := []*musicService.FetchMusicRankingDto{
		{Id: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1", Ranking: []*musicService.Ranking{{Rank: "1", VoteYear: "2024/3/12"}}},
	}

	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		param param.RequestParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*musicService.FetchMusicRankingDto
		wantErr bool
	}{
		{name: "DBからモンスターデータを複数件取得できる", fields: fields{conn: conn}, args: args{param: param1}, want: wantBgm1, wantErr: false},
		{name: "DBからモンスターデータをbgmIdを複数件指定して取得できる", fields: fields{conn: conn}, args: args{param: param2}, want: wantBgm2, wantErr: false},
		{name: "DBからモンスターの名前を部分一致検索で指定して取得できる", fields: fields{conn: conn}, args: args{param: param3}, want: wantBgm3, wantErr: false},
		{name: "DBからモンスターデータをbgmIdでソート（昇順）して取得できる", fields: fields{conn: conn}, args: args{param: param4}, want: wantBgm4, wantErr: false},
		{name: "DBからモンスターデータをbgmIdでソート（降順）して取得できる", fields: fields{conn: conn}, args: args{param: param5}, want: wantBgm5, wantErr: false},
		{name: "DBからモンスターデータを取得できない場合、NotFoundErrorで返す", fields: fields{conn: conn}, args: args{param: param.RequestParam{}}, want: nil, wantErr: true},
		{name: "DBからモンスターデータをmusicIdとmusicNameを指定して取得できる", fields: fields{conn: conn}, args: args{param: param6}, want: wantBgm7, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx = context.WithValue(context.Background(), paramKey, tt.args.param)
			s := &musicQueryService{
				conn: tt.fields.conn,
			}
			got, err := s.FetchRank(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("monsterQueryService.FetchRank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); !tt.wantErr && diff != "" {
				t.Errorf("monsterQueryService.FetchRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
