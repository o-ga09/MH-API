package monster

import (
	"context"
	"fmt"

	"log/slog"
	"mh-api/internal/service/monsters"
	"mh-api/pkg/constant"
	"mh-api/pkg/ptr"
	"mh-api/pkg/validator"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const YOUTUBE_URL = "https://www.youtube.com/watch?v="

type MonsterHandler struct {
	monsterService monsters.IMonsterService
}

func NewMonsterHandler(s monsters.IMonsterService) *MonsterHandler {
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
		slog.Log(c, constant.SeverityError, "param marshal error", "error message", err)
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	if param.Limit == 0 {
		param.Limit = 100
	}
	validate := validator.GetValidator()
	err = validate.Struct(&param)
	if err != nil {
		slog.Log(c, constant.SeverityError, "validation error", "error message", err)
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	id, ok := c.Params.Get("id")
	if ok {
		id = ""
	}
	ctx = context.WithValue(ctx, "param", param)
	result, err := m.monsterService.FetchMonsterDetail(ctx, id)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, constant.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, constant.SeverityError, "database error", "error", err)
		return
	}

	monsters := []ResponseJson{}
	for _, r := range result.Monsters {
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
			Description:        ptr.StrToPtr(r.Description),
			AnotherName:        ptr.StrToPtr(r.AnotherName),
			NameEn:             ptr.StrToPtr(r.NameEn),
			Location:           ptr.StrArrayToPtr(r.Location),
			Category:           r.Category,
			Title:              ptr.StrArrayToPtr(r.Title),
			FirstWeak_Attack:   ptr.StrToPtr(r.FirstWeak_Attack),
			FirstWeak_Element:  ptr.StrToPtr(r.FirstWeak_Element),
			SecondWeak_Attack:  ptr.StrToPtr(r.SecondWeak_Attack),
			SecondWeak_Element: ptr.StrToPtr(r.SecondWeak_Element),
			Weakness_attack:    wa,
			Weakness_element:   we,
			Ranking:            ranking,
			ImageUrl:           ptr.CreateImageURL(r.Id),
			BGM:                bgm,
			Element:            r.Element,
		})
	}
	response := Monsters{
		Total:    result.Total,
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

	span := sentry.StartSpan(ctx, "handler.GetById")
	span.SetTag("handler", "monsterHandler")
	defer span.Finish()

	id, ok := c.Params.Get("id")
	if !ok {
		slog.Log(c, constant.SeverityError, "path parameter required")
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	result, err := m.monsterService.FetchMonsterDetail(ctx, id)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, constant.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, constant.SeverityError, "database error", "error", err)
		return
	}

	monster := ResponseJson{}
	for _, r := range result.Monsters {
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
			AnotherName:        ptr.StrToPtr(r.Description),
			Location:           ptr.StrArrayToPtr(r.Location),
			Category:           r.Category,
			Title:              ptr.StrArrayToPtr(r.Title),
			FirstWeak_Attack:   ptr.StrToPtr(r.FirstWeak_Attack),
			FirstWeak_Element:  ptr.StrToPtr(r.FirstWeak_Element),
			SecondWeak_Attack:  ptr.StrToPtr(r.SecondWeak_Attack),
			SecondWeak_Element: ptr.StrToPtr(r.SecondWeak_Element),
			Weakness_attack:    wa,
			Weakness_element:   we,
			Ranking:            ranking,
			ImageUrl:           ptr.CreateImageURL(r.Id),
			BGM:                bgm,
			Element:            r.Element, // Added Element
		}
	}
	response := Monster{
		Monster: monster,
	}
	c.JSON(http.StatusOK, response)
}
