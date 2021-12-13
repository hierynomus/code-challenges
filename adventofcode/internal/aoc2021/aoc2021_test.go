package aoc2021

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
		1:  {s: Day01, part1: "1387", part2: "1362"},
		2:  {s: Day02, part1: "1636725", part2: "1872757425"},
		3:  {s: Day03, part1: "4138664", part2: "4273224"},
		4:  {s: Day04, part1: "28082", part2: "8224"},
		5:  {s: Day05, part1: "7674", part2: "20898"},
		6:  {s: Day06, part1: "360268", part2: "1632146183902"},
		7:  {s: Day07, part1: "359648", part2: "100727924"},
		8:  {s: Day08, part1: "387", part2: "986034"},
		9:  {s: Day09, part1: "541", part2: "847504"},
		10: {s: Day10, part1: "315693", part2: "1870887234"},
		11: {s: Day11, part1: "1743", part2: "364"},
		12: {s: Day12, part1: "3421", part2: "84870"},
		13: {s: Day13, part1: "708", part2: `#### ###  #    #  # ###  ###  #### #  #
#    #  # #    #  # #  # #  # #    #  #
###  ###  #    #  # ###  #  # ###  ####
#    #  # #    #  # #  # ###  #    #  #
#    #  # #    #  # #  # # #  #    #  #
#### ###  ####  ##  ###  #  # #    #  #
`},
		// 14: {s: Day14, part1: "8566770985168", part2: ""},
		// 15: {s: Day15, part1: "376", part2: "323780"},
		// 16: {s: Day16, part1: "27870", part2: "3173135507987"},
		// 17: {s: Day17, part1: "223", part2: "1884"},
		// 18: {s: Day18, part1: "3885386961962", part2: "112899558798666"},
		// 19: {s: Day19, part1: "113", part2: "253"},
		// // 20: {s: Day20, part1: "", part2: ""},
		// // 21: {s: Day21, part1: "", part2: ""},
		// 22: {s: Day22, part1: "33010", part2: "32769"},
		// 23: {s: Day23, part1: "89372645", part2: "21273394210"},
		// 24: {s: Day24, part1: "351", part2: "3869"},
		// 25: {s: Day25, part1: "", part2: ""},
	}

	for d, s := range tests {
		t.Run(fmt.Sprintf("Day%02d", d), func(t *testing.T) {
			day := day.TestDay(t, s.s)
			day.WithFile(fmt.Sprintf("../../input/aoc2021/day%02d.in", d), s.part1, s.part2)
		})
	}
}
