package item

type MessageRequest struct {
	Message string `json:"message"`
}

type RequestJson struct {
	Req []Json `json:"req"`
}

type Json struct {
	ItemId string `json:"item_id,omitempty"`
	Name   string `json:"item_name,omitempty"`
}

type RequestParam struct {
	MonsterIds   string `json:"monster_id"`
	ItemName     string `json:"item_name"`
	ItemNameKana string `json:"item_name_kana"`
	Limit        int
	Offset       int
	Sort         string
	Order        int
}
