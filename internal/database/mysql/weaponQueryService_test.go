package mysql

import (
"context"
"errors"
"testing"

"mh-api/internal/domain/weapons"
"mh-api/pkg/ptr"

"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/require"
"gorm.io/gorm"
)

func createTestWeapons(t *testing.T, ctx context.Context) []*weapons.Weapon {
testWeapons := []*weapons.Weapon{
{WeaponID: "weapon001", Name: "炎剣リオレウス", ImageUrl: "https://example.com/weapon001.png", Rarerity: "5", Attack: "100", ElementAttack: "30", Shapness: "40", Critical: "20", Description: "炎の力を宿した剣"},
{WeaponID: "weapon002", Name: "氷剣イヴェルカーナ", ImageUrl: "https://example.com/weapon002.png", Rarerity: "5", Attack: "110", ElementAttack: "40", Shapness: "30", Critical: "25", Description: "氷の力を宿した剣"},
{WeaponID: "weapon003", Name: "雷剣ジンオウガ", ImageUrl: "https://example.com/weapon003.png", Rarerity: "5", Attack: "105", ElementAttack: "35", Shapness: "45", Critical: "30", Description: "雷の力を宿した剣"},
}

db := CtxFromTestDB(ctx)
for _, weapon := range testWeapons {
if err := db.Create(weapon).Error; err != nil {
t.Fatalf("failed to create test weapon: %v", err)
}
}
return testWeapons
}

func TestWeaponRepository_Find(t *testing.T) {
ctx := context.Background()
ctx = setupTestDB(ctx)
testDB.Begin()
defer testDB.Rollback()

_ = createTestWeapons(t, ctx)

tests := []struct {
name      string
params    weapons.SearchParams
wantCount int
wantErr   bool
}{
{
name:      "正常系: すべての武器を取得",
params:    weapons.SearchParams{Limit: ptr.IntToPtr(10)},
wantCount: 3,
},
{
name:      "正常系: 武器IDで絞り込み",
params:    weapons.SearchParams{WeaponID: ptr.StrToPtr("weapon001")},
wantCount: 1,
},
{
name:      "正常系: 名前で部分一致検索",
params:    weapons.SearchParams{Name: ptr.StrToPtr("剣"), Limit: ptr.IntToPtr(10)},
wantCount: 3,
},
{
name:      "正常系: limitで件数制限",
params:    weapons.SearchParams{Limit: ptr.IntToPtr(1)},
wantCount: 1,
},
{
name:      "正常系: offsetでスキップ",
params:    weapons.SearchParams{Offset: ptr.IntToPtr(1), Limit: ptr.IntToPtr(10)},
wantCount: 2,
},
{
name:      "正常系: 昇順ソート",
params:    weapons.SearchParams{Sort: ptr.StrToPtr("asc"), Limit: ptr.IntToPtr(10)},
wantCount: 3,
},
{
name:      "正常系: 降順ソート",
params:    weapons.SearchParams{Sort: ptr.StrToPtr("desc"), Limit: ptr.IntToPtr(10)},
wantCount: 3,
},
{
name:      "正常系: 検索結果0件",
params:    weapons.SearchParams{Name: ptr.StrToPtr("存在しない武器")},
wantCount: 0,
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
repo := NewWeaponRepository()
got, err := repo.Find(ctx, tt.params)

if tt.wantErr {
assert.Error(t, err)
return
}
require.NoError(t, err)
assert.Len(t, got.Weapons, tt.wantCount)
})
}
}

func TestWeaponRepository_FindByID(t *testing.T) {
ctx := context.Background()
ctx = setupTestDB(ctx)
testDB.Begin()
defer testDB.Rollback()

testWeapons := createTestWeapons(t, ctx)

tests := []struct {
name     string
weaponID string
want     *weapons.Weapon
wantErr  bool
errType  error
}{
{
name:     "正常系: 存在するIDの場合",
weaponID: testWeapons[0].WeaponID,
want:     testWeapons[0],
wantErr:  false,
},
{
name:     "異常系: 存在しないIDの場合",
weaponID: "non-existent-id",
wantErr:  true,
errType:  gorm.ErrRecordNotFound,
},
{
name:     "異常系: 空のIDの場合",
weaponID: "",
wantErr:  true,
errType:  gorm.ErrRecordNotFound,
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
repo := NewWeaponRepository().(*weaponRepository)
got, err := repo.FindByID(ctx, tt.weaponID)

if tt.wantErr {
require.Error(t, err)
if tt.errType != nil {
assert.True(t, errors.Is(err, tt.errType))
}
return
}
require.NoError(t, err)
assert.Equal(t, tt.want.WeaponID, got.WeaponID)
assert.Equal(t, tt.want.Name, got.Name)
})
}
}
