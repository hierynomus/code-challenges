package aoc2022

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type ElfPairing struct {
	elf1, elf2 CleanAssignment
}
type CleanAssignment struct {
	from, to int
}

func (c CleanAssignment) FullyContains(o CleanAssignment) bool {
	return c.from <= o.from && c.to >= o.from && c.from <= o.to && c.to >= o.to
}

func (c CleanAssignment) Contains(o CleanAssignment) bool {
	return (c.from <= o.from && c.to >= o.from) || (c.from <= o.to && c.to >= o.to)
}

func Day04(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	pairings := []ElfPairing{}
	for _, l := range aoc.ReadStringArray(reader) {
		s := strings.Split(l, ",")
		e1, e2 := strings.Split(s[0], "-"), strings.Split(s[1], "-")
		pairings = append(pairings, ElfPairing{
			elf1: CleanAssignment{
				from: aoc.ToInt(e1[0]),
				to:   aoc.ToInt(e1[1]),
			},
			elf2: CleanAssignment{
				from: aoc.ToInt(e2[0]),
				to:   aoc.ToInt(e2[1]),
			},
		})
	}

	for _, pair := range pairings {
		if pair.elf1.FullyContains(pair.elf2) || pair.elf2.FullyContains(pair.elf1) {
			part1++
		}
		if pair.elf1.Contains(pair.elf2) || pair.elf2.Contains(pair.elf1) {
			part2++
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
