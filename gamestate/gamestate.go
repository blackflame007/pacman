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
	board                    board.Board
	pacman                   pacman.Pacman
	ghosts                   []ghost.Ghost
	score                    int
	powerPellet              int
	lives                    int
	spawn                    pacman.Coordinates // New field to store spawn point coordinates
	frightenedMode           bool               // Indicates if ghosts are in frightened mode
	frightenedTurns          int                // Number of turns left in frightened mode
	ghostSpawnX, ghostSpawnY int                // Coordinates of the ghost spawn cage
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
		board:           board,
		pacman:          pacman.Pacman{X: 13, Y: 23},
		ghosts:          ghosts,
		score:           0,
		powerPellet:     0,
		lives:           3,
		spawn:           pacman.Coordinates{X: 13, Y: 23}, // Set spawn coordinates to initial Pacman position
		frightenedMode:  false,
		frightenedTurns: 0,
		ghostSpawnX:     14, // Example coordinates for the ghost spawn cage
		ghostSpawnY:     12,
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

	// Add Win condition here
	// TODO: if all dots are collected bt the player they win the game
	// Check if all dots are collected
	if g.areAllDotsCollected() {
		// Player wins the game
		// Return line to the player saying they won
		return g, tea.Quit
	}

	// Handle frightened mode turn countdown
	if g.frightenedMode {
		g.frightenedTurns--
		if g.frightenedTurns <= 0 {
			g.frightenedMode = false
		}
	}

	return g, nil
}

func (g *gameState) areAllDotsCollected() bool {
	for i := range g.board.Cells {
		for j := range g.board.Cells[i] {
			if g.board.Cells[i][j] == '.' {
				return false
			}
		}
	}
	return true
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
	out += fmt.Sprintf("Power Pellets Timer: %d\n", g.frightenedTurns)
	out += fmt.Sprintf("Lives: %d\n", g.lives)
	return out
}
