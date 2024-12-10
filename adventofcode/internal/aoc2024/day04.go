package aoc2024

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day04(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	xmas := aoc.ReadRuneGrid(reader)

	for i := 0; i < len(xmas); i++ {
		for j := 0; j < len(xmas[i]); j++ {
			if xmas[i][j] == 'X' {
				if i < len(xmas)-3 && xmas[i+1][j] == 'M' && xmas[i+2][j] == 'A' && xmas[i+3][j] == 'S' {
					part1++
				}

				if j < len(xmas[i])-3 && xmas[i][j+1] == 'M' && xmas[i][j+2] == 'A' && xmas[i][j+3] == 'S' {
					part1++
				}

				if i >= 3 && xmas[i-1][j] == 'M' && xmas[i-2][j] == 'A' && xmas[i-3][j] == 'S' {
					part1++
				}

				if j >= 3 && xmas[i][j-1] == 'M' && xmas[i][j-2] == 'A' && xmas[i][j-3] == 'S' {
					part1++
				}

				if i < len(xmas)-3 && j < len(xmas[i])-3 && xmas[i+1][j+1] == 'M' && xmas[i+2][j+2] == 'A' && xmas[i+3][j+3] == 'S' {
					part1++
				}

				if i >= 3 && j >= 3 && xmas[i-1][j-1] == 'M' && xmas[i-2][j-2] == 'A' && xmas[i-3][j-3] == 'S' {
					part1++
				}

				if i < len(xmas)-3 && j >= 3 && xmas[i+1][j-1] == 'M' && xmas[i+2][j-2] == 'A' && xmas[i+3][j-3] == 'S' {
					part1++
				}

				if i >= 3 && j < len(xmas[i])-3 && xmas[i-1][j+1] == 'M' && xmas[i-2][j+2] == 'A' && xmas[i-3][j+3] == 'S' {
					part1++
				}
			}
		}
	}

	for i := 1; i < len(xmas)-1; i++ {
		for j := 1; j < len(xmas[i])-1; j++ {
			if xmas[i][j] == 'A' {
				if ((xmas[i-1][j-1] == 'M' && xmas[i+1][j+1] == 'S') || (xmas[i-1][j-1] == 'S' && xmas[i+1][j+1] == 'M')) &&
					((xmas[i-1][j+1] == 'M' && xmas[i+1][j-1] == 'S') || (xmas[i-1][j+1] == 'S' && xmas[i+1][j-1] == 'M')) {
					part2++
				}
			}
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
