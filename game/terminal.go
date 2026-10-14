package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
)

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

type ViewPort struct {
	Width    int
	Height   int
	Position BlockPos
}

type model struct {
	Seed     int64
	Rng      *rand.Rand
	Crops    map[string]*Block `json:"player_crops"`
	Blocks   map[BlockPos]*Block
	ViewPort ViewPort
	Player   Player
}

func (m model) GenerateTile(pos BlockPos) {
	roll := m.Rng.Float64()

	var tile Tile

	switch {
	case roll < .05:
		tile = WheatTile

	default:
		tile = GrassTile
	}

	if !ok {
		fmt.Println("error: block not found in block map")
		os.Exit(1)
	}

	for i, block := range tile.Blocks {
	}
}

func (m model) RenderTiles() model {
	m.ViewPort.Position = m.Player.Pos
	halfwidth := m.ViewPort.Width / 2
	halfheight := m.ViewPort.Height / 2
	for x := (m.ViewPort.Position.x - halfwidth); x < m.ViewPort.Width; x++ {
		for y := (m.ViewPort.Position.y - halfheight); y < m.ViewPort.Height; y++ {
			pos := BlockPos{
				x: x,
				y: y,
			}

			block, ok := m.Blocks[pos]
			if !ok {

			}
		}
	}
}

func newGameModel() model {

	seed := int64(12345)

	rng := rand.New(rand.NewSource(seed))

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
		Seed:   seed,
		Rng:    rng,
	}

	return model
}
