// main.go
package main

import (
	"pacman/gamestate"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	initialModel := gamestate.NewGameState()

	program := tea.NewProgram(initialModel)
	if err := program.Start(); err != nil {
		panic(err)
	}
}
