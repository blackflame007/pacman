// gamestate.go
package gamestate

import (
	"fmt"
	"pacman/board"
	"pacman/ghost"
	"pacman/pacman"

	tea "github.com/charmbracelet/bubbletea"
)

type gameState struct {
	board       board.Board
	pacman      pacman.Pacman
	ghosts      []ghost.Ghost
	score       int
	powerPellet int
}

func NewGameState() *gameState {
	board := board.NewBoard()

	// Add Pacman to the board
	// board.Cells[13][23] = 'P'

	// Add Ghosts to the board
	ghosts := []ghost.Ghost{
		{X: 14, Y: 11},
		{X: 13, Y: 15},
		{X: 14, Y: 14},
		{X: 15, Y: 15},
	}

	for _, g := range ghosts {
		board.Cells[g.Y][g.X] = 'G'
	}

	return &gameState{
		board:       board,
		pacman:      pacman.Pacman{X: 13, Y: 23},
		ghosts:      ghosts,
		score:       0,
		powerPellet: 0,
	}
}

func (g gameState) Init() tea.Cmd {
	return nil
}

func (g *gameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return g, tea.Quit
		case "w":
			g.movePacman(0, -1)
		case "a":
			g.movePacman(-1, 0)
		case "s":
			g.movePacman(0, 1)
		case "d":
			g.movePacman(1, 0)
		}
	}

	return g, nil
}

func (g *gameState) movePacman(dx, dy int) {
	newX := g.pacman.X + dx
	newY := g.pacman.Y + dy
	if g.board.Cells[newY][newX] != '#' {
		switch g.board.Cells[newY][newX] {
		case '.':
			g.score++
		case 'O':
			g.score += 10
			g.powerPellet++
		}
		g.board.Cells[g.pacman.Y][g.pacman.X] = ' '
		g.pacman.X = newX
		g.pacman.Y = newY
		g.board.Cells[g.pacman.Y][g.pacman.X] = 'P'
	}
}

func (g gameState) View() string {
	var out string
	for i := range g.board.Cells {
		for j := range g.board.Cells[i] {
			out += string(g.board.Cells[i][j])
		}
		out += "\n"
	}
	out += fmt.Sprintf("Score: %d\n", g.score)
	out += fmt.Sprintf("Power Pellets: %d\n", g.powerPellet)
	return out
}
