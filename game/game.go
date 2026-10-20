package game

type BlockName string

const Wheat BlockName = "Wheat"
const Grass BlockName = "Grass"
const PlayerBlock BlockName = "Player"
const Wood BlockName = "Wood"

var BlockMap = map[BlockName]Block{
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
		Name:       Wood,
		Amount:     1,
		Value:      1,
		IsCrop:     true,
		IsSolid:    true,
		HP:         3,
		Invincible: false,
		Appearance: textures[Wood],
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
		Wood,
		Wood,
		Wood,
	},
	Height: 3,
	Width:  1,
}
