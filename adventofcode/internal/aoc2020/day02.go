package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day02(reader *bufio.Scanner) (string, string) {
	passwords := []struct {
		Min      int
		Max      int
		Rune     rune
		Password []rune
	}{}

	for reader.Scan() {
		l := reader.Text()
		split := strings.Split(l, " ")
		mm := strings.Split(split[0], "-")

		passwords = append(passwords, struct {
			Min      int
			Max      int
			Rune     rune
			Password []rune
		}{
			Min:      aoc.ToInt(mm[0]),
			Max:      aoc.ToInt(mm[1]),
			Rune:     rune(split[1][0]),
			Password: []rune(split[2]),
		})
	}

	part1 := 0
	for _, p := range passwords {
		h := aoc.MakeRuneHistogram(p.Password)
		if p.Min <= h[p.Rune] && h[p.Rune] <= p.Max {
			part1 += 1
		}
	}

	part2 := 0
	for _, p := range passwords {
		pos1 := p.Password[p.Min-1] == p.Rune
		pos2 := p.Password[p.Max-1] == p.Rune

		if (pos1 || pos2) && !(pos1 && pos2) {
			part2 += 1
		}
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
