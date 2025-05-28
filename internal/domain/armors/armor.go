package armors

type Armor struct {
	armorId              ArmorId
	name                 ArmorName
	slot                 ArmorSlot
	defense              ArmorDefense
	fireResistance       ArmorFireResistance
	waterResistance      ArmorWaterResistance
	lightningResistance  ArmorLightningResistance
	iceResistance        ArmorIceResistance
	dragonResistance     ArmorDragonResistance
	skills               []Skill
	requiredItems        []RequiredItem
}

func newArmor(
	armorId ArmorId,
	name ArmorName,
	slot ArmorSlot,
	defense ArmorDefense,
	fireResistance ArmorFireResistance,
	waterResistance ArmorWaterResistance,
	lightningResistance ArmorLightningResistance,
	iceResistance ArmorIceResistance,
	dragonResistance ArmorDragonResistance,
	skills []Skill,
	requiredItems []RequiredItem,
) *Armor {
	return &Armor{
		armorId:             armorId,
		name:                name,
		slot:                slot,
		defense:             defense,
		fireResistance:      fireResistance,
		waterResistance:     waterResistance,
		lightningResistance: lightningResistance,
		iceResistance:       iceResistance,
		dragonResistance:    dragonResistance,
		skills:              skills,
		requiredItems:       requiredItems,
	}
}

func NewArmor(
	id string,
	name string,
	slot string,
	defense int,
	fireResistance int,
	waterResistance int,
	lightningResistance int,
	iceResistance int,
	dragonResistance int,
	skills []Skill,
	requiredItems []RequiredItem,
) *Armor {
	return newArmor(
		ArmorId{value: id},
		ArmorName{value: name},
		ArmorSlot{value: slot},
		ArmorDefense{value: defense},
		ArmorFireResistance{value: fireResistance},
		ArmorWaterResistance{value: waterResistance},
		ArmorLightningResistance{value: lightningResistance},
		ArmorIceResistance{value: iceResistance},
		ArmorDragonResistance{value: dragonResistance},
		skills,
		requiredItems,
	)
}