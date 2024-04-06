package weapon

type MessageRequest struct {
	Message string `json:"message"`
}

type RequestJson struct {
	Req []Json `json:"req"`
}

type Json struct {
	MonsterId     string `json:"monster_id,omitempty"`
	Name          string `json:"name,omitempty"`
	ImageUrl      string `json:"image_url,omitempty"`
	Rare          string `json:"rare,omitempty"`
	Attack        string `json:"attack,omitempty"`
	ElemantAttaxk string `json:"elemant_attaxk,omitempty"`
	Shapness      string `json:"shapness,omitempty"`
	Critical      string `json:"critical,omitempty"`
	Description   string `json:"description,omitempty"`
}

type RequestParam struct {
	WeaponIds      string `json:"monster_id,omitempty"`
	WeaponName     string `json:"name,omitempty"`
	WeaponNameKana string `json:"name_kana,omitempty"`
	Limit          int    `json:"limit,omitempty"`
	Offset         int    `json:"offset,omitempty"`
	Sort           string `json:"sort,omitempty"`
	Order          int    `json:"order,omitempty"`
}
