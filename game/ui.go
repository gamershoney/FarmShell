package game

import (
	"charm.land/lipgloss/v2"
)

func (m model) RenderHeader() string {
	name := "Player: " + m.Player.Name

	style := lipgloss.NewStyle().Foreground(TextBlue).Padding(2)
	return style.Render(name)
}
