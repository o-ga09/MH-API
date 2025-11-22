package items

import (
	"context"
	"errors"
	"mh-api/internal/domain/items"
	"mh-api/internal/service/monsters"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_GetAllItems(t *testing.T) {
	tests := []struct {
		name       string
		mockFunc   func() (items.Items, error)
		want       *ItemListResponseDTO
		wantErr    bool
		checkCalls bool // モック関数の呼び出し回数を確認する
	}{
		{
			name: "正常系: アイテム一覧を取得できる",
			mockFunc: func() (items.Items, error) {
				return items.Items{
					*items.NewItem("0000000001", "ポーション", "images/potion.png"),
					*items.NewItem("0000000002", "グレートポーション", "images/great_potion.png"),
					*items.NewItem("0000000003", "メガポーション", "images/mega_potion.png"),
				}, nil
			},
			want: &ItemListResponseDTO{
				Items: []ItemDTO{
					{ItemID: "0000000001", ItemName: "ポーション"},
					{ItemID: "0000000002", ItemName: "グレートポーション"},
					{ItemID: "0000000003", ItemName: "メガポーション"},
				},
			},
			wantErr:    false,
			checkCalls: true,
		},
		{
			name: "正常系: アイテムが存在しない場合は空のリストが返る",
			mockFunc: func() (items.Items, error) {
				return items.Items{}, nil
			},
			want: &ItemListResponseDTO{
				Items: []ItemDTO{},
			},
			wantErr: false,
		},
		{
			name: "異常系: リポジトリでエラーが発生する",
			mockFunc: func() (items.Items, error) {
				return nil, errors.New("repository error")
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックリポジトリの作成
			mockRepo := &items.RepositoryMock{
				FindAllFunc: func(ctx context.Context) (items.Items, error) {
					return tt.mockFunc()
				},
			}
			mockMonster := &monsters.MonsterQueryServiceMock{
				FetchListFunc: func(ctx context.Context, monsterID string) (*monsters.FetchMonsterListResult, error) {
					return nil, nil
				},
			}

			// テスト対象のサービスを初期化
			service := NewService(mockMonster, mockRepo)

			// テスト実行
			got, err := service.GetAllItems(context.Background())

			// アサーション
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantErr {
				return
			}

			require.NotNil(t, got)
			assert.Equal(t, len(tt.want.Items), len(got.Items))
			for i, item := range got.Items {
				assert.Equal(t, tt.want.Items[i].ItemID, item.ItemID)
				assert.Equal(t, tt.want.Items[i].ItemName, item.ItemName)
			}

			// モックの呼び出しを確認
			if tt.checkCalls {
				assert.Equal(t, 1, len(mockRepo.FindAllCalls()))
			}
		})
	}
}

func TestService_GetItemByID(t *testing.T) {
	tests := []struct {
		name       string
		itemID     string
		mockFunc   func(itemID string) (*items.Item, error)
		want       *ItemDTO
		wantErr    bool
		checkCalls bool
	}{
		{
			name:   "正常系: 存在するIDの場合",
			itemID: "0000000001",
			mockFunc: func(itemID string) (*items.Item, error) {
				return items.NewItem("0000000001", "ポーション", "images/potion.png"), nil
			},
			want: &ItemDTO{
				ItemID:   "0000000001",
				ItemName: "ポーション",
			},
			wantErr:    false,
			checkCalls: true,
		},
		{
			name:   "異常系: 存在しないIDの場合",
			itemID: "non-existent-id",
			mockFunc: func(itemID string) (*items.Item, error) {
				return nil, errors.New("item not found")
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "異常系: リポジトリでエラーが発生する",
			itemID: "0000000001",
			mockFunc: func(itemID string) (*items.Item, error) {
				return nil, errors.New("repository error")
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックリポジトリの作成
			mockRepo := &items.RepositoryMock{
				FindByIDFunc: func(ctx context.Context, itemID string) (*items.Item, error) {
					return tt.mockFunc(itemID)
				},
			}
			mockMonster := &monsters.MonsterQueryServiceMock{
				FetchListFunc: func(ctx context.Context, monsterID string) (*monsters.FetchMonsterListResult, error) {
					return nil, nil
				},
			}

			// テスト対象のサービスを初期化
			service := NewService(mockMonster, mockRepo)

			// テスト実行
			got, err := service.GetItemByID(context.Background(), tt.itemID)

			// アサーション
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantErr {
				return
			}

			require.NotNil(t, got)
			assert.Equal(t, tt.want.ItemID, got.ItemID)
			assert.Equal(t, tt.want.ItemName, got.ItemName)

			// モックの呼び出しを確認
			if tt.checkCalls {
				calls := mockRepo.FindByIDCalls()
				assert.Equal(t, 1, len(calls))
				assert.Equal(t, tt.itemID, calls[0].ItemID)
			}
		})
	}
}

func TestService_GetItemByMonsterID(t *testing.T) {
	tests := []struct {
		name       string
		monsterID  string
		mockFunc   func(monsterID string) (items.Items, error)
		want       *ItemListResponseDTO
		wantErr    bool
		checkCalls bool
	}{
		{
			name:      "正常系: モンスターIDに関連するアイテムが存在する場合",
			monsterID: "0000000001",
			mockFunc: func(monsterID string) (items.Items, error) {
				return items.Items{
					*items.NewItem("0000000001", "ポーション", "images/potion.png"),
					*items.NewItem("0000000002", "グレートポーション", "images/great_potion.png"),
				}, nil
			},
			want: &ItemListResponseDTO{
				Items: []ItemDTO{
					{ItemID: "0000000001", ItemName: "ポーション"},
					{ItemID: "0000000002", ItemName: "グレートポーション"},
				},
			},
			wantErr:    false,
			checkCalls: true,
		},
		{
			name:      "正常系: モンスターIDに関連するアイテムが存在しない場合",
			monsterID: "0000000099",
			mockFunc: func(monsterID string) (items.Items, error) {
				return items.Items{}, nil
			},
			want: &ItemListResponseDTO{
				Items: []ItemDTO{},
			},
			wantErr: false,
		},
		{
			name:      "異常系: リポジトリでエラーが発生する",
			monsterID: "0000000001",
			mockFunc: func(monsterID string) (items.Items, error) {
				return nil, errors.New("repository error")
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックリポジトリの作成
			mockRepo := &items.RepositoryMock{
				FindByMonsterIDFunc: func(ctx context.Context, monsterID string) (items.Items, error) {
					return tt.mockFunc(monsterID)
				},
			}
			mockMonster := &monsters.MonsterQueryServiceMock{
				FetchListFunc: func(ctx context.Context, monsterID string) (*monsters.FetchMonsterListResult, error) {
					return &monsters.FetchMonsterListResult{
						Monsters: []*monsters.FetchMonsterListDto{
							{
								Id:   "0000000001",
								Name: "リオレウス",
							},
						},
						Total: 1,
					}, nil
				},
			}

			// テスト対象のサービスを初期化
			service := NewService(mockMonster, mockRepo)

			// テスト実行
			got, err := service.GetItemByMonsterID(context.Background(), tt.monsterID)

			// アサーション
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantErr {
				return
			}

			require.NotNil(t, got)
			assert.Equal(t, len(tt.want.Items), len(got.Item))
			for i, item := range got.Item {
				assert.Equal(t, tt.want.Items[i].ItemID, item.ItemID)
				assert.Equal(t, tt.want.Items[i].ItemName, item.ItemName)
			}

			// モックの呼び出しを確認
			if tt.checkCalls {
				calls := mockRepo.FindByMonsterIDCalls()
				assert.Equal(t, 1, len(calls))
				assert.Equal(t, tt.monsterID, calls[0].MonsterID)
			}
		})
	}
}
