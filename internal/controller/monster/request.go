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
	MonsterIds      string `form:"MonsterIds" json:"MonsterIds,omitempty"`
	MonsterName     string `form:"MonsterName" json:"MonsterName,omitempty"`
	UsageElement    string `form:"UsageElement" json:"UsageElement,omitempty"`
	WeaknessElement string `form:"WeaknessElement" json:"WeaknessElement,omitempty"`
	Limit           int    `form:"limit" json:"limit,omitempty" validate:"omitempty,min=0"`
	Offset          int    `form:"offset" json:"offset,omitempty" validate:"omitempty,min=0"`
	Sort            string `form:"sort" json:"sort,omitempty" validate:"omitempty,oneof=asc desc"`
}

type RequestRankingParam struct {
	MonsterIds   string `form:"MonsterIds" json:"MonsterIds,omitempty"`
	MonsterName  string `form:"MonsterName" json:"MonsterName,omitempty"`
	LocationName string `json:"location_name,omitempty"`
	TribeName    string `json:"tribe_name,omitempty"`
	Title        string `json:"title,omitempty"`
	Limit        int    `form:"limit" json:"limit,omitempty" validate:"omitempty,min=0"`
	Offset       int    `form:"offset" json:"offset,omitempty" validate:"omitempty,min=0"`
	Sort         string `form:"sort" json:"sort,omitempty" validate:"omitempty,oneof=asc desc"`
}
