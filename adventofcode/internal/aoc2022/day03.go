package aoc2022

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func GetItemPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - int('a') + 1
	}

	return int(r) - int('A') + 27
}

func Day03(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	rucksacks := aoc.ReadStringArray(reader)

	for _, r := range rucksacks {
		l := len(r)
		s := aoc.NewRuneSet([]rune(r[0 : l/2]))
		for _, i := range r[l/2 : l] {
			if s.Contains(i) {
				part1 += GetItemPriority(i)
				break
			}
		}
	}

	for idx := 0; idx < len(rucksacks); idx += 3 {
		h := aoc.MakeRuneHistogram([]rune{})
		for i := idx; i < idx+3; i++ {
			s := aoc.NewRuneSet([]rune(rucksacks[i]))
			h.Adds(s.Keys())
		}

		badge := h.Max()
		part2 += GetItemPriority(badge)
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
