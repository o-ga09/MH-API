package itemdomain

import monsterdomain "mh-api/app/internal/domain/monsterDomain"

type Item struct {
	itemId      ItemId
	monsterId   monsterdomain.MonsterId
	name        ItemName
	descripiton ItemDescription
	imageUrl    ItemImageUrl
}

func newItem(itemId ItemId, monster_id monsterdomain.MonsterId, name ItemName, ItemDescription ItemDescription, imageUrl ItemImageUrl) *Item {
	return &Item{itemId, monster_id, name, ItemDescription, imageUrl}
}

func NewItem(itemId string, monster_id string, name string, description string, imageUrl string) *Item {
	return newItem(
		ItemId{value: itemId},
		monsterdomain.NewMonsterId(monster_id),
		ItemName{value: name},
		ItemDescription{value: description},
		ItemImageUrl{value: imageUrl},
	)
}

type Items []Item

type ItemId struct{ value string }
type ItemName struct{ value string }
type ItemDescription struct{ value string }
type ItemImageUrl struct{ value string }

func (f *Item) ItemID() string          { return f.itemId.value }
func (f *Item) MonsterID() string       { return f.monsterId.GetID() }
func (f *Item) ItemName() string        { return f.name.value }
func (f *Item) ItemDescription() string { return f.descripiton.value }
func (f *Item) ItemImageURL() string    { return f.imageUrl.value }
