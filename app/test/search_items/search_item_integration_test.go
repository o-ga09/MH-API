// ここを参考にテスト書いていく
// https://zenn.dev/media_engine/articles/testing-go-applications
// https://qiita.com/takehanKosuke/items/849c732c5892d149b50a
package test

import (
	"encoding/json"
	"mh-api/app/internal/controller/item"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/presenter"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
)

func TestGetItem(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	r, err := presenter.NewServer()
	if err != nil {
		t.Fatal(err)
	}

	expectedItem1 := item.Items{
		Total:  11,
		Limit:  100,
		Offset: 0,
		Item: []item.ResponseJson{
			{Id: "0000000001", ItemName: "回復薬"},
			{Id: "0000000002", ItemName: "回復薬グレート"},
			{Id: "0000000003", ItemName: "秘薬"},
			{Id: "0000000004", ItemName: "砥石"},
			{Id: "0000000005", ItemName: "おとし穴"},
			{Id: "0000000006", ItemName: "毒ビン"},
			{Id: "0000000007", ItemName: "麻痺ビン"},
			{Id: "0000000008", ItemName: "眠りビン"},
			{Id: "0000000009", ItemName: "爆弾"},
			{Id: "0000000010", ItemName: "大タル爆弾"},
			{Id: "0000000011", ItemName: "閃光玉"},
		},
	}

	expectedItem2 := item.Items{
		Total:  6,
		Limit:  10,
		Offset: 5,
		Item: []item.ResponseJson{
			{Id: "0000000006", ItemName: "毒ビン"},
			{Id: "0000000007", ItemName: "麻痺ビン"},
			{Id: "0000000008", ItemName: "眠りビン"},
			{Id: "0000000009", ItemName: "爆弾"},
			{Id: "0000000010", ItemName: "大タル爆弾"},
			{Id: "0000000011", ItemName: "閃光玉"},
		},
	}

	expectedItem3 := item.Items{
		Total:  3,
		Limit:  100,
		Offset: 0,
		Item: []item.ResponseJson{
			{Id: "0000000001", ItemName: "回復薬"},
			{Id: "0000000002", ItemName: "回復薬グレート"},
			{Id: "0000000003", ItemName: "秘薬"},
		},
	}

	expectedItem4 := item.Items{
		Total:  2,
		Limit:  100,
		Offset: 0,
		Item: []item.ResponseJson{
			{Id: "0000000001", ItemName: "回復薬"},
			{Id: "0000000002", ItemName: "回復薬グレート"},
		},
	}

	expectedItem5 := item.Items{
		Total:  2,
		Limit:  100,
		Offset: 0,
		Item: []item.ResponseJson{
			{Id: "0000000001", ItemName: "回復薬"},
			{Id: "0000000002", ItemName: "回復薬グレート"},
		},
	}

	expectedItem6 := item.MessageResponse{
		Message: "NOT FOUND",
	}

	cases := []struct {
		name            string
		path            string
		expected_status int
		expected_body   interface{}
	}{
		{name: "アイテムを複数件取得できる(limit , offset 指定なし)", path: "/v1/items", expected_status: 200, expected_body: expectedItem1},
		{name: "アイテムを複数件取得できる(limit=10 , offset=5)", path: "/v1/items?limit=10&offset=5", expected_status: 200, expected_body: expectedItem2},
		{name: "ItemIdを複数件指定して取得できる", path: "/v1/items?itemIds=0000000001,0000000002,0000000003", expected_status: 200, expected_body: expectedItem3},
		{name: "ItemNameを指定して取得できる", path: "/v1/items?itemName=回復薬", expected_status: 200, expected_body: expectedItem4},
		{name: "ItemNameKanaを指定して取得できる", path: "/v1/items?itemNameKana=カイフクヤク", expected_status: 200, expected_body: expectedItem5},
		{name: "取得結果が0件の場合、404になる", path: "/v1/items?itemIds=1111111111,2222222222,3333333333", expected_status: 404, expected_body: expectedItem6},
		// 500のケースは保留
		// {name: "どこかでエラーの場合、500になる", path: "/v1/items?itemIds=$$$select%%%%", expected_status: 500, expected_body: `{"item":"expected item"}`},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			// テスト用のレスポンスライターを作成します
			rr := httptest.NewRecorder()
			// リクエストを実行します
			r.ServeHTTP(rr, req)
			// レスポンスステータスが期待通りであることを確認します
			assert.Equal(t, tt.expected_status, rr.Code)
			// レスポンスボディが期待通りであることを確認します
			// ここでは、期待するレスポンスボディを`expected`に設定します
			assert.Equal(t, tt.expected_status, rr.Code)
			if tt.expected_status == 200 {
				jsonBytes := rr.Body.Bytes()
				data := new(item.Items)

				if err := json.Unmarshal(jsonBytes, data); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expected_body, *data)
			} else {
				jsonBytes := rr.Body.Bytes()
				data := new(item.MessageResponse)

				if err := json.Unmarshal(jsonBytes, data); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expected_body, *data)
			}
		})
	}
}

