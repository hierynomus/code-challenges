package aoc

func Transpose[A any](slice [][]A) [][]A {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]A, xl)

	for i := range result {
		result[i] = make([]A, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}

	return result
}

func RotateCW[A any](grid [][]A) [][]A {
	yl := len(grid)
	xl := len(grid[0])
	result := make([][]A, xl)

	for i := range result {
		result[i] = make([]A, yl)
	}

	for y := 0; y < yl; y++ {
		for x := 0; x < xl; x++ {
			result[x][yl-y-1] = grid[y][x]
		}
	}

	return result
}

func RotateCCW[A any](grid [][]A) [][]A {
	yl := len(grid)
	xl := len(grid[0])
	result := make([][]A, xl)

	for i := range result {
		result[i] = make([]A, yl)
	}

	for y := 0; y < yl; y++ {
		for x := 0; x < xl; x++ {
			result[x][y] = grid[y][xl-x-1]
		}
	}

	return result
}

func TransposeString(slice []string) []string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([]string, xl)

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i] += string(slice[j][i])
		}
	}

	return result
}
