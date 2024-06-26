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

func TestGetItemValidation(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	r, err := presenter.NewServer()
	if err != nil {
		t.Fatal(err)
	}

	expectedItem1 := item.Items{
		Total:  1,
		Limit:  1,
		Offset: 0,
		Item: []item.ResponseJson{
			{Id: "0000000001", ItemName: "回復薬"},
		},
	}

	expectedItem2 := item.Items{
		Total:  9,
		Limit:  1000,
		Offset: 2,
		Item: []item.ResponseJson{
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

	expectedNotFound := item.MessageResponse{
		Message: "NOT FOUND",
	}

	expectedBadRequest := item.MessageResponse{
		Message: "BAD REQUEST",
	}

	cases := []struct {
		name            string
		path            string
		expected_status int
		expected_body   interface{}
	}{
		{name: "アイテムを1件取得できる(limit=1 , offset=0)", path: "/v1/items?limit=1&offset=0", expected_status: 200, expected_body: expectedItem1},
		{name: "アイテムを複数件取得できる(limit=1000 , offset=2)", path: "/v1/items?limit=1000&offset=2", expected_status: 200, expected_body: expectedItem2},
		{name: "limitが異常値の場合(limit=-1)", path: "/v1/items?limit=-1&offset=0", expected_status: 400, expected_body: expectedBadRequest},
		{name: "limitが異常値の場合(limit=1001)", path: "/v1/items?limit=1001&offset=0", expected_status: 400, expected_body: expectedBadRequest},
		{name: "limitが0の場合、404が返る(limit=0)", path: "/v1/items?limit=0&offset=0", expected_status: 404, expected_body: expectedNotFound},
		{name: "offsetが異常値の場合(offset=-1)", path: "/v1/items?limit=1&offset=-1", expected_status: 400, expected_body: expectedBadRequest},
		{name: "offsetが異常値の場合(offset=1001)", path: "/v1/items?limit=1&offset=1001", expected_status: 400, expected_body: expectedBadRequest},
		{name: "sortが異常値の場合(sort=2)", path: "/v1/items?sort=2", expected_status: 400, expected_body: expectedBadRequest},
		{name: "orderが異常値の場合(order=2)", path: "/v1/items?order=2", expected_status: 400, expected_body: expectedBadRequest},
		{name: "itemIdsが異常値の場合(itemIds=1111111111,2222222222,3333333333)", path: "/v1/items?itemIds=1111111111111,2222222222222,333333333333333", expected_status: 400, expected_body: expectedBadRequest},
		{name: "itemNameが異常値の場合(itemName=回復薬%%##@@)", path: "/v1/items?itemName=回復薬%%##@@", expected_status: 404, expected_body: expectedNotFound},
		{name: "itemNameKanaが異常値の場合(itemNameKana=カイフクヤク%%##@@)", path: "/v1/items?itemNameKana=カイフクヤク%%##@@", expected_status: 404, expected_body: expectedNotFound},
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
