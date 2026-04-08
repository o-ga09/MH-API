package mysql

import (
"context"
"errors"
"testing"

"mh-api/internal/domain/armors"

"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/require"
"gorm.io/gorm"
)

func createTestArmors(t *testing.T, ctx context.Context) []*armors.Armor {
testArmors := []*armors.Armor{
{ArmorId: "armor001", Name: "レウスヘルム", Slot: "①②③", Defense: 100, FireResistance: 10, WaterResistance: 5, LightningResistance: -10, IceResistance: 5, DragonResistance: 15},
{ArmorId: "armor002", Name: "レウスメイル", Slot: "①①②", Defense: 120, FireResistance: 15, WaterResistance: 0, LightningResistance: -5, IceResistance: 0, DragonResistance: 20},
{ArmorId: "armor003", Name: "レウスアーム", Slot: "①②", Defense: 80, FireResistance: 8, WaterResistance: 3, LightningResistance: -8, IceResistance: 3, DragonResistance: 12},
}

testSkills := []*armors.ArmorSkill{
{ArmorId: "armor001", SkillId: "1", SkillName: "攻撃LV1"},
{ArmorId: "armor001", SkillId: "2", SkillName: "火属性攻撃強化LV1"},
{ArmorId: "armor002", SkillId: "1", SkillName: "攻撃LV2"},
{ArmorId: "armor002", SkillId: "3", SkillName: "体力増強LV1"},
{ArmorId: "armor003", SkillId: "4", SkillName: "見切りLV1"},
}

testRequiredItems := []*armors.ArmorRequiredItem{
{ArmorId: "armor001", ItemId: "ITM0019", ItemName: "リオレウスの鱗"},
{ArmorId: "armor001", ItemId: "ITM0016", ItemName: "ドラグライト鉱石"},
{ArmorId: "armor002", ItemId: "ITM0019", ItemName: "リオレウスの鱗"},
{ArmorId: "armor002", ItemId: "ITM0017", ItemName: "大地の結晶"},
{ArmorId: "armor003", ItemId: "ITM0019", ItemName: "リオレウスの鱗"},
}

db := CtxFromTestDB(ctx)
for _, armor := range testArmors {
require.NoError(t, db.Create(armor).Error)
}
for _, skill := range testSkills {
require.NoError(t, db.Create(skill).Error)
}
for _, item := range testRequiredItems {
require.NoError(t, db.Create(item).Error)
}

return testArmors
}

func TestArmorRepository_GetAll(t *testing.T) {
ctx := t.Context()
ctx = setupTestDB(ctx)
db := ctx.Value(CtxKey).(*gorm.DB)
db.Begin()
defer db.Rollback()

_ = createTestArmors(t, ctx)

tests := []struct {
name      string
wantCount int
wantErr   bool
}{
{name: "正常系: すべての防具を取得", wantCount: 3},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
repo := NewArmorRepository()
got, err := repo.GetAll(ctx)

if tt.wantErr {
assert.Error(t, err)
return
}
require.NoError(t, err)
assert.Len(t, got, tt.wantCount)

if len(got) > 0 {
first := got[0]
assert.NotEmpty(t, first.ArmorId)
assert.NotEmpty(t, first.Name)
assert.NotEmpty(t, first.Slot)
assert.Greater(t, first.Defense, 0)
assert.GreaterOrEqual(t, len(first.Skills), 1)
assert.GreaterOrEqual(t, len(first.RequiredItems), 1)
}
})
}
}

func TestArmorRepository_GetByID(t *testing.T) {
ctx := t.Context()
ctx = setupTestDB(ctx)
db := ctx.Value(CtxKey).(*gorm.DB)
db.Begin()
defer db.Rollback()

testArmors := createTestArmors(t, ctx)

tests := []struct {
name    string
armorID string
want    *armors.Armor
wantErr bool
errType error
}{
{
name:    "正常系: 存在するIDの場合",
armorID: testArmors[0].ArmorId,
want:    testArmors[0],
},
{
name:    "異常系: 存在しないIDの場合",
armorID: "non-existent-id",
wantErr: true,
errType: gorm.ErrRecordNotFound,
},
{
name:    "異常系: 空のIDの場合",
armorID: "",
wantErr: true,
errType: gorm.ErrRecordNotFound,
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
repo := NewArmorRepository()
got, err := repo.GetByID(ctx, tt.armorID)

if tt.wantErr {
require.Error(t, err)
if tt.errType != nil {
assert.True(t, errors.Is(err, tt.errType))
}
return
}
require.NoError(t, err)
assert.Equal(t, tt.want.ArmorId, got.ArmorId)
assert.Equal(t, tt.want.Name, got.Name)
assert.Equal(t, tt.want.Slot, got.Slot)
assert.Equal(t, tt.want.Defense, got.Defense)
assert.Equal(t, tt.want.FireResistance, got.FireResistance)
assert.Equal(t, tt.want.WaterResistance, got.WaterResistance)
assert.Equal(t, tt.want.LightningResistance, got.LightningResistance)
assert.Equal(t, tt.want.IceResistance, got.IceResistance)
assert.Equal(t, tt.want.DragonResistance, got.DragonResistance)
assert.GreaterOrEqual(t, len(got.Skills), 1)
assert.GreaterOrEqual(t, len(got.RequiredItems), 1)
})
}
}
