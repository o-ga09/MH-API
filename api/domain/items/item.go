package items

type Item struct {
	itemId   ItemId
	name     ItemName
	imageUrl ItemImageUrl
}

func newItem(itemId ItemId, name ItemName, imageUrl ItemImageUrl) *Item {
	return &Item{itemId, name, imageUrl}
}

func NewItem(itemId string, name string, imageUrl string) *Item {
	return newItem(
		ItemId{value: itemId},
		ItemName{value: name},
		ItemImageUrl{value: imageUrl},
	)
}
