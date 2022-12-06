package aoc2022

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
		2: {S: Day02, Part1: "14264", Part2: "12382"},
		3: {S: Day03, Part1: "8394", Part2: "2413"},
		4: {S: Day04, Part1: "571", Part2: "917"},
		5: {S: Day05, Part1: "DHBJQJCCW", Part2: "WJVRLSJJT"},
		6: {S: Day06, Part1: "1651", Part2: "3837"},
		// 7:  {S: Day07, Part1: "", Part2: ""},
		// 8:  {S: Day08, Part1: "", Part2: ""},
		// 9:  {S: Day09, Part1: "", Part2: ""},
		// 10: {S: Day10, Part1: "", Part2: ""},
		// 11: {S: Day11, Part1: "", Part2: ""},
		// 12: {S: Day12, Part1: "", Part2: ""},
		// 13: {S: Day13, Part1: "", Part2: ``},
		// 14: {S: Day14, Part1: "", Part2: ""},
		// 15: {S: Day15, Part1: "", Part2: ""},
		// 16: {S: Day16, Part1: "", Part2: ""},
		// 17: {S: Day17, Part1: "", Part2: ""},
		// 18: {S: Day18, Part1: "", Part2: ""},
		// 19: {S: Day19, Part1: "", Part2: ""},
		// 20: {S: Day20, Part1: "", Part2: ""},
		// 21: {S: Day21, Part1: "", Part2: ""},
		// 22: {S: Day22, Part1: "", Part2: ""},
		// 23: {S: Day23, Part1: "", Part2: ""},
		// 24: {S: Day24, Part1: "", Part2: ""},
		// 25: {S: Day25, Part1: "", Part2: ""},
	}

	day.RunDays(t, tests)
}
