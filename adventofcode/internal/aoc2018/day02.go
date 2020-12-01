package aoc2018

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day02(reader *bufio.Scanner) (string, string) {
	boxes := []string{}
	for reader.Scan() {
		boxes = append(boxes, reader.Text())
	}

	twos, threes := 0, 0
	for _, b := range boxes {
		h := aoc.MakeRuneHistogram([]rune(b))
		two, three := false, false

		for _, v := range h {
			two = two || v == 2
			three = three || v == 3
		}

		if two {
			twos += 1
		}

		if three {
			threes += 1
		}
	}

	part1 := twos * threes

	part2 := make([]rune, 0)
	for _, comb := range aoc.StringCombinations(boxes) {
		if aoc.HammingDistance(comb[0], comb[1]) == 1 {
			for i, c := range comb[0] {
				if c == rune(comb[1][i]) {
					part2 = append(part2, c)
				}
			}
		}
	}

	return strconv.Itoa(part1), string(part2)
}
