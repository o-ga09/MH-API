package armor

type GetArmorByIDRequest struct {
	ArmorID string `uri:"id" binding:"required"`
}

type SearchArmorsRequest struct {
	Name      string `form:"name"`
	SkillName string `form:"skill_name"`
	Slot      string `form:"slot"`
	Limit     int    `form:"limit" binding:"omitempty,min=0"`
	Offset    int    `form:"offset" binding:"omitempty,min=0"`
	Sort      string `form:"sort" binding:"omitempty,oneof=asc desc"`
}
