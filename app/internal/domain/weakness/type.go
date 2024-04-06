package weakness

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

func (f *FireRV) GetFire() string                         { return f.value }
func (f *WaterRV) GetWater() string                       { return f.value }
func (f *LightningRV) GetLightning() string               { return f.value }
func (f *IceRV) GetIce() string                           { return f.value }
func (f *DragonRV) GetDragon() string                     { return f.value }
func (f *SlashingRV) GetSlashing() string                 { return f.value }
func (f *BlowRV) GetBlow() string                         { return f.value }
func (f *BulletRV) GetBullet() string                     { return f.value }
func (f *FirstWeakAttack) GetFirstWeakAttack() string     { return f.value }
func (f *SecondWeakAttack) GetSecondWeakAttack() string   { return f.value }
func (f *FirstWeakElement) GetFirstWeakElement() string   { return f.value }
func (f *SecondWeakElement) GetSecondWeakElement() string { return f.value }
