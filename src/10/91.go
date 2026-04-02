package main

type XY struct {
	X int
	Y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	if grid[0][0] == 1 {
		return -1
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				grid[i][j] = -1
			}
		}
	}

	grid[0][0] = 1
	queue := make([]XY, 0)
	queue = append(queue, XY{0, 0})

	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]

		nextDistance := grid[c.X][c.Y] + 1

		if c.X > 0 && c.Y > 0 && grid[c.X-1][c.Y-1] == 0 {
			grid[c.X-1][c.Y-1] = nextDistance
			queue = append(queue, XY{c.X - 1, c.Y - 1})
		}
		if c.X > 0 && grid[c.X-1][c.Y] == 0 {
			grid[c.X-1][c.Y] = nextDistance
			queue = append(queue, XY{c.X - 1, c.Y})
		}
		if c.X > 0 && c.Y < m-1 && grid[c.X-1][c.Y+1] == 0 {
			grid[c.X-1][c.Y+1] = nextDistance
			queue = append(queue, XY{c.X - 1, c.Y + 1})
		}
		if c.Y < m-1 && grid[c.X][c.Y+1] == 0 {
			grid[c.X][c.Y+1] = nextDistance
			queue = append(queue, XY{c.X, c.Y + 1})
		}
		if c.X < n-1 && c.Y < m-1 && grid[c.X+1][c.Y+1] == 0 {
			grid[c.X+1][c.Y+1] = nextDistance
			queue = append(queue, XY{c.X + 1, c.Y + 1})
		}
		if c.X < n-1 && grid[c.X+1][c.Y] == 0 {
			grid[c.X+1][c.Y] = nextDistance
			queue = append(queue, XY{c.X + 1, c.Y})
		}
		if c.X < n-1 && c.Y > 0 && grid[c.X+1][c.Y-1] == 0 {
			grid[c.X+1][c.Y-1] = nextDistance
			queue = append(queue, XY{c.X + 1, c.Y - 1})
		}
		if c.Y > 0 && grid[c.X][c.Y-1] == 0 {
			grid[c.X][c.Y-1] = nextDistance
			queue = append(queue, XY{c.X, c.Y - 1})
		}
	}

	if grid[n-1][m-1] < 1 {
		return -1
	}
	return grid[n-1][m-1]
}
