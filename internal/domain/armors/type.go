package armors

type Armors []Armor

type ArmorId struct{ value string }

func (id ArmorId) Value() string { return id.value }

type ArmorName struct{ value string }

func (name ArmorName) Value() string { return name.value }

type ArmorSlot struct{ value string }

func (slot ArmorSlot) Value() string { return slot.value }

type ArmorDefense struct{ value int }

func (defense ArmorDefense) Value() int { return defense.value }

type ArmorFireResistance struct{ value int }

func (resistance ArmorFireResistance) Value() int { return resistance.value }

type ArmorWaterResistance struct{ value int }

func (resistance ArmorWaterResistance) Value() int { return resistance.value }

type ArmorLightningResistance struct{ value int }

func (resistance ArmorLightningResistance) Value() int { return resistance.value }

type ArmorIceResistance struct{ value int }

func (resistance ArmorIceResistance) Value() int { return resistance.value }

type ArmorDragonResistance struct{ value int }

func (resistance ArmorDragonResistance) Value() int { return resistance.value }

type SkillId struct{ value string }

func (id SkillId) Value() string { return id.value }

type SkillName struct{ value string }

func (name SkillName) Value() string { return name.value }

type RequiredItemId struct{ value string }

func (id RequiredItemId) Value() string { return id.value }

type RequiredItemName struct{ value string }

func (name RequiredItemName) Value() string { return name.value }

type Skill struct {
	skillId SkillId
	name    SkillName
}

func NewSkill(id string, name string) *Skill {
	return &Skill{
		skillId: SkillId{value: id},
		name:    SkillName{value: name},
	}
}

func (s *Skill) GetID() string   { return s.skillId.value }
func (s *Skill) GetName() string { return s.name.value }

type RequiredItem struct {
	itemId RequiredItemId
	name   RequiredItemName
}

func NewRequiredItem(id string, name string) *RequiredItem {
	return &RequiredItem{
		itemId: RequiredItemId{value: id},
		name:   RequiredItemName{value: name},
	}
}

func (r *RequiredItem) GetID() string   { return r.itemId.value }
func (r *RequiredItem) GetName() string { return r.name.value }

func (a *Armor) GetID() string                { return a.armorId.value }
func (a *Armor) GetName() string              { return a.name.value }
func (a *Armor) GetSlot() string              { return a.slot.value }
func (a *Armor) GetDefense() int              { return a.defense.value }
func (a *Armor) GetFireResistance() int       { return a.fireResistance.value }
func (a *Armor) GetWaterResistance() int      { return a.waterResistance.value }
func (a *Armor) GetLightningResistance() int  { return a.lightningResistance.value }
func (a *Armor) GetIceResistance() int        { return a.iceResistance.value }
func (a *Armor) GetDragonResistance() int     { return a.dragonResistance.value }
func (a *Armor) GetSkills() []Skill           { return a.skills }
func (a *Armor) GetRequiredItems() []RequiredItem { return a.requiredItems }