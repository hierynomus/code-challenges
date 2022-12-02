package aoc2019

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
		1: {S: Day01, Part1: "3388015", Part2: "5079140"},
		2: {S: Day02, Part1: "4138687", Part2: "6635"},
		3: {S: Day03, Part1: "446", Part2: "9006"},
		4: {S: Day04, Part1: "945", Part2: "617"},
		// 5: {s: Day05, part1: "15426686", part2: "11430197"},
		6: {S: Day06, Part1: "117672", Part2: "277"},
		7: {S: Day07, Part1: "929800", Part2: "15432220"},
		8: {S: Day08, Part1: "1485", Part2: `
XXX..X.....XX..X..X.XXXX.
X..X.X....X..X.X.X..X....
X..X.X....X..X.XX...XXX..
XXX..X....XXXX.X.X..X....
X.X..X....X..X.X.X..X....
X..X.XXXX.X..X.X..X.X....
`},
		9:  {S: Day09, Part1: "2427443564", Part2: "87221"},
		10: {S: Day10, Part1: "263", Part2: "1110"},
		11: {S: Day11, Part1: "2172", Part2: `
...##.####.#....####.####..##..#..#.###....
....#.#....#....#....#....#..#.#..#.#..#...
....#.###..#....###..###..#....####.#..#...
....#.#....#....#....#....#.##.#..#.###....
.#..#.#....#....#....#....#..#.#..#.#......
..##..####.####.####.#.....###.#..#.#......
`},
		12: {S: Day12, Part1: "7202", Part2: ""},
		13: {S: Day13, Part1: "361", Part2: "17590"},
		// 14: {s: Day14, part1: "337862", part2: ""},
		15: {S: Day15, Part1: "258", Part2: "372"},
		16: {S: Day16, Part1: "90744714", Part2: ""},
		// 17: {s: Day17, part1: "3888", part2: "927809"},
		// 18: {s: Day18, part1: "90744714", part2: ""},
		19: {S: Day19, Part1: "234", Part2: "9290812"},
		// 20: {s: Day20, part1: "90744714", part2: ""},
		21: {S: Day21, Part1: "19357335", Part2: "1140147758"},
		22: {S: Day22, Part1: "7096", Part2: ""},
		// 23: {s: Day23, part1: "24106", part2: ""},
		24: {S: Day24, Part1: "18370591", Part2: ""},
		// 25: {s: Day25, part1: "", part2: ""},
	}

	day.RunDays(t, tests)
}
