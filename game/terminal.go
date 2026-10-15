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
	Name       BlockName `json:"crop_name"`
	Amount     int       `json:"amount"`
	Value      float64   `json:"crop_value"`
	IsCrop     bool
	IsSolid    bool
	HP         float64
	Pos        BlockPos
	Appearance Appearance
}

type BlockPos struct {
	x int
	y int
}

func (b BlockPos) Add(n BlockPos) (r BlockPos) {
	r.x = b.x + n.x
	r.y = b.y + n.y

	return r
}

type Tile struct {
	Blocks []BlockName
	Height int
	Width  int
}

type Direction = BlockPos

var (
	North = Direction{x: 0, y: -1}
	South = Direction{x: 0, y: 1}
	East  = Direction{x: 1, y: 0}
	West  = Direction{x: -1, y: 0}
)

type Player struct {
	Pos       BlockPos
	Name      string
	Direction Direction
	Texture   Appearance
}

func IsNavigable(m *model, pos BlockPos) bool {
	_, ok := m.Blocks[pos]
	if !ok {
		model := m.GenerateTile(pos)
		m.Blocks = model.Blocks
	}
	block := m.Blocks[pos]

	if block.IsSolid {
		return false
	}
	return true
}

func (m model) MovePlayer(dir Direction) model {
	requested_pos := m.Player.Pos.Add(dir)
	navigable := IsNavigable(&m, requested_pos)

	if !navigable {
		return m
	}
	m.Player.Pos = requested_pos
	return m

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
		item.Appearance = textures[item.Name]
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
			style := lipgloss.NewStyle().Background(block.Appearance.BG).Foreground(block.Appearance.FG)

			b.WriteString(style.Render(string(block.Appearance.Char)))
		}
		b.WriteRune('\n')
	}
	return tea.NewView(b.String())
}

func (m *model) RenderPlayer() {
	_, ok := m.Blocks[m.Player.Pos]

	if !ok {
		m.Blocks = m.GenerateTile(m.Player.Pos).Blocks
	}

	block := m.Blocks[m.Player.Pos]

	block.Appearance.FG = m.Player.Texture.FG
	block.Appearance.Char = m.Player.Texture.Char
}

func NewGameModel() model {
	seed := int64(12345)

	rng := rand.New(rand.NewSource(seed))

	defUser := Player{
		Pos: BlockPos{
			x: 0,
			y: 0,
		},
		Texture: textures[PlayerBlock],
		Name:    "New Player",
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
	m.RenderPlayer()
	view := m.RenderTiles()

	return view
}
