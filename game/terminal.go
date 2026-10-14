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
}

type Tile struct {
	Blocks []BlockName
	Height int
	Width  int
}

type TilePos struct {
	x int
	y int
}

type model struct {
	Playername string            `json:"username"`
	Crops      map[string]*Block `json:"player_crops"`
	Tiles      map[TilePos]*Tile
}

func newGameModel() model {
	defUser := "big_player"

	crops := make(map[string]*Block)

	Tiles := make(map[TilePos]*Tile)
	model := model{
		Playername: defUser,
		Crops:      crops,
		Tiles:      Tiles,
	}

	return model
}
