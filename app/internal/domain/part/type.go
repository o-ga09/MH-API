package part

type Parts []Part

type PartId struct{ Value string }
type PartName struct{ value string }
type PartDescription struct{ value string }

func (f *Part) GetMonsterID() string   { return f.monsterId.Value }
func (f *Part) GetID() string          { return f.partId.Value }
func (f *Part) GetName() string        { return f.name.value }
func (f *Part) GetDescription() string { return f.description.value }
