package monster

import (
	"fmt"
	"log/slog"
	"net/http"

	"mh-api/internal/domain/monsters"
	"mh-api/pkg/constant"
	"mh-api/pkg/ptr"
	"mh-api/pkg/validator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const YOUTUBE_URL = "https://www.youtube.com/watch?v="

type MonsterHandler struct {
	repo monsters.Repository
}

func NewMonsterHandler(repo monsters.Repository) *MonsterHandler {
	return &MonsterHandler{repo: repo}
}

func (m *MonsterHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

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

	searchParams := monsters.SearchParams{
		MonsterIds:      param.MonsterIds,
		MonsterName:     param.MonsterName,
		UsageElement:    param.UsageElement,
		WeaknessElement: param.WeaknessElement,
		Limit:           param.Limit,
		Offset:          param.Offset,
		Sort:            param.Sort,
	}
	result, err := m.repo.FindAll(ctx, searchParams)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, constant.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "can not get records"})
		slog.Log(c, constant.SeverityError, "database error", "error", err)
		return
	}

	list := make([]ResponseJson, 0, len(result.Monsters))
	for _, r := range result.Monsters {
		list = append(list, toResponseJson(r))
	}
	c.JSON(http.StatusOK, Monsters{
		Total:    result.Total,
		Monsters: list,
	})
}

// GetById godoc
// @Summary モンスター検索（1件）
// @Description モンスターを検索して、条件に合致するモンスターを1件取得する
// @Tags モンスター検索
// @Accept json
// @Produce json
// @Param request path string true "モンスターID"
// @Success 200 {object} Monster
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /monsters/{id} [get]
func (m *MonsterHandler) GetById(c *gin.Context) {
	ctx := c.Request.Context()

	id, ok := c.Params.Get("id")
	if !ok {
		slog.Log(c, constant.SeverityError, "path parameter required")
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "BAD REQUEST"})
		return
	}

	r, err := m.repo.FindById(ctx, id)
	if err == gorm.ErrRecordNotFound {
		slog.Log(c, constant.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "can not get records"})
		slog.Log(c, constant.SeverityError, "database error", "error", err)
		return
	}

	c.JSON(http.StatusOK, Monster{Monster: toResponseJson(r)})
}

func toResponseJson(r *monsters.Monster) ResponseJson {
	var wa []*Weakness_attack
	var we []*Weakness_element
	var ranking []*Ranking
	var bgm []*Music
	var locations []string
	var titles []string

	for _, w := range r.Weakness {
		wa = append(wa, &Weakness_attack{
			Slashing: w.Slashing,
			Blow:     w.Blow,
			Bullet:   w.Bullet,
		})
		we = append(we, &Weakness_element{
			Fire:    w.Fire,
			Water:   w.Water,
			Thunder: w.Lightning,
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
			Name: bg.Name,
			Url:  fmt.Sprintf("%s%s", YOUTUBE_URL, bg.Url),
		})
	}
	for _, f := range r.Field {
		locations = append(locations, f.Name)
	}
	for _, p := range r.Product {
		titles = append(titles, p.Name)
	}

	category := ""
	if r.Tribe != nil {
		category = r.Tribe.Name_ja
	}

	var firstWeakAttack, secondWeakAttack, firstWeakElement, secondWeakElement *string
	if len(r.Weakness) > 0 {
		firstWeakAttack = ptr.StrToPtr(r.Weakness[0].FirstWeakAttack)
		secondWeakAttack = ptr.StrToPtr(r.Weakness[0].SecondWeakAttack)
		firstWeakElement = ptr.StrToPtr(r.Weakness[0].FirstWeakElement)
		secondWeakElement = ptr.StrToPtr(r.Weakness[0].SecondWeakElement)
	}

	return ResponseJson{
		Id:                 r.MonsterId,
		Name:               r.Name,
		Description:        ptr.StrToPtr(r.Description),
		AnotherName:        ptr.StrToPtr(r.AnotherName),
		NameEn:             ptr.StrToPtr(r.NameEn),
		Location:           ptr.StrArrayToPtr(locations),
		Category:           category,
		Title:              ptr.StrArrayToPtr(titles),
		FirstWeak_Attack:   firstWeakAttack,
		FirstWeak_Element:  firstWeakElement,
		SecondWeak_Attack:  secondWeakAttack,
		SecondWeak_Element: secondWeakElement,
		Weakness_attack:    wa,
		Weakness_element:   we,
		Ranking:            ranking,
		ImageUrl:           ptr.CreateImageURL(r.MonsterId),
		BGM:                bgm,
		Element:            r.Element,
	}
}
