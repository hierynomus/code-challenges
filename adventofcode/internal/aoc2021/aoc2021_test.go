package aoc2021

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
		1:  {S: Day01, Part1: "1387", Part2: "1362"},
		2:  {S: Day02, Part1: "1636725", Part2: "1872757425"},
		3:  {S: Day03, Part1: "4138664", Part2: "4273224"},
		4:  {S: Day04, Part1: "28082", Part2: "8224"},
		5:  {S: Day05, Part1: "7674", Part2: "20898"},
		6:  {S: Day06, Part1: "360268", Part2: "1632146183902"},
		7:  {S: Day07, Part1: "359648", Part2: "100727924"},
		8:  {S: Day08, Part1: "387", Part2: "986034"},
		9:  {S: Day09, Part1: "541", Part2: "847504"},
		10: {S: Day10, Part1: "315693", Part2: "1870887234"},
		11: {S: Day11, Part1: "1743", Part2: "364"},
		12: {S: Day12, Part1: "3421", Part2: "84870"},
		13: {S: Day13, Part1: "708", Part2: `#### ###  #    #  # ###  ###  #### #  #
#    #  # #    #  # #  # #  # #    #  #
###  ###  #    #  # ###  #  # ###  ####
#    #  # #    #  # #  # ###  #    #  #
#    #  # #    #  # #  # # #  #    #  #
#### ###  ####  ##  ###  #  # #    #  #
`},
		14: {S: Day14, Part1: "2602", Part2: "2942885922173"},
		15: {S: Day15, Part1: "390", Part2: "2814"},
		16: {S: Day16, Part1: "989", Part2: "7936430475134"},
		17: {S: Day17, Part1: "6786", Part2: "2313"},
		18: {S: Day18, Part1: "4347", Part2: "4721"},
		// 19: {s: Day19, part1: "113", part2: "253"},
		// // 20: {s: Day20, part1: "", part2: ""},
		21: {S: Day21, Part1: "906093", Part2: "274291038026362"},
		// 22: {s: Day22, part1: "33010", part2: ""},
		// 23: {s: Day23, part1: "89372645", part2: "21273394210"},
		// 24: {s: Day24, part1: "351", part2: "3869"},
		// 25: {s: Day25, part1: "", part2: ""},
	}

	day.RunDays(t, tests)
}
