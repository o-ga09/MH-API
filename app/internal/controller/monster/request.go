package monster

type MessageRequest struct {
	Message string `json:"message"`
}

type RequestJson struct {
	Req []Json `json:"req"`
}

type Json struct {
	MonsterId        string `json:"monster_id,omitempty"`
	Name             string `json:"name,omitempty"`
	Desc             string `json:"desc,omitempty"`
	Location         string `json:"location,omitempty"`
	Category         string `json:"category,omitempty"`
	Title            string `json:"title,omitempty"`
	Weakness_attack  string `json:"weakness_attack,omitempty"`
	Weakness_element string `json:"weakness_element,omitempty"`
}

type RequestParam struct {
	MonsterIds      string `json:"monster_id"`
	MonsterName     string `json:"name"`
	MonsterNameKana string `json:"name_kana"`
	Limit           int
	Offset          int
	Sort            string
	Order           int
}
