package fields

type FieldId struct{ value string }
type FieldName struct{ value string }
type FieldImageUrl struct{ value string }

func (f *FieldId) GetID() string        { return f.value }
func (f *FieldName) GetName() string    { return f.value }
func (f *FieldImageUrl) GetURL() string { return f.value }
