package item

type MessageRequest struct {
	Message string `form:"message"`
}

type Requestform struct {
	Req []Json `form:"req"`
}

type Json struct {
	ItemId string `form:"item_id,omitempty"`
	Name   string `form:"item_name,omitempty"`
}

type RequestParam struct {
	ItemIds      string `form:"itemIds" binding:"validateId,max=9999"`
	ItemName     string `form:"itemName" binding:"max=9999"`
	ItemNameKana string `form:"itemNameKana" binding:"max=9999"`
	Limit        int    `form:"limit" binding:"min=0,max=1000"`
	Offset       int    `form:"offset" binding:"min=0,max=1000"`
	Sort         int    `form:"sort" binding:"min=0,max=1"`
	Order        int    `form:"order" binding:"min=0,max=1"`
}
