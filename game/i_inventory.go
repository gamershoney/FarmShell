package game

type InventoryMap map[BlockName]int

type Inventory interface {
	AddToInventory(block BlockName, qty int)
	RemoveFromInventory(qty int)
}
