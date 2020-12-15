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
		1:  {s: Day01, part1: "224436", part2: "303394260"},
		2:  {s: Day02, part1: "556", part2: "605"},
		3:  {s: Day03, part1: "145", part2: "3424528800"},
		4:  {s: Day04, part1: "230", part2: "156"},
		5:  {s: Day05, part1: "965", part2: "524"},
		6:  {s: Day06, part1: "6809", part2: "3394"},
		7:  {s: Day07, part1: "192", part2: "12128"},
		8:  {s: Day08, part1: "1134", part2: "1205"},
		9:  {s: Day09, part1: "3199139634", part2: "438559930"},
		10: {s: Day10, part1: "2310", part2: "64793042714624"},
		11: {s: Day11, part1: "2354", part2: "2072"},
		12: {s: Day12, part1: "1838", part2: "89936"},
		13: {s: Day13, part1: "3215", part2: "1001569619313439"},
		14: {s: Day14, part1: "8566770985168", part2: ""},
		15: {s: Day15, part1: "376", part2: "323780"},
		// 16: {s: Day16, part1: "", part2: ""},
		// 17: {s: Day17, part1: "", part2: ""},
		// 18: {s: Day18, part1: "", part2: ""},
		// 19: {s: Day19, part1: "", part2: ""},
		// 20: {s: Day20, part1: "", part2: ""},
		// 21: {s: Day21, part1: "", part2: ""},
		// 22: {s: Day22, part1: "", part2: ""},
		// 23: {s: Day23, part1: "", part2: ""},
		// 24: {s: Day24, part1: "", part2: ""},
		// 25: {s: Day25, part1: "", part2: ""},
	}

	for d, s := range tests {
		t.Run(fmt.Sprintf("Day%02d", d), func(t *testing.T) {
			day := day.TestDay(t, s.s)
			day.WithFile(fmt.Sprintf("../../input/aoc2020/day%02d.in", d), s.part1, s.part2)
		})
	}
}
