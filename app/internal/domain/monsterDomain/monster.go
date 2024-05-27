package monsterdomain

func NewMonster(
	id string,
	name string,
	desc string,
) *Monster {
	return newMonster(
		MonsterId{value: id},
		MonsterName{value: name},
		MonsterDesc{value: desc},
	)
}

func newMonster(id MonsterId, name MonsterName, desc MonsterDesc) *Monster {
	monster := &Monster{
		id:   id,
		name: name,
		desc: desc,
	}
	return monster
}

type Monster struct {
	id       MonsterId
	name     MonsterName
	desc     MonsterDesc
	field    []Field
	tribe    Tribe
	product  []Product
	ranking  []Ranking
	weakness []Weakness
}

func NewField(
	fieldId string,
	monsterId string,
	name string,
	imageUrl string,
) *Field {
	return &Field{
		fieldId:   FieldId{value: fieldId},
		monsterId: MonsterId{value: monsterId},
		name:      FieldName{value: name},
		imageUrl:  FieldImageUrl{value: imageUrl},
	}
}

func NewProduct(
	productId string,
	name string,
	publishYear string,
	totalSales string,
) *Product {
	return &Product{
		productId:   ProductId{value: productId},
		name:        ProductName{value: name},
		publishYear: ProductPublishYear{value: publishYear},
		totalSales:  ProductTotalSales{value: totalSales},
	}
}

func NewRanking(
	monsterId string,
	ranking string,
	voteYear string,
) *Ranking {
	return &Ranking{
		monsterId: MonsterId{value: monsterId},
		ranking:   Rank{value: ranking},
		voteYear:  VoteYear{value: voteYear},
	}
}

func NewTribe(
	tribeId string,
	monsterId string,
	nameJA string,
	nameEN string,
	description string,
) *Tribe {
	return &Tribe{
		tribeId:     TribeId{value: tribeId},
		monsterId:   MonsterId{value: monsterId},
		name_ja:     TribeNameJA{value: nameJA},
		name_en:     TribeNameEN{value: nameEN},
		description: TribeDescription{value: description},
	}
}

func NewWeakness(
	monsterId string,
	partId string,
	fire string,
	water string,
	lightning string,
	ice string,
	dragon string,
	slashing string,
	blow string,
	bullet string,
	firstWeakAttack string,
	secondWeakAttack string,
	firstWeakElement string,
	secondWeakElement string,
) *Weakness {
	return &Weakness{
		monsterId:         MonsterId{value: monsterId},
		partId:            PartId{Value: partId},
		fire:              FireRV{value: fire},
		water:             WaterRV{value: water},
		lightning:         LightningRV{value: lightning},
		ice:               IceRV{value: ice},
		dragon:            DragonRV{value: dragon},
		slashing:          SlashingRV{value: slashing},
		blow:              BlowRV{value: blow},
		bullet:            BulletRV{value: bullet},
		firstWeakAttack:   FirstWeakAttack{value: firstWeakAttack},
		secondWeakAttack:  SecondWeakAttack{value: secondWeakAttack},
		firstWeakElement:  FirstWeakElement{value: firstWeakElement},
		secondWeakElement: SecondWeakElement{value: secondWeakElement},
	}
}

type Field struct {
	fieldId   FieldId
	monsterId MonsterId
	name      FieldName
	imageUrl  FieldImageUrl
}

type Product struct {
	productId   ProductId
	name        ProductName
	publishYear ProductPublishYear
	totalSales  ProductTotalSales
}

type Ranking struct {
	monsterId MonsterId
	ranking   Rank
	voteYear  VoteYear
}

type Tribe struct {
	tribeId     TribeId
	monsterId   MonsterId
	name_ja     TribeNameJA
	name_en     TribeNameEN
	description TribeDescription
}

type Weakness struct {
	monsterId         MonsterId
	partId            PartId
	fire              FireRV
	water             WaterRV
	lightning         LightningRV
	ice               IceRV
	dragon            DragonRV
	slashing          SlashingRV
	blow              BlowRV
	bullet            BulletRV
	firstWeakAttack   FirstWeakAttack
	secondWeakAttack  SecondWeakAttack
	firstWeakElement  FirstWeakElement
	secondWeakElement SecondWeakElement
}

