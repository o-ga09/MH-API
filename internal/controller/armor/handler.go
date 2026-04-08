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
// @Description 全ての防具のリストを返します。
// @Tags Armors
// @Accept json
// @Produce json
// @Success 200 {object} ListArmorsResponse "防具のリスト"
// @Failure 500 {object} ErrorResponse "サーバ内部エラー"
// @Router /armors [get]
func (h *ArmorHandler) GetAllArmors(c *gin.Context) {
armorList, err := h.repo.GetAll(c.Request.Context())
if err != nil {
c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to get armors: " + err.Error()})
return
}

c.JSON(http.StatusOK, ListArmorsResponse{
Armors: toArmorDetailResponses(armorList),
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
