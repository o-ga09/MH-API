package weapon

type Weapons struct {
	Total  int            `json:"total,omitempty"`
	Limit  int            `json:"limit,omitempty"`
	Offset int            `json:"offset,omitempty"`
	Weapon []ResponseJson `json:"monsters,omitempty"`
}

type Weapon struct {
	Weapon ResponseJson `json:"monster"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ResponseJson struct {
	Id            string `json:"monster_id,omitempty"`
	Name          string `json:"name,omitempty"`
	ImageUrl      string `json:"image_url,omitempty"`
	Rare          string `json:"rare,omitempty"`
	Attack        string `json:"attack,omitempty"`
	ElemantAttaxk string `json:"elemant_attaxk,omitempty"`
	Shapness      string `json:"shapness,omitempty"`
	Critical      string `json:"critical,omitempty"`
	Description   string `json:"description,omitempty"`
}
