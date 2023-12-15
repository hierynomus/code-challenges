package aoc2023

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day14(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	grid := aoc.ReadRuneGrid(reader)
	maxWeight := len(grid)
	w := 0
	// Iterate over the grid transposed
	for x := 0; x < len(grid[0]); x++ {
		blockedAt := 0
		for y := 0; y < len(grid); y++ {
			switch grid[y][x] {
			case '#':
				blockedAt = y + 1
			case '.':
				continue
			case 'O':
				w += (maxWeight - blockedAt)
				blockedAt++
			}
		}
	}

	part1 = w

	// Damn, for part2 we need to do the simulation
	// As we start north, first rotate CW, then move boulders, repeat 3 times for all directions, then calculate hash
	// We need to find the cycle, and then calculate the offset
	hashes := map[string]int{}
	weights := []int{}
	rotationCycles := 0
	nxt := grid
	hash := aoc.HashRuneGrid(nxt)
	for _, ok := hashes[hash]; !ok; _, ok = hashes[hash] { // while not in hashes
		w := SupportBeamWeight(nxt)
		hashes[hash] = rotationCycles
		weights = append(weights, w)
		for i := 0; i < 4; i++ { // do 4 times
			nxt = aoc.RotateCW(nxt)
			for x := 0; x < len(nxt); x++ {
				MoveBouldersLine(nxt[x])
			}
		}

		rotationCycles++

		hash = aoc.HashRuneGrid(nxt)
	}

	cycleLen := rotationCycles - hashes[hash]
	offset := (1000000000 - hashes[hash]) % cycleLen
	part2 = weights[hashes[hash]+offset]

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

// This func will move the boulders packed to the right on the line
func MoveBouldersLine(line []rune) {
	blockedPos := len(line)
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == '#' {
			blockedPos = i
		} else if line[i] == 'O' {
			line[i] = '.'
			line[blockedPos-1] = 'O'
			blockedPos--
		}
	}
}

func SupportBeamWeight(grid [][]rune) int {
	w := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'O' {
				w += len(grid) - y
			}
		}
	}

	return w
}
