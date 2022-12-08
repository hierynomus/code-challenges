package aoc2022

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day08(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	trees := aoc.ReadIntGrid(reader, "")
	visible := map[aoc.Point]int{}

	for y := 0; y < len(trees); y++ {
		maxHeight := -1
		for x := 0; x < len(trees[y]); x++ {
			if trees[y][x] > maxHeight {
				maxHeight = trees[y][x]
				visible[aoc.Point{X: x, Y: y}] = trees[y][x]
			}
		}
		maxHeight = -1
		for x := len(trees[y]) - 1; x >= 0; x-- {
			if trees[y][x] > maxHeight {
				maxHeight = trees[y][x]
				visible[aoc.Point{X: x, Y: y}] = trees[y][x]
			}
		}
	}

	for x := 0; x < len(trees[0]); x++ {
		maxHeight := -1
		for y := 0; y < len(trees); y++ {
			if trees[y][x] > maxHeight {
				maxHeight = trees[y][x]
				visible[aoc.Point{X: x, Y: y}] = trees[y][x]
			}
		}
		maxHeight = -1
		for y := len(trees) - 1; y >= 0; y-- {
			if trees[y][x] > maxHeight {
				maxHeight = trees[y][x]
				visible[aoc.Point{X: x, Y: y}] = trees[y][x]
			}
		}
	}

	part1 = len(visible)
	maxScenicScore := 0
	for y := 1; y < len(trees)-1; y++ {
		for x := 1; x < len(trees[y])-1; x++ {
			height := trees[y][x]
			up, down, left, right := 0, 0, 0, 0
			for i := y - 1; i >= 0; i-- {
				if trees[i][x] < height {
					up++
				} else {
					up++
					break
				}
			}

			for i := y + 1; i < len(trees); i++ {
				if trees[i][x] < height {
					down++
				} else {
					down++
					break
				}
			}

			for i := x - 1; i >= 0; i-- {
				if trees[y][i] < height {
					left++
				} else {
					left++
					break
				}
			}

			for i := x + 1; i < len(trees[y]); i++ {
				if trees[y][i] < height {
					right++
				} else {
					right++
					break
				}
			}

			scenicScore := up * down * left * right
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	part2 = maxScenicScore
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
