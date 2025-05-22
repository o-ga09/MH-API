package item

import (
	"mh-api/internal/service/monsters"

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

// NOT IMPLEMENT GetItems godoc
// NOT IMPLEMENT @Summary アイテム検索（複数件）
// NOT IMPLEMENT @Description アイテムを検索して、条件に合致するアイテムを複数件取得する
// NOT IMPLEMENT @Tags アイテム検索
// NOT IMPLEMENT @Accept json
// NOT IMPLEMENT @Produce json
// NOT IMPLEMENT @Param request query RequestParam true  "クエリパラメータ"
// NOT IMPLEMENT @Success 200 {object} Items
// NOT IMPLEMENT @Failure      400  {object}  MessageResponse
// NOT IMPLEMENT @Failure      404  {object}  MessageResponse
// NOT IMPLEMENT @Failure      500  {object}  MessageResponse
// NOT IMPLEMENT @Router /items [get]
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

// NOT IMPLEMENT GetItems godoc
// NOT IMPLEMENT @Summary アイテム検索（モンスター別）
// NOT IMPLEMENT @Description 指定のアイテムが取得可能なモンスターの一覧
// NOT IMPLEMENT @Tags アイテム検索
// NOT IMPLEMENT @Accept json
// NOT IMPLEMENT @Produce json
// NOT IMPLEMENT @Param request query RequestParam true  "クエリパラメータ"
// NOT IMPLEMENT @Success 200 {object} ItemsByMonster
// NOT IMPLEMENT @Failure      400  {object}  MessageResponse
// NOT IMPLEMENT @Failure      404  {object}  MessageResponse
// NOT IMPLEMENT @Failure      500  {object}  MessageResponse
// NOT IMPLEMENT @Router /items/monsters [get]
func (h *ItemHandler) GetItemByMonster(c *gin.Context) {}
