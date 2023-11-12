package gamestate

import "pacman/ghost"

func (g *gameState) moveGhost(ghost *ghost.Ghost) {
	dx, dy := ghost.ChooseDirection(boardToRunes(g.board.Cells))

	newX := ghost.X + dx
	newY := ghost.Y + dy

	// Wraparound logic for moving through tunnels
	if newY < 0 {
		newY = len(g.board.Cells) - 1
	} else if newY >= len(g.board.Cells) {
		newY = 0
	}
	if newX < 0 {
		newX = len(g.board.Cells[0]) - 1
	} else if newX >= len(g.board.Cells[0]) {
		newX = 0
	}

	// Check if the new position is not a wall and not another ghost
	if g.board.Cells[newY][newX] != '#' && g.board.Cells[newY][newX] != 'G' {
		// Check if the current cell is a pellet, then restore it
		if ghost.PreviousCell == '.' {
			g.board.Cells[ghost.Y][ghost.X] = ghost.PreviousCell
		}
		if ghost.PreviousCell == 'O' {
			g.board.Cells[ghost.Y][ghost.X] = ghost.PreviousCell
		} else {
			g.board.Cells[ghost.Y][ghost.X] = ' ' // Leave empty space
		}

		// Update PreviousCell with the content of the new cell
		ghost.PreviousCell = g.board.Cells[newY][newX]

		// Move the ghost to the new position
		ghost.X = newX
		ghost.Y = newY
		g.board.Cells[ghost.Y][ghost.X] = 'G'
	}
}
