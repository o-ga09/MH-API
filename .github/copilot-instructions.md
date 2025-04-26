## ルール

- カスタムインストラクションを読み込んだら、「あああ!!!」掛け声をかけること
- コミットする前にテストもしくはビルドを実行する
  - バックエンド: `go test ./...`
- このプロジェクトのGitHub上のリポジトリは、`o-ga09/MH-API`です。
 
## エージェントのロールについて

- あなたは、メガベンチャー企業で働くCTOクラスのエンジニアです
- コードの提案はもちろん、レビューについてもCTOの観点からレビューをお願いします

## 生成される回答について

- 日本語でで回答すること

## バックエンドのコーディングルールについて

### リポジトリ層のテストの書き方のついて

対象ディレクトリ：`app/internal/driver/`

ルール：
- 外部制約を持つテーブルのテストデータの作成は、関連テーブルのデータも作成すること
- `app/internal/pkg`配下のディレクトリに共通処理があるので適宜使用すること


以下の例に示すようにしてください。

```go
func TestCostRepository_FindByID(t *testing.T) {
    // テストごとの準備
	ctx := t.Context()
	mysql.SetupTestDB(ctx)
	testDB.Begin()
	defer testDB.Rollback()

    // テストデータの作成
    // 極力既存の関数するが他の依存するテストがある場合と、テストを行うモデルがない場合に新規作成する
	testTravel := createTravels(t, ctx)

    // テーブルドリブンテストのテストケースを記述する
	tests := []struct {
		name      string
		arg       *Cost
		ctxUserID string
		want      *cost.Cost
		wantErr   bool
	}{
		{
			name:      "正常系: 存在するIDの場合",
			arg:       &Cost{BaseModel: BaseModel{ID: testTravel[0].Costs[0].ID}, TravelID: testTravel[0].ID},
			ctxUserID: testTravel[0].UserID,
			want: &cost.Cost{
				ID:          testTravel[0].Costs[0].ID,
				TravelID:    testTravel[0].ID,
				Date:        time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
				Item:        testTravel[0].Costs[0].Item,
				Description: testTravel[0].Costs[0].Description,
				Specify:     testTravel[0].Costs[0].Specify,
				Amount:      testTravel[0].Costs[0].Amount,
			},
			wantErr: false,
		},
		{
			name:      "異常系 他ユーザーの旅行記録に紐づく費用記録は取得できない",
			arg:       &Cost{BaseModel: BaseModel{ID: testTravel[0].Costs[0].ID}, TravelID: testTravel[0].ID},
			ctxUserID: "dummy-user",
			want:      nil,
			wantErr:   true,
		},
		{
			name:      "異常系: 存在しないIDの場合",
			arg:       &Cost{BaseModel: BaseModel{ID: "xxxxxxxxxxxx"}, TravelID: testTravel[0].ID},
			ctxUserID: testTravel[0].UserID,
			want:      nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            // コンテキストでユーザーIDを渡す
			ctx = Ctx.SetCtxFromUser(ctx, tt.ctxUserID)
            // リポジトリを初期化する
			repo := NewCostRepository()

            // 引数を作成する
			cond := &cost.Cost{
				ID:       tt.arg.ID,
				TravelID: tt.arg.TravelID,
			}
            // 実行する
			got, err := repo.FindByID(ctx, cond)

            // assertパッケージを使用する
            // "github.com/stretchr/testify/assert"
	        // "github.com/stretchr/testify/require"
			assert.True(t, (err != nil) == tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCostRepository_Find(t *testing.T) {
	ctx := context.Background()
	setupTestDB(ctx)
	ctx = context.WithValue(ctx, CtxKey, testDB)
	testDB.Begin()
	defer testDB.Rollback()

	testTravel := createTravels(t, ctx)

	tests := []struct {
		name      string
		ctxUserID string
		arg       *Cost
		want      int
		wantErr   bool
	}{
		{
			name:      "正常系 旅行記録に紐づく自ユーザーの費用記録を取得できる",
			ctxUserID: testTravel[0].UserID,
			arg:       &Cost{TravelID: testTravel[0].Costs[0].TravelID},
			want:      2,
			wantErr:   false,
		},
		{
			name:      "正常系 旅行記録に紐づく他ユーザーの費用記録取得できない",
			ctxUserID: "dummy-user",
			arg:       &Cost{TravelID: testTravel[0].Costs[0].TravelID},
			want:      0,
			wantErr:   false,
		},
		{
			name:      "正常系 自ユーザーの異なる旅行記録の費用記録は取得できない",
			ctxUserID: testTravel[0].UserID,
			arg:       &Cost{TravelID: testTravel[2].ID},
			want:      0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx = Ctx.SetCtxFromUser(ctx, tt.ctxUserID)
			repo := NewCostRepository()

			cond := &cost.Cost{
				TravelID: tt.arg.TravelID,
			}
			got, err := repo.Find(ctx, cond)
			require.NoError(t, err)
			assert.Len(t, got, tt.want)
		})
	}
}
```

