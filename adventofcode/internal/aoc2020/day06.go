package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day06(reader *bufio.Scanner) (string, string) {
	anyone := []aoc.RuneSet{}
	everyone := []aoc.RuneSet{}
	orGroup := aoc.RuneSet{}
	andGroup := aoc.NewRuneSet([]rune("abcdefghijklmnopqrstuvwxyz"))

	for reader.Scan() {
		l := reader.Text()
		if len(l) == 0 {
			anyone = append(anyone, orGroup)
			everyone = append(everyone, andGroup)
			orGroup = aoc.RuneSet{}
			andGroup = aoc.NewRuneSet([]rune("abcdefghijklmnopqrstuvwxyz"))
			continue
		}

		orGroup.Adds([]rune(l))
		andGroup.Intersect([]rune(l))
	}

	anyone = append(anyone, orGroup)
	everyone = append(everyone, andGroup)

	part1 := 0
	for _, g := range anyone {
		part1 += len(g)
	}

	part2 := 0
	for _, g := range everyone {
		part2 += len(g)
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
