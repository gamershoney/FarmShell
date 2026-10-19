package game

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

const TickRate = time.Second / 20

type GameTickMsg time.Time

func GameTick() tea.Cmd {
	return tea.Tick(TickRate, func(t time.Time) tea.Msg {
		return GameTickMsg(t)
	})
}
