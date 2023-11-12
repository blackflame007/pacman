// main.go
package main

import (
	"pacman/gamestate"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	initialModel := gamestate.NewGameState()

	program := tea.NewProgram(initialModel)
	if _, err := program.Run(); err != nil {
		panic(err)
	}
}
