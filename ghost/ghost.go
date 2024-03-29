package ghost

import (
	"math/rand"
	"pacman/board"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Ghost struct {
	X, Y         int
	Direction    Direction
	PreviousCell board.Cell // New field to store the previous cell value
}

func NewGhost(x, y int) Ghost {
	return Ghost{
		X:         x,
		Y:         y,
		Direction: Direction(rand.Intn(4)),
	}
}

func (g *Ghost) Move(dx, dy int) {
	g.X += dx
	g.Y += dy
}

func (g Ghost) ChooseDirection(board [][]rune) (int, int) {
	// Randomly choose a direction for now. Later this can be replaced with more sophisticated AI logic.
	direction := Direction(rand.Intn(4))

	switch direction {
	case Up:
		return 0, -1
	case Down:
		return 0, 1
	case Left:
		return -1, 0
	default: // Right
		return 1, 0
	}
}

func (g *Ghost) View() string {
	return "G"
}
