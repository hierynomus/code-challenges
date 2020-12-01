package aoc2018

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day05(reader *bufio.Scanner) (string, string) {
	if !reader.Scan() {
		panic("No input")
	}

	input := reader.Text()

	part1 := len(fold(input))

	part2 := len(input)
	for i := 0; i < 26; i++ {
		filtered := []rune{}

		for _, c := range input {
			if c == rune(97+i) || c == rune(65+i) {
				continue
			}
			filtered = append(filtered, c)
		}

		l := len(fold(string(filtered)))
		if part2 > l {
			part2 = l
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func fold(inp string) []int {
	reacted := []int{}
	r := -1

	for i := 0; i < len(inp); i++ {
		n := int(inp[i])
		if r < 0 || aoc.Abs(n-reacted[r]) != 32 {
			reacted = append(reacted, n)
			r += 1
		} else {
			reacted = reacted[:len(reacted)-1]
			r -= 1
		}
	}

	return reacted
}
