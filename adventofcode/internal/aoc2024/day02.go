package aoc2024

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type ReindeerReport []int

func (r ReindeerReport) Safe() bool {
	dec := r[0] > r[1]
	for i := 1; i < len(r); i++ {
		d := aoc.Abs(r[i] - r[i-1])
		if 1 > d || d > 3 {
			return false
		}

		if (r[i-1] > r[i]) && !dec {
			return false
		} else if (r[i-1] < r[i]) && dec {
			return false
		} else if r[i-1] == r[i] {
			return false
		}
	}

	return true
}

func (r ReindeerReport) DampenedSafe() bool {
	if r.Safe() {
		return true
	}

	for i := 0; i < len(r); i++ {
		rep := make(ReindeerReport, 0)
		rep = append(rep, r[:i]...)
		rep = append(rep, r[i+1:]...)
		if rep.Safe() {
			return true
		}
	}

	return false
}

func Day02(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	g := aoc.ReadIntGrid(reader, " ")
	reports := make([]ReindeerReport, len(g))
	for i, line := range g {
		reports[i] = ReindeerReport(line)
	}

	for _, r := range reports {
		if r.Safe() {
			part1++
			part2++
		} else if r.DampenedSafe() {
			part2++
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
