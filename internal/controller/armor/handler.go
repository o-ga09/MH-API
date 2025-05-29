package armor

import (
	"context"
	"errors"
	"mh-api/internal/service/armors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IArmorService interface {
	GetAllArmors(ctx context.Context) (*armors.ListArmorsResponse, error)
	GetArmorByID(ctx context.Context, armorId string) (*armors.ArmorData, error)
}

type ArmorHandler struct {
	service IArmorService
}

func NewArmorHandler(s IArmorService) *ArmorHandler {
	return &ArmorHandler{service: s}
}

// GetAllArmors godoc
// @Summary 防具一覧を取得します
// @Description 全ての防具のリストを返します。
// @Tags Armors
// @Accept json
// @Produce json
// @Success 200 {object} ListArmorsResponse "防具のリスト"
// @Failure 500 {object} ErrorResponse "サーバ内部エラー"
// @Router /skills [get]
func (h *ArmorHandler) GetAllArmors(c *gin.Context) {
	appCtx := c.Request.Context()

	serviceRes, err := h.service.GetAllArmors(appCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to get armors: " + err.Error()})
		return
	}

	ctrlResponseArmors := make([]ArmorDetailResponse, len(serviceRes.Armors))
	for i, sa := range serviceRes.Armors {
		skills := make([]SkillResponse, len(sa.Skill))
		for j, skill := range sa.Skill {
			skills[j] = SkillResponse{
				ID:   skill.ID,
				Name: skill.Name,
			}
		}

		requiredItems := make([]RequiredItemResponse, len(sa.Required))
		for k, item := range sa.Required {
			requiredItems[k] = RequiredItemResponse{
				ID:   item.ID,
				Name: item.Name,
			}
		}

		ctrlResponseArmors[i] = ArmorDetailResponse{
			ID:      sa.ID,
			Name:    sa.Name,
			Skills:  skills,
			Slot:    sa.Slot,
			Defense: sa.Defense,
			Resistance: ResistanceResponse{
				Fire:      sa.Resistance.Fire,
				Water:     sa.Resistance.Water,
				Lightning: sa.Resistance.Lightning,
				Ice:       sa.Resistance.Ice,
				Dragon:    sa.Resistance.Dragon,
			},
			Required: requiredItems,
		}
	}

	ctrlResponse := ListArmorsResponse{
		Armors: ctrlResponseArmors,
	}

	c.JSON(http.StatusOK, ctrlResponse)
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
// @Router /skills/{id} [get]
func (h *ArmorHandler) GetArmorByID(c *gin.Context) {
	appCtx := c.Request.Context()

	var req GetArmorByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request parameters"})
		return
	}
	if req.ArmorID == " " {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Armor ID cannot be empty"})
		return
	}

	serviceRes, err := h.service.GetArmorByID(appCtx, req.ArmorID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: "Armor not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to get armor: " + err.Error()})
		return
	}

	skills := make([]SkillResponse, len(serviceRes.Skill))
	for j, skill := range serviceRes.Skill {
		skills[j] = SkillResponse{
			ID:   skill.ID,
			Name: skill.Name,
		}
	}

	requiredItems := make([]RequiredItemResponse, len(serviceRes.Required))
	for k, item := range serviceRes.Required {
		requiredItems[k] = RequiredItemResponse{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	ctrlResponse := ArmorDetailResponse{
		ID:      serviceRes.ID,
		Name:    serviceRes.Name,
		Skills:  skills,
		Slot:    serviceRes.Slot,
		Defense: serviceRes.Defense,
		Resistance: ResistanceResponse{
			Fire:      serviceRes.Resistance.Fire,
			Water:     serviceRes.Resistance.Water,
			Lightning: serviceRes.Resistance.Lightning,
			Ice:       serviceRes.Resistance.Ice,
			Dragon:    serviceRes.Resistance.Dragon,
		},
		Required: requiredItems,
	}

	c.JSON(http.StatusOK, ctrlResponse)
}
