package weapons

import "github.com/o-ga09/MH-API/internal/domain/weapons" // ドメイン層のweaponsをインポート

// WeaponResponse はAPIから返される武器情報のレスポンス形式です。
// Issue #83 のレスポンスJSONに対応します。
// domain層のweapons.Weapon構造体を直接使用するか、ここで再定義するか検討。
// ここではドメイン層のものをラップせずに、API仕様に厳密に合わせたフィールドを持つことを想定します。
// ただし、今回はドメイン層のWeapon構造体が既にJSONタグを含んでいるため、それをそのまま利用することを考えます。
// その場合、この WeaponResponse は実質的に weapons.Weapon のエイリアスまたはそれを内包する形になります。
// ここでは、複数の武器情報を返すためのリストと、ページネーションのための総件数情報を持つ構造体を定義します。

type WeaponData struct {
	Attack         string `json:"attack"`
	Critical       string `json:"critical"`
	Description    string `json:"description"`
	ElementAttack  string `json:"element_attack"`
	ImageURL       string `json:"image_url"`
	MonsterID      string `json:"monster_id"` // 武器のID
	Name           string `json:"name"`
	Rare           string `json:"rare"`
	Sharpness      string `json:"sharpness"`
}

type ListWeaponsResponse struct {
	Weapons    []WeaponData `json:"weapons"` // Issueでは "weapon" 単数形だが、リストなので複数形 "weapons" が一般的。Issueのレスポンスは1件の例だったため、それに合わせるなら weapon: WeaponData とする。今回はリスト取得APIなので複数形を想定。
	TotalCount int          `json:"total_count"`
	Offset     int          `json:"offset"`
	Limit      int          `json:"limit"`
}

// SearchWeaponsParams は武器検索サービスへの入力パラメータです。
// これは controller から service へ渡される際に使われます。
type SearchWeaponsParams struct {
	Limit      *int
	Offset     *int
	Sort       *string
	Order      *int    // 0: Asc, 1: Desc (仮)
	MonsterID  *string // 武器のIDによるフィルタ
	Name       *string
	NameKana   *string
}

// この関数は domain.Weapon を WeaponData に変換します。
// ドメインオブジェクトを直接APIレスポンスとして返さない場合に必要です。
func ToWeaponData(weapon *weapons.Weapon) WeaponData { // weapon は internal/domain/weapons/weapon.go の Weapon 型
	return WeaponData{
		MonsterID:     weapon.MonsterId.Value(),
		Name:          weapon.Name.Value(),
		ImageURL:      weapon.ImageUrl.Value(),
		Rare:          weapon.Rare.Value(),
		Attack:        weapon.Attack.Value(),
		ElementAttack: weapon.ElementAttack.Value(),
		Sharpness:     weapon.Sharpness.Value(), // Sharpness フィールド (型は WeaponShapness) から Value() を呼び出す
		Critical:      weapon.Critical.Value(),
		Description:   weapon.Description.Value(),
	}
}

func ToWeaponDataList(domainWeapons []*weapons.Weapon) []WeaponData {
	weaponDataList := make([]WeaponData, len(domainWeapons))
	for i, dw := range domainWeapons {
		weaponDataList[i] = ToWeaponData(dw)
	}
	return weaponDataList
}
