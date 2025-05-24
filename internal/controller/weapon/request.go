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
	WeaponID *string `form:"weapon_id"`
	Name     *string `form:"name"`
	Limit    *int    `form:"limit" binding:"omitempty,min=1"`
	Offset   *int    `form:"offset" binding:"omitempty,min=0"`
	Sort     *string `form:"sort"`
	Order    *int    `form:"order"`
}
