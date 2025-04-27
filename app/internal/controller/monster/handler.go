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
		slog.Log(c, middleware.SeverityError, "database error", "error", err)
		return
	}

	monsters := []ResponseJson{}
	for _, r := range res {
		var wa []*Weakness_attack
		var we []*Weakness_element
		var ranking []*Ranking
		for _, w := range r.Weakness_attack {
			wa = append(wa, &Weakness_attack{
				Slashing: w.Slashing,
				Blow:     w.Blow,
				Bullet:   w.Bullet,
			})
		}

		for _, w := range r.Weakness_element {
			we = append(we, &Weakness_element{
				Fire:    w.Fire,
				Water:   w.Water,
				Thunder: w.Thunder,
				Ice:     w.Ice,
				Dragon:  w.Dragon,
			})
		}
		for _, r := range r.Ranking {
			ranking = append(ranking, &Ranking{
				Ranking:  r.Ranking,
				VoteYear: r.VoteYear,
			})
		}

		monsters = append(monsters, ResponseJson{
			Id:                 r.Id,
			Name:               r.Name,
			AnotherName:        pkg.StrToPtr(r.Description),
			Location:           pkg.StrArrayToPtr(r.Location),
			Category:           r.Category,
			Title:              pkg.StrArrayToPtr(r.Title),
			FirstWeak_Attack:   pkg.StrToPtr(r.FirstWeak_Attack),
			FirstWeak_Element:  pkg.StrToPtr(r.FirstWeak_Element),
			SecondWeak_Attack:  pkg.StrToPtr(r.SecondWeak_Attack),
			SecondWeak_Element: pkg.StrToPtr(r.SecondWeak_Element),
			Weakness_attack:    wa,
			Weakness_element:   we,
			Ranking:            ranking,
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
		slog.Log(c, middleware.SeverityError, "database error", "error", err)
		return
	}

	monster := ResponseJson{}
	for _, r := range res {
		var wa []*Weakness_attack
		var we []*Weakness_element
		var ranking []*Ranking
		for _, w := range r.Weakness_attack {
			wa = append(wa, &Weakness_attack{
				Slashing: w.Slashing,
				Blow:     w.Blow,
				Bullet:   w.Bullet,
			})
		}

		for _, w := range r.Weakness_element {
			we = append(we, &Weakness_element{
				Fire:    w.Fire,
				Water:   w.Water,
				Thunder: w.Thunder,
				Ice:     w.Ice,
				Dragon:  w.Dragon,
			})
		}

		for _, rank := range r.Ranking {
			ranking = append(ranking, &Ranking{
				Ranking:  rank.Ranking,
				VoteYear: rank.VoteYear,
			})
		}

		monster = ResponseJson{
			Id:                 r.Id,
			Name:               r.Name,
			AnotherName:        pkg.StrToPtr(r.Description),
			Location:           pkg.StrArrayToPtr(r.Location),
			Category:           r.Category,
			Title:              pkg.StrArrayToPtr(r.Title),
			FirstWeak_Attack:   pkg.StrToPtr(r.FirstWeak_Attack),
			FirstWeak_Element:  pkg.StrToPtr(r.FirstWeak_Element),
			SecondWeak_Attack:  pkg.StrToPtr(r.SecondWeak_Attack),
			SecondWeak_Element: pkg.StrToPtr(r.SecondWeak_Element),
			Weakness_attack:    wa,
			Weakness_element:   we,
			Ranking:            ranking,
			ImageUrl:           pkg.CreateImageURL(r.Id),
		}
	}
	response := Monster{
		Monster: monster,
	}
	c.JSON(http.StatusOK, response)
}
