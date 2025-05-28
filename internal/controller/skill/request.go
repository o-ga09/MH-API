package skill

type MessageRequest struct {
	Message string `json:"message"`
}

type RequestSkillByID struct {
	SkillId string `uri:"skillId" form:"skillId" validate:"required" binding:"required"`
}