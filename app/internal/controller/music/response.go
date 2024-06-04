package music

type BGMs struct {
	Total  int            `json:"total,omitempty"`
	Limit  int            `json:"limit,omitempty"`
	Offset int            `json:"offset,omitempty"`
	BGM    []ResponseJson `json:"bgm,omitempty"`
}

type BGM struct {
	BGM ResponseJson `json:"bgm"`
}

type BGMRankings struct {
	Total   int                   `json:"total,omitempty"`
	Limit   int                   `json:"limit,omitempty"`
	Offset  int                   `json:"offset,omitempty"`
	Ranking []ResponseRankingJson `json:"ranking,omitempty"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ResponseJson struct {
	Id   string `json:"music_id,omitempty"`
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type ResponseRankingJson struct {
	BgmId   string    `json:"music_id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Url     string    `json:"url,omitempty"`
	Ranking []Ranking `json:"ranking,omitempty"`
}

type Ranking struct {
	Ranking  string `json:"ranking,omitempty"`
	VoteYear string `json:"vote_year,omitempty"`
}
