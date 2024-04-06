package music

type MessageRequest struct {
	Message string `json:"message"`
}

type RequestJson struct {
	Req []Json `json:"req"`
}

type Json struct {
	BgmId string `json:"monster_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Url   string `json:"url,omitempty"`
}

type RequestParam struct {
	BgmIds      string `json:"monster_id,omitempty"`
	BgmName     string `json:"name,omitempty"`
	BgmNameKana string `json:"name_kana,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	Offset      int    `json:"offset,omitempty"`
	Sort        string `json:"sort,omitempty"`
	Order       int    `json:"order,omitempty"`
}
