package item

import (
	"mh-api/internal/domain/items"
	"mh-api/internal/domain/monsters"
	"mh-api/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemRepo    items.Repository
	monsterRepo monsters.Repository
}

func NewItemHandler(itemRepo items.Repository, monsterRepo monsters.Repository) *ItemHandler {
	return &ItemHandler{
		itemRepo:    itemRepo,
		monsterRepo: monsterRepo,
	}
}

// GetItems godoc
// @Summary アイテム名の一覧を取得する
// @Description 全てのアイテム名とIDの一覧を取得する
// @Tags アイテム検索
// @Produce json
// @Success 200 {object} Items
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /items [get]
func (h *ItemHandler) GetItems(c *gin.Context) {
	itemList, err := h.itemRepo.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get items"})
		return
	}
	c.JSON(http.StatusOK, toItemListResponse(itemList))
}

// GetItem godoc
// @Summary アイテム検索（1件）
// @Description アイテムを検索して、条件に合致するアイテムを1件取得する
// @Tags アイテム検索
// @Accept json
// @Produce json
// @Param itemId path string true "アイテムID"
// @Success 200 {object} Item
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /items/{itemId} [get]
func (h *ItemHandler) GetItem(c *gin.Context) {
	var req RequestItemByID
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Invalid item ID"})
		return
	}

	validate := validator.GetValidator()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Validation failed: " + err.Error()})
		return
	}

	item, err := h.itemRepo.FindByID(c.Request.Context(), req.ItemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get item"})
		return
	}

	if item == nil {
		c.JSON(http.StatusNotFound, MessageResponse{Message: "Item not found"})
		return
	}

	c.JSON(http.StatusOK, toItemResponse(item))
}

// GetItemByMonster godoc
// @Summary アイテム検索（モンスター別）
// @Description 指定のアイテムが取得可能なモンスターの一覧
// @Tags アイテム検索
// @Accept json
// @Produce json
// @Param monsterId path string true "モンスターID"
// @Success 200 {object} ItemsByMonster
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /items/monsters/{monsterId} [get]
func (h *ItemHandler) GetItemByMonster(c *gin.Context) {
	var req RequestItemByMonster
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Invalid monster ID"})
		return
	}

	validate := validator.GetValidator()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Validation failed: " + err.Error()})
		return
	}

	itemList, err := h.itemRepo.FindByMonsterID(c.Request.Context(), req.MonsterId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get item by monster ID"})
		return
	}

	if len(itemList) == 0 {
		c.JSON(http.StatusNotFound, MessageResponse{Message: "Item not found"})
		return
	}

	monsterName := ""
	monster, err := h.monsterRepo.FindById(c.Request.Context(), req.MonsterId)
	if err == nil && monster != nil {
		monsterName = monster.Name
	}

	c.JSON(http.StatusOK, toItemByMonsterResponse(req.MonsterId, monsterName, itemList))
}
