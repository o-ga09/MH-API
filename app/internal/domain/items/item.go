package items

type Item struct {
	itemId   ItemId
	name     ItemName
	nameKana ItemNameKana
	imageUrl ItemImageUrl
}

func newItem(itemId ItemId, name ItemName, nameKana ItemNameKana, imageUrl ItemImageUrl) *Item {
	return &Item{itemId, name, nameKana, imageUrl}
}

func NewItem(itemId string, name string, nameKana string, imageUrl string) *Item {
	return newItem(
		ItemId{value: itemId},
		ItemName{value: name},
		ItemNameKana{value: nameKana},
		ItemImageUrl{value: imageUrl},
	)
}
