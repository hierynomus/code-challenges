package aoc2018

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
		1: {s: Day01, part1: "402", part2: "481"},
		2: {s: Day02, part1: "5166", part2: "cypueihajytordkgzxfqplbwn"},
		3: {s: Day03, part1: "105071", part2: "222"},
		4: {s: Day04, part1: "39422", part2: "65474"},
		5: {s: Day05, part1: "11118", part2: "6948"},
	}

	for d, s := range tests {
		t.Run(fmt.Sprintf("Day%02d", d), func(t *testing.T) {
			day := day.TestDay(t, s.s)
			day.WithFile(fmt.Sprintf("../../input/aoc2018/day%02d.in", d), s.part1, s.part2)
		})
	}
}
