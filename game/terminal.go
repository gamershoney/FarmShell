package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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

type Direction int

const North Direction = 0
const East Direction = 1
const South Direction = 2
const West Direction = 3

type Player struct {
	Pos       BlockPos
	Name      string
	Direction Direction
	Texture   Appearance
}

func (m model) MovePlayer(dir Direction) {
	switch dir {
	case condition:

	}

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

func (m model) GenerateTile(pos BlockPos) model {
	roll := m.Rng.Float64()

	var tile Tile

	switch {
	case roll < .05:
		tile = WheatTile

	default:
		tile = GrassTile
	}

	cursor := BlockPos{x: 0, y: 0}
	for _, block := range tile.Blocks {
		if cursor.x >= tile.Width {
			cursor.y++
		}
		if cursor.y >= tile.Height {
			break
		}
		item, ok := BlockMap[block]
		if !ok {
			fmt.Println("error could not find block in map: ", block)
		}
		m.Blocks[BlockPos{x: pos.x + cursor.x, y: pos.y + cursor.y}] = &item

	}

	return m
}

func (m model) RenderTiles() tea.View {
	m.ViewPort.Position = m.Player.Pos
	halfwidth := m.ViewPort.Width / 2
	halfheight := m.ViewPort.Height / 2

	var b strings.Builder

	for x := (m.ViewPort.Position.x - halfwidth); x < m.ViewPort.Width; x++ {
		for y := (m.ViewPort.Position.y - halfheight); y < m.ViewPort.Height; y++ {
			pos := BlockPos{
				x: x,
				y: y,
			}

			block, ok := m.Blocks[pos]
			if !ok {
				m = m.GenerateTile(pos)
			}
			block, ok = m.Blocks[pos]
			if !ok {
				m = m.GenerateTile(pos)
			}
			texture, ok := textures[block.Name]
			style := lipgloss.NewStyle().Background(texture.BG).Foreground(texture.FG)
			b.WriteString(style.Render(string(texture.Char)))
		}
		b.WriteRune('\n')
	}
	return tea.NewView(b.String())
}

func NewGameModel() model {
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
		ViewPort: ViewPort{
			Width:  100,
			Height: 30,
			Position: BlockPos{
				x: 0,
				y: 0,
			},
		},
	}

	return model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() tea.View {
	view := m.RenderTiles()

	return view
}
