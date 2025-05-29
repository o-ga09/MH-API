package armors

import (
	"context"
	"errors"
	"mh-api/internal/domain/armors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArmorService_GetAllArmors_Success(t *testing.T) {
	ctx := context.Background()
	mockQueryService := &IArmorQueryServiceMock{
		GetAllFunc: func(ctx context.Context) (armors.Armors, error) {
			skills1 := []armors.Skill{
				*armors.NewSkill("1", "攻撃LV1"),
				*armors.NewSkill("2", "火属性攻撃強化LV1"),
			}
			requiredItems1 := []armors.RequiredItem{
				*armors.NewRequiredItem("ITM0019", "リオレウスの鱗"),
				*armors.NewRequiredItem("ITM0016", "ドラグライト鉱石"),
			}

			skills2 := []armors.Skill{
				*armors.NewSkill("1", "攻撃LV2"),
				*armors.NewSkill("3", "体力増強LV1"),
			}
			requiredItems2 := []armors.RequiredItem{
				*armors.NewRequiredItem("ITM0019", "リオレウスの鱗"),
				*armors.NewRequiredItem("ITM0017", "大地の結晶"),
			}

			dummyArmors := armors.Armors{
				*armors.NewArmor("1", "レウスヘルム", "①②③", 100, 10, 5, -10, 5, 15, skills1, requiredItems1),
				*armors.NewArmor("2", "レウスメイル", "①①②", 120, 15, 0, -5, 0, 20, skills2, requiredItems2),
			}
			return dummyArmors, nil
		},
	}

	armorService := NewArmorService(mockQueryService)

	result, err := armorService.GetAllArmors(ctx)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result.Armors))
	assert.Equal(t, "レウスヘルム", result.Armors[0].Name)
	assert.Equal(t, "レウスメイル", result.Armors[1].Name)

	// スキルの確認
	assert.Equal(t, 2, len(result.Armors[0].Skill))
	assert.Equal(t, "攻撃LV1", result.Armors[0].Skill[0].Name)

	// 必要素材の確認
	assert.Equal(t, 2, len(result.Armors[0].Required))
	assert.Equal(t, "リオレウスの鱗", result.Armors[0].Required[0].Name)

	// 防御力・耐性の確認
	assert.Equal(t, 100, result.Armors[0].Defense)
	assert.Equal(t, 10, result.Armors[0].Resistance.Fire)
	assert.Equal(t, 5, result.Armors[0].Resistance.Water)
	assert.Equal(t, -10, result.Armors[0].Resistance.Lightning)
	assert.Equal(t, 5, result.Armors[0].Resistance.Ice)
	assert.Equal(t, 15, result.Armors[0].Resistance.Dragon)
}

func TestArmorService_GetAllArmors_Empty(t *testing.T) {
	ctx := context.Background()
	mockQueryService := &IArmorQueryServiceMock{
		GetAllFunc: func(ctx context.Context) (armors.Armors, error) {
			return armors.Armors{}, nil
		},
	}

	armorService := NewArmorService(mockQueryService)

	result, err := armorService.GetAllArmors(ctx)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result.Armors))
}

func TestArmorService_GetAllArmors_QueryServiceError(t *testing.T) {
	ctx := context.Background()
	mockQueryService := &IArmorQueryServiceMock{
		GetAllFunc: func(ctx context.Context) (armors.Armors, error) {
			return nil, errors.New("データベース接続エラー")
		},
	}

	armorService := NewArmorService(mockQueryService)

	result, err := armorService.GetAllArmors(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "データベース接続エラー", err.Error())
}

func TestArmorService_GetArmorByID_Success(t *testing.T) {
	ctx := context.Background()
	targetArmorID := "1"

	mockQueryService := &IArmorQueryServiceMock{
		GetByIDFunc: func(ctx context.Context, armorId string) (*armors.Armor, error) {
			if armorId == targetArmorID {
				skills := []armors.Skill{
					*armors.NewSkill("1", "攻撃LV1"),
					*armors.NewSkill("2", "火属性攻撃強化LV1"),
				}
				requiredItems := []armors.RequiredItem{
					*armors.NewRequiredItem("ITM0019", "リオレウスの鱗"),
					*armors.NewRequiredItem("ITM0016", "ドラグライト鉱石"),
				}

				return armors.NewArmor("1", "レウスヘルム", "①②③", 100, 10, 5, -10, 5, 15, skills, requiredItems), nil
			}
			return nil, errors.New("armor not found")
		},
	}

	armorService := NewArmorService(mockQueryService)

	result, err := armorService.GetArmorByID(ctx, targetArmorID)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "レウスヘルム", result.Name)
	assert.Equal(t, "①②③", result.Slot)
	assert.Equal(t, 100, result.Defense)

	// スキルの確認
	assert.Equal(t, 2, len(result.Skill))
	assert.Equal(t, "攻撃LV1", result.Skill[0].Name)

	// 必要素材の確認
	assert.Equal(t, 2, len(result.Required))
	assert.Equal(t, "リオレウスの鱗", result.Required[0].Name)

	// 耐性の確認
	assert.Equal(t, 10, result.Resistance.Fire)
	assert.Equal(t, 5, result.Resistance.Water)
	assert.Equal(t, -10, result.Resistance.Lightning)
	assert.Equal(t, 5, result.Resistance.Ice)
	assert.Equal(t, 15, result.Resistance.Dragon)
}

func TestArmorService_GetArmorByID_NotFound(t *testing.T) {
	ctx := context.Background()
	mockQueryService := &IArmorQueryServiceMock{
		GetByIDFunc: func(ctx context.Context, armorId string) (*armors.Armor, error) {
			return nil, errors.New("armor not found")
		},
	}

	armorService := NewArmorService(mockQueryService)

	result, err := armorService.GetArmorByID(ctx, "non-existent-id")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "armor not found", err.Error())
}

func TestArmorService_GetArmorByID_QueryServiceError(t *testing.T) {
	ctx := context.Background()
	mockQueryService := &IArmorQueryServiceMock{
		GetByIDFunc: func(ctx context.Context, armorId string) (*armors.Armor, error) {
			return nil, errors.New("データベース接続エラー")
		},
	}

	armorService := NewArmorService(mockQueryService)

	result, err := armorService.GetArmorByID(ctx, "1")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "データベース接続エラー", err.Error())
}
