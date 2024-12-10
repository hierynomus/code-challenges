package aoc2024

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
		1: {S: Day01, Part1: "1603498", Part2: "25574739"},
		2: {S: Day02, Part1: "670", Part2: "700"},
		3: {S: Day03, Part1: "184122457", Part2: "107862689"},
		4: {S: Day04, Part1: "2483", Part2: "1925"},
		5: {S: Day05, Part1: "5732", Part2: "4716"},
		// 6: {S: Day06, Part1: "4973", Part2: "28228952"},
		7: {S: Day07, Part1: "303876485655", Part2: "146111650210682"},
		8: {S: Day08, Part1: "295", Part2: "1034"},
		9: {S: Day09, Part1: "6320029754031", Part2: "6347435485773"},
		// 10: {S: Day10, Part1: "6768", Part2: "351"},
		// 11: {S: Day11, Part1: "9769724", Part2: "603020563700"},
		// // 		12: {S: Day12, Part1: "462", Part2: "451"},
		// 13: {S: Day13, Part1: "28651", Part2: "25450"},
		// 14: {S: Day14, Part1: "105784", Part2: "91286"},
		// 15: {S: Day15, Part1: "506891", Part2: "230462"},
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
