package weapon

// type MessageRequest struct {
// 	Message string `json:"message"`
// }
//
// type RequestJson struct {
// 	Req []Json `json:"req"`
// }
//
// type Json struct {
// 	MonsterId     string `json:"monster_id,omitempty"`
// 	Name          string `json:"name,omitempty"`
// 	ImageUrl      string `json:"image_url,omitempty"`
// 	Rare          string `json:"rare,omitempty"`
// 	Attack        string `json:"attack,omitempty"`
// 	ElemantAttaxk string `json:"elemant_attaxk,omitempty"`
// 	Shapness      string `json:"shapness,omitempty"`
// 	Critical      string `json:"critical,omitempty"`
// 	Description   string `json:"description,omitempty"`
// }

// SearchWeaponsRequest は武器検索APIのクエリパラメータを保持します。
type SearchWeaponsRequest struct {
	MonsterID *string `query:"monster_id"`
	Name      *string `query:"name"`
	NameKana  *string `query:"name_kana"`
	Limit     *int    `query:"limit"`
	Offset    *int    `query:"offset"`
	Sort      *string `query:"sort"`
	Order     *int    `query:"order"` // 例: 0=asc, 1=desc など具体的な仕様に応じてバリデーションを追加することも検討
}
