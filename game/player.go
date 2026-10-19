package game

type Player struct {
	Pos           BlockPos
	Name          string
	Direction     Direction
	Texture       Appearance
	InteractBox   BlockPos
	DefaultDamage float64
	Inventory     InventoryMap
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
		p.Inventory = make(InventoryMap)
	}

	p.Inventory[block] = p.Inventory[block] + qty

}

func (p *Player) RemoveFromInventory(qty int) {}
