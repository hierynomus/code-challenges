package aoc2019

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
		1: {s: Day01, part1: "3388015", part2: "5079140"},
		2: {s: Day02, part1: "4138687", part2: "6635"},
		3: {s: Day03, part1: "446", part2: "9006"},
		4: {s: Day04, part1: "945", part2: "617"},
		// 5: {s: Day05, part1: "15426686", part2: "11430197"},
		6: {s: Day06, part1: "117672", part2: "277"},
		7: {s: Day07, part1: "929800", part2: "15432220"},
		8: {s: Day08, part1: "1485", part2: `
XXX..X.....XX..X..X.XXXX.
X..X.X....X..X.X.X..X....
X..X.X....X..X.XX...XXX..
XXX..X....XXXX.X.X..X....
X.X..X....X..X.X.X..X....
X..X.XXXX.X..X.X..X.X....
`},
		9:  {s: Day09, part1: "2427443564", part2: "87221"},
		10: {s: Day10, part1: "263", part2: "1110"},
		11: {s: Day11, part1: "2172", part2: `
...##.####.#....####.####..##..#..#.###....
....#.#....#....#....#....#..#.#..#.#..#...
....#.###..#....###..###..#....####.#..#...
....#.#....#....#....#....#.##.#..#.###....
.#..#.#....#....#....#....#..#.#..#.#......
..##..####.####.####.#.....###.#..#.#......
`},
		// 12: {s: Day12, part1: "7202", part2: "87221"},
		13: {s: Day13, part1: "361", part2: "17590"},
		// 14: {s: Day14, part1: "337862", part2: ""},
		15: {s: Day15, part1: "258", part2: "372"},
		// 16: {s: Day16, part1: "90744714", part2: ""},
		// 17: {s: Day17, part1: "3888", part2: "927809"},
		// 18: {s: Day18, part1: "90744714", part2: ""},
		19: {s: Day19, part1: "234", part2: "9290812"},
		// 20: {s: Day20, part1: "90744714", part2: ""},
		21: {s: Day21, part1: "19357335", part2: "1140147758"},
		22: {s: Day22, part1: "7096", part2: ""},
		// 23: {s: Day23, part1: "24106", part2: "1140147758"},
		24: {s: Day24, part1: "18370591", part2: ""},
		// 25: {s: Day25, part1: "", part2: ""},
	}

	for d, s := range tests {
		t.Run(fmt.Sprintf("Day%02d", d), func(t *testing.T) {
			day := day.TestDay(t, s.s)
			day.WithFile(fmt.Sprintf("../../input/aoc2019/day%02d.in", d), s.part1, s.part2)
		})
	}
}
