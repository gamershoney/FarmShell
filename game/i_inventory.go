package game

type InventoryBlock struct {
	Name BlockName
	Qty  int
}

type InventorySlice []InventoryBlock

type Inventory interface {
	AddToInventory(block BlockName, qty int)
	RemoveFromInventory(block BlockName, qty int) (Success bool)
}
