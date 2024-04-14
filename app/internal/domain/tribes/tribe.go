package Tribes

import "mh-api/app/internal/domain/monsters"

type Tribe struct {
	tribeId     TribeId
	monsterId   monsters.MonsterId
	name_ja     TribeNameJA
	name_en     TribeNameEN
	description TribeDescription
}

func newTribe(TribeId TribeId, monsterId monsters.MonsterId, name_ja TribeNameJA, name_en TribeNameEN, description TribeDescription) *Tribe {
	return &Tribe{
		tribeId:     TribeId,
		monsterId:   monsterId,
		name_ja:     name_ja,
		name_en:     name_en,
		description: description,
	}
}

func NewTribe(tribeId string, monsterId string, name_ja string, name_en string, description string) *Tribe {
	return newTribe(
		TribeId{value: tribeId},
		monsters.MonsterId{Value: monsterId},
		TribeNameJA{value: name_ja},
		TribeNameEN{value: name_en},
		TribeDescription{value: description},
	)
}
