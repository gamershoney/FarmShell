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

func (m model) RenderRight() string {
	title := "INVENTORY"
	title = lipgloss.NewStyle().Foreground(FancyRed).Bold(true).Render(title)

	nameStyle := lipgloss.NewStyle().Width(10).Align(lipgloss.Left).Foreground(TextBlue)
	qtyStyle := lipgloss.NewStyle().Width(8).Align(lipgloss.Right).Foreground(FancyOrange)

	var b strings.Builder

	for item, qty := range m.Player.Inventory {
		name := nameStyle.Render(string(item))
		qty := qtyStyle.Render(strconv.Itoa(qty))

		line := lipgloss.JoinHorizontal(lipgloss.Top, name, qty)

		b.WriteString(line)
		b.WriteByte('\n')
	}

	inv := lipgloss.JoinVertical(lipgloss.Center, title, b.String())
	return inv
}
