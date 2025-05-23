package weapon

// ListWeaponsResponse は武器リストAPIのレスポンス形式です。
// (旧 Weapons)
type ListWeaponsResponse struct {
	TotalCount int                    `json:"total_count,omitempty"` // Total -> TotalCount (JSONキーも変更)
	Limit      int                    `json:"limit,omitempty"`
	Offset     int                    `json:"offset,omitempty"`
	Weapons    []WeaponDetailResponse `json:"weapons,omitempty"`     // monsters -> weapons, 型も更新
}

// type Weapon struct {
// 	Weapon ResponseJson `json:"monster"`
// }
//
// type MessageResponse struct {
// 	Message string `json:"message"`
// }

// WeaponDetailResponse はAPIから返される個々の武器情報の詳細です。
// (旧 ResponseJson)
type WeaponDetailResponse struct {
	MonsterID     string `json:"monster_id,omitempty"`
	Name          string `json:"name,omitempty"`
	ImageURL      string `json:"image_url,omitempty"` // ImageUrl -> ImageURL
	Rare          string `json:"rare,omitempty"`
	Attack        string `json:"attack,omitempty"`
	ElementAttack string `json:"element_attack,omitempty"` // ElemantAttaxk -> ElementAttack
	Sharpness     string `json:"sharpness,omitempty"`    // Shapness -> Sharpness
	Critical      string `json:"critical,omitempty"`
	Description   string `json:"description,omitempty"`
}
