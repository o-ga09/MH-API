package music

type BGMs struct {
	Total  int            `json:"total,omitempty"`
	Limit  int            `json:"limit,omitempty"`
	Offset int            `json:"offset,omitempty"`
	BGM    []ResponseJson `json:"bgm,omitempty"`
}

type BGM struct {
	BGM ResponseJson `json:"monster"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ResponseJson struct {
	Id   string `json:"monster_id,omitempty"`
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}
