package monster

import (
	"log/slog"

	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/internal/service/monsters"
	"mh-api/app/pkg"

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
	var param RequestParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		slog.Log(c, middleware.SeverityError, "param marshal error", err)
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	validate := pkg.GetValidator()
	err = validate.Struct(&param)
	if err != nil {
		slog.Log(c, middleware.SeverityError, "validation error", err)
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	id, ook := c.Params.Get("id")
	if ook {
		id = ""
	}
	res, err := m.monsterService.FetchMonsterDetail(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	monsters := []ResponseJson{}
	for _, r := range res {
		var wa []Weakness_attack
		var we []Weakness_element
		for _, w := range r.Weakness_attack {
			wa = append(wa, Weakness_attack{
				Slashing: w.Slashing,
				Blow:     w.Blow,
				Bullet:   w.Bullet,
			})
		}

		for _, w := range r.Weakness_element {
			we = append(we, Weakness_element{
				Fire:    w.Fire,
				Water:   w.Water,
				Thunder: w.Thunder,
				Ice:     w.Ice,
				Dragon:  w.Dragon,
			})
		}
		monsters = append(monsters, ResponseJson{
			Id:                 r.Id,
			Name:               r.Name,
			Desc:               r.Description,
			Location:           Location{Name: r.Location},
			Category:           r.Category,
			Title:              Title{Name: r.Title},
			FirstWeak_Attack:   r.FirstWeak_Attack,
			FirstWeak_Element:  r.FirstWeak_Element,
			SecondWeak_Attack:  r.SecondWeak_Attack,
			SecondWeak_Element: r.SecondWeak_Element,
			Weakness_attack:    wa,
			Weakness_element:   we,
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
			Location:         Location{},
			Category:         "dummy category",
			Title:            Title{},
			Weakness_attack:  []Weakness_attack{},
			Weakness_element: []Weakness_element{},
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
