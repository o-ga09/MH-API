package Tribes

type Tribes []Tribe

type TribeId struct{ value string }
type TribeNameJA struct{ value string }
type TribeNameEN struct{ value string }
type TribeDescription struct{ value string }

func (f *Tribe) GetID() string          { return f.tribeId.value }
func (f *Tribe) GetMonsterID() string   { return f.monsterId.Value }
func (f *Tribe) GetNameJA() string      { return f.name_ja.value }
func (f *Tribe) GetNameEN() string      { return f.name_en.value }
func (f *Tribe) GetDescription() string { return f.description.value }
