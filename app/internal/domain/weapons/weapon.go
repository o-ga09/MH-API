package weapons

type Weapon struct {
	weaponId      WeaponId
	name          WeaponName
	imageUrl      WeaponImageUrl
	rare          WeaponRarity
	attack        WeaponAttack
	elemantAttaxk WeaponElementAttack
	shapness      WeaponShapness
	critical      WeaponCritical
	description   WeaponDescription
}

func newWeapon(
	weaponId WeaponId,
	name WeaponName,
	imageUrl WeaponImageUrl,
	rerarity WeaponRarity,
	attack WeaponAttack,
	elementattack WeaponElementAttack,
	shapness WeaponShapness,
	critical WeaponCritical,
	description WeaponDescription,
) *Weapon {
	return &Weapon{weaponId, name, imageUrl, rerarity, attack, elementattack, shapness, critical, description}
}

func NewWeapon(weaponId string, name string, imageUrl string, rerarity string, attack string, elementAttac string, shapness string, critical string, description string) *Weapon {
	return newWeapon(
		WeaponId{value: weaponId},
		WeaponName{value: name},
		WeaponImageUrl{value: imageUrl},
		WeaponRarity{value: rerarity},
		WeaponAttack{value: attack},
		WeaponElementAttack{value: elementAttac},
		WeaponShapness{value: shapness},
		WeaponCritical{value: critical},
		WeaponDescription{value: description},
	)
}
