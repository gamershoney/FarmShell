package game

type BlockName string

const Empty BlockName = "Empty"
const Wheat BlockName = "Wheat"
const Grass BlockName = "Grass"
const PlayerBlock BlockName = "Player"
const Wood BlockName = "Wood"
const Leaf BlockName = "Leaf"

var BlockMap = map[BlockName]Block{
	Empty: {
		Name:         Empty,
		Amount:       0,
		Value:        0,
		IsCrop:       false,
		IsSolid:      false,
		HP:           1000,
		Invincible:   true,
		Appearance:   textures[Empty],
		ReturnAmount: 0,
	},
	Wheat: {
		Name:         Wheat,
		Amount:       1,
		Value:        1,
		IsCrop:       true,
		IsSolid:      true,
		HP:           1,
		Appearance:   textures[Wheat],
		ReturnAmount: 1,
	},

	Grass: {
		Name:       Grass,
		Amount:     1,
		Value:      0,
		IsCrop:     false,
		IsSolid:    false,
		HP:         100000,
		Invincible: true,
		Appearance: textures[Grass],
	},

	Wood: {
		Name:         Wood,
		Amount:       1,
		Value:        1,
		IsCrop:       true,
		IsSolid:      true,
		HP:           3,
		Invincible:   false,
		Appearance:   textures[Wood],
		ReturnAmount: 1,
	},
	Leaf: {
		Name:         Leaf,
		Amount:       1,
		Value:        1,
		IsCrop:       true,
		IsSolid:      true,
		HP:           1,
		Invincible:   false,
		Appearance:   textures[Leaf],
		ReturnAmount: 1,
	},
}

var WheatTile = Tile{
	Blocks: []BlockName{
		Wheat,
	},
	Height: 1,
	Width:  1,
}

var GrassTile = Tile{
	Blocks: []BlockName{
		Grass,
	},
	Height: 1,
	Width:  1,
}

var Tree1 = Tile{
	Blocks: []BlockName{
		Leaf, Leaf, Leaf,
		Empty, Wood, Empty,
		Empty, Wood, Empty,
		Empty, Wood, Empty,
	},
	Height: 4,
	Width:  3,
}
