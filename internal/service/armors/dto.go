package armors

import (
	"mh-api/internal/domain/armors"
)

type SkillData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RequiredItemData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResistanceData struct {
	Fire      int `json:"fire"`
	Water     int `json:"water"`
	Lightning int `json:"lightning"`
	Ice       int `json:"ice"`
	Dragon    int `json:"dragon"`
}

type ArmorData struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Skill     []SkillData        `json:"skill"`
	Slot      string             `json:"slot"`
	Defense   int                `json:"defense"`
	Resistance ResistanceData    `json:"resistance"`
	Required  []RequiredItemData `json:"required"`
}

type ListArmorsResponse struct {
	Armors []ArmorData `json:"armors"`
}

func ToArmorData(armor *armors.Armor) ArmorData {
	skills := make([]SkillData, len(armor.GetSkills()))
	for i, skill := range armor.GetSkills() {
		skills[i] = SkillData{
			ID:   skill.GetID(),
			Name: skill.GetName(),
		}
	}

	requiredItems := make([]RequiredItemData, len(armor.GetRequiredItems()))
	for i, item := range armor.GetRequiredItems() {
		requiredItems[i] = RequiredItemData{
			ID:   item.GetID(),
			Name: item.GetName(),
		}
	}

	return ArmorData{
		ID:   armor.GetID(),
		Name: armor.GetName(),
		Skill: skills,
		Slot: armor.GetSlot(),
		Defense: armor.GetDefense(),
		Resistance: ResistanceData{
			Fire:      armor.GetFireResistance(),
			Water:     armor.GetWaterResistance(),
			Lightning: armor.GetLightningResistance(),
			Ice:       armor.GetIceResistance(),
			Dragon:    armor.GetDragonResistance(),
		},
		Required: requiredItems,
	}
}

func ToArmorDataList(domainArmors armors.Armors) []ArmorData {
	armorDataList := make([]ArmorData, len(domainArmors))
	for i, da := range domainArmors {
		armorDataList[i] = ToArmorData(&da)
	}
	return armorDataList
}