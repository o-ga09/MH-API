package armor

type GetArmorByIDRequest struct {
	ArmorID string `uri:"id" binding:"required"`
}