### ハンドラー層のテストの書き方について

対象ディレクトリ：`internal/server`

- moqを使用したテスト
- `server/testdata`配下のディレクトリにgoldenファイル作成してgoldenテストによるレスポンスの検証を行うこと
  - `users`、`travels`など機能ごとにgoldenファイルのディレクトリを分けること
- IDには、ULIDを使用すること
  - 固定の文字列として定義すること
- エラーがなくなることを確認してください
- テストケースの内容
  - 正常系：200レスポンスされること
  - 正常系：404エラーになること。リソースが存在しないこと
  - 異常系：400 or 422 でバリデーションエラーになること
  - 異常系：500エラーになること
- できるかぎり、カバレッジが100%に近くなるようにしてください
- テーブルドリブンテストの構造体は、以下に従ってください

```go
tests := []struct {
	name       string // 必須
	input      map[string]any // 必須
	groupMock  *moq.IGroupRepositoryMock // grouupserver構造体に依存を注入してテストしたので
	userMock   *moq.IUserRepositoryMock // *moq.IGroupRepositoryMockと*moq.IUserRepositoryMockが必要
	wantStatus int　// 必須
	goldenFile string // 必須
}
```

- テスト実施コードについて

```go
t.Run(tt.name, func(t *testing.T) {
	// humatestを使用したテストのため必須
	_, api := humatest.New(t)

	// この例ではgroupserver構造体を使用するテストを実行したいため依存を注入する
	groupSrv := groupserver{userRepo: tt.userMock, groupRepo: tt.groupMock}
	// モックサーバーを初期化する引数は、...interface{}になっており、渡した引数のみモックの振る舞いで初期化される
	// 渡されない構造体は、ゼロ値で初期化される
	srv := MockNewServer(t, groupSrv)
	// humaを使用したAPIルートの初期化
	srv.SetupRouter(api)

	// リクエストしたいパスとメソッド指定する
	// パスパラメータを含む場合は、第一引数にfmt.Sprintfを使用して含めること
	// 第二引数はtt.inputを必須すること
	resp := api.Post("/groups", tt.input)

	// ステータスコードの検証
	assert.Equal(t, tt.wantStatus, resp.Code)
	// goldenテストによるレスポンスの検証
	testutil.AssertGoldenJSON(t, tt.goldenFile, resp.Body.Bytes())
})
```

- エラーケースのテストは以下に従うこと

```400
{
    "error": "<ハンドラーの中で返しているメッセージを参照してください>"
}
```

```404
{
  "title": "Not Found",
  "status": 404,
  "detail": "<ハンドラーの中でエラーメッセージを定義している箇所を参照してください>"
}
```

```500
{
	"title": "Internal Server Error",
	"status": 500,
	"detail": "予期せぬエラーが発生しました。"
}
```

