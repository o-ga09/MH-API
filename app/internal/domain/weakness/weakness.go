package weakness

import (
	"mh-api/app/internal/domain/monsters"
	"mh-api/app/internal/domain/part"
)

type Weakness struct {
	monsterId         monsters.MonsterId
	partId            part.PartId
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

func newWeakness(
	monsterId monsters.MonsterId,
	pointId part.PartId,
	fire FireRV,
	water WaterRV,
	lightning LightningRV,
	ice IceRV,
	dragon DragonRV,
	slashing SlashingRV,
	blow BlowRV,
	bullet BulletRV,
	firstWeakAttack FirstWeakAttack,
	secondWeakAttack SecondWeakAttack,
	firstWeakElement FirstWeakElement,
	secondWeakElement SecondWeakElement,
) *Weakness {
	return &Weakness{
		monsterId:         monsterId,
		partId:            pointId,
		fire:              fire,
		water:             water,
		lightning:         lightning,
		ice:               ice,
		dragon:            dragon,
		slashing:          slashing,
		blow:              blow,
		bullet:            bullet,
		firstWeakAttack:   firstWeakAttack,
		secondWeakAttack:  secondWeakAttack,
		firstWeakElement:  firstWeakElement,
		secondWeakElement: secondWeakElement,
	}
}

func NewWeakness(
	monsterId string,
	partId string,
	fireRV string,
	waterRV string,
	lightningRV string,
	iceRV string,
	dragonRV string,
	slashingRV string,
	blowRV string,
	bulletRV string,
	firstWeakAttack string,
	secondWeakAttack string,
	firstWeakElement string,
	secondWeakElement string,
) *Weakness {
	return newWeakness(
		monsters.MonsterId{Value: monsterId},
		part.PartId{},
		FireRV{value: fireRV},
		WaterRV{value: waterRV},
		LightningRV{value: lightningRV},
		IceRV{value: iceRV},
		DragonRV{value: dragonRV},
		SlashingRV{value: slashingRV},
		BlowRV{value: blowRV},
		BulletRV{value: bulletRV},
		FirstWeakAttack{value: firstWeakAttack},
		SecondWeakAttack{value: secondWeakAttack},
		FirstWeakElement{value: firstWeakElement},
		SecondWeakElement{value: secondWeakElement},
	)
}
