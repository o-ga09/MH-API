package items

// Item アイテム情報
type ItemDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ItemsByMonster モンスター別アイテム情報
type ItemsByMonster struct {
	ItemId   string
	ItemName string
	Monster  []Monster
}

type Monster struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