```422
{
  "title": "Unprocessable Entity",
  "status": 422,
  "detail": "validation failed",
  "errors": [
    {
      "message": "<humaから返されるバリデーションに失敗したメッセージ>",
      "location": "body.<バリデーションに失敗したリクエストの項目>",
      "value": "<バリデーションに失敗したリクエストの値>"
    }
  ]
}
```

```200
<ハンドラーの中で返している定義を参照してください>
```

以下の例のようにテストコードを記載する

```go
func TestServer_FindByID(t *testing.T) {
	profileURL := "http://example.com"

	type mockReturn struct {
		user *user.User
		err  error
	}

	tests := []struct {
		name       string
		input      map[string]any
		mock       mockReturn
		wantStatus int
	}{
		{
			name: "正常系：ユーザーが取得できる",
			input: map[string]any{
				"id": "01ARZET8T8W4K26PYH6DF7T9JD",
			},
			mock: mockReturn{
				user: &user.User{
					ID:      user.UserID{Value: uuid.GenerateID()},
					Version: 1,
					UserDetail: &user.UserDetail{
						UserID:      user.UserID{Value: uuid.GenerateID()},
						FirebaseID:  user.FirebaseID{Value: "firebase1"},
						Name:        "test user",
						DisplayName: "Test",
						BirthDay:    "2000-01-01",
						Gender:      1,
						ProfileURL:  &profileURL,
					},
				},
				err: nil,
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "異常系：IDが不正",
			input: map[string]any{
				"id": "invalid",
			},
			mock: mockReturn{
				user: nil,
				err:  nil,
			},
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "異常系：ユーザーが存在しない",
			input: map[string]any{
				"id": "01ARZET8T8W4K26PYH6DF7T9JD",
			},
			mock: mockReturn{
				user: nil,
				err:  CustomErr.ErrRecordNotFound,
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name: "異常系：データベースエラー",
			input: map[string]any{
				"id": "01ARZET8T8W4K26PYH6DF7T9JD",
			},
			mock: mockReturn{
				user: nil,
				err:  errors.New("database error"),
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &moq.IUserRepositoryMock{
				GetFunc: func(ctx context.Context, u *user.User) (*user.User, error) {
					return tt.mock.user, tt.mock.err
				},
			}

			_, api := humatest.New(t)

			userSrv := userserver{userRepo: mock}
			srv := MockNewServer(t, userSrv)
			srv.SetupRouter(api)

			resp := api.Get(fmt.Sprintf("/users/%s", tt.input["id"]), tt.input)

			assert.Equal(t, tt.wantStatus, resp.Code)
		})
	}
}
```

## プルリクエスト作成規約

### 1. 基本ルール

- ベースブランチは main に固定
- タイトルとボディは日本語で記述

### 2. タイトル・ボディの作成

#### タイトル
- ブランチに含まれるコミット内容を簡潔に要約
- フォーマット: `コミットタイプ: 変更内容の要約`
- 例：`feature: ドキュメントレビュー承認機能の追加`

#### ボディ
- コミット履歴から主要な変更点を抽出してリスト形式で記述
- 変更の背景や目的を含める
- テスト実行結果や動作確認結果を記載

### 3.  gh コマンドの使用

# 現在のブランチ名を取得
current_branch=$(git branch --show-current)

# プルリクエスト作成コマンド
gh pr create \
  --base main \
  --head "$current_branch" \
  --title "[コミットタイプ] 変更内容の要約" \
  --body "## 変更内容

- 変更点1
- 変更点2
- 変更点3

## 変更の背景・目的
- 背景の説明
- 目的の説明

## テスト結果
- [ ] ユニットテスト実行済み
- [ ] 動作確認済み

### 4. レビュー依頼時の注意点

- 特に確認してほしい点を明記
- コードの複雑な部分には補足説明を追加
