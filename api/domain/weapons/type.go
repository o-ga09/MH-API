package weapons

type WeaponId struct{ value string }
type WeaponName struct{ value string }
type WeaponImageUrl struct{ value string }
type WeaponRarity struct{ value string }
type WeaponAttack struct{ value string }
type WeaponElementAttack struct{ value string }
type WeaponShapness struct{ value string }
type WeaponCritical struct{ value string }
type WeaponDescription struct{ value string }

func (f *WeaponId) GetID() string                       { return f.value }
func (f *WeaponName) GetName() string                   { return f.value }
func (f *WeaponImageUrl) GetURL() string                { return f.value }
func (f *WeaponRarity) GetRERATY() string               { return f.value }
func (f *WeaponAttack) GetAttack() string               { return f.value }
func (f *WeaponElementAttack) GetElementAttack() string { return f.value }
func (f *WeaponShapness) GetShapness() string           { return f.value }
func (f *WeaponCritical) GetCritical() string           { return f.value }
func (f *WeaponDescription) GetDescription() string     { return f.value }
