package item

type Items struct {
	Total  int            `json:"total,omitempty"`
	Limit  int            `json:"limit,omitempty"`
	Offset int            `json:"offset,omitempty"`
	Item   []ResponseJson `json:"items,omitempty"`
}

type Item struct {
	Item ResponseJson `json:"item"`
}

type ItemsByMonsterList struct {
	ItemId   string    `json:"item_id,omitempty"`
	ItemName string    `json:"item_name,omitempty"`
	Monsters []Monster `json:"monsters,omitempty"`
}

type ItemsByMonster struct {
	MonsterId   string         `json:"monster_id,omitempty"`
	MonsterName string         `json:"monster_name,omitempty"`
	Items       []ResponseJson `json:"items,omitempty"`
}

type Monster struct {
	MonsterId   string `json:"monster_id,omitempty"`
	MonsterName string `json:"monster_name,omitempty"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ResponseJson struct {
	Id       string `json:"item_id,omitempty"`
	ItemName string `json:"item_name,omitempty"`
}
