package monsters

type Monsters []Monster

type Monster struct {
	id          MonsterId
	name        MonsterName
	desc        MonsterDesc
	anotherName MonsterName
	nameEn      MonsterName
	element     MonsterElement
}

func newMonster(id MonsterId, name MonsterName, desc MonsterDesc, element MonsterElement) Monster {
	return Monster{
		id:      id,
		name:    name,
		desc:    desc,
		element: element,
	}
}

func NewMonster(id string, name string, desc string, element string) Monster {
	return newMonster(
		MonsterId{Value: id},
		MonsterName{Value: name},
		MonsterDesc{Value: desc},
		MonsterElement{Value: element},
	)
}

func (m *Monster) GetId() string {
	return m.id.Value
}

func (m *Monster) GetName() string {
	return m.name.Value
}

func (m *Monster) GetDesc() string {
	return m.desc.Value
}

func (m *Monster) GetAnotherName() string {
	return m.anotherName.Value
}

func (m *Monster) GetNameEn() string {
	return m.nameEn.Value
}

func (m *Monster) GetElement() string {
	return m.element.Value
}
