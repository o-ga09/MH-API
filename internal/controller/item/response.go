package item

import (
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
	MonsterId   string `json:"monster_id,omitempty"`
	MonsterName string `json:"monster_name,omitempty"`
	Item        []Item `json:"items,omitempty"`
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

func ToItemByMonsterResponse(item items.ItemByMonster) ItemsByMonster {
	resItem := make([]Item, len(item.Item))
	for i, it := range item.Item {
		resItem[i] = Item{
			Item: ResponseJson{
				Id:       it.ItemID,
				ItemName: it.ItemName,
			},
		}
	}

	return ItemsByMonster{
		MonsterId:   item.MonsterID,
		MonsterName: item.MonsterName,
		Item:        resItem,
	}
}
