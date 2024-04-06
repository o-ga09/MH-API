package monster

import (
	"log/slog"

	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/internal/service/monsters"

	"net/http"

	"github.com/gin-gonic/gin"
)

type MonsterHandler struct {
	monsterService monsters.MonsterService
}

func NewMonsterHandler(s monsters.MonsterService) *MonsterHandler {
	return &MonsterHandler{
		monsterService: s,
	}
}

// GetAll godoc
// @Summary モンスター検索（複数件）
// @Description モンスターを検索して、条件に合致するモンスターを複数件取得する
// @Tags モンスター検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} Monsters
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /monsters [get]
func (m *MonsterHandler) GetAll(c *gin.Context) {
	id, ook := c.Params.Get("id")
	if ook {
		id = ""
	}
	res, err := m.monsterService.GetMonster(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	monsters := []ResponseJson{}
	for _, r := range res {
		monsters = append(monsters, ResponseJson{
			Id:               r.ID,
			Name:             r.Name,
			Desc:             r.Description,
			Location:         "dummy location",
			Category:         "dummy category",
			Title:            "dummy title",
			Weakness_attack:  "dummy weakness attack",
			Weakness_element: "dummy weakness element",
		})
	}
	response := Monsters{
		Total:    len(res),
		Monsters: monsters,
	}
	c.JSON(http.StatusOK, response)
}

// GetById godoc
// @Summary モンスター検索（1件）
// @Description モンスターを検索して、条件に合致するモンスターを1件取得する
// @Tags モンスター検索
// @Accept json
// @Produce json
// @Param request path string ture "モンスターID"
// @Success 200 {object} Monster
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /monsters/:monsterid [get]
func (m *MonsterHandler) GetById(c *gin.Context) {
	id, ook := c.Params.Get("id")
	if ook {
		id = ""
	}
	res, err := m.monsterService.GetMonster(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	response := Monster{
		Monster: ResponseJson{
			Id:               res[0].ID,
			Name:             res[0].Name,
			Desc:             res[0].Description,
			Location:         "dummy location",
			Category:         "dummy category",
			Title:            "dummy title",
			Weakness_attack:  "dummy weakness attack",
			Weakness_element: "dummy weakness element",
		},
	}
	c.JSON(http.StatusOK, response)
}

// GetRankingMonster godoc
// @Summary モンスター人気投票結果検索
// @Description 人気投票ランキングの結果を検索する
// @Tags モンスター検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} Monsters
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /monsters/ranking [get]
func (m *MonsterHandler) GetRankingMonster(c *gin.Context) {}
