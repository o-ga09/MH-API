package skill

import (
	"errors"
	"mh-api/internal/domain/skills"
	"mh-api/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SkillHandler struct {
	repo skills.Repository
}

func NewSkillHandler(repo skills.Repository) *SkillHandler {
	return &SkillHandler{repo: repo}
}

// GetSkills godoc
// @Summary スキル一覧を取得する
// @Description スキルを検索して一覧を取得する（名前・説明文による絞り込み・ページネーション対応）
// @Tags スキル検索
// @Produce json
// @Param name query string false "スキル名（部分一致）"
// @Param description query string false "説明文（部分一致）"
// @Param limit query int false "取得件数" default(100)
// @Param offset query int false "取得開始位置" default(0)
// @Success 200 {object} Skills
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /skills [get]
func (h *SkillHandler) GetSkills(c *gin.Context) {
	var req SearchSkillsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Invalid request parameters: " + err.Error()})
		return
	}

	params := skills.SearchParams{
		Name:        req.Name,
		Description: req.Description,
		Limit:       req.Limit,
		Offset:      req.Offset,
	}

	result, err := h.repo.Find(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get skills"})
		return
	}
	c.JSON(http.StatusOK, ToSkillSearchResponse(result))
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
// @Router /skills/{skillId} [get]
func (h *SkillHandler) GetSkill(c *gin.Context) {
	var req RequestSkillByID
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Invalid skill ID"})
		return
	}

	validate := validator.GetValidator()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Validation failed: " + err.Error()})
		return
	}
	if req.SkillId == " " {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "Skill ID is required"})
		return
	}

	skill, err := h.repo.FindById(c.Request.Context(), req.SkillId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, MessageResponse{Message: "Skill not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Failed to get skill"})
		return
	}

	if skill == nil {
		c.JSON(http.StatusNotFound, MessageResponse{Message: "Skill not found"})
		return
	}

	c.JSON(http.StatusOK, ToSkillResponse(*skill))
}
