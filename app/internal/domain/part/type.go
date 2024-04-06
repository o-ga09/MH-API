package part

type PartId struct{ Value string }
type PartName struct{ value string }
type PartDescription struct{ value string }

func (f *PartId) GetID() string           { return f.Value }
func (f *PartName) GetName() string       { return f.value }
func (f *PartDescription) GetURL() string { return f.value }
