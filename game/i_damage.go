package game

type Damage interface {
	ApplyDamage(This any, m *model, Damage float64) bool
}
