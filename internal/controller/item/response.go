package item

import "mh-api/internal/domain/items"

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

func toItemListResponse(itemList items.Items) Items {
	res := make([]ResponseJson, len(itemList))
	for i, item := range itemList {
		res[i] = ResponseJson{
			Id:       item.ItemId,
			ItemName: item.Name,
		}
	}
	return Items{
		Total: len(itemList),
		Item:  res,
	}
}

func toItemSearchResponse(result *items.SearchResult, limit, offset int) Items {
	res := make([]ResponseJson, len(result.Items))
	for i, item := range result.Items {
		res[i] = ResponseJson{
			Id:       item.ItemId,
			ItemName: item.Name,
		}
	}
	return Items{
		Total:  result.Total,
		Limit:  limit,
		Offset: offset,
		Item:   res,
	}
}

func toItemResponse(item *items.Item) Item {
	return Item{
		Item: ResponseJson{
			Id:       item.ItemId,
			ItemName: item.Name,
		},
	}
}

func toItemByMonsterResponse(monsterId, monsterName string, itemList items.Items) ItemsByMonster {
	resItems := make([]Item, len(itemList))
	for i, it := range itemList {
		resItems[i] = Item{
			Item: ResponseJson{
				Id:       it.ItemId,
				ItemName: it.Name,
			},
		}
	}
	return ItemsByMonster{
		MonsterId:   monsterId,
		MonsterName: monsterName,
		Item:        resItems,
	}
}
