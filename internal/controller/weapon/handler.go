package weapon

import (
"context"
"mh-api/internal/domain/weapons"
"net/http"

"github.com/gin-gonic/gin"
)

type WeaponHandler struct {
repo weapons.Repository
}

func NewWeaponHandler(repo weapons.Repository) *WeaponHandler {
return &WeaponHandler{repo: repo}
}

// SearchWeapons godoc
// @Summary 武器リストを検索します
// @Description 指定されたクエリパラメータに基づいて武器のリストを返します。
// @Tags Weapons
// @Accept json
// @Produce json
// @Param weapon_id query string false "武器ID (完全一致)"
// @Param name query string false "武器名 (部分一致を想定)"
// @Param limit query int false "取得件数" default(20)
// @Param offset query int false "取得開始位置" default(0)
// @Param sort query string false "ソートフィールド (asc/desc)"
// @Param order query int false "ソート順 (0:昇順, 1:降順)"
// @Success 200 {object} ListWeaponsResponse "武器のリストとページネーション情報"
// @Failure 400 {object} ErrorResponse "リクエストパラメータが不正な場合"
// @Failure 500 {object} ErrorResponse "サーバ内部エラー"
// @Router /weapons [get]
func (h *WeaponHandler) SearchWeapons(c *gin.Context) {
var req SearchWeaponsRequest
if err := c.ShouldBindQuery(&req); err != nil {
c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request parameters: " + err.Error()})
return
}

params := weapons.SearchParams{
Limit:    req.Limit,
Offset:   req.Offset,
Sort:     req.Sort,
Order:    req.Order,
WeaponID: req.WeaponID,
Name:     req.Name,
}

result, err := h.repo.Find(context.Background(), params)
if err != nil {
c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to search weapons: " + err.Error()})
return
}

weaponResponses := make([]WeaponDetailResponse, len(result.Weapons))
for i, w := range result.Weapons {
weaponResponses[i] = WeaponDetailResponse{
WeaponID:      w.WeaponID,
Name:          w.Name,
ImageURL:      w.ImageUrl,
Rare:          w.Rarerity,
Attack:        w.Attack,
ElementAttack: w.ElementAttack,
Sharpness:     w.Shapness,
Critical:      w.Critical,
Description:   w.Description,
}
}

c.JSON(http.StatusOK, ListWeaponsResponse{
TotalCount: result.TotalCount,
Limit:      result.Limit,
Offset:     result.Offset,
Weapons:    weaponResponses,
})
}

type ErrorResponse struct {
Message string `json:"message"`
}