func NewMonsterId(id string) MonsterId { return MonsterId{value: id} }
func (m *MonsterId) GetID() string     { return m.value }

func (m *Monster) GetID() string   { return m.id.value }
func (m *Monster) GetName() string { return m.name.value }
func (m *Monster) GetDesc() string { return m.desc.value }
func (m *Monster) GetField() []Field {
	return m.field
}
func (m *Monster) GetTribe() Tribe {
	return m.tribe
}
func (m *Monster) GetProduct() []Product {
	return m.product
}
func (m *Monster) GetRanking() []Ranking {
	return m.ranking
}
func (m *Monster) GetWeakness() []Weakness {
	return m.weakness
}

type MonsterId struct{ value string }   //モンスターID
type MonsterName struct{ value string } //モンスターの名前
type MonsterDesc struct{ value string } //モンスターの説明文
type FieldId struct{ value string }
type FieldName struct{ value string }
type FieldImageUrl struct{ value string }
type ProductId struct{ value string }
type ProductName struct{ value string }
type ProductPublishYear struct{ value string }
type ProductTotalSales struct{ value string }
type Rank struct{ value string }
type VoteYear struct{ value string }
type TribeId struct{ value string }
type TribeNameJA struct{ value string }
type TribeNameEN struct{ value string }
type TribeDescription struct{ value string }
type FireRV struct{ value string }
type WaterRV struct{ value string }
type LightningRV struct{ value string }
type IceRV struct{ value string }
type DragonRV struct{ value string }
type SlashingRV struct{ value string }
type BlowRV struct{ value string }
type BulletRV struct{ value string }
type FirstWeakAttack struct{ value string }
type SecondWeakAttack struct{ value string }
type FirstWeakElement struct{ value string }
type SecondWeakElement struct{ value string }
type PartId struct{ Value string }

func (f *Field) GetID() string                   { return f.fieldId.value }
func (f *Field) GetMonsterID() string            { return f.monsterId.value }
func (f *Field) GetName() string                 { return f.name.value }
func (f *Field) GetURL() string                  { return f.imageUrl.value }
func (f *Product) GetID() string                 { return f.productId.value }
func (f *Product) GetName() string               { return f.name.value }
func (f *Product) GetYear() string               { return f.publishYear.value }
func (f *Product) GetSales() string              { return f.totalSales.value }
func (r *Ranking) GetID() string                 { return r.monsterId.value }
func (r *Ranking) GetRank() string               { return r.ranking.value }
func (r *Ranking) GetVoteYear() string           { return r.voteYear.value }
func (f *Tribe) GetID() string                   { return f.tribeId.value }
func (f *Tribe) GetMonsterID() string            { return f.monsterId.value }
func (f *Tribe) GetNameJA() string               { return f.name_ja.value }
func (f *Tribe) GetNameEN() string               { return f.name_en.value }
func (f *Tribe) GetDescription() string          { return f.description.value }
func (f *Weakness) GetMonsterID() string         { return f.monsterId.value }
func (f *Weakness) GetPartID() string            { return f.partId.Value }
func (f *Weakness) GetFire() string              { return f.fire.value }
func (f *Weakness) GetWater() string             { return f.water.value }
func (f *Weakness) GetLightning() string         { return f.lightning.value }
func (f *Weakness) GetIce() string               { return f.ice.value }
func (f *Weakness) GetDragon() string            { return f.dragon.value }
func (f *Weakness) GetSlashing() string          { return f.slashing.value }
func (f *Weakness) GetBlow() string              { return f.blow.value }
func (f *Weakness) GetBullet() string            { return f.bullet.value }
func (f *Weakness) GetFirstWeakAttack() string   { return f.firstWeakAttack.value }
func (f *Weakness) GetSecondWeakAttack() string  { return f.secondWeakAttack.value }
func (f *Weakness) GetFirstWeakElement() string  { return f.firstWeakElement.value }
func (f *Weakness) GetSecondWeakElement() string { return f.secondWeakElement.value }
