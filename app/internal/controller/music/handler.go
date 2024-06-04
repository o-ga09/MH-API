package music

import (
	"context"
	"log/slog"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/internal/service/music"
	"mh-api/app/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type contextKey string

const ParamKey contextKey = "param"

type BGMHandler struct {
	musicService music.MusicService
}

func NewBGMHandler(s music.MusicService) *BGMHandler {
	return &BGMHandler{
		musicService: s,
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
func (h *BGMHandler) GetBGM(c *gin.Context) {
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
	ctx := context.WithValue(c.Request.Context(), ParamKey, param)
	res, err := h.musicService.FetchList(ctx, id)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, middleware.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	bgms := []ResponseJson{}
	for _, r := range res {
		bgms = append(bgms, ResponseJson{
			Id:   r.Id,
			Name: r.Name,
			Url:  r.Url,
		})
	}
	response := BGMs{
		Total:  len(bgms),
		Limit:  param.Limit,
		Offset: param.Offset,
		BGM:    bgms,
	}
	c.JSON(http.StatusOK, response)
}

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
func (h *BGMHandler) GetBGMById(c *gin.Context) {
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
	if !ok {
		id = ""
	}

	ctx := context.WithValue(c.Request.Context(), ParamKey, param)
	res, err := h.musicService.FetchList(ctx, id)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, middleware.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	bgms := ResponseJson{}
	for _, r := range res {
		bgms = ResponseJson{
			Id:   r.Id,
			Name: r.Name,
			Url:  r.Url,
		}
	}
	response := BGM{
		BGM: bgms,
	}
	c.JSON(http.StatusOK, response)
}

// GetBGM godoc
// @Summary BGM人気投票結果検索
// @Description 人気投票ランキングの結果を検索する
// @Tags BGM検索
// @Accept json
// @Produce json
// @Param request query RequestParam true  "クエリパラメータ"
// @Success 200 {object} BGMRankings
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /bgms/ranking [get]
func (h *BGMHandler) GetRankingBGM(c *gin.Context) {
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

	ctx := context.WithValue(c.Request.Context(), ParamKey, param)
	res, err := h.musicService.FetchRank(ctx)

	if err == gorm.ErrRecordNotFound {
		slog.Log(c, middleware.SeverityError, "Record Not Found", "error message", err)
		c.JSON(http.StatusNotFound, MessageResponse{Message: "NOT FOUND"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	bgms := []ResponseRankingJson{}
	for _, r := range res {
		var rankings []Ranking
		for _, rank := range r.Ranking {
			ranking := Ranking{
				Ranking:  rank.Rank,
				VoteYear: rank.VoteYear,
			}
			rankings = append(rankings, ranking)
		}
		bgms = append(bgms, ResponseRankingJson{
			BgmId:   r.Id,
			Name:    r.Name,
			Url:     r.Url,
			Ranking: rankings,
		})
	}
	response := BGMRankings{
		Total:   len(bgms),
		Limit:   param.Limit,
		Offset:  param.Offset,
		Ranking: bgms,
	}
	c.JSON(http.StatusOK, response)
}
