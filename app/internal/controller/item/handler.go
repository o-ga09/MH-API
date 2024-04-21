package item

import (
	"context"
	"log/slog"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/internal/service/items"
	"mh-api/app/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemHandler struct {
	fetchItemService          *items.FetchItemList
	fetchItemByMonsterService *items.FetchItemByMonster
	fetchItemByIdService      *items.FetchItemById
	saveItemService           *items.SaveItem
	removeItemService         *items.RemoveItem
}

func NewItemHandler(f items.FetchItemList, m items.FetchItemByMonster, b items.FetchItemById, s items.SaveItem, r items.RemoveItem) *ItemHandler {
	return &ItemHandler{
		fetchItemService:          &f,
		fetchItemByMonsterService: &m,
		fetchItemByIdService:      &b,
		saveItemService:           &s,
		removeItemService:         &r,
	}
}

// GetItems godoc
// @Summary アイテム検索（複数件）
// @Description アイテムを検索して、条件に合致するアイテムを複数件取得する
// @Tags アイテム検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} Items
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /items [get]
func (h *ItemHandler) GetItems(c *gin.Context) {
	var param RequestParam
	err := c.ShouldBindQuery(&param)
	if err != nil {
		slog.Log(c, middleware.SeverityError, "param marshal error", "error message", err)
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	if param.Limit == 0 {
		param.Limit = 100
	}
	validate := pkg.GetValidator()
	err = validate.Struct(&param)
	if err != nil {
		slog.Log(c, middleware.SeverityError, "validation error", "error message", err)
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	ctx := context.WithValue(c.Request.Context(), "param", param)
	res, err := h.fetchItemService.Run(ctx)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, middleware.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	items := []ResponseJson{}
	for _, r := range res {
		items = append(items, ResponseJson{
			Id:       r.ID,
			ItemName: r.Name,
		})
	}
	response := Items{
		Total: len(res),
		Item:  items,
	}
	c.JSON(http.StatusOK, response)
}

// GetItems godoc
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
	var param RequestParam
	id, ook := c.Params.Get("id")
	if !ook {
		slog.Log(c, middleware.SeverityError, "path parameter required")
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	param.MonsterIds = id
	param.Limit = 1
	ctx := context.WithValue(c.Request.Context(), "param", param)
	res, err := h.fetchItemByIdService.Run(ctx, id)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, middleware.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	item := ResponseJson{
		Id:       res.ID,
		ItemName: res.Name,
	}

	response := Item{
		Item: item,
	}
	c.JSON(http.StatusOK, response)
}

// GetItems godoc
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
// @Router /items/:itemId/monsters [get]
func (h *ItemHandler) GetItemByMonster(c *gin.Context) {
	var param RequestParam

	id, ook := c.Params.Get("id")
	if !ook {
		slog.Log(c, middleware.SeverityError, "path parameter required")
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	err := c.ShouldBindQuery(&param)
	if err != nil {
		slog.Log(c, middleware.SeverityError, "param marshal error", "error message", err)
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	if param.Limit == 0 {
		param.Limit = 100
	}
	validate := pkg.GetValidator()
	err = validate.Struct(&param)
	if err != nil {
		slog.Log(c, middleware.SeverityError, "validation error", "error message", err)
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	ctx := context.WithValue(c.Request.Context(), "param", param)
	res, err := h.fetchItemByMonsterService.Run(ctx, id)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, middleware.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	monster := []Monster{}
	for _, r := range res.Monster {
		monster = append(monster, Monster{
			MonsterId:   r.ID,
			MonsterName: r.Name,
		})
	}

	response := ItemsByMonster{
		ItemId:   res.ItemId,
		ItemName: res.ItemName,
		Monsters: monster,
	}
	c.JSON(http.StatusOK, response)
}
