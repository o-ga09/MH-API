package part

import "mh-api/app/internal/domain/monsters"

type Part struct {
	partId      PartId
	monsterId   monsters.MonsterId
	name        PartName
	description PartDescription
}

func newPart(PartId PartId, monsterId monsters.MonsterId, name PartName, description PartDescription) *Part {
	return &Part{PartId, monsterId, name, description}
}

func NewPart(partId string, monsterId string, name string, imageUrl string) *Part {
	return newPart(
		PartId{Value: partId},
		monsters.MonsterId{Value: monsterId},
		PartName{value: name},
		PartDescription{value: imageUrl},
	)
}
