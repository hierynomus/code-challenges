package aoc2021

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day11(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	grid := make([][]int, 0)
	for reader.Scan() {
		line := aoc.AsIntArrayS(reader.Text(), "")
		grid = append(grid, line)
	}

	var flashes int
	i := 0

	for flashes != 100 {
		grid, flashes = OctopusStep(grid)
		i++
		if i <= 100 {
			part1 += flashes
		}
	}
	part2 = i

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func OctopusStep(grid [][]int) ([][]int, int) {
	flashes := []aoc.Point{}
	flashed := map[string]bool{}

	for y, row := range grid {
		for x := range row {
			grid[y][x]++
			if grid[y][x] > 9 {
				flashes = append(flashes, aoc.Point{X: x, Y: y})
			}
		}
	}

	for len(flashes) > 0 {
		flash := flashes[0]
		flashes = flashes[1:]
		if _, ok := flashed[flash.Coords()]; ok {
			// Already flashed
			continue
		}

		flashed[flash.Coords()] = true
		for _, neighbor := range flash.Neighbours8() {
			if neighbor.X < 0 || neighbor.Y < 0 || neighbor.X >= len(grid[0]) || neighbor.Y >= len(grid) {
				continue
			}
			grid[neighbor.Y][neighbor.X]++
			if grid[neighbor.Y][neighbor.X] > 9 && !flashed[neighbor.Coords()] {
				flashes = append(flashes, neighbor)
			}
		}
	}

	for y, row := range grid {
		for x, cell := range row {
			if cell > 9 {
				grid[y][x] = 0
			}
		}
	}

	return grid, len(flashed)
}
