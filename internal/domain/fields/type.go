package fields

type Fields []Field

type FieldId struct{ value string }
type FieldName struct{ value string }
type FieldImageUrl struct{ value string }

func (f *Field) GetID() string        { return f.fieldId.value }
func (f *Field) GetMonsterID() string { return f.monsterId.Value }
func (f *Field) GetName() string      { return f.name.value }
func (f *Field) GetURL() string       { return f.imageUrl.value }
