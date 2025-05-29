package mysql

import (
	"context"
	"errors"
	"mh-api/internal/domain/armors"

	"gorm.io/gorm"
)

type ArmorQueryService struct{}

func NewArmorQueryService() *ArmorQueryService {
	return &ArmorQueryService{}
}

func (qs *ArmorQueryService) GetAll(ctx context.Context) (armors.Armors, error) {
	var armorsList []*Armor

	db := CtxFromDB(ctx)
	result := db.Preload("Skills").Preload("RequiredItems").Find(&armorsList)
	if result.Error != nil {
		return nil, result.Error
	}

	armorDomainList := ToArmorList(armorsList)
	return armorDomainList, nil
}

func (qs *ArmorQueryService) GetByID(ctx context.Context, armorId string) (*armors.Armor, error) {
	var armor Armor
	result := CtxFromDB(ctx).Preload("Skills").Preload("RequiredItems").Where("armor_id = ?", armorId).First(&armor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	res := ToArmor(&armor)
	return res, nil
}

func ToArmor(armor *Armor) *armors.Armor {
	skills := make([]armors.Skill, len(armor.Skills))
	for i, skill := range armor.Skills {
		skills[i] = *armors.NewSkill(skill.SkillId, skill.SkillName)
	}

	requiredItems := make([]armors.RequiredItem, len(armor.RequiredItems))
	for i, item := range armor.RequiredItems {
		requiredItems[i] = *armors.NewRequiredItem(item.ItemId, item.ItemName)
	}

	return armors.NewArmor(
		armor.ArmorId,
		armor.Name,
		armor.Slot,
		armor.Defense,
		armor.FireResistance,
		armor.WaterResistance,
		armor.LightningResistance,
		armor.IceResistance,
		armor.DragonResistance,
		skills,
		requiredItems,
	)
}

func ToArmorList(armorsList []*Armor) armors.Armors {
	armorDomainList := make(armors.Armors, len(armorsList))
	for i, a := range armorsList {
		armorDomainList[i] = *ToArmor(a)
	}
	return armorDomainList
}
