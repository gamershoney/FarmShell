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

	for index, value := range p.Inventory {
		if value.Name == block {
			p.Inventory[index].Qty = p.Inventory[index].Qty + qty
			return
		}
	}

	item := InventoryBlock{
		Name: block,
		Qty:  qty,
	}

	p.Inventory = append(p.Inventory, item)
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

	next := p.SelectedBlock + 1

	if len(p.Inventory)-1 > next {
		p.SelectedBlock = next
		return
	}
	p.SelectedBlock = 0

}

func (p *Player) SelectPrev() {
	prev := p.SelectedBlock - 1

	if prev < 0 {
		p.SelectedBlock = len(p.Inventory) - 1
		return
	}

	p.SelectedBlock = prev
}

func (m *model) PlaceBlock() {
	block := m.Player.Inventory[m.Player.SelectedBlock]
	newBlock := BlockMap[block.Name]
	m.Blocks[m.Player.InteractBox] = &newBlock
}
