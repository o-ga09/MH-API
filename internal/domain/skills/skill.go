package skills

type Skills []Skill

type Skill struct {
	id          SkillId
	name        SkillName
	description SkillDesc
	levels      []SkillLevelDetail
}

type SkillLevelDetail struct {
	levelId     SkillLevelId
	skillId     SkillId
	level       SkillLevel
	description SkillLevelDesc
}

func newSkill(id SkillId, name SkillName, description SkillDesc, levels []SkillLevelDetail) Skill {
	return Skill{
		id:          id,
		name:        name,
		description: description,
		levels:      levels,
	}
}

func NewSkill(id string, name string, description string, levels []SkillLevelDetail) Skill {
	return newSkill(
		SkillId{Value: id},
		SkillName{Value: name},
		SkillDesc{Value: description},
		levels,
	)
}

func NewSkillLevelDetail(levelId string, skillId string, level int, description string) SkillLevelDetail {
	return SkillLevelDetail{
		levelId:     SkillLevelId{Value: levelId},
		skillId:     SkillId{Value: skillId},
		level:       SkillLevel{Value: level},
		description: SkillLevelDesc{Value: description},
	}
}

func (s *Skill) GetId() string {
	return s.id.Value
}

func (s *Skill) GetName() string {
	return s.name.Value
}

func (s *Skill) GetDescription() string {
	return s.description.Value
}

func (s *Skill) GetLevels() []SkillLevelDetail {
	return s.levels
}

func (sld *SkillLevelDetail) GetLevelId() string {
	return sld.levelId.Value
}

func (sld *SkillLevelDetail) GetSkillId() string {
	return sld.skillId.Value
}

func (sld *SkillLevelDetail) GetLevel() int {
	return sld.level.Value
}

func (sld *SkillLevelDetail) GetDescription() string {
	return sld.description.Value
}
