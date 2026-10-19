package game

type Block struct {
	Name         BlockName `json:"crop_name"`
	Amount       int       `json:"amount"`
	Value        float64   `json:"crop_value"`
	IsCrop       bool
	IsSolid      bool
	Invincible   bool
	ReturnAmount int
	HP           float64
	Pos          BlockPos
	Appearance   Appearance
}

func (b *Block) ApplyDamage(This any, m *model, Damage float64) (Destroyed bool) {

	if b.Invincible {
		return false
	}

	if Damage >= b.HP || b.HP <= 0 {
		b.HP = 0
		grass := BlockMap[Grass]
		grass.Pos = b.Pos

		m.Blocks[b.Pos] = &grass

		if v, ok := This.(Inventory); ok {
			v.AddToInventory(b.Name, b.ReturnAmount)
		}

		return true
	}
	b.HP -= Damage

	return false
}
