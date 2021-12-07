package aoc2021

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day07(reader *bufio.Scanner) (string, string) {
	fuelNeeded := func(crabs []int, pos int) int {
		x := 0
		for _, c := range crabs {
			x += aoc.Abs(c - pos)
		}
		return x
	}

	fuelNeeded2 := func(crabs []int, pos int) int {
		x := 0
		for _, c := range crabs {
			n := aoc.Abs(c - pos)
			x += n + (n*(n-1))/2
		}
		return x
	}

	var part1, part2 int

	crabs := aoc.AsIntArray(aoc.Read(reader))

	part1 = fuelNeeded(crabs, findMin(crabs, fuelNeeded))
	part2 = fuelNeeded2(crabs, findMin(crabs, fuelNeeded2))

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func findMin(crabs []int, fuelFunc func([]int, int) int) int {
	left := aoc.Min(crabs)
	right := aoc.Max(crabs)

	for left < right {
		mean := (left + right) / 2
		if fuelFunc(crabs, mean) > fuelFunc(crabs, mean+1) {
			left = mean + 1
		} else {
			right = mean
		}
	}

	return left
}
