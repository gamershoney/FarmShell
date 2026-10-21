package game

type Block struct {
	Name         BlockName `json:"crop_name"`
	Amount       int       `json:"amount"`
	Value        float64   `json:"crop_value"`
	IsCrop       bool
	IsSolid      bool
	Invincible   bool
	ReturnAmount int
	CanTick      bool
	Tick         func(*Block, *model)
	State        uint64
	HP           float64
	Pos          BlockPos
	ReturnItem   BlockName
	Appearance   Appearance
	Destroy      func(*Block) (BlockName, int)
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
			var name BlockName
			var qty int
			if b.Destroy != nil {
				name, qty = b.Destroy(b)
				v.AddToInventory(name, qty)
			} else {

				v.AddToInventory(b.Name, b.ReturnAmount)
			}
		}

		return true
	}
	b.HP -= Damage

	return false
}
