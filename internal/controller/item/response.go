package item

import (
	"mh-api/internal/controller/monster"
	"mh-api/internal/service/items"
)

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

func ToItemListResponse(items items.ItemListResponseDTO) Items {
	res := make([]ResponseJson, len(items.Items))
	for i, item := range items.Items {
		res[i] = ResponseJson{
			Id:       item.ItemID,
			ItemName: item.ItemName,
		}
	}
	return Items{
		Total:  len(items.Items),
		Limit:  items.Limit,
		Offset: items.Offset,
		Item:   res,
	}
}

func ToItemResponse(item items.ItemDTO) Item {
	return Item{
		Item: ResponseJson{
			Id:       item.ItemID,
			ItemName: item.ItemName,
		},
	}
}
