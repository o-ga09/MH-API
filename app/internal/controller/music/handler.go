package music

import (
	"mh-api/app/internal/service/monsters"

	"github.com/gin-gonic/gin"
)

type BGMHandler struct {
	monsterService monsters.MonsterService
}

func NewBGMHandler(s monsters.MonsterService) *BGMHandler {
	return &BGMHandler{
		monsterService: s,
	}
}

// GetBGM godoc
// @Summary BGM検索（複数件）
// @Description モンスターのBGMを検索して、条件に合致するモンスターのBGMを複数件取得する
// @Tags BGM検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} BGMs
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /bgms [get]
func (h *BGMHandler) GetBGM(c *gin.Context) {}

// GetBGMById godoc
// @Summary BGM検索（1件）
// @Description モンスターのBGMを検索して、条件に合致するモンスターのBGMを1件取得する
// @Tags BGM検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} BGM
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /bgms/:bgmid [get]
func (h *BGMHandler) GetBGMById(c *gin.Context) {}

// GetBGM godoc
// @Summary BGM人気投票結果検索
// @Description 人気投票ランキングの結果を検索する
// @Tags BGM検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} BGM
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /bgms/ranking [get]
func (m *BGMHandler) GetRankingBGM(c *gin.Context) {}
