package game

func InitGameLoop() {

}

type Crop struct {
	Name   string  `json:"crop_name"`
	Amount int     `json:"amount"`
	Value  float64 `json:"crop_value"`
}

const TileHeight int = 8

const TileWidth int = 8

type Tile struct {
	Type *Crop `json:"crop"`
}

type TilePos struct {
	x int
	y int
}

type model struct {
	Playername string           `json:"username"`
	Crops      map[string]*Crop `json:"player_crops"`
	Tiles      map[TilePos]*Tile
}

func newGameModel() model {
	defUser := "big_player"

	crops := make(map[string]*Crop)

	Tiles := make(map[TilePos]*Tile)
	model := model{
		Playername: defUser,
		Crops:      crops,
		Tiles:      Tiles,
	}

	return model
}