func TestGetItemById(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	r, err := presenter.NewServer()
	if err != nil {
		t.Fatal(err)
	}

	expectedItem1 := item.ResponseJson{
		Id:       "0000000001",
		ItemName: "回復薬",
	}

	expectedItem2 := item.MessageResponse{
		Message: "NOT FOUND",
	}

	cases := []struct {
		name            string
		path            string
		expected_status int
		expected_body   interface{}
	}{
		{name: "アイテムを1件取得できる", path: "/v1/items/0000000001", expected_status: 200, expected_body: expectedItem1},
		{name: "取得結果が0件の場合、404になる", path: "/v1/items/1111111111", expected_status: 404, expected_body: expectedItem2},
		// 500のケースは保留
		// {name: "どこかでエラーの場合、500になる", path: "/v1/items/2222222222", expected_status: 500, expected_body: expectedItem3},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			// テスト用のレスポンスライターを作成します
			rr := httptest.NewRecorder()
			// リクエストを実行します
			r.ServeHTTP(rr, req)
			// レスポンスステータスが期待通りであることを確認します
			// レスポンスボディが期待通りであることを確認します
			// ここでは、期待するレスポンスボディを`expected`に設定します
			assert.Equal(t, tt.expected_status, rr.Code)
			if tt.expected_status == 200 {
				jsonBytes := rr.Body.Bytes()
				data := new(item.ResponseJson)

				if err := json.Unmarshal(jsonBytes, data); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expected_body, *data)
			} else {
				jsonBytes := rr.Body.Bytes()
				data := new(item.MessageResponse)

				if err := json.Unmarshal(jsonBytes, data); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expected_body, *data)
			}
		})
	}
}

func TestGetItemByMonster(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	r, err := presenter.NewServer()
	if err != nil {
		t.Fatal(err)
	}

	expectedItem1 := item.ItemsByMonsterList{
		Total:  3,
		Limit:  100,
		Offset: 0,
		MonsterList: []item.MonsterList{
			{ItemId: "0000000001", ItemName: "回復薬", Monsters: []item.Monster{{MonsterId: "0000000001"}, {MonsterId: "0000000002"}, {MonsterId: "0000000003"}}},
			{ItemId: "0000000002", ItemName: "回復薬グレート", Monsters: []item.Monster{{MonsterId: "0000000001"}, {MonsterId: "0000000002"}, {MonsterId: "0000000003"}}},
			{ItemId: "0000000003", ItemName: "秘薬", Monsters: []item.Monster{{MonsterId: "0000000001"}, {MonsterId: "0000000002"}, {MonsterId: "0000000003"}}},
		},
	}

	expectedItem2 := item.MessageResponse{
		Message: "NOT FOUND",
	}

	cases := []struct {
		name            string
		path            string
		expected_status int
		expected_body   interface{}
	}{
		{name: "アイテムと取得可能なモンスターの一覧を取得できる", path: "/v1/items/monsters", expected_status: 200, expected_body: expectedItem1},
		{name: "取得結果が0件の場合、404になる", path: "/v1/items/monsters?itemIds=1111111111", expected_status: 404, expected_body: expectedItem2},
		// 500のケースは保留
		// {name: "どこかでエラーの場合、500になる", path: "/v1/items/monsters/2222222222", expected_status: 500, expected_body: expectedItem3},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			// テスト用のレスポンスライターを作成します
			rr := httptest.NewRecorder()
			// リクエストを実行します
			r.ServeHTTP(rr, req)
			// レスポンスステータスが期待通りであることを確認します
			// レスポンスボディが期待通りであることを確認します
			// ここでは、期待するレスポンスボディを`expected`に設定します
			assert.Equal(t, tt.expected_status, rr.Code)
			if tt.expected_status == 200 {
				jsonBytes := rr.Body.Bytes()
				data := new(item.ItemsByMonsterList)

				if err := json.Unmarshal(jsonBytes, data); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expected_body, *data)
			} else {
				jsonBytes := rr.Body.Bytes()
				data := new(item.MessageResponse)

				if err := json.Unmarshal(jsonBytes, data); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expected_body, *data)
			}
		})
	}
}

func TestGetItemByMonsterId(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	r, err := presenter.NewServer()
	if err != nil {
		t.Fatal(err)
	}

	expectedItem1 := item.ItemsByMonster{
		MonsterId:   "0000000001",
		MonsterName: "リオレウス",
		Items: []item.ResponseJson{
			{Id: "0000000001"},
			{Id: "0000000002"},
			{Id: "0000000003"},
		},
	}

	expectedItem2 := item.MessageResponse{
		Message: "NOT FOUND",
	}

	cases := []struct {
		name            string
		path            string
		expected_status int
		expected_body   interface{}
	}{
		{name: "指定したモンスターから取得できるアイテム一覧を取得できる", path: "/v1/items/monsters/0000000001", expected_status: 200, expected_body: expectedItem1},
		{name: "取得結果が0件の場合、404になる", path: "/v1/items/monsters/1111111111", expected_status: 404, expected_body: expectedItem2},
		// 500のケースは保留
		// {name: "どこかでエラーの場合、500になる", path: "/v1/items/monsters/2222222222", expected_status: 500, expected_body: expectedItem3},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			// テスト用のレスポンスライターを作成します
			rr := httptest.NewRecorder()
			// リクエストを実行します
			r.ServeHTTP(rr, req)
			// レスポンスステータスが期待通りであることを確認します
			// レスポンスボディが期待通りであることを確認します
			// ここでは、期待するレスポンスボディを`expected`に設定します
			assert.Equal(t, tt.expected_status, rr.Code)
			if tt.expected_status == 200 {
				jsonBytes := rr.Body.Bytes()
				data := new(item.ItemsByMonster)

				if err := json.Unmarshal(jsonBytes, data); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expected_body, *data)
			} else {
				jsonBytes := rr.Body.Bytes()
				data := new(item.MessageResponse)

				if err := json.Unmarshal(jsonBytes, data); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expected_body, *data)
			}
		})
	}
}