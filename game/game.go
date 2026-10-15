package game

type BlockName string

const Wheat BlockName = "Wheat"
const Grass BlockName = "Grass"
const PlayerBlock BlockName = "Player"

var BlockMap = map[BlockName]Block{
	Wheat: {
		Name:    Wheat,
		Amount:  1,
		Value:   1,
		IsCrop:  true,
		IsSolid: false,
		HP:      1,
	},
	Grass: {
		Name:    Grass,
		Amount:  1,
		Value:   0,
		IsCrop:  false,
		IsSolid: false,
		HP:      100000,
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
