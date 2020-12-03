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
		6: {s: Day06, part1: "4186", part2: "45509"},
		7: {s: Day07, part1: "ABGKCMVWYDEHFOPQUILSTNZRJX", part2: "898"},
		8: {s: Day08, part1: "40746", part2: "37453"},
		9: {s: Day09, part1: "404611", part2: "3350093681"},
	}

	for d, s := range tests {
		t.Run(fmt.Sprintf("Day%02d", d), func(t *testing.T) {
			day := day.TestDay(t, s.s)
			day.WithFile(fmt.Sprintf("../../input/aoc2018/day%02d.in", d), s.part1, s.part2)
		})
	}
}
