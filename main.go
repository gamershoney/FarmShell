package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/gamershoney/FarmShell/game"
)

func main() {
	prog := tea.NewProgram(game.NewGameModel())
	if _, err := prog.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
