package game

type Damage interface {
	ApplyDamage(Damage float64) bool
}
