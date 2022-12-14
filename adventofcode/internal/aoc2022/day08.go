package aoc2022

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Forest [][]int

func (f Forest) ScenicScore(x, y int) int {
	p := aoc.Point{X: x, Y: y}
	return aoc.Manhattan(p, f.LineOfSight(x, y, -1, 0)) *
		aoc.Manhattan(p, f.LineOfSight(x, y, 1, 0)) *
		aoc.Manhattan(p, f.LineOfSight(x, y, 0, -1)) *
		aoc.Manhattan(p, f.LineOfSight(x, y, 0, 1))
}

func (f Forest) LineOfSight(ox, oy, dx, dy int) aoc.Point {
	if ox <= 0 || oy <= 0 || ox >= len(f[oy])-1 || oy >= len(f)-1 {
		return aoc.Point{X: ox, Y: oy}
	}

	height := f[oy][ox]
	x, y := ox, oy
	for {
		x, y = x+dx, y+dy
		if x <= 0 || y <= 0 || x >= len(f[y])-1 || y >= len(f)-1 {
			return aoc.Point{X: x, Y: y}
		}
		if f[y][x] >= height {
			return aoc.Point{X: x, Y: y}
		}
	}
}

func CompactedCalc(forest Forest) [2]int {
	visible := map[aoc.Point]int{}
	maxScenicScore := 0

	for y := 0; y < len(forest); y++ {
		visible[aoc.Point{X: 0, Y: y}] = forest[y][0]
		visible[aoc.Point{X: len(forest[y]) - 1, Y: y}] = forest[y][len(forest[y])-1]
	}
	for x := 0; x < len(forest[0]); x++ {
		visible[aoc.Point{X: x, Y: 0}] = forest[0][x]
		visible[aoc.Point{X: x, Y: len(forest) - 1}] = forest[len(forest)-1][x]
	}

	for y := 1; y < len(forest)-1; y++ {
		for x := 1; x < len(forest[y])-1; x++ {
			scenicScore := 1
			for _, p := range aoc.Origin.Neighbours4() {
				los := forest.LineOfSight(x, y, p.X, p.Y)
				if (p.X == -1 && los.X == 0 && forest[y][x] > forest[y][0]) ||
					(p.X == 1 && los.X == len(forest[y])-1 && forest[y][x] > forest[y][len(forest[y])-1]) ||
					(p.Y == -1 && los.Y == 0 && forest[y][x] > forest[0][x]) ||
					(p.Y == 1 && los.Y == len(forest)-1 && forest[y][x] > forest[len(forest)-1][x]) {
					visible[aoc.Point{X: x, Y: y}] = forest[y][x]
				}

				scenicScore *= aoc.Manhattan(aoc.Point{X: x, Y: y}, los)
			}

			if scenicScore > maxScenicScore {
				// fmt.Printf("%d,%d -> %d\n", x, y, scenicScore)
				maxScenicScore = scenicScore
			}
		}
	}

	return [2]int{len(visible), maxScenicScore}
}

func OriginalCalc(forest Forest) [2]int {
	visible := map[aoc.Point]int{}
	maxScenicScore := 0

	for y := 0; y < len(forest); y++ {
		maxHeight := -1
		for x := 0; x < len(forest[y]); x++ {
			if forest[y][x] > maxHeight {
				maxHeight = forest[y][x]
				visible[aoc.Point{X: x, Y: y}] = forest[y][x]
			}
		}
		maxHeight = -1
		for x := len(forest[y]) - 1; x >= 0; x-- {
			if forest[y][x] > maxHeight {
				maxHeight = forest[y][x]
				visible[aoc.Point{X: x, Y: y}] = forest[y][x]
			}
		}
	}

	for x := 0; x < len(forest[0]); x++ {
		maxHeight := -1
		for y := 0; y < len(forest); y++ {
			if forest[y][x] > maxHeight {
				maxHeight = forest[y][x]
				visible[aoc.Point{X: x, Y: y}] = forest[y][x]
			}
		}
		maxHeight = -1
		for y := len(forest) - 1; y >= 0; y-- {
			if forest[y][x] > maxHeight {
				maxHeight = forest[y][x]
				visible[aoc.Point{X: x, Y: y}] = forest[y][x]
			}
		}
	}

	for y := 0; y < len(forest); y++ {
		for x := 0; x < len(forest[y]); x++ {
			scenicScore := forest.ScenicScore(x, y)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	return [2]int{len(visible), maxScenicScore}
}

func Day08(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	forest := Forest(aoc.ReadIntGrid(reader, ""))
	pts := OriginalCalc(forest)
	part1, part2 = pts[0], pts[1]
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
