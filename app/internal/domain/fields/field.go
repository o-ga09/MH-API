package fields

type Field struct {
	fieldId  FieldId
	name     FieldName
	imageUrl FieldImageUrl
}

func newField(fieldId FieldId, name FieldName, imageUrl FieldImageUrl) *Field {
	return &Field{fieldId, name, imageUrl}
}

func NewFiled(fieldId string, name string, imageUrl string) *Field {
	return newField(
		FieldId{value: fieldId},
		FieldName{value: name},
		FieldImageUrl{value: imageUrl},
	)
}
