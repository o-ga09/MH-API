package weapons

type Weapons []Weapon

type WeaponId struct{ value string }
type WeaponName struct{ value string }
type WeaponImageUrl struct{ value string }
type WeaponRarity struct{ value string }
type WeaponAttack struct{ value string }
type WeaponElementAttack struct{ value string }
type WeaponShapness struct{ value string }
type WeaponCritical struct{ value string }
type WeaponDescription struct{ value string }

func (f *Weapon) GetID() string            { return f.weaponId.value }
func (f *Weapon) GetName() string          { return f.name.value }
func (f *Weapon) GetURL() string           { return f.imageUrl.value }
func (f *Weapon) GetRERATY() string        { return f.rare.value }
func (f *Weapon) GetAttack() string        { return f.attack.value }
func (f *Weapon) GetElementAttack() string { return f.elemantAttaxk.value }
func (f *Weapon) GetShapness() string      { return f.shapness.value }
func (f *Weapon) GetCritical() string      { return f.critical.value }
func (f *Weapon) GetDescription() string   { return f.description.value }
