package item

import (
	"context"
	"log/slog"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/internal/service/item"
	"mh-api/app/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type contextKey string

const ParamKey contextKey = "param"

type ItemHandler struct {
	itemService item.ItemService
}

func NewItemHandler(s item.ItemService) *ItemHandler {
	return &ItemHandler{
		itemService: s,
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

	ctx := context.WithValue(c.Request.Context(), ParamKey, param)
	res, err := h.itemService.GetItems(ctx)

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
			Id:       r.Id,
			ItemName: r.Name,
		})
	}
	response := Items{
		Total:  len(items),
		Limit:  param.Limit,
		Offset: param.Offset,
		Item:   items,
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

	id := c.Param("id")

	ctx := context.WithValue(c.Request.Context(), ParamKey, param)
	res, err := h.itemService.GetItemById(ctx, id)

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

	response := ResponseJson{
		Id:       res[0].Id,
		ItemName: res[0].Name,
	}

	c.JSON(http.StatusOK, response)
}

// GetItems godoc
// @Summary 取得可能なアイテムの一覧(アイテム別)
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

	id := c.Param("id")

	ctx := context.WithValue(c.Request.Context(), ParamKey, param)
	res, err := h.itemService.GetItemsByMonsterList(ctx, id)

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
	itemsWithMonster := []MonsterList{}
	for _, r := range res {
		for _, m := range r.Monster {
			monster = append(monster, Monster{
				MonsterId: m.MonsterId,
			})
		}
		itemsWithMonster = append(itemsWithMonster, MonsterList{
			ItemId:   r.Id,
			ItemName: r.Name,
			Monsters: monster,
		})
		monster = nil
	}
	response := ItemsByMonsterList{
		Total:       len(itemsWithMonster),
		Limit:       param.Limit,
		Offset:      param.Offset,
		MonsterList: itemsWithMonster,
	}
	c.JSON(http.StatusOK, response)
}

// GetItems godoc
// @Summary 取得可能なアイテムの一覧(モンスター別)
// @Description 指定のモンスターから取得可能なアイテムの一覧
// @Tags アイテム検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} ItemsByMonster
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /items/monsters/:id [get]
func (h *ItemHandler) GetItemByMonsterId(c *gin.Context) {
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

	id := c.Param("id")

	ctx := context.WithValue(c.Request.Context(), ParamKey, param)
	res, err := h.itemService.GetItemByMonsterId(ctx, id)

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
	for _, r := range res[0].Item {
		items = append(items, ResponseJson{
			Id:       r.Id,
			ItemName: r.Name,
		})
	}
	response := ItemsByMonster{
		MonsterId:   res[0].MonsterId,
		MonsterName: res[0].MonsterName,
		Items:       items,
	}
	c.JSON(http.StatusOK, response)
}
