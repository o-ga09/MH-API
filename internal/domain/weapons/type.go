package weapons

type Weapons []Weapon

type WeaponId struct{ value string }
func (id WeaponId) Value() string { return id.value }

type WeaponName struct{ value string }
func (name WeaponName) Value() string { return name.value }

type WeaponImageUrl struct{ value string }
func (img WeaponImageUrl) Value() string { return img.value }

type WeaponRarity struct{ value string }
func (r WeaponRarity) Value() string { return r.value }

type WeaponAttack struct{ value string }
func (a WeaponAttack) Value() string { return a.value }

type WeaponElementAttack struct{ value string }
func (ea WeaponElementAttack) Value() string { return ea.value }

type WeaponShapness struct{ value string }
func (s WeaponShapness) Value() string { return s.value }

type WeaponCritical struct{ value string }
func (c WeaponCritical) Value() string { return c.value }

type WeaponDescription struct{ value string }
func (d WeaponDescription) Value() string { return d.value }

func (f *Weapon) GetID() string            { return f.weaponId.value }
func (f *Weapon) GetName() string          { return f.name.value }
func (f *Weapon) GetURL() string           { return f.imageUrl.value }
func (f *Weapon) GetRERATY() string        { return f.rare.value }
func (f *Weapon) GetAttack() string        { return f.attack.value }
func (f *Weapon) GetElementAttack() string { return f.elemantAttaxk.value }
func (f *Weapon) GetShapness() string      { return f.shapness.value }
func (f *Weapon) GetCritical() string      { return f.critical.value }
func (f *Weapon) GetDescription() string   { return f.description.value }
