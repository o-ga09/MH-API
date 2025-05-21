package monster

import (
	"fmt"
	"log/slog"

	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/internal/service/monsters"
	"mh-api/app/pkg"

	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

const YOUTUBE_URL = "https://www.youtube.com/watch?v="

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
	ctx := c.Request.Context()

	span := sentry.StartSpan(ctx, "handler.GetAll")
	span.Description = "GetAll"
	defer span.Finish()

	var param RequestParam
	err := c.ShouldBindQuery(&param)
	if err != nil {
		slog.Log(c, pkg.SeverityError, "param marshal error", "error message", err)
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	if param.Limit == 0 {
		param.Limit = 100
	}
	validate := pkg.GetValidator()
	err = validate.Struct(&param)
	if err != nil {
		slog.Log(c, pkg.SeverityError, "validation error", "error message", err)
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	id, ok := c.Params.Get("id")
	if ok {
		id = ""
	}
	ctx = context.WithValue(ctx, "param", param)
	res, err := m.monsterService.FetchMonsterDetail(ctx, id)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, pkg.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, pkg.SeverityError, "database error", "error", err)
		return
	}

	monsters := []ResponseJson{}
	for _, r := range res {
		var wa []*Weakness_attack
		var we []*Weakness_element
		var ranking []*Ranking
		var bgm []*Music
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
		for _, bg := range r.BGM {
			bgm = append(bgm, &Music{
				Name: bg.GetName(),
				Url:  fmt.Sprintf("%s%s", YOUTUBE_URL, bg.GetURL()),
			})
		}

		monsters = append(monsters, ResponseJson{
			Id:                 r.Id,
			Name:               r.Name,
			Description:        pkg.StrToPtr(r.Description),
			AnotherName:        pkg.StrToPtr(r.AnotherName),
			NameEn:             pkg.StrToPtr(r.NameEn),
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
			BGM:                bgm,
			Element:            r.Element,
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
	// トレーシング用のスパンを作成（親トランザクションの子スパンとして）
	ctx := c.Request.Context()
	db := middleware.GetDB(ctx)
	if db == nil {
		slog.Log(c, pkg.SeverityError, "database connection not found in context")
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "INTERNAL SERVER ERROR"})
		return
	}

	span := sentry.StartSpan(ctx, "handler.GetById")
	span.SetTag("handler", "monsterHandler")
	defer span.Finish()

	id, ook := c.Params.Get("id")
	if !ook {
		slog.Log(c, pkg.SeverityError, "path parameter required")
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	res, err := m.monsterService.FetchMonsterDetail(ctx, id)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, pkg.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, pkg.SeverityError, "database error", "error", err)
		return
	}

	monster := ResponseJson{}
	for _, r := range res {
		var wa []*Weakness_attack
		var we []*Weakness_element
		var ranking []*Ranking
		var bgm []*Music
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

		for _, bg := range r.BGM {
			bgm = append(bgm, &Music{
				Name: bg.GetName(),
				Url:  fmt.Sprintf("%s%s", YOUTUBE_URL, bg.GetURL()),
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
			BGM:                bgm,
			Element:            r.Element, // Added Element
		}
	}
	response := Monster{
		Monster: monster,
	}
	c.JSON(http.StatusOK, response)
}
