package gamestate

func (g *gameState) respawnPlayer() {
	g.board.Cells[g.pacman.Y][g.pacman.X] = ' '
	g.pacman.X = g.spawn.X
	g.pacman.Y = g.spawn.Y
	g.board.Cells[g.pacman.Y][g.pacman.X] = 'P'
}

func (g *gameState) movePacman(dx, dy int) {
	newX := g.pacman.X + dx
	newY := g.pacman.Y + dy

	// Check if the new position is within the jail area
	if newX >= 12 && newX <= 15 && newY >= 12 && newY <= 14 {
		return
	}

	// Wraparound Pacman's position when moving through tunnels
	// Wraparound Pacman's position when moving through tunnels
	if newY < 0 {
		newY = len(g.board.Cells) - 1
	} else if newY >= len(g.board.Cells) {
		newY = 0
	} else if newX < 0 {
		newX = len(g.board.Cells[0]) - 1
	} else if newX >= len(g.board.Cells[0]) {
		newX = 0
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
