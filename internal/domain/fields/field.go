package fields

import "mh-api/internal/domain/monsters"

type Field struct {
	fieldId   FieldId
	monsterId monsters.MonsterId
	name      FieldName
	imageUrl  FieldImageUrl
}

func newField(fieldId FieldId, monsterId monsters.MonsterId, name FieldName, imageUrl FieldImageUrl) *Field {
	return &Field{fieldId, monsterId, name, imageUrl}
}

func NewField(fieldId string, monsterId string, name string, imageUrl string) *Field {
	return newField(
		FieldId{value: fieldId},
		monsters.MonsterId{Value: monsterId},
		FieldName{value: name},
		FieldImageUrl{value: imageUrl},
	)
}
