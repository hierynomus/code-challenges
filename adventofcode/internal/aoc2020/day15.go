package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day15(reader *bufio.Scanner) (string, string) {
	if !reader.Scan() {
		panic("No input")
	}

	input := aoc.AsIntArray(reader.Text())

	part1 := playMemoryGame(input, 2020)
	part2 := playMemoryGame(input, 30000000)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func playMemoryGame(input []int, rounds int) int {
	game1 := map[int]int{}
	for i, x := range input {
		if i != len(input)-1 {
			game1[x] = i
		}
	}
	last := input[len(input)-1]

	for i := len(input); i < rounds; i++ {
		if x, ok := game1[last]; ok {
			game1[last] = i - 1
			last = i - 1 - x
		} else {
			game1[last] = i - 1
			last = 0
		}
	}

	return last
}
