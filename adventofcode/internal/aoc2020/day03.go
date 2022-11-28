package aoc2020

import (
	"bufio"
	"strconv"
)

func Day03(reader *bufio.Scanner) (string, string) {
	lines := []string{}
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}

	p := 1
	for _, dxy := range [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	} {
		p *= traverseForest(lines, dxy[0], dxy[1])
	}

	return strconv.Itoa(traverseForest(lines, 3, 1)), strconv.Itoa(p)
}

func traverseForest(forest []string, dx, dy int) int {
	x, y := 0, 0
	s := 0
	for y < len(forest) {
		if forest[y][x] == '#' {
			s += 1
		}

		x = (x + dx) % len(forest[y])
		y += dy
	}

	return s
}
