package monster

import (
	"log/slog"

	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/internal/service/monsters"
	"mh-api/app/pkg"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"gorm.io/gorm"
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

	id, ok := c.Params.Get("id")
	if ok {
		id = ""
	}
	ctx := context.WithValue(c.Request.Context(), "param", param)
	res, err := m.monsterService.FetchMonsterDetail(ctx, id)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, middleware.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "internal server error", "error message", err)
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
	if !ook {
		slog.Log(c, middleware.SeverityError, "path parameter required")
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	res, err := m.monsterService.FetchMonsterDetail(c.Request.Context(), id)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, middleware.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "internal server error", "error message", err)
		return
	}

	monster := ResponseJson{}
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
		monster = ResponseJson{
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
		}
	}
	response := Monster{
		Monster: monster,
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
func (m *MonsterHandler) GetRankingMonster(c *gin.Context) {
	var param RequestRankingParam
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
	res, err := m.monsterService.FetchMonsterRanking(ctx)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, middleware.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "internal server error", "error message", err)
		return
	}

	monsters := []ResponseRankingJson{}
	for _, r := range res {
		var rankings []Ranking
		for _, rank := range r.Ranking {
			ranking := Ranking{
				Ranking:  rank.Ranking,
				VoteYear: rank.VoteYear,
			}
			rankings = append(rankings, ranking)
		}
		monsters = append(monsters, ResponseRankingJson{
			Id:       r.Id,
			Name:     r.Name,
			Desc:     r.Description,
			Location: Location{Name: r.Location},
			Category: r.Category,
			Title:    Title{Name: r.Title},
			Ranking:  rankings,
		})
	}
	response := MonsterRanking{
		Ranking: monsters,
	}
	c.JSON(http.StatusOK, response)
}
