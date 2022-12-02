package aoc2020

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDays(t *testing.T) {
	tests := map[int]struct {
		S     day.Solver
		Part1 string
		Part2 string
	}{
		1:  {S: Day01, Part1: "224436", Part2: "303394260"},
		2:  {S: Day02, Part1: "556", Part2: "605"},
		3:  {S: Day03, Part1: "145", Part2: "3424528800"},
		4:  {S: Day04, Part1: "230", Part2: "156"},
		5:  {S: Day05, Part1: "965", Part2: "524"},
		6:  {S: Day06, Part1: "6809", Part2: "3394"},
		7:  {S: Day07, Part1: "192", Part2: "12128"},
		8:  {S: Day08, Part1: "1134", Part2: "1205"},
		9:  {S: Day09, Part1: "3199139634", Part2: "438559930"},
		10: {S: Day10, Part1: "2310", Part2: "64793042714624"},
		11: {S: Day11, Part1: "2354", Part2: "2072"},
		12: {S: Day12, Part1: "1838", Part2: "89936"},
		13: {S: Day13, Part1: "3215", Part2: "1001569619313439"},
		14: {S: Day14, Part1: "8566770985168", Part2: ""},
		15: {S: Day15, Part1: "376", Part2: "323780"},
		16: {S: Day16, Part1: "27870", Part2: "3173135507987"},
		17: {S: Day17, Part1: "223", Part2: "1884"},
		18: {S: Day18, Part1: "3885386961962", Part2: "112899558798666"},
		19: {S: Day19, Part1: "113", Part2: "253"},
		// 20: {s: Day20, part1: "", part2: ""},
		// 21: {s: Day21, part1: "", part2: ""},
		22: {S: Day22, Part1: "33010", Part2: "32769"},
		23: {S: Day23, Part1: "89372645", Part2: "21273394210"},
		24: {S: Day24, Part1: "351", Part2: "3869"},
		// 25: {s: Day25, part1: "", part2: ""},
	}

	day.RunDays(t, tests)
}
