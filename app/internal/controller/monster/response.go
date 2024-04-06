package monster

type Monsters struct {
	Total    int            `json:"total,omitempty"`
	Limit    int            `json:"limit,omitempty"`
	Offset   int            `json:"offset,omitempty"`
	Monsters []ResponseJson `json:"monsters,omitempty"`
}

type Monster struct {
	Monster ResponseJson `json:"monster"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ResponseJson struct {
	Id               string `json:"monster_id,omitempty"`
	Name             string `json:"name,omitempty"`
	Desc             string `json:"desc,omitempty"`
	Location         string `json:"location,omitempty"`
	Category         string `json:"category,omitempty"`
	Title            string `json:"title,omitempty"`
	Weakness_attack  string `json:"weakness_attack,omitempty"`
	Weakness_element string `json:"weakness_element,omitempty"`
}
