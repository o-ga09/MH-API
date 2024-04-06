package Tribes

type Tribe struct {
	TribeId     TribeId
	name_ja     TribeNameJA
	name_en     TribeNameEN
	description TribeDescription
}

func newTribe(TribeId TribeId, name_ja TribeNameJA, name_en TribeNameEN, description TribeDescription) *Tribe {
	return &Tribe{
		TribeId:     TribeId,
		name_ja:     name_ja,
		name_en:     name_en,
		description: description,
	}
}

func NewFiled(tribeId string, name_ja string, name_en string, description string) *Tribe {
	return newTribe(
		TribeId{value: tribeId},
		TribeNameJA{value: name_ja},
		TribeNameEN{value: name_en},
		TribeDescription{value: description},
	)
}
