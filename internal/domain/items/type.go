package items

type Items []Item

type ItemId struct{ value string }
type ItemName struct{ value string }
type ItemImageUrl struct{ value string }

func (f *Item) GetID() string   { return f.itemId.value }
func (f *Item) GetName() string { return f.name.value }
func (f *Item) GetURL() string  { return f.imageUrl.value }
