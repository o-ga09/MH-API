package item

import (
	"mh-api/internal/service/items"
	"mh-api/pkg/validator"
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
// @Success 200 {object} Items
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /v1/items [get]
func (h *ItemHandler) GetItems(c *gin.Context) {
	// このエンドポイントではクエリパラメータは必要ないが、将来的にページネーションなどを追加する可能性があるため
	// バリデーションのフレームワークのみ用意しておく
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
// @Param itemId path string true "アイテムID"
// @Success 200 {object} Item
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /v1/items/{itemId} [get]
func (h *ItemHandler) GetItem(c *gin.Context) {
	var req RequestItemByID
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Invalid item ID"})
		return
	}

	// バリデーション実行
	validate := validator.GetValidator()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Validation failed: " + err.Error()})
		return
	}

	item, err := h.service.GetItemByID(c.Request.Context(), req.ItemId)
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
// @Param monsterId path string true "モンスターID"
// @Success 200 {object} ItemsByMonster
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /v1/items/monsters/{monsterId} [get]
func (h *ItemHandler) GetItemByMonster(c *gin.Context) {
	// リクエスト構造体にURIパラメータをバインド
	var req RequestItemByMonster
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Invalid monster ID"})
		return
	}

	// バリデーション実行
	validate := validator.GetValidator()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Validation failed: " + err.Error()})
		return
	}

	item, err := h.service.GetItemByMonsterID(c.Request.Context(), req.MonsterId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get item by monster ID"})
		return
	}

	if item == nil {
		c.JSON(http.StatusNotFound, MessageResponse{Message: "Item not found"})
		return
	}

	c.JSON(http.StatusOK, ToItemByMonsterResponse(*item))
}
