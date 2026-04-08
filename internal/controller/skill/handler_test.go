package skill

import (
"context"
"errors"
"net/http"
"net/http/httptest"
"testing"

"mh-api/internal/domain/skills"
"mh-api/pkg/testutil"

"github.com/gin-gonic/gin"
"github.com/stretchr/testify/assert"
"gorm.io/gorm"
)

func MockNewServer(t *testing.T, handler *SkillHandler) *gin.Engine {
t.Helper()
gin.SetMode(gin.TestMode)
return gin.New()
}

func (h *SkillHandler) SetupRouter(r *gin.Engine) {
r.GET("/v1/skills", h.GetSkills)
r.GET("/v1/skills/:skillId", h.GetSkill)
}

func TestSkillHandler_GetSkills(t *testing.T) {
tests := []struct {
name       string
mock       func() skills.Repository
wantStatus int
goldenFile string
}{
{
name: "正常系：スキルの一覧が取得できる",
mock: func() skills.Repository {
return &skills.RepositoryMock{
FindAllFunc: func(ctx context.Context) (skills.Skills, error) {
return skills.Skills{
{
SkillId:     "0000000001",
Name:        "攻撃",
Description: "攻撃力が上昇する",
Levels: []skills.SkillLevel{
{Level: 1, Description: "攻撃力+3"},
{Level: 2, Description: "攻撃力+6"},
},
},
{
SkillId:     "0000000002",
Name:        "防御",
Description: "防御力が上昇する",
Levels: []skills.SkillLevel{
{Level: 1, Description: "防御力+5"},
},
},
}, nil
},
FindByIdFunc: func(ctx context.Context, skillId string) (*skills.Skill, error) {
return nil, nil
},
}
},
wantStatus: http.StatusOK,
goldenFile: "skills/get_skills_success.json.golden",
},
{
name: "異常系：リポジトリでエラーが発生する",
mock: func() skills.Repository {
return &skills.RepositoryMock{
FindAllFunc: func(ctx context.Context) (skills.Skills, error) {
return nil, errors.New("repository error")
},
FindByIdFunc: func(ctx context.Context, skillId string) (*skills.Skill, error) {
return nil, nil
},
}
},
wantStatus: http.StatusInternalServerError,
goldenFile: "skills/get_skills_error.json.golden",
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
handler := NewSkillHandler(tt.mock())
router := MockNewServer(t, handler)
handler.SetupRouter(router)

req, _ := http.NewRequest("GET", "/v1/skills", nil)
rec := httptest.NewRecorder()
router.ServeHTTP(rec, req)

assert.Equal(t, tt.wantStatus, rec.Code)
testutil.AssertGoldenJSON(t, tt.goldenFile, rec.Body.Bytes())
})
}
}

func TestSkillHandler_GetSkill(t *testing.T) {
tests := []struct {
name       string
skillId    string
mock       func() skills.Repository
wantStatus int
goldenFile string
}{
{
name:    "正常系：存在するスキルIDの場合",
skillId: "0000000001",
mock: func() skills.Repository {
return &skills.RepositoryMock{
FindAllFunc: func(ctx context.Context) (skills.Skills, error) {
return nil, nil
},
FindByIdFunc: func(ctx context.Context, skillId string) (*skills.Skill, error) {
return &skills.Skill{
SkillId:     "0000000001",
Name:        "攻撃",
Description: "攻撃力が上昇する",
Levels: []skills.SkillLevel{
{Level: 1, Description: "攻撃力+3"},
{Level: 2, Description: "攻撃力+6"},
},
}, nil
},
}
},
wantStatus: http.StatusOK,
goldenFile: "skills/get_skill_success.json.golden",
},
{
name:    "準正常系：不正なスキルIDの場合",
skillId: " ",
mock: func() skills.Repository {
return &skills.RepositoryMock{
FindAllFunc: func(ctx context.Context) (skills.Skills, error) {
return nil, nil
},
FindByIdFunc: func(ctx context.Context, skillId string) (*skills.Skill, error) {
return nil, nil
},
}
},
wantStatus: http.StatusBadRequest,
goldenFile: "skills/get_skill_bad_request.json.golden",
},
{
name:    "異常系：存在しないスキルIDの場合",
skillId: "9999999999",
mock: func() skills.Repository {
return &skills.RepositoryMock{
FindAllFunc: func(ctx context.Context) (skills.Skills, error) {
return nil, nil
},
FindByIdFunc: func(ctx context.Context, skillId string) (*skills.Skill, error) {
return nil, gorm.ErrRecordNotFound
},
}
},
wantStatus: http.StatusNotFound,
goldenFile: "skills/get_skill_error.json.golden",
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
handler := NewSkillHandler(tt.mock())
router := MockNewServer(t, handler)
handler.SetupRouter(router)

url := "/v1/skills/" + tt.skillId
req, _ := http.NewRequest("GET", url, nil)
rec := httptest.NewRecorder()
router.ServeHTTP(rec, req)

assert.Equal(t, tt.wantStatus, rec.Code)
testutil.AssertGoldenJSON(t, tt.goldenFile, rec.Body.Bytes())
})
}
}
