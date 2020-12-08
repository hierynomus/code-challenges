package aoc2020

import (
	"fmt"
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDays(t *testing.T) {
	tests := map[int]struct {
		s     day.Solver
		part1 string
		part2 string
	}{
		1: {s: Day01, part1: "224436", part2: "303394260"},
		2: {s: Day02, part1: "556", part2: "605"},
		3: {s: Day03, part1: "145", part2: "3424528800"},
		4: {s: Day04, part1: "230", part2: "156"},
		5: {s: Day05, part1: "965", part2: "524"},
		6: {s: Day06, part1: "6809", part2: "3394"},
		7: {s: Day07, part1: "192", part2: "12128"},
		8: {s: Day08, part1: "1134", part2: "1205"},
	}

	for d, s := range tests {
		t.Run(fmt.Sprintf("Day%02d", d), func(t *testing.T) {
			day := day.TestDay(t, s.s)
			day.WithFile(fmt.Sprintf("../../input/aoc2020/day%02d.in", d), s.part1, s.part2)
		})
	}
}
