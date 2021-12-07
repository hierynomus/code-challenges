package aoc2021

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day06(reader *bufio.Scanner) (string, string) {
	evolve := func(fish []int) {
		t := fish[0]
		for j := 1; j < 9; j++ {
			fish[j-1] = fish[j]
		}
		fish[8] = t
		fish[6] += t
	}

	var part1, part2 int

	fish := make([]int, 9)
	inp := aoc.AsIntArray(aoc.Read(reader))
	for _, v := range inp {
		fish[v]++
	}

	for i := 0; i < 80; i++ {
		evolve(fish)
	}

	part1 = aoc.Sum(fish)

	for i := 0; i < 256-80; i++ {
		evolve(fish)
	}

	part2 = aoc.Sum(fish)
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
