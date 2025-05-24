package item

import (
	"mh-api/internal/service/items"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	service items.IitemService
}

func NewItemHandler(s items.IitemService) *ItemHandler {
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
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get items"})
		return
	}
	c.JSON(http.StatusOK, ToItemListResponse(*itemsResponse))
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
	itemID := c.Param("itemId")
	if itemID == "" {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Item ID is required"})
		return
	}

	item, err := h.service.GetItemByID(c.Request.Context(), itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get item"})
		return
	}

	if item == nil {
		c.JSON(http.StatusNotFound, MessageResponse{Message: "Item not found"})
		return
	}

	c.JSON(http.StatusOK, ToItemResponse(*item))
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
	monsterID := c.Query("monster_id")
	if monsterID == "" {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Monster ID is required"})
		return
	}

	item, err := h.service.GetItemByMonsterID(c.Request.Context(), monsterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get item by monster ID"})
		return
	}

	if item == nil {
		c.JSON(http.StatusNotFound, MessageResponse{Message: "Item not found"})
		return
	}

	c.JSON(http.StatusOK, ToItemListResponse(*item))
}
