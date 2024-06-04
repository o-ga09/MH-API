package item

import (
	"mh-api/app/internal/service/monsters"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	monsterService monsters.MonsterService
}

func NewItemHandler(s monsters.MonsterService) *ItemHandler {
	return &ItemHandler{
		monsterService: s,
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
func (h *ItemHandler) GetItems(c *gin.Context) {}

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
func (h *ItemHandler) GetItem(c *gin.Context) {}

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
func (h *ItemHandler) GetItemByMonster(c *gin.Context) {}

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
func (h *ItemHandler) GetItemByMonsterId(c *gin.Context) {}
