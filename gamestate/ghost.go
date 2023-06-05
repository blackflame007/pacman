package gamestate

import "pacman/ghost"

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
