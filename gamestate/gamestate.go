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
	lives       int
	spawn       pacman.Coordinates // New field to store spawn point coordinates
}

func NewGameState() *gameState {
	board := board.NewBoard()

	// Add Pacman to the board
	board.Cells[23][13] = 'P'

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
		lives:       3,
		spawn:       pacman.Coordinates{X: 13, Y: 23}, // Set spawn coordinates to initial Pacman position
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

	// Move the ghosts
	for i := range g.ghosts {
		g.moveGhost(&g.ghosts[i])
	}

	return g, nil
}

func (g *gameState) movePacman(dx, dy int) {
	newX := g.pacman.X + dx
	newY := g.pacman.Y + dy

	// Check if the new position is within the jail area
	if newX >= 12 && newX <= 15 && newY >= 12 && newY <= 14 {
		return
	}

	if g.board.Cells[newY][newX] != '#' {
		switch g.board.Cells[newY][newX] {
		case '.':
			g.score++
			g.board.Cells[newY][newX] = ' ' // Remove the dot from the board
		case 'O':
			g.score += 10
			g.powerPellet++
		case 'G':
			g.lives--
			if g.lives == 0 {
				g.resetGame()
				return
			} else {
				g.respawnPlayer()
				return
			}
		}
		g.board.Cells[g.pacman.Y][g.pacman.X] = ' '
		g.pacman.X = newX
		g.pacman.Y = newY
		g.board.Cells[g.pacman.Y][g.pacman.X] = 'P'
	}
}

func boardToRunes(cells [][]board.Cell) [][]rune {
	runes := make([][]rune, len(cells))
	for i := range cells {
		runes[i] = make([]rune, len(cells[i]))
		for j, cell := range cells[i] {
			runes[i][j] = rune(cell)
		}
	}
	return runes
}

func (g *gameState) moveGhost(ghost *ghost.Ghost) {
	dx, dy := ghost.ChooseDirection(boardToRunes(g.board.Cells))

	newX := ghost.X + dx
	newY := ghost.Y + dy

	if g.board.Cells[newY][newX] != '#' && g.board.Cells[newY][newX] != 'G' {
		// Save the current cell value before moving the ghost
		previousCell := g.board.Cells[ghost.Y][ghost.X]

		// Move the ghost to the new position
		g.board.Cells[ghost.Y][ghost.X] = ' '
		ghost.X = newX
		ghost.Y = newY
		g.board.Cells[ghost.Y][ghost.X] = 'G'

		// Restore the previous cell value
		if previousCell == '.' {
			g.board.Cells[ghost.Y][ghost.X] = previousCell
		}
	}
}

func (g *gameState) respawnPlayer() {
	g.board.Cells[g.pacman.Y][g.pacman.X] = ' '
	g.pacman.X = g.spawn.X
	g.pacman.Y = g.spawn.Y
	g.board.Cells[g.pacman.Y][g.pacman.X] = 'P'
}

func (g *gameState) resetGame() {
	g.score = 0
	g.powerPellet = 0
	g.lives = 3

	// Reset Pacman's position
	g.board.Cells[g.pacman.Y][g.pacman.X] = ' '
	g.pacman.X = g.spawn.X
	g.pacman.Y = g.spawn.Y
	g.board.Cells[g.pacman.Y][g.pacman.X] = 'P'

	// Reset Ghosts' positions
	for i := range g.ghosts {
		g.board.Cells[g.ghosts[i].Y][g.ghosts[i].X] = ' '
		g.ghosts[i].X = 12 + i
		g.ghosts[i].Y = 14
		g.board.Cells[g.ghosts[i].Y][g.ghosts[i].X] = 'G'
	}

	g.resetDots() // Reset the dots on the board
}

func (g *gameState) resetDots() {
	for i := range g.board.Cells {
		for j := range g.board.Cells[i] {
			if g.board.Cells[i][j] != '#' && !(i >= 11 && i <= 15 && j >= 11 && j <= 16) {
				g.board.Cells[i][j] = '.'
			}
		}
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
	out += fmt.Sprintf("Lives: %d\n", g.lives)
	return out
}
