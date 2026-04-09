package skill

type MessageRequest struct {
	Message string `json:"message"`
}

type RequestSkillByID struct {
	SkillId string `uri:"skillId" form:"skillId" validate:"required" binding:"required"`
}

type SearchSkillsRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Limit       int    `form:"limit" binding:"omitempty,min=0"`
	Offset      int    `form:"offset" binding:"omitempty,min=0"`
}
