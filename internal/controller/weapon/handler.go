package weapon

import (
	"context"
	"mh-api/internal/service/weapons"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IWeaponService interface {
	SearchWeapons(ctx context.Context, params weapons.SearchWeaponsParams) (*weapons.ListWeaponsResponse, error)
}

type WeaponHandler struct {
	service IWeaponService
}

func NewWeaponHandler(s IWeaponService) *WeaponHandler {
	return &WeaponHandler{service: s}
}

// SearchWeapons godoc
// @Summary 武器リストを検索します (Gin版)
// @Description 指定されたクエリパラメータに基づいて武器のリストを返します。
// @Tags Weapons
// @Accept json
// @Produce json
// @Param monster_id query string false "武器ID (完全一致)"
// @Param name query string false "武器名 (部分一致を想定)"
// @Param name_kana query string false "武器名かな (部分一致を想定)"
// @Param limit query int false "取得件数" default(20)
// @Param offset query int false "取得開始位置" default(0)
// @Param sort query string false "ソート対象フィールド (例: attack, rare)"
// @Param order query int false "ソート順 (0:昇順, 1:降順)"
// @Success 200 {object} ListWeaponsResponse "武器のリストとページネーション情報" // response.goで定義したListWeaponsResponse
// @Failure 400 {object} ErrorResponse "リクエストパラメータが不正な場合"
// @Failure 500 {object} ErrorResponse "サーバ内部エラー"
// @Router /weapons [get]
func (h *WeaponHandler) SearchWeapons(c *gin.Context) { // 引数を *gin.Context に変更
	appCtx := c.Request.Context() // Ginのコンテキストから標準のcontext.Contextを取得

	var req SearchWeaponsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request parameters: " + err.Error()}) // Ginのレスポンス方法
		return
	}

	serviceParams := weapons.SearchWeaponsParams{
		Limit:    req.Limit,
		Offset:   req.Offset,
		Sort:     req.Sort,
		Order:    req.Order,
		WeaponID: req.WeaponID,
		Name:     req.Name,
	}

	serviceRes, err := h.service.SearchWeapons(appCtx, serviceParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to search weapons: " + err.Error()})
		return
	}

	ctrlResponseWeapons := make([]WeaponDetailResponse, len(serviceRes.Weapons))
	for i, sw := range serviceRes.Weapons {
		ctrlResponseWeapons[i] = WeaponDetailResponse{
			WeaponID:      sw.WeaponID,
			Name:          sw.Name,
			ImageURL:      sw.ImageURL,
			Rare:          sw.Rare,
			Attack:        sw.Attack,
			ElementAttack: sw.ElementAttack,
			Sharpness:     sw.Sharpness,
			Critical:      sw.Critical,
			Description:   sw.Description,
		}
	}

	ctrlResponse := ListWeaponsResponse{
		TotalCount: serviceRes.TotalCount,
		Limit:      serviceRes.Limit,
		Offset:     serviceRes.Offset,
		Weapons:    ctrlResponseWeapons,
	}

	c.JSON(http.StatusOK, ctrlResponse)
}

type ErrorResponse struct {
	Message string `json:"message"`
}
