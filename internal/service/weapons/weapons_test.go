package weapons

import (
	"context"
	"errors"
	"mh-api/internal/domain/weapons"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// テストケース1: 正常系 - パラメータ指定なしでWeaponsを取得できる
func TestWeaponService_SearchWeapons_Success(t *testing.T) {
	// テスト準備
	ctx := t.Context()
	mockQueryService := &IWeaponQueryServiceMock{
		FindWeaponsFunc: func(ctx context.Context, params SearchWeaponsParams) ([]*weapons.Weapon, int, error) {
			dummyWeapons := []*weapons.Weapon{
				weapons.NewWeapon("1", "テスト武器1", "http://example.com/weapon1.jpg", "8", "100", "火10", "青40", "10%", "テスト武器1の説明"),
				weapons.NewWeapon("2", "テスト武器2", "http://example.com/weapon2.jpg", "9", "120", "水15", "白20", "15%", "テスト武器2の説明"),
			}
			return dummyWeapons, len(dummyWeapons), nil
		},
	}

	// サービスの初期化
	weaponService := NewWeaponService(mockQueryService)

	// テスト実行
	params := SearchWeaponsParams{}
	result, err := weaponService.SearchWeapons(ctx, params)

	// アサーション
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.TotalCount)
	assert.Equal(t, 2, len(result.Weapons))
	assert.Equal(t, "テスト武器1", result.Weapons[0].Name)
	assert.Equal(t, "テスト武器2", result.Weapons[1].Name)
	assert.Equal(t, 0, result.Offset)
	assert.Equal(t, 0, result.Limit)
}

// テストケース2: 正常系 - Limit, Offsetを指定してWeaponsを取得できる
func TestWeaponService_SearchWeapons_WithPagination(t *testing.T) {
	// テスト準備
	ctx := t.Context()
	mockQueryService := &IWeaponQueryServiceMock{
		FindWeaponsFunc: func(ctx context.Context, params SearchWeaponsParams) ([]*weapons.Weapon, int, error) {
			dummyWeapons := []*weapons.Weapon{
				weapons.NewWeapon("3", "テスト武器3", "http://example.com/weapon3.jpg", "7", "90", "雷20", "緑50", "0%", "テスト武器3の説明"),
			}
			// 実際にはクエリサービスでLimitとOffsetを使った処理が行われる想定
			return dummyWeapons, 10, nil // 10はトータル件数
		},
	}

	// サービスの初期化
	weaponService := NewWeaponService(mockQueryService)

	// テストパラメータの設定
	limit := 5
	offset := 5
	params := SearchWeaponsParams{
		Limit:  &limit,
		Offset: &offset,
	}

	// テスト実行
	result, err := weaponService.SearchWeapons(ctx, params)

	// アサーション
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 10, result.TotalCount)
	assert.Equal(t, 1, len(result.Weapons))
	assert.Equal(t, "テスト武器3", result.Weapons[0].Name)
	assert.Equal(t, 5, result.Offset)
	assert.Equal(t, 5, result.Limit)
}

// テストケース3: 正常系 - モンスターIDを指定してWeaponsを取得できる
func TestWeaponService_SearchWeapons_ByMonsterID(t *testing.T) {
	// テスト準備
	ctx := t.Context()
	mockQueryService := &IWeaponQueryServiceMock{
		FindWeaponsFunc: func(ctx context.Context, params SearchWeaponsParams) ([]*weapons.Weapon, int, error) {
			if params.WeaponID != nil && *params.WeaponID == "weapon_1" {
				dummyWeapons := []*weapons.Weapon{
					weapons.NewWeapon("4", "モンスター1の武器", "http://example.com/weapon4.jpg", "10", "150", "龍30", "紫10", "20%", "モンスター1の強力な武器"),
				}
				return dummyWeapons, len(dummyWeapons), nil
			}
			return []*weapons.Weapon{}, 0, nil
		},
	}

	// サービスの初期化
	weaponService := NewWeaponService(mockQueryService)

	// テストパラメータの設定
	weaponID := "weapon_1"
	params := SearchWeaponsParams{
		WeaponID: &weaponID,
	}

	// テスト実行
	result, err := weaponService.SearchWeapons(ctx, params)

	// アサーション
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.TotalCount)
	assert.Equal(t, 1, len(result.Weapons))
	assert.Equal(t, "モンスター1の武器", result.Weapons[0].Name)
}

// テストケース4: 正常系 - 武器名を指定して検索できる
func TestWeaponService_SearchWeapons_ByName(t *testing.T) {
	// テスト準備
	ctx := t.Context()
	mockQueryService := &IWeaponQueryServiceMock{
		FindWeaponsFunc: func(ctx context.Context, params SearchWeaponsParams) ([]*weapons.Weapon, int, error) {
			// パラメータチェック
			if params.Name != nil && *params.Name == "太刀" {
				dummyWeapons := []*weapons.Weapon{
					weapons.NewWeapon("5", "真・太刀", "http://example.com/weapon5.jpg", "9", "140", "氷20", "白30", "10%", "太刀の最終強化形態"),
					weapons.NewWeapon("6", "豪剣・太刀", "http://example.com/weapon6.jpg", "8", "130", "火15", "青50", "5%", "古の伝説の太刀"),
				}
				return dummyWeapons, len(dummyWeapons), nil
			}
			return []*weapons.Weapon{}, 0, nil
		},
	}

	// サービスの初期化
	weaponService := NewWeaponService(mockQueryService)

	// テストパラメータの設定
	name := "太刀"
	params := SearchWeaponsParams{
		Name: &name,
	}

	// テスト実行
	result, err := weaponService.SearchWeapons(ctx, params)

	// アサーション
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.TotalCount)
	assert.Equal(t, 2, len(result.Weapons))
	assert.Equal(t, "真・太刀", result.Weapons[0].Name)
	assert.Equal(t, "豪剣・太刀", result.Weapons[1].Name)
}

// テストケース5: 異常系 - クエリサービスでエラーが発生する場合
func TestWeaponService_SearchWeapons_QueryServiceError(t *testing.T) {
	// テスト準備
	ctx := t.Context()
	mockQueryService := &IWeaponQueryServiceMock{
		FindWeaponsFunc: func(ctx context.Context, params SearchWeaponsParams) ([]*weapons.Weapon, int, error) {
			return nil, 0, errors.New("データベース接続エラー")
		},
	}

	// サービスの初期化
	weaponService := NewWeaponService(mockQueryService)

	// テスト実行
	params := SearchWeaponsParams{}
	result, err := weaponService.SearchWeapons(ctx, params)

	// アサーション
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "データベース接続エラー", err.Error())
}
