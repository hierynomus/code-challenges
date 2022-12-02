package aoc2021

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day09(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	grid := aoc.ReadIntGrid(reader, "")
	lows := make([]aoc.Point, 0)

	for y, row := range grid {
		for x, cell := range row {
			neigh := aoc.Neighbours4(x, y)
			heights := make([]int, 0)
			for _, n := range neigh {
				if n.X < 0 || n.Y < 0 || n.X >= len(grid) || n.Y >= len(grid[0]) {
					continue
				}
				heights = append(heights, grid[n.Y][n.X])
			}
			if cell < aoc.Min(heights) {
				lows = append(lows, aoc.Point{X: x, Y: y})
				part1 += 1 + cell
			}
		}
	}

	basins := make(map[aoc.Point]aoc.PointSet)
	for _, p := range lows {
		basins[p] = aoc.PointSet{}
		basins[p].Add(p)
		frontier := p.Neighbours4()
		expand(grid, basins[p], frontier)
	}

	lengths := make([]int, 0)
	for _, basin := range basins {
		lengths = append(lengths, len(basin))
	}

	sort.Ints(lengths)
	part2 = lengths[len(lengths)-1] * lengths[len(lengths)-2] * lengths[len(lengths)-3]

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func expand(grid [][]int, basin aoc.PointSet, frontier []aoc.Point) {
	for _, n := range frontier {
		if basin.Contains(n) || n.X < 0 || n.Y < 0 || n.X >= len(grid) || n.Y >= len(grid[0]) || grid[n.Y][n.X] == 9 {
			continue
		}

		basin.Add(n)
		expand(grid, basin, n.Neighbours4())
	}
}
