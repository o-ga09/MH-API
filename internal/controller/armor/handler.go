package armor

import (
	"errors"
	"mh-api/internal/domain/armors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArmorHandler struct {
	repo armors.Repository
}

func NewArmorHandler(repo armors.Repository) *ArmorHandler {
	return &ArmorHandler{repo: repo}
}

// GetAllArmors godoc
// @Summary 防具一覧を取得します
// @Description 防具を検索して一覧を返します（名前・スキル名・スロットによる絞り込み・ページネーション対応）
// @Tags Armors
// @Accept json
// @Produce json
// @Param name query string false "防具名（部分一致）"
// @Param skill_name query string false "スキル名（部分一致）"
// @Param slot query string false "スロット（完全一致）"
// @Param limit query int false "取得件数" default(100)
// @Param offset query int false "取得開始位置" default(0)
// @Param sort query string false "ソート順 (asc/desc)"
// @Success 200 {object} ListArmorsResponse "防具のリスト"
// @Failure 400 {object} ErrorResponse "リクエストパラメータが不正な場合"
// @Failure 500 {object} ErrorResponse "サーバ内部エラー"
// @Router /armors [get]
func (h *ArmorHandler) GetAllArmors(c *gin.Context) {
	var req SearchArmorsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request parameters: " + err.Error()})
		return
	}

	params := armors.SearchParams{
		Name:      req.Name,
		SkillName: req.SkillName,
		Slot:      req.Slot,
		Limit:     req.Limit,
		Offset:    req.Offset,
		Sort:      req.Sort,
	}

	result, err := h.repo.Find(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to get armors: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, ListArmorsResponse{
		Total:  result.Total,
		Armors: toArmorDetailResponses(result.Armors),
	})
}

// GetArmorByID godoc
// @Summary 防具詳細を取得します
// @Description 指定されたIDの防具の詳細を返します。
// @Tags Armors
// @Accept json
// @Produce json
// @Param id path string true "防具ID"
// @Success 200 {object} ArmorDetailResponse "防具の詳細"
// @Failure 400 {object} ErrorResponse "リクエストパラメータが不正な場合"
// @Failure 404 {object} ErrorResponse "防具が見つからない場合"
// @Failure 500 {object} ErrorResponse "サーバ内部エラー"
// @Router /armors/{id} [get]
func (h *ArmorHandler) GetArmorByID(c *gin.Context) {
	var req GetArmorByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request parameters"})
		return
	}
	if req.ArmorID == " " {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Armor ID cannot be empty"})
		return
	}

	armor, err := h.repo.GetByID(c.Request.Context(), req.ArmorID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: "Armor not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to get armor: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, toArmorDetailResponse(armor))
}

func toArmorDetailResponse(a *armors.Armor) ArmorDetailResponse {
	skillResponses := make([]SkillResponse, len(a.Skills))
	for i, s := range a.Skills {
		skillResponses[i] = SkillResponse{
			ID:   s.SkillId,
			Name: s.SkillName,
		}
	}

	requiredItems := make([]RequiredItemResponse, len(a.RequiredItems))
	for i, item := range a.RequiredItems {
		requiredItems[i] = RequiredItemResponse{
			ID:   item.ItemId,
			Name: item.ItemName,
		}
	}

	return ArmorDetailResponse{
		ID:      a.ArmorId,
		Name:    a.Name,
		Skills:  skillResponses,
		Slot:    a.Slot,
		Defense: a.Defense,
		Resistance: ResistanceResponse{
			Fire:      a.FireResistance,
			Water:     a.WaterResistance,
			Lightning: a.LightningResistance,
			Ice:       a.IceResistance,
			Dragon:    a.DragonResistance,
		},
		Required: requiredItems,
	}
}

func toArmorDetailResponses(armorList armors.Armors) []ArmorDetailResponse {
	res := make([]ArmorDetailResponse, len(armorList))
	for i, a := range armorList {
		res[i] = toArmorDetailResponse(a)
	}
	return res
}
