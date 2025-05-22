package weapons

import (
	"context"

	"github.com/o-ga09/MH-API/internal/database/mysql" // データベース層のmysqlパッケージをインポート
	"github.com/o-ga09/MH-API/internal/domain/weapons"  // ドメイン層のweaponsパッケージをインポート
)

// IWeaponQueryService はデータベース層の武器クエリサービスのインターフェースです。
//これにより、テスト時にモックを注入できます。
type IWeaponQueryService interface {
	FindWeapons(ctx context.Context, params mysql.FindWeaponsParams) ([]*weapons.Weapon, int, error)
	// FindWeaponByID(ctx context.Context, weaponID string) (*weapons.Weapon, error) // 必要であればこちらも
}

// WeaponService は武器に関するビジネスロジックを提供します。
type WeaponService struct {
	queryService IWeaponQueryService
}

// NewWeaponService は新しい WeaponService をインスタンス化します。
func NewWeaponService(qs IWeaponQueryService) *WeaponService {
	return &WeaponService{queryService: qs}
}

// SearchWeapons は武器を検索し、結果を ListWeaponsResponse DTO として返します。
func (s *WeaponService) SearchWeapons(ctx context.Context, params SearchWeaponsParams) (*ListWeaponsResponse, error) {
	// サービス層のSearchWeaponsParamsをデータベース層のFindWeaponsParamsに変換
	dbParams := mysql.FindWeaponsParams{
		Limit:     params.Limit,
		Offset:    params.Offset,
		Sort:      params.Sort,
		Order:     params.Order,
		MonsterID: params.MonsterID,
		Name:      params.Name,
		NameKana:  params.NameKana,
	}

	domainWeapons, totalCount, err := s.queryService.FindWeapons(ctx, dbParams)
	if err != nil {
		// エラーの種類に応じて適切なエラーハンドリングを行う
		// (例: ログ出力、カスタムエラー型への変換など)
		return nil, err // とりあえずそのまま返す
	}

	// ドメインオブジェクトのリストをレスポンスDTOのリストに変換
	weaponDataList := ToWeaponDataList(domainWeapons)

	// デフォルトのオフセットとリミットを設定 (もしnilの場合)
	currentOffset := 0
	if params.Offset != nil {
		currentOffset = *params.Offset
	}
	currentLimit := 0 // 0は「指定なし」または「デフォルト値」を示す。コントローラーで設定されるべきか、ここで定義すべきか検討。
	                  // DB層のダミー実装ではLimit未指定だと全件なので、ここではその挙動に合わせる。
	if params.Limit != nil {
		currentLimit = *params.Limit
	}


	return &ListWeaponsResponse{
		Weapons:    weaponDataList,
		TotalCount: totalCount,
		Offset:     currentOffset,
		Limit:      currentLimit, // 返却するLimitはリクエストされたもの。実際の取得件数ではない。
	}, nil
}

// GetWeaponByID のような単一取得のサービスメソッドも必要に応じてここに追加できます。
// func (s *WeaponService) GetWeaponByID(ctx context.Context, weaponID string) (*WeaponData, error) {
//     domainWeapon, err := s.queryService.FindWeaponByID(ctx, weaponID)
//     if err != nil {
//         return nil, err
//     }
//     if domainWeapon == nil {
//         return nil, nil // or a specific "not found" error
//     }
//     weaponData := ToWeaponData(domainWeapon)
//     return &weaponData, nil
// }
