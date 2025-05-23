package mysql

import (
	"context"
	"mh-api/internal/domain/weapons"
	weaponService "mh-api/internal/service/weapons"

	"gorm.io/gorm"
)

// WeaponQueryService は武器データに関するクエリサービスです。
type WeaponQueryService struct {
	db *gorm.DB
}

// NewWeaponQueryService は新しい WeaponQueryService をインスタンス化します。
func NewWeaponQueryService(db *gorm.DB) *WeaponQueryService {
	return &WeaponQueryService{db: db}
}

// FindWeaponsParams は武器検索時のパラメータです。
// Issue #83 で定義されたリクエストパラメータに対応します。
type FindWeaponsParams struct {
	Limit     *int    `json:"limit"`
	Offset    *int    `json:"offset"`
	Sort      *string `json:"sort"`
	Order     *int    `json:"order"` // 0: Asc, 1: Desc (仮)
	MonsterID *string `json:"monster_id"`
	Name      *string `json:"name"`
	NameKana  *string `json:"name_kana"`
	// 他にも必要なパラメータがあればここに追加
}

func (qs *WeaponQueryService) FindWeapons(ctx context.Context, params weaponService.SearchWeaponsParams) ([]*weapons.Weapon, int, error) {
	dummyWeapons := []*weapons.Weapon{}

	dummyTotalCount := len(dummyWeapons)

	start := 0
	if params.Offset != nil {
		start = *params.Offset
	}
	end := dummyTotalCount
	if params.Limit != nil {
		if start+*params.Limit < dummyTotalCount {
			end = start + *params.Limit
		}
	}

	if start > dummyTotalCount {
		return []*weapons.Weapon{}, dummyTotalCount, nil
	}
	if end > dummyTotalCount {
		end = dummyTotalCount
	}

	return dummyWeapons[start:end], dummyTotalCount, nil
}

// FindWeaponByID は指定されたIDで武器を1件検索します。(これはIssueにはないが、あると便利かもしれないので枠だけ)
// (注意: この初期実装ではダミーデータを返します)
func (qs *WeaponQueryService) FindWeaponByID(ctx context.Context, weaponID string) (*weapons.Weapon, error) {
	return nil, nil // 見つからない場合はnilを返す
}
