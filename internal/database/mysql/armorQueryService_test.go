package mysql

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func createTestArmors(t *testing.T, ctx context.Context) []*Armor {
	testArmors := []*Armor{
		{
			ArmorId:             "armor001",
			Name:                "レウスヘルム",
			Slot:                "①②③",
			Defense:             100,
			FireResistance:      10,
			WaterResistance:     5,
			LightningResistance: -10,
			IceResistance:       5,
			DragonResistance:    15,
		},
		{
			ArmorId:             "armor002",
			Name:                "レウスメイル",
			Slot:                "①①②",
			Defense:             120,
			FireResistance:      15,
			WaterResistance:     0,
			LightningResistance: -5,
			IceResistance:       0,
			DragonResistance:    20,
		},
		{
			ArmorId:             "armor003",
			Name:                "レウスアーム",
			Slot:                "①②",
			Defense:             80,
			FireResistance:      8,
			WaterResistance:     3,
			LightningResistance: -8,
			IceResistance:       3,
			DragonResistance:    12,
		},
	}

	testSkills := []*ArmorSkill{
		{ArmorId: "armor001", SkillId: "1", SkillName: "攻撃LV1"},
		{ArmorId: "armor001", SkillId: "2", SkillName: "火属性攻撃強化LV1"},
		{ArmorId: "armor002", SkillId: "1", SkillName: "攻撃LV2"},
		{ArmorId: "armor002", SkillId: "3", SkillName: "体力増強LV1"},
		{ArmorId: "armor003", SkillId: "4", SkillName: "見切りLV1"},
	}

	testRequiredItems := []*ArmorRequiredItem{
		{ArmorId: "armor001", ItemId: "ITM0019", ItemName: "リオレウスの鱗"},
		{ArmorId: "armor001", ItemId: "ITM0016", ItemName: "ドラグライト鉱石"},
		{ArmorId: "armor002", ItemId: "ITM0019", ItemName: "リオレウスの鱗"},
		{ArmorId: "armor002", ItemId: "ITM0017", ItemName: "大地の結晶"},
		{ArmorId: "armor003", ItemId: "ITM0019", ItemName: "リオレウスの鱗"},
	}

	db := CtxFromTestDB(ctx)

	for _, armor := range testArmors {
		if err := db.Create(armor).Error; err != nil {
			t.Fatalf("failed to create test armor: %v", err)
		}
	}

	for _, skill := range testSkills {
		if err := db.Create(skill).Error; err != nil {
			t.Fatalf("failed to create test armor skill: %v", err)
		}
	}

	for _, item := range testRequiredItems {
		if err := db.Create(item).Error; err != nil {
			t.Fatalf("failed to create test armor required item: %v", err)
		}
	}

	return testArmors
}

func TestArmorQueryService_GetAll(t *testing.T) {
	ctx := context.Background()

	ctx = setupTestDB(ctx)
	testDB.Begin()
	defer testDB.Rollback()

	_ = createTestArmors(t, ctx)

	tests := []struct {
		name      string
		wantCount int
		wantErr   bool
	}{
		{
			name:      "正常系: すべての防具を取得",
			wantCount: 3,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := NewArmorQueryService()

			got, err := qs.GetAll(ctx)

			assert.True(t, (err != nil) == tt.wantErr)
			if !tt.wantErr {
				assert.Len(t, got, tt.wantCount)
				
				// 最初の防具の詳細をチェック
				if len(got) > 0 {
					firstArmor := got[0]
					assert.NotEmpty(t, firstArmor.GetID())
					assert.NotEmpty(t, firstArmor.GetName())
					assert.NotEmpty(t, firstArmor.GetSlot())
					assert.Greater(t, firstArmor.GetDefense(), 0)
					
					// スキルがロードされているかチェック
					skills := firstArmor.GetSkills()
					assert.GreaterOrEqual(t, len(skills), 1)
					
					// 必要素材がロードされているかチェック
					required := firstArmor.GetRequiredItems()
					assert.GreaterOrEqual(t, len(required), 1)
				}
			}
		})
	}
}

func TestArmorQueryService_GetByID(t *testing.T) {
	ctx := context.Background()

	ctx = setupTestDB(ctx)
	testDB.Begin()
	defer testDB.Rollback()

	testArmors := createTestArmors(t, ctx)

	tests := []struct {
		name    string
		armorID string
		want    *Armor
		wantErr bool
		errType error
	}{
		{
			name:    "正常系: 存在するIDの場合",
			armorID: testArmors[0].ArmorId,
			want:    testArmors[0],
			wantErr: false,
			errType: nil,
		},
		{
			name:    "異常系: 存在しないIDの場合",
			armorID: "non-existent-id",
			want:    nil,
			wantErr: true,
			errType: gorm.ErrRecordNotFound,
		},
		{
			name:    "異常系: 空のIDの場合",
			armorID: "",
			want:    nil,
			wantErr: true,
			errType: gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := NewArmorQueryService()

			got, err := qs.GetByID(ctx, tt.armorID)

			if tt.wantErr {
				require.Error(t, err)
				if tt.errType != nil {
					assert.True(t, errors.Is(err, tt.errType), "expected error type: %v, got: %v", tt.errType, err)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want.ArmorId, got.GetID())
				assert.Equal(t, tt.want.Name, got.GetName())
				assert.Equal(t, tt.want.Slot, got.GetSlot())
				assert.Equal(t, tt.want.Defense, got.GetDefense())
				assert.Equal(t, tt.want.FireResistance, got.GetFireResistance())
				assert.Equal(t, tt.want.WaterResistance, got.GetWaterResistance())
				assert.Equal(t, tt.want.LightningResistance, got.GetLightningResistance())
				assert.Equal(t, tt.want.IceResistance, got.GetIceResistance())
				assert.Equal(t, tt.want.DragonResistance, got.GetDragonResistance())
				
				// スキルがロードされているかチェック
				skills := got.GetSkills()
				assert.GreaterOrEqual(t, len(skills), 1)
				
				// 必要素材がロードされているかチェック
				required := got.GetRequiredItems()
				assert.GreaterOrEqual(t, len(required), 1)
			}
		})
	}
}