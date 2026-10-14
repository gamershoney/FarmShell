package game

type BlockName string

const Wheat BlockName = "Wheat"

var BlockMap = map[BlockName]Block{
	Wheat: {
		Name:    Wheat,
		Amount:  1,
		Value:   1,
		IsCrop:  true,
		IsSolid: false,
		HP:      1,
	},
}

var WheatTile = Tile{
	Blocks: []BlockName{
		Wheat,
	},
	Height: 1,
	Width:  1,
}
