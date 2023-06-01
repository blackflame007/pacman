package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nsf/termbox-go"
)

type cell rune

type board struct {
	cells [][]cell
}

type pacman struct {
	x, y int
}

type ghost struct {
	x, y int
}

type gameState struct {
	board       board
	pacman      pacman
	ghosts      []ghost
	score       int
	powerPellet int
}

func newGameState() gameState {
	cells := make([][]cell, 20)
	for i := range cells {
		cells[i] = make([]cell, 20)
	}
	// Initialize the board with dots
	for i := range cells {
		for j := range cells[i] {
			cells[i][j] = '.'
		}
	}
	// Add borders to the board
	for i := range cells[0] {
		cells[0][i] = '-'
		cells[len(cells)-1][i] = '-'
	}
	for i := range cells {
		cells[i][0] = '|'
		cells[i][len(cells[i])-1] = '|'
	}

	// Add internal walls
	for i := 5; i < 15; i++ {
		cells[i][8] = '|'
		cells[i][12] = '|'
	}

	// Add Pacman to the board
	cells[10][10] = 'P'
	// Add Ghosts to the board
	ghosts := []ghost{
		{x: 5, y: 5},
		{x: 15, y: 5},
		{x: 5, y: 15},
		{x: 15, y: 15},
	}
	for _, g := range ghosts {
		cells[g.y][g.x] = 'G'
	}
	return gameState{
		board:       board{cells: cells},
		pacman:      pacman{x: 10, y: 10},
		ghosts:      ghosts,
		score:       0,
		powerPellet: 0,
	}
}

func (g gameState) Init() tea.Cmd {
	return nil
}

func (g gameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	newX := g.pacman.x + dx
	newY := g.pacman.y + dy
	if g.board.cells[newY][newX] != '-' && g.board.cells[newY][newX] != '|' {
		switch g.board.cells[newY][newX] {
		case '.':
			g.score++
		case 'O':
			g.score += 10
			g.powerPellet++
		}
		g.board.cells[g.pacman.y][g.pacman.x] = ' '
		g.pacman.x = newX
		g.pacman.y = newY
		g.board.cells[g.pacman.y][g.pacman.x] = 'P'
	}
}

func (g gameState) View() string {
	var out string
	for i := range g.board.cells {
		for j := range g.board.cells[i] {
			out += string(g.board.cells[i][j])
		}
		out += "\n"
	}
	out += fmt.Sprintf("Score: %d\n", g.score)
	out += fmt.Sprintf("Power Pellets: %d\n", g.powerPellet)
	return out
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	initialModel := newGameState()

	program := tea.NewProgram(initialModel)
	if err := program.Start(); err != nil {
		panic(err)
	}
}
