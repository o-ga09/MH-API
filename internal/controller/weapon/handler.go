package weapon

import (
	"mh-api/internal/service/monsters"

	"github.com/gin-gonic/gin"
)

type WeaponHandler struct {
	monsterService monsters.MonsterService
}

func NewWeaponHandler(s monsters.MonsterService) *WeaponHandler {
	return &WeaponHandler{
		monsterService: s,
	}
}

// GetWeapon godoc
// @Summary 武器検索（複数件）
// @Description 武器を検索して、条件に合致する武器を複数件取得する
// @Tags 武器検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} Weapon
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /weapons [get]
func (h *WeaponHandler) GetBGM(c *gin.Context) {}

// GetWeaponById godoc
// @Summary 武器検索（1件）
// @Description 武器を検索して、条件に合致する武器を1件取得する
// @Tags 武器検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} Weapon
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /weapons/:bgmid [get]
func (h *WeaponHandler) GetBGMById(c *gin.Context) {}
