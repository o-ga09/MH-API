package skill

import (
	"mh-api/internal/service/skills"
	"mh-api/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	service skills.ISkillService
}

func NewSkillHandler(s skills.ISkillService) *SkillHandler {
	return &SkillHandler{
		service: s,
	}
}

// GetSkills godoc
// @Summary スキル一覧を取得する
// @Description 全てのスキルとその情報の一覧を取得する
// @Tags スキル検索
// @Produce json
// @Success 200 {object} Skills
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /v1/skills [get]
func (h *SkillHandler) GetSkills(c *gin.Context) {
	skillsResponse, err := h.service.GetAllSkills(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get skills"})
		return
	}
	c.JSON(http.StatusOK, ToSkillListResponse(*skillsResponse))
}

// GetSkill godoc
// @Summary スキル検索（1件）
// @Description スキルを検索して、条件に合致するスキルを1件取得する
// @Tags スキル検索
// @Accept json
// @Produce json
// @Param skillId path string true "スキルID"
// @Success 200 {object} Skill
// @Failure      400  {object}  MessageResponse
// @Failure      404  {object}  MessageResponse
// @Failure      500  {object}  MessageResponse
// @Router /v1/skills/{skillId} [get]
func (h *SkillHandler) GetSkill(c *gin.Context) {
	var req RequestSkillByID
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Invalid skill ID"})
		return
	}

	// バリデーション実行
	validate := validator.GetValidator()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Validation failed: " + err.Error()})
		return
	}

	skill, err := h.service.GetSkillByID(c.Request.Context(), req.SkillId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get skill"})
		return
	}

	if skill == nil {
		c.JSON(http.StatusNotFound, MessageResponse{Message: "Skill not found"})
		return
	}

	c.JSON(http.StatusOK, ToSkillResponse(*skill))
}