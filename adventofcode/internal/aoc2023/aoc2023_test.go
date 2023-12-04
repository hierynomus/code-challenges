package aoc2023

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
		1: {S: Day01, Part1: "67016", Part2: "200116"},
		// 		2: {S: Day02, Part1: "14264", Part2: "12382"},
		// 		3: {S: Day03, Part1: "8394", Part2: "2413"},
		// 		4: {S: Day04, Part1: "571", Part2: "917"},
		// 		5: {S: Day05, Part1: "DHBJQJCCW", Part2: "WJVRLSJJT"},
		// 		6: {S: Day06, Part1: "1651", Part2: "3837"},
		// 		7: {S: Day07, Part1: "1644735", Part2: "1300850"},
		// 		8: {S: Day08, Part1: "1845", Part2: "230112"},
		// 		9: {S: Day09, Part1: "6266", Part2: "2369"},
		// 		10: {S: Day10, Part1: "17840", Part2: `XXXX..XX..X.....XX..X..X.X....XXX...XX..
		// X....X..X.X....X..X.X..X.X....X..X.X..X.
		// XXX..X..X.X....X....X..X.X....X..X.X....
		// X....XXXX.X....X.XX.X..X.X....XXX..X.XX.
		// X....X..X.X....X..X.X..X.X....X....X..X.
		// XXXX.X..X.XXXX..XXX..XX..XXXX.X.....XXX.`},
		// 		11: {S: Day11, Part1: "113232", Part2: "29703395016"},
		// 		12: {S: Day12, Part1: "462", Part2: "451"},
		// 		13: {S: Day13, Part1: "4894", Part2: "24180"},
		// 		14: {S: Day14, Part1: "979", Part2: "29044"},
		// 		15: {S: Day15, Part1: "4737443", Part2: "11482462818989"},
		// 		// 		// 16: {S: Day16, Part1: "", Part2: ""},
		// 		// 		17: {S: Day17, Part1: "3144", Part2: ""},
		// 		18: {S: Day18, Part1: "4482", Part2: "2576"},
		// 		// 19: {S: Day19, Part1: "", Part2: ""},
		// 		20: {S: Day20, Part1: "2275", Part2: "4090409331120"},
		// 		// 21: {S: Day21, Part1: "158731561459602", Part2: ""},
		// 		// 22: {S: Day22, Part1: "", Part2: ""},
		// 		23: {S: Day23, Part1: "3815", Part2: "893"},
		// 		// 24: {S: Day24, Part1: "", Part2: ""},
		// 		25: {S: Day25, Part1: "", Part2: ""},
	}

	day.RunDays(t, tests)
}
