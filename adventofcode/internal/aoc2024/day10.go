package aoc2024

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day10(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	topography := aoc.ReadIntGrid(reader, "")
	trailheads := []aoc.Point{}
	for y, row := range topography {
		for x, cell := range row {
			if cell == 0 {
				trailheads = append(trailheads, aoc.Point{X: x, Y: y})
			}
		}
	}

	for _, trailhead := range trailheads {
		ends := FindEnds(topography, trailhead)
		part2 += len(ends)
		ps := aoc.NewPointSet(ends)
		part1 += len(ps)
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func FindEnds(topography [][]int, start aoc.Point) []aoc.Point {
	ends := []aoc.Point{}
	height := topography[start.Y][start.X]

	if height == 9 {
		return []aoc.Point{start}
	}

	for _, n := range start.Neighbours4() {
		if aoc.InBounds(n, topography) && topography[n.Y][n.X] == height+1 {
			ends = append(ends, FindEnds(topography, n)...)
		}
	}

	return ends
}
