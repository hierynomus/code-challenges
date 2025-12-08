package aoc2025

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day06(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)
	part1 = d6p1(lines)
	part2 = d6p2(lines)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func d6p1(lines []string) int {
	sum := 0
	grid := [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Fields(line))
	}

	for x := 0; x < len(grid[0]); x++ {
		f := grid[len(grid)-1][x]
		s := aoc.ToInt(grid[0][x])
		for y := 1; y < len(grid)-1; y++ {
			v := aoc.ToInt(grid[y][x])
			switch f {
			case "*":
				s *= v
			case "+":
				s += v
			}
		}

		sum += s
	}
	return sum
}

func d6p2(lines []string) int {
	sum := 0
	fs := strings.Fields(lines[len(lines)-1])
	fidx := 0
	s := 0
	for x := 0; x < len(lines[0]); x++ {
		i := 0
		for y := 0; y < len(lines)-1; y++ {
			c := string(lines[y][x])
			if c == " " {
				continue
			}
			i = i*10 + aoc.ToInt(c)
		}

		if i == 0 {
			fidx++
			sum += s
			s = 0
			continue
		}

		switch fs[fidx] {
		case "*":
			if s == 0 {
				s = 1
			}
			s *= i
		case "+":
			s += i
		}
	}
	sum += s
	return sum
}
