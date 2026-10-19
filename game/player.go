package game

type Player struct {
	Pos           BlockPos
	Name          string
	Direction     Direction
	Texture       Appearance
	InteractBox   BlockPos
	DefaultDamage float64
	Inventory     InventorySlice
	SelectedBlock int
}

func (p *Player) Attack(m model) {
	block, ok := m.Blocks[p.InteractBox]
	if !ok {
		return
	}

	block.ApplyDamage(p, &m, p.DefaultDamage)
}

func (p *Player) AddToInventory(block BlockName, qty int) {
	if p.Inventory == nil {
		p.Inventory = make(InventorySlice, 0)
	}

	found := -1
	for index, value := range p.Inventory {
		if value.Name == block {
			p.Inventory[index].Qty = p.Inventory[index].Qty + qty
			found
		}
	}

}

func (p *Player) RemoveFromInventory(block BlockName, qty int) (Success bool) {
	if p.Inventory == nil {
		p.Inventory = make(InventorySlice, 0)
	}

	for index, value := range p.Inventory {
		if value.Name == block {
			if p.Inventory[index].Qty < qty {
				return false
			}
			p.Inventory[index].Qty = p.Inventory[index].Qty - qty
			return true
		}
	}
	return false
}

func (p *Player) SelectNext() {

	//iter.Pull(keys[])
}

func (m *model) PlaceBlock() {

}
