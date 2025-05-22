package weapons

type Weapon struct {
	MonsterId     WeaponId            `json:"monster_id"`
	Name          WeaponName          `json:"name"`
	ImageUrl      WeaponImageUrl      `json:"image_url"`
	Rare          WeaponRarity        `json:"rare"`
	Attack        WeaponAttack        `json:"attack"`
	ElementAttack WeaponElementAttack `json:"element_attack"`
	Sharpness     WeaponShapness      `json:"sharpness"`
	Critical      WeaponCritical      `json:"critical"`
	Description   WeaponDescription   `json:"description"`
}

func newWeapon(
	monsterId WeaponId,
	name WeaponName,
	imageUrl WeaponImageUrl,
	rarity WeaponRarity,
	attack WeaponAttack,
	elementAttack WeaponElementAttack,
	sharpness WeaponShapness,
	critical WeaponCritical,
	description WeaponDescription,
) *Weapon {
	return &Weapon{monsterId, name, imageUrl, rarity, attack, elementAttack, sharpness, critical, description}
}

func NewWeapon(monsterId string, name string, imageUrl string, rarity string, attack string, elementAttack string, sharpness string, critical string, description string) *Weapon {
	return newWeapon(
		WeaponId{value: monsterId},
		WeaponName{value: name},
		WeaponImageUrl{value: imageUrl},
		WeaponRarity{value: rarity},
		WeaponAttack{value: attack},
		WeaponElementAttack{value: elementAttack},
		WeaponShapness{value: sharpness},
		WeaponCritical{value: critical},
		WeaponDescription{value: description},
	)
}
