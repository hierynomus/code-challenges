package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day11(reader *bufio.Scanner) (string, string) {
	grid := aoc.ReadRuneGrid(reader)

	newGrid := grid
	evolved := true
	for evolved {
		newGrid, evolved = evolve(newGrid, 4, countOccupied1)
	}

	part1 := aoc.CountRuneGridOccurrences(newGrid, '#')

	newGrid = grid
	evolved = true
	for evolved {
		newGrid, evolved = evolve(newGrid, 5, countOccupied2)
	}

	part2 := aoc.CountRuneGridOccurrences(newGrid, '#')

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func evolve(grid [][]rune, tolerance int, occupancyCount func([][]rune, int, int) int) ([][]rune, bool) {
	newGrid := make([][]rune, len(grid))
	evolved := false
	for y := 0; y < len(grid); y++ {
		line := grid[y]
		newLine := make([]rune, len(line))
		for x := 0; x < len(line); x++ {
			occ := occupancyCount(grid, x, y)
			switch {
			case line[x] == 'L' && occ == 0:
				newLine[x] = '#'
				evolved = true
			case line[x] == '#' && occ >= tolerance:
				newLine[x] = 'L'
				evolved = true
			default:
				newLine[x] = line[x]
			}
		}
		newGrid[y] = newLine
	}

	return newGrid, evolved
}

func countOccupied1(grid [][]rune, x, y int) int {
	occupied := 0
	for _, p := range aoc.Neighbours8(x, y) {
		if p.Y < 0 || p.Y >= len(grid) {
			continue
		}
		if p.X < 0 || p.X >= len(grid[p.Y]) {
			continue
		}

		if grid[p.Y][p.X] == '#' {
			occupied += 1
		}
	}

	return occupied
}

func countOccupied2(grid [][]rune, x, y int) int {
	occupied := 0
	for _, d := range aoc.DeltaNeighbours8() {
		dx, dy := d.X, d.Y

		for y+dy >= 0 && y+dy < len(grid) && x+dx >= 0 && x+dx < len(grid[y+dy]) {
			if grid[y+dy][x+dx] == '.' {
				dx += d.X
				dy += d.Y
			} else {
				break
			}
		}

		// Bounds check
		if y+dy >= 0 && y+dy < len(grid) && x+dx >= 0 && x+dx < len(grid[y+dy]) {
			if grid[y+dy][x+dx] == '#' {
				occupied += 1
			}
		}
	}

	return occupied
}
