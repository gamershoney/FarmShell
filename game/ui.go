package game

import (
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"
)

func (m model) RenderHeader() string {
	name := "Player: " + m.Player.Name

	style := lipgloss.NewStyle().Foreground(TextBlue).Padding(2)
	return style.Render(name)
}

func (m *model) RenderRight() string {
	title := "INVENTORY"
	title = lipgloss.NewStyle().Foreground(FancyRed).Bold(true).Render(title)

	nameStyle := lipgloss.NewStyle().
		Width(10).
		Align(lipgloss.Left).
		Foreground(TextBlue)

	qtyStyle := lipgloss.NewStyle().
		Width(8).
		Align(lipgloss.Right).
		Foreground(FancyOrange)

	highlightBG := FancyGray

	var b strings.Builder

	for index, block := range m.Player.Inventory {

		ns := nameStyle
		qs := qtyStyle

		if index == m.Player.SelectedBlock {
			ns = ns.Background(highlightBG)
			qs = qs.Background(highlightBG)
		}

		name := ns.Render(string(block.Name))
		quantity := qs.Render(strconv.Itoa(block.Qty))

		line := name + quantity

		b.WriteString(line)
		b.WriteByte('\n')
	}

	inv := lipgloss.JoinVertical(lipgloss.Center, title, b.String())
	return inv
}
