package part

import "mh-api/app/internal/domain/monsters"

type Part struct {
	partId      PartId
	monsterId   monsters.MonsterId
	description PartDescription
}

func newPart(PartId PartId, monsterId monsters.MonsterId, description PartDescription) *Part {
	return &Part{PartId, monsterId, description}
}

func NewPart(partId string, monsterId string, imageUrl string) *Part {
	return newPart(
		PartId{Value: partId},
		monsters.MonsterId{Value: monsterId},
		PartDescription{value: imageUrl},
	)
}
