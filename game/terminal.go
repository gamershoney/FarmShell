package game

import "image/color"

func InitGameLoop() bool {

	return true
}

type Appearance struct {
	FG   color.Color
	BG   color.Color
	Char rune
}

type Block struct {
	Name    BlockName `json:"crop_name"`
	Amount  int       `json:"amount"`
	Value   float64   `json:"crop_value"`
	IsCrop  bool
	IsSolid bool
	HP      float64
	Pos     BlockPos
}

type BlockPos struct {
	x int
	y int
}

type Tile struct {
	Blocks []BlockName
	Height int
	Width  int
}

type Player struct {
	Pos  BlockPos
	Name string
}
type model struct {
	Crops    map[string]*Block `json:"player_crops"`
	Blocks   map[BlockPos]*Block
	ViewPort ViewPort
	Player   Player
}

type ViewPort struct {
	Width    int
	Height   int
	Position int
}

func newGameModel() model {
	defUser := Player{
		Pos: BlockPos{
			x: 0,
			y: 0,
		},
		Name: "New Player",
	}

	crops := make(map[string]*Block)

	Block := make(map[BlockPos]*Block)
	model := model{
		Player: defUser,
		Crops:  crops,
		Blocks: Block,
	}

	return model
}
