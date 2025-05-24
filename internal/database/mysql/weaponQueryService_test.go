package mysql

import (
	"context"
	"errors"
	"testing"

	"mh-api/pkg/ptr"

	weaponService "mh-api/internal/service/weapons"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// テスト用の武器データを作成する関数
func createTestWeapons(t *testing.T, ctx context.Context) []*Weapon {
	testWeapons := []*Weapon{
		{
			WeaponID:      "weapon001",
			Name:          "炎剣リオレウス",
			ImageUrl:      "https://example.com/weapon001.png",
			Rarerity:      "5",
			Attack:        "100",
			ElementAttack: "30",
			Shapness:      "40",
			Critical:      "20",
			Description:   "炎の力を宿した剣",
		},
		{
			WeaponID:      "weapon002",
			Name:          "氷剣イヴェルカーナ",
			ImageUrl:      "https://example.com/weapon002.png",
			Rarerity:      "5",
			Attack:        "110",
			ElementAttack: "40",
			Shapness:      "30",
			Critical:      "25",
			Description:   "氷の力を宿した剣",
		},
		{
			WeaponID:      "weapon003",
			Name:          "雷剣ジンオウガ",
			ImageUrl:      "https://example.com/weapon003.png",
			Rarerity:      "5",
			Attack:        "105",
			ElementAttack: "35",
			Shapness:      "45",
			Critical:      "30",
			Description:   "雷の力を宿した剣",
		},
	}

	db := CtxFromTestDB(ctx)

	for _, weapon := range testWeapons {
		if err := db.Create(weapon).Error; err != nil {
			t.Fatalf("failed to create test weapon: %v", err)
		}
	}

	return testWeapons
}

func TestWeaponQueryService_FindWeapons(t *testing.T) {
	ctx := context.Background()

	// テストDBをセットアップ
	ctx = setupTestDB(ctx)
	testDB.Begin()
	defer testDB.Rollback()

	// テストデータ作成
	_ = createTestWeapons(t, ctx)

	// テストケース定義
	tests := []struct {
		name      string
		params    weaponService.SearchWeaponsParams
		wantCount int
		wantTotal int
		wantErr   bool
	}{
		{
			name: "正常系: すべての武器を取得",
			params: weaponService.SearchWeaponsParams{
				Limit: ptr.IntToPtr(10),
			},
			wantCount: 3,
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name: "正常系: 武器IDで絞り込み",
			params: weaponService.SearchWeaponsParams{
				WeaponID: ptr.StrToPtr("weapon001"),
			},
			wantCount: 1,
			wantTotal: 1,
			wantErr:   false,
		},
		{
			name: "正常系: 名前で部分一致検索",
			params: weaponService.SearchWeaponsParams{
				Name: ptr.StrToPtr("剣"),
			},
			wantCount: 3,
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name: "正常系: limitで件数制限",
			params: weaponService.SearchWeaponsParams{
				Limit: ptr.IntToPtr(1),
			},
			wantCount: 1,
			wantTotal: 1,
			wantErr:   false,
		},
		{
			name: "正常系: offsetでスキップ",
			params: weaponService.SearchWeaponsParams{
				Offset: ptr.IntToPtr(1),
				Limit:  ptr.IntToPtr(10),
			},
			wantCount: 2,
			wantTotal: 2,
			wantErr:   false,
		},
		{
			name: "正常系: 昇順ソート",
			params: weaponService.SearchWeaponsParams{
				Sort:  ptr.StrToPtr("asc"),
				Limit: ptr.IntToPtr(10),
			},
			wantCount: 3,
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name: "正常系: 降順ソート",
			params: weaponService.SearchWeaponsParams{
				Sort:  ptr.StrToPtr("desc"),
				Limit: ptr.IntToPtr(10),
			},
			wantCount: 3,
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name: "正常系: Orderによるソート(1:昇順)",
			params: weaponService.SearchWeaponsParams{
				Order: ptr.IntToPtr(1),
				Limit: ptr.IntToPtr(10),
			},
			wantCount: 3,
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name: "正常系: Orderによるソート(2:降順)",
			params: weaponService.SearchWeaponsParams{
				Order: ptr.IntToPtr(2),
				Limit: ptr.IntToPtr(10),
			},
			wantCount: 3,
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name: "正常系: 複合条件(武器ID + 名前)",
			params: weaponService.SearchWeaponsParams{
				WeaponID: ptr.StrToPtr("weapon001"),
				Name:     ptr.StrToPtr("炎"),
			},
			wantCount: 1,
			wantTotal: 1,
			wantErr:   false,
		},
		{
			name: "正常系: 検索結果0件",
			params: weaponService.SearchWeaponsParams{
				Name: ptr.StrToPtr("存在しない武器"),
			},
			wantCount: 0,
			wantTotal: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := NewWeaponQueryService()

			got, total, err := qs.FindWeapons(ctx, tt.params)

			assert.True(t, (err != nil) == tt.wantErr)
			if !tt.wantErr {
				assert.Len(t, got, tt.wantCount)
				assert.Equal(t, tt.wantTotal, total)
			}
		})
	}
}

func TestWeaponQueryService_FindWeaponByID(t *testing.T) {
	ctx := context.Background()

	// テストDBをセットアップ
	ctx = setupTestDB(ctx)
	testDB.Begin()
	defer testDB.Rollback()

	// テストデータ作成
	testWeapons := createTestWeapons(t, ctx)

	// テストケース定義
	tests := []struct {
		name     string
		weaponID string
		want     *Weapon
		wantErr  bool
		errType  error
	}{
		{
			name:     "正常系: 存在するIDの場合",
			weaponID: testWeapons[0].WeaponID,
			want:     testWeapons[0],
			wantErr:  false,
			errType:  nil,
		},
		{
			name:     "異常系: 存在しないIDの場合",
			weaponID: "non-existent-id",
			want:     nil,
			wantErr:  true,
			errType:  gorm.ErrRecordNotFound,
		},
		{
			name:     "異常系: 空のIDの場合",
			weaponID: "",
			want:     nil,
			wantErr:  true,
			errType:  gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// クエリサービスの初期化
			qs := NewWeaponQueryService()

			// テスト対象メソッド実行
			got, err := qs.FindWeaponByID(ctx, tt.weaponID)

			// アサーション
			if tt.wantErr {
				require.Error(t, err)
				if tt.errType != nil {
					assert.True(t, errors.Is(err, tt.errType), "expected error type: %v, got: %v", tt.errType, err)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want.WeaponID, got.GetID())
				assert.Equal(t, tt.want.Name, got.GetName())
			}
		})
	}
}
