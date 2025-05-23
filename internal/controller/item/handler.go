package item

import (
	"mh-api/internal/service/items"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	// 依存するサービスを itemService.Service に変更
	service *items.Service
}

func NewItemHandler(s *items.Service) *ItemHandler {
	return &ItemHandler{
		service: s,
	}
}

// GetItems godoc
// @Summary アイテム名の一覧を取得する
// @Description 全てのアイテム名とIDの一覧を取得する
// @Tags アイテム検索
// @Produce json
// @Success 200 {object} itemService.ItemListResponseDTO
// @Failure 500 {object} response.MessageResponse
// @Router /v1/items [get]
func (h *ItemHandler) GetItems(c *gin.Context) {
	itemsResponse, err := h.service.GetAllItems(c.Request.Context())
	if err != nil {
		// TODO: エラーの種類に応じてステータスコードを出し分ける (例: 404 Not Found)
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get items"})
		return
	}
	c.JSON(http.StatusOK, itemsResponse)
}

// GetItem godoc
// @Summary アイテム検索（1件）
// @Description アイテムを検索して、条件に合致するアイテムを1件取得する
// @Tags アイテム検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} Item
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /items/:itemId [get]
func (h *ItemHandler) GetItem(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, MessageResponse{Message: "Not Implemented"})
}

// GetItemByMonster godoc
// @Summary アイテム検索（モンスター別）
// @Description 指定のアイテムが取得可能なモンスターの一覧
// @Tags アイテム検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} ItemsByMonster
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /items/monsters [get]
func (h *ItemHandler) GetItemByMonster(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, MessageResponse{Message: "Not Implemented"})
}
