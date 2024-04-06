package monsters

type Monsters []Monster

type Monster struct {
	id   MonsterId
	name MonsterName
	desc MonsterDesc
}

func newMonster(id MonsterId, name MonsterName, desc MonsterDesc) Monster {
	return Monster{
		id:   id,
		name: name,
		desc: desc,
	}
}

func NewMonster(id string, name string, desc string) Monster {
	return newMonster(
		MonsterId{Value: id},
		MonsterName{Value: name},
		MonsterDesc{Value: desc},
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
