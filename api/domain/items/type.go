package items

type ItemId struct{ value string }
type ItemName struct{ value string }
type ItemImageUrl struct{ value string }

func (f *ItemId) GetID() string        { return f.value }
func (f *ItemName) GetName() string    { return f.value }
func (f *ItemImageUrl) GetURL() string { return f.value }
