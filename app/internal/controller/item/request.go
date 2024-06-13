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
	ItemIds      string `form:"itemIds,omitempty"`
	ItemName     string `form:"itemName,omitempty"`
	ItemNameKana string `form:"itemNameKana,omitempty"`
	Limit        int    `form:"limit,omitempty"`
	Offset       int    `form:"offset,omitempty"`
	Sort         int    `form:"sort,omitempty"`
	Order        int    `form:"order,omitempty"`
}
