package weapon

import (
"context"
"errors"
"net/http"
"net/http/httptest"
"testing"

"mh-api/internal/domain/weapons"
"mh-api/pkg/testutil"

"github.com/gin-gonic/gin"
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/require"
)

func setupTestRouter(t *testing.T, repo weapons.Repository) *gin.Engine {
t.Helper()
gin.SetMode(gin.TestMode)
r := gin.New()
handler := NewWeaponHandler(repo)
g := r.Group("/weapons")
g.GET("", handler.SearchWeapons)
return r
}

func TestWeaponHandler_SearchWeapons(t *testing.T) {
defaultLimit := 20
defaultOffset := 0

tests := []struct {
name       string
mockSetup  func() weapons.Repository
query      string
wantStatus int
goldenFile string
}{
{
name: "正常系: 武器一覧が取得できる",
mockSetup: func() weapons.Repository {
return &weapons.RepositoryMock{
FindFunc: func(ctx context.Context, params weapons.SearchParams) (*weapons.SearchResult, error) {
assert.Equal(t, defaultLimit, *params.Limit)
assert.Equal(t, defaultOffset, *params.Offset)
return createTestWeaponsResult(), nil
},
}
},
query:      "?limit=20&offset=0",
wantStatus: http.StatusOK,
goldenFile: "weapon/weapon_search_200.json",
},
{
name: "異常系: バリデーションエラー",
mockSetup: func() weapons.Repository {
return &weapons.RepositoryMock{
FindFunc: func(ctx context.Context, params weapons.SearchParams) (*weapons.SearchResult, error) {
return nil, nil
},
}
},
query:      "?limit=-1",
wantStatus: http.StatusBadRequest,
goldenFile: "weapon/weapon_search_400.json",
},
{
name: "異常系: 内部エラー",
mockSetup: func() weapons.Repository {
return &weapons.RepositoryMock{
FindFunc: func(ctx context.Context, params weapons.SearchParams) (*weapons.SearchResult, error) {
return nil, errors.New("database error")
},
}
},
query:      "?limit=20&offset=0",
wantStatus: http.StatusInternalServerError,
goldenFile: "weapon/weapon_search_500.json",
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
router := setupTestRouter(t, tt.mockSetup())
req, err := http.NewRequest(http.MethodGet, "/weapons"+tt.query, nil)
require.NoError(t, err)
w := httptest.NewRecorder()
router.ServeHTTP(w, req)
assert.Equal(t, tt.wantStatus, w.Code)
testutil.AssertGoldenJSON(t, tt.goldenFile, w.Body.Bytes())
})
}
}

func createTestWeaponsResult() *weapons.SearchResult {
return &weapons.SearchResult{
Weapons: []*weapons.Weapon{
{
WeaponID:      "1",
Name:          "リオレウス剣",
ImageUrl:      "http://example.com/weapon1.jpg",
Rarerity:      "5",
Attack:        "800",
ElementAttack: "火 200",
Shapness:      "青",
Critical:      "0%",
Description:   "リオレウスの素材から作られた剣",
},
{
WeaponID:      "2",
Name:          "ジンオウガ太刀",
ImageUrl:      "http://example.com/weapon2.jpg",
Rarerity:      "6",
Attack:        "750",
ElementAttack: "雷 300",
Shapness:      "白",
Critical:      "10%",
Description:   "ジンオウガの素材から作られた太刀",
},
},
TotalCount: 2,
Offset:     0,
Limit:      20,
}
}
