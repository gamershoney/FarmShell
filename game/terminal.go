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
	m.Player.InteractBox = m.Player.Pos.Add(BlockPos{1, 0})

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
	Ticks    uint64
	Paused   bool
}

func (m model) GenerateTile(pos BlockPos) model {
	roll := m.Rng.Float64()

	var tile Tile

	switch {
	case roll < .01:
		tile = Tree1
	case roll < .05:
		tile = WheatTile
	default:
		tile = GrassTile
	}

	cursor := BlockPos{x: 0, y: 0}
	for _, block := range tile.Blocks {
		if cursor.x >= tile.Width {
			cursor.x = 0
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
		item.Pos = BlockPos{x: pos.x + cursor.x, y: pos.y + cursor.y}
		m.Blocks[item.Pos] = &item

		cursor.x++

	}

	return m
}

func (m model) RenderTiles() string {
	m.ViewPort.Position = m.Player.Pos
	halfwidth := m.ViewPort.Width / 2
	halfheight := m.ViewPort.Height / 2
	startX := m.ViewPort.Position.x - halfwidth
	startY := m.ViewPort.Position.y - halfheight

	var b strings.Builder

	for y := startY; y < startY+m.ViewPort.Height; y++ {
		for x := startX; x < startX+m.ViewPort.Width; x++ {
			pos := BlockPos{
				x: x,
				y: y,
			}

			if pos == m.Player.Pos {
				style := lipgloss.NewStyle().Background(m.Player.Texture.BG).Foreground(m.Player.Texture.FG)
				b.WriteString(style.Render(string(m.Player.Texture.Char)))
				continue
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
			if pos == m.Player.InteractBox {
				highlight := lipgloss.Lighten(block.Appearance.BG, .2)
				style = style.Background(highlight)
			}
			b.WriteString(style.Render(string(block.Appearance.Char)))
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func NewGameModel() model {
	seed := int64(12345)

	rng := rand.New(rand.NewSource(seed))

	defUser := Player{
		Pos: BlockPos{
			x: 0,
			y: 0,
		},
		Texture:       textures[PlayerBlock],
		Name:          "New Player",
		DefaultDamage: 1,
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
	return GameTick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case GameTickMsg:
		if !m.Paused {
			m.Ticks++
		}
		return m, GameTick()

	case tea.KeyPressMsg:

		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "up", "w":
			m = m.MovePlayer(North)
			return m, nil

		case "down", "s":
			m = m.MovePlayer(South)
			return m, nil

		case "left", "a":
			m = m.MovePlayer(West)
			return m, nil

		case "right", "d":
			m = m.MovePlayer(East)
			return m, nil

		case "space":
			m.Player.Attack(m)
			return m, nil

		case "]":
			m.Player.SelectNext()
			return m, nil

		case "[":
			m.Player.SelectPrev()
			return m, nil
		case "e":
			m.PlaceBlock()
			return m, nil
		}

	}
	m.Player.InteractBox = m.Player.Pos.Add(BlockPos{1, 0})
	return m, nil
}

func (m model) View() tea.View {
	game := m.RenderTiles()

	top := m.RenderHeader()
	right := m.RenderRight()
	screen := lipgloss.JoinVertical(lipgloss.Left, top, game)
	screen = lipgloss.JoinHorizontal(lipgloss.Left, screen, right)
	return tea.NewView(screen)
}
