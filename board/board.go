// board.go
package board

const (
	Height = 31
	Width  = 28
)

type Cell rune

type Board struct {
	Cells [][]Cell
}

// Initialize the Board
func NewBoard() Board {
	cells := make([][]Cell, Height)
	for i := range cells {
		cells[i] = make([]Cell, Width)
	}

	// Add borders to the Board
	for i := range cells[0] {
		cells[0][i] = '-'
		cells[len(cells)-1][i] = '-'
	}
	for i := range cells {
		cells[i][0] = '|'
		cells[i][len(cells[i])-1] = '|'
	}

	// Generate Maze
	GenerateMaze(cells)

	return Board{Cells: cells}
}

// GenerateMaze generates a simple maze layout for the board
func GenerateMaze(cells [][]Cell) {
	maze := []string{
		"############################",
		"#............##............#",
		"#.####.#####.##.#####.####.#",
		"#O####.#####.##.#####.####O#",
		"#.####.#####.##.#####.####.#",
		"#..........................#",
		"#.####.##.########.##.####.#",
		"#.####.##.########.##.####.#",
		"#......##....##....##......#",
		"######.##### ## #####.######",
		"######.##### ## #####.######",
		"######.##          ##.######",
		"######.## ##    ## ##.######",
		"######.## #      # ##.######",
		".......   #      #   .......",
		"######.## #      # ##.######",
		"######.## ######## ##.######",
		"######.##          ##.######",
		"######.## ######## ##.######",
		"######.## ######## ##.######",
		"#............##............#",
		"#.####.#####.##.#####.####.#",
		"#.####.#####.##.#####.####.#",
		"#O..##................##..O#",
		"###.##.##.########.##.##.###",
		"###.##.##.########.##.##.###",
		"#......##....##....##......#",
		"#.##########.##.##########.#",
		"#.##########.##.##########.#",
		"#..........................#",
		"############################",
	}

	for i, row := range maze {
		for j, col := range row {
			switch col {
			case '#':
				cells[i][j] = '#'
			case '.':
				cells[i][j] = '.'
			case 'O':
				cells[i][j] = 'O'
			default:
				cells[i][j] = ' '
			}
		}
	}
}
