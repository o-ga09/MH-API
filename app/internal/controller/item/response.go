package item

import "mh-api/app/internal/controller/monster"

type Items struct {
	Total  int            `json:"total,omitempty"`
	Limit  int            `json:"limit,omitempty"`
	Offset int            `json:"offset,omitempty"`
	Item   []ResponseJson `json:"items,omitempty"`
}

type Item struct {
	Item ResponseJson `json:"item"`
}

type ItemsByMonster struct {
	ItemId   string           `json:"item_id,omitempty"`
	ItemName string           `json:"item_name,omitempty"`
	Monsters monster.Monsters `json:"monsters,omitempty"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ResponseJson struct {
	Id       string `json:"item_id,omitempty"`
	ItemName string `json:"item_name,omitempty"`
}
