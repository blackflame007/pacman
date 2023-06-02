// pacman.go
package pacman

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Coordinates struct {
	X, Y int
}

type Pacman struct {
	X, Y int
}

func NewPacman(x, y int) Pacman {
	return Pacman{X: x, Y: y}
}

func (d Direction) Delta() (int, int) {
	switch d {
	case Up:
		return 0, -1
	case Down:
		return 0, 1
	case Left:
		return -1, 0
	case Right:
		return 1, 0
	default:
		return 0, 0
	}
}

func (p *Pacman) View() string {
	return "P"
}
