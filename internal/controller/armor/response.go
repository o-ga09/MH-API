package armor

type SkillResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RequiredItemResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResistanceResponse struct {
	Fire      int `json:"fire"`
	Water     int `json:"water"`
	Lightning int `json:"lightning"`
	Ice       int `json:"ice"`
	Dragon    int `json:"dragon"`
}

type ArmorDetailResponse struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Skills     []SkillResponse        `json:"skills"`
	Slot       string                 `json:"slot"`
	Defense    int                    `json:"defense"`
	Resistance ResistanceResponse     `json:"resistance"`
	Required   []RequiredItemResponse `json:"required"`
}

type ListArmorsResponse struct {
	Armors []ArmorDetailResponse `json:"armors"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
