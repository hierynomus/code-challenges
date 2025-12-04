package aoc2025

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day04(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	rolls := aoc.ReadRuneGrid(reader)
	removable := Removable(rolls)
	part1 = len(removable)
	part2 = part1
	for len(removable) > 0 {
		for _, p := range removable {
			rolls[p.Y][p.X] = '.'
		}
		removable = Removable(rolls)
		part2 += len(removable)
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func Removable(rolls [][]rune) []aoc.Point {
	removable := []aoc.Point{}
	for y := 0; y < len(rolls); y++ {
		for x := 0; x < len(rolls[y]); x++ {
			if rolls[y][x] == '@' {
				neighbours := aoc.Neighbours8(x, y)
				cnt := 0
				for _, n := range neighbours {
					if aoc.InBounds(n, rolls) && rolls[n.Y][n.X] == '@' {
						cnt++
					}
				}

				if cnt < 4 {
					removable = append(removable, aoc.Point{X: x, Y: y})
				}
			}
		}
	}
	return removable
}
