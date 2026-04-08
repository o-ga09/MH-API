package armor

import (
"context"
"errors"
"net/http"
"net/http/httptest"
"testing"

"mh-api/internal/domain/armors"
"mh-api/pkg/testutil"

"github.com/gin-gonic/gin"
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/require"
"gorm.io/gorm"
)

func setupTestRouter(t *testing.T, repo armors.Repository) *gin.Engine {
t.Helper()
gin.SetMode(gin.TestMode)
r := gin.New()
handler := NewArmorHandler(repo)
g := r.Group("/armors")
g.GET("", handler.GetAllArmors)
g.GET("/:id", handler.GetArmorByID)
return r
}

func TestArmorHandler_GetAllArmors(t *testing.T) {
tests := []struct {
name       string
mockSetup  func() armors.Repository
wantStatus int
goldenFile string
}{
{
name: "正常系: 防具一覧が取得できる",
mockSetup: func() armors.Repository {
return &armors.RepositoryMock{
GetAllFunc: func(ctx context.Context) (armors.Armors, error) {
return createTestArmors(), nil
},
GetByIDFunc: func(ctx context.Context, armorId string) (*armors.Armor, error) {
return nil, nil
},
}
},
wantStatus: http.StatusOK,
goldenFile: "armor/armor_get_all_200.json.golden",
},
{
name: "正常系: 防具一覧が空の場合",
mockSetup: func() armors.Repository {
return &armors.RepositoryMock{
GetAllFunc: func(ctx context.Context) (armors.Armors, error) {
return armors.Armors{}, nil
},
GetByIDFunc: func(ctx context.Context, armorId string) (*armors.Armor, error) {
return nil, nil
},
}
},
wantStatus: http.StatusOK,
goldenFile: "armor/armor_get_all_empty.json.golden",
},
{
name: "異常系: 内部エラー",
mockSetup: func() armors.Repository {
return &armors.RepositoryMock{
GetAllFunc: func(ctx context.Context) (armors.Armors, error) {
return nil, errors.New("database error")
},
GetByIDFunc: func(ctx context.Context, armorId string) (*armors.Armor, error) {
return nil, nil
},
}
},
wantStatus: http.StatusInternalServerError,
goldenFile: "armor/armor_get_all_500.json.golden",
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
router := setupTestRouter(t, tt.mockSetup())
req, err := http.NewRequest(http.MethodGet, "/armors", nil)
require.NoError(t, err)
w := httptest.NewRecorder()
router.ServeHTTP(w, req)
assert.Equal(t, tt.wantStatus, w.Code)
testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
})
}
}

func TestArmorHandler_GetArmorByID(t *testing.T) {
tests := []struct {
name       string
mockSetup  func() armors.Repository
armorID    string
wantStatus int
goldenFile string
}{
{
name: "正常系: 防具詳細が取得できる",
mockSetup: func() armors.Repository {
return &armors.RepositoryMock{
GetAllFunc: func(ctx context.Context) (armors.Armors, error) {
return nil, nil
},
GetByIDFunc: func(ctx context.Context, armorId string) (*armors.Armor, error) {
return createTestArmors()[0], nil
},
}
},
armorID:    "1",
wantStatus: http.StatusOK,
goldenFile: "armor/armor_get_by_id_200.json.golden",
},
{
name: "異常系: バリデーションエラー（IDが空）",
mockSetup: func() armors.Repository {
return &armors.RepositoryMock{
GetAllFunc: func(ctx context.Context) (armors.Armors, error) {
return nil, nil
},
GetByIDFunc: func(ctx context.Context, armorId string) (*armors.Armor, error) {
return nil, nil
},
}
},
armorID:    " ",
wantStatus: http.StatusBadRequest,
goldenFile: "armor/armor_get_by_id_400.json.golden",
},
{
name: "異常系: 防具が見つからない",
mockSetup: func() armors.Repository {
return &armors.RepositoryMock{
GetAllFunc: func(ctx context.Context) (armors.Armors, error) {
return nil, nil
},
GetByIDFunc: func(ctx context.Context, armorId string) (*armors.Armor, error) {
return nil, gorm.ErrRecordNotFound
},
}
},
armorID:    "999",
wantStatus: http.StatusNotFound,
goldenFile: "armor/armor_get_by_id_404.json.golden",
},
{
name: "異常系: 内部エラー",
mockSetup: func() armors.Repository {
return &armors.RepositoryMock{
GetAllFunc: func(ctx context.Context) (armors.Armors, error) {
return nil, nil
},
GetByIDFunc: func(ctx context.Context, armorId string) (*armors.Armor, error) {
return nil, errors.New("database error")
},
}
},
armorID:    "1",
wantStatus: http.StatusInternalServerError,
goldenFile: "armor/armor_get_by_id_500.json.golden",
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
router := setupTestRouter(t, tt.mockSetup())
url := "/armors/" + tt.armorID
req, err := http.NewRequest(http.MethodGet, url, nil)
require.NoError(t, err)
w := httptest.NewRecorder()
router.ServeHTTP(w, req)
assert.Equal(t, tt.wantStatus, w.Code)
testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
})
}
}

func createTestArmors() armors.Armors {
return armors.Armors{
{
ArmorId: "1",
Name:    "レウスヘルム",
Slot:    "①②③",
Defense: 100,
FireResistance:      10,
WaterResistance:     5,
LightningResistance: -10,
IceResistance:       5,
DragonResistance:    15,
Skills: []*armors.ArmorSkill{
{SkillId: "1", SkillName: "攻撃LV1"},
{SkillId: "2", SkillName: "火属性攻撃強化LV1"},
},
RequiredItems: []*armors.ArmorRequiredItem{
{ItemId: "ITM0019", ItemName: "リオレウスの鱗"},
{ItemId: "ITM0016", ItemName: "ドラグライト鉱石"},
},
},
{
ArmorId: "2",
Name:    "レウスメイル",
Slot:    "①①②",
Defense: 120,
FireResistance:      15,
WaterResistance:     0,
LightningResistance: -5,
IceResistance:       0,
DragonResistance:    20,
Skills: []*armors.ArmorSkill{
{SkillId: "1", SkillName: "攻撃LV2"},
{SkillId: "3", SkillName: "体力増強LV1"},
},
RequiredItems: []*armors.ArmorRequiredItem{
{ItemId: "ITM0019", ItemName: "リオレウスの鱗"},
{ItemId: "ITM0017", ItemName: "大地の結晶"},
},
},
}
}
