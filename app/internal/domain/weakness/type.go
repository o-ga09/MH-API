package weakness

type Weaknesses []Weakness

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

func (f *Weakness) GetMonsterID() string         { return f.monsterId.Value }
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
