package aoc2023

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type ScratchCard struct {
	Nr      int
	Winning aoc.IntSet
	Numbers []int
}

func Day04(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	cards := []ScratchCard{}
	for _, l := range aoc.ReadStringArray(reader) {
		cards = append(cards, parseScratchCard(l))
	}

	copies := map[int]int{}
	for _, c := range cards {
		copies[c.Nr] = 1
	}

	for _, c := range cards {
		value := 0
		matches := 0
		for _, n := range c.Numbers {
			if c.Winning.Contains(n) {
				matches++
				if value == 0 {
					value = 1
				} else {
					value *= 2
				}
			}
		}

		part1 += value

		for x := 1; x <= matches; x++ {
			copies[c.Nr+x] += copies[c.Nr] // Add as many winning copies as there are copies of this card
		}
	}

	for _, c := range copies {
		part2 += c
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func parseScratchCard(l string) ScratchCard {
	card := ScratchCard{}
	parts := strings.Split(l, ":")
	card.Nr = aoc.ToInt(strings.TrimSpace(parts[0][4:]))
	nrSets := strings.Split(parts[1], " | ")
	card.Winning = aoc.NewIntSet(aoc.AsIntArraySpace(nrSets[0]))
	card.Numbers = aoc.AsIntArraySpace(nrSets[1])
	return card
}